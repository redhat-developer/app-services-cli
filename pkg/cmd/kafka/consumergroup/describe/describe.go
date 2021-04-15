package describe

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"sort"

	cgutil "github.com/redhat-developer/app-services-cli/pkg/kafka/consumergroup"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/internal/localizer"
	strimziadminclient "github.com/redhat-developer/app-services-cli/pkg/api/strimzi-admin/client"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	"github.com/redhat-developer/app-services-cli/pkg/color"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type Options struct {
	kafkaID      string
	outputFormat string
	id           string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
}

type consumerRow struct {
	MemberID      string `json:"memberId,omitempty" header:"Member ID"`
	Partition     int    `json:"partition,omitempty" header:"Partition"`
	Topic         string `json:"topic,omitempty" header:"Topic"`
	LogEndOffset  int    `json:"logEndOffset,omitempty" header:"Log end offset"`
	CurrentOffset int    `json:"offset,omitempty" header:"Current offset"`
	OffsetLag     int    `json:"lag,omitempty" header:"Offset lag"`
}

// NewDescribeConsumerGroupCommand gets a new command for describing a consumer group.
func NewDescribeConsumerGroupCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Connection: f.Connection,
		Config:     f.Config,
		IO:         f.IOStreams,
	}
	cmd := &cobra.Command{
		Use:     localizer.MustLocalizeFromID("kafka.consumerGroup.describe.cmd.use"),
		Short:   localizer.MustLocalizeFromID("kafka.consumerGroup.describe.cmd.shortDescription"),
		Long:    localizer.MustLocalizeFromID("kafka.consumerGroup.describe.cmd.longDescription"),
		Example: localizer.MustLocalizeFromID("kafka.consumerGroup.describe.cmd.example"),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			opts.id = args[0]

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
				return errors.New(localizer.MustLocalizeFromID("kafka.consumerGroup.common.error.noKafkaSelected"))
			}

			opts.kafkaID = cfg.Services.Kafka.ClusterID

			return runCmd(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "", localizer.MustLocalize(&localizer.Config{
		MessageID: "kafka.consumerGroup.common.flag.output.description",
	}))

	return cmd
}

func runCmd(opts *Options) error {
	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	api, kafkaInstance, err := conn.API().TopicAdmin(opts.kafkaID)
	if err != nil {
		return err
	}

	ctx := context.Background()

	consumerGroupData, httpRes, err := api.GetConsumerGroupById(ctx, opts.id).Execute()

	if err != nil {
		if httpRes == nil {
			return err
		}

		switch httpRes.StatusCode {
		case 404:
			return errors.New(localizer.MustLocalize(&localizer.Config{
				MessageID: "kafka.consumerGroup.common.error.notFoundError",
				TemplateData: map[string]interface{}{
					"ID":           opts.id,
					"InstanceName": kafkaInstance.GetName(),
				},
			}))
		case 401:
			return errors.New(localizer.MustLocalize(&localizer.Config{
				MessageID: "kafka.consumerGroup.common.error.unauthorized",
				TemplateData: map[string]interface{}{
					"Operation": "view",
				},
			}))
		case 403:
			return errors.New(localizer.MustLocalize(&localizer.Config{
				MessageID: "kafka.consumerGroup.common.error.forbidden",
				TemplateData: map[string]interface{}{
					"Operation": "view",
				},
			}))
		case 500:
			return errors.New(localizer.MustLocalizeFromID("kafka.consumerGroup.common.error.internalServerError"))
		case 503:
			return fmt.Errorf("%v: %w", localizer.MustLocalize(&localizer.Config{
				MessageID: "kafka.consumerGroup.common.error.unableToConnectToKafka",
				TemplateData: map[string]interface{}{
					"Name": kafkaInstance.GetName(),
				},
			}), err)
		default:
			return err
		}
	}

	stdout := opts.IO.Out
	switch opts.outputFormat {
	case "json":
		data, _ := json.Marshal(consumerGroupData)
		_ = dump.JSON(stdout, data)
	case "yaml", "yml":
		data, _ := yaml.Marshal(consumerGroupData)
		_ = dump.YAML(stdout, data)
	default:
		printConsumerGroupDetails(stdout, consumerGroupData)
	}

	return nil
}

func mapConsumerGroupDescribeToTableFormat(consumers []strimziadminclient.Consumer) []consumerRow {
	var rows []consumerRow = []consumerRow{}

	for _, consumer := range consumers {

		row := consumerRow{
			Partition:     int(consumer.GetPartition()),
			Topic:         consumer.GetTopic(),
			MemberID:      consumer.GetMemberId(),
			LogEndOffset:  int(consumer.GetLogEndOffset()),
			CurrentOffset: int(consumer.GetOffset()),
			OffsetLag:     int(consumer.GetLag()),
		}
		rows = append(rows, row)
	}

	// sort members by partition number
	sort.Slice(rows, func(i, j int) bool {
		return rows[i].Partition < rows[j].Partition
	})

	return rows
}

func printConsumerGroupDetails(w io.Writer, consumerGroupData strimziadminclient.ConsumerGroup) {
	fmt.Fprintln(w, "")
	consumers := consumerGroupData.GetConsumers()

	activeMembersCount := cgutil.GetActiveConsumersCount(consumers)
	partitionsWithLagCount := cgutil.GetPartitionsWithLag(consumers)

	fmt.Fprintln(w, color.Bold(localizer.MustLocalizeFromID("kafka.consumerGroup.describe.output.activeMembers")), activeMembersCount, "\t", color.Bold(localizer.MustLocalizeFromID("kafka.consumerGroup.describe.output.partitionsWithLag")), partitionsWithLagCount)
	fmt.Fprintln(w, "")

	rows := mapConsumerGroupDescribeToTableFormat(consumers)
	dump.Table(w, rows)
}
