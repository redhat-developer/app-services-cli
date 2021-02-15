package cluster

import (
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/cluster/connect"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/cluster/status"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/spf13/cobra"
)

// NewServiceAccountCommand creates a new command sub-group to manage service accounts
func NewClusterCommand(f *factory.Factory) *cobra.Command {
	localizer.LoadMessageFiles("cmd/cluster")

	cmd := &cobra.Command{
		Use:   localizer.MustLocalizeFromID("cluster.status.cmd.use"),
		Short: localizer.MustLocalizeFromID("cluster.status.cmd.shortDescription"),
		Args:  cobra.ExactArgs(1),
	}

	cmd.AddCommand(
		status.NewStatusCommand(f),
		connect.NewConnectCommand(f),
	)

	return cmd
}
