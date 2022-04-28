package status

import (
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/redhat-developer/app-services-cli/pkg/shared/servicespec"

	"github.com/spf13/cobra"
)

type options struct {
	f *factory.Factory

	outputFormat string
	name         string
	services     []string
}

func NewStatusCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		f: f,
	}

	cmd := &cobra.Command{
		Use:       "status [args]",
		Short:     f.Localizer.MustLocalize("status.cmd.shortDescription"),
		Long:      f.Localizer.MustLocalize("status.cmd.longDescription"),
		Example:   f.Localizer.MustLocalize("status.cmd.example"),
		ValidArgs: servicespec.AllServiceLabels,
		Args:      cobra.RangeArgs(0, len(servicespec.AllServiceLabels)),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				for _, s := range args {
					if !flagutil.IsValidInput(s, servicespec.AllServiceLabels...) {
						return f.Localizer.MustLocalizeError("common.error.args.error.unknownServiceError", localize.NewEntry("ServiceName", s))
					}
				}

				opts.services = args
			} else {
				opts.services = servicespec.AllServiceLabels
			}

			validOutputFormats := flagutil.ValidOutputFormats
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flagutil.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			return runStatus(opts)
		},
	}

	flags := flagutil.NewFlagSet(cmd, f.Localizer)

	flags.StringVar(&opts.name, "name", "", f.Localizer.MustLocalize("context.common.flag.name"))
	flags.AddOutput(&opts.outputFormat)

	flagutil.EnableOutputFlagCompletion(cmd)

	return cmd
}

func runStatus(opts *options) error {
	factory := opts.f

	if len(opts.services) > 0 {
		opts.f.Logger.Debug(factory.Localizer.MustLocalize("status.log.debug.requestingStatusOfServices"), opts.services)
	}

	svcContext, err := factory.ServiceContext.Load()
	if err != nil {
		return err
	}

	var svcConfig *servicecontext.ServiceConfig

	if opts.name == "" {
		svcConfig, err = contextutil.GetCurrentContext(svcContext, opts.f.Localizer)
		if err != nil {
			return err
		}
	} else {
		svcConfig, err = contextutil.GetContext(svcContext, opts.f.Localizer, opts.name)
		if err != nil {
			return err
		}
	}

	statusClient := newStatusClient(&clientConfig{
		f:             factory,
		serviceConfig: svcConfig,
	})

	status, err := statusClient.BuildStatus(opts.services)
	if err != nil {
		return err
	}
	status.Name = svcContext.CurrentContext
	status.Location, _ = factory.ServiceContext.Location()

	if !status.hasStatus() {
		factory.Logger.Info("")
		factory.Logger.Info(factory.Localizer.MustLocalize("status.log.info.noStatusesAreUsed"))
		return nil
	}

	stdout := factory.IOStreams.Out
	if opts.outputFormat != "" {
		if err = dump.Formatted(stdout, opts.outputFormat, status); err != nil {
			return err
		}
	} else {
		Print(stdout, status)
	}

	return nil
}
