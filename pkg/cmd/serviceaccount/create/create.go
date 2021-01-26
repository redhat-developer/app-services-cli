package create

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/color"
	"os"

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

	output      string
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
	}

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a service account",
		Long: heredoc.Doc(`
			Create a service account with credentials which are saved to a file.
			
			The credentials generated as part of a service account can be by applications 
			and tools to authenticate and interact with services.
			
			You must specify an output format to which the credentials will be stored as.
				- env (default: Store credentials in an env file as environment variables
				- json: Store credentials in a JSON file
				- kube: Store credentials as a Kubernetes secret file (the secret is not created for you)
				- properties: Store credentials in a properties file, mainly used in Java-related technologies.
		`),
		Example: heredoc.Doc(`
			# create a service account through an interactive prompt
			$ rhoas serviceaccount create

			# create a service account and save the credentials in a Kubernetes secret file format
			$ rhoas serviceaccount create --output kube

			# create a service account and forcefully overwrite the credentials file if it exists already
			$ rhoas serviceaccount create --overwrite

			# create a service account and save credentials to a custom file location
			$ rhoas serviceaccount create --file-location=./service-acct-credentials.json
		`),
		RunE: func(cmd *cobra.Command, _ []string) error {
			if !opts.IO.CanPrompt() && opts.name == "" {
				return fmt.Errorf("--name required when not running interactively")
			} else if opts.name == "" && opts.description == "" {
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
			return runCreate(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.output, "output", "o", "", fmt.Sprintf("Format of the credentials file: %q", flagutil.CredentialsOutputFormats))
	cmd.Flags().StringVar(&opts.name, "name", "", "Name of the service account")
	cmd.Flags().StringVar(&opts.description, "description", "", "Description for the service account")
	cmd.Flags().BoolVar(&opts.overwrite, "overwrite", false, "Force overwrite a credentials file if it already exists")
	cmd.Flags().StringVar(&opts.filename, "file-location", "", "Sets a custom file location to save the credentials")

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

	if opts.interactive {
		// run the create command interactively
		err = runInteractivePrompt(opts)
		if err != nil {
			return err
		}
	} else {
		// obtain the absolute path to where credentials will be saved
		opts.filename = credentials.AbsolutePath(opts.output, opts.filename)
	}

	// If the credentials file already exists, and the --overwrite flag is not set then return an error
	// indicating that the user should explicitly request overwriting of the file
	_, err = os.Stat(opts.filename)
	if err == nil && !opts.overwrite {
		return fmt.Errorf("file %v already exists. Use --overwrite to overwrite the file, or --file-location flag to choose a custom location", color.Info(opts.filename))
	}

	// create the service account
	serviceAccountPayload := &serviceapi.ServiceAccountRequest{Name: opts.name, Description: &opts.description}

	api := connection.API()
	a := api.Kafka.CreateServiceAccount(context.Background())
	a = a.ServiceAccountRequest(*serviceAccountPayload)
	serviceacct, _, apiErr := a.Execute()

	if apiErr.Error() != "" {
		return fmt.Errorf("Could not create service account: %w", apiErr)
	}

	logger.Infof("Service account %v created", color.Info(serviceacct.GetName()))

	creds := &credentials.Credentials{
		ClientID:     serviceacct.GetClientID(),
		ClientSecret: serviceacct.GetClientSecret(),
	}

	if logger.DebugEnabled() {
		b, _ := json.Marshal(creds)
		logger.Debug("Credentials created:", string(b))
	}

	// save the credentials to a file
	err = credentials.Write(opts.output, opts.filename, creds)
	if err != nil {
		return fmt.Errorf("Could not save credentials to file: %w", err)
	}

	logger.Info("Credentials saved to", opts.filename)

	return nil
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

	promptName := &survey.Input{Message: "Name:", Help: "Give your service account an easily identifiable name"}

	err = survey.AskOne(promptName, &opts.name, survey.WithValidator(survey.Required))
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

	promptDescription := &survey.Multiline{Message: "Description (optional):"}

	err = survey.AskOne(promptDescription, &opts.description)
	if err != nil {
		return err
	}

	return nil
}
