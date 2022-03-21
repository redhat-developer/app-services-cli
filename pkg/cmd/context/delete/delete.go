package delete

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/context/contextcmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/redhat-developer/app-services-cli/pkg/shared/profileutil"
	"github.com/spf13/cobra"
)

type options struct {
	IO             *iostreams.IOStreams
	Logger         logging.Logger
	Connection     factory.ConnectionFunc
	localizer      localize.Localizer
	Context        context.Context
	ServiceContext servicecontext.IContext

	name string
}

// NewDeleteCommand command for deleting service contexts
func NewDeleteCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Connection:     f.Connection,
		Logger:         f.Logger,
		IO:             f.IOStreams,
		localizer:      f.Localizer,
		Context:        f.Context,
		ServiceContext: f.ServiceContext,
	}

	cmd := &cobra.Command{
		Use:     "delete",
		Short:   f.Localizer.MustLocalize("context.delete.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("context.delete.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("context.delete.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDelete(opts)
		},
	}

	flags := contextcmdutil.NewFlagSet(cmd, f)

	flags.AddContextName(&opts.name)

	return cmd
}

func runDelete(opts *options) error {

	svcContext, err := opts.ServiceContext.Load()
	if err != nil {
		return err
	}

	profileHandler := &profileutil.ContextHandler{
		Context:   svcContext,
		Localizer: opts.localizer,
	}

	if opts.name == "" {

		currCtx, newErr := profileHandler.GetCurrentContext()
		if newErr != nil {
			return newErr
		}

		opts.name = currCtx

		svcContext.CurrentContext = ""
	}

	if _, err = profileHandler.GetContext(opts.name); err != nil {
		return err
	}

	delete(svcContext.Contexts, opts.name)

	err = opts.ServiceContext.Save(svcContext)
	if err != nil {
		return err
	}

	opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("context.delete.log.successMessage"))

	return nil

}
