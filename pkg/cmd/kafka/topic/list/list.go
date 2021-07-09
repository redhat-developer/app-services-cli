package list

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/redhat-developer/app-services-cli/pkg/cmdutil"
	topicutil "github.com/redhat-developer/app-services-cli/pkg/kafka/topic"
	"github.com/redhat-developer/app-services-cli/pkg/localize"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	"github.com/redhat-developer/app-services-cli/pkg/connection"

	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal/client"

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
	search  string
	page    int32
	size    int32
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

			if opts.page < 1 {
				return errors.New(opts.localizer.MustLocalize("kafka.topic.list.validation.page.error.invalid.minValue", localize.NewEntry("Page", opts.page)))
			}

			if opts.size < 1 {
				return errors.New(opts.localizer.MustLocalize("kafka.topic.list.validation.size.error.invalid.minValue", localize.NewEntry("Size", opts.size)))
			}

			if opts.search != "" {
				validator := topicutil.Validator{
					Localizer: opts.localizer,
				}
				if err := validator.ValidateSearchInput(opts.search); err != nil {
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
	cmd.Flags().StringVarP(&opts.search, "search", "", "", opts.localizer.MustLocalize("kafka.topic.list.flag.search.description"))
	cmd.Flags().Int32VarP(&opts.page, "page", "", int32(cmdutil.DefaultPageNumber), opts.localizer.MustLocalize("kafka.topic.list.flag.page.description"))
	cmd.Flags().Int32VarP(&opts.size, "size", "", int32(cmdutil.DefaultPageSize), opts.localizer.MustLocalize("kafka.topic.list.flag.size.description"))

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

	api, kafkaInstance, err := conn.API().KafkaAdmin(opts.kafkaID)
	if err != nil {
		return err
	}

	a := api.GetTopics(context.Background())

	if opts.search != "" {
		logger.Debug(opts.localizer.MustLocalize("kafka.topic.list.log.debug.filteringTopicList", localize.NewEntry("Search", opts.search)))
		a = a.Filter(opts.search)
	}

	a = a.Size(opts.size)

	a = a.Page(opts.page)

	topicData, httpRes, err := a.Execute()
	if err != nil {
		if httpRes == nil {
			return err
		}

		operationTemplatePair := localize.NewEntry("Operation", "list")

		switch httpRes.StatusCode {
		case http.StatusUnauthorized:
			return errors.New(opts.localizer.MustLocalize("kafka.topic.list.error.unauthorized", operationTemplatePair))
		case http.StatusForbidden:
			return errors.New(opts.localizer.MustLocalize("kafka.topic.list.error.forbidden", operationTemplatePair))
		case http.StatusInternalServerError:
			return errors.New(opts.localizer.MustLocalize("kafka.topic.common.error.internalServerError"))
		case http.StatusServiceUnavailable:
			return errors.New(opts.localizer.MustLocalize("kafka.topic.common.error.unableToConnectToKafka", localize.NewEntry("Name", kafkaInstance.GetName())))
		default:
			return err
		}
	}

	defer httpRes.Body.Close()

	if topicData.GetTotal() == 0 && opts.output == "" {
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

func mapTopicResultsToTableFormat(topics []kafkainstanceclient.Topic) []topicRow {
	rows := []topicRow{}

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
