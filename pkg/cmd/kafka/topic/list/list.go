package list

import (
	"context"
	"net/http"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/topic/topiccmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal/client"

	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/internal/build"
)

type options struct {
	Config     config.IConfig
	IO         *iostreams.IOStreams
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	localizer  localize.Localizer
	Context    context.Context

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
	opts := &options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "list",
		Short:   opts.localizer.MustLocalize("kafka.topic.list.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("kafka.topic.list.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("kafka.topic.list.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if opts.output != "" {
				if err := flagutil.ValidateOutput(opts.output); err != nil {
					return err
				}
			}

			if opts.page < 1 {
				return opts.localizer.MustLocalizeError("kafka.common.validation.page.error.invalid.minValue", localize.NewEntry("Page", opts.page))
			}

			if opts.size < 1 {
				return opts.localizer.MustLocalizeError("kafka.common.validation.size.error.invalid.minValue", localize.NewEntry("Size", opts.size))
			}

			if opts.search != "" {
				validator := topiccmdutil.Validator{
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

			instanceID, ok := cfg.GetKafkaIdOk()
			if !ok {
				return opts.localizer.MustLocalizeError("kafka.topic.common.error.noKafkaSelected")
			}

			opts.kafkaID = instanceID

			return runCmd(opts)
		},
	}

	flags := flagutil.NewFlagSet(cmd, opts.localizer)

	flags.AddOutput(&opts.output)
	flags.StringVar(&opts.search, "search", "", opts.localizer.MustLocalize("kafka.topic.list.flag.search.description"))
	flags.Int32VarP(&opts.page, "page", "", cmdutil.ConvertPageValueToInt32(build.DefaultPageNumber), opts.localizer.MustLocalize("kafka.topic.list.flag.page.description"))
	flags.Int32VarP(&opts.size, "size", "", cmdutil.ConvertSizeValueToInt32(build.DefaultPageSize), opts.localizer.MustLocalize("kafka.topic.list.flag.size.description"))

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

	a := api.TopicsApi.GetTopics(opts.Context)

	if opts.search != "" {
		opts.Logger.Debug(opts.localizer.MustLocalize("kafka.topic.list.log.debug.filteringTopicList", localize.NewEntry("Search", opts.search)))
		a = a.Filter(opts.search)
	}

	a = a.Size(opts.size)

	a = a.Page(opts.page)

	topicData, httpRes, err := a.Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}
	if err != nil {
		if httpRes == nil {
			return err
		}

		operationTemplatePair := localize.NewEntry("Operation", "list")

		switch httpRes.StatusCode {
		case http.StatusUnauthorized:
			return opts.localizer.MustLocalizeError("kafka.topic.list.error.unauthorized", operationTemplatePair)
		case http.StatusForbidden:
			return opts.localizer.MustLocalizeError("kafka.topic.list.error.forbidden", operationTemplatePair)
		case http.StatusInternalServerError:
			return opts.localizer.MustLocalizeError("kafka.topic.common.error.internalServerError")
		case http.StatusServiceUnavailable:
			return opts.localizer.MustLocalizeError("kafka.topic.common.error.unableToConnectToKafka", localize.NewEntry("Name", kafkaInstance.GetName()))
		default:
			return err
		}
	}

	if topicData.GetTotal() == 0 && opts.output == "" {
		opts.Logger.Info(opts.localizer.MustLocalize("kafka.topic.list.log.info.noTopics", localize.NewEntry("InstanceName", kafkaInstance.GetName())))

		return nil
	}

	stdout := opts.IO.Out
	switch opts.output {
	case dump.TableFormat:
		topics := topicData.GetItems()
		rows := mapTopicResultsToTableFormat(topics)
		dump.Table(stdout, rows)
	default:
		return dump.Formatted(stdout, opts.output, topicData)
	}

	return nil
}

func mapTopicResultsToTableFormat(topics []kafkainstanceclient.Topic) []topicRow {
	rows := make([]topicRow, len(topics))

	for i, t := range topics {

		row := topicRow{
			Name:            t.GetName(),
			PartitionsCount: len(t.GetPartitions()),
		}
		for _, conf := range t.GetConfig() {
			unlimitedVal := "-1 (Unlimited)"

			if conf.Key == topiccmdutil.RetentionMsKey {
				val := conf.GetValue()
				if val == "-1" {
					row.RetentionTime = unlimitedVal
				} else {
					row.RetentionTime = val
				}
			}
			if conf.Key == topiccmdutil.RetentionSizeKey {
				val := conf.GetValue()
				if val == "-1" {
					row.RetentionSize = unlimitedVal
				} else {
					row.RetentionSize = val
				}
			}
		}

		rows[i] = row
	}

	return rows
}
