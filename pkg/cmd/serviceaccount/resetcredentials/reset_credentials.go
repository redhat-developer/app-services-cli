package resetcredentials

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil"

	"github.com/AlecAivazis/survey/v2"
	flagutil "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil/flags"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/iostreams"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/serviceaccount/credentials"

	"github.com/MakeNowJust/heredoc"
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	serviceapi "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/serviceapi/client"
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

	id        string
	output    string
	overwrite bool
	filename  string

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

	cmd := &cobra.Command{
		Use:   "reset-credentials",
		Short: "Reset credentials for a service account",
		Long:  "Generate new SASL/PLAIN credentials for a service account and revoke the old credentials",
		Example: heredoc.Doc(`
			$ rhoas serviceaccount reset-credentials
			$ rhoas serviceaccount reset-credentials --id 173c1ad9-932d-4007-ae0f-4da74f4d2ccd -o json
		`),
		RunE: func(cmd *cobra.Command, _ []string) error {
			if !opts.IO.CanPrompt() && opts.id == "" {
				return fmt.Errorf("--id required when not running interactively")
			} else if opts.id == "" {
				opts.interactive = true
			}

			if !opts.interactive {
				if opts.output == "" {
					return fmt.Errorf("--output is a required flag")
				}
			}

			if opts.output != "" {
				// check that a valid --output flag value is used
				validOutput := flagutil.IsValidInput(opts.output, flagutil.CredentialsOutputFormats...)
				if !validOutput {
					return fmt.Errorf("Invalid value for --output. Valid values: %q", flagutil.CredentialsOutputFormats)
				}
			}

			return runResetCredentials(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.output, "output", "o", "", fmt.Sprintf("Format of the credentials file: %q", flagutil.CredentialsOutputFormats))
	cmd.Flags().StringVar(&opts.id, "id", "", "The unique ID of the service account to delete")
	cmd.Flags().BoolVar(&opts.overwrite, "overwrite", false, "Force overwrite a file if it already exists")
	cmd.Flags().StringVar(&opts.filename, "file-location", "", "Sets a custom file location to save the credentials")

	return cmd
}

func runResetCredentials(opts *Options) (err error) {
	_, err = opts.Connection()
	if err != nil {
		return err
	}

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	if opts.interactive {
		err = runInteractivePrompt(opts)
		if err = cmdutil.CheckSurveyError(err); err != nil {
			return err
		}
	} else {
		// obtain the absolute path to where credentials will be saved
		opts.filename = credentials.AbsolutePath(opts.output, opts.filename)

		// If the credentials file already exists, and the --overwrite flag is not set then return an error
		// indicating that the user should explicitly request overwriting of the file
		if _, err = os.Stat(opts.filename); err == nil && !opts.overwrite {
			return fmt.Errorf("file '%v' already exists. Use --overwrite to overwrite the file, or --file-location to choose a custom location", opts.filename)
		}
	}

	// prompt the user to confirm their wish to proceed with this action
	var confirmReset bool
	promptConfirmDelete := &survey.Confirm{
		Message: fmt.Sprintf("Are you sure you want to reset the credentials for the service account with ID '%v'?", opts.id),
	}

	if err = survey.AskOne(promptConfirmDelete, &confirmReset); err != nil {
		return err
	}
	if !confirmReset {
		logger.Debug("You have chosen to not reset the service account credentials")
		return nil
	}

	serviceacct, err := resetCredentials(opts)
	if err != nil {
		return err
	}

	creds := &credentials.Credentials{
		ClientID:     serviceacct.GetClientID(),
		ClientSecret: serviceacct.GetClientSecret(),
	}
	if logger.DebugEnabled() {
		b, _ := json.Marshal(creds)
		logger.Debug("Service account credentials reset:", string(b))
	} else {
		logger.Info("Service account credentials reset")
	}

	// save the credentials to a file
	err = credentials.Write(opts.output, opts.filename, creds)
	if err != nil {
		return err
	}

	logger.Info("Credentials saved to", opts.filename)

	return nil
}

func resetCredentials(opts *Options) (*serviceapi.ServiceAccount, error) {
	connection, err := opts.Connection()
	if err != nil {
		return nil, err
	}

	logger, err := opts.Logger()
	if err != nil {
		return nil, err
	}

	api := connection.API()
	a := api.Kafka.ResetServiceAccountCreds(context.Background(), opts.id)

	logger.Debug("Resetting credentials for service account with ID", opts.id)
	serviceacct, _, apiErr := a.Execute()

	if apiErr.Error() != "" {
		return nil, fmt.Errorf("Unable to reset service account credentials: %w", apiErr)
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

	logger.Debug("Beginning interactive prompt")

	promptID := &survey.Input{Message: "Service account ID:", Help: "What is the ID of the service account?"}

	err = survey.AskOne(promptID, &opts.id, survey.WithValidator(survey.Required))
	if err != nil {
		return err
	}

	// if the --output flag was not used, ask in the prompt
	if opts.output == "" {
		logger.Debug("--output flag is not set, prompting user to choose a value")

		outputPrompt := &survey.Select{
			Message: "Credentials output format:",
			Help:    "Output format to save the service account credentials",
			Options: flagutil.CredentialsOutputFormats,
			Default: "env",
		}

		err = survey.AskOne(outputPrompt, &opts.output)
		if err != nil {
			return err
		}
	}

	opts.filename, opts.overwrite, err = credentials.ChooseFileLocation(opts.output, opts.filename, opts.overwrite)
	if err != nil {
		return err
	}

	return err
}
