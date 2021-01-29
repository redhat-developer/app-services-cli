package topics

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/topic/create"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/topic/delete"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/topic/list"
)

const (
	Name      = "name"
	Operation = "operation"
)

// NewTopicCommand gives commands that manages Kafka topics.
func NewTopicCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "topic",
		Short: "Create, list and delete Kafka topics",
		Long: heredoc.Doc(`
			Create, list and delete topics for a Kafka instance.
		`),
		Example: heredoc.Doc(`
			# create a topic in the current Kafka instance
			$ rhoas kafka topic create --name "my-example-topic"

			# list all topics for a Kafka instance
			$ rhoas kafka topic list

			# delete a topic for the current Kafka instance
			$ rhoas kafka topic delete --name "my-example-topic"
		`),
	}

	cmd.AddCommand(
		create.NewCreateTopicCommand(f),
		list.NewListTopicCommand(f),
		delete.NewDeleteTopicCommand(f),
	)

	return cmd
}
