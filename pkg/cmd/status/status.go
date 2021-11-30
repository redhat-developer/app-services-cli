package status

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/factory"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/connection"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	pkgStatus "github.com/redhat-developer/app-services-cli/pkg/core/status"

	"github.com/spf13/cobra"
)

const (
	kafkaSvcName    = "kafka"
	registrySvcName = "service-registry"
)

var validServices = []string{kafkaSvcName, registrySvcName}

type options struct {
	IO         *iostreams.IOStreams
	Config     config.IConfig
	Logger     logging.Logger
	Connection factory.ConnectionFunc
	localizer  localize.Localizer
	Context    context.Context

	outputFormat string
	services     []string
}

func NewStatusCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		IO:         f.IOStreams,
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		services:   validServices,
		localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:       "status [args]",
		Short:     opts.localizer.MustLocalize("status.cmd.shortDescription"),
		Long:      opts.localizer.MustLocalize("status.cmd.longDescription"),
		Example:   opts.localizer.MustLocalize("status.cmd.example"),
		ValidArgs: validServices,
		Args:      cobra.RangeArgs(0, len(validServices)),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				for _, s := range args {
					if !flagutil.IsValidInput(s, validServices...) {
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

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "", opts.localizer.MustLocalize("status.flag.output.description"))

	flagutil.EnableOutputFlagCompletion(cmd)

	return cmd
}

func runStatus(opts *options) error {
	conn, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	pkgOpts := &pkgStatus.Options{
		Config:     opts.Config,
		Connection: conn,
		Logger:     opts.Logger,
		Services:   opts.services,
	}

	if len(opts.services) > 0 {
		opts.Logger.Debug(opts.localizer.MustLocalize("status.log.debug.requestingStatusOfServices"), opts.services)
	}

	status, ok, err := pkgStatus.Get(opts.Context, pkgOpts)
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
		pkgStatus.Print(stdout, status)
	}

	return nil
}
