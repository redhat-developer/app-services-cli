package use

import (
	"context"
	"fmt"
	"os"

	"github.com/MakeNowJust/heredoc"

	"github.com/spf13/cobra"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
)

type options struct {
	id string

	Config func() (config.Config, error)
}

func NewUseCommand(f *factory.Factory) *cobra.Command {
	opts := &options{}

	cmd := &cobra.Command{
		Use:   "use",
		Short: "Set the current Kafka instance context",
		Long:  "Sets a Kafka instance in context by its unique identifier",
		Example: heredoc.Doc(`
			$ rhoas kafka use
			$ rhoas kafka use --id=1iSY6RQ3JKI8Q0OTmjQFd3ocFRg`,
		),
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runUse(opts)
		},
	}

	cmd.Flags().StringVar(&opts.id, "id", "", "ID of the Kafka instance you want to use")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func runUse(opts *options) error {
	cfg, err := opts.Config()
	if err != nil {
		return fmt.Errorf("Error loading config: %w", err)
	}

	connection, err := cfg.Connection()
	if err != nil {
		return fmt.Errorf("Can't create connection: %w", err)
	}

	client := connection.NewMASClient()

	res, _, err := client.DefaultApi.GetKafkaById(context.Background(), opts.id)
	if err != nil {
		return fmt.Errorf("Unable to retrieve Kafka instance \"%v\": %w", opts.id, err)
	}

	// build Kafka config object from the response
	var kafkaConfig config.KafkaConfig = config.KafkaConfig{
		ClusterID: res.Id,
	}

	cfg.Services.Kafka = &kafkaConfig
	if err := cfg.Save(); err != nil {
		return fmt.Errorf("Unable to update config: %w", err)
	}

	fmt.Fprintf(os.Stderr, "Using Kafka instance \"%v\"\n", res.Id)

	return nil
}
