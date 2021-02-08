package create

import (
	"fmt"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	topicflags "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/topic/flags"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/sdk/kafka/topics"

	"github.com/segmentio/kafka-go"
	"github.com/spf13/cobra"
)

const (
	Partitions = "partitions"
	Replicas   = "replicas"
)

type Options struct {
	topicName  string
	insecure   bool
	partitions int32
	replicas   int32

	Config     config.IConfig
	Connection func() (connection.Connection, error)
	Logger     func() (logging.Logger, error)
}

// NewCreateTopicCommand gets a new command for creating kafka topic.
func NewCreateTopicCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Connection: f.Connection,
		Config:     f.Config,
		Logger:     f.Logger,
	}

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a Kafka topic",
		Long:  "Create topic in the current Kafka instance",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return createTopic(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.topicName, topicflags.FlagName, "n", "", "Topic name.")
	cmd.Flags().Int32VarP(&opts.partitions, Partitions, "p", 1, "Set number of partitions.")
	cmd.Flags().Int32VarP(&opts.replicas, Replicas, "r", 1, "Set number of replicas.")
	cmd.Flags().BoolVar(&opts.insecure, "insecure", false, "Enables insecure communication with the server. This disables verification of TLS certificates and host names.")

	_ = cmd.MarkFlagRequired(topicflags.FlagName)

	return cmd
}

func createTopic(opts *Options) error {
	_, err := opts.Connection()
	if err != nil {
		return err
	}

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	topicOpts := &topics.Options{
		Connection: opts.Connection,
		Config:     opts.Config,
		Insecure:   opts.insecure,
		Logger:     opts.Logger,
	}

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             opts.topicName,
			NumPartitions:     int(opts.partitions),
			ReplicationFactor: int(opts.replicas),
		},
	}
	err = topics.ValidateCredentials(topicOpts)
	if err != nil {
		return fmt.Errorf("Unable to create credentials for topic: %w", err)
	}
	err = topics.CreateKafkaTopic(topicConfigs, topicOpts)
	if err != nil {
		return fmt.Errorf("Unable to create topic: %w", err)
	}

	logger.Infof("Topic %v created", opts.topicName)

	return nil
}
