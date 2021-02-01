package delete

import (
	"context"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil"

	"github.com/MakeNowJust/heredoc"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/color"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/kas"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/iostreams"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"

	"github.com/spf13/cobra"
)

type Options struct {
	topicName string
	kafkaID   string
	force     bool

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection func() (connection.Connection, error)
	Logger     func() (logging.Logger, error)
}

// NewDeleteTopicCommand gets a new command for deleting a kafka topic.
func NewDeleteTopicCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Connection: f.Connection,
		Config:     f.Config,
		Logger:     f.Logger,
		IO:         f.IOStreams,
	}

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a Kafka topic",
		Long:  "Delete a topic in the current Kafka instance",
		Example: heredoc.Doc(`
			# delete Kafka topic "topic-1"
			$ rhoas kafka delete topic-1
		`),
		Args: cobra.ExactArgs(1),
		// Dynamic completion of the topic name
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			validNames := []string{}

			cfg, err := opts.Config.Load()
			if err != nil {
				return validNames, cobra.ShellCompDirectiveError
			}

			if !cfg.HasKafka() {
				return validNames, cobra.ShellCompDirectiveError
			}

			opts.kafkaID = cfg.Services.Kafka.ClusterID

			return cmdutil.FilterValidTopicNameArgs(f, opts.kafkaID, toComplete)
		},
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if !opts.IO.CanPrompt() {
				return fmt.Errorf("Cannot delete Kafka topics when not running interactively")
			}

			if len(args) > 0 {
				opts.topicName = args[0]
			}

			if opts.kafkaID != "" {
				return runCmd(opts)
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			if !cfg.HasKafka() {
				return fmt.Errorf("No Kafka instance selected. To use a Kafka instance run %v", color.CodeSnippet("rhoas kafka use"))
			}

			opts.kafkaID = cfg.Services.Kafka.ClusterID

			return runCmd(opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.force, "force", "f", false, "Skip confirmation to force delete this topic")

	return cmd
}

func runCmd(opts *Options) error {
	conn, err := opts.Connection()
	if err != nil {
		return err
	}

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	api := conn.API()
	ctx := context.Background()

	// check if the Kafka instance exists
	kafkaInstance, _, apiErr := api.Kafka().GetKafkaById(ctx, opts.kafkaID).Execute()
	if kas.IsErr(apiErr, kas.ErrorNotFound) {
		return fmt.Errorf("Kafka instance with ID '%v' not found", opts.kafkaID)
	} else if apiErr.Error() != "" {
		return apiErr
	}

	if !opts.force {
		var promptConfirmName = &survey.Input{
			Message: "Confirm the name of the topic you want to delete:",
		}

		var userConfirmedName string

		if err := survey.AskOne(promptConfirmName, &userConfirmedName); err != nil {
			return err
		}

		if userConfirmedName != opts.topicName {
			logger.Infof("The topic name entered (%v) does not match the name of the topic you tried to delete (%v)", userConfirmedName, opts.topicName)
			return nil
		}
	}

	// perform delete topic API request
	httpRes, topicErr := api.TopicAdmin(opts.kafkaID).
		DeleteTopic(ctx, opts.topicName).
		Execute()

	if topicErr.Error() != "" {
		switch httpRes.StatusCode {
		case 404:
			return fmt.Errorf("topic '%v' not found in Kafka instance '%v'", opts.topicName, kafkaInstance.GetName())
		case 401:
			return fmt.Errorf("you are unauthorized to delete this topic")
		case 500:
			return fmt.Errorf("internal server error: %w", topicErr)
		case 503:
			return fmt.Errorf("unable to connect to Kafka instance '%v': %w", kafkaInstance.GetName(), topicErr)
		default:
			return topicErr
		}
	}

	logger.Infof("Topic %v in Kafka instance %v has been deleted", color.Info(opts.topicName), color.Info(kafkaInstance.GetName()))

	return nil
}
