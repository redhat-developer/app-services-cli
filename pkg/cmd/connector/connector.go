package connector

import (
	"github.com/redhat-developer/app-services-cli/internal/doc"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/connector/addon"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/connector/cluster"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/connector/create"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/connector/delete"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/connector/describe"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/connector/list"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/connector/namespaces"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/connector/update"

	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

func NewConnectorsCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:         "connector",
		Hidden:      true,
		Annotations: map[string]string{doc.AnnotationName: "Connector commands"},
		Short:       f.Localizer.MustLocalize("connector.cmd.shortDescription"),
		Long:        f.Localizer.MustLocalize("connector.cmd.longDescription"),
		Example:     f.Localizer.MustLocalize("connector.cmd.example"),
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
