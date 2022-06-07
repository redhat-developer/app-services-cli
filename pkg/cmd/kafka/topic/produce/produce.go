package produce

import (
	"context"
	"fmt"

	kafkaflagutil "github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"
	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal/client"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/kafkacmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"

	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

type options struct {
	topicName string
	kafkaID   string

	IO             *iostreams.IOStreams
	Connection     factory.ConnectionFunc
	Logger         logging.Logger
	localizer      localize.Localizer
	Context        context.Context
	ServiceContext servicecontext.IContext
}

func NewProduceCommand(f *factory.Factory) *cobra.Command {
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
		Short:   "Produce to topic",
		Long:    "Produce to topic",
		Example: "",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if !opts.IO.CanPrompt() {
				return opts.localizer.MustLocalizeError("flag.error.requiredWhenNonInteractive", localize.NewEntry("Flag", "yes"))
			}

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

	_ = cmd.MarkFlagRequired("name")

	_ = cmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return kafkacmdutil.FilterValidTopicNameArgs(f, toComplete)
	})

	flags.AddInstanceID(&opts.kafkaID)

	return cmd
}

func runCmd(opts *options) error {
	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	api, _, err := conn.API().KafkaAdmin(opts.kafkaID)
	if err != nil {
		return err
	}

	recordToProduce := kafkainstanceclient.Record{

		Value: "Hello topic",
	}

	request := api.RecordsApi.ProduceRecord(opts.Context, opts.topicName)
	request = request.Record(recordToProduce)

	// perform delete topic API request
	record, _, err := request.Execute()

	if err != nil {
		// TODO handle errors
		return err
	}

	opts.Logger.Info(fmt.Sprintf("produced %#v", record))

	return nil
}
