package topics

import (
	"github.com/spf13/cobra"

	"github.com/bf2fc6cc711aee1a0c2a/cli/cmd/rhmas/kafka/topics/create"
	"github.com/bf2fc6cc711aee1a0c2a/cli/cmd/rhmas/kafka/topics/delete"
	"github.com/bf2fc6cc711aee1a0c2a/cli/cmd/rhmas/kafka/topics/list"
	"github.com/bf2fc6cc711aee1a0c2a/cli/cmd/rhmas/kafka/topics/update"
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
		Long:  "Manage Kafka topics for the current selected Managed Kafka Cluster",
	}

	cmd.AddCommand(
		create.NewCreateTopicCommand(),
		list.NewListTopicCommand(),
		update.NewUpdateTopicCommand(),
		delete.NewDeleteTopicCommand(),
	)
	return cmd
}
