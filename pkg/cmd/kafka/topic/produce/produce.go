package produce

import (
	"context"
	"os"

	kafkaflagutil "github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"
	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal/client"

	"bufio"
	"io/ioutil"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/kafkacmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

type options struct {
	topicName string
	kafkaID   string
	key       string
	partition int32

	IO             *iostreams.IOStreams
	Connection     factory.ConnectionFunc
	Logger         logging.Logger
	localizer      localize.Localizer
	Context        context.Context
	ServiceContext servicecontext.IContext
}

// NewProduceTopicCommand gets a new command for producing to a kafka topic.
func NewProduceTopicCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Connection:     f.Connection,
		Logger:         f.Logger,
		IO:             f.IOStreams,
		localizer:      f.Localizer,
		Context:        f.Context,
		ServiceContext: f.ServiceContext,
	}

	cmd := &cobra.Command{
		Use:     "produce",
		Short:   opts.localizer.MustLocalize("kafka.topic.produce.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("kafka.topic.produce.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("kafka.topic.produce.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// if !opts.IO.CanPrompt() {
			// 	return opts.localizer.MustLocalizeError("flag.error.requiredWhenNonInteractive", localize.NewEntry("Flag", "yes"))
			// }

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

	flags := kafkaflagutil.NewFlagSet(cmd, opts.localizer)

	flags.StringVar(&opts.topicName, "name", "", opts.localizer.MustLocalize("kafka.topic.common.flag.name.description"))
	flags.StringVar(&opts.key, "key", "", opts.localizer.MustLocalize("kafka.topic.produce.flag.key.description"))
	flags.Int32Var(&opts.partition, "partition", 0, opts.localizer.MustLocalize("kafka.topic.produce.flag.partition.description"))

	_ = cmd.MarkFlagRequired("name")

	_ = cmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return kafkacmdutil.FilterValidTopicNameArgs(f, toComplete)
	})

	flags.AddInstanceID(&opts.kafkaID)

	return cmd
}

// nolint:funlen
func runCmd(opts *options) error {
	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	api, _, err := conn.API().KafkaAdmin(opts.kafkaID)
	if err != nil {
		return err
	}

	var value string

	// if value being piped then cannot have delimeter as \n
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

	record := kafkainstanceclient.Record{
		Key:       &opts.key,
		Value:     value,
		Partition: &opts.partition,
	}

	_, _, err = api.RecordsApi.ProduceRecord(opts.Context, opts.topicName).Record(record).Execute()
	if err != nil {
		return err
	}

	return nil
}
