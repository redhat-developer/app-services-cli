package status

import (
	"context"
	"fmt"
	"os"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/dump"

	"github.com/spf13/cobra"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
)

type Options struct {
	id string

	Config     config.IConfig
	Connection func() (connection.IConnection, error)
}

func NewStatusCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
	}

	cmd := &cobra.Command{
		Use:   "status",
		Short: "Get status of current Kafka instance",
		Long:  "Gets the status of the current Kafka instance context",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			cfg, err := opts.Config.Load()
			if err != nil {
				return err
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
	connection, err := opts.Connection()
	if err != nil {
		return err
	}

	client := connection.NewMASClient()

	res, _, apiErr := client.DefaultApi.GetKafkaById(context.Background(), opts.id).Execute()
	if apiErr.Error() != "" {
		return fmt.Errorf("Unable to get Kafka instance: %w", apiErr)
	}

	type kafkaStatus struct {
		ID     string `header:"ID"`
		Name   string `header:"Name"`
		Status string `header:"Status"`
	}

	statusInfo := &kafkaStatus{
		ID:     *res.Id,
		Name:   *res.Name,
		Status: *res.Status,
	}

	dump.Table(os.Stdout, statusInfo)

	return nil
}
