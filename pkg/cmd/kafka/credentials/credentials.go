package credentials

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/sdk/kafka/credentials"
	"github.com/spf13/cobra"
)

var outputFlagValue string
var force bool

// NewCredentialsCommand gets a new command for
// generating credentials to connect to your Kafka cluster
func NewCredentialsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "credentials",
		Short: "Generate credentials to connect to your cluster",
		Long: heredoc.Doc(`
			Generate SASL/PLAIN credentials to connect your application to the Kafka cluster.
			The credentials are saved to an output format of your choice from the list below with the --output flag:
				
				- env (default): Saves the credentials to a .env file as environment variables
				- kafka, properties: Saves the cluster credentials to a kafka.properties file
				- json: Saves the cluster credentials in a credentials.json JSON file
		`),
		Example: heredoc.Doc(`
			$ rhoas kafka credentials
			$ rhoas kafka credentials --force
			$ rhoas kafka credentials --output kafka
			$ rhoas kafka credentials --output=env
			$ rhoas kafka credentials -o=properties
			$ rhoas kafka credentials -o json`,
		),
		Run: runCredentials,
	}
	cmd.Flags().StringVarP(&outputFlagValue, "output", "o", "env", "Format of the config [env, kafka, properties, json]")
	cmd.Flags().BoolVarP(&force, "force", "f", false, "Force overwrite a file if it already exists")
	return cmd
}

func runCredentials(cmd *cobra.Command, _ []string) {
	credentials.RunCredentials(outputFlagValue, force)
}
