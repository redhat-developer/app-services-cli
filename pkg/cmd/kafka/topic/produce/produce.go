package produce

import (
	"os"

	kafkaflagutil "github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal/client"

	"bufio"
	"io/ioutil"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/kafkacmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

type options struct {
	topicName string
	kafkaID   string
	key       string
	file      string
	partition int32

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

// NewProduceTopicCommand creates a new command for producing to a kafka topic.
func NewProduceTopicCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		f: f,
	}

	cmd := &cobra.Command{
		Use:     "produce",
		Short:   f.Localizer.MustLocalize("kafka.topic.produce.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("kafka.topic.produce.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("kafka.topic.produce.cmd.example"),
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
	flags.StringVar(&opts.key, "key", "", f.Localizer.MustLocalize("kafka.topic.produce.flag.key.description"))
	flags.Int32Var(&opts.partition, "partition", 0, f.Localizer.MustLocalize("kafka.topic.produce.flag.partition.description"))
	flags.StringVar(&opts.file, "file", "", f.Localizer.MustLocalize("kafka.topic.produce.flag.key.file"))

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

	var value string

	if opts.file != "" {
		bytes, err := ioutil.ReadFile(opts.file)
		if err != nil {
			return err
		}

		value = string(bytes)
	} else {
		// if value is being piped then cannot have delimeter as \n
		info, _ := os.Stdin.Stat()
		if info.Mode()&os.ModeCharDevice == 0 {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}

			value = string(bytes)
		} else {
			value, err = bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				return err
			}

		}
	}

	record := kafkainstanceclient.Record{
		Key:       &opts.key,
		Value:     value,
		Partition: &opts.partition,
	}

	record, _, err = api.RecordsApi.ProduceRecord(opts.f.Context, opts.topicName).Record(record).Execute()
	if err != nil {
		return err
	}

	dump.Table(opts.f.IOStreams.Out, mapRecordToRow(opts.topicName, &record))
	opts.f.Logger.Info("")

	opts.f.Logger.Info(icon.SuccessPrefix(), opts.f.Localizer.MustLocalize("kafka.topic.produce.info.produceSuccess",
		localize.NewEntry("Topic", opts.topicName),
		localize.NewEntry("Offset", record.Offset)))

	return nil
}

func mapRecordToRow(topic string, record *kafkainstanceclient.Record) kafkaRow {
	return kafkaRow{
		Topic:     topic,
		Key:       *record.Key,
		Value:     record.Value,
		Partition: *record.Partition,
		Offset:    *record.Offset,
	}
}
