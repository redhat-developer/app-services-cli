package list

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/core/profile"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

type options struct {
	IO         *iostreams.IOStreams
	Config     config.IConfig
	Logger     logging.Logger
	Connection factory.ConnectionFunc
	localizer  localize.Localizer
	Context    context.Context
	Profiles   profile.IContext
}

func NewListCommand(f *factory.Factory) *cobra.Command {

	opts := &options{
		Config:     f.Config,
		Connection: f.Connection,
		IO:         f.IOStreams,
		Logger:     f.Logger,
		localizer:  f.Localizer,
		Profiles:   f.Profile,
	}

	cmd := &cobra.Command{
		Use:     "list",
		Short:   f.Localizer.MustLocalize("context.list.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("context.list.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("context.list.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runList(opts)
		},
	}

	return cmd
}

func runList(opts *options) error {

	context, err := opts.Profiles.Load()
	if err != nil {
		return err
	}

	profiles := context.Contexts

	profileNames := make([]string, 0, len(profiles))

	for name := range profiles {
		profileNames = append(profileNames, name)
	}

	opts.Logger.Info(profileNames)

	return nil
}
