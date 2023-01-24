package use

import (
	connectormgmtclient "github.com/jackdelahunt/app-services-sdk-core/app-services-sdk-go/connectormgmt/apiv1/client"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"

	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/redhat-developer/app-services-cli/pkg/shared/namespaceutil"

	"github.com/spf13/cobra"
)

type options struct {
	id          string
	name        string
	interactive bool

	f *factory.Factory
}

func NewUseCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		f: f,
	}

	cmd := &cobra.Command{
		Use:     "use",
		Short:   opts.f.Localizer.MustLocalize("namespace.use.cmd.shortDescription"),
		Long:    opts.f.Localizer.MustLocalize("namespace.use.cmd.longDescription"),
		Example: opts.f.Localizer.MustLocalize("namespace.use.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {

			if opts.id == "" && opts.name == "" {
				if !opts.f.IOStreams.CanPrompt() {
					return opts.f.Localizer.MustLocalizeError("namespace.use.error.idOrNameRequired")
				}
				opts.interactive = true
			}

			if opts.name != "" && opts.id != "" {
				return opts.f.Localizer.MustLocalizeError("service.error.idAndNameCannotBeUsed")
			}

			return runUse(opts)
		},
	}

	flags := flagutil.NewFlagSet(cmd, opts.f.Localizer)
	flags.StringVar(&opts.id, "id", "", opts.f.Localizer.MustLocalize("namespace.use.flag.id"))
	flags.StringVar(&opts.name, "name", "", opts.f.Localizer.MustLocalize("namespace.use.flag.name"))

	return cmd
}

func runUse(opts *options) error {

	svcContext, err := opts.f.ServiceContext.Load()
	if err != nil {
		return err
	}

	currCtx, err := contextutil.GetCurrentContext(svcContext, opts.f.Localizer)
	if err != nil {
		return err
	}

	conn, err := opts.f.Connection()
	if err != nil {
		return err
	}

	if opts.interactive {
		// run the use command interactively
		err = runInteractivePrompt(opts, &conn)
		if err != nil {
			return err
		}
		// no connector was found, exit program
		if opts.name == "" {
			return nil
		}
	}

	var namespace *connectormgmtclient.ConnectorNamespace
	api := conn.API().ConnectorsMgmt()

	if opts.id != "" {
		namespace, err = namespaceutil.GetNamespaceByID(&api, opts.id, opts.f)
		if err != nil {
			return err
		}
	}

	if opts.name != "" {
		namespace, err = namespaceutil.GetNamespaceByName(&api, opts.name, opts.f)
		if err != nil {
			return err
		}
	}

	currCtx.NamespaceID = namespace.GetId()
	svcContext.Contexts[svcContext.CurrentContext] = *currCtx

	if err := opts.f.ServiceContext.Save(svcContext); err != nil {
		return err
	}

	opts.f.Logger.Info(icon.SuccessPrefix(), opts.f.Localizer.MustLocalize("namespace.use.log.info.useSuccess", localize.NewEntry("Name", namespace.GetName())))

	return nil
}

func runInteractivePrompt(opts *options, conn *connection.Connection) error {
	opts.f.Logger.Debug(opts.f.Localizer.MustLocalize("common.log.debug.startingInteractivePrompt"))

	selectedNamespace, err := namespaceutil.InteractiveSelect(*conn, opts.f)
	if err != nil {
		return err
	}

	if selectedNamespace != nil {
		opts.name = selectedNamespace.Name

	} else {
		opts.name = ""
	}

	return nil
}
