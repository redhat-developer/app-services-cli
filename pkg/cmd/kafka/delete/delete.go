package delete

import (
	"context"
	"fmt"
	"os"

	"github.com/MakeNowJust/heredoc"

	"github.com/AlecAivazis/survey/v2"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/config"
	"github.com/spf13/cobra"
)

type options struct {
	id string

	cfg *config.Config
}

// NewDeleteCommand command for deleting kafkas.
func NewDeleteCommand() *cobra.Command {
	opts := &options{}

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a Kafka cluster",
		Long:  "Permanently delete a Kafka cluster",
		Example: heredoc.Doc(`
			$ rhoas kafka delete
			$ rhoas kafka delete --id=1iSY6RQ3JKI8Q0OTmjQFd3ocFRg
		`),
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			cfg, err := config.Load()
			if err != nil {
				return fmt.Errorf("Error loading config: %w", err)
			}
			opts.cfg = cfg

			if opts.id != "" {
				return runDelete(opts)
			}

			var kafkaConfig *config.KafkaConfig
			if cfg.Services.Kafka == kafkaConfig || cfg.Services.Kafka.ClusterID == "" {
				return fmt.Errorf("No Kafka cluster selected. Use the '--id' flag or set one in context with the 'use' command")
			}

			opts.id = cfg.Services.Kafka.ClusterID

			return runDelete(opts)
		},
	}

	cmd.Flags().StringVar(&opts.id, "id", "", "ID of the Kafka cluster you want to delete. If not set, the currently selected Kafka cluster will be used")

	return cmd
}

func runDelete(opts *options) error {
	cfg := opts.cfg
	connection, err := cfg.Connection()
	if err != nil {
		return fmt.Errorf("Can't create connection: %w", err)
	}

	client := connection.NewMASClient()

	kafkaID := opts.id

	var confirmDeleteAction bool
	var promptConfirmAction = &survey.Confirm{
		Message: "Once a Kafka cluster is deleted it is gone forever and cannot be recovered, are you sure you want to proceed?",
	}

	err = survey.AskOne(promptConfirmAction, &confirmDeleteAction)
	if err != nil {
		return err
	}
	if !confirmDeleteAction {
		return nil
	}

	var promptConfirmID = &survey.Input{
		Message: "Please confirm the ID of the Kafka cluster you wish to permanently delete:",
	}

	var confirmedKafkaID string
	err = survey.AskOne(promptConfirmID, &confirmedKafkaID)
	if err != nil {
		return err
	}

	if confirmedKafkaID != kafkaID {
		fmt.Fprintln(os.Stderr, "The ID you entered does not match the ID of the Kafka cluster that you are trying to delete. Please check that it correct and try again.")
		return nil
	}

	fmt.Fprintf(os.Stderr, "Deleting Kafka cluster with ID '%v'.\n", kafkaID)
	_, _, err = client.DefaultApi.DeleteKafkaById(context.Background(), kafkaID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error deleting Kafka cluster: %v", err)
	}

	fmt.Fprint(os.Stderr, "Kafka cluster has successfully been deleted.")

	currentKafka := cfg.Services.Kafka
	// this is not the current cluster, our work here is done
	if currentKafka.ClusterID != kafkaID {
		return nil
	}

	// the Kafka that was deleted is set as the user's current cluster
	// since it was deleted it should be removed from the confgi
	cfg.Services.RemoveKafka()
	err = config.Save(cfg)
	if err != nil {
		return fmt.Errorf("Could not save config: %w", err)
	}

	return nil
}
