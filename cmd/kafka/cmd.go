// Package cluster contains commands for interacting with cluster logic of the service directly instead of through the
// REST API exposed via the serve command.
package kafka

import (
	"github.com/spf13/cobra"
)

func NewKafkaCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "streams",
		Short: "Manage your OpenShift Streams instances",
		Long:  "Manage your OpenShift Streams instances",
	}

	// add sub-commands
	cmd.AddCommand(NewCreateCommand(), NewGetCommand(), NewDeleteCommand(),
		NewListCommand(),
		NewUseCommand(),
		NewTopicsCommand(),
		NewCredentialsCommand())

	return cmd
}
