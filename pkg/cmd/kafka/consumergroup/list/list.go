package list

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"gopkg.in/yaml.v2"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	"github.com/redhat-developer/app-services-cli/pkg/cmdutil"
	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/kafka/consumergroup"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal/client"
	"github.com/spf13/cobra"
)

type Options struct {
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     func() (logging.Logger, error)
	IO         *iostreams.IOStreams
	localizer  localize.Localizer

	output  string
	kafkaID string
	limit   int32
	topic   string
}

type consumerGroupRow struct {
	ConsumerGroupID   string `json:"groupId,omitempty" header:"Consumer group ID"`
	ActiveMembers     int    `json:"active_members,omitempty" header:"Active members"`
	PartitionsWithLag int    `json:"lag,omitempty" header:"Partitions with lag"`
}

// NewListConsumerGroupCommand creates a new command to list consumer groups
func NewListConsumerGroupCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
	}

	cmd := &cobra.Command{
		Use:     opts.localizer.MustLocalize("kafka.consumerGroup.list.cmd.use"),
		Short:   opts.localizer.MustLocalize("kafka.consumerGroup.list.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("kafka.consumerGroup.list.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("kafka.consumerGroup.list.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if opts.output != "" && !flagutil.IsValidInput(opts.output, flagutil.ValidOutputFormats...) {
				return flag.InvalidValueError("output", opts.output, flagutil.ValidOutputFormats...)
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			if !cfg.HasKafka() {
				return fmt.Errorf(opts.localizer.MustLocalize("kafka.consumerGroup.common.error.noKafkaSelected"))
			}

			opts.kafkaID = cfg.Services.Kafka.ClusterID

			return runList(opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.limit, "limit", "", 1000, opts.localizer.MustLocalize("kafka.consumerGroup.list.flag.limit"))
	cmd.Flags().StringVarP(&opts.output, "output", "o", "", opts.localizer.MustLocalize("kafka.consumerGroup.list.flag.output.description"))
	cmd.Flags().StringVar(&opts.topic, "topic", "", opts.localizer.MustLocalize("kafka.consumerGroup.list.flag.topic.description"))

	_ = cmd.RegisterFlagCompletionFunc("topic", func(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return cmdutil.FilterValidTopicNameArgs(f, toComplete)
	})

	flagutil.EnableOutputFlagCompletion(cmd)

	return cmd
}

func runList(opts *Options) (err error) {
	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	ctx := context.Background()

	api, kafkaInstance, err := conn.API().KafkaAdmin(opts.kafkaID)
	if err != nil {
		return err
	}

	req := api.GetConsumerGroups(ctx)
	req = req.Limit(opts.limit)
	if opts.topic != "" {
		req = req.Topic(opts.topic)
	}
	consumerGroupData, httpRes, err := req.Execute()
	if err != nil {
		if httpRes == nil {
			return err
		}

		operationTmplPair := localize.NewEntry("Operation", "list")

		switch httpRes.StatusCode {
		case 401:
			return errors.New(opts.localizer.MustLocalize("kafka.consumerGroup.list.common.error.unauthorized", operationTmplPair))
		case 403:
			return errors.New(opts.localizer.MustLocalize("kafka.consumerGroup.list.common.error.forbidden", operationTmplPair))
		case 500:
			return errors.New(opts.localizer.MustLocalize("kafka.consumerGroup.common.error.internalServerError"))
		case 503:
			return errors.New(opts.localizer.MustLocalize("kafka.consumerGroup.common.error.unableToConnectToKafka", localize.NewEntry("Name", kafkaInstance.GetName())))
		default:
			return err
		}
	}

	ok, err := checkForConsumerGroups(int(consumerGroupData.GetTotal()), opts, kafkaInstance.GetName())
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}

	switch opts.output {
	case "json":
		data, _ := json.Marshal(consumerGroupData)
		_ = dump.JSON(opts.IO.Out, data)
	case "yaml", "yml":
		data, _ := yaml.Marshal(consumerGroupData)
		_ = dump.YAML(opts.IO.Out, data)
	default:
		logger.Info("")
		topics := consumerGroupData.GetItems()
		rows := mapConsumerGroupResultsToTableFormat(topics)
		dump.Table(opts.IO.Out, rows)

		return nil
	}

	return nil

}

func mapConsumerGroupResultsToTableFormat(consumerGroups []kafkainstanceclient.ConsumerGroup) []consumerGroupRow {
	var rows []consumerGroupRow = []consumerGroupRow{}

	for _, t := range consumerGroups {
		consumers := t.GetConsumers()
		row := consumerGroupRow{
			ConsumerGroupID:   t.GetGroupId(),
			ActiveMembers:     consumergroup.GetActiveConsumersCount(consumers),
			PartitionsWithLag: consumergroup.GetPartitionsWithLag(consumers),
		}
		rows = append(rows, row)
	}

	return rows
}

// checks if there are any consumer groups available
// prints to stderr if not
func checkForConsumerGroups(count int, opts *Options, kafkaName string) (hasCount bool, err error) {
	logger, err := opts.Logger()
	if err != nil {
		return false, err
	}
	kafkaNameTmplPair := localize.NewEntry("InstanceName", kafkaName)
	if count == 0 && opts.output == "" {
		if opts.topic == "" {
			logger.Info(opts.localizer.MustLocalize("kafka.consumerGroup.list.log.info.noConsumerGroups", kafkaNameTmplPair))
		} else {
			logger.Info(opts.localizer.MustLocalize("kafka.consumerGroup.list.log.info.noConsumerGroupsForTopic", kafkaNameTmplPair, localize.NewEntry("TopicName", opts.topic)))
		}

		return false, nil
	}

	return true, nil
}
