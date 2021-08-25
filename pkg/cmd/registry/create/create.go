package create

import (
	"context"
	"errors"
	"fmt"

	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/serviceregistry"

	"github.com/redhat-developer/app-services-cli/pkg/ams"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"
	"github.com/redhat-developer/app-services-cli/pkg/connection"

	srsmgmtv1 "github.com/redhat-developer/app-services-sdk-go/registrymgmt/apiv1/client"

	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"

	"github.com/redhat-developer/app-services-cli/pkg/logging"

	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
)

type Options struct {
	name string

	outputFormat string
	autoUse      bool

	interactive bool

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	localizer  localize.Localizer
}

// NewCreateCommand creates a new command for creating registry.
func NewCreateCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		IO:         f.IOStreams,
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		localizer:  f.Localizer,
	}

	cmd := &cobra.Command{
		Use:     "create",
		Short:   f.Localizer.MustLocalize("registry.cmd.create.shortDescription"),
		Long:    f.Localizer.MustLocalize("registry.cmd.create.longDescription"),
		Example: f.Localizer.MustLocalize("registry.cmd.create.example"),
		Args:    cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				opts.name = args[0]

				if err := serviceregistry.ValidateName(opts.name); err != nil {
					return err
				}
			}

			if !opts.IO.CanPrompt() && opts.name == "" {
				return errors.New(opts.localizer.MustLocalize("registry.cmd.create.error.name.requiredWhenNonInteractive"))
			} else if opts.name == "" {
				opts.interactive = true
			}

			validOutputFormats := flagutil.ValidOutputFormats
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flag.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			return runCreate(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "json", opts.localizer.MustLocalize("registry.cmd.flag.output.description"))
	cmd.Flags().BoolVar(&opts.autoUse, "use", true, opts.localizer.MustLocalize("registry.cmd.create.flag.use.description"))

	flagutil.EnableOutputFlagCompletion(cmd)

	return cmd
}

func runCreate(opts *Options) error {
	cfg, err := opts.Config.Load()
	if err != nil {
		return err
	}

	var payload *srsmgmtv1.RegistryCreateRest
	if opts.interactive {
		opts.Logger.Debug()

		payload, err = promptPayload(opts)
		if err != nil {
			return err
		}
	} else {
		payload = &srsmgmtv1.RegistryCreateRest{
			Name: &opts.name,
		}
	}

	connection, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	// the user must have accepted the terms and conditions from the provider
	// before they can create a registry instance
	termsAccepted, termsURL, err := ams.CheckTermsAccepted(connection)
	if err != nil {
		return err
	}
	if !termsAccepted && termsURL != "" {
		opts.Logger.Info(opts.localizer.MustLocalize("service.info.termsCheck", localize.NewEntry("TermsURL", termsURL)))
		return nil
	}

	opts.Logger.Info(opts.localizer.MustLocalize("registry.cmd.create.info.action", localize.NewEntry("Name", payload.GetName())))

	response, _, err := connection.API().
		ServiceRegistryMgmt().
		CreateRegistry(context.Background()).
		RegistryCreateRest(*payload).
		Execute()
	if err != nil {
		return err
	}

	opts.Logger.Info(opts.localizer.MustLocalize("registry.cmd.create.info.successMessage"))

	dump.PrintDataInFormat(opts.outputFormat, response, opts.IO.Out)

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
func promptPayload(opts *Options) (payload *srsmgmtv1.RegistryCreateRest, err error) {
	if err != nil {
		return nil, err
	}

	// set type to store the answers from the prompt with defaults
	answers := struct {
		Name string
	}{}

	promptName := &survey.Input{
		Message: opts.localizer.MustLocalize("registry.cmd.create.input.name.message"),
		Help:    opts.localizer.MustLocalize("registry.cmd.create.input.name.help"),
	}

	err = survey.AskOne(promptName, &answers.Name, survey.WithValidator(serviceregistry.ValidateName))
	if err != nil {
		return nil, err
	}

	payload = &srsmgmtv1.RegistryCreateRest{
		Name: &answers.Name,
	}

	return payload, nil
}
