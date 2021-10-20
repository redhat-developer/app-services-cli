package resetoffset

import (
	"context"
	"net/http"

	"github.com/redhat-developer/app-services-cli/pkg/icon"

	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmdutil"
	cmdFlagUtil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/kafka/consumergroup"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal/client"
	"github.com/spf13/cobra"
)

type options struct {
	kafkaID     string
	id          string
	skipConfirm bool
	value       string
	offset      string
	topic       string
	partitions  []int32

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	localizer  localize.Localizer
	Context    context.Context
}

var validator consumergroup.Validator

// NewResetOffsetConsumerGroupCommand gets a new command for resetting offset for a consumer group.
func NewResetOffsetConsumerGroupCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Connection: f.Connection,
		Config:     f.Config,
		IO:         f.IOStreams,
		Logger:     f.Logger,
		localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "reset-offset",
		Short:   opts.localizer.MustLocalize("kafka.consumerGroup.resetOffset.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("kafka.consumerGroup.resetOffset.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("kafka.consumerGroup.resetOffset.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			validator = consumergroup.Validator{
				Localizer: opts.localizer,
			}

			if opts.offset != "" {
				if err = validator.ValidateOffset(opts.offset); err != nil {
					return err
				}
			}

			if opts.value == "" && (opts.offset == consumergroup.OffsetAbsolute || opts.offset == consumergroup.OffsetTimestamp) {
				return opts.localizer.MustLocalizeError("kafka.consumerGroup.resetOffset.error.valueRequired", localize.NewEntry("Offset", opts.offset))
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
				return opts.localizer.MustLocalizeError("kafka.consumerGroup.common.error.noKafkaSelected")
			}

			opts.kafkaID = instanceID

			return runCmd(opts)
		},
	}

	flags := flagutil.NewFlagSet(cmd, opts.localizer)
	flags.BoolVarP(&opts.skipConfirm, "yes", "y", false, opts.localizer.MustLocalize("kafka.consumerGroup.resetOffset.flag.yes"))
	flags.StringVar(&opts.id, "id", "", opts.localizer.MustLocalize("kafka.consumerGroup.common.flag.id.description", localize.NewEntry("Action", "reset-offset")))
	flags.StringVar(&opts.value, "value", "", opts.localizer.MustLocalize("kafka.consumerGroup.resetOffset.flag.value"))
	flags.StringVar(&opts.offset, "offset", "", opts.localizer.MustLocalize("kafka.consumerGroup.resetOffset.flag.offset"))
	flags.StringVar(&opts.topic, "topic", "", opts.localizer.MustLocalize("kafka.consumerGroup.resetOffset.flag.topic"))
	flags.Int32SliceVar(&opts.partitions, "partitions", []int32{}, opts.localizer.MustLocalize("kafka.consumerGroup.resetOffset.flag.partitions"))

	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("offset")
	_ = cmd.MarkFlagRequired("topic")

	// flag based completions for ID
	_ = cmd.RegisterFlagCompletionFunc("id", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return cmdutil.FilterValidConsumerGroupIDs(f, toComplete)
	})

	// flag based completions for topic
	_ = cmd.RegisterFlagCompletionFunc("topic", func(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return cmdutil.FilterValidTopicNameArgs(f, toComplete)
	})

	cmdFlagUtil.EnableOutputFlagCompletion(cmd)
	cmdFlagUtil.EnableStaticFlagCompletion(cmd, "offset", consumergroup.ValidOffsets)

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

	offsetResetParams := kafkainstanceclient.ConsumerGroupResetOffsetParameters{
		Offset: opts.offset,
	}

	if opts.value != "" {
		offsetResetParams.Value = &opts.value
	}

	if opts.offset == consumergroup.OffsetAbsolute || opts.offset == consumergroup.OffsetTimestamp {
		if err = validator.ValidateOffsetValue(opts.offset, opts.value); err != nil {
			return err
		}
	}

	if opts.id != "" {
		_, httpRes, newErr := api.GroupsApi.GetConsumerGroupById(opts.Context, opts.id).Execute()
		if httpRes != nil {
			defer httpRes.Body.Close()
		}

		if newErr != nil {
			cgIDPair := localize.NewEntry("ID", opts.id)
			kafkaNameTmplPair := localize.NewEntry("InstanceName", kafkaInstance.GetName())
			if httpRes == nil {
				return newErr
			}
			if httpRes.StatusCode == http.StatusNotFound {
				return opts.localizer.MustLocalizeError("kafka.consumerGroup.common.error.notFoundError", cgIDPair, kafkaNameTmplPair)
			}
			return newErr
		}
	}

	if opts.topic != "" {
		_, httpRes, newErr := api.TopicsApi.GetTopic(opts.Context, opts.topic).Execute()
		if httpRes != nil {
			defer httpRes.Body.Close()
		}

		if newErr != nil {
			if httpRes == nil {
				return newErr
			}
			topicNameTmplPair := localize.NewEntry("TopicName", opts.topic)
			kafkaNameTmplPair := localize.NewEntry("InstanceName", kafkaInstance.GetName())
			if httpRes.StatusCode == http.StatusNotFound {
				return opts.localizer.MustLocalizeError("kafka.topic.common.error.notFoundError", topicNameTmplPair, kafkaNameTmplPair)
			}
			return newErr
		}

		topicToReset := kafkainstanceclient.TopicsToResetOffset{
			Topic: opts.topic,
		}

		if len(opts.partitions) != 0 {
			topicToReset.Partitions = &opts.partitions
		}

		topicsToResetArr := []kafkainstanceclient.TopicsToResetOffset{topicToReset}

		offsetResetParams.Topics = &topicsToResetArr
	}

	a := api.GroupsApi.ResetConsumerGroupOffset(opts.Context, opts.id).ConsumerGroupResetOffsetParameters(offsetResetParams)

	if !opts.skipConfirm {

		var confirmReset bool
		opts.localizer.MustLocalize("kafka.consumerGroup.resetOffset.input.confirmReset.message", localize.NewEntry("ID", opts.id))
		promptConfirmReset := &survey.Confirm{
			Message: opts.localizer.MustLocalize("kafka.consumerGroup.resetOffset.input.confirmReset.message", localize.NewEntry("ID", opts.id)),
		}

		if err = survey.AskOne(promptConfirmReset, &confirmReset); err != nil {
			return err
		}
		if !confirmReset {
			opts.Logger.Debug(opts.localizer.MustLocalize("kafka.consumerGroup.resetOffset.log.debug.cancelledReset"))
			return nil
		}
	}

	_, httpRes, err := a.Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	if err != nil {

		if httpRes == nil {
			return err
		}

		operationTmplPair := localize.NewEntry("Operation", "reset offset")

		switch httpRes.StatusCode {
		case http.StatusUnauthorized:
			return opts.localizer.MustLocalizeError("kafka.consumerGroup.common.error.unauthorized", operationTmplPair)
		case http.StatusForbidden:
			return opts.localizer.MustLocalizeError("kafka.consumerGroup.common.error.forbidden", operationTmplPair)
		case http.StatusInternalServerError:
			return opts.localizer.MustLocalizeError("kafka.consumerGroup.common.error.internalServerError")
		case http.StatusServiceUnavailable:
			return opts.localizer.MustLocalizeError("kafka.consumerGroup.common.error.unableToConnectToKafka", localize.NewEntry("Name", kafkaInstance.GetName()))
		default:
			return err
		}
	}

	defer httpRes.Body.Close()

	opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize(
		"kafka.consumerGroup.resetOffset.log.info.successful",
		localize.NewEntry("ConsumerGroupID", opts.id),
		localize.NewEntry("InstanceName", kafkaInstance.GetName())),
	)

	return nil
}
