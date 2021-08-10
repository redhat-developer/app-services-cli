package resetoffset

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	"github.com/redhat-developer/app-services-cli/pkg/cmdutil"
	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal/client"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type Options struct {
	kafkaID     string
	id          string
	skipConfirm bool
	value       string
	offset      string
	topic       string
	partitions  []int32
	output      string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     func() (logging.Logger, error)
	localizer  localize.Localizer
}

type updatedConsumerRow struct {
	Topic     string `json:"groupId,omitempty" header:"Topic"`
	Partition int32  `json:"active_members,omitempty" header:"Partition"`
	Offset    int32  `json:"lag,omitempty" header:"Offset"`
}

// NewResetOffsetConsumerGroupCommand gets a new command for resetting offset for a consumer group.
func NewResetOffsetConsumerGroupCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Connection: f.Connection,
		Config:     f.Config,
		IO:         f.IOStreams,
		Logger:     f.Logger,
		localizer:  f.Localizer,
	}

	cmd := &cobra.Command{
		Use:     "reset-offset",
		Short:   opts.localizer.MustLocalize("kafka.consumerGroup.resetOffset.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("kafka.consumerGroup.resetOffset.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("kafka.consumerGroup.resetOffset.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			if opts.output != "" && !flagutil.IsValidInput(opts.output, flagutil.ValidOutputFormats...) {
				return flag.InvalidValueError("output", opts.output, flagutil.ValidOutputFormats...)
			}

			if opts.offset != "" {
				if err = flag.ValidateOffset(opts.offset); err != nil {
					return err
				}
			}

			if opts.value == "" && (opts.offset == "absolute" || opts.offset == "timestamp") {
				return errors.New(opts.localizer.MustLocalize("kafka.consumerGroup.resetOffset.error.valueRequired", localize.NewEntry("Offset", opts.offset)))
			}

			if opts.kafkaID != "" {
				return runCmd(opts)
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			if !cfg.HasKafka() {
				return errors.New(opts.localizer.MustLocalize("kafka.consumerGroup.common.error.noKafkaSelected"))
			}

			opts.kafkaID = cfg.Services.Kafka.ClusterID

			return runCmd(opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.skipConfirm, "yes", "y", false, opts.localizer.MustLocalize("kafka.consumerGroup.resetOffset.flag.yes"))
	cmd.Flags().StringVar(&opts.id, "id", "", opts.localizer.MustLocalize("kafka.consumerGroup.common.flag.id.description", localize.NewEntry("Action", "reset-offset")))
	cmd.Flags().StringVar(&opts.value, "value", "", opts.localizer.MustLocalize("kafka.consumerGroup.resetOffset.flag.value"))
	cmd.Flags().StringVar(&opts.offset, "offset", "", opts.localizer.MustLocalize("kafka.consumerGroup.resetOffset.flag.offset"))
	cmd.Flags().StringVar(&opts.topic, "topic", "", opts.localizer.MustLocalize("kafka.consumerGroup.resetOffset.flag.topic"))
	cmd.Flags().Int32SliceVar(&opts.partitions, "partitions", []int32{}, opts.localizer.MustLocalize("kafka.consumerGroup.resetOffset.flag.partitions"))
	cmd.Flags().StringVarP(&opts.output, "output", "o", "", opts.localizer.MustLocalize("kafka.consumerGroup.resetOffset.flag.output"))

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

	flagutil.EnableOutputFlagCompletion(cmd)
	flagutil.EnableStaticFlagCompletion(cmd, "offset", flagutil.ValidOffsets)

	return cmd
}

// nolint:funlen
func runCmd(opts *Options) error {

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	api, kafkaInstance, err := conn.API().KafkaAdmin(opts.kafkaID)
	if err != nil {
		return err
	}

	ctx := context.Background()

	offsetResetParams := kafkainstanceclient.ConsumerGroupResetOffsetParameters{
		Offset: opts.offset,
	}

	if opts.value != "" {
		offsetResetParams.Value = &opts.value
	}

	if opts.topic != "" {
		topicToReset := kafkainstanceclient.TopicsToResetOffset{
			Topic: opts.topic,
		}

		if len(opts.partitions) != 0 {
			topicToReset.Partitions = &opts.partitions
		}

		topicsToResetArr := []kafkainstanceclient.TopicsToResetOffset{topicToReset}

		offsetResetParams.Topics = &topicsToResetArr
	}

	a := api.GroupsApi.ResetConsumerGroupOffset(ctx, opts.id).ConsumerGroupResetOffsetParameters(offsetResetParams)

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
			logger.Debug(opts.localizer.MustLocalize("kafka.consumerGroup.resetOffset.log.debug.cancelledReset"))
			return nil
		}
	}

	updatedConsumers, httpRes, err := a.Execute()

	if err != nil {

		if httpRes == nil {
			return err
		}

		operationTmplPair := localize.NewEntry("Operation", "reset offset")

		switch httpRes.StatusCode {
		case http.StatusUnauthorized:
			return errors.New(opts.localizer.MustLocalize("kafka.consumerGroup.common.error.unauthorized", operationTmplPair))
		case http.StatusForbidden:
			return errors.New(opts.localizer.MustLocalize("kafka.consumerGroup.common.error.forbidden", operationTmplPair))
		case http.StatusInternalServerError:
			return errors.New(opts.localizer.MustLocalize("kafka.consumerGroup.common.error.internalServerError"))
		case http.StatusServiceUnavailable:
			return errors.New(opts.localizer.MustLocalize("kafka.consumerGroup.common.error.unableToConnectToKafka", localize.NewEntry("Name", kafkaInstance.GetName())))
		default:
			return err
		}
	}

	defer httpRes.Body.Close()

	logger.Info(opts.localizer.MustLocalize(
		"kafka.consumerGroup.resetOffset.log.info.successful",
		localize.NewEntry("ConsumerGroupID", opts.id),
		localize.NewEntry("InstanceName", kafkaInstance.GetName())),
	)

	switch opts.output {
	case "json":
		data, _ := json.Marshal(updatedConsumers)
		_ = dump.JSON(opts.IO.Out, data)
	case "yaml", "yml":
		data, _ := yaml.Marshal(updatedConsumers)
		_ = dump.YAML(opts.IO.Out, data)
	default:
		logger.Info("")
		consumers := updatedConsumers.GetItems()
		rows := mapResetOffsetResultToTableFormat(consumers)
		dump.Table(opts.IO.Out, rows)

		return nil
	}

	return nil

}

func mapResetOffsetResultToTableFormat(consumers []kafkainstanceclient.ConsumerGroupResetOffsetResultItem) []updatedConsumerRow {
	rows := []updatedConsumerRow{}

	for _, t := range consumers {

		row := updatedConsumerRow{
			Topic:     t.GetTopic(),
			Partition: t.GetPartition(),
			Offset:    t.GetOffset(),
		}
		rows = append(rows, row)
	}

	return rows
}
