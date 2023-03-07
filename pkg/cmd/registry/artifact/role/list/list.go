package list

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact/util"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/registrycmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	registryinstanceclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal/client"

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

	IO             *iostreams.IOStreams
	Connection     factory.ConnectionFunc
	Logger         logging.Logger
	localizer      localize.Localizer
	Context        context.Context
	ServiceContext servicecontext.IContext
}

// NewListCommand creates a new command for listing principal roles
func NewListCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Connection:     f.Connection,
		Logger:         f.Logger,
		IO:             f.IOStreams,
		localizer:      f.Localizer,
		Context:        f.Context,
		ServiceContext: f.ServiceContext,
	}

	cmd := &cobra.Command{
		Use:     "list",
		Short:   f.Localizer.MustLocalize("registry.role.cmd.list.shortDescription"),
		Long:    f.Localizer.MustLocalize("registry.role.cmd.list.longDescription"),
		Example: f.Localizer.MustLocalize("registry.role.cmd.list.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.registryID != "" {
				return runList(opts)
			}

			registryInstance, err := contextutil.GetCurrentRegistryInstance(f)
			if err != nil {
				return err
			}

			opts.registryID = registryInstance.GetId()

			return runList(opts)
		},
	}

	cmd.Flags().StringVar(&opts.registryID, "instance-id", "", opts.localizer.MustLocalize("registry.common.flag.instance.id"))
	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "table", opts.localizer.MustLocalize("artifact.common.message.output.format"))

	flagutil.EnableOutputFlagCompletion(cmd)

	return cmd
}

func runList(opts *options) error {
	format := util.OutputFormatFromString(opts.outputFormat)
	if format == util.UnknownOutputFormat {
		return opts.localizer.MustLocalizeError("artifact.common.error.invalidOutputFormat")
	}

	conn, err := opts.Connection()
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

	if len(mappings) == 0 && format == util.TableOutputFormat {
		opts.Logger.Info(opts.localizer.MustLocalize("registry.role.cmd.nomappings", localize.NewEntry("Registry", opts.registryID)))
		return nil
	}

	return util.Dump(opts.IO.Out, format, mapResponseItemsToRows(mappings), mappings)
}

func mapResponseItemsToRows(artifacts []registryinstanceclient.RoleMapping) []registryRow {
	rows := []registryRow{}

	for i := range artifacts {
		k := (artifacts)[i]
		row := registryRow{
			Principal: k.GetPrincipalId(),
			Role:      util.GetRoleLabel(k.GetRole()),
		}

		rows = append(rows, row)
	}

	return rows
}
