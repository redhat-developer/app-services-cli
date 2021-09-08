package use

import (
	"context"
	"errors"
	"fmt"
	"github.com/redhat-developer/app-services-cli/pkg/icon"

	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	kafkacmdutil "github.com/redhat-developer/app-services-cli/pkg/kafka/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"

	"github.com/redhat-developer/app-services-cli/pkg/kafka"

	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
)

type Options struct {
	id          string
	name        string
	interactive bool

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	localizer  localize.Localizer
}

func NewUseCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
	}

	cmd := &cobra.Command{
		Use:     opts.localizer.MustLocalize("kafka.use.cmd.use"),
		Short:   opts.localizer.MustLocalize("kafka.use.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("kafka.use.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("kafka.use.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.id == "" && opts.name == "" {
				if !opts.IO.CanPrompt() {
					return errors.New(opts.localizer.MustLocalize("kafka.use.error.idOrNameRequired"))
				}
				opts.interactive = true
			}

			if opts.name != "" && opts.id != "" {
				return errors.New(opts.localizer.MustLocalize("service.error.idAndNameCannotBeUsed"))
			}

			return runUse(opts)
		},
	}

	cmd.Flags().StringVar(&opts.id, "id", "", opts.localizer.MustLocalize("kafka.use.flag.id"))
	cmd.Flags().StringVar(&opts.name, "name", "", opts.localizer.MustLocalize("kafka.use.flag.name"))

	if err := kafkacmdutil.RegisterNameFlagCompletionFunc(cmd, f); err != nil {
		opts.Logger.Debug(opts.localizer.MustLocalize("kafka.common.error.load.completions.name.flag"), err)
	}

	return cmd
}

func runUse(opts *Options) error {
	if opts.interactive {
		// run the use command interactively
		err := runInteractivePrompt(opts)
		if err != nil {
			return err
		}
		// no Kafka was selected, exit program
		if opts.name == "" {
			return nil
		}
	}

	cfg, err := opts.Config.Load()
	if err != nil {
		return err
	}

	conn, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	api := conn.API()

	var res *kafkamgmtclient.KafkaRequest
	ctx := context.Background()
	if opts.name != "" {
		res, _, err = kafka.GetKafkaByName(ctx, api.Kafka(), opts.name)
		if err != nil {
			return err
		}
	} else {
		res, _, err = kafka.GetKafkaByID(ctx, api.Kafka(), opts.id)
		if err != nil {
			return err
		}
	}

	// build Kafka config object from the response
	var kafkaConfig = config.KafkaConfig{
		ClusterID: res.GetId(),
	}

	nameTmplEntry := localize.NewEntry("Name", res.GetName())
	cfg.Services.Kafka = &kafkaConfig
	if err := opts.Config.Save(cfg); err != nil {
		saveErrMsg := opts.localizer.MustLocalize("kafka.use.error.saveError", nameTmplEntry)
		return fmt.Errorf("%v: %w", saveErrMsg, err)
	}

	opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("kafka.use.log.info.useSuccess", nameTmplEntry))

	return nil
}

func runInteractivePrompt(opts *Options) error {
	conn, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	opts.Logger.Debug(opts.localizer.MustLocalize("common.log.debug.startingInteractivePrompt"))

	selectedKafka, err := kafka.InteractiveSelect(conn, opts.Logger)
	if err != nil {
		return err
	}

	opts.name = selectedKafka.GetName()

	return nil
}
