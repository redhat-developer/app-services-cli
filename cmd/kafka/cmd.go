// Package kafka cluster contains commands for interacting with cluster logic of the service directly instead of through the
// REST API exposed via the serve command.
package kafka

import (
	"github.com/bf2fc6cc711aee1a0c2a/cli/cmd/kafka/topics"
	"github.com/spf13/cobra"
)

func NewKafkaCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kafka",
		Short: "Manage your Kafka clusters",
		Long:  "Manage your Kafka clusters",
	}

	// add sub-commands
	cmd.AddCommand(
		NewCreateCommand(), 
		NewGetCommand(), 
		NewDeleteCommand(),
		NewListCommand(),
		NewUseCommand(),
		NewStatusCommand(),
		topics.NewTopicsCommand(),
		NewCredentialsCommand())

	return cmd
}
