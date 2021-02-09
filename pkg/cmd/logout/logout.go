// Package cluster contains commands for interacting with cluster logic of the service directly instead of through the
// REST API exposed via the serve command.
package logout

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"
)

type Options struct {
	Config     config.IConfig
	Connection func() (connection.Connection, error)
	Logger     func() (logging.Logger, error)
}

// NewLogoutCommand gets the command that's logs the current logged in user
func NewLogoutCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
	}

	localizer.LoadMessageFiles("cmd/logout")

	cmd := &cobra.Command{
		Use:   localizer.MustLocalizeFromID("logout.cmd.use"),
		Short: localizer.MustLocalizeFromID("logout.cmd.shortDescription"),
		Long:  localizer.MustLocalizeFromID("logout.cmd.longDescription"),
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runLogout(opts)
		},
	}
	return cmd
}

func runLogout(opts *Options) error {
	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	cfg, err := opts.Config.Load()
	if err != nil {
		return err
	}

	connection, err := opts.Connection()
	if err != nil {
		return err
	}

	err = connection.Logout(context.TODO())

	if err != nil {
		return fmt.Errorf("%v: %w", localizer.MustLocalizeFromID("logout.error.unableToLogout"), err)
	}

	logger.Info(localizer.MustLocalizeFromID("logout.log.info.logoutSuccess"))

	cfg.AccessToken = ""
	cfg.RefreshToken = ""

	err = opts.Config.Save(cfg)
	if err != nil {
		return err
	}

	return nil
}
