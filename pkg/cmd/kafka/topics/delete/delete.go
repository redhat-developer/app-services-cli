package delete

import (
	"fmt"
	"os"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/topics/flags"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/sdk/kafka/topics"

	"github.com/spf13/cobra"
)

var topicName string

// NewDeleteTopicCommand gets a new command for deleting kafka topic.
func NewDeleteTopicCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete topic",
		Long:  "Delete topic from the current selected Managed Kafka instance",
		Run:   deleteTopic,
	}

	cmd.Flags().StringVarP(&topicName, flags.FlagName, "n", "", "Topic name (required)")
	_ = cmd.MarkFlagRequired(flags.FlagName)
	return cmd
}

func deleteTopic(cmd *cobra.Command, _ []string) {
	fmt.Fprintf(os.Stderr, "Deleting topic %v\n", topicName)
	err := topics.DeleteKafkaTopic(topicName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error deleting topic: %v\n", topicName)
		return
	}

	fmt.Fprintf(os.Stderr, "Topic %v deleted\n", topicName)
}
