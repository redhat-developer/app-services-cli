package produce

import (
	"net/http"
	"os"

	kafkaflagutil "github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal/client"

	"io/ioutil"

	"github.com/AlecAivazis/survey/v2"

	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/kafkacmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

var outputFormatTypes = []string{dump.JSONFormat, dump.YAMLFormat}

type options struct {
	topicName    string
	kafkaID      string
	key          string
	file         string
	outputFormat string
	partition    int32

	f *factory.Factory
}

// row is the details of a record produced needed to print to a table
type kafkaRow struct {
	Topic     string `json:"topic"`
	Key       string `json:"key"`
	Value     string `json:"value"`
	Partition int32  `json:"partition"`
	Offset    int64  `json:"offset"`
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
		Hidden:  true,
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
	flags.StringVar(&opts.outputFormat, "format", "json", f.Localizer.MustLocalize("kafka.topic.produce.flag.format.description"))

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

	var value string

	if opts.file != "" {
		bytes, readError := ioutil.ReadFile(opts.file)
		if readError != nil {
			return readError
		}

		value = string(bytes)
	} else {
		// if value is being piped then cannot have delimeter as \n
		info, _ := os.Stdin.Stat()
		if info.Mode()&os.ModeCharDevice == 0 {
			bytes, readError := ioutil.ReadAll(os.Stdin)
			if readError != nil {
				return readError
			}

			value = string(bytes)
		} else {
			promptName := &survey.Input{
				Message: opts.f.Localizer.MustLocalize("kafka.topic.produce.input.value"),
				Help:    opts.f.Localizer.MustLocalize("kafka.topic.produce.input.help"),
			}

			surveyErr := survey.AskOne(promptName, &value)
			if surveyErr != nil {
				return surveyErr
			}
		}
	}

	record := kafkainstanceclient.Record{
		Key:       &opts.key,
		Value:     value,
		Partition: &opts.partition,
	}

	record, httpRes, err := api.RecordsApi.ProduceRecord(opts.f.Context, opts.topicName).Record(record).Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	if err != nil {

		if httpRes == nil {
			return err
		}

		if httpRes.StatusCode == http.StatusRequestEntityTooLarge {
			return opts.f.Localizer.MustLocalizeError("kafka.topic.produce.error.messageTooLarge")
		}

		return err
	}

	format := opts.outputFormat
	if format == dump.EmptyFormat {
		format = dump.JSONFormat
	}

	_ = dump.Formatted(opts.f.IOStreams.Out, format, mapRecordToRow(opts.topicName, &record))
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
