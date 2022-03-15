package create

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/context/contextcmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/redhat-developer/app-services-cli/pkg/shared/kafkautil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/profileutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/serviceregistryutil"
	"github.com/spf13/cobra"

	"github.com/AlecAivazis/survey/v2"
)

type options struct {
	IO             *iostreams.IOStreams
	Logger         logging.Logger
	Connection     factory.ConnectionFunc
	localizer      localize.Localizer
	Context        context.Context
	ServiceContext servicecontext.IContext

	name        string
	kafkaID     string
	registryID  string
	interactive bool
	autoUse     bool
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

			if !opts.IO.CanPrompt() && opts.name == "" {
				return flagutil.RequiredWhenNonInteractiveError("name")
			} else if opts.name == "" {
				opts.interactive = true
			}

			return runCreate(opts)
		},
	}

	flags := contextcmdutil.NewFlagSet(cmd, f)

	flags.AddContextName(&opts.name)
	flags.BoolVar(&opts.autoUse, "use", true, opts.localizer.MustLocalize("context.create.flag.use"))
	flags.StringVar(&opts.kafkaID, "kafka-id", "", opts.localizer.MustLocalize("context.create.flag.kafkaID"))
	flags.StringVar(&opts.registryID, "registry-id", "", opts.localizer.MustLocalize("context.create.flag.registryID"))

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

	if opts.interactive {
		if err = runInteractive(opts); err != nil {
			return err
		}
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

	services := servicecontext.ServiceConfig{
		KafkaID:           opts.kafkaID,
		ServiceRegistryID: opts.registryID,
	}

	profiles[opts.name] = services

	svcContext.Contexts = profiles

	if opts.autoUse {
		opts.Logger.Debug("Auto-use is set, updating the current service context")
		svcContext.CurrentContext = opts.name
	} else {
		opts.Logger.Debug("Auto-use is not set, skipping updating the current service context")
	}

	err = opts.ServiceContext.Save(svcContext)
	if err != nil {
		return err
	}

	opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("context.create.log.successMessage"))

	return nil
}

func runInteractive(opts *options) error {

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

	conn, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	promptName := &survey.Input{
		Message: opts.localizer.MustLocalize("context.create.input.name.message"),
	}

	err = survey.AskOne(
		promptName,
		&opts.name,
		survey.WithValidator(survey.Required),
		survey.WithValidator(profileValidator.ValidateName),
		survey.WithValidator(profileValidator.ValidateNameIsAvailable),
	)

	if err != nil {
		return err
	}

	selectedKafka, err := kafkautil.InteractiveSelect(opts.Context, conn, opts.Logger, opts.localizer)
	if err != nil {
		return err
	}

	opts.kafkaID = selectedKafka.GetId()

	selectedRegistry, err := serviceregistryutil.InteractiveSelect(opts.Context, conn, opts.Logger)
	if err != nil {
		return err
	}

	opts.registryID = selectedRegistry.GetId()

	return nil
}
