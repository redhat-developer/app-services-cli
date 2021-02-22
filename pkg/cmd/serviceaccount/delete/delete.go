package delete

import (
	"context"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/iostreams"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"
	"github.com/spf13/cobra"
)

type Options struct {
	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection func() (connection.Connection, error)
	Logger     func() (logging.Logger, error)

	id    string
	force bool
}

// NewDeleteCommand creates a new command to delete a service account
func NewDeleteCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
	}

	cmd := &cobra.Command{
		Use:     localizer.MustLocalizeFromID("serviceAccount.delete.cmd.use"),
		Short:   localizer.MustLocalizeFromID("serviceAccount.delete.cmd.shortDescription"),
		Long:    localizer.MustLocalizeFromID("serviceAccount.delete.cmd.longDescription"),
		Example: localizer.MustLocalizeFromID("serviceAccount.delete.cmd.example"),
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runDelete(opts)
		},
	}

	cmd.Flags().StringVar(&opts.id, "id", "", localizer.MustLocalizeFromID("serviceAccount.delete.flag.id.description"))
	cmd.Flags().BoolVarP(&opts.force, "force", "f", false, localizer.MustLocalizeFromID("serviceAccount.delete.flag.force.description"))

	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func runDelete(opts *Options) (err error) {
	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	if !opts.force {
		var confirmDelete bool
		promptConfirmDelete := &survey.Confirm{
			Message: localizer.MustLocalize(&localizer.Config{
				MessageID: "serviceAccount.delete.input.confirmDelete.message",
				TemplateData: map[string]interface{}{
					"ID": opts.id,
				},
			}),
		}

		err = survey.AskOne(promptConfirmDelete, &confirmDelete)
		if err != nil {
			return err
		}

		if !confirmDelete {
			logger.Debug(localizer.MustLocalizeFromID("serviceAccount.delete.log.debug.deleteNotConfirmed"))
			return nil
		}
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

	a := api.Kafka().DeleteServiceAccount(context.Background(), opts.id)
	_, _, apiErr := a.Execute()

	if apiErr.Error() != "" {
		return fmt.Errorf("%v: %w", localizer.MustLocalizeFromID("serviceAccount.delete.error.unableToDelete"), apiErr)
	}

	logger.Info(localizer.MustLocalizeFromID("serviceAccount.delete.log.info.deleteSuccess"))

	return nil
}
