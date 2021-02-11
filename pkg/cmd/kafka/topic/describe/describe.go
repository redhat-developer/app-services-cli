package describe

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil"

	"github.com/MakeNowJust/heredoc"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/color"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/flag"
	flagutil "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil/flags"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/kas"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/dump"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/iostreams"
	"gopkg.in/yaml.v2"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"

	"github.com/spf13/cobra"
)

type Options struct {
	topicName    string
	kafkaID      string
	outputFormat string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection func() (connection.Connection, error)
	Logger     func() (logging.Logger, error)
}

// NewDescribeTopicCommand gets a new command for describing a kafka topic.
func NewDescribeTopicCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Connection: f.Connection,
		Config:     f.Config,
		Logger:     f.Logger,
		IO:         f.IOStreams,
	}

	cmd := &cobra.Command{
		Use:   "describe",
		Short: "Describe a Kafka topic",
		Long:  "Print detailed configuration information for a Kafka topic",
		Example: heredoc.Doc(`
			# describe Kafka topic "topic-1"
			$ rhoas kafka describe topic-1
		`),
		Args: cobra.ExactArgs(1),
		// dynamic completion of topic names
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
			if len(args) > 0 {
				opts.topicName = args[0]
			}

			if err = flag.ValidateOutput(opts.outputFormat); err != nil {
				return err
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

	fs := cmd.Flags()
	flag.AddOutput(fs, &opts.outputFormat, "json", flagutil.ValidOutputFormats)

	return cmd
}

func runCmd(opts *Options) error {
	conn, err := opts.Connection()
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

	// fetch the topic
	topicResponse, httpRes, topicErr := api.TopicAdmin(opts.kafkaID).
		GetTopic(ctx, opts.topicName).
		Execute()

	if topicErr.Error() != "" {
		switch httpRes.StatusCode {
		case 404:
			return fmt.Errorf("topic '%v' not found in Kafka instance '%v'", opts.topicName, kafkaInstance.GetName())
		case 401:
			return fmt.Errorf("you are unauthorized to view this topic")
		case 500:
			return fmt.Errorf("internal server error: %w", topicErr)
		case 503:
			return fmt.Errorf("unable to connect to Kafka instance '%v': %w", kafkaInstance.GetName(), topicErr)
		default:
			return topicErr
		}
	}

	switch opts.outputFormat {
	case "json":
		data, _ := json.Marshal(topicResponse)
		_ = dump.JSON(opts.IO.Out, data)
	case "yaml", "yml":
		data, _ := yaml.Marshal(topicResponse)
		_ = dump.YAML(opts.IO.Out, data)
	}

	return nil
}
