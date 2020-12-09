package create

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/MakeNowJust/heredoc"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/managedservices"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/config"
	"github.com/spf13/cobra"
)

// Templates
var (
	templateProperties = heredoc.Doc(`
	## Generated by rhoas cli
	kafka_user=%v
	kafka_password=%v
	`)

	templateEnv = heredoc.Doc(`
	## Generated by rhoas cli
	KAFKA_USER=%v
	KAFKA_PASSWORD=%v
	`)

	templateKafkaPlain = heredoc.Doc(`
	## Generated by rhoas cli
	kafka.sasl.jaas.config=org.apache.kafka.common.security.plain.PlainLoginModule required username="%v" password="%v";
	kafka.sasl.mechanism=PLAIN
	kafka.security.protocol=SASL_SSL
	kafka.ssl.protocol=TLSv1.2
	`)

	templateJSON = heredoc.Doc(`
	{ 
		"name":"%v",
		"user":"%v", 
		"password":"%v" 
	}`)

	templateSecret = heredoc.Doc(`
	kind: Secret
	apiVersion: v1
	metadata:
	  name: "%v"
	stringData:
	  clientID: "%v"
	  clientSecret: "%v"
	type: Opaque
	`)
)

type options struct {
	output      string
	force       bool
	name        string
	description string
	filename    string

	cfg *config.Config
}

// NewCreateCommand creates a new command to create service accounts
func NewCreateCommand() *cobra.Command {
	opts := &options{}

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
			$ rhoas service-account create --output kafka
			$ rhoas service-account create --output=env
			$ rhoas service-account create -o=properties
			$ rhoas service-account create -o env --force
			$ rhoas service-account create -o json`,
		),
		RunE: func(cmd *cobra.Command, _ []string) error {
			if opts.output != "env" && opts.output != "properties" &&
				opts.output != "json" && opts.output != "kafka" {
				return fmt.Errorf("Invalid output format '%v'", opts.output)
			}

			cfg, err := config.Load()
			if err != nil {
				return fmt.Errorf("Error loading config: %w", err)
			}
			opts.cfg = cfg

			return runCreate(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.output, "output", "o", "env", "Format of the config [env, kafka, properties, json, kube]")
	cmd.Flags().StringVar(&opts.name, "name", "", "Name of the service account")
	cmd.Flags().StringVar(&opts.description, "description", "", "Description for the service account")
	cmd.Flags().BoolVarP(&opts.force, "force", "f", false, "Force overwrite a file if it already exists")
	cmd.Flags().StringVar(&opts.filename, "output-file", "", "Sets a custom file location to save the credentials")

	_ = cmd.MarkFlagRequired("output")
	_ = cmd.MarkFlagRequired("name")

	return cmd
}

func runCreate(opts *options) error {
	cfg := opts.cfg
	connection, err := cfg.Connection()
	if err != nil {
		return fmt.Errorf("Can't create connection: %w", err)
	}

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
	case "kafka":
		fileFormat = templateKafkaPlain
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

	fmt.Fprintf(os.Stderr, "Writing credentials to %v \n", fileName)
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

	fmt.Fprintf(os.Stderr, "Successfully saved credentials to %v \n", fileName)

	return nil
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
