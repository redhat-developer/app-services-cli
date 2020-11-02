// Package cluster contains commands for interacting with cluster logic of the service directly instead of through the
// REST API exposed via the serve command.
package auth

import (
	"github.com/spf13/cobra"
)

func NewAuthGroupCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Authentication and Authorization",
		Long:  "Authentication and Authorization",
	}

	// add sub-commands
	cmd.AddCommand(NewAuthorizationCommand())

	return cmd
}
