package status

import (
	"context"
	"fmt"
	"os"

	"github.com/MakeNowJust/heredoc"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/dump"
	pkgKafka "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/kafka"

	"github.com/spf13/cobra"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
)

type Options struct {
	id string

	Config     config.IConfig
	Connection func() (connection.Connection, error)
}

func NewStatusCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
	}

	cmd := &cobra.Command{
		Use:   "status",
		Short: "View status of a Kafka instance",
		Long: heredoc.Doc(`
			View the status of a Kafka instance.

			The values shown as part of the status are: ID, Name, Status, Bootstrap Server Host
		`),
		Example: heredoc.Doc(`
			# view the status of the current Kafka instance
			$ rhoas kafka status

			# view the status of a Kafka instance using its ID
			$ rhoas kafka status --id "1nYlgkt87xelHT1wnOdEyGhFhaO"
		`),
		Args: cobra.ExactArgs(0),
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

	api := connection.API()

	res, _, apiErr := api.Kafka.GetKafkaById(context.Background(), opts.id).Execute()
	pkgKafka.TransformKafkaRequest(&res)

	if apiErr.Error() != "" {
		return fmt.Errorf("Unable to get Kafka instance: %w", apiErr)
	}

	type kafkaStatus struct {
		ID            string `header:"ID"`
		Name          string `header:"Name"`
		Status        string `header:"Status"`
		BootstrapHost string `header:"Bootstrap Server Host"`
	}

	statusInfo := &kafkaStatus{
		ID:            res.GetId(),
		Name:          res.GetName(),
		Status:        res.GetStatus(),
		BootstrapHost: res.GetBootstrapServerHost(),
	}

	dump.Table(os.Stdout, statusInfo)

	return nil
}
