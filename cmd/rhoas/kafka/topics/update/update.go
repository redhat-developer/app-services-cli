package update

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/bf2fc6cc711aee1a0c2a/cli/cmd/rhoas/kafka/flags"
)

var topicName string
var config string

const Config = "config"

// NewUpdateTopicCommand gets a new command for updating kafkas topics.
func NewUpdateTopicCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update topic",
		Long:  "Update topic in the current selected Managed Kafka cluster",
		Run:   updateTopic,
	}

	cmd.Flags().StringVarP(&topicName, flags.FlagName, "n", "", "Topic name (required)")
	_ = cmd.MarkFlagRequired(flags.FlagName)
	cmd.Flags().StringVarP(&config, Config, "c", "", "A comma-separated list of configuration to override e.g 'key1=value1,key2=value2'. (required)")
	_ = cmd.MarkFlagRequired(Config)
	return cmd
}

func updateTopic(cmd *cobra.Command, _ []string) {
	fmt.Fprintln(os.Stderr, "Updating topic "+topicName+" ("+config+") ...")
	// Mimick operation happening by sleeping for a while
	time.Sleep(500 * time.Millisecond)
	fmt.Fprintln(os.Stderr, "Topic "+topicName+" updated")
}
