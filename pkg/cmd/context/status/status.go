package status

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/pkg/core/profile"
	"github.com/redhat-developer/app-services-cli/pkg/shared/profileutil"
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

func NewStatusCommand(f *factory.Factory) *cobra.Command {

	opts := &options{
		Config:     f.Config,
		Connection: f.Connection,
		IO:         f.IOStreams,
		Logger:     f.Logger,
		localizer:  f.Localizer,
		Profiles:   f.Profile,
	}

	cmd := &cobra.Command{
		Use:     "status",
		Short:   f.Localizer.MustLocalize("context.status.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("context.status.cmd.shortDescription"),
		Example: f.Localizer.MustLocalize("context.status.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runStatus(opts)
		},
	}

	flags := flagutil.NewFlagSet(cmd, opts.localizer)

	flags.StringVar(&opts.name, "name", "", opts.localizer.MustLocalize("context.common.flag.name"))

	return cmd
}

func runStatus(opts *options) error {

	stdout := opts.IO.Out
	var currentCtx *profile.ServiceConfig
	var err error

	context, err := opts.Profiles.Load()
	if err != nil {
		return err
	}

	profileHandler := &profileutil.ContextHandler{
		Context:   context,
		Localizer: opts.localizer,
	}

	if opts.name != "" {
		currentCtx, err = profileHandler.GetContext(opts.name)
		if err != nil {
			return err
		}
	} else {
		currentCtx, err = profileHandler.GetContext(context.CurrentContext)
		if err != nil {
			return err
		}
	}

	err = dump.Formatted(stdout, "yml", currentCtx)
	if err != nil {
		return err
	}

	return nil
}
