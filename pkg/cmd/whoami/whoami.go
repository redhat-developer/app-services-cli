package whoami

import (
	"fmt"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/auth/token"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"
	"github.com/spf13/cobra"
)

type Options struct {
	Config config.IConfig
	Logger func() (logging.Logger, error)
}

func NewWhoAmICmd(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config: f.Config,
		Logger: f.Logger,
	}

	localizer.LoadMessageFiles("cmd/whoami")

	cmd := &cobra.Command{
		Use:     localizer.MustLocalizeFromID("whoami.cmd.use"),
		Short:   localizer.MustLocalizeFromID("whoami.cmd.shortDescription"),
		Long:    localizer.MustLocalizeFromID("whoami.cmd.longDescription"),
		Example: localizer.MustLocalizeFromID("whoami.cmd.example"),
		RunE: func(cmd *cobra.Command, _ []string) error {
			return getCurrentUser(opts)
		},
	}

	return cmd
}

func getCurrentUser(opts *Options) (err error) {
	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	cfg, err := opts.Config.Load()
	if err != nil {
		logger.Error(err)
		return err
	}

	accessTkn, _ := token.Parse(cfg.AccessToken)

	if accessTkn == nil {
		logger.Info(localizer.MustLocalizeFromID("whoami.error.notLoggedInError"))
		return nil
	}

	tknClaims, _ := token.MapClaims(accessTkn)

	userName, _ := tknClaims["preferred_username"]

	logger.Info(fmt.Sprintf("%v", userName))

	return nil
}
