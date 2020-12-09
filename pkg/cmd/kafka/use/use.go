package use

import (
	"context"
	"fmt"
	"os"

	"github.com/MakeNowJust/heredoc"

	"github.com/spf13/cobra"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/config"
)

type options struct {
	id string

	cfg *config.Config
}

func NewUseCommand() *cobra.Command {
	opts := &options{}

	cmd := &cobra.Command{
		Use:   "use",
		Short: "Set the current Kafka cluster context",
		Long:  "Sets a Kafka cluster in context by its unique identifier",
		Example: heredoc.Doc(`
			$ rhoas kafka use
			$ rhoas kafka use --id=1iSY6RQ3JKI8Q0OTmjQFd3ocFRg`,
		),
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			cfg, err := config.Load()
			if err != nil {
				return fmt.Errorf("Error loading config: %w", err)
			}
			opts.cfg = cfg

			return runUse(opts)
		},
	}

	cmd.Flags().StringVar(&opts.id, "id", "", "ID of the Kafka cluster you want to use")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func runUse(opts *options) error {
	cfg := *opts.cfg
	connection, err := cfg.Connection()
	if err != nil {
		return fmt.Errorf("Can't create connection: %w", err)
	}

	client := connection.NewMASClient()

	res, _, err := client.DefaultApi.GetKafkaById(context.Background(), opts.id)
	if err != nil {
		return fmt.Errorf("Unable to retrieve Kafka cluster \"%v\": %w", opts.id, err)
	}

	// build Kafka config object from the response
	var kafkaConfig config.KafkaConfig = config.KafkaConfig{
		ClusterID: res.Id,
	}

	cfg.Services.SetKafka(&kafkaConfig)
	if err := config.Save(&cfg); err != nil {
		return fmt.Errorf("Unable to update config: %w", err)
	}

	fmt.Fprintf(os.Stderr, "Using Kafka cluster \"%v\"", res.Id)

	return nil
}
