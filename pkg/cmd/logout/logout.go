package logout

import (
	"context"
	"fmt"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/factory"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/connection"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/spf13/cobra"
)

type options struct {
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	localizer  localize.Localizer
	Context    context.Context
}

// NewLogoutCommand gets the command that's logs the current logged in user
func NewLogoutCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "logout",
		Short:   opts.localizer.MustLocalize("logout.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("logout.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("logout.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runLogout(opts)
		},
	}
	return cmd
}

func runLogout(opts *options) error {
	conn, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	err = conn.Logout(opts.Context)

	if err != nil {
		return fmt.Errorf("%v: %w", opts.localizer.MustLocalize("logout.error.unableToLogout"), err)
	}

	opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("logout.log.info.logoutSuccess"))

	return nil
}
