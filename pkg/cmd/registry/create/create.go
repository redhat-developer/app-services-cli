package create

import (
	"context"
	"fmt"

	"github.com/redhat-developer/app-services-cli/pkg/accountmgmtutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/registrycmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/factory"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/connection"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/remote"

	srsmgmtv1 "github.com/redhat-developer/app-services-sdk-go/registrymgmt/apiv1/client"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

type options struct {
	name        string
	description string

	outputFormat string
	autoUse      bool

	interactive      bool
	bypassTermsCheck bool

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	localizer  localize.Localizer
	Context    context.Context
}

// NewCreateCommand creates a new command for creating registry.
func NewCreateCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		IO:         f.IOStreams,
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "create",
		Short:   f.Localizer.MustLocalize("registry.cmd.create.shortDescription"),
		Long:    f.Localizer.MustLocalize("registry.cmd.create.longDescription"),
		Example: f.Localizer.MustLocalize("registry.cmd.create.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.name != "" {
				if err := registrycmdutil.ValidateName(opts.name); err != nil {
					return err
				}
			}

			if !opts.IO.CanPrompt() && opts.name == "" {
				return opts.localizer.MustLocalizeError("registry.cmd.create.error.name.requiredWhenNonInteractive")
			} else if opts.name == "" {
				opts.interactive = true
			}

			validOutputFormats := flagutil.ValidOutputFormats
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flagutil.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			return runCreate(opts)
		},
	}

	flags := flagutil.NewFlagSet(cmd, opts.localizer)

	flags.StringVar(&opts.name, "name", "", opts.localizer.MustLocalize("registry.cmd.create.flag.name.description"))
	flags.StringVarP(&opts.outputFormat, "output", "o", "json", opts.localizer.MustLocalize("registry.cmd.flag.output.description"))
	flags.StringVar(&opts.description, "description", "", opts.localizer.MustLocalize("registry.cmd.create.flag.description.description"))
	flags.BoolVar(&opts.autoUse, "use", true, opts.localizer.MustLocalize("registry.cmd.create.flag.use.description"))
	flags.AddBypassTermsCheck(&opts.bypassTermsCheck)

	flagutil.EnableOutputFlagCompletion(cmd)

	return cmd
}

func runCreate(opts *options) error {
	cfg, err := opts.Config.Load()
	if err != nil {
		return err
	}

	var payload *srsmgmtv1.RegistryCreate
	if opts.interactive {
		opts.Logger.Debug()

		payload, err = promptPayload(opts)
		if err != nil {
			return err
		}
	} else {
		payload = &srsmgmtv1.RegistryCreate{
			Name: &opts.name,
		}

		if opts.description != "" {
			payload.SetDescription(opts.description)
		}
	}

	var conn connection.Connection
	if conn, err = opts.Connection(connection.DefaultConfigSkipMasAuth); err != nil {
		return err
	}

	if !opts.bypassTermsCheck {
		opts.Logger.Debug("Checking if terms and conditions have been accepted")
		// the user must have accepted the terms and conditions from the provider
		// before they can create a registry instance
		err, constants := remote.GetRemoteServiceConstants(opts.Context, opts.Logger)
		if err != nil {
			return err
		}
		var termsAccepted bool
		var termsURL string
		termsAccepted, termsURL, err = accountmgmtutil.CheckTermsAccepted(opts.Context, constants.ServiceRegistry.Ams, conn)
		if err != nil {
			return err
		}
		if !termsAccepted && termsURL != "" {
			opts.Logger.Info(opts.localizer.MustLocalize("service.info.termsCheck", localize.NewEntry("TermsURL", termsURL)))
			return nil
		}
	}

	opts.Logger.Info(opts.localizer.MustLocalize("registry.cmd.create.info.action", localize.NewEntry("Name", payload.GetName())))

	response, _, err := conn.API().
		ServiceRegistryMgmt().
		CreateRegistry(opts.Context).
		RegistryCreate(*payload).
		Execute()
	if err != nil {
		return err
	}

	opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("registry.cmd.create.info.successMessage"))

	if err = dump.Formatted(opts.IO.Out, opts.outputFormat, response); err != nil {
		return err
	}

	registryConfig := &config.ServiceRegistryConfig{
		InstanceID: response.GetId(),
		Name:       response.GetName(),
	}

	if opts.autoUse {
		opts.Logger.Debug("Auto-use is set, updating the current instance")
		cfg.Services.ServiceRegistry = registryConfig
		if err := opts.Config.Save(cfg); err != nil {
			return fmt.Errorf("%v: %w", opts.localizer.MustLocalize("registry.cmd.create.error.couldNotUse"), err)
		}
	} else {
		opts.Logger.Debug("Auto-use is not set, skipping updating the current instance")
	}

	return nil
}

// Show a prompt to allow the user to interactively insert the data
func promptPayload(opts *options) (payload *srsmgmtv1.RegistryCreate, err error) {
	if err != nil {
		return nil, err
	}

	// set type to store the answers from the prompt with defaults
	answers := struct {
		Name        string
		Description string
	}{}

	promptName := &survey.Input{
		Message: opts.localizer.MustLocalize("registry.cmd.create.input.name.message"),
		Help:    opts.localizer.MustLocalize("registry.cmd.create.input.name.help"),
	}

	err = survey.AskOne(promptName, &answers.Name, survey.WithValidator(registrycmdutil.ValidateName))
	if err != nil {
		return nil, err
	}

	promptDescription := &survey.Multiline{
		Message: opts.localizer.MustLocalize("registry.cmd.create.input.description.message"),
		Help:    opts.localizer.MustLocalize("registry.cmd.create.input.description.help"),
	}

	err = survey.AskOne(promptDescription, &answers.Description)
	if err != nil {
		return nil, err
	}

	payload = &srsmgmtv1.RegistryCreate{
		Name: &answers.Name,
	}

	if answers.Description != "" {
		payload.SetDescription(answers.Description)
	}

	return payload, nil
}
