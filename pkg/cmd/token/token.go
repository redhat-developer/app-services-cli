package token

import (
	"fmt"

	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

type options struct {
	Config     config.IConfig
	Connection factory.ConnectionFunc
	IO         *iostreams.IOStreams
	Logger     logging.Logger
	localizer  localize.Localizer
}

func NewAuthTokenCmd(f *factory.Factory) *cobra.Command {
	opts := &options{
		Config:     f.Config,
		Connection: f.Connection,
		IO:         f.IOStreams,
		Logger:     f.Logger,
		localizer:  f.Localizer,
	}

	cmd := &cobra.Command{
		Use:     "authtoken",
		Short:   f.Localizer.MustLocalize("token.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("token.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("token.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runCmd(opts)
		},
	}

	return cmd
}

func runCmd(opts *options) (err error) {
	cfg, err := opts.Config.Load()
	if err != nil {
		return err
	}

	_, err = opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	if cfg.AccessToken != "" {
		fmt.Fprintln(opts.IO.Out, cfg.AccessToken)
	} else {
		opts.Logger.Info(opts.localizer.MustLocalize("token.log.info.tokenUnavailable"))
	}

	return nil
}
