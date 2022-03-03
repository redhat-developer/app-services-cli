package describe

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/redhat-developer/app-services-cli/pkg/shared/profileutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/serviceregistryutil"
	srsmgmtv1 "github.com/redhat-developer/app-services-sdk-go/registrymgmt/apiv1/client"

	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/pkg/core/profile"
)

type options struct {
	id           string
	name         string
	outputFormat string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	localizer  localize.Localizer
	Context    context.Context
	Profiles   profile.IContext
}

// NewDescribeCommand describes a service instance, either by passing an `--id flag`
// or by using the service instance set in the config, if any
func NewDescribeCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Config:     f.Config,
		Connection: f.Connection,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
		Context:    f.Context,
		Profiles:   f.Profile,
	}

	cmd := &cobra.Command{
		Use:     "describe",
		Short:   f.Localizer.MustLocalize("registry.cmd.describe.shortDescription"),
		Long:    f.Localizer.MustLocalize("registry.cmd.describe.longDescription"),
		Example: f.Localizer.MustLocalize("registry.cmd.describe.example"),
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

			context, err := opts.Profiles.Load()
			if err != nil {
				return err
			}

			profileHandler := &profileutil.ContextHandler{
				Context:   context,
				Localizer: opts.localizer,
			}

			conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
			if err != nil {
				return err
			}

			registryInstance, err := profileHandler.GetCurrentRegistryInstance(conn.API().ServiceRegistryMgmt())
			if err != nil {
				return err
			}

			opts.id = registryInstance.GetId()

			return runDescribe(opts)
		},
	}

	cmd.Flags().StringVar(&opts.name, "name", "", opts.localizer.MustLocalize("registry.cmd.describe.flag.name.description"))
	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "json", opts.localizer.MustLocalize("registry.cmd.flag.output.description"))
	cmd.Flags().StringVar(&opts.id, "id", "", opts.localizer.MustLocalize("registry.describe.flag.id"))

	flagutil.EnableOutputFlagCompletion(cmd)

	return cmd
}

func runDescribe(opts *options) error {
	conn, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	api := conn.API()

	var registry *srsmgmtv1.Registry
	if opts.name != "" {
		registry, _, err = serviceregistryutil.GetServiceRegistryByName(opts.Context, api.ServiceRegistryMgmt(), opts.name)
		if err != nil {
			return err
		}
	} else {
		registry, _, err = serviceregistryutil.GetServiceRegistryByID(opts.Context, api.ServiceRegistryMgmt(), opts.id)
		if err != nil {
			return err
		}
	}

	return dump.Formatted(opts.IO.Out, opts.outputFormat, registry)
}
