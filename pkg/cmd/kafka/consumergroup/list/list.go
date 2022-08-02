package list

import (
	"context"
	"net/http"

	kafkaflagutil "github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/kafkacmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"

	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1/client"

	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/internal/build"
)

type options struct {
	Connection     factory.ConnectionFunc
	Logger         logging.Logger
	IO             *iostreams.IOStreams
	localizer      localize.Localizer
	Context        context.Context
	ServiceContext servicecontext.IContext

	output  string
	kafkaID string
	topic   string
	search  string
	page    int32
	size    int32
}

type consumerGroupRow struct {
	ConsumerGroupID   string                                 `json:"groupId,omitempty" header:"Consumer group ID"`
	ActiveMembers     int32                                  `json:"active_members,omitempty" header:"Active members"`
	PartitionsWithLag int32                                  `json:"lag,omitempty" header:"Partitions with lag"`
	State             kafkainstanceclient.ConsumerGroupState `json:"state,omitempty" header:"State"`
}

// NewListConsumerGroupCommand creates a new command to list consumer groups
func NewListConsumerGroupCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Connection:     f.Connection,
		Logger:         f.Logger,
		IO:             f.IOStreams,
		localizer:      f.Localizer,
		Context:        f.Context,
		ServiceContext: f.ServiceContext,
	}

	cmd := &cobra.Command{
		Use:     "list",
		Short:   opts.localizer.MustLocalize("kafka.consumerGroup.list.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("kafka.consumerGroup.list.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("kafka.consumerGroup.list.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if opts.output != "" && !flagutil.IsValidInput(opts.output, flagutil.ValidOutputFormats...) {
				return flagutil.InvalidValueError("output", opts.output, flagutil.ValidOutputFormats...)
			}

			if opts.page < 1 {
				return opts.localizer.MustLocalizeError("kafka.common.validation.page.error.invalid.minValue", localize.NewEntry("Page", opts.page))
			}

			if opts.size < 1 {
				return opts.localizer.MustLocalizeError("kafka.common.validation.size.error.invalid.minValue", localize.NewEntry("Size", opts.size))
			}

			if opts.kafkaID == "" {

				kafkaInstance, err := contextutil.GetCurrentKafkaInstance(f)
				if err != nil {
					return err
				}

				opts.kafkaID = kafkaInstance.GetId()
			}

			return runList(opts)
		},
	}

	flags := kafkaflagutil.NewFlagSet(cmd, opts.localizer)

	flags.AddInstanceID(&opts.kafkaID)

	flags.AddOutput(&opts.output)
	flags.StringVar(&opts.topic, "topic", "", opts.localizer.MustLocalize("kafka.consumerGroup.list.flag.topic.description"))
	flags.StringVar(&opts.search, "search", "", opts.localizer.MustLocalize("kafka.consumerGroup.list.flag.search"))
	flags.Int32VarP(&opts.page, "page", "", cmdutil.ConvertPageValueToInt32(build.DefaultPageNumber), opts.localizer.MustLocalize("kafka.consumerGroup.list.flag.page"))
	flags.Int32VarP(&opts.size, "size", "", cmdutil.ConvertSizeValueToInt32(build.DefaultPageSize), opts.localizer.MustLocalize("kafka.consumerGroup.list.flag.size"))

	_ = cmd.RegisterFlagCompletionFunc("topic", func(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return kafkacmdutil.FilterValidTopicNameArgs(f, toComplete)
	})

	flagutil.EnableOutputFlagCompletion(cmd)

	return cmd
}

// nolint:funlen
func runList(opts *options) (err error) {
	conn, err := opts.Connection()
	if err != nil {
		return err
	}

	api, kafkaInstance, err := conn.API().KafkaAdmin(opts.kafkaID)
	if err != nil {
		return err
	}

	req := api.GroupsApi.GetConsumerGroups(opts.Context)

	if opts.topic != "" {
		req = req.Topic(opts.topic)
	}
	if opts.search != "" {
		req = req.GroupIdFilter(opts.search)
	}

	req = req.Size(opts.size)

	req = req.Page(opts.page)

	consumerGroupData, httpRes, err := req.Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}
	if err != nil {
		if httpRes == nil {
			return err
		}

		operationTmplPair := localize.NewEntry("Operation", "list")

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

	if !checkForConsumerGroups(int(consumerGroupData.GetTotal()), opts, kafkaInstance.GetName()) {
		return nil
	}

	switch opts.output {
	case dump.EmptyFormat:
		opts.Logger.Info("")
		consumerGroups := consumerGroupData.GetItems()
		rows := mapConsumerGroupResultsToTableFormat(consumerGroups)
		dump.Table(opts.IO.Out, rows)
	default:
		return dump.Formatted(opts.IO.Out, opts.output, consumerGroupData)
	}

	return nil
}

func mapConsumerGroupResultsToTableFormat(consumerGroups []kafkainstanceclient.ConsumerGroup) []consumerGroupRow {
	rows := make([]consumerGroupRow, len(consumerGroups))

	for i, t := range consumerGroups {
		metrics := t.GetMetrics()
		row := consumerGroupRow{
			ConsumerGroupID:   t.GetGroupId(),
			ActiveMembers:     metrics.GetActiveConsumers(),
			PartitionsWithLag: metrics.GetLaggingPartitions(),
			State:             t.GetState(),
		}
		rows[i] = row
	}

	return rows
}

// checks if there are any consumer groups available
// prints to stderr if not
func checkForConsumerGroups(count int, opts *options, kafkaName string) (hasCount bool) {
	kafkaNameTmplPair := localize.NewEntry("InstanceName", kafkaName)
	if count == 0 && opts.output == "" {
		if opts.topic == "" {
			opts.Logger.Info(opts.localizer.MustLocalize("kafka.consumerGroup.list.log.info.noConsumerGroups", kafkaNameTmplPair))
		} else {
			opts.Logger.Info(opts.localizer.MustLocalize("kafka.consumerGroup.list.log.info.noConsumerGroupsForTopic", kafkaNameTmplPair, localize.NewEntry("TopicName", opts.topic)))
		}

		return false
	}

	return true
}
