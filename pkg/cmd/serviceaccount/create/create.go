package create

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path"

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

	output      string
	overwrite   bool
	name        string
	description string
	filename    string
}

// NewCreateCommand creates a new command to create service accounts
func NewCreateCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
	}

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new service account with credentials",
		Long: heredoc.Doc(`
			Creates a new service account with SASL/PLAIN credentials to connect your application to services.
			The credentials are saved to an output format of your choice from the list below with the --output flag:
				
				- env (default): Saves the credentials to a .env file as environment variables
				- kafka, properties: Saves the credentials to a kafka.properties file
				- json: Saves the credentials in a credentials.json JSON file
				- kube: Saves credentials as Kubernetes secret
		`),
		Example: heredoc.Doc(`
			$ rhoas serviceaccount create --output kafka
			$ rhoas serviceaccount create --output=env
			$ rhoas serviceaccount create -o=properties
			$ rhoas serviceaccount create -o env --force
			$ rhoas serviceaccount create -o json`,
		),
		RunE: func(cmd *cobra.Command, _ []string) error {
			validOutput := flagutil.IsValidInput(opts.output, flagutil.CredentialsOutputFormats...)
			if !validOutput {
				return fmt.Errorf("Invalid output format '%v'", opts.output)
			}

			return runCreate(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.output, "output", "o", "env", "Format of the config [env, kafka, properties, json, kube]")
	cmd.Flags().StringVar(&opts.name, "name", "", "Name of the service account")
	cmd.Flags().StringVar(&opts.description, "description", "", "Description for the service account")
	cmd.Flags().BoolVar(&opts.overwrite, "overwrite", false, "Force overwrite a file if it already exists")
	cmd.Flags().StringVar(&opts.filename, "file-location", "", "Sets a custom file location to save the credentials")

	_ = cmd.MarkFlagRequired("output")
	_ = cmd.MarkFlagRequired("name")

	return cmd
}

func runCreate(opts *Options) error {
	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	connection, err := opts.Connection()
	if err != nil {
		return err
	}

	client := connection.NewAPIClient()

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
	_, err = os.Stat(fileLoc)
	if err == nil && !opts.overwrite {
		return fmt.Errorf("file '%v' already exists. Use --force-overwrite to overwrite the file, or --file-location flag to choose a custom location", fileLoc)
	}

	// create the service account
	svcAcctPayload := &managedservices.ServiceAccountRequest{Name: opts.name, Description: &opts.description}
	a := client.DefaultApi.CreateServiceAccount(context.Background())
	a = a.ServiceAccountRequest(*svcAcctPayload)
	serviceacct, _, apiErr := a.Execute()

	if apiErr.Error() != "" {
		return fmt.Errorf("Could not create service account: %w", apiErr)
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
		return fmt.Errorf("Could not save credentials to file: %w", err)
	}

	logger.Info("Credentials saved to", fileLoc)

	return nil
}
