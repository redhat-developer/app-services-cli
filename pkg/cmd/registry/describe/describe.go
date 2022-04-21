package describe

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/registrycmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/redhat-developer/app-services-cli/pkg/shared/serviceregistryutil"
	srsmgmtv1 "github.com/redhat-developer/app-services-sdk-go/registrymgmt/apiv1/client"

	"github.com/spf13/cobra"
)

type options struct {
	id           string
	name         string
	outputFormat string

	IO             *iostreams.IOStreams
	Connection     factory.ConnectionFunc
	localizer      localize.Localizer
	logger         logging.Logger
	Context        context.Context
	ServiceContext servicecontext.IContext
}

// NewDescribeCommand describes a service instance, either by passing an `--id flag`
// or by using the service instance set in the current context, if any
func NewDescribeCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Connection:     f.Connection,
		IO:             f.IOStreams,
		localizer:      f.Localizer,
		Context:        f.Context,
		logger:         f.Logger,
		ServiceContext: f.ServiceContext,
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

			registryInstance, err := contextutil.GetCurrentRegistryInstance(f)
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

	if err := dump.Formatted(opts.IO.Out, opts.outputFormat, registry); err != nil {
		return err
	}

	compatibleEndpoints := registrycmdutil.GetCompatibilityEndpoints(registry.GetRegistryUrl())

	opts.logger.Error(opts.localizer.MustLocalize(
		"registry.common.log.message.compatibleAPIs",
		localize.NewEntry("CoreRegistryAPI", compatibleEndpoints.CoreRegistry),
		localize.NewEntry("SchemaRegistryAPI", compatibleEndpoints.SchemaRegistry),
		localize.NewEntry("CncfSchemaRegistryAPI", compatibleEndpoints.CncfSchemaRegistry),
	))

	return nil
}
