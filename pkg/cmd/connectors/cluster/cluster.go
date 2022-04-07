package cluster

import (
	"github.com/redhat-developer/app-services-cli/pkg/cmd/connectors/cluster/create"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/connectors/cluster/delete"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/connectors/cluster/list"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

func NewConnectorClusterCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cluster",
		Short:   f.Localizer.MustLocalize("connectors.cluster.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("connectors.cluster.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("connectors.cluster.cmd.example"),
		Args:    cobra.MinimumNArgs(1),
	}

	// add sub-commands
	cmd.AddCommand(
		create.NewCreateCommand(f),
		delete.NewDeleteCommand(f),
		list.NewListCommand(f),
	)

	return cmd
}
