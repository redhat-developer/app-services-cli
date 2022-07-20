package consume

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	kafkaflagutil "github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/kafkacmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal/client"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
	"github.com/spf13/cobra"
)

const (
	DefaultOffset    = ""
	DefaultLimit     = 20
	DefaultTimestamp = ""
	DefaultPartition = -1
	FormatKeyValue   = "key-value"
)

var outputFormatTypes = []string{dump.JSONFormat, dump.YAMLFormat, FormatKeyValue}

type options struct {
	topicName    string
	kafkaID      string
	partition    int32
	date         string
	timestamp    string
	limit        int32
	offset       string
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
		Short:   f.Localizer.MustLocalize("kafka.topic.consume.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("kafka.topic.consume.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("kafka.topic.consume.cmd.example"),
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
	flags.Int32Var(&opts.partition, "partition", DefaultPartition, f.Localizer.MustLocalize("kafka.topic.consume.flag.partition.description"))
	flags.StringVar(&opts.date, "from-date", DefaultTimestamp, f.Localizer.MustLocalize("kafka.topic.consume.flag.date.description"))
	flags.StringVar(&opts.timestamp, "from-timestamp", DefaultTimestamp, f.Localizer.MustLocalize("kafka.topic.consume.flag.timestamp.description"))
	flags.BoolVar(&opts.wait, "wait", false, f.Localizer.MustLocalize("kafka.topic.consume.flag.wait.description"))
	flags.StringVar(&opts.offset, "offset", DefaultOffset, f.Localizer.MustLocalize("kafka.topic.consume.flag.offset.description"))
	flags.Int32Var(&opts.limit, "limit", DefaultLimit, f.Localizer.MustLocalize("kafka.topic.consume.flag.limit.description"))
	flags.StringVar(&opts.outputFormat, "format", FormatKeyValue, f.Localizer.MustLocalize("kafka.topic.produce.flag.format.description"))

	_ = cmd.MarkFlagRequired("name")

	_ = cmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return kafkacmdutil.FilterValidTopicNameArgs(f, toComplete)
	})

	flags.AddInstanceID(&opts.kafkaID)
	flagutil.EnableStaticFlagCompletion(cmd, "format", outputFormatTypes)

	return cmd
}

func runCmd(opts *options) error {

	conn, err := opts.f.Connection()
	if err != nil {
		return err
	}

	api, kafkaInstance, err := conn.API().KafkaAdmin(opts.kafkaID)
	if err != nil {
		return err
	}

	// cannot set unix timestamp and the date
	if opts.date != DefaultTimestamp && opts.timestamp != DefaultTimestamp {
		return opts.f.Localizer.MustLocalizeError("kafka.topic.consume.error.dateAndTimestampConflict")
	}

	// check for flags that are exclusive to eachother
	if opts.offset != DefaultOffset && (opts.date != DefaultTimestamp || opts.timestamp != DefaultTimestamp) {
		return opts.f.Localizer.MustLocalizeError("kafka.topic.consume.error.offsetAndFromConflict")
	}

	if opts.wait {
		err := consumeAndWait(opts, api, kafkaInstance)
		if err != nil {
			return err
		}
	} else {
		records, err := consume(opts, api, kafkaInstance)
		if err != nil {
			return err
		}

		outputRecords(opts, records)
	}

	return nil
}

func consumeAndWait(opts *options, api *kafkainstanceclient.APIClient, kafkaInstance *kafkamgmtclient.KafkaRequest) error {

	if opts.partition == DefaultPartition {
		return opts.f.Localizer.MustLocalizeError("kafka.topic.consume.error.noPartitionWhenWaiting")
	}

	if opts.limit != DefaultLimit {
		opts.f.Logger.Info(opts.f.Localizer.MustLocalize("kafka.topic.consume.log.info.limitIgnored", localize.NewEntry("Limit", DefaultLimit)))
		opts.limit = DefaultLimit
	}

	if opts.offset != DefaultOffset {
		opts.f.Logger.Info(opts.f.Localizer.MustLocalize("kafka.topic.consume.log.info.offsetIgnored", localize.NewEntry("Offset", DefaultOffset)))
		opts.offset = DefaultOffset
	}

	if opts.date == DefaultTimestamp && opts.timestamp == DefaultTimestamp {
		// get current time in ISO 8601
		opts.date = time.Now().Format(time.RFC3339)
	}

	var max_offset int64
	first_consume := true
	for true {

		records, err := consume(opts, api, kafkaInstance)
		if err != nil {
			return err
		}

		record_count := len(records.Items)
		if record_count > 0 {
			max_offset = *(records.Items[record_count-1].Offset) + 1
			outputRecords(opts, records)

			if first_consume {
				// reset timestamp and date after first consume as it will
				// conflict with the max offset we are setting to only get new records
				opts.date = DefaultTimestamp
				opts.timestamp = DefaultTimestamp
				first_consume = false
			}
			opts.offset = fmt.Sprint(max_offset)
		}

		time.Sleep(1 * time.Second)
	}

	return nil
}

func consume(opts *options, api *kafkainstanceclient.APIClient, kafkaInstance *kafkamgmtclient.KafkaRequest) (*kafkainstanceclient.RecordList, error) {
	request := api.RecordsApi.ConsumeRecords(opts.f.Context, opts.topicName).Limit(opts.limit)

	if opts.partition != DefaultPartition {
		request = request.Partition(opts.partition)
	}

	if opts.offset != DefaultOffset {
		intOffset, err := strconv.ParseInt(opts.offset, 10, 64)
		if err != nil {
			return nil, opts.f.Localizer.MustLocalizeError("kafka.topic.comman.error.offsetInvalid", localize.NewEntry("Offset", opts.offset))
		}

		if intOffset < 0 {
			return nil, opts.f.Localizer.MustLocalizeError("kafka.topic.comman.error.offsetNegative")
		}

		request = request.Offset(int32(intOffset))
	}

	if opts.date != DefaultTimestamp {
		_, err := time.Parse(time.RFC3339, opts.date)
		if err != nil {
			return nil, opts.f.Localizer.MustLocalizeError("kafka.topic.comman.error.timeFormat", localize.NewEntry("Time", opts.date))
		}

		request = request.Timestamp(opts.date)
	}

	if opts.timestamp != DefaultTimestamp {
		digits, err := strconv.ParseInt(opts.timestamp, 10, 64)
		if err != nil {
			return nil, err
		}

		opts.timestamp = time.Unix(digits, 0).Format(time.RFC3339)
	}

	list, httpRes, err := request.Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	if err != nil {

		if httpRes == nil {
			return nil, err
		}

		if httpRes.StatusCode == http.StatusNotFound {
			return nil, opts.f.Localizer.MustLocalizeError("kafka.topic.common.error.topicNotFoundError",
				localize.NewEntry("TopicName", opts.topicName),
				localize.NewEntry("InstanceName", kafkaInstance.GetName()))
		}

		if httpRes.StatusCode == 400 {
			return nil, opts.f.Localizer.MustLocalizeError("kafka.topic.common.error.partitionNotFoundError",
				localize.NewEntry("Topic", opts.topicName),
				localize.NewEntry("Partition", opts.partition))
		}

		return nil, err
	}

	// fill in fields not set in message
	for i := 0; i < len(list.Items); i++ {
		record := &list.Items[i]

		if record.Key == nil {
			defaultKey := ""
			record.Key = &defaultKey
		}
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
		row := &recordsAsRows[i]
		if format == FormatKeyValue {
			if row.Key == "" {
				opts.f.Logger.Info(fmt.Sprintf("Message: %v", row.Value))
			} else {
				opts.f.Logger.Info(fmt.Sprintf("Key: %v\nMessage: %v", row.Key, row.Value))
			}
			opts.f.Logger.Info(fmt.Sprintf("Offset: %v", row.Offset))
			if opts.partition == DefaultPartition {
				opts.f.Logger.Info(fmt.Sprintf("Partition: %v", row.Partition))
			}
		} else {
			_ = dump.Formatted(opts.f.IOStreams.Out, format, row)
			opts.f.Logger.Info("")
		}
	}
}

func mapRecordsToRows(topic string, records *[]kafkainstanceclient.Record) []kafkaRow {

	rows := make([]kafkaRow, len(*records))

	for i := 0; i < len(*records); i++ {
		record := &(*records)[i]
		row := kafkaRow{
			Topic:     topic,
			Key:       record.GetKey(),
			Value:     record.Value,
			Partition: record.GetPartition(),
			Offset:    record.GetOffset(),
		}

		rows[i] = row
	}

	return rows
}
