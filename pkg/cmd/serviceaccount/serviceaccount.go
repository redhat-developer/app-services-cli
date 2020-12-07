package serviceaccount

import (
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/serviceaccount/create"
	"github.com/spf13/cobra"
)

// NewServiceAccountCommand creates a new command sub-group to manage service accounts
func NewServiceAccountCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "service-account",
		Short: "Manage your service accounts",
		Long:  "Manage your service accounts which can be used to connect your application to a Kafka instances",
		Args:  cobra.ExactArgs(1),
	}

	cmd.AddCommand(
		create.NewCreateCommand(),
	)

	return cmd
}
