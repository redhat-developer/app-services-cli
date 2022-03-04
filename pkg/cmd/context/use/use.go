package use

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/redhat-developer/app-services-cli/pkg/shared/profileutil"
	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/pkg/core/profile"
)

type options struct {
	IO         *iostreams.IOStreams
	Config     config.IConfig
	Logger     logging.Logger
	Connection factory.ConnectionFunc
	localizer  localize.Localizer
	Context    context.Context
	Profiles   profile.IContext

	name string
}

func NewUseCommand(f *factory.Factory) *cobra.Command {

	opts := &options{
		Config:     f.Config,
		Connection: f.Connection,
		IO:         f.IOStreams,
		Logger:     f.Logger,
		localizer:  f.Localizer,
		Profiles:   f.Profile,
	}

	cmd := &cobra.Command{
		Use:     "use",
		Short:   f.Localizer.MustLocalize("context.use.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("context.use.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("context.use.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runUse(opts)
		},
	}

	flags := flagutil.NewFlagSet(cmd, opts.localizer)

	flags.StringVar(&opts.name, "name", "", opts.localizer.MustLocalize("context.common.flag.name"))

	return cmd
}

func runUse(opts *options) error {

	context, err := opts.Profiles.Load()
	if err != nil {
		return err
	}

	profileHandler := &profileutil.ContextHandler{
		Context:   context,
		Localizer: opts.localizer,
	}

	_, err = profileHandler.GetContext(context.CurrentContext)
	if err != nil {
		return err
	}

	context.CurrentContext = opts.name

	err = opts.Profiles.Save(context)
	if err != nil {
		return err
	}

	opts.Logger.Info(opts.localizer.MustLocalize("context.use.successMessage", localize.NewEntry("Name", opts.name)))

	return nil
}
