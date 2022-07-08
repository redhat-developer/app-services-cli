package use

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"
	connectormgmtclient "github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1/client"

	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
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

	IO             *iostreams.IOStreams
	Connection     factory.ConnectionFunc
	Logger         logging.Logger
	localizer      localize.Localizer
	Context        context.Context
	ServiceContext servicecontext.IContext
}

func NewUseCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Connection:     f.Connection,
		Logger:         f.Logger,
		IO:             f.IOStreams,
		localizer:      f.Localizer,
		Context:        f.Context,
		ServiceContext: f.ServiceContext,
	}

	cmd := &cobra.Command{
		Use:     "use",
		Short:   opts.localizer.MustLocalize("connector.use.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("connector.use.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("connector.use.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {

			if opts.id == "" && opts.name == "" {
				if !opts.IO.CanPrompt() {
					return opts.localizer.MustLocalizeError("connector.use.error.idOrNameRequired")
				}
				opts.interactive = true
			}

			if opts.name != "" && opts.id != "" {
				return opts.localizer.MustLocalizeError("service.error.idAndNameCannotBeUsed")
			}

			return runUse(opts)
		},
	}

	flags := flagutil.NewFlagSet(cmd, opts.localizer)
	flags.StringVar(&opts.id, "id", "", opts.localizer.MustLocalize("connector.use.flag.id"))
	flags.StringVar(&opts.name, "name", "", opts.localizer.MustLocalize("connector.use.flag.name"))

	return cmd
}

func runUse(opts *options) error {

	svcContext, err := opts.ServiceContext.Load()
	if err != nil {
		return err
	}

	currCtx, err := contextutil.GetCurrentContext(svcContext, opts.localizer)
	if err != nil {
		return err
	}

	conn, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	if opts.interactive {
		// run the use command interactively
		err = runInteractivePrompt(opts, &conn)
		if err != nil {
			return err
		}
		// no Kafka was connector, exit program
		if opts.name == "" {
			return nil
		}
	}

	var connectorInstance *connectormgmtclient.Connector
	api := conn.API().ConnectorsMgmt()

	if opts.id != "" {
		connectorInstance, err = connectorutil.GetConnectorByID(opts.Context, &api, opts.id)
		if err != nil {
			return err
		}
	}

	if opts.name != "" {
		connectorInstance, err = connectorutil.GetConnectorByName(opts.Context, &api, opts.name, opts.localizer)
		if err != nil {
			return err
		}
	}

	currCtx.ConnectorID = *connectorInstance.Id
	svcContext.Contexts[svcContext.CurrentContext] = *currCtx

	nameTmplEntry := localize.NewEntry("Name", connectorInstance.Name)

	if err := opts.ServiceContext.Save(svcContext); err != nil {
		return err
	}

	opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("connector.use.log.info.useSuccess", nameTmplEntry))

	return nil
}

func runInteractivePrompt(opts *options, conn *connection.Connection) error {
	opts.Logger.Debug(opts.localizer.MustLocalize("common.log.debug.startingInteractivePrompt"))

	selectedConnector, err := connectorutil.InteractiveSelect(opts.Context, *conn, opts.Logger, opts.localizer)
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
