package unset

import (
	"github.com/redhat-developer/app-services-cli/pkg/cmd/context/contextcmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/redhat-developer/app-services-cli/pkg/shared/servicespec"
	"github.com/spf13/cobra"
)

type options struct {
	f *factory.Factory

	services []string
	name     string
}

// NewUnsetCommand creates a new command to unset services in current command
func NewUnsetCommand(f *factory.Factory) *cobra.Command {

	opts := &options{
		f: f,
	}

	cmd := &cobra.Command{
		Use:     "unset",
		Short:   f.Localizer.MustLocalize("context.unset.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("context.unset.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("context.unset.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {

			if len(opts.services) == 0 {
				return opts.f.Localizer.MustLocalizeError("context.unset.cmd.error.noServices")
			}

			for _, s := range opts.services {
				if !flagutil.IsValidInput(s, servicespec.AllServiceLabels...) {
					return f.Localizer.MustLocalizeError("common.error.args.error.unknownServiceError", localize.NewEntry("ServiceName", s))
				}
			}

			return runUnset(opts)
		},
	}

	flags := contextcmdutil.NewFlagSet(cmd, f)

	flags.AddContextName(&opts.name)
	flags.StringSliceVar(&opts.services, "services", []string{}, "context.unset.flag.services.description")

	return cmd

}

func runUnset(opts *options) error {

	svcContext, err := opts.f.ServiceContext.Load()
	if err != nil {
		return err
	}

	var svcConfig *servicecontext.ServiceConfig
	var ctxName string

	if opts.name == "" {
		svcConfig, err = contextutil.GetCurrentContext(svcContext, opts.f.Localizer)
		if err != nil {
			return err
		}
		ctxName = svcContext.CurrentContext
	} else {
		svcConfig, err = contextutil.GetContext(svcContext, opts.f.Localizer, opts.name)
		if err != nil {
			return err
		}
		ctxName = opts.name
	}

	if flagutil.StringInSlice(servicespec.KafkaServiceName, opts.services) {
		svcConfig.KafkaID = ""
	}

	if flagutil.StringInSlice(servicespec.ServiceRegistryServiceName, opts.services) {
		svcConfig.ServiceRegistryID = ""
	}

	svcContext.Contexts[ctxName] = *svcConfig

	err = opts.f.ServiceContext.Save(svcContext)
	if err != nil {
		return err
	}

	opts.f.Logger.Info(icon.SuccessPrefix(), opts.f.Localizer.MustLocalize("context.unset.log.info.success"))

	return nil

}
