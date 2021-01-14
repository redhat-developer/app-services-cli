package delete

import (
	"context"
	"fmt"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"

	"github.com/MakeNowJust/heredoc"

	"github.com/AlecAivazis/survey/v2"
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/spf13/cobra"
)

type options struct {
	id string

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
		Long:  "Permanently deletes a Kafka instance",
		Example: heredoc.Doc(`
			$ rhoas kafka delete
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

	cmd.Flags().StringVar(&opts.id, "id", "", "ID of the Kafka instance you want to delete. If not set, the currently selected Kafka instance will be used")

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

	client := connection.NewAPIClient()

	response, _, apiErr := client.DefaultApi.GetKafkaById(context.Background(), opts.id).Execute()

	if apiErr.Error() != "" {
		return fmt.Errorf("Unable to get Kafka instance: %w", apiErr)
	}

	kafkaName := response.GetName()

	var confirmDeleteAction bool
	var promptConfirmAction = &survey.Confirm{
		Message: "Once a Kafka instance is deleted it cannot be recovered, are you sure you want to proceed?",
	}

	err = survey.AskOne(promptConfirmAction, &confirmDeleteAction)
	if err != nil {
		return err
	}
	if !confirmDeleteAction {
		return nil
	}

	var promptConfirmName = &survey.Input{
		Message: "Please confirm the name of the Kafka instance you wish to permanently delete:",
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

	logger.Debug("Deleting Kafka instance", kafkaName)
	_, _, apiErr = client.DefaultApi.DeleteKafkaById(context.Background(), opts.id).Execute()

	if apiErr.Error() != "" {
		return fmt.Errorf("Unable to delete Kafka instance: %w", apiErr)
	}

	logger.Info("Kafka instance has successfully been deleted")

	currentKafka := cfg.Services.Kafka
	// this is not the current cluster, our work here is done
	if currentKafka.ClusterID != kafkaName {
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
