package topics

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/topic/create"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/topic/delete"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/topic/describe"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/topic/list"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/topic/update"
)

const (
	Name      = "name"
	Operation = "operation"
)

// NewTopicCommand gives commands that manages Kafka topics.
func NewTopicCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "topic",
		Short: "Create, describe, update, list and delete Kafka topics",
		Long: heredoc.Doc(`
			Create, describe, update, list and delete topics for a Kafka instance.
		`),
	}

	cmd.AddCommand(
		create.NewCreateTopicCommand(f),
		list.NewListTopicCommand(f),
		delete.NewDeleteTopicCommand(f),
		describe.NewDescribeTopicCommand(f),
		update.NewUpdateTopicCommand(f),
	)

	return cmd
}
