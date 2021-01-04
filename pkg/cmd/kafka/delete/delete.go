package delete

import (
	"context"
	"fmt"
	"os"

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
	Connection func() (connection.IConnection, error)
}

// NewDeleteCommand command for deleting kafkas.
func NewDeleteCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Config:     f.Config,
		Connection: f.Connection,
	}

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a Kafka instance",
		Long:  "Permanently delete a Kafka instance",
		Example: heredoc.Doc(`
			$ rhoas kafka delete
			$ rhoas kafka delete --id=1iSY6RQ3JKI8Q0OTmjQFd3ocFRg
		`),
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			cfg, err := opts.Config.Load()
			if err != nil {
				return fmt.Errorf("Error loading config: %w", err)
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
	cfg, err := opts.Config.Load()
	if err != nil {
		return fmt.Errorf("Error loading config: %w", err)
	}

	connection, err := opts.Connection()
	if err != nil {
		return fmt.Errorf("Can't create connection: %w", err)
	}

	client := connection.NewMASClient()

	kafkaID := opts.id

	var confirmDeleteAction bool
	var promptConfirmAction = &survey.Confirm{
		Message: "Once a Kafka instance is deleted it is gone forever and cannot be recovered, are you sure you want to proceed?",
	}

	err = survey.AskOne(promptConfirmAction, &confirmDeleteAction)
	if err != nil {
		return err
	}
	if !confirmDeleteAction {
		return nil
	}

	var promptConfirmID = &survey.Input{
		Message: "Please confirm the ID of the Kafka instance you wish to permanently delete:",
	}

	var confirmedKafkaID string
	err = survey.AskOne(promptConfirmID, &confirmedKafkaID)
	if err != nil {
		return err
	}

	if confirmedKafkaID != kafkaID {
		fmt.Fprintln(os.Stderr, "The ID you entered does not match the ID of the Kafka instance that you are trying to delete. Please check that it correct and try again.")
		return nil
	}

	fmt.Fprintf(os.Stderr, "Deleting Kafka instance with ID '%v'.\n", kafkaID)
	_, _, apiErr := client.DefaultApi.DeleteKafkaById(context.Background(), kafkaID).Execute()

	if apiErr.Error() != "" {
		return fmt.Errorf("Error deleting Kafka instance: %w", apiErr)
	}

	fmt.Fprint(os.Stderr, "\nKafka instance has successfully been deleted.\n")

	currentKafka := cfg.Services.Kafka
	// this is not the current cluster, our work here is done
	if currentKafka.ClusterID != kafkaID {
		return nil
	}

	// the Kafka that was deleted is set as the user's current cluster
	// since it was deleted it should be removed from the confgi
	cfg.Services.Kafka = nil
	err = opts.Config.Save(cfg)
	if err != nil {
		return fmt.Errorf("Could not save config: %w", err)
	}

	return nil
}
