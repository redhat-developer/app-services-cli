package use

import (
	"context"
	"fmt"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/kas"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/kafka"

	"github.com/spf13/cobra"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
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
		Use:     localizer.MustLocalizeFromID("kafka.use.cmd.use"),
		Short:   localizer.MustLocalizeFromID("kafka.use.cmd.shortDescription"),
		Long:    localizer.MustLocalizeFromID("kafka.use.cmd.longDescription"),
		Example: localizer.MustLocalizeFromID("kafka.use.cmd.example"),
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runUse(opts)
		},
	}

	cmd.Flags().StringVar(&opts.id, "id", "", localizer.MustLocalizeFromID("kafka.use.flag.id"))
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

	res, _, apiErr := api.Kafka().GetKafkaById(context.Background(), opts.id).Execute()
	if kas.IsErr(apiErr, kas.ErrorNotFound) {
		return kafka.ErrorNotFound(opts.id)
	}

	if apiErr.Error() != "" {
		return apiErr
	}

	// build Kafka config object from the response
	var kafkaConfig config.KafkaConfig = config.KafkaConfig{
		ClusterID: *res.Id,
	}

	cfg.Services.Kafka = &kafkaConfig
	if err := opts.Config.Save(cfg); err != nil {
		saveErrMsg := localizer.MustLocalize(&localizer.Config{
			MessageID: "kafka.use.error.saveError",
			TemplateData: map[string]interface{}{
				"Name": res.GetName(),
			},
		})
		return fmt.Errorf("%v: %w", saveErrMsg, err)
	}

	logger.Info(localizer.MustLocalize(&localizer.Config{
		MessageID: "kafka.use.log.info.useSuccess",
		TemplateData: map[string]interface{}{
			"Name": res.GetName(),
		},
	}))

	return nil
}
