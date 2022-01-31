package resetcredentials

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/serviceaccount/svcaccountcmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/serviceaccount/svcaccountcmdutil/credentials"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/serviceaccount/svcaccountcmdutil/validation"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/factory"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/color"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
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

	id         string
	fileFormat string
	overwrite  bool
	filename   string

	interactive bool
	force       bool
}

// NewResetCredentialsCommand creates a new command to delete a service account
func NewResetCredentialsCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		IO:         f.IOStreams,
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "reset-credentials",
		Short:   opts.localizer.MustLocalize("serviceAccount.resetCredentials.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("serviceAccount.resetCredentials.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("serviceAccount.resetCredentials.cmd.example"),
		RunE: func(cmd *cobra.Command, _ []string) error {
			if !opts.IO.CanPrompt() && opts.id == "" {
				return opts.localizer.MustLocalizeError("flag.error.requiredWhenNonInteractive", localize.NewEntry("Flag", "id"))
			} else if opts.id == "" {
				opts.interactive = true
			}

			if !opts.interactive && opts.fileFormat == "" {
				return opts.localizer.MustLocalizeError("flag.error.requiredWhenNonInteractive", localize.NewEntry("Flag", "file-format"))
			}

			validOutput := flagutil.IsValidInput(opts.fileFormat, svcaccountcmdutil.CredentialsOutputFormats...)
			if !validOutput && opts.fileFormat != "" {
				return flagutil.InvalidValueError("file-format", opts.fileFormat, svcaccountcmdutil.CredentialsOutputFormats...)
			}

			if !opts.interactive {
				validator := &validation.Validator{
					Localizer: opts.localizer,
				}

				validID := validator.ValidateUUID(opts.id)
				if validID != nil {
					return validID
				}
			}

			return runResetCredentials(opts)
		},
	}

	cmd.Flags().StringVar(&opts.id, "id", "", opts.localizer.MustLocalize("serviceAccount.resetCredentials.flag.id.description"))
	cmd.Flags().BoolVar(&opts.overwrite, "overwrite", false, opts.localizer.MustLocalize("serviceAccount.common.flag.overwrite.description"))
	cmd.Flags().StringVar(&opts.filename, "output-file", "", opts.localizer.MustLocalize("serviceAccount.common.flag.fileLocation.description"))
	cmd.Flags().StringVar(&opts.fileFormat, "file-format", "", opts.localizer.MustLocalize("serviceAccount.common.flag.fileFormat.description"))
	cmd.Flags().BoolVarP(&opts.force, "yes", "y", false, opts.localizer.MustLocalize("serviceAccount.resetCredentials.flag.yes.description"))

	flagutil.EnableStaticFlagCompletion(cmd, "file-format", svcaccountcmdutil.CredentialsOutputFormats)

	return cmd
}

// nolint:funlen
func runResetCredentials(opts *options) (err error) {
	conn, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	api := conn.API()

	_, httpRes, err := api.ServiceAccountMgmt().GetServiceAccountById(opts.Context, opts.id).Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	if err != nil {
		return err
	}
	if opts.interactive {
		err = runInteractivePrompt(opts)
		if err != nil {
			return err
		}
	} else if opts.filename == "" {
		// obtain the default absolute path to where credentials will be saved
		opts.filename = credentials.GetDefaultPath(opts.fileFormat)
	}

	// If the credentials file already exists, and the --overwrite flag is not set then return an error
	// indicating that the user should explicitly request overwriting of the file
	if _, err = os.Stat(opts.filename); err == nil && !opts.overwrite {
		return opts.localizer.MustLocalizeError("serviceAccount.common.error.credentialsFileAlreadyExists", localize.NewEntry("FilePath", color.CodeSnippet(opts.filename)))
	}

	if !opts.force {
		// prompt the user to confirm their wish to proceed with this action
		var confirmReset bool
		opts.localizer.MustLocalize("serviceAccount.resetCredentials.input.confirmReset.message", localize.NewEntry("ID", opts.id))
		promptConfirmDelete := &survey.Confirm{
			Message: opts.localizer.MustLocalize("serviceAccount.resetCredentials.input.confirmReset.message", localize.NewEntry("ID", opts.id)),
		}

		if err = survey.AskOne(promptConfirmDelete, &confirmReset); err != nil {
			return err
		}
		if !confirmReset {
			opts.Logger.Debug(opts.localizer.MustLocalize("serviceAccount.resetCredentials.log.debug.cancelledReset"))
			return nil
		}
	}

	updatedServiceAccount, err := resetCredentials(opts)
	if err != nil {
		return fmt.Errorf("%v: %w", opts.localizer.MustLocalize("serviceAccount.resetCredentials.error.resetError", localize.NewEntry("ID", opts.id)), err)
	}

	opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("serviceAccount.resetCredentials.log.info.resetSuccess", localize.NewEntry("ID", updatedServiceAccount.GetId())))

	cfg, err := opts.Config.Load()
	if err != nil {
		return err
	}

	creds := &credentials.Credentials{
		ClientID:     updatedServiceAccount.GetClientId(),
		ClientSecret: updatedServiceAccount.GetClientSecret(),
		TokenURL:     cfg.MasAuthURL + "/protocol/openid-connect/token",
	}

	// save the credentials to a file
	err = credentials.Write(opts.fileFormat, opts.filename, creds)
	if err != nil {
		return err
	}

	opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("serviceAccount.common.log.info.credentialsSaved", localize.NewEntry("FilePath", opts.filename)))

	return nil
}

func resetCredentials(opts *options) (*kafkamgmtclient.ServiceAccount, error) {
	conn, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return nil, err
	}

	// check if the service account exists
	api := conn.API()

	opts.Logger.Debug(opts.localizer.MustLocalize("serviceAccount.resetCredentials.log.debug.resettingCredentials", localize.NewEntry("ID", opts.id)))

	serviceacct, httpRes, err := api.ServiceAccountMgmt().ResetServiceAccountCreds(opts.Context, opts.id).Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}
	if err != nil {
		if httpRes == nil {
			return nil, err
		}

		switch httpRes.StatusCode {
		case http.StatusForbidden:
			opts.localizer.MustLocalize("serviceAccount.common.error.forbidden", localize.NewEntry("Operation", "update"))
			return nil, fmt.Errorf("%v: %w", opts.localizer.MustLocalize("serviceAccount.common.error.forbidden", localize.NewEntry("Operation", "update")), err)
		case http.StatusInternalServerError:
			return nil, opts.localizer.MustLocalizeError("serviceAccount.common.error.internalServerError")
		default:
			return nil, err
		}
	}

	return &serviceacct, nil
}

func runInteractivePrompt(opts *options) (err error) {
	_, err = opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	opts.Logger.Debug(opts.localizer.MustLocalize("common.log.debug.startingInteractivePrompt"))

	promptID := &survey.Input{
		Message: opts.localizer.MustLocalize("serviceAccount.resetCredentials.input.id.message"),
		Help:    opts.localizer.MustLocalize("serviceAccount.resetCredentials.input.id.help"),
	}

	validator := &validation.Validator{
		Localizer: opts.localizer,
	}

	err = survey.AskOne(promptID, &opts.id, survey.WithValidator(survey.Required), survey.WithValidator(validator.ValidateUUID))
	if err != nil {
		return err
	}

	// if the --output flag was not used, ask in the prompt
	if opts.fileFormat == "" {
		opts.Logger.Debug(opts.localizer.MustLocalize("serviceAccount.common.log.debug.interactive.fileFormatNotSet"))

		fileFormatPrompt := &survey.Select{
			Message: opts.localizer.MustLocalize("serviceAccount.resetCredentials.input.fileFormat.message"),
			Help:    opts.localizer.MustLocalize("serviceAccount.resetCredentials.input.fileFormat.help"),
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

	return err
}
