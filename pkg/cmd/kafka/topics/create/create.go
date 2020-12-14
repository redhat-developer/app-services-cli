package create

import (
	"fmt"
	"os"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	topicflags "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/topics/flags"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/sdk/kafka/topics"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/cobra"
)

var partitions, replicas int32

const (
	Partitions = "partitions"
	Replicas   = "replicas"
)

type Options struct {
	topicName  string
	insecure   bool
	partitions int32
	replicas   int32

	Config func() (config.Config, error)
}

var topicName string
var insecure bool

// NewCreateTopicCommand gets a new command for creating kafka topic.
func NewCreateTopicCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config: f.Config,
	}

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create topic",
		Long:  "Create topic in the current selected Managed Kafka cluster",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return createTopic(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.topicName, topicflags.FlagName, "n", "", "Topic name (required)")
	_ = cmd.MarkFlagRequired(topicflags.FlagName)
	cmd.Flags().Int32VarP(&opts.partitions, Partitions, "p", 1, "Set number of partitions")
	cmd.Flags().Int32VarP(&opts.replicas, Replicas, "r", 1, "Set number of replicas")
	cmd.Flags().BoolVar(&opts.insecure, "insecure", false, "Enables insecure communication with the server. This disables verification of TLS certificates and host names.")

	// TODO define file format etc
	return cmd
}

func createTopic(opts *Options) error {
	cfg, err := opts.Config()
	if err != nil {
		return fmt.Errorf("Error loading config: %w", err)
	}

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             opts.topicName,
			NumPartitions:     int(opts.partitions),
			ReplicationFactor: int(opts.replicas),
		},
	}
	err = topics.ValidateCredentials(&cfg)
	if err != nil {
		return fmt.Errorf("Error creating credentials for topic: %w", err)
	}
	err = topics.CreateKafkaTopic(topicConfigs, &cfg, insecure)
	if err != nil {
		return fmt.Errorf("Error creating topic: %w", err)
	}

	fmt.Fprintf(os.Stderr, "Topic %v created\n", err)

	return nil
}
