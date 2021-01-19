package create

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
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/managedservices"
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
			$ rhoas serviceaccount create
			$ rhoas serviceaccount create --output kafka
			$ rhoas serviceaccount create -o env --force
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
				// check that a valid --output flag value is used
				validOutput := flagutil.IsValidInput(opts.output, flagutil.CredentialsOutputFormats...)
				if !validOutput {
					return fmt.Errorf("Invalid value for --output. Valid values: %q", flagutil.CredentialsOutputFormats)
				}
			}

			return runCreate(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.output, "output", "o", "", "Format of the config [env, kafka, properties, json, kube]")
	cmd.Flags().StringVar(&opts.name, "name", "", "Name of the service account")
	cmd.Flags().StringVar(&opts.description, "description", "", "Description for the service account")
	cmd.Flags().BoolVar(&opts.overwrite, "overwrite", false, "Force overwrite a file if it already exists")
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

	client := connection.NewAPIClient()

	var serviceAccountPayload *managedservices.ServiceAccountRequest

	if opts.interactive {
		// run the create command interactively
		serviceAccountPayload, err = promptRequestPayload(opts)
		if err != nil {
			return err
		}
	} else {
		// obtain the absolute path to where credentials will be saved
		opts.filename = credentials.AbsolutePath(opts.output, opts.filename)

		// If the credentials file already exists, and the --overwrite flag is not set then return an error
		// indicating that the user should explicitly request overwriting of the file
		_, err = os.Stat(opts.filename)
		if err == nil && !opts.overwrite {
			return fmt.Errorf("file '%v' already exists. Use --overwrite to overwrite the file, or --file-location flag to choose a custom location", opts.filename)
		}

		// create the service account
		serviceAccountPayload = &managedservices.ServiceAccountRequest{Name: opts.name, Description: &opts.description}
	}

	a := client.DefaultApi.CreateServiceAccount(context.Background())
	a = a.ServiceAccountRequest(*serviceAccountPayload)
	serviceacct, _, apiErr := a.Execute()

	if apiErr.Error() != "" {
		return fmt.Errorf("Could not create service account: %w", apiErr)
	}

	logger.Infof("Service account '%v' created", serviceacct.GetName())

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

func promptRequestPayload(opts *Options) (payload *managedservices.ServiceAccountRequest, err error) {
	_, err = opts.Connection()
	if err != nil {
		return nil, err
	}

	logger, err := opts.Logger()
	if err != nil {
		return nil, err
	}

	answers := struct {
		Name        string
		Description string
	}{}

	logger.Debug("Beginning interactive prompt")

	qs := []*survey.Question{
		{
			Name:     "name",
			Prompt:   &survey.Input{Message: "Name:", Help: "Give your service account an easily identifiable name"},
			Validate: survey.Required,
		},
		{
			Name:   "description",
			Prompt: &survey.Multiline{Message: "Description"},
		},
	}

	err = survey.Ask(qs, &answers)
	if err = cmdutil.CheckSurveyError(err); err != nil {
		return nil, err
	}

	// if the --output flag was not used, ask in the prompt
	if opts.output == "" {
		logger.Debug("--output flag is not set, prompting user to choose a value")

		outputPrompt := &survey.Select{
			Message: "Credentials output format",
			Help:    "Output format to save the service account credentials",
			Options: flagutil.CredentialsOutputFormats,
			Default: "json",
		}

		err = survey.AskOne(outputPrompt, &opts.output)
		if err = cmdutil.CheckSurveyError(err); err != nil {
			return nil, err
		}
	}

	opts.filename, err = chooseFileLocation(opts)
	if err != nil {
		return nil, err
	}

	serviceacct := &managedservices.ServiceAccountRequest{
		Name:        answers.Name,
		Description: &answers.Description,
	}

	if opts.overwrite {
		return serviceacct, nil
	}

	return serviceacct, err
}

// start an interactive prompt to get the path to the credentials file
// a while loop will be entered as it can take multiple attempts to find a suitable location
// if the file already exists
func chooseFileLocation(opts *Options) (filePath string, err error) {
	chooseFileLocation := true
	filePath = opts.filename

	defaultPath := credentials.AbsolutePath(opts.output, filePath)

	logger, err := opts.Logger()
	if err != nil {
		return "", err
	}

	var attempts = 0
	for chooseFileLocation {
		attempts++
		logger.Debug("Choosing location to save service account credentials")
		logger.Debug("Attempt number", attempts)

		// choose location
		fileNamePrompt := &survey.Input{
			Message: "Credentials file location",
			Help:    "Enter the path to the file where the service account credentials will be saved to",
			Default: defaultPath,
		}
		if filePath == "" {
			err = survey.AskOne(fileNamePrompt, &filePath, survey.WithValidator(survey.Required))
			if err = cmdutil.CheckSurveyError(err); err != nil {
				return "", err
			}
		}

		// check if the file selected already exists
		// if so ask the user to confirm if they would like to have it overwritten
		_, err = os.Stat(filePath)
		// file does not exist, we will create it
		if os.IsNotExist(err) {
			return filePath, nil
		}
		// another error occurred
		if err != nil {
			return "", err
		}

		if opts.overwrite {
			return filePath, nil
		}

		overwriteFilePrompt := &survey.Confirm{
			Message: fmt.Sprintf("The file '%v' already exists. Do you want to overwrite it?", filePath),
		}

		err = survey.AskOne(overwriteFilePrompt, &opts.overwrite)

		if err = cmdutil.CheckSurveyError(err); err != nil {
			return "", err
		}

		if opts.overwrite {
			return filePath, nil
		}

		filePath = ""

		diffLocationPrompt := &survey.Confirm{
			Message: "Would you like to specify a different file location?",
		}
		err = survey.AskOne(diffLocationPrompt, &chooseFileLocation)
		if err = cmdutil.CheckSurveyError(err); err != nil {
			return "", err
		}
		defaultPath = ""
	}

	if filePath == "" {
		return "", fmt.Errorf("You must specify a file to save the service account credentials")
	}

	return "", nil
}
