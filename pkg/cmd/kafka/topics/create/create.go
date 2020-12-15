package create

import (
	"fmt"
	"os"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/topics/flags"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/sdk/kafka/topics"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/cobra"
)

var partitions, replicas int32

const (
	Partitions = "partitions"
	Replicas   = "replicas"
)

var topicName string

// NewCreateTopicCommand gets a new command for creating kafka topic.
func NewCreateTopicCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create topic",
		Long:  "Create topic in the current selected Managed Kafka instance",
		Run:   createTopic,
	}

	cmd.Flags().StringVarP(&topicName, flags.FlagName, "n", "", "Topic name (required)")
	_ = cmd.MarkFlagRequired(flags.FlagName)
	cmd.Flags().Int32VarP(&partitions, Partitions, "p", 1, "Set number of partitions")
	cmd.Flags().Int32VarP(&replicas, Replicas, "r", 1, "Set number of replicas")

	// TODO define file format etc
	return cmd
}

func createTopic(cmd *cobra.Command, _ []string) {
	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topicName,
			NumPartitions:     int(partitions),
			ReplicationFactor: int(replicas),
		},
	}
	err := topics.ValidateCredentials()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating credentials for topic: %v\n", err)
		return
	}
	err = topics.CreateKafkaTopic(topicConfigs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating topic: %v\n", err)
		return
	}

	fmt.Fprintf(os.Stderr, "Topic %v created\n", err)
}
