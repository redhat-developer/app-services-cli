package connectors

import (
	"github.com/redhat-developer/app-services-cli/internal/doc"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/connectors/addon"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/connectors/cluster"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/connectors/create"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/connectors/delete"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/connectors/describe"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/connectors/list"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/connectors/namespaces"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/connectors/update"

	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

func NewConnectorsCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:         "connector",
		Hidden:      true,
		Annotations: map[string]string{doc.AnnotationName: "Connectors commands"},
		Short:       f.Localizer.MustLocalize("connectors.cmd.shortDescription"),
		Long:        f.Localizer.MustLocalize("connectors.cmd.longDescription"),
		Example:     f.Localizer.MustLocalize("connectors.cmd.example"),
		Args:        cobra.MinimumNArgs(1),
	}

	// add sub-commands
	cmd.AddCommand(
		cluster.NewConnectorClusterCommand(f),
		create.NewCreateCommand(f),
		delete.NewDeleteCommand(f),
		list.NewListCommand(f),
		describe.NewDescribeCommand(f),
		namespaces.NewNameSpaceCommand(f),
		addon.NewParametersCommand(f),
		update.NewUpdateCommand(f),
	)

	return cmd
}
