package cluster

import (
	"github.com/redhat-developer/app-services-cli/pkg/cmd/cluster/bind"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/cluster/connect"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/cluster/status"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/spf13/cobra"
)

// NewServiceAccountCommand creates a new command sub-group to manage service accounts
func NewClusterCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     f.Localizer.LoadMessage("cluster.cmd.use"),
		Short:   f.Localizer.LoadMessage("cluster.cmd.shortDescription"),
		Example: f.Localizer.LoadMessage("cluster.cmd.example"),
		Args:    cobra.ExactArgs(1),
	}

	cmd.AddCommand(
		status.NewStatusCommand(f),
		connect.NewConnectCommand(f),
		bind.NewBindCommand(f),
	)

	return cmd
}
