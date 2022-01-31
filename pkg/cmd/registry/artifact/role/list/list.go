package list

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/registrycmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	registryinstanceclient "github.com/redhat-developer/app-services-sdk-go/registryinstance/apiv1internal/client"

	"github.com/spf13/cobra"
)

// row is the details of a Service Registry instance needed to print to a table
type registryRow struct {
	Principal string `json:"principal" header:"Principal"`
	Role      string `json:"role" header:"Role"`
}

type options struct {
	outputFormat string
	registryID   string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	localizer  localize.Localizer
	Context    context.Context
}

// NewListCommand creates a new command for listing principal roles
func NewListCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "list",
		Short:   f.Localizer.MustLocalize("registry.role.cmd.list.shortDescription"),
		Long:    f.Localizer.MustLocalize("registry.role.cmd.list.longDescription"),
		Example: f.Localizer.MustLocalize("registry.role.cmd.list.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, flagutil.ValidOutputFormats...) {
				return flagutil.InvalidValueError("output", opts.outputFormat, flagutil.ValidOutputFormats...)
			}

			if opts.registryID != "" {
				return runList(opts)
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			instanceID, ok := cfg.GetServiceRegistryIdOk()
			if !ok {
				return opts.localizer.MustLocalizeError("artifact.cmd.common.error.noServiceRegistrySelected")
			}

			opts.registryID = instanceID

			return runList(opts)
		},
	}

	cmd.Flags().StringVar(&opts.registryID, "instance-id", "", opts.localizer.MustLocalize("artifact.common.instance.id"))
	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "", opts.localizer.MustLocalize("artifact.common.message.output.format"))

	flagutil.EnableOutputFlagCompletion(cmd)

	return cmd
}

func runList(opts *options) error {

	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	api := conn.API()

	a, _, err := api.ServiceRegistryInstance(opts.registryID)
	if err != nil {
		return err
	}
	mappings, _, err := a.AdminApi.ListRoleMappings(opts.Context).Execute()
	if err != nil {
		return registrycmdutil.TransformInstanceError(err)
	}

	if len(mappings) == 0 && opts.outputFormat == "" {
		opts.Logger.Info(opts.localizer.MustLocalize("registry.role.cmd.nomappings", localize.NewEntry("Registry", opts.registryID)))
		return nil
	}

	stdout := opts.IO.Out

	switch opts.outputFormat {
	case dump.EmptyFormat:
		rows := mapResponseItemsToRows(mappings)
		dump.Table(opts.IO.Out, rows)
		opts.Logger.Info("")
	default:
		return dump.Formatted(stdout, opts.outputFormat, mappings)
	}

	return nil
}

func mapResponseItemsToRows(artifacts []registryinstanceclient.RoleMapping) []registryRow {
	rows := []registryRow{}

	for i := range artifacts {
		k := (artifacts)[i]
		row := registryRow{
			Principal: k.GetPrincipalId(),
			Role:      string(k.GetRole()),
		}

		rows = append(rows, row)
	}

	return rows
}
