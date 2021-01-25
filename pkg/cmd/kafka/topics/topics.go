package topics

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/topics/create"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/topics/delete"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/topics/list"
)

const (
	Name      = "name"
	Operation = "operation"
)

// NewTopicsCommand gives commands that manages Kafka topics.
func NewTopicsCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "topics",
		Short: "Create, list and delete Kafka topics",
		Long: heredoc.Doc(`
			Create, list and delete topics for a Kafka instance.
		`),
		Example: heredoc.Doc(`
			# create a topic in the current Kafka instance
			$ rhoas kafka topics create --name "my-example-topic"

			# list all topics for a Kafka instance
			$ rhoas kafka topics list

			# delete a topic for the current Kafka instance
			$ rhoas kafka topics delete --name "my-example-topic"
		`),
	}

	cmd.AddCommand(
		create.NewCreateTopicCommand(f),
		list.NewListTopicCommand(f),
		delete.NewDeleteTopicCommand(f),
	)

	return cmd
}
