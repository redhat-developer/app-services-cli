package topics

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewDeleteTopicCommand gets a new command for deleting kafka topic.
func NewDeleteTopicCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete topic",
		Long:  "Delete topic from the current selected Managed Kafka cluster",
		Run:   deleteTopic,
	}

	cmd.Flags().StringVarP(&topicName, Name, "n", "", "Topic name (required)")
	cmd.MarkFlagRequired(Name)
	return cmd
}

func deleteTopic(cmd *cobra.Command, _ []string) {
	fmt.Println("Deleting topic " + topicName + " ...")
	doRemoteOperation()
	fmt.Println("Topic " + topicName + " deleted")
}
