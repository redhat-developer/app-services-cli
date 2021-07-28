package delete

import (
	"context"
	"errors"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/pkg/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/localize"

	"github.com/redhat-developer/app-services-cli/pkg/iostreams"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/logging"

	"github.com/spf13/cobra"
)

type Options struct {
	topicName string
	kafkaID   string
	force     bool

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     func() (logging.Logger, error)
	localizer  localize.Localizer
}

// NewDeleteTopicCommand gets a new command for deleting a kafka topic.
func NewDeleteTopicCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Connection: f.Connection,
		Config:     f.Config,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
	}

	cmd := &cobra.Command{
		Use:     opts.localizer.MustLocalize("kafka.topic.delete.cmd.use"),
		Short:   opts.localizer.MustLocalize("kafka.topic.delete.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("kafka.topic.delete.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("kafka.topic.delete.cmd.example"),
		Args:    cobra.ExactValidArgs(1),
		// Dynamic completion of the topic name
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return cmdutil.FilterValidTopicNameArgs(f, toComplete)
		},
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if !opts.IO.CanPrompt() && !opts.force {
				return errors.New(opts.localizer.MustLocalize("flag.error.requiredWhenNonInteractive", localize.NewEntry("Flag", "yes")))
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
				return fmt.Errorf(opts.localizer.MustLocalize("kafka.topic.common.error.noKafkaSelected"))
			}

			opts.kafkaID = cfg.Services.Kafka.ClusterID

			return runCmd(opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.force, "yes", "y", false, opts.localizer.MustLocalize("kafka.topic.delete.flag.yes.description"))

	return cmd
}

// nolint:funlen
func runCmd(opts *Options) error {
	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	api, kafkaInstance, err := conn.API().KafkaAdmin(opts.kafkaID)
	if err != nil {
		return err
	}

	// perform delete topic API request
	_, httpRes, err := api.TopicsApi.GetTopic(context.Background(), opts.topicName).
		Execute()

	topicNameTmplPair := localize.NewEntry("TopicName", opts.topicName)
	kafkaNameTmplPair := localize.NewEntry("InstanceName", kafkaInstance.GetName())
	if err != nil {
		if httpRes == nil {
			return err
		}
		if httpRes.StatusCode == 404 {
			return errors.New(opts.localizer.MustLocalize("kafka.topic.common.error.topicNotFoundError", topicNameTmplPair, kafkaNameTmplPair))
		}
	}

	if !opts.force {
		promptConfirmName := &survey.Input{
			Message: opts.localizer.MustLocalize("kafka.topic.delete.input.name.message"),
		}
		var userConfirmedName string
		if err = survey.AskOne(promptConfirmName, &userConfirmedName); err != nil {
			return err
		}

		if userConfirmedName != opts.topicName {
			return errors.New(opts.localizer.MustLocalize("kafka.topic.delete.error.mismatchedNameConfirmation", localize.NewEntry("ConfirmedName", userConfirmedName), localize.NewEntry("ActualName", opts.topicName)))
		}
	}

	// perform delete topic API request
	httpRes, err = api.TopicsApi.DeleteTopic(context.Background(), opts.topicName).
		Execute()
	if err != nil {
		if httpRes == nil {
			return err
		}

		operationTmplPair := localize.NewEntry("Operation", "delete")
		switch httpRes.StatusCode {
		case 404:
			return errors.New(opts.localizer.MustLocalize("kafka.topic.common.error.notFoundError", topicNameTmplPair, kafkaNameTmplPair))
		case 401:
			return errors.New(opts.localizer.MustLocalize("kafka.topic.common.error.unauthorized", operationTmplPair))
		case 403:
			return errors.New(opts.localizer.MustLocalize("kafka.topic.common.error.forbidden", operationTmplPair))
		case 500:
			return errors.New(opts.localizer.MustLocalize("kafka.topic.common.error.internalServerError"))
		case 503:
			return errors.New(opts.localizer.MustLocalize("kafka.topic.common.error.unableToConnectToKafka", localize.NewEntry("Name", kafkaInstance.GetName())))
		default:
			return err
		}
	}

	logger.Info(opts.localizer.MustLocalize("kafka.topic.delete.log.info.topicDeleted", topicNameTmplPair, kafkaNameTmplPair))

	return nil
}
