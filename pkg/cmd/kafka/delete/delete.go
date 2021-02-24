package delete

import (
	"context"
	"errors"
	"fmt"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/kas"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/iostreams"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/kafka"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"

	"github.com/AlecAivazis/survey/v2"
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/spf13/cobra"
)

type options struct {
	id    string
	force bool

	IO         *iostreams.IOStreams
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
		IO:         f.IOStreams,
	}

	cmd := &cobra.Command{
		Use:     localizer.MustLocalizeFromID("kafka.delete.cmd.use"),
		Short:   localizer.MustLocalizeFromID("kafka.delete.cmd.shortDescription"),
		Long:    localizer.MustLocalizeFromID("kafka.delete.cmd.longDescription"),
		Example: localizer.MustLocalizeFromID("kafka.delete.cmd.example"),
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			if !opts.IO.CanPrompt() && !opts.force {
				return fmt.Errorf(localizer.MustLocalize(&localizer.Config{
					MessageID: "flag.error.requiredWhenNonInteractive",
					TemplateData: map[string]interface{}{
						"Flag": "force",
					},
				}))
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			if opts.id != "" {
				return runDelete(opts)
			}

			var kafkaConfig *config.KafkaConfig
			if cfg.Services.Kafka == kafkaConfig || cfg.Services.Kafka.ClusterID == "" {
				return errors.New(localizer.MustLocalizeFromID("kafka.common.error.noKafkaSelected"))
			}

			opts.id = cfg.Services.Kafka.ClusterID

			return runDelete(opts)
		},
	}

	cmd.Flags().StringVar(&opts.id, "id", "", localizer.MustLocalizeFromID("kafka.delete.flag.id"))
	cmd.Flags().BoolVarP(&opts.force, "force", "f", false, localizer.MustLocalizeFromID("kafka.delete.flag.force"))

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
		return apiErr
	}

	kafkaName := response.GetName()

	logger.Info(localizer.MustLocalize(&localizer.Config{
		MessageID: "kafka.delete.log.info.deleting",
		TemplateData: map[string]interface{}{
			"Name": kafkaName,
		},
	}), "\n")

	if !opts.force {
		promptConfirmName := &survey.Input{
			Message: localizer.MustLocalizeFromID("kafka.delete.input.confirmName.message"),
		}

		var confirmedKafkaName string
		err = survey.AskOne(promptConfirmName, &confirmedKafkaName)
		if err != nil {
			return err
		}

		if confirmedKafkaName != kafkaName {
			logger.Info(localizer.MustLocalizeFromID("kafka.delete.log.info.incorrectNameConfirmation"))
			return nil
		}
	}

	logger.Debug(localizer.MustLocalizeFromID("kafka.delete.log.debug.deletingKafka"), fmt.Sprintf("\"%s\"", kafkaName))
	a := api.Kafka().DeleteKafkaById(context.Background(), opts.id)
	a = a.Async(true)
	_, _, apiErr = a.Execute()

	if apiErr.Error() != "" {
		return apiErr
	}

	logger.Info(localizer.MustLocalize(&localizer.Config{
		MessageID: "kafka.delete.log.info.deleteSuccess",
		TemplateData: map[string]interface{}{
			"Name": kafkaName,
		},
	}))

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
