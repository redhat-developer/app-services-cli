package delete

import (
	"github.com/bf2fc6cc711aee1a0c2a/cli/cmd/rhmas/kafka/flags"
	"time"
	"os"
	"fmt"

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

	cmd.Flags().StringVarP(&topicName, flags.FlagName, "n", "", "Topic name (required)")
	cmd.MarkFlagRequired(flags.FlagName)
	return cmd
}

func deleteTopic(cmd *cobra.Command, _ []string) {
	fmt.Fprintln(os.Stderr, "Deleting topic " + topicName + " ...")
	// Mimick operation happening by sleeping for a while
	time.Sleep(500 * time.Millisecond)
	fmt.Fprintln(os.Stderr, "Topic " + topicName + " deleted")
}
