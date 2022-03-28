package status

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/redhat-developer/app-services-cli/pkg/shared/servicespec"

	"github.com/spf13/cobra"
)

type options struct {
	IO             *iostreams.IOStreams
	Config         config.IConfig
	Logger         logging.Logger
	Connection     factory.ConnectionFunc
	localizer      localize.Localizer
	Context        context.Context
	ServiceContext servicecontext.IContext

	outputFormat string
	name         string
	services     []string
}

func NewStatusCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		IO:             f.IOStreams,
		Config:         f.Config,
		Connection:     f.Connection,
		Logger:         f.Logger,
		services:       servicespec.AllServiceLabels,
		localizer:      f.Localizer,
		Context:        f.Context,
		ServiceContext: f.ServiceContext,
	}

	cmd := &cobra.Command{
		Use:       "status [args]",
		Short:     opts.localizer.MustLocalize("status.cmd.shortDescription"),
		Long:      opts.localizer.MustLocalize("status.cmd.longDescription"),
		Example:   opts.localizer.MustLocalize("status.cmd.example"),
		ValidArgs: servicespec.AllServiceLabels,
		Args:      cobra.RangeArgs(0, len(servicespec.AllServiceLabels)),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				for _, s := range args {
					if !flagutil.IsValidInput(s, servicespec.AllServiceLabels...) {
						return opts.localizer.MustLocalizeError("status.error.args.error.unknownServiceError", localize.NewEntry("ServiceName", s))
					}
				}

				opts.services = args
			}

			validOutputFormats := flagutil.ValidOutputFormats
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flagutil.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			return runStatus(opts)
		},
	}

	flags := flagutil.NewFlagSet(cmd, opts.localizer)

	flags.StringVar(&opts.name, "name", "", opts.localizer.MustLocalize("context.common.flag.name"))
	flags.AddOutput(&opts.outputFormat)

	flagutil.EnableOutputFlagCompletion(cmd)

	return cmd
}

func runStatus(opts *options) error {
	conn, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	if len(opts.services) > 0 {
		opts.Logger.Debug(opts.localizer.MustLocalize("status.log.debug.requestingStatusOfServices"), opts.services)
	}

	svcContext, err := opts.ServiceContext.Load()
	if err != nil {
		return err
	}

	profileHandler := &contextutil.ContextHandler{
		Context:   svcContext,
		Localizer: opts.localizer,
	}

	var svcConfig *servicecontext.ServiceConfig

	if opts.name == "" {
		opts.name, err = profileHandler.GetCurrentContext()
		if err != nil {
			return err
		}
	}

	svcConfig, err = profileHandler.GetContext(opts.name)
	if err != nil {
		return err
	}

	statusClient := newStatusClient(&clientConfig{
		config:        opts.Config,
		connection:    conn,
		Logger:        opts.Logger,
		localizer:     opts.localizer,
		context:       opts.Context,
		serviceConfig: svcConfig,
	})

	status, ok, err := statusClient.BuildStatus(opts.name, opts.services)
	if err != nil {
		return err
	}

	if !ok {
		opts.Logger.Info("")
		opts.Logger.Info(opts.localizer.MustLocalize("status.log.info.noStatusesAreUsed"))
		return nil
	}

	stdout := opts.IO.Out
	if opts.outputFormat != "" {
		if err = dump.Formatted(stdout, opts.outputFormat, status); err != nil {
			return err
		}
	} else {
		Print(stdout, status)
	}

	return nil
}
