package whoami

import (
	"fmt"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/auth/token"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/iostreams"

	"github.com/spf13/cobra"
)

type Options struct {
	Config     config.IConfig
	Connection func() (connection.Connection, error)
	IO         *iostreams.IOStreams
}

func NewWhoAmICmd(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		IO:         f.IOStreams,
	}

	cmd := &cobra.Command{
		Use:     localizer.MustLocalizeFromID("whoami.cmd.use"),
		Short:   localizer.MustLocalizeFromID("whoami.cmd.shortDescription"),
		Long:    localizer.MustLocalizeFromID("whoami.cmd.longDescription"),
		Example: localizer.MustLocalizeFromID("whoami.cmd.example"),
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

	_, err = opts.Connection()
	if err != nil {
		return err
	}

	accessTkn, _ := token.Parse(cfg.AccessToken)

	tknClaims, _ := token.MapClaims(accessTkn)

	userName, _ := tknClaims["preferred_username"]

	fmt.Fprintln(opts.IO.Out, userName)

	return nil
}
