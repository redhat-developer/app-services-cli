package resetoffset

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	"github.com/redhat-developer/app-services-cli/pkg/cmdutil"
	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal/client"
	"github.com/spf13/cobra"
)

type Options struct {
	kafkaID     string
	id          string
	skipConfirm bool
	value       string
	offset      string
	topic       string
	partitions  string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     func() (logging.Logger, error)
	localizer  localize.Localizer
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

			if opts.offset != "" {
				if err = flag.ValidateOffset(opts.offset); err != nil {
					return err
				}
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
	cmd.Flags().StringVar(&opts.partitions, "partitions", "", opts.localizer.MustLocalize("kafka.consumerGroup.resetOffset.flag.partitions"))

	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("offset")

	// flag based completions for ID
	_ = cmd.RegisterFlagCompletionFunc("id", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return cmdutil.FilterValidConsumerGroupIDs(f, toComplete)
	})

	// flag based completions for topic
	_ = cmd.RegisterFlagCompletionFunc("topic", func(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return cmdutil.FilterValidTopicNameArgs(f, toComplete)
	})

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

		if opts.partitions != "" {
			partitionsStr := strings.Fields(opts.partitions)

			partitionsArr := []int32{}

			for _, partition := range partitionsStr {
				partitionInt, convErr := strconv.ParseInt(partition, 10, 32)
				if convErr != nil {
					return convErr
				}
				partitionsArr = append(partitionsArr, int32(partitionInt))
			}

			topicToReset.Partitions = &partitionsArr
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
	// display the response object(todo)
	fmt.Println(updatedConsumers)

	if err != nil {

		if httpRes == nil {
			return err
		}

		operationTmplPair := localize.NewEntry("Operation", "reset offset")

		switch httpRes.StatusCode {
		case 401:
			return errors.New(opts.localizer.MustLocalize("kafka.consumerGroup.common.error.unauthorized", operationTmplPair))
		case 403:
			return errors.New(opts.localizer.MustLocalize("kafka.consumerGroup.common.error.forbidden", operationTmplPair))
		case 500:
			return errors.New(opts.localizer.MustLocalize("kafka.consumerGroup.common.error.internalServerError"))
		case 503:
			return errors.New(opts.localizer.MustLocalize("kafka.consumerGroup.common.error.unableToConnectToKafka", localize.NewEntry("Name", kafkaInstance.GetName())))
		default:
			return err
		}
	}

	return nil

}
