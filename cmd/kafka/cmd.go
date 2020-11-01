// Package cluster contains commands for interacting with cluster logic of the service directly instead of through the
// REST API exposed via the serve command.
package kafka

import (
	"github.com/spf13/cobra"
)

func NewKafkaCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "streaming",
		Short: "Manage your OpenShift Streaming instances",
		Long:  "Manage your OpenShift Streaming instances",
	}

	// add sub-commands
	cmd.AddCommand(NewCreateCommand(), NewGetCommand(), NewDeleteCommand(), NewListCommand(), NewUseCommand())

	return cmd
}
