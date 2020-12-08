package cluster

import (
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/cluster/connect"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/cluster/info"
	"github.com/spf13/cobra"
)

// NewServiceAccountCommand creates a new command sub-group to manage service accounts
func NewClusterCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cluster",
		Short: "Managed your services in OpenShift Cluster",
		Long:  "Managed your services in OpenShift Cluster",
		Args:  cobra.ExactArgs(1),
	}

	cmd.AddCommand(
		info.NewInfoCommand(),
		connect.NewConnectCommand(),
	)

	return cmd
}
