package cluster

import (
	"github.com/redhat-developer/app-services-cli/pkg/cmd/connector/cluster/addon"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/connector/cluster/create"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/connector/cluster/delete"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/connector/cluster/list"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/connector/cluster/update"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

// NewConnectorClusterCommand creates a new command to manage connector clusters
func NewConnectorClusterCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cluster",
		Short:   f.Localizer.MustLocalize("connector.cluster.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("connector.cluster.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("connector.cluster.cmd.example"),
		Args:    cobra.MinimumNArgs(1),
	}

	// add sub-commands
	cmd.AddCommand(
		create.NewCreateCommand(f),
		delete.NewDeleteCommand(f),
		list.NewListCommand(f),
		addon.NewParametersCommand(f),
		update.NewUpdateCommand(f),
	)

	return cmd
}
