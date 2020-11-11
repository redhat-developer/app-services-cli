package list

import (
	"fmt"
	"os"
	"time"

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
	fmt.Fprintln(os.Stderr, "Listing topics ...")

	// Mimick operation happening by sleeping for a while
	time.Sleep(500 * time.Millisecond)

	fmt.Println(`
3 topics:
topic "topic-1" with 3 partitions:
	partition 0, leader 3, replicas: 1,2,3, isrs: 1,2,3
	partition 1, leader 1, replicas: 1,2,3, isrs: 1,2,3
	partition 2, leader 1, replicas: 1,2, isrs: 1,2
topic "auto_49f744a4327b1b1e" with 2 partitions:
	partition 0, leader 3, replicas: 3, isrs: 3
	partition 1, leader 1, replicas: 1, isrs: 1
topic "auto_e02f58f2c581cba" with 2 partitions:
	partition 0, leader 3, replicas: 3, isrs: 3
	partition 1, leader 1, replicas: 1, isrs: 1`)
}
