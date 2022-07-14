package delete

import (
	"context"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/redhat-developer/app-services-cli/pkg/shared/serviceregistryutil"
	"github.com/spf13/cobra"

	srsmgmtv1client "github.com/redhat-developer/app-services-sdk-go/registrymgmt/apiv1/client"
)

type options struct {
	id    string
	name  string
	force bool

	IO             *iostreams.IOStreams
	Connection     factory.ConnectionFunc
	Logger         logging.Logger
	localizer      localize.Localizer
	Context        context.Context
	ServiceContext servicecontext.IContext
}

func NewDeleteCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Connection:     f.Connection,
		Logger:         f.Logger,
		IO:             f.IOStreams,
		localizer:      f.Localizer,
		Context:        f.Context,
		ServiceContext: f.ServiceContext,
	}

	cmd := &cobra.Command{
		Use:     "delete",
		Short:   f.Localizer.MustLocalize("registry.cmd.delete.shortDescription"),
		Long:    f.Localizer.MustLocalize("registry.cmd.delete.longDescription"),
		Example: f.Localizer.MustLocalize("registry.cmd.delete.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if !opts.IO.CanPrompt() && !opts.force {
				return flagutil.RequiredWhenNonInteractiveError("yes")
			}

			if opts.name != "" && opts.id != "" {
				return opts.localizer.MustLocalizeError("service.error.idAndNameCannotBeUsed")
			}

			if opts.id != "" || opts.name != "" {
				return runDelete(opts)
			}

			registryInstance, err := contextutil.GetCurrentRegistryInstance(f)
			if err != nil {
				return err
			}

			opts.id = registryInstance.GetId()

			return runDelete(opts)
		},
	}

	cmd.Flags().StringVar(&opts.name, "name", "", opts.localizer.MustLocalize("registry.cmd.delete.flag.name.description"))
	cmd.Flags().StringVar(&opts.id, "id", "", opts.localizer.MustLocalize("registry.delete.flag.id"))
	cmd.Flags().BoolVarP(&opts.force, "yes", "y", false, opts.localizer.MustLocalize("registry.delete.flag.yes"))

	return cmd
}

func runDelete(opts *options) error {

	svcContext, err := opts.ServiceContext.Load()
	if err != nil {
		return err
	}

	currCtx, err := contextutil.GetCurrentContext(svcContext, opts.localizer)
	if err != nil {
		return err
	}

	conn, err := opts.Connection()
	if err != nil {
		return err
	}

	api := conn.API()

	var registry *srsmgmtv1client.Registry
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

	a := api.ServiceRegistryMgmt().DeleteRegistry(opts.Context, registry.GetId())
	_, err = a.Execute()

	if err != nil {
		return err
	}

	opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("registry.delete.log.info.deleteSuccess", localize.NewEntry("Name", registryName)))

	currentContextRegistry := currCtx.ServiceRegistryID
	// this is not the current cluster, our work here is done
	if currentContextRegistry != opts.id {
		return nil
	}

	// the service that was deleted is set as the user's current instance
	// since it was deleted it should be removed from the context

	currCtx.ServiceRegistryID = ""
	svcContext.Contexts[svcContext.CurrentContext] = *currCtx

	if err := opts.ServiceContext.Save(svcContext); err != nil {
		return err
	}

	return nil
}
