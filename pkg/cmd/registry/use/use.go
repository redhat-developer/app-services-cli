package use

import (
	"context"
	"errors"
	"fmt"

	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/kafka"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/serviceregistry"

	"github.com/redhat-developer/app-services-cli/internal/config"
	srsclient "github.com/redhat-developer/app-services-cli/pkg/api/srs/client"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	"github.com/spf13/cobra"
)

type Options struct {
	id          string
	name        string
	interactive bool

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     func() (logging.Logger, error)
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
		Use:     "use",
		Short:   "Use service registry",
		Long:    "",
		Example: "",
		Args:    cobra.RangeArgs(0, 1),
		// TODO
		// ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		// 	return cmdutil.FilterValidKafkas(f, toComplete)
		// },
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				opts.name = args[0]
			} else if opts.id == "" {
				if !opts.IO.CanPrompt() {
					return errors.New(opts.localizer.MustLocalize("kafka.use.error.idOrNameRequired"))
				}
				opts.interactive = true
			}

			if opts.name != "" && opts.id != "" {
				return errors.New(opts.localizer.MustLocalize("kafka.common.error.idAndNameCannotBeUsed"))
			}

			return runUse(opts)
		},
	}

	// TODO - why not using name here but other commands do use it?
	cmd.Flags().StringVar(&opts.id, "id", "", opts.localizer.MustLocalize("kafka.use.flag.id"))

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

	var registry *srsclient.Registry
	ctx := context.Background()
	if opts.name != "" {
		registry, _, err = serviceregistry.GetServiceRegistryByName(ctx, api.ServiceRegistry(), opts.name)
		if err != nil {
			return err
		}
	} else {
		registry, _, err = serviceregistry.GetServiceRegistryByID(ctx, api.ServiceRegistry(), opts.id)
		if err != nil {
			return err
		}
	}

	registryConfig := &config.ServiceRegistryConfig{
		InstanceID: registry.GetId(),
		Name:       *registry.Name,
	}

	// FIXME Use as helper?
	logger.Debug(opts.localizer.MustLocalize("kafka.create.debug.autoUseSetMessage"))
	cfg.Services.ServiceRegistry = registryConfig
	if err := opts.Config.Save(cfg); err != nil {
		saveErrMsg := "Cannot save config"
		return fmt.Errorf("%v: %w", saveErrMsg, err)
	}

	logger.Info("Successfully setup registry to " + registry.GetName())

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

	logger.Debug(opts.localizer.MustLocalize("common.log.debug.startingInteractivePrompt"))

	selectedKafka, err := kafka.InteractiveSelect(connection, logger)
	if err != nil {
		return err
	}

	opts.name = selectedKafka.GetName()

	return nil
}
