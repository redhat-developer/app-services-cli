package delete

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	srsclient "github.com/redhat-developer/app-services-cli/pkg/api/srs/client"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/serviceregistry"

	"github.com/redhat-developer/app-services-cli/pkg/iostreams"

	"github.com/redhat-developer/app-services-cli/pkg/logging"

	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	"github.com/spf13/cobra"
)

type options struct {
	id    string
	name  string
	force bool

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     func() (logging.Logger, error)
	localizer  localize.Localizer
}

func NewDeleteCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
	}

	cmd := &cobra.Command{
		Use:     "delete",
		Short:   "Delete service registry",
		Long:    "",
		Example: "",
		Args:    cobra.RangeArgs(0, 1),
		// TODO make this more generic?
		// ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		// 	return cmdutil.FilterValidKafkas(f, toComplete)
		// },
		RunE: func(cmd *cobra.Command, args []string) error {
			if !opts.IO.CanPrompt() && !opts.force {
				return flag.RequiredWhenNonInteractiveError("yes")
			}

			if len(args) > 0 {
				opts.name = args[0]
			}

			if opts.name != "" && opts.id != "" {
				return errors.New(opts.localizer.MustLocalize("kafka.common.error.idAndNameCannotBeUsed"))
			}

			if opts.id != "" || opts.name != "" {
				return runDelete(opts)
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			var serviceRegistryConfig *config.ServiceRegistryConfig
			// TODO replace int with string
			if cfg.Services.ServiceRegistry == serviceRegistryConfig || cfg.Services.ServiceRegistry.InstanceID == 0 {
				return errors.New(opts.localizer.MustLocalize("kafka.common.error.noKafkaSelected"))
			}

			opts.id = string(cfg.Services.ServiceRegistry.InstanceID)

			return runDelete(opts)
		},
	}

	cmd.Flags().StringVar(&opts.id, "id", "", opts.localizer.MustLocalize("kafka.delete.flag.id"))
	cmd.Flags().BoolVarP(&opts.force, "yes", "y", false, opts.localizer.MustLocalize("kafka.delete.flag.yes"))

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

	connection, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	api := connection.API()

	// TODO temporary until id will be moved to string

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

	registryName := registry.GetName()

	// TODO unique i18n or common?
	logger.Info(opts.localizer.MustLocalize("kafka.delete.log.info.deleting", localize.NewEntry("Name", registryName)))
	logger.Info("")

	if !opts.force {
		promptConfirmName := &survey.Input{
			Message: opts.localizer.MustLocalize("kafka.delete.input.confirmName.message"),
		}

		var confirmedName string
		err = survey.AskOne(promptConfirmName, &confirmedName)
		if err != nil {
			return err
		}

		if confirmedName != registryName {
			logger.Info(opts.localizer.MustLocalize("kafka.delete.log.info.incorrectNameConfirmation"))
			return nil
		}
	}

	logger.Debug("Deleting Service registry", fmt.Sprintf("\"%s\"", registryName))
	// TODO temporary change to int (requires api change)
	rgInt, _ := strconv.Atoi(opts.id)
	registryID := int32(rgInt)
	a := api.ServiceRegistry().DeleteRegistry(context.Background(), registryID)
	_, err = a.Execute()

	if err != nil {
		return err
	}

	logger.Info(opts.localizer.MustLocalize("kafka.delete.log.info.deleteSuccess", localize.NewEntry("Name", registryName)))

	// TODO this should be helper
	currentContextRegistry := cfg.Services.ServiceRegistry
	// this is not the current cluster, our work here is done
	if currentContextRegistry == nil || currentContextRegistry.InstanceID != registryID {
		return nil
	}

	// the Kafka that was deleted is set as the user's current cluster
	// since it was deleted it should be removed from the config
	cfg.Services.ServiceRegistry = nil
	err = opts.Config.Save(cfg)
	if err != nil {
		return err
	}

	return nil
}
