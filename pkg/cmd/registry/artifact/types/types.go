package types

import (
	"context"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/registrycmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	registryinstanceclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal/client"
	"github.com/spf13/cobra"
	"sort"
)

type options struct {
	registryID string

	IO             *iostreams.IOStreams
	Logger         logging.Logger
	Connection     factory.ConnectionFunc
	localizer      localize.Localizer
	context        context.Context
	ServiceContext servicecontext.IContext
}

func NewGetTypesCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Connection:     f.Connection,
		IO:             f.IOStreams,
		localizer:      f.Localizer,
		Logger:         f.Logger,
		context:        f.Context,
		ServiceContext: f.ServiceContext,
	}

	cmd := &cobra.Command{
		Use:     "types",
		Short:   f.Localizer.MustLocalize("artifact.cmd.types.description.short"),
		Long:    f.Localizer.MustLocalize("artifact.cmd.types.description.long"),
		Example: f.Localizer.MustLocalize("artifact.cmd.types.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {

			if opts.registryID != "" {

				return run(opts)

			} else {

				registryInstance, err := contextutil.GetCurrentRegistryInstance(f)
				if err != nil {
					return err
				}

				opts.registryID = registryInstance.GetId()
				return run(opts)
			}
		},
	}
	return cmd
}

func run(opts *options) error {
	conn, err := opts.Connection()
	if err != nil {
		return err
	}

	dataAPI, _, err := conn.API().ServiceRegistryInstance(opts.registryID)
	if err != nil {
		return err
	}

	types, err := GetArtifactTypes(dataAPI, opts.context)
	for _, v := range types {
		opts.Logger.Info(v)
	}
	return nil
}

func GetArtifactTypes(dataAPI *registryinstanceclient.APIClient, ctx context.Context) ([]string, error) {
	response, _, err := dataAPI.AdminApi.ListArtifactTypes(ctx).Execute()
	if err != nil {
		return nil, registrycmdutil.TransformInstanceError(err)
	}

	types := make([]string, 0)
	for _, v := range response {
		types = append(types, *v.Name)
	}
	sort.Strings(types)

	return types, nil
}
