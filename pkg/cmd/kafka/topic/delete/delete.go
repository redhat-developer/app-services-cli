package delete

import (
	"context"
	"net/http"

	"github.com/AlecAivazis/survey/v2"

	"github.com/redhat-developer/app-services-cli/pkg/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/localize"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/logging"

	"github.com/spf13/cobra"
)

type options struct {
	topicName string
	kafkaID   string
	force     bool

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	localizer  localize.Localizer
	Context    context.Context
}

// NewDeleteTopicCommand gets a new command for deleting a kafka topic.
func NewDeleteTopicCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Connection: f.Connection,
		Config:     f.Config,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "delete",
		Short:   opts.localizer.MustLocalize("kafka.topic.delete.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("kafka.topic.delete.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("kafka.topic.delete.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if !opts.IO.CanPrompt() && !opts.force {
				return opts.localizer.MustLocalizeError("flag.error.requiredWhenNonInteractive", localize.NewEntry("Flag", "yes"))
			}

			if opts.kafkaID != "" {
				return runCmd(opts)
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			instanceID, ok := cfg.GetKafkaIdOk()
			if !ok {
				return opts.localizer.MustLocalizeError("kafka.topic.common.error.noKafkaSelected")
			}

			opts.kafkaID = instanceID

			return runCmd(opts)
		},
	}

	flags := flagutil.NewFlagSet(cmd, f.Localizer)
	flags.StringVar(&opts.topicName, "name", "", opts.localizer.MustLocalize("kafka.topic.common.flag.name.description"))

	_ = cmd.MarkFlagRequired("name")

	_ = cmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return cmdutil.FilterValidTopicNameArgs(f, toComplete)
	})
	flags.BoolVarP(&opts.force, "yes", "y", false, opts.localizer.MustLocalize("kafka.topic.delete.flag.yes.description"))

	return cmd
}

// nolint:funlen
func runCmd(opts *options) error {
	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	api, kafkaInstance, err := conn.API().KafkaAdmin(opts.kafkaID)
	if err != nil {
		return err
	}

	// perform delete topic API request
	_, httpRes, err := api.TopicsApi.GetTopic(opts.Context, opts.topicName).Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	topicNameTmplPair := localize.NewEntry("TopicName", opts.topicName)
	kafkaNameTmplPair := localize.NewEntry("InstanceName", kafkaInstance.GetName())
	if err != nil {
		if httpRes == nil {
			return err
		}
		if httpRes.StatusCode == http.StatusNotFound {
			return opts.localizer.MustLocalizeError("kafka.topic.common.error.topicNotFoundError", topicNameTmplPair, kafkaNameTmplPair)
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
			return opts.localizer.MustLocalizeError("kafka.topic.delete.error.mismatchedNameConfirmation", localize.NewEntry("ConfirmedName", userConfirmedName), localize.NewEntry("ActualName", opts.topicName))
		}
	}

	// perform delete topic API request
	httpRes, err = api.TopicsApi.DeleteTopic(opts.Context, opts.topicName).Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}
	if err != nil {
		if httpRes == nil {
			return err
		}

		operationTmplPair := localize.NewEntry("Operation", "delete")
		switch httpRes.StatusCode {
		case http.StatusNotFound:
			return opts.localizer.MustLocalizeError("kafka.topic.common.error.notFoundError", topicNameTmplPair, kafkaNameTmplPair)
		case http.StatusUnauthorized:
			return opts.localizer.MustLocalizeError("kafka.topic.common.error.unauthorized", operationTmplPair)
		case http.StatusForbidden:
			return opts.localizer.MustLocalizeError("kafka.topic.common.error.forbidden", operationTmplPair)
		case http.StatusInternalServerError:
			return opts.localizer.MustLocalizeError("kafka.topic.common.error.internalServerError")
		case http.StatusServiceUnavailable:
			return opts.localizer.MustLocalizeError("kafka.topic.common.error.unableToConnectToKafka", localize.NewEntry("Name", kafkaInstance.GetName()))
		default:
			return err
		}
	}

	opts.Logger.Info(opts.localizer.MustLocalize("kafka.topic.delete.log.info.topicDeleted", topicNameTmplPair, kafkaNameTmplPair))

	return nil
}
