package cluster

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/cluster/connect"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/cluster/info"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/spf13/cobra"
)

// NewServiceAccountCommand creates a new command sub-group to manage service accounts
func NewClusterCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cluster",
		Short: "View and perform operations on your Kubernetes or OpenShift Cluster",
		Long: heredoc.Doc(`
			View information about your Kubernetes or OpenShift cluster 
			and perform operations related to your application services.
		`),
		Example: heredoc.Doc(`
			# view information about your cluster
			$ rhoas cluster info

			# connect a service to your cluster
			$ rhoas cluster connect
		`),
		Args: cobra.ExactArgs(1),
	}

	cmd.AddCommand(
		info.NewInfoCommand(f),
		connect.NewConnectCommand(f),
	)

	return cmd
}
