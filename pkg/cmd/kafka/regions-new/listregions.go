package regions

import (
	"context"
	"fmt"

	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
	"github.com/spf13/cobra"
)

type options struct {
	id              string
	name            string
	bootstrapServer bool
	outputFormat    string

	IO             *iostreams.IOStreams
	Connection     factory.ConnectionFunc
	Logger         logging.Logger
	localizer      localize.Localizer
	Context        context.Context
	ServiceContext servicecontext.IContext
}

// NewDescribeCommand describes a Kafka instance, either by passing an `--id flag`
// or by using the kafka instance set in the current context, if any
func NewlistregionsCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Connection:     f.Connection,
		IO:             f.IOStreams,
		Logger:         f.Logger,
		localizer:      f.Localizer,
		Context:        f.Context,
		ServiceContext: f.ServiceContext,
	}

	cmd := &cobra.Command{
		Use:     "list-regions",
		Short:   opts.localizer.MustLocalize("kafka.list-regions.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("kafka.list-regions.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("kafka.list-regions.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			validOutputFormats := flagutil.ValidOutputFormats
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flagutil.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			kafkaInstance, err := contextutil.GetCurrentKafkaInstance(f)
			if err != nil {
				return err
			}

			opts.id = kafkaInstance.GetId()
			return runListRegions(opts)
		},
	}

	flags := flagutil.NewFlagSet(cmd, opts.localizer)

	table := dump.TableFormat
	flags.AddOutputFormatted(&opts.outputFormat, true, &table)

	flagutil.EnableOutputFlagCompletion(cmd)

	return cmd
}

func runListRegions(opts *options) error {
	conn, err := opts.Connection(connection.DefaultConfigSkipMasAuth)

	if err != nil {
		return err
	}

	api := conn.API()

	var kafkaInstance *kafkamgmtclient.KafkaRequest

	providers, _, err := api.KafkaMgmt().GetCloudProviders(opts.Context).Execute()

	if opts.bootstrapServer {
		if host, ok := kafkaInstance.GetBootstrapServerHostOk(); ok {
			fmt.Fprintln(opts.IO.Out, *host)
			return nil
		}
		opts.Logger.Info(opts.localizer.MustLocalize("kafka.describe.bootstrapserver.not.available", localize.NewEntry("Name", kafkaInstance.GetName())))
		return nil
	}

	return dump.Formatted(opts.IO.Out, opts.outputFormat, providers)
}
