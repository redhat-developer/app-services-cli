package topics

import (
	"time"

	"github.com/spf13/cobra"
)

const (
	Name      = "name"
	Operation = "operation"
)

var topicName string

// NewTopicsCommand gives commands that manages Kafka topics.
func NewTopicsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "topics",
		Short: "Manage topics",
		Long:  "Manage Kafka topics for the current selected Managed Kafka Cluster",
	}

	cmd.AddCommand(NewCreateTopicCommand(), NewListTopicCommand(), NewUpdateTopicCommand(), NewDeleteTopicCommand())
	return cmd
}

func doRemoteOperation() {
	// Mimick operation happening by sleeping for a while
	time.Sleep(500 * time.Millisecond)
}
