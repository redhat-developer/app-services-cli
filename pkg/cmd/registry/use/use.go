package use

import (
	"context"
	"fmt"

	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/redhat-developer/app-services-cli/pkg/shared/profileutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/serviceregistryutil"

	srsmgmtv1 "github.com/redhat-developer/app-services-sdk-go/registrymgmt/apiv1/client"
	"github.com/spf13/cobra"
)

type options struct {
	id          string
	name        string
	interactive bool

	IO             *iostreams.IOStreams
	Config         config.IConfig
	Connection     factory.ConnectionFunc
	Logger         logging.Logger
	localizer      localize.Localizer
	Context        context.Context
	ServiceContext servicecontext.IContext
}

func NewUseCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Config:         f.Config,
		Connection:     f.Connection,
		Logger:         f.Logger,
		IO:             f.IOStreams,
		localizer:      f.Localizer,
		Context:        f.Context,
		ServiceContext: f.ServiceContext,
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
					return opts.localizer.MustLocalizeError("registry.use.error.idOrNameRequired")
				}
				opts.interactive = true
			}

			if opts.name != "" && opts.id != "" {
				return opts.localizer.MustLocalizeError("service.error.idAndNameCannotBeUsed")
			}

			return runUse(opts)
		},
	}

	cmd.Flags().StringVar(&opts.id, "id", "", opts.localizer.MustLocalize("registry.use.flag.id"))
	cmd.Flags().StringVar(&opts.name, "name", "", opts.localizer.MustLocalize("registry.use.flag.name"))

	return cmd
}

func runUse(opts *options) error {
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

	svcContext, err := opts.ServiceContext.Load()
	if err != nil {
		return err
	}

	profileHandler := &profileutil.ContextHandler{
		Context:   svcContext,
		Localizer: opts.localizer,
	}

	currCtx, err := profileHandler.GetCurrentContext()
	if err != nil {
		return err
	}

	svcConfig, err := profileHandler.GetContext(currCtx)
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
		registry, _, err = serviceregistryutil.GetServiceRegistryByName(opts.Context, api.ServiceRegistryMgmt(), opts.name)
		if err != nil {
			return err
		}
	} else {
		registry, _, err = serviceregistryutil.GetServiceRegistryByID(opts.Context, api.ServiceRegistryMgmt(), opts.id)
		if err != nil {
			return err
		}
	}

	nameTmplEntry := localize.NewEntry("Name", registry.GetName())
	svcConfig.ServiceRegistryID = registry.GetId()
	svcContext.Contexts[svcContext.CurrentContext] = *svcConfig

	if err := opts.ServiceContext.Save(svcContext); err != nil {
		saveErrMsg := opts.localizer.MustLocalize("registry.use.error.saveError", nameTmplEntry)
		return fmt.Errorf("%v: %w", saveErrMsg, err)
	}

	opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("registry.use.log.info.useSuccess", nameTmplEntry))

	return nil
}

func runInteractivePrompt(opts *options) error {
	conn, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	opts.Logger.Debug(opts.localizer.MustLocalize("common.log.debug.startingInteractivePrompt"))

	selectedRegistry, err := serviceregistryutil.InteractiveSelect(opts.Context, conn, opts.Logger)
	if err != nil {
		return err
	}

	opts.name = selectedRegistry.GetName()

	return nil
}
