package whoami

import (
	"fmt"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/auth/token"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/logging"

	"github.com/spf13/cobra"
)

type args struct {
	Config     config.IConfig
	Connection factory.ConnectionFunc
	IO         *iostreams.IOStreams
	Logger     logging.Logger
	localizer  localize.Localizer
}

func NewWhoAmICmd(f *factory.Factory) *cobra.Command {
	opts := &args{
		Config:     f.Config,
		Connection: f.Connection,
		IO:         f.IOStreams,
		Logger:     f.Logger,
		localizer:  f.Localizer,
	}

	cmd := &cobra.Command{
		Use:     "whoami",
		Short:   f.Localizer.MustLocalize("whoami.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("whoami.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("whoami.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runCmd(opts)
		},
	}

	return cmd
}

func runCmd(opts *args) (err error) {
	cfg, err := opts.Config.Load()
	if err != nil {
		return err
	}

	_, err = opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	userName, ok := token.GetUsername(cfg.AccessToken)
	if !ok {
		userName = "unknown"
	}

	if ok {
		fmt.Fprintln(opts.IO.Out, userName)
	} else {
		opts.Logger.Info(opts.localizer.MustLocalize("whoami.log.info.tokenHasNoUsername"))
	}

	return nil
}
