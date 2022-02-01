package create

import (
	"context"
	"fmt"
	"os"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/serviceaccount/svcaccountcmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/serviceaccount/svcaccountcmdutil/credentials"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/serviceaccount/svcaccountcmdutil/validation"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/color"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/spinner"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"

	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

type options struct {
	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	localizer  localize.Localizer
	Context    context.Context

	fileFormat       string
	overwrite        bool
	shortDescription string
	filename         string

	interactive bool
}

// NewCreateCommand creates a new command to create service accounts
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
		Short:   opts.localizer.MustLocalize("serviceAccount.create.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("serviceAccount.create.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("serviceAccount.create.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) (err error) {
			if !opts.IO.CanPrompt() && opts.shortDescription == "" {
				return opts.localizer.MustLocalizeError("flag.error.requiredWhenNonInteractive", localize.NewEntry("Flag", "short-description"))
			} else if opts.shortDescription == "" {
				opts.interactive = true
			}

			if !opts.interactive {
				validator := &validation.Validator{
					Localizer: opts.localizer,
				}

				if opts.fileFormat == "" {
					return opts.localizer.MustLocalizeError("flag.error.requiredWhenNonInteractive", localize.NewEntry("Flag", "file-format"))
				}

				if err = validator.ValidateShortDescription(opts.shortDescription); err != nil {
					return err
				}
			}

			// check that a valid --file-format flag value is used
			validOutput := flagutil.IsValidInput(opts.fileFormat, svcaccountcmdutil.CredentialsOutputFormats...)
			if !validOutput && opts.fileFormat != "" {
				return flagutil.InvalidValueError("file-format", opts.fileFormat, svcaccountcmdutil.CredentialsOutputFormats...)
			}

			return runCreate(opts)
		},
	}

	cmd.Flags().StringVar(&opts.shortDescription, "short-description", "", opts.localizer.MustLocalize("serviceAccount.create.flag.shortDescription.description"))
	cmd.Flags().BoolVar(&opts.overwrite, "overwrite", false, opts.localizer.MustLocalize("serviceAccount.common.flag.overwrite.description"))
	cmd.Flags().StringVar(&opts.filename, "output-file", "", opts.localizer.MustLocalize("serviceAccount.common.flag.fileLocation.description"))
	cmd.Flags().StringVar(&opts.fileFormat, "file-format", "", opts.localizer.MustLocalize("serviceAccount.common.flag.fileFormat.description"))

	flagutil.EnableStaticFlagCompletion(cmd, "file-format", svcaccountcmdutil.CredentialsOutputFormats)

	return cmd
}

// nolint:funlen
func runCreate(opts *options) error {
	conn, err := opts.Connection()
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
		return opts.localizer.MustLocalizeError("serviceAccount.common.error.credentialsFileAlreadyExists", localize.NewEntry("FilePath", opts.filename))
	}

	spinner := spinner.New(opts.IO.ErrOut, opts.localizer)
	spinner.SetSuffix(opts.localizer.MustLocalize("serviceAccount.create.log.info.creating"))
	spinner.Start()
	// create the service account
	serviceAccountPayload := kafkamgmtclient.ServiceAccountRequest{Name: opts.shortDescription}

	serviceacct, httpRes, err := conn.API().
		ServiceAccountMgmt().
		CreateServiceAccount(opts.Context).
		ServiceAccountRequest(serviceAccountPayload).
		Execute()
	spinner.Stop()

	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	if err != nil {
		return err
	}

	opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("serviceAccount.create.log.info.createdSuccessfully", localize.NewEntry("ID", serviceacct.GetId())))

	creds := &credentials.Credentials{
		ClientID:     serviceacct.GetClientId(),
		ClientSecret: serviceacct.GetClientSecret(),
		// TODO new location of the token url? From where we can take this value?
		// TokenURL:     cfg.MasAuthURL + "/protocol/openid-connect/token",
	}

	// save the credentials to a file
	err = credentials.Write(opts.fileFormat, opts.filename, creds)
	if err != nil {
		return fmt.Errorf("%v: %w", opts.localizer.MustLocalize("serviceAccount.common.error.couldNotSaveCredentialsFile"), err)
	}

	opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("serviceAccount.common.log.info.credentialsSaved",
		localize.NewEntry("FilePath", color.CodeSnippet(opts.filename)),
		localize.NewEntry("ClientID", color.Success(creds.ClientID)),
	))

	return nil
}

func runInteractivePrompt(opts *options) (err error) {
	_, err = opts.Connection()
	if err != nil {
		return err
	}

	validator := &validation.Validator{
		Localizer: opts.localizer,
	}

	opts.Logger.Debug(opts.localizer.MustLocalize("common.log.debug.startingInteractivePrompt"))

	promptName := &survey.Input{
		Message: opts.localizer.MustLocalize("serviceAccount.create.input.shortDescription.message"),
		Help:    opts.localizer.MustLocalize("serviceAccount.create.input.shortDescription.help"),
	}

	err = survey.AskOne(promptName, &opts.shortDescription, survey.WithValidator(survey.Required), survey.WithValidator(validator.ValidateShortDescription))
	if err != nil {
		return err
	}

	// if the --file-format flag was not used, ask in the prompt
	if opts.fileFormat == "" {
		opts.Logger.Debug(opts.localizer.MustLocalize("serviceAccount.common.log.debug.interactive.fileFormatNotSet"))

		fileFormatPrompt := &survey.Select{
			Message: opts.localizer.MustLocalize("serviceAccount.create.input.fileFormat.message"),
			Help:    opts.localizer.MustLocalize("serviceAccount.create.input.fileFormat.help"),
			Options: svcaccountcmdutil.CredentialsOutputFormats,
			Default: credentials.EnvFormat,
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

	return nil
}
