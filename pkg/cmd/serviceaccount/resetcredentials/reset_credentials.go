package resetcredentials

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/AlecAivazis/survey/v2"
	flagutil "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil/flags"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/serviceaccount/credentials"

	"github.com/MakeNowJust/heredoc"
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/managedservices"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"
	"github.com/spf13/cobra"
)

type Options struct {
	Config     config.IConfig
	Connection func() (connection.Connection, error)
	Logger     func() (logging.Logger, error)

	id        string
	force     bool
	output    string
	overwrite bool
	filename  string
}

// NewResetCredentialsCommand creates a new command to delete a service account
func NewResetCredentialsCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
	}

	cmd := &cobra.Command{
		Use:   "reset-credentials",
		Short: "Reset credentials for a service account",
		Long:  "Generate new SASL/PLAIN credentials for a service account and revoke the old credentials",
		Example: heredoc.Doc(`
			$ rhoas serviceaccount reset-credentials --id 173c1ad9-932d-4007-ae0f-4da74f4d2ccd -o json
		`),
		RunE: func(cmd *cobra.Command, _ []string) error {
			validOutput := flagutil.IsValidInput(opts.output, flagutil.CredentialsOutputFormats...)
			if !validOutput {
				return fmt.Errorf("Invalid output format '%v'", opts.output)
			}

			return runResetCredentials(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.output, "output", "o", "env", "Format of the config [env, kafka, properties, json, kube]")
	cmd.Flags().StringVar(&opts.id, "id", "", "The unique ID of the service account to delete")
	cmd.Flags().BoolVarP(&opts.force, "force", "f", false, "Forcefully reset credentials for the service account")
	cmd.Flags().BoolVar(&opts.overwrite, "overwrite", false, "Force overwrite a file if it already exists")
	cmd.Flags().StringVar(&opts.filename, "file-location", "", "Sets a custom file location to save the credentials")

	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("output")

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

	// obtain the absolute path to where credentials will be saved
	fileLoc := credentials.AbsolutePath(opts.output, opts.filename)

	// check if the parent directory exists. If not return an error
	fileDir := path.Dir(fileLoc)
	_, err = os.Stat(fileDir)
	if err != nil {
		return err
	}

	// If the credentials file already exists, and the --overwrite flag is not set then return an error
	// indicating that the user should explicitly request overwriting of the file
	if _, err = os.Stat(fileLoc); err == nil && !opts.overwrite {
		return fmt.Errorf("file '%v' already exists. Use --overwrite to overwrite the file, or --file-location to choose a custom location", fileLoc)
	}

	var serviceacct *managedservices.ServiceAccount
	if opts.force {
		// If the --force flag is set, skip the interactive confirmation prompt
		// and go straight to resetting credentials
		logger.Debug("Forcefully resetting credentials for service account with ID", opts.id)
		serviceacct, err = resetCredentials(opts)
		if err != nil {
			return err
		}
	} else {
		// prompt the user to confirm their wish to proceed with this action
		var confirmDelete bool
		promptConfirmDelete := &survey.Confirm{
			Message: "Are you sure you want to reset the credentials for this service account?",
		}

		if err = survey.AskOne(promptConfirmDelete, &confirmDelete); err != nil {
			return err
		}

		serviceacct, err = resetCredentials(opts)
		if err != nil {
			return err
		}
	}

	creds := &credentials.Credentials{
		ClientID:     serviceacct.GetClientID(),
		ClientSecret: serviceacct.GetClientSecret(),
	}
	if logger.DebugEnabled() {
		b, _ := json.Marshal(creds)
		logger.Debug("Credentials reset:", string(b))
	}

	// save the credentials to a file
	err = credentials.Write(opts.output, fileLoc, creds)
	if err != nil {
		return err
	}

	logger.Info("Credentials saved to", fileLoc)

	return nil
}

func resetCredentials(opts *Options) (*managedservices.ServiceAccount, error) {
	connection, err := opts.Connection()
	if err != nil {
		return nil, err
	}

	logger, err := opts.Logger()
	if err != nil {
		return nil, err
	}

	client := connection.NewAPIClient()
	a := client.DefaultApi.ResetServiceAccountCreds(context.Background(), opts.id)

	logger.Debug("Resetting credentials for service account with ID", opts.id)
	serviceacct, _, apiErr := a.Execute()

	if apiErr.Error() != "" {
		return nil, apiErr
	}

	return &serviceacct, nil
}
