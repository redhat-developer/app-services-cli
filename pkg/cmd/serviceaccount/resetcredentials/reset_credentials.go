package resetcredentials

import (
	"context"
	"fmt"
	"os"

	kasclient "github.com/redhat-developer/app-services-cli/pkg/api/kas/client"
	"github.com/redhat-developer/app-services-cli/pkg/connection"

	"github.com/AlecAivazis/survey/v2"
	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/serviceaccount/credentials"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/internal/localizer"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	"github.com/spf13/cobra"
)

type Options struct {
	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     func() (logging.Logger, error)

	id         string
	fileFormat string
	overwrite  bool
	filename   string

	interactive bool
	force       bool
}

// NewResetCredentialsCommand creates a new command to delete a service account
func NewResetCredentialsCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		IO:         f.IOStreams,
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
	}

	cmd := &cobra.Command{
		Use:     localizer.MustLocalizeFromID("serviceAccount.resetCredentials.cmd.use"),
		Short:   localizer.MustLocalizeFromID("serviceAccount.resetCredentials.cmd.shortDescription"),
		Long:    localizer.MustLocalizeFromID("serviceAccount.resetCredentials.cmd.longDescription"),
		Example: localizer.MustLocalizeFromID("serviceAccount.resetCredentials.cmd.example"),
		RunE: func(cmd *cobra.Command, _ []string) error {
			if !opts.IO.CanPrompt() && opts.id == "" {
				return fmt.Errorf(localizer.MustLocalize(&localizer.Config{
					MessageID: "flag.error.requiredWhenNonInteractive",
					TemplateData: map[string]interface{}{
						"Flag": "id",
					},
				}))
			} else if opts.id == "" {
				opts.interactive = true
			}

			if !opts.interactive && opts.fileFormat == "" {
				return fmt.Errorf(localizer.MustLocalize(&localizer.Config{
					MessageID: "flag.error.requiredWhenNonInteractive",
					TemplateData: map[string]interface{}{
						"Flag": "file-format",
					},
				}))
			}

			validOutput := flagutil.IsValidInput(opts.fileFormat, flagutil.CredentialsOutputFormats...)
			if !validOutput && opts.fileFormat != "" {
				return flag.InvalidValueError("file-format", opts.fileFormat, flagutil.CredentialsOutputFormats...)
			}

			return runResetCredentials(opts)
		},
	}

	cmd.Flags().StringVar(&opts.id, "id", "", localizer.MustLocalizeFromID("serviceAccount.resetCredentials.flag.id.description"))
	cmd.Flags().BoolVar(&opts.overwrite, "overwrite", false, localizer.MustLocalizeFromID("serviceAccount.common.flag.overwrite.description"))
	cmd.Flags().StringVar(&opts.filename, "file-location", "", localizer.MustLocalizeFromID("serviceAccount.common.flag.fileLocation.description"))
	cmd.Flags().StringVar(&opts.fileFormat, "file-format", "", localizer.MustLocalizeFromID("serviceAccount.common.flag.fileFormat.description"))
	cmd.Flags().BoolVarP(&opts.force, "yes", "y", false, localizer.MustLocalizeFromID("serviceAccount.resetCredentials.flag.yes.description"))

	return cmd
}

// nolint:funlen
func runResetCredentials(opts *Options) (err error) {
	connection, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	api := connection.API()

	serviceacct, _, apiErr := api.Kafka().GetServiceAccountById(context.Background(), opts.id).Execute()
	if apiErr.Error() != "" {
		return apiErr
	}
	serviceAcctName := serviceacct.GetName()

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
		return fmt.Errorf(localizer.MustLocalize(&localizer.Config{
			MessageID: "serviceAccount.common.error.credentialsFileAlreadyExists",
			TemplateData: map[string]interface{}{
				"FilePath": opts.filename,
			},
		}))
	}

	if !opts.force {
		// prompt the user to confirm their wish to proceed with this action
		var confirmReset bool
		promptConfirmDelete := &survey.Confirm{
			Message: localizer.MustLocalize(&localizer.Config{
				MessageID: "serviceAccount.resetCredentials.input.confirmReset.message",
				TemplateData: map[string]interface{}{
					"ID": opts.id,
				},
			}),
		}

		if err = survey.AskOne(promptConfirmDelete, &confirmReset); err != nil {
			return err
		}
		if !confirmReset {
			logger.Debug(localizer.MustLocalizeFromID("serviceAccount.resetCredentials.log.debug.cancelledReset"))
			return nil
		}
	}

	updatedServiceAccount, err := resetCredentials(serviceAcctName, opts)

	if err != nil {
		return fmt.Errorf("%v: %w", localizer.MustLocalize(&localizer.Config{
			MessageID: "serviceAccount.resetCredentials.error.resetError",
			TemplateData: map[string]interface{}{
				"Name": updatedServiceAccount.GetName(),
			},
		}), err)
	}

	logger.Info(localizer.MustLocalize(&localizer.Config{
		MessageID: "serviceAccount.resetCredentials.log.info.resetSuccess",
		TemplateData: map[string]interface{}{
			"Name": updatedServiceAccount.GetName(),
		},
	}))

	creds := &credentials.Credentials{
		ClientID:     updatedServiceAccount.GetClientID(),
		ClientSecret: updatedServiceAccount.GetClientSecret(),
	}

	// save the credentials to a file
	err = credentials.Write(opts.fileFormat, opts.filename, creds)
	if err != nil {
		return err
	}

	logger.Info(localizer.MustLocalize(&localizer.Config{
		MessageID: "serviceAccount.common.log.info.credentialsSaved",
		TemplateData: map[string]interface{}{
			"FilePath": opts.filename,
		},
	}))

	return nil
}

func resetCredentials(name string, opts *Options) (*kasclient.ServiceAccount, error) {
	connection, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return nil, err
	}

	// check if the service account exists
	api := connection.API()

	logger, err := opts.Logger()
	if err != nil {
		return nil, err
	}

	logger.Debug(localizer.MustLocalize(&localizer.Config{
		MessageID: "serviceAccount.resetCredentials.log.debug.resettingCredentials",
		TemplateData: map[string]interface{}{
			"Name": name,
		},
	}))

	serviceacct, httpRes, apiErr := api.Kafka().ResetServiceAccountCreds(context.Background(), opts.id).Execute()

	if apiErr.Error() != "" {
		if httpRes == nil {
			return nil, apiErr
		}

		switch httpRes.StatusCode {
		case 403:
			return nil, fmt.Errorf("%v: %w", localizer.MustLocalize(&localizer.Config{
				MessageID: "serviceAccount.common.error.forbidden",
				TemplateData: map[string]interface{}{
					"Operation": "update",
				},
			}), apiErr)
		case 500:
			return nil, fmt.Errorf("%v: %w", localizer.MustLocalizeFromID("serviceAccount.common.error.internalServerError"), apiErr)
		default:
			return nil, apiErr
		}
	}

	return &serviceacct, nil
}

func runInteractivePrompt(opts *Options) (err error) {
	_, err = opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	logger.Debug(localizer.MustLocalizeFromID("common.log.debug.startingInteractivePrompt"))

	promptID := &survey.Input{
		Message: localizer.MustLocalizeFromID("serviceAccount.resetCredentials.input.id.message"),
		Help:    localizer.MustLocalizeFromID("serviceAccount.resetCredentials.input.id.help"),
	}

	err = survey.AskOne(promptID, &opts.id, survey.WithValidator(survey.Required))
	if err != nil {
		return err
	}

	// if the --output flag was not used, ask in the prompt
	if opts.fileFormat == "" {
		logger.Debug(localizer.MustLocalizeFromID("serviceAccount.common.log.debug.interactive.fileFormatNotSet"))

		fileFormatPrompt := &survey.Select{
			Message: localizer.MustLocalizeFromID("serviceAccount.resetCredentials.input.fileFormat.message"),
			Help:    localizer.MustLocalizeFromID("serviceAccount.resetCredentials.input.fileFormat.help"),
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

	return err
}
