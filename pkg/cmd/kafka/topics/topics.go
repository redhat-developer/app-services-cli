package topics

import (
	"github.com/spf13/cobra"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/topics/create"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/topics/delete"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/topics/list"
)

const (
	Name      = "name"
	Operation = "operation"
)

// NewTopicsCommand gives commands that manages Kafka topics.
func NewTopicsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "topics",
		Short: "Manage topics",
		Long:  "Manage Kafka topics for the current selected Managed Kafka instance",
	}

	cmd.AddCommand(
		create.NewCreateTopicCommand(),
		list.NewListTopicCommand(),
		delete.NewDeleteTopicCommand(),
	)
	return cmd
}
