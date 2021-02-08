package delete

import (
	"fmt"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/topic/flags"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/sdk/kafka/topics"

	"github.com/spf13/cobra"
)

type Options struct {
	topicName string
	insecure  bool

	Config     config.IConfig
	Connection func() (connection.Connection, error)
	Logger     func() (logging.Logger, error)
}

// NewDeleteTopicCommand gets a new command for deleting kafka topic.
func NewDeleteTopicCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config: f.Config,
		Logger: f.Logger,
	}

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete Kafka topic",
		Long:  "Delete a topic from the current Kafka instance",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return deleteTopic(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.topicName, flags.FlagName, "n", "", "Topic name.")
	cmd.Flags().BoolVar(&opts.insecure, "insecure", false, "Enables insecure communication with the server. This disables verification of TLS certificates and host names.")

	_ = cmd.MarkFlagRequired(flags.FlagName)
	return cmd
}

func deleteTopic(opts *Options) error {
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

	logger.Infof("Deleting topic %v", opts.topicName)
	err = topics.ValidateCredentials(topicOpts)
	if err != nil {
		return fmt.Errorf("Unable to create credentials for topic: %w", err)
	}
	err = topics.DeleteKafkaTopic(opts.topicName, topicOpts)
	if err != nil {
		return fmt.Errorf("Unable to delete topic: %w", err)
	}

	logger.Infof("Topic %v deleted", opts.topicName)

	return nil
}
