package delete

import (
	"fmt"
	"os"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/broker"

	"github.com/spf13/cobra"
)

var topicName string

// NewDeleteTopicCommand gets a new command for deleting kafka topic.
func NewDeleteTopicCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete topic",
		Long:  "Delete topic from the current selected Managed Kafka cluster",
		Run:   deleteTopic,
	}

	cmd.Flags().StringVarP(&topicName, "id", "n", "", "Topic name (required)")
	_ = cmd.MarkFlagRequired("id")
	return cmd
}

func deleteTopic(cmd *cobra.Command, _ []string) {
	fmt.Fprintf(os.Stderr, "Deleting topic %v ", topicName)
	err := broker.DeleteKafkaTopic(topicName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error deleting topic: %v", topicName)
		return
	}

	fmt.Fprintf(os.Stderr, "Topic %v deleted", topicName)
}
