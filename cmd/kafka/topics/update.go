package topics

import (
	"fmt"

	"github.com/spf13/cobra"
)

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

	cmd.Flags().StringVarP(&topicName, Name, "n", "", "Topic name (required)")
	cmd.MarkFlagRequired(Name)
	cmd.Flags().StringVarP(&config, Config, "c", "", "A comma-separated list of configuration to override e.g 'key1=value1,key2=value2'. (required)")
	cmd.MarkFlagRequired(Config)
	return cmd
}

func updateTopic(cmd *cobra.Command, _ []string) {
	fmt.Println("Updating topic " + topicName + " (" + config + ") ...")
	doRemoteOperation()
	fmt.Println("Topic " + topicName + " updated")
}
