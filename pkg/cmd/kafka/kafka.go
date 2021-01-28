// Package kafka instance contains commands for interacting with cluster logic of the service directly instead of through the
// REST API exposed via the serve command.
package kafka

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/create"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/delete"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/describe"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/list"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/topics"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/use"
)

func NewKafkaCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kafka",
		Short: "Create, view, use and manage your Kafka instances",
		Long: heredoc.Doc(`
			Perform various operations on your Kafka instances.
		`),
		Example: heredoc.Doc(`
			# create a Kafka instance
			$ rhoas kafka create

			# list Kafka instances
			$ rhoas kafka list

			# create a Kafka topic
			$ rhoas kafka topics create --name "my-kafka-topic"
		`),
	}

	// add sub-commands
	cmd.AddCommand(
		create.NewCreateCommand(f),
		describe.NewDescribeCommand(f),
		delete.NewDeleteCommand(f),
		list.NewListCommand(f),
		use.NewUseCommand(f),
		topics.NewTopicsCommand(f),
	)
	return cmd
}
