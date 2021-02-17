package resetcredentials

import (
	"context"
	"fmt"
	"os"

	kasclient "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/kas/client"

	"github.com/AlecAivazis/survey/v2"
	flagutil "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil/flags"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/iostreams"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/serviceaccount/credentials"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"
	"github.com/spf13/cobra"
)

type Options struct {
	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection func() (connection.Connection, error)
	Logger     func() (logging.Logger, error)

	id         string
	fileFormat string
	overwrite  bool
	filename   string

	interactive bool
}

// NewResetCredentialsCommand creates a new command to delete a service account
func NewResetCredentialsCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		IO:         f.IOStreams,
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
	}

	localizer.LoadMessageFiles("cmd/common", "cmd/serviceaccount", "cmd/serviceaccount/reset-credentials")

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
				return fmt.Errorf(localizer.MustLocalize(&localizer.Config{
					MessageID: "flag.error.invalidValue",
					TemplateData: map[string]interface{}{
						"Flag":  "file-format",
						"Value": opts.fileFormat,
					},
				}))
			}

			return runResetCredentials(opts)
		},
	}

	cmd.Flags().StringVar(&opts.id, "id", "", localizer.MustLocalizeFromID("serviceAccount.resetCredentials.flag.id.description"))
	cmd.Flags().BoolVar(&opts.overwrite, "overwrite", false, localizer.MustLocalizeFromID("serviceAccount.common.flag.overwrite.description"))
	cmd.Flags().StringVar(&opts.filename, "file-location", "", localizer.MustLocalizeFromID("serviceAccount.common.flag.fileLocation.description"))
	cmd.Flags().StringVar(&opts.fileFormat, "file-format", "", localizer.MustLocalizeFromID("serviceAccount.common.flag.fileFormat.description"))

	return cmd
}

// nolint:funlen
func runResetCredentials(opts *Options) (err error) {
	connection, err := opts.Connection()
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
	} else {
		// obtain the absolute path to where credentials will be saved
		opts.filename = credentials.AbsolutePath(opts.fileFormat, opts.filename)

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
	}

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
	connection, err := opts.Connection()
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

	serviceacct, _, apiErr := api.Kafka().ResetServiceAccountCreds(context.Background(), opts.id).Execute()

	if apiErr.Error() != "" {
		return nil, apiErr
	}

	return &serviceacct, nil
}

func runInteractivePrompt(opts *Options) (err error) {
	_, err = opts.Connection()
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
