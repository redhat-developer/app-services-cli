// Package cluster contains commands for interacting with cluster logic of the service directly instead of through the
// REST API exposed via the serve command.
package logout

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/config"
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
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v", err)
		os.Exit(1)
	}

	connection, err := cfg.Connection()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	err = connection.Logout(context.TODO())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to log out: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully logged out.")

	cfg.SetAccessToken("")
	cfg.SetRefreshToken("")

	err = config.Save(cfg)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not save config file: %v\n", err)
		os.Exit(1)
	}
}
