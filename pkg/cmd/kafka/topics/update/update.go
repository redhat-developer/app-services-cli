package update

import (
	"fmt"
	"os"

	"github.com/segmentio/kafka-go"

	"github.com/spf13/cobra"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/topics/flags"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/sdk/kafka/topics"
)

var topicName string
var config string

const Config = "config"

// NewUpdateTopicCommand gets a new command for updating kafkas topics.
func NewUpdateTopicCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update topic",
		Long:  "Update topic in the current selected Managed Kafka instance",
		Run:   updateTopic,
	}

	cmd.Flags().StringVarP(&topicName, flags.FlagName, "n", "", "Topic name (required)")
	_ = cmd.MarkFlagRequired(flags.FlagName)
	cmd.Flags().StringVarP(&config, Config, "c", "", "A comma-separated list of configuration to override e.g 'key1=value1,key2=value2'. (required)")
	_ = cmd.MarkFlagRequired(Config)
	return cmd
}

func updateTopic(cmd *cobra.Command, _ []string) {
	// TODO not sure about format
	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topicName,
			NumPartitions:     int(1),
			ReplicationFactor: int(1),
		},
	}

	err := topics.CreateKafkaTopic(&topicConfigs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error updating topic: %v\n", topicName)
		return
	}

	fmt.Fprintf(os.Stderr, "Topic %v updated\n", topicName)
}
