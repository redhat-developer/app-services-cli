package delete

import (
	"fmt"
	"os"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/topics/flags"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/sdk/kafka/topics"

	"github.com/spf13/cobra"
)

var topicName string
var insecure bool

type Options struct {
	topicName string
	insecure  bool

	Config     config.IConfig
	Connection func() (connection.IConnection, error)
}

// NewDeleteTopicCommand gets a new command for deleting kafka topic.
func NewDeleteTopicCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config: f.Config,
	}

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete topic",
		Long:  "Delete topic from the current selected Managed Kafka cluster",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return deleteTopic(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.topicName, flags.FlagName, "n", "", "Topic name (required)")
	cmd.Flags().BoolVar(&opts.insecure, "insecure", false, "Enables insecure communication with the server. This disables verification of TLS certificates and host names.")

	_ = cmd.MarkFlagRequired(flags.FlagName)
	return cmd
}

func deleteTopic(opts *Options) error {
	topicOpts := &topics.Options{
		Connection: opts.Connection,
		Config:     opts.Config,
		Insecure:   opts.insecure,
	}

	fmt.Fprintf(os.Stderr, "Deleting topic %v\n", topicName)
	err := topics.ValidateCredentials(topicOpts)
	if err != nil {
		return fmt.Errorf("Error creating credentials for topic: %w", err)
	}
	err = topics.DeleteKafkaTopic(opts.topicName, topicOpts)
	if err != nil {
		return fmt.Errorf("Error deleting topic: %w", err)
	}

	fmt.Fprintf(os.Stderr, "Topic %v deleted\n", topicName)

	return nil
}
