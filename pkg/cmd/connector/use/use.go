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
		Short:   "set connector",
		Long:    "set connector",
		Example: "",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {

			if opts.id == "" && opts.name == "" {
				if !opts.IO.CanPrompt() {
					// TODO
					return opts.localizer.MustLocalizeError("kafka.use.error.idOrNameRequired")
				}
				opts.interactive = true
			}

			if opts.name != "" && opts.id != "" {
				// TODO
				return opts.localizer.MustLocalizeError("service.error.idAndNameCannotBeUsed")
			}

			return runUse(opts)
		},
	}

	flags := flagutil.NewFlagSet(cmd, opts.localizer)
	flags.StringVar(&opts.id, "id", "", opts.localizer.MustLocalize("kafka.use.flag.id"))       // TODO
	flags.StringVar(&opts.name, "name", "", opts.localizer.MustLocalize("kafka.use.flag.name")) // TODO

	return cmd
}

func runUse(opts *options) error {

	if opts.interactive {
		// run the use command interactively
		err := runInteractivePrompt(opts)
		if err != nil {
			return err
		}
		// no Kafka was connector, exit program
		if opts.name == "" {
			return nil
		}
	}

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

	var connectorInstance *connectormgmtclient.Connector

	if opts.id != "" {
		connectorInstance, err = connectorutil.GetConnectorByID(opts.Context, conn.API().ConnectorsMgmt(), opts.id)
		if err != nil {
			return err
		}
	}

	if opts.name != "" {
		connectorInstance, err = connectorutil.GetConnectorByName(opts.Context, conn.API().ConnectorsMgmt(), opts.name)
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

	// TODO
	opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("kafka.use.log.info.useSuccess", nameTmplEntry))

	return nil
}

func runInteractivePrompt(opts *options) error {
	conn, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	opts.Logger.Debug(opts.localizer.MustLocalize("common.log.debug.startingInteractivePrompt"))

	selectedConnector, err := connectorutil.InteractiveSelect(opts.Context, conn, opts.Logger, opts.localizer)
	if err != nil {
		return err
	}

	opts.name = selectedConnector.Name
	return nil
}
