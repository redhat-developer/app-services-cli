package credentials

import (
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/sdk/kafka/credentials"
	"github.com/spf13/cobra"
)

var outputFlagValue string
var force bool

// NewGetCommand gets a new command for getting kafkas.
func NewCredentialsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "credentials",
		Short: "Generate credentials to connect to your cluster",
		Long:  "Generate credentials to connect your application to the Kafka cluster",
		Run:   runCredentials,
	}
	cmd.Flags().StringVarP(&outputFlagValue, "output", "o", "env", "Format of the config [env, kafka, properties, json]")
	cmd.Flags().BoolVarP(&force, "force", "f", false, "Force overwrite existing files")
	return cmd
}

func runCredentials(cmd *cobra.Command, _ []string) {
	credentials.RunCredentials(outputFlagValue, force)
}
