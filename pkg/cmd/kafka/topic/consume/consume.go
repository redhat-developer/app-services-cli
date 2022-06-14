package consume

import (
	kafkaflagutil "github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"

	"strings"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/kafkacmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal/client"
	"github.com/spf13/cobra"
)

type options struct {
	topicName string
	kafkaID   string
	partition int32
	timestamp string
	limit     int32

	f *factory.Factory
}

// row is the details of a record produced needed to print to a table
type kafkaRow struct {
	Topic     string `json:"topic" header:"Topic"`
	Key       string `json:"key" header:"Key"`
	Value     string `json:"value" header:"Value"`
	Partition int32  `json:"partition" header:"Partition"`
}

// NewComsumeTopicCommand creates a new command for producing to a kafka topic.
func NewConsumeTopicCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		f: f,
	}

	cmd := &cobra.Command{
		Use:     "consume",
		Short:   "consume short",
		Long:    "consume long",
		Example: "consume example",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if opts.kafkaID == "" {

				kafkaInstance, err := contextutil.GetCurrentKafkaInstance(f)
				if err != nil {
					return err
				}

				opts.kafkaID = kafkaInstance.GetId()
			}

			return runCmd(opts)
		},
	}

	flags := kafkaflagutil.NewFlagSet(cmd, f.Localizer)

	flags.StringVar(&opts.topicName, "name", "", f.Localizer.MustLocalize("kafka.topic.common.flag.name.description"))
	flags.Int32Var(&opts.partition, "partition", 0, f.Localizer.MustLocalize("kafka.topic.consume.flag.partition.description"))
	flags.StringVar(&opts.timestamp, "timestamp", "", f.Localizer.MustLocalize("kafka.topic.consume.flag.timestamp.description"))
	flags.Int32Var(&opts.limit, "limit", 20, f.Localizer.MustLocalize("kafka.topic.consume.flag.limit.description"))

	_ = cmd.MarkFlagRequired("name")

	_ = cmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return kafkacmdutil.FilterValidTopicNameArgs(f, toComplete)
	})

	flags.AddInstanceID(&opts.kafkaID)

	return cmd
}

func runCmd(opts *options) error {
	conn, err := opts.f.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	api, _, err := conn.API().KafkaAdmin(opts.kafkaID)
	if err != nil {
		return err
	}

	list, _, err := api.RecordsApi.ConsumeRecords(opts.f.Context, opts.topicName).Limit(opts.limit).Partition(opts.partition).Timestamp(opts.timestamp).Execute()
	if err != nil {
		return err
	}

	dump.Table(opts.f.IOStreams.Out, mapRecordsToRows(opts.topicName, &list.Items))
	opts.f.Logger.Info("")

	return nil
}

func mapRecordsToRows(topic string, records *[]kafkainstanceclient.Record) []kafkaRow {

	rows := make([]kafkaRow, len(*records))

	for i := 0; i < len(*records); i++ {
		record := &(*records)[i]
		row := kafkaRow{
			Topic:     topic,
			Key:       *record.Key,
			Value:     strings.TrimSuffix(record.Value, "\n"), // trailing new line gives weird printing of table
			Partition: *record.Partition,
		}

		rows[i] = row
	}

	return rows
}
