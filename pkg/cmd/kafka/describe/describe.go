package describe

import (
	"context"
	"fmt"
	"net/http"

	kafkaflagutil "github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/redhat-developer/app-services-cli/pkg/shared/kafkautil"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
	"github.com/spf13/cobra"
)

type options struct {
	id              string
	name            string
	bootstrapServer bool
	outputFormat    string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	localizer  localize.Localizer
	Context    context.Context
}

// NewDescribeCommand describes a Kafka instance, either by passing an `--id flag`
// or by using the kafka instance set in the config, if any
func NewDescribeCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Config:     f.Config,
		Connection: f.Connection,
		IO:         f.IOStreams,
		Logger:     f.Logger,
		localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "describe",
		Short:   opts.localizer.MustLocalize("kafka.describe.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("kafka.describe.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("kafka.describe.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			validOutputFormats := flagutil.ValidOutputFormats
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flagutil.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			if opts.name != "" && opts.id != "" {
				return opts.localizer.MustLocalizeError("service.error.idAndNameCannotBeUsed")
			}

			if opts.id != "" || opts.name != "" {
				return runDescribe(opts)
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			instanceID, ok := cfg.GetKafkaIdOk()
			if !ok {
				return opts.localizer.MustLocalizeError("kafka.common.error.noKafkaSelected")
			}
			opts.id = instanceID

			opts.id = cfg.Services.Kafka.ClusterID

			return runDescribe(opts)
		},
	}

	flags := flagutil.NewFlagSet(cmd, opts.localizer)

	flags.AddOutput(&opts.outputFormat)
	flags.StringVar(&opts.id, "id", "", opts.localizer.MustLocalize("kafka.describe.flag.id"))
	flags.StringVar(&opts.name, "name", "", opts.localizer.MustLocalize("kafka.describe.flag.name"))
	flags.BoolVar(&opts.bootstrapServer, "bootstrap-server", false, opts.localizer.MustLocalize("kafka.describe.flag.bootstrapserver"))

	if err := kafkaflagutil.RegisterNameFlagCompletionFunc(cmd, f); err != nil {
		opts.Logger.Debug(opts.localizer.MustLocalize("kafka.common.error.load.completions.name.flag"), err)
	}
	flagutil.EnableOutputFlagCompletion(cmd)

	return cmd
}

func runDescribe(opts *options) error {
	conn, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	api := conn.API()

	var kafkaInstance *kafkamgmtclient.KafkaRequest
	var httpRes *http.Response
	if opts.name != "" {
		kafkaInstance, httpRes, err = kafkautil.GetKafkaByName(opts.Context, api.KafkaMgmt(), opts.name)
		if httpRes != nil {
			defer httpRes.Body.Close()
		}
		if err != nil {
			return err
		}
	} else {
		kafkaInstance, httpRes, err = kafkautil.GetKafkaByID(opts.Context, api.KafkaMgmt(), opts.id)
		if httpRes != nil {
			defer httpRes.Body.Close()
		}
		if err != nil {
			return err
		}
	}

	if opts.bootstrapServer {
		if host, ok := kafkaInstance.GetBootstrapServerHostOk(); ok {
			fmt.Fprintln(opts.IO.Out, *host)
			return nil
		}
		opts.Logger.Info(opts.localizer.MustLocalize("kafka.describe.bootstrapserver.not.available", localize.NewEntry("Name", kafkaInstance.GetName())))
		return nil
	}

	return dump.Formatted(opts.IO.Out, opts.outputFormat, kafkaInstance)
}
