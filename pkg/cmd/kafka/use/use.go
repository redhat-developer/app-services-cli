package use

import (
	"context"
	"errors"
	"fmt"

	kasclient "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/kas/client"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/iostreams"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/kafka"

	"github.com/spf13/cobra"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"
)

type Options struct {
	id          string
	name        string
	interactive bool

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     func() (logging.Logger, error)
}

func NewUseCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
	}

	cmd := &cobra.Command{
		Use:     localizer.MustLocalizeFromID("kafka.use.cmd.use"),
		Short:   localizer.MustLocalizeFromID("kafka.use.cmd.shortDescription"),
		Long:    localizer.MustLocalizeFromID("kafka.use.cmd.longDescription"),
		Example: localizer.MustLocalizeFromID("kafka.use.cmd.example"),
		Args:    cobra.RangeArgs(0, 1),
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var searchName string
			if len(args) > 0 {
				searchName = args[0]
			}
			return cmdutil.FilterValidKafkas(f, searchName)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				opts.name = args[0]
			} else if opts.id == "" {
				if !opts.IO.CanPrompt() {
					return errors.New(localizer.MustLocalizeFromID("kafka.use.error.idOrNameRequired"))
				}
				opts.interactive = true
			}

			if opts.name != "" && opts.id != "" {
				return errors.New(localizer.MustLocalizeFromID("kafka.common.error.idAndNameCannotBeUsed"))
			}

			return runUse(opts)
		},
	}

	cmd.Flags().StringVar(&opts.id, "id", "", localizer.MustLocalizeFromID("kafka.use.flag.id"))

	return cmd
}

func runUse(opts *Options) error {

	if opts.interactive {
		// run the use command interactively
		err := runInteractivePrompt(opts)
		if err != nil {
			return err
		}
	}

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	cfg, err := opts.Config.Load()
	if err != nil {
		return err
	}

	connection, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	api := connection.API()

	var res *kasclient.KafkaRequest
	ctx := context.Background()
	if opts.name != "" {
		res, _, err = kafka.GetKafkaByName(ctx, api.Kafka(), opts.name)
		if err.Error() != "" {
			return err
		}
	} else {
		res, _, err = kafka.GetKafkaByID(ctx, api.Kafka(), opts.id)
		if err.Error() != "" {
			return err
		}
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

func runInteractivePrompt(opts *Options) error {
	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	connection, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	logger.Debug(localizer.MustLocalizeFromID("common.log.debug.startingInteractivePrompt"))

	selectedKafka, err := kafka.InteractiveSelect(connection, logger)
	if err != nil {
		return err
	}

	opts.name = selectedKafka.GetName()

	return nil
}
