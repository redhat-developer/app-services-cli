package create

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

// NewCreateCommand creates a new command to create contexts
func NewCreateCommand(f *factory.Factory) *cobra.Command {

	opts := &options{
		Connection:     f.Connection,
		IO:             f.IOStreams,
		Logger:         f.Logger,
		localizer:      f.Localizer,
		ServiceContext: f.ServiceContext,
	}

	cmd := &cobra.Command{
		Use:     "create",
		Short:   f.Localizer.MustLocalize("context.create.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("context.create.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("context.create.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {

			return runCreate(opts)
		},
	}

	flags := contextcmdutil.NewFlagSet(cmd, f)

	_ = flags.AddContextName(&opts.name)

	return cmd

}

func runCreate(opts *options) error {

	svcContext, err := opts.ServiceContext.Load()
	if err != nil {
		return err
	}

	profileHandler := &profileutil.ContextHandler{
		Context:   svcContext,
		Localizer: opts.localizer,
	}

	profileValidator := &contextcmdutil.Validator{
		Localizer:      opts.localizer,
		ProfileHandler: profileHandler,
	}

	profiles := svcContext.Contexts

	if profiles == nil {
		profiles = make(map[string]servicecontext.ServiceConfig)
	}

	err = profileValidator.ValidateName(opts.name)
	if err != nil {
		return err
	}

	err = profileValidator.ValidateNameIsAvailable(opts.name)
	if err != nil {
		return err
	}

	context, _ := profileHandler.GetContext(opts.name)
	if context != nil {
		return opts.localizer.MustLocalizeError("context.create.log.alreadyExists", localize.NewEntry("Name", opts.name))
	}

	profiles[opts.name] = servicecontext.ServiceConfig{}

	svcContext.Contexts = profiles

	err = opts.ServiceContext.Save(svcContext)
	if err != nil {
		return err
	}

	opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("context.create.log.successMessage", localize.NewEntry("Name", opts.name)))

	return nil
}
