package status

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/pkg/shared/profileutil"
)

type options struct {
	IO             *iostreams.IOStreams
	Logger         logging.Logger
	Connection     factory.ConnectionFunc
	localizer      localize.Localizer
	Context        context.Context
	ServiceContext servicecontext.IContext

	name         string
	outputFormat string
}

// NewStatusCommand creates a new command to display status of a context
func NewStatusCommand(f *factory.Factory) *cobra.Command {

	opts := &options{
		Connection:     f.Connection,
		IO:             f.IOStreams,
		Logger:         f.Logger,
		localizer:      f.Localizer,
		ServiceContext: f.ServiceContext,
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
	flags.AddOutput(&opts.outputFormat)

	return cmd
}

func runStatus(opts *options) error {

	var currentCtx *servicecontext.ServiceConfig
	var err error

	context, err := opts.ServiceContext.Load()
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

	stdout := opts.IO.Out
	err = dump.Formatted(stdout, opts.outputFormat, currentCtx)
	if err != nil {
		return err
	}

	return nil
}
