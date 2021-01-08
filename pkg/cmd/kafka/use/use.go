package use

import (
	"context"
	"fmt"

	"github.com/MakeNowJust/heredoc"

	"github.com/spf13/cobra"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"
)

type options struct {
	id string

	Config     config.IConfig
	Connection func() (connection.Connection, error)
	Logger     func() (logging.Logger, error)
}

func NewUseCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
	}

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
	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	cfg, err := opts.Config.Load()
	if err != nil {
		return err
	}

	connection, err := opts.Connection()
	if err != nil {
		return err
	}

	client := connection.NewAPIClient()

	res, _, apiErr := client.DefaultApi.GetKafkaById(context.Background(), opts.id).Execute()
	if apiErr.Error() != "" {
		return fmt.Errorf("Unable to retrieve Kafka instance \"%v\": %w", opts.id, apiErr)
	}

	// build Kafka config object from the response
	var kafkaConfig config.KafkaConfig = config.KafkaConfig{
		ClusterID: *res.Id,
	}

	cfg.Services.Kafka = &kafkaConfig
	if err := opts.Config.Save(cfg); err != nil {
		return fmt.Errorf("Unable to use Kafka instance: %w", err)
	}

	logger.Infof("Using Kafka instance \"%v\"\n", *res.Id)

	return nil
}
