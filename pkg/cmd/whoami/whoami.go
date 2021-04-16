package whoami

import (
	"fmt"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/internal/localizer"
	"github.com/redhat-developer/app-services-cli/pkg/auth/token"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/logging"

	"github.com/spf13/cobra"
)

type Options struct {
	Config     config.IConfig
	Connection factory.ConnectionFunc
	IO         *iostreams.IOStreams
	Logger     func() (logging.Logger, error)
}

func NewWhoAmICmd(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		IO:         f.IOStreams,
		Logger:     f.Logger,
	}

	cmd := &cobra.Command{
		Use:     localizer.MustLocalizeFromID("whoami.cmd.use"),
		Short:   localizer.MustLocalizeFromID("whoami.cmd.shortDescription"),
		Long:    localizer.MustLocalizeFromID("whoami.cmd.longDescription"),
		Example: localizer.MustLocalizeFromID("whoami.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runCmd(opts)
		},
	}

	return cmd
}

func runCmd(opts *Options) (err error) {
	cfg, err := opts.Config.Load()
	if err != nil {
		return err
	}

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	_, err = opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	accessTkn, _ := token.Parse(cfg.AccessToken)

	tknClaims, _ := token.MapClaims(accessTkn)

	userName, ok := tknClaims["preferred_username"]

	if ok {
		fmt.Fprintln(opts.IO.Out, userName)
	} else {
		logger.Info(localizer.MustLocalizeFromID("whoami.log.info.tokenHasNoUsername"))
	}

	return nil
}
