package connectors

import (
	"github.com/redhat-developer/app-services-cli/internal/doc"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/connectors/cluster"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/connectors/create"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/connectors/delete"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/connectors/list"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

func NewConnectorsCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:         "connectors",
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
	)

	return cmd
}
