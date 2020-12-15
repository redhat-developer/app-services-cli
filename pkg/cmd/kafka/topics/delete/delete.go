package delete

import (
	"fmt"
	"os"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/topics/flags"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/sdk/kafka/topics"

	"github.com/spf13/cobra"
)

var topicName string
var insecure bool

// NewDeleteTopicCommand gets a new command for deleting kafka topic.
func NewDeleteTopicCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete topic",
		Long:  "Delete topic from the current selected Managed Kafka cluster",
		Run: func(cmd *cobra.Command, _ []string) {
			deleteTopic(insecure)
		},
	}

	cmd.Flags().StringVarP(&topicName, flags.FlagName, "n", "", "Topic name (required)")
	cmd.Flags().BoolVar(&insecure, "insecure", false, "Enables insecure communication with the server. This disables verification of TLS certificates and host names.")

	_ = cmd.MarkFlagRequired(flags.FlagName)
	return cmd
}

func deleteTopic(insecure bool) {
	fmt.Fprintf(os.Stderr, "Deleting topic %v\n", topicName)
	err := topics.ValidateCredentials()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating credentials for topic: %v\n", topicName)
		return
	}
	err = topics.DeleteKafkaTopic(topicName, insecure)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error deleting topic: %v\n", err)
		return
	}

	fmt.Fprintf(os.Stderr, "Topic %v deleted\n", topicName)
}
