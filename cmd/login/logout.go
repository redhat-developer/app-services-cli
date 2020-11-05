// Package cluster contains commands for interacting with cluster logic of the service directly instead of through the
// REST API exposed via the serve command.
package login

import (
	"fmt"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/config"
	"github.com/spf13/cobra"
)

// NewLogoutCommand gets the command that's logs the current logged in user
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
	err := config.Remove()
	if err == nil {
		fmt.Println("Successfully logged out")
		return
	}

	fmt.Errorf("Unable to logout %v", err)
}
