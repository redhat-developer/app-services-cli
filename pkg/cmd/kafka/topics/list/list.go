package list

import (
	"fmt"
	"os"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/sdk/kafka/topics"
	"github.com/spf13/cobra"
)

type Options struct {
	Config     config.IConfig
	Connection func() (connection.IConnection, error)

	output   string
	insecure bool
}

// NewListTopicCommand gets a new command for getting kafkas.
func NewListTopicCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
	}

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List topics",
		Long:  "List all topics in the current selected Managed Kafka cluster",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return listTopic(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.output, "Output", "o", "plain-text", "The output format as 'plain-text', 'json', or 'yaml'")
	cmd.Flags().BoolVar(&opts.insecure, "insecure", false, "Enables insecure communication with the server. This disables verification of TLS certificates and host names.")
	return cmd
}

func listTopic(opts *Options) error {
	topicOpts := &topics.Options{
		Connection: opts.Connection,
		Config:     opts.Config,
		Insecure:   opts.insecure,
	}

	err := topics.ValidateCredentials(topicOpts)
	if err != nil {
		return fmt.Errorf("Unable to create credentials: %w", err)
	}
	fmt.Fprintln(os.Stderr, "Topics:")
	err = topics.ListKafkaTopics(topicOpts)
	if err != nil {
		return fmt.Errorf("Failed to perform list operation: %w", err)
	}

	return err
}
