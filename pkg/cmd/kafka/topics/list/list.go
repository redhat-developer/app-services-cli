package list

import (
	"fmt"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	flagutil "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil/flags"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/sdk/kafka/topics"
	"github.com/spf13/cobra"
)

type Options struct {
	Config     config.IConfig
	Connection func() (connection.Connection, error)
	Logger     func() (logging.Logger, error)

	output   string
	insecure bool
}

// NewListTopicCommand gets a new command for getting kafkas.
func NewListTopicCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
	}

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List topics",
		Long:  "List all topics in the current selected Managed Kafka cluster",
		RunE: func(cmd *cobra.Command, _ []string) error {
			logger, err := opts.Logger()
			if err != nil {
				return err
			}
			if flagutil.IsValidInput(opts.output, flagutil.AllowedListFormats...) {
				logger.Infof("Unknown flag value '%v' for --output. Using table format instead", opts.output)
				opts.output = "plain"
			}
			return listTopic(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.output, "output", "o", "plain", fmt.Sprintf("Output format of the results. Choose from %q", flagutil.AllowedListFormats))
	cmd.Flags().BoolVar(&opts.insecure, "insecure", false, "Enables insecure communication with the server. This disables verification of TLS certificates and host names.")
	return cmd
}

func listTopic(opts *Options) error {
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

	err = topics.ValidateCredentials(topicOpts)
	if err != nil {
		return fmt.Errorf("Unable to create credentials: %w", err)
	}
	logger.Info("Topics:")
	err = topics.ListKafkaTopics(topicOpts)
	if err != nil {
		return fmt.Errorf("Failed to perform list operation: %w", err)
	}

	return err
}
