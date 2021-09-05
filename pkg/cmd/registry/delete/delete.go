package delete

import (
	"context"
	"errors"
	"fmt"
	"github.com/redhat-developer/app-services-cli/internal/build"

	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/serviceregistry"

	"github.com/redhat-developer/app-services-cli/pkg/logging"

	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	"github.com/spf13/cobra"

	srsmgmtv1client "github.com/redhat-developer/app-services-sdk-go/registrymgmt/apiv1/client"
)

type options struct {
	id    string
	name  string
	force bool

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
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
		Short:   f.Localizer.MustLocalize("registry.cmd.delete.shortDescription"),
		Long:    f.Localizer.MustLocalize("registry.cmd.delete.longDescription"),
		Example: f.Localizer.MustLocalize("registry.cmd.delete.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if !opts.IO.CanPrompt() && !opts.force {
				return flag.RequiredWhenNonInteractiveError("yes")
			}

			if opts.name != "" && opts.id != "" {
				return errors.New(opts.localizer.MustLocalize("service.error.idAndNameCannotBeUsed"))
			}

			if opts.id != "" || opts.name != "" {
				return runDelete(opts)
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			var serviceRegistryConfig *config.ServiceRegistryConfig
			if cfg.Services.ServiceRegistry == serviceRegistryConfig || cfg.Services.ServiceRegistry.InstanceID == "" {
				return errors.New(opts.localizer.MustLocalize("registry.common.error.noServiceSelected"))
			}

			opts.id = cfg.Services.ServiceRegistry.InstanceID

			return runDelete(opts)
		},
	}

	cmd.Flags().StringVar(&opts.name, "name", "", opts.localizer.MustLocalize("registry.cmd.delete.flag.name.description"))
	cmd.Flags().StringVar(&opts.id, "id", "", opts.localizer.MustLocalize("registry.delete.flag.id"))
	cmd.Flags().BoolVarP(&opts.force, "yes", "y", false, opts.localizer.MustLocalize("registry.delete.flag.yes"))

	return cmd
}

func runDelete(opts *options) error {
	cfg, err := opts.Config.Load()
	if err != nil {
		return err
	}

	conn, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	api := conn.API()

	var registry *srsmgmtv1client.Registry
	ctx := context.Background()
	if opts.name != "" {
		registry, _, err = serviceregistry.GetServiceRegistryByName(ctx, api.ServiceRegistryMgmt(), opts.name)
		if err != nil {
			return err
		}
	} else {
		registry, _, err = serviceregistry.GetServiceRegistryByID(ctx, api.ServiceRegistryMgmt(), opts.id)
		if err != nil {
			return err
		}
	}

	registryName := registry.GetName()
	opts.Logger.Info(opts.localizer.MustLocalize("registry.delete.log.info.deletingService", localize.NewEntry("Name", registryName)))
	opts.Logger.Info("")

	if !opts.force {
		promptConfirmName := &survey.Input{
			Message: opts.localizer.MustLocalize("registry.delete.input.confirmName.message"),
		}

		var confirmedName string
		err = survey.AskOne(promptConfirmName, &confirmedName)
		if err != nil {
			return err
		}

		if confirmedName != registryName {
			opts.Logger.Info(opts.localizer.MustLocalize("registry.delete.log.info.incorrectNameConfirmation"))
			return nil
		}
	}

	opts.Logger.Debug("Deleting Service registry", fmt.Sprintf("\"%s\"", registryName))

	a := api.ServiceRegistryMgmt().DeleteRegistry(context.Background(), registry.GetId())
	_, err = a.Execute()

	if err != nil {
		return err
	}

	opts.Logger.Info(build.GetEmojiSuccess(), opts.localizer.MustLocalize("registry.delete.log.info.deleteSuccess", localize.NewEntry("Name", registryName)))

	currentContextRegistry := cfg.Services.ServiceRegistry
	// this is not the current cluster, our work here is done
	if currentContextRegistry == nil || currentContextRegistry.InstanceID != opts.id {
		return nil
	}

	// the service that was deleted is set as the user's current cluster
	// since it was deleted it should be removed from the config
	cfg.Services.ServiceRegistry = nil
	err = opts.Config.Save(cfg)
	if err != nil {
		return err
	}

	return nil
}
