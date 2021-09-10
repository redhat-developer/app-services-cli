package use

import (
	"context"
	"errors"
	"fmt"
	"github.com/redhat-developer/app-services-cli/pkg/icon"

	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/serviceregistry"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	srsmgmtv1 "github.com/redhat-developer/app-services-sdk-go/registrymgmt/apiv1/client"
	"github.com/spf13/cobra"
)

type args struct {
	id          string
	name        string
	interactive bool

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	localizer  localize.Localizer
	Context    context.Context
}

func NewUseCommand(f *factory.Factory) *cobra.Command {
	opts := &args{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "use",
		Short:   f.Localizer.MustLocalize("registry.cmd.use.shortDescription"),
		Long:    f.Localizer.MustLocalize("registry.cmd.use.longDescription"),
		Example: f.Localizer.MustLocalize("registry.cmd.use.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.id == "" && opts.name == "" {
				if !opts.IO.CanPrompt() {
					return errors.New(opts.localizer.MustLocalize("registry.use.error.idOrNameRequired"))
				}
				opts.interactive = true
			}

			if opts.name != "" && opts.id != "" {
				return errors.New(opts.localizer.MustLocalize("service.error.idAndNameCannotBeUsed"))
			}

			return runUse(opts)
		},
	}

	cmd.Flags().StringVar(&opts.id, "id", "", opts.localizer.MustLocalize("registry.use.flag.id"))
	cmd.Flags().StringVar(&opts.name, "name", "", opts.localizer.MustLocalize("registry.use.flag.name"))

	return cmd
}

func runUse(opts *args) error {
	if opts.interactive {
		// run the use command interactively
		err := runInteractivePrompt(opts)
		if err != nil {
			return err
		}
		// no service was selected, exit program
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

	var registry *srsmgmtv1.Registry
	if opts.name != "" {
		registry, _, err = serviceregistry.GetServiceRegistryByName(opts.Context, api.ServiceRegistryMgmt(), opts.name)
		if err != nil {
			return err
		}
	} else {
		registry, _, err = serviceregistry.GetServiceRegistryByID(opts.Context, api.ServiceRegistryMgmt(), opts.id)
		if err != nil {
			return err
		}
	}

	registryConfig := &config.ServiceRegistryConfig{
		InstanceID: registry.GetId(),
		Name:       *registry.Name,
	}

	nameTmplEntry := localize.NewEntry("Name", registry.GetName())
	cfg.Services.ServiceRegistry = registryConfig
	if err := opts.Config.Save(cfg); err != nil {
		saveErrMsg := opts.localizer.MustLocalize("registry.use.error.saveError", nameTmplEntry)
		return fmt.Errorf("%v: %w", saveErrMsg, err)
	}

	opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("registry.use.log.info.useSuccess", nameTmplEntry))

	return nil
}

func runInteractivePrompt(opts *args) error {
	conn, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	opts.Logger.Debug(opts.localizer.MustLocalize("common.log.debug.startingInteractivePrompt"))

	selectedRegistry, err := serviceregistry.InteractiveSelect(opts.Context, conn, opts.Logger)
	if err != nil {
		return err
	}

	opts.name = selectedRegistry.GetName()

	return nil
}
