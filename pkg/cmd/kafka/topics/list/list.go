package list

import (
	"fmt"
	"os"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/sdk/kafka/topics"
	"github.com/spf13/cobra"
)

var output string

const Output = "output"

// NewListTopicCommand gets a new command for getting kafkas.
func NewListTopicCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List topics",
		Long:  "List all topics in the current selected Managed Kafka cluster",
		Run:   listTopic,
	}

	cmd.Flags().StringVarP(&output, Output, "o", "plain-text", "The output format as 'plain-text', 'json', or 'yaml'")
	return cmd
}

func listTopic(cmd *cobra.Command, _ []string) {
	fmt.Fprintln(os.Stderr, "Topics:")

	err := topics.ListKafkaTopics()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to perform list operation")
	}

}
