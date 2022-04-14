package describe

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sort"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/consumergroup/groupcmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/kafkacmdutil"

	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/color"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal/client"

	"github.com/spf13/cobra"
)

type options struct {
	kafkaID      string
	outputFormat string
	id           string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	localizer  localize.Localizer
	Context    context.Context
}

type consumerRow struct {
	MemberID      string `json:"memberId,omitempty" header:"Consumer ID"`
	Partition     int    `json:"partition,omitempty" header:"Partition"`
	Topic         string `json:"topic,omitempty" header:"Topic"`
	LogEndOffset  int    `json:"logEndOffset,omitempty" header:"Log end offset"`
	CurrentOffset int    `json:"offset,omitempty" header:"Current offset"`
	OffsetLag     int    `json:"lag,omitempty" header:"Offset lag"`
}

// NewDescribeConsumerGroupCommand gets a new command for describing a consumer group.
func NewDescribeConsumerGroupCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Connection: f.Connection,
		Config:     f.Config,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
		Context:    f.Context,
	}
	cmd := &cobra.Command{
		Use:     "describe",
		Short:   opts.localizer.MustLocalize("kafka.consumerGroup.describe.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("kafka.consumerGroup.describe.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("kafka.consumerGroup.describe.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if opts.outputFormat != "" {
				if err = flagutil.ValidateOutput(opts.outputFormat); err != nil {
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

			instanceID, ok := cfg.GetKafkaIdOk()
			if !ok {
				return opts.localizer.MustLocalizeError("kafka.consumerGroup.common.error.noKafkaSelected")
			}

			opts.kafkaID = instanceID

			return runCmd(opts)
		},
	}

	flags := flagutil.NewFlagSet(cmd, opts.localizer)

	flags.AddOutput(&opts.outputFormat)
	flags.StringVar(&opts.id, "id", "", opts.localizer.MustLocalize("kafka.consumerGroup.common.flag.id.description", localize.NewEntry("Action", "view")))
	_ = cmd.MarkFlagRequired("id")

	// flag based completions for ID
	_ = cmd.RegisterFlagCompletionFunc("id", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return kafkacmdutil.FilterValidConsumerGroupIDs(f, toComplete)
	})

	flagutil.EnableOutputFlagCompletion(cmd)

	return cmd
}

func runCmd(opts *options) error {
	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	api, kafkaInstance, err := conn.API().KafkaAdmin(opts.kafkaID)
	if err != nil {
		return err
	}

	consumerGroupData, httpRes, err := api.GroupsApi.GetConsumerGroupById(opts.Context, opts.id).Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	if err != nil {
		if httpRes == nil {
			return err
		}

		cgIDPair := localize.NewEntry("ID", opts.id)
		kafkaNameTmplPair := localize.NewEntry("InstanceName", kafkaInstance.GetName())
		operationTmplPair := localize.NewEntry("Operation", "view")

		switch httpRes.StatusCode {
		case http.StatusNotFound:
			return opts.localizer.MustLocalizeError("kafka.consumerGroup.common.error.notFoundError", cgIDPair, kafkaNameTmplPair)
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

	stdout := opts.IO.Out

	switch opts.outputFormat {
	case dump.EmptyFormat:
		printConsumerGroupDetails(stdout, consumerGroupData, opts.localizer)
	default:
		return dump.Formatted(stdout, opts.outputFormat, consumerGroupData)
	}

	return nil
}

func mapConsumerGroupDescribeToTableFormat(consumers []kafkainstanceclient.Consumer) []consumerRow {
	rows := make([]consumerRow, len(consumers))

	for i, consumer := range consumers {

		row := consumerRow{
			Partition:     int(consumer.GetPartition()),
			Topic:         consumer.GetTopic(),
			MemberID:      consumer.GetMemberId(),
			LogEndOffset:  int(consumer.GetLogEndOffset()),
			CurrentOffset: int(consumer.GetOffset()),
			OffsetLag:     int(consumer.GetLag()),
		}

		if consumer.GetMemberId() == "" {
			row.MemberID = color.Italic("unassigned")
		}

		rows[i] = row
	}

	// sort members by partition number
	sort.Slice(rows, func(i, j int) bool {
		return rows[i].Partition < rows[j].Partition
	})

	return rows
}

// print the consumer group details
func printConsumerGroupDetails(w io.Writer, consumerGroupData kafkainstanceclient.ConsumerGroup, localizer localize.Localizer) {
	fmt.Fprintln(w, "")
	consumers := consumerGroupData.GetConsumers()

	activeMembersCount := groupcmdutil.GetActiveConsumersCount(consumers)
	partitionsWithLagCount := groupcmdutil.GetPartitionsWithLag(consumers)
	unassignedPartitions := groupcmdutil.GetUnassignedPartitions(consumers)

	fmt.Fprintln(w, color.Bold(localizer.MustLocalize("kafka.consumerGroup.describe.output.activeMembers")), activeMembersCount, "\t", color.Bold(localizer.MustLocalize("kafka.consumerGroup.describe.output.partitionsWithLag")), partitionsWithLagCount, "\t", color.Bold(localizer.MustLocalize("kafka.consumerGroup.describe.output.unassignedPartitions")), unassignedPartitions)
	fmt.Fprintln(w, "")

	rows := mapConsumerGroupDescribeToTableFormat(consumers)
	dump.Table(w, rows)
}
