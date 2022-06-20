package consume

import (
	kafkaflagutil "github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"

	"strings"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/kafkacmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal/client"
	"github.com/spf13/cobra"
)

var outputFormatTypes = []string{dump.JSONFormat, dump.YAMLFormat}

type options struct {
	topicName    string
	kafkaID      string
	partition    int32
	timestamp    string
	limit        int32
	offset       int64
	wait         bool
	outputFormat string

	f *factory.Factory
}

// row is the details of a record produced needed to print to a table
type kafkaRow struct {
	Topic     string `json:"topic" header:"Topic"`
	Key       string `json:"key" header:"Key"`
	Value     string `json:"value" header:"Value"`
	Partition int32  `json:"partition" header:"Partition"`
	Offset    int64  `json:"offset" header:"Offset"`
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
	flags.BoolVar(&opts.wait, "wait", false, f.Localizer.MustLocalize("kafka.topic.consume.flag.wait.description"))
	flags.Int64Var(&opts.offset, "offset", 0, f.Localizer.MustLocalize("kafka.topic.consume.flag.offset.description"))
	flags.Int32Var(&opts.limit, "limit", 20, f.Localizer.MustLocalize("kafka.topic.consume.flag.limit.description"))
	flags.StringVar(&opts.outputFormat, "format", "json", f.Localizer.MustLocalize("kafka.topic.consume.flag.format.description"))

	_ = cmd.MarkFlagRequired("name")

	_ = cmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return kafkacmdutil.FilterValidTopicNameArgs(f, toComplete)
	})

	flags.AddInstanceID(&opts.kafkaID)
	flagutil.EnableStaticFlagCompletion(cmd, "format", outputFormatTypes)

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

	if opts.wait {

		max_offset := opts.offset
		for true {

			records, err := consume(opts, api)
			if err != nil {
				return err
			}

			record_count := len(records.Items)
			if record_count > 0 {
				max_offset = *(records.Items[record_count-1].Offset) + 1
				outputRecords(opts, records)
			}

			opts.offset = max_offset
		}
	} else {
		records, err := consume(opts, api)
		if err != nil {
			return err
		}

		outputRecords(opts, records)
	}

	return nil
}

func consume(opts *options, api *kafkainstanceclient.APIClient) (*kafkainstanceclient.RecordList, error) {

	request := api.RecordsApi.ConsumeRecords(opts.f.Context, opts.topicName).Limit(opts.limit).Partition(opts.partition).Offset(int32(opts.offset))
	if opts.timestamp != "" {
		// setting timestamp as "" (not set by user) is not valid
		// not setting timestamp is handled by the request
		request = request.Timestamp(opts.timestamp)
	}

	list, _, err := request.Execute()
	if err != nil {
		return nil, err
	}

	return &list, nil
}

func outputRecords(opts *options, records *kafkainstanceclient.RecordList) {
	recordsAsRows := mapRecordsToRows(opts.topicName, &records.Items)

	if len(records.Items) == 0 {
		opts.f.Logger.Info(opts.f.Localizer.MustLocalize("kafka.common.log.info.noRecords"))
		return
	}

	format := opts.outputFormat
	if format == dump.EmptyFormat {
		format = dump.JSONFormat
	}

	for i := 0; i < len(recordsAsRows); i++ {
		dump.Formatted(opts.f.IOStreams.Out, format, recordsAsRows[i])
	}
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
			Offset:    *record.Offset,
		}

		rows[i] = row
	}

	return rows
}
