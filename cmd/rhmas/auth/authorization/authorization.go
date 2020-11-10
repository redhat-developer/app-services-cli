// Package cluster contains commands for interacting with cluster logic of the service directly instead of through the
// REST API exposed via the serve command.
package authorization

import (
	"github.com/spf13/cobra"
)

func NewAuthorizationCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "authorization",
		Short: "Manage authorization rules",
		Long:  "Manage authorization rules",
		Run:   runAuthz,
	}

	return cmd
}

func runAuthz(cmd *cobra.Command, _ []string) {

}
