package delete

import (
	"context"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/MakeNowJust/heredoc"
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/color"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"
	"github.com/spf13/cobra"
)

type Options struct {
	Config     config.IConfig
	Connection func() (connection.Connection, error)
	Logger     func() (logging.Logger, error)

	id string
}

// NewDeleteCommand creates a new command to delete a service account
func NewDeleteCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
	}

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a service account",
		Long: heredoc.Doc(`
			Permanently delete a service account.
		`),
		Example: heredoc.Doc(`
			# delete the service account with ID
			$ rhoas serviceaccount delete --id 173c1ad9-932d-4007-ae0f-4da74f4d2ccd
		`),
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runDelete(opts)
		},
	}

	cmd.Flags().StringVar(&opts.id, "id", "", "The unique ID of the service account to delete")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func runDelete(opts *Options) (err error) {
	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	var confirmDelete bool
	promptConfirmDelete := &survey.Confirm{
		Message: fmt.Sprintf("Are you sure you want to delete the service account with ID %v?", color.Info(opts.id)),
	}

	err = survey.AskOne(promptConfirmDelete, &confirmDelete)
	if err != nil {
		return err
	}

	if !confirmDelete {
		logger.Debug("Service account delete action was not confirmed. Exiting silently")
		return nil
	}

	return deleteServiceAccount(opts)
}

func deleteServiceAccount(opts *Options) error {
	connection, err := opts.Connection()
	if err != nil {
		return err
	}

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	api := connection.API()

	a := api.Kafka.DeleteServiceAccount(context.Background(), opts.id)
	_, _, apiErr := a.Execute()

	if apiErr.Error() != "" {
		return fmt.Errorf("Unable to delete service account: %w", apiErr)
	}

	logger.Info("Service account deleted")

	return nil
}
