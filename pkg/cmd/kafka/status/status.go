package status

import (
	"context"
	"fmt"
	"os"

	"github.com/landoop/tableprinter"

	"github.com/spf13/cobra"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
)

type Options struct {
	id string

	Config func() (config.Config, error)
}

func NewStatusCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config: f.Config,
	}

	cmd := &cobra.Command{
		Use:   "status",
		Short: "Get status of current Kafka instance",
		Long:  "Gets the status of the current Kafka instance context",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			cfg, err := opts.Config()
			if err != nil {
				return fmt.Errorf("Error loading config: %w", err)
			}

			if opts.id != "" {
				return runStatus(opts)
			}

			var kafkaConfig *config.KafkaConfig
			if cfg.Services.Kafka == kafkaConfig || cfg.Services.Kafka.ClusterID == "" {
				return fmt.Errorf("No Kafka instance selected. Use the '--id' flag or set one in context with the 'use' command")
			}

			opts.id = cfg.Services.Kafka.ClusterID

			return runStatus(opts)
		},
	}

	cmd.Flags().StringVar(&opts.id, "id", "", "ID of the Kafka instance you want to get the status from")

	return cmd
}

func runStatus(opts *Options) error {
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
		return fmt.Errorf("Error retrieving Kafka instance: %w", err)
	}

	type kafkaStatus struct {
		ID     string `header:"ID"`
		Name   string `header:"Name"`
		Status string `header:"Status"`
	}

	statusInfo := &kafkaStatus{
		ID:     res.Id,
		Name:   res.Name,
		Status: res.Status,
	}

	printer := tableprinter.New(os.Stdout)
	printer.Print(statusInfo)

	return nil
}
