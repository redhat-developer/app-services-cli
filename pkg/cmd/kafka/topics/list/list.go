package list

import (
	"fmt"
	"os"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/sdk/kafka/topics"
	"github.com/spf13/cobra"
)

var output string
var insecure bool

type Options struct {
	Config func() (config.Config, error)

	output   string
	insecure bool
}

// NewListTopicCommand gets a new command for getting kafkas.
func NewListTopicCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config: f.Config,
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
	cfg, err := opts.Config()
	if err != nil {
		return fmt.Errorf("Error loading config: %w", err)
	}

	err = topics.ValidateCredentials(&cfg)
	if err != nil {
		return fmt.Errorf("Error creating credentials for list: %w", err)
	}
	fmt.Fprintln(os.Stderr, "Topics:")
	err = topics.ListKafkaTopics(&cfg, insecure)
	if err != nil {
		return fmt.Errorf("Failed to perform list operation: %w", err)
	}

	return err
}
