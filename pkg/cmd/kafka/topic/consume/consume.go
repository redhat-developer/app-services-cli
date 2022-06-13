package consume

import (
	kafkaflagutil "github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/kafkacmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
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

	list, _, err := api.RecordsApi.ConsumeRecords(opts.f.Context, opts.topicName).Limit(opts.limit).Partition(opts.partition).Execute()
	if err != nil {
		return err
	}

	for i := int32(0); i < list.Total; i++ {
		opts.f.Logger.Info(list.Items[i].Value)
	}

	return nil
}
