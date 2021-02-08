package delete

import (
	"context"
	"fmt"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/kas"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/kafka"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/color"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"

	"github.com/MakeNowJust/heredoc"

	"github.com/AlecAivazis/survey/v2"
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/spf13/cobra"
)

type options struct {
	id    string
	force bool

	Config     config.IConfig
	Connection func() (connection.Connection, error)
	Logger     func() (logging.Logger, error)
}

// NewDeleteCommand command for deleting kafkas.
func NewDeleteCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
	}

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a Kafka instance",
		Long: heredoc.Doc(`
			Permanently delete a Kafka instance, including all topics.

			When this command is run, you will be asked to confirm the name of the instance you want to delete.
			Otherwise you can pass "--force" to forcefully delete the instance.
		`),
		Example: heredoc.Doc(`
			# delete the current Kafka instance
			$ rhoas kafka delete

			# delete a Kafka instance with a specific ID
			$ rhoas kafka delete --id=1iSY6RQ3JKI8Q0OTmjQFd3ocFRg
		`),
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			if opts.id != "" {
				return runDelete(opts)
			}

			var kafkaConfig *config.KafkaConfig
			if cfg.Services.Kafka == kafkaConfig || cfg.Services.Kafka.ClusterID == "" {
				return fmt.Errorf("No Kafka instance selected. Use the '--id' flag or set one in context with the 'use' command")
			}

			opts.id = cfg.Services.Kafka.ClusterID

			return runDelete(opts)
		},
	}

	cmd.Flags().StringVar(&opts.id, "id", "", "ID of the Kafka instance you want to delete. If not set, the current Kafka instance will be used.")
	cmd.Flags().BoolVarP(&opts.force, "force", "f", false, "Skip confirmation to forcibly delete this Kafka instance.")

	return cmd
}

func runDelete(opts *options) error {
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

	response, _, apiErr := api.Kafka().GetKafkaById(context.Background(), opts.id).Execute()
	if kas.IsErr(apiErr, kas.ErrorNotFound) {
		return kafka.ErrorNotFound(opts.id)
	}

	if apiErr.Error() != "" {
		return fmt.Errorf("Unable to get Kafka instance: %w", apiErr)
	}

	kafkaName := response.GetName()

	logger.Info("Deleting Kafka instance", color.Info(kafkaName), "\n")

	if !opts.force {
		var promptConfirmName = &survey.Input{
			Message: "Confirm the name of the instance you want to delete:",
		}

		var confirmedKafkaName string
		err = survey.AskOne(promptConfirmName, &confirmedKafkaName)
		if err != nil {
			return err
		}

		if confirmedKafkaName != kafkaName {
			logger.Info("The name you entered does not match the name of the Kafka instance that you are trying to delete. Please check that it correct and try again.")
			return nil
		}
	}

	logger.Debug("Deleting Kafka instance", kafkaName)
	a := api.Kafka().DeleteKafkaById(context.Background(), opts.id)
	a = a.Async(true)
	_, _, apiErr = a.Execute()

	if apiErr.Error() != "" {
		return fmt.Errorf("Unable to delete Kafka instance: %w", apiErr)
	}

	logger.Infof("Kafka instance %v has successfully been deleted", color.Info(kafkaName))

	currentKafka := cfg.Services.Kafka
	// this is not the current cluster, our work here is done
	if currentKafka == nil || currentKafka.ClusterID != response.GetId() {
		return nil
	}

	// the Kafka that was deleted is set as the user's current cluster
	// since it was deleted it should be removed from the config
	cfg.Services.Kafka = nil
	err = opts.Config.Save(cfg)
	if err != nil {
		return err
	}

	return nil
}
