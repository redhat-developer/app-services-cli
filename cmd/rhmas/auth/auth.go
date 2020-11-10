// Package cluster contains commands for interacting with cluster logic of the service directly instead of through the
// REST API exposed via the serve command.
package auth

import (
	"github.com/bf2fc6cc711aee1a0c2a/cli/cmd/rhmas/auth/authorization"
	"github.com/spf13/cobra"
)

func NewAuthGroupCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Authentication and Authorization",
		Long:  "Authentication and Authorization",
	}

	// add sub-commands
	cmd.AddCommand(authorization.NewAuthorizationCommand())

	return cmd
}
