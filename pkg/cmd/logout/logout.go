// Package cluster contains commands for interacting with cluster logic of the service directly instead of through the
// REST API exposed via the serve command.
package logout

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
)

type LogoutOptions struct {
	Config     config.IConfig
	Connection func() (connection.IConnection, error)
}

// NewLogoutCommand gets the command that's logs the current logged in user
func NewLogoutCommand(f *factory.Factory) *cobra.Command {
	opts := &LogoutOptions{
		Config:     f.Config,
		Connection: f.Connection,
	}

	cmd := &cobra.Command{
		Use:   "logout",
		Short: "Logout from connected Managed Application Services cluster",
		Long:  "Logout from connected Managed Application Services cluster",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runLogout(opts)
		},
	}
	return cmd
}

func runLogout(opts *LogoutOptions) error {
	cfg, err := opts.Config.Load()
	if err != nil {
		return fmt.Errorf("Error loading config: %w", err)
	}

	connection, err := opts.Connection()
	if err != nil {
		return fmt.Errorf("Could not create connection: %w", err)
	}

	err = connection.Logout(context.TODO())

	if err != nil {
		return fmt.Errorf("Unable to log out: %w", err)
	}

	fmt.Fprintln(os.Stderr, "Successfully logged out")

	cfg.AccessToken = ""
	cfg.RefreshToken = ""

	err = opts.Config.Save(cfg)
	if err != nil {
		return fmt.Errorf("Could not save config file: %w", err)
	}

	return nil
}
