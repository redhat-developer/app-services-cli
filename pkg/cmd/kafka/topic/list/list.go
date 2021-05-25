package list

import (
	"context"
	"encoding/json"
	"errors"

	topicutil "github.com/redhat-developer/app-services-cli/pkg/kafka/topic"
	"github.com/redhat-developer/app-services-cli/pkg/localize"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	"github.com/redhat-developer/app-services-cli/pkg/connection"

	strimziadminclient "github.com/redhat-developer/app-services-cli/pkg/api/strimzi-admin/client"

	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"

	"gopkg.in/yaml.v2"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	"github.com/spf13/cobra"
)

type Options struct {
	Config     config.IConfig
	IO         *iostreams.IOStreams
	Connection factory.ConnectionFunc
	Logger     func() (logging.Logger, error)
	localizer  localize.Localizer

	kafkaID string
	output  string
}

type topicRow struct {
	Name            string `json:"name,omitempty" header:"Name"`
	PartitionsCount int    `json:"partitions_count,omitempty" header:"Partitions"`
	RetentionTime   string `json:"retention.ms,omitempty" header:"Retention time (ms)"`
	RetentionSize   string `json:"retention.bytes,omitempty" header:"Retention size (bytes)"`
}

// NewListTopicCommand gets a new command for getting kafkas.
func NewListTopicCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
	}

	cmd := &cobra.Command{
		Use:     opts.localizer.MustLocalize("kafka.topic.list.cmd.use"),
		Short:   opts.localizer.MustLocalize("kafka.topic.list.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("kafka.topic.list.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("kafka.topic.list.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if opts.output != "" {
				if err := flag.ValidateOutput(opts.output); err != nil {
					return err
				}
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			if !cfg.HasKafka() {
				return errors.New(opts.localizer.MustLocalize("kafka.topic.common.error.noKafkaSelected"))
			}

			opts.kafkaID = cfg.Services.Kafka.ClusterID

			return runCmd(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.output, "output", "o", "", opts.localizer.MustLocalize("kafka.topic.list.flag.output.description"))

	flagutil.EnableOutputFlagCompletion(cmd)

	return cmd
}

func runCmd(opts *Options) error {
	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	api, kafkaInstance, err := conn.API().TopicAdmin(opts.kafkaID)
	if err != nil {
		return err
	}

	a := api.GetTopicsList(context.Background())
	topicData, httpRes, err := a.Execute()

	if err != nil {
		if httpRes == nil {
			return err
		}

		operationTemplatePair := localize.NewEntry("Operation", "list")
		switch httpRes.StatusCode {
		case 401:
			return errors.New(opts.localizer.MustLocalize("kafka.topic.list.error.unauthorized", operationTemplatePair))
		case 403:
			return errors.New(opts.localizer.MustLocalize("kafka.topic.list.error.forbidden", operationTemplatePair))
		case 500:
			return errors.New(opts.localizer.MustLocalize("kafka.topic.common.error.internalServerError"))
		case 503:
			return errors.New(opts.localizer.MustLocalize("kafka.topic.common.error.unableToConnectToKafka", localize.NewEntry("Name", kafkaInstance.GetName())))
		default:
			return err
		}
	}

	if topicData.GetCount() == 0 && opts.output == "" {
		logger.Info(opts.localizer.MustLocalize("kafka.topic.list.log.info.noTopics", localize.NewEntry("InstanceName", kafkaInstance.GetName())))

		return nil
	}

	stdout := opts.IO.Out
	switch opts.output {
	case "json":
		data, _ := json.Marshal(topicData)
		_ = dump.JSON(stdout, data)
	case "yaml", "yml":
		data, _ := yaml.Marshal(topicData)
		_ = dump.YAML(stdout, data)
	default:
		topics := topicData.GetItems()
		rows := mapTopicResultsToTableFormat(topics)
		dump.Table(stdout, rows)
	}

	return nil
}

func mapTopicResultsToTableFormat(topics []strimziadminclient.Topic) []topicRow {
	var rows []topicRow = []topicRow{}

	for _, t := range topics {

		row := topicRow{
			Name:            t.GetName(),
			PartitionsCount: len(t.GetPartitions()),
		}
		for _, config := range t.GetConfig() {
			unlimitedVal := "-1 (Unlimited)"

			if *config.Key == topicutil.RetentionMsKey {
				val := config.GetValue()
				if val == "-1" {
					row.RetentionTime = unlimitedVal
				} else {
					row.RetentionTime = val
				}
			}
			if *config.Key == topicutil.RetentionSizeKey {
				val := config.GetValue()
				if val == "-1" {
					row.RetentionSize = unlimitedVal
				} else {
					row.RetentionSize = val
				}
			}
		}

		rows = append(rows, row)
	}

	return rows
}
