package describe

import (
	"context"
	"errors"
	"fmt"
	"github.com/redhat-developer/app-services-cli/pkg/cmdutil"
	cgutil "github.com/redhat-developer/app-services-cli/pkg/kafka/consumergroup"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal/client"
	"io"
	"net/http"
	"sort"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"
	"github.com/redhat-developer/app-services-cli/pkg/color"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
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
				if err = flag.ValidateOutput(opts.outputFormat); err != nil {
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

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "", opts.localizer.MustLocalize("kafka.consumerGroup.common.flag.output.description"))
	cmd.Flags().StringVar(&opts.id, "id", "", opts.localizer.MustLocalize("kafka.consumerGroup.common.flag.id.description", localize.NewEntry("Action", "view")))
	_ = cmd.MarkFlagRequired("id")

	// flag based completions for ID
	_ = cmd.RegisterFlagCompletionFunc("id", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return cmdutil.FilterValidConsumerGroupIDs(f, toComplete)
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
			return errors.New(opts.localizer.MustLocalize("kafka.consumerGroup.common.error.notFoundError", cgIDPair, kafkaNameTmplPair))
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

	stdout := opts.IO.Out

	switch opts.outputFormat {
	case "table":
		printConsumerGroupDetails(stdout, consumerGroupData, opts.localizer)
	default:
		dump.PrintDataInFormat(opts.outputFormat, consumerGroupData, stdout)
	}

	return nil
}

func mapConsumerGroupDescribeToTableFormat(consumers []kafkainstanceclient.Consumer) []consumerRow {
	rows := []consumerRow{}

	for _, consumer := range consumers {

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

		rows = append(rows, row)
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

	activeMembersCount := cgutil.GetActiveConsumersCount(consumers)
	partitionsWithLagCount := cgutil.GetPartitionsWithLag(consumers)
	unassignedPartitions := cgutil.GetUnassignedPartitions(consumers)

	fmt.Fprintln(w, color.Bold(localizer.MustLocalize("kafka.consumerGroup.describe.output.activeMembers")), activeMembersCount, "\t", color.Bold(localizer.MustLocalize("kafka.consumerGroup.describe.output.partitionsWithLag")), partitionsWithLagCount, "\t", color.Bold(localizer.MustLocalize("kafka.consumerGroup.describe.output.unassignedPartitions")), unassignedPartitions)
	fmt.Fprintln(w, "")

	rows := mapConsumerGroupDescribeToTableFormat(consumers)
	dump.Table(w, rows)
}
