package status

import (
	"errors"
	"strconv"
	"strings"

	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/profile"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/logging"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/spf13/cobra"
)

type Options struct {
	IO        *iostreams.IOStreams
	Config    config.IConfig
	Logger    func() (logging.Logger, error)
	localizer localize.Localizer

	devPreview bool
}

const DevPreviewConfigArgument = "devPreview"

var validArguments = []string{DevPreviewConfigArgument}

func NewConfigCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		IO:        f.IOStreams,
		Config:    f.Config,
		Logger:    f.Logger,
		localizer: f.Localizer,
	}

	cmd := &cobra.Command{
		Use:       "config",
		Short:     opts.localizer.MustLocalize("config.cmd.shortDescription"),
		Long:      opts.localizer.MustLocalize("config.cmd.longDescription"),
		Example:   opts.localizer.MustLocalize("config.cmd.example"),
		ValidArgs: validArguments,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return errors.New("Command accepts two arguments")
			}

			invalidConfigArg := true
			for _, configArg := range validArguments {
				if configArg == args[0] {
					invalidConfigArg = false
				}
			}

			if invalidConfigArg {
				return errors.New("First argument should contain: " + strings.Join(validArguments, ","))
			}

			if _, err := strconv.ParseBool(args[1]); err != nil {
				return errors.New("Second argument should be boolean")
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if args[0] == DevPreviewConfigArgument {
				devPreview, _ := strconv.ParseBool(args[1])
				_, err := profile.EnableDevPreview(f, devPreview)
				return err
			}
			return nil
		},
	}

	return cmd
}
