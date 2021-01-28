package use

import (
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/kafka"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/kas"
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
		Short: "Set the current Kafka instance",
		Long: heredoc.Doc(`
			Select a Kafka instance and set it in the config as the current Kafka instance.

			Once a Kafka instance is used, it is saved as the current instance.
			When an ID is not specified in other Kafka commands, the current Kafka instance is used.
		`),
		Example: heredoc.Doc(`
			# use the Kafka instance that has an ID of "1iSY6RQ3JKI8Q0OTmjQFd3ocFRg"
			$ rhoas kafka use --id=1iSY6RQ3JKI8Q0OTmjQFd3ocFRg`,
		),
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runUse(opts)
		},
	}

	cmd.Flags().StringVar(&opts.id, "id", "", "ID of the Kafka instance to use")
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

	api := connection.API()

	res, _, apiErr := api.Kafka.GetKafkaById(context.Background(), opts.id).Execute()
	if kas.IsErr(apiErr, kas.ErrorNotFound) {
		return kafka.ErrorNotFound(opts.id)
	}
	
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
