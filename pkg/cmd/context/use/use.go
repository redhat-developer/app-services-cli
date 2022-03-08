package use

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
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

// NewUseCommand creates a new command to set the current context
func NewUseCommand(f *factory.Factory) *cobra.Command {

	opts := &options{
		Connection:     f.Connection,
		IO:             f.IOStreams,
		Logger:         f.Logger,
		localizer:      f.Localizer,
		ServiceContext: f.ServiceContext,
	}

	cmd := &cobra.Command{
		Use:     "use",
		Short:   f.Localizer.MustLocalize("context.use.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("context.use.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("context.use.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {

			if !opts.IO.CanPrompt() && opts.name == "" {
				return flagutil.RequiredWhenNonInteractiveError("name")
			}

			return runUse(opts)
		},
	}

	flags := flagutil.NewFlagSet(cmd, opts.localizer)

	flags.StringVar(&opts.name, "name", "", opts.localizer.MustLocalize("context.common.flag.name"))

	return cmd
}

func runUse(opts *options) error {

	context, err := opts.ServiceContext.Load()
	if err != nil {
		return err
	}

	if opts.name == "" {
		opts.name, err = runInteractivePrompt(opts, context)
		if err != nil {
			return err
		}
	}

	profileHandler := &profileutil.ContextHandler{
		Context:   context,
		Localizer: opts.localizer,
	}

	_, err = profileHandler.GetContext(opts.name)
	if err != nil {
		return err
	}

	context.CurrentContext = opts.name

	err = opts.ServiceContext.Save(context)
	if err != nil {
		return err
	}

	opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("context.use.successMessage", localize.NewEntry("Name", opts.name)))

	return nil
}

func runInteractivePrompt(opts *options, context *servicecontext.Context) (string, error) {

	profiles := context.Contexts

	if profiles == nil {
		profiles = map[string]servicecontext.ServiceConfig{}
	}

	profileNames := make([]string, 0, len(profiles))

	for name := range profiles {
		profileNames = append(profileNames, name)
	}

	if len(profileNames) == 0 {
		opts.Logger.Info(opts.localizer.MustLocalize("context.list.log.noContexts"))
		return "", nil
	}

	prompt := &survey.Select{
		Message:  opts.localizer.MustLocalize("context.common.flag.name"),
		Options:  profileNames,
		PageSize: 10,
	}

	var selectedServiceContext string
	err := survey.AskOne(prompt, &selectedServiceContext)
	if err != nil {
		return "", err
	}

	return selectedServiceContext, nil

}
