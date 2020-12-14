package create

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/MakeNowJust/heredoc"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/managedservices"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/spf13/cobra"
)

// Templates
var (
	templateProperties = heredoc.Doc(`
	## Generated by rhoas cli
	user=%v
	password=%v
	`)

	templateEnv = heredoc.Doc(`
	## Generated by rhoas cli
	USER=%v
	PASSWORD=%v
	`)

	templateJSON = heredoc.Doc(`
	{ 
		"user":"%v", 
		"password":"%v" 
	}`)

	templateSecret = heredoc.Doc(`
	kind: Secret
	apiVersion: v1
	metadata:
	  name: "rhoas-service-account-secret"
	stringData:
	  clientID: "%v"
	  clientSecret: "%v"
	type: Opaque
	`)
)

type Options struct {
	Config func() (config.Config, error)

	output      string
	force       bool
	name        string
	description string
	scopes      []string
	filename    string
}

// NewCreateCommand creates a new command to create service accounts
func NewCreateCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config: f.Config,
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
			if opts.output != "env" &&
				opts.output != "properties" &&
				opts.output != "json" &&
				opts.output != "kafka" &&
				opts.output != "kube" {
				return fmt.Errorf("Invalid output format '%v'", opts.output)
			}

			return runCreate(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.output, "output", "o", "env", "Format of the config [env, kafka, properties, json, kube]")
	cmd.Flags().StringVar(&opts.name, "name", "", "Name of the service account")
	cmd.Flags().StringVar(&opts.description, "description", "", "Description for the service account")
	cmd.Flags().StringArrayVar(&opts.scopes, "scopes", []string{"kafka-all"},
		"Number of supported scopes (permissions) for this service account")
	cmd.Flags().BoolVarP(&opts.force, "force", "f", false, "Force overwrite a file if it already exists")
	cmd.Flags().StringVar(&opts.filename, "output-file", "", "Sets a custom file location to save the credentials")

	_ = cmd.MarkFlagRequired("output")
	_ = cmd.MarkFlagRequired("name")

	return cmd
}

func runCreate(opts *Options) error {
	cfg, err := opts.Config()
	if err != nil {
		return fmt.Errorf("Could not load config: %w", err)
	}
	
	connection, err := cfg.Connection()
	if err != nil {
		return fmt.Errorf("Can't create connection: %w", err)
	}

	fmt.Fprintf(os.Stderr, "Creating service account with the following permissions: %v\n", opts.scopes)

	client := connection.NewMASClient()

	var fileFormat string
	var fileName string
	switch opts.output {
	case "env":
		fileFormat = templateEnv
		fileName = ".env"
	case "properties":
		fileFormat = templateProperties
		fileName = "kafka.properties"
	case "json":
		fileFormat = templateJSON
		fileName = "credentials.json"
	case "kube":
		fileFormat = templateSecret
		fileName = "credentials.yaml"
	}

	if opts.filename != "" {
		fileName = path.Clean(opts.filename)
	}

	fileDir := path.Dir(fileName)
	if !pathExists(fileDir) {
		return fmt.Errorf("Directory '%v' does not exist", fileDir)
	}

	svcAcctPayload := &managedservices.ServiceAccountRequest{Name: opts.name, Description: opts.description}
	response, _, err := client.DefaultApi.CreateServiceAccount(context.Background(), *svcAcctPayload)

	if err != nil {
		return fmt.Errorf("Could not create service account: %w", err)
	}

	fmt.Fprintf(os.Stderr, "Writing credentials to %v\n", fileName)
	fileContent := fmt.Sprintf(fileFormat, response.ClientID, response.ClientSecret)

	dataToWrite := []byte(fileContent)

	if pathExists(fileName) && !opts.force {
		fmt.Fprintf(os.Stderr, "File '%v' already exist. Use --force flag to overwrite the file, or use the --output-file flag to choose a custom location\n", fileName)
		return nil
	}

	err = ioutil.WriteFile(fileName, dataToWrite, 0600)
	if err != nil {
		return fmt.Errorf("Could not save file: %w", err)
	}

	fmt.Fprintf(os.Stderr, "Successfully saved credentials to %v\n", fileName)

	return nil
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
