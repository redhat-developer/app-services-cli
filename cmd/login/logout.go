// Package cluster contains commands for interacting with cluster logic of the service directly instead of through the
// REST API exposed via the serve command.
package login

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewLogoutCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logout",
		Short: "Logout from connected Managed Application Services cluster",
		Long:  "Logout from connected Managed Application Services cluster",
		Run:   runLogout,
	}
	return cmd
}

func runLogout(cmd *cobra.Command, _ []string) {
	fmt.Printf("Successfully logged out")
}
