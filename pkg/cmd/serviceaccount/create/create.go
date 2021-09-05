package create

import (
	"context"
	"errors"
	"fmt"
	"github.com/redhat-developer/app-services-cli/internal/build"
	"os"

	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/serviceaccount/validation"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"

	"github.com/redhat-developer/app-services-cli/pkg/connection"

	"github.com/AlecAivazis/survey/v2"
	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/serviceaccount/credentials"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	"github.com/spf13/cobra"
)

type Options struct {
	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	localizer  localize.Localizer

	fileFormat  string
	overwrite   bool
	name        string
	description string
	filename    string

	interactive bool
}

// NewCreateCommand creates a new command to create service accounts
func NewCreateCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		IO:         f.IOStreams,
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		localizer:  f.Localizer,
	}

	cmd := &cobra.Command{
		Use:     opts.localizer.MustLocalize("serviceAccount.create.cmd.use"),
		Short:   opts.localizer.MustLocalize("serviceAccount.create.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("serviceAccount.create.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("serviceAccount.create.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) (err error) {
			if !opts.IO.CanPrompt() && opts.name == "" {
				return errors.New(opts.localizer.MustLocalize("flag.error.requiredWhenNonInteractive", localize.NewEntry("Flag", "name")))
			} else if opts.name == "" && opts.description == "" {
				opts.interactive = true
			}

			if !opts.interactive {

				validator := &validation.Validator{
					Localizer: opts.localizer,
				}

				if opts.fileFormat == "" {
					return errors.New(opts.localizer.MustLocalize("flag.error.requiredWhenNonInteractive", localize.NewEntry("Flag", "file-format")))
				}

				if err = validator.ValidateName(opts.name); err != nil {
					return err
				}
				if err = validator.ValidateDescription(opts.description); err != nil {
					return err
				}
			}

			// check that a valid --file-format flag value is used
			validOutput := flagutil.IsValidInput(opts.fileFormat, flagutil.CredentialsOutputFormats...)
			if !validOutput && opts.fileFormat != "" {
				return flag.InvalidValueError("file-format", opts.fileFormat, flagutil.CredentialsOutputFormats...)
			}

			return runCreate(opts)
		},
	}

	cmd.Flags().StringVar(&opts.name, "name", "", opts.localizer.MustLocalize("serviceAccount.create.flag.name.description"))
	cmd.Flags().StringVar(&opts.description, "description", "", opts.localizer.MustLocalize("serviceAccount.create.flag.description.description"))
	cmd.Flags().BoolVar(&opts.overwrite, "overwrite", false, opts.localizer.MustLocalize("serviceAccount.common.flag.overwrite.description"))
	cmd.Flags().StringVar(&opts.filename, "output-file", "", opts.localizer.MustLocalize("serviceAccount.common.flag.fileLocation.description"))
	cmd.Flags().StringVar(&opts.fileFormat, "file-format", "", opts.localizer.MustLocalize("serviceAccount.common.flag.fileFormat.description"))

	flagutil.EnableStaticFlagCompletion(cmd, "file-format", flagutil.CredentialsOutputFormats)

	return cmd
}

// nolint:funlen
func runCreate(opts *Options) error {
	conn, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	if opts.interactive {
		// run the create command interactively
		err = runInteractivePrompt(opts)
		if err != nil {
			return err
		}
	} else if opts.filename == "" {
		// obtain the absolute path to where credentials will be saved
		opts.filename = credentials.GetDefaultPath(opts.fileFormat)
	}

	// If the credentials file already exists, and the --overwrite flag is not set then return an error
	// indicating that the user should explicitly request overwriting of the file
	_, err = os.Stat(opts.filename)
	if err == nil && !opts.overwrite {
		return errors.New(opts.localizer.MustLocalize("serviceAccount.common.error.credentialsFileAlreadyExists", localize.NewEntry("FilePath", opts.filename)))
	}

	// create the service account
	serviceAccountPayload := &kafkamgmtclient.ServiceAccountRequest{Name: opts.name, Description: &opts.description}

	a := conn.API().ServiceAccount().CreateServiceAccount(context.Background())
	a = a.ServiceAccountRequest(*serviceAccountPayload)
	serviceacct, _, err := a.Execute()
	if err != nil {
		return err
	}

	opts.Logger.Info(build.GetEmojiSuccess(), opts.localizer.MustLocalize("serviceAccount.create.log.info.createdSuccessfully", localize.NewEntry("ID", serviceacct.GetId()), localize.NewEntry("Name", serviceacct.GetName())))

	creds := &credentials.Credentials{
		ClientID:     serviceacct.GetClientId(),
		ClientSecret: serviceacct.GetClientSecret(),
	}

	// save the credentials to a file
	err = credentials.Write(opts.fileFormat, opts.filename, creds)
	if err != nil {
		return fmt.Errorf("%v: %w", opts.localizer.MustLocalize("serviceAccount.common.error.couldNotSaveCredentialsFile"), err)
	}

	opts.Logger.Info(opts.localizer.MustLocalize("serviceAccount.common.log.info.credentialsSaved", localize.NewEntry("FilePath", opts.filename)))

	return nil
}

func runInteractivePrompt(opts *Options) (err error) {
	_, err = opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	validator := &validation.Validator{
		Localizer: opts.localizer,
	}

	opts.Logger.Debug(opts.localizer.MustLocalize("common.log.debug.startingInteractivePrompt"))

	promptName := &survey.Input{
		Message: opts.localizer.MustLocalize("serviceAccount.create.input.name.message"),
		Help:    opts.localizer.MustLocalize("serviceAccount.create.input.name.help"),
	}

	err = survey.AskOne(promptName, &opts.name, survey.WithValidator(survey.Required), survey.WithValidator(validator.ValidateName))
	if err != nil {
		return err
	}

	// if the --file-format flag was not used, ask in the prompt
	if opts.fileFormat == "" {
		opts.Logger.Debug(opts.localizer.MustLocalize("serviceAccount.common.log.debug.interactive.fileFormatNotSet"))

		fileFormatPrompt := &survey.Select{
			Message: opts.localizer.MustLocalize("serviceAccount.create.input.fileFormat.message"),
			Help:    opts.localizer.MustLocalize("serviceAccount.create.input.fileFormat.help"),
			Options: flagutil.CredentialsOutputFormats,
			Default: "env",
		}

		err = survey.AskOne(fileFormatPrompt, &opts.fileFormat)
		if err != nil {
			return err
		}
	}

	opts.filename, opts.overwrite, err = credentials.ChooseFileLocation(opts.fileFormat, opts.filename, opts.overwrite)
	if err != nil {
		return err
	}

	promptDescription := &survey.Multiline{
		Message: opts.localizer.MustLocalize("serviceAccount.create.input.description.message"),
		Help:    opts.localizer.MustLocalize("serviceAccount.create.flag.description.description"),
	}

	err = survey.AskOne(promptDescription, &opts.description, survey.WithValidator(validator.ValidateDescription))
	if err != nil {
		return err
	}

	opts.Logger.Info(opts.localizer.MustLocalize("serviceAccount.create.log.info.creating", localize.NewEntry("Name", opts.name)))

	return nil
}
