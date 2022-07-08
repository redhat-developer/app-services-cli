package use

import (
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"
	connectormgmtclient "github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1/client"

	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connectorutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"

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
		Short:   opts.f.Localizer.MustLocalize("connector.use.cmd.shortDescription"),
		Long:    opts.f.Localizer.MustLocalize("connector.use.cmd.longDescription"),
		Example: opts.f.Localizer.MustLocalize("connector.use.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {

			if opts.id == "" && opts.name == "" {
				if !opts.f.IOStreams.CanPrompt() {
					return opts.f.Localizer.MustLocalizeError("connector.use.error.idOrNameRequired")
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
	flags.StringVar(&opts.id, "id", "", opts.f.Localizer.MustLocalize("connector.use.flag.id"))
	flags.StringVar(&opts.name, "name", "", opts.f.Localizer.MustLocalize("connector.use.flag.name"))

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

	conn, err := opts.f.Connection(connection.DefaultConfigSkipMasAuth)
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

	var connectorInstance *connectormgmtclient.Connector
	api := conn.API().ConnectorsMgmt()

	if opts.id != "" {
		connectorInstance, err = connectorutil.GetConnectorByID(&api, opts.id, opts.f)
		if err != nil {
			return err
		}
	}

	if opts.name != "" {
		connectorInstance, err = connectorutil.GetConnectorByName(&api, opts.name, opts.f)
		if err != nil {
			return err
		}
	}

	currCtx.ConnectorID = *connectorInstance.Id
	svcContext.Contexts[svcContext.CurrentContext] = *currCtx

	nameTmplEntry := localize.NewEntry("Name", connectorInstance.Name)

	if err := opts.f.ServiceContext.Save(svcContext); err != nil {
		return err
	}

	opts.f.Logger.Info(icon.SuccessPrefix(), opts.f.Localizer.MustLocalize("connector.use.log.info.useSuccess", nameTmplEntry))

	return nil
}

func runInteractivePrompt(opts *options, conn *connection.Connection) error {
	opts.f.Logger.Debug(opts.f.Localizer.MustLocalize("common.log.debug.startingInteractivePrompt"))

	selectedConnector, err := connectorutil.InteractiveSelect(*conn, opts.f)
	if err != nil {
		return err
	}

	if selectedConnector != nil {
		opts.name = selectedConnector.Name

	} else {
		// conector instance found
		opts.name = ""
	}

	return nil
}
