package list

import (
	"fmt"
	"os"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/sdk/kafka/topics"
	"github.com/spf13/cobra"
)

var output string
var insecure bool

const Output = "output"

// NewListTopicCommand gets a new command for getting kafkas.
func NewListTopicCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List topics",
		Long:  "List all topics in the current selected Managed Kafka cluster",
		Run: func(cmd *cobra.Command, _ []string) {
			listTopic(insecure)
		},
	}

	cmd.Flags().StringVarP(&output, Output, "o", "plain-text", "The output format as 'plain-text', 'json', or 'yaml'")
	cmd.Flags().BoolVar(&insecure, "insecure", false, "Enables insecure communication with the server. This disables verification of TLS certificates and host names.")
	return cmd
}

func listTopic(insecure bool) {

	err := topics.ValidateCredentials()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating credentials for list: %v\n", err)
		return
	}
	fmt.Fprintln(os.Stderr, "Topics:")
	err = topics.ListKafkaTopics(insecure)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to perform list operation: %v\n", err)
	}
}
