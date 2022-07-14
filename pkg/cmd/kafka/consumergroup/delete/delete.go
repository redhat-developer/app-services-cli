package delete

import (
	"context"
	"net/http"

	kafkaflagutil "github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/kafkacmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"

	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

type options struct {
	kafkaID     string
	id          string
	skipConfirm bool

	IO             *iostreams.IOStreams
	Connection     factory.ConnectionFunc
	Logger         logging.Logger
	localizer      localize.Localizer
	Context        context.Context
	ServiceContext servicecontext.IContext
}

// NewDeleteConsumerGroupCommand gets a new command for deleting a consumer group.
func NewDeleteConsumerGroupCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Connection:     f.Connection,
		IO:             f.IOStreams,
		Logger:         f.Logger,
		localizer:      f.Localizer,
		Context:        f.Context,
		ServiceContext: f.ServiceContext,
	}

	cmd := &cobra.Command{
		Use:     "delete",
		Short:   opts.localizer.MustLocalize("kafka.consumerGroup.delete.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("kafka.consumerGroup.delete.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("kafka.consumerGroup.delete.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if opts.kafkaID != "" {
				return runCmd(opts)
			}

			kafkaInstance, err := contextutil.GetCurrentKafkaInstance(f)
			if err != nil {
				return err
			}

			opts.kafkaID = kafkaInstance.GetId()

			return runCmd(opts)
		},
	}

	flags := kafkaflagutil.NewFlagSet(cmd, opts.localizer)

	flags.AddYes(&opts.skipConfirm)
	flags.AddInstanceID(&opts.kafkaID)
	flags.StringVar(&opts.id, "id", "", opts.localizer.MustLocalize("kafka.consumerGroup.common.flag.id.description", localize.NewEntry("Action", "delete")))
	_ = cmd.MarkFlagRequired("id")

	// flag based completions for ID
	_ = cmd.RegisterFlagCompletionFunc("id", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return kafkacmdutil.FilterValidConsumerGroupIDs(f, toComplete)
	})

	return cmd
}

// nolint:funlen
func runCmd(opts *options) error {
	conn, err := opts.Connection()
	if err != nil {
		return err
	}

	api, kafkaInstance, err := conn.API().KafkaAdmin(opts.kafkaID)
	if err != nil {
		return err
	}

	_, httpRes, err := api.GroupsApi.GetConsumerGroupById(opts.Context, opts.id).Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	cgIDPair := localize.NewEntry("ID", opts.id)
	kafkaNameTmplPair := localize.NewEntry("InstanceName", kafkaInstance.GetName())
	if err != nil {
		if httpRes == nil {
			return err
		}
		if httpRes.StatusCode == http.StatusNotFound {
			return opts.localizer.MustLocalizeError("kafka.consumerGroup.common.error.notFoundError", cgIDPair, kafkaNameTmplPair)
		}
	}

	if !opts.skipConfirm {
		var confirmedID string
		promptConfirmDelete := &survey.Input{
			Message: opts.localizer.MustLocalize("kafka.consumerGroup.delete.input.name.message"),
		}

		err = survey.AskOne(promptConfirmDelete, &confirmedID)
		if err != nil {
			return err
		}

		if confirmedID != opts.id {
			return opts.localizer.MustLocalizeError("kafka.consumerGroup.delete.error.mismatchedIDConfirmation", localize.NewEntry("ConfirmedID", confirmedID), cgIDPair)
		}
	}

	httpRes, err = api.GroupsApi.DeleteConsumerGroupById(opts.Context, opts.id).Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	if err != nil {
		if httpRes == nil {
			return err
		}

		operationTmplPair := localize.NewEntry("Operation", "delete")

		switch httpRes.StatusCode {
		case http.StatusUnauthorized:
			return opts.localizer.MustLocalizeError("kafka.consumerGroup.common.error.unauthorized", operationTmplPair)
		case http.StatusForbidden:
			return opts.localizer.MustLocalizeError("kafka.consumerGroup.common.error.forbidden", operationTmplPair)
		case http.StatusLocked:
			return opts.localizer.MustLocalizeError("kafka.consumerGroup.delete.error.locked")
		case http.StatusInternalServerError:
			return opts.localizer.MustLocalizeError("kafka.consumerGroup.common.error.internalServerError")
		case http.StatusServiceUnavailable:
			return opts.localizer.MustLocalizeError("kafka.consumerGroup.common.error.unableToConnectToKafka", localize.NewEntry("Name", kafkaInstance.GetName()))
		default:
			return err
		}
	}

	opts.Logger.Info(opts.localizer.MustLocalize("kafka.consumerGroup.delete.log.info.consumerGroupDeleted", localize.NewEntry("ConsumerGroupID", opts.id), kafkaNameTmplPair))

	return nil
}
