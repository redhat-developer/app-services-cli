package start

import (
	"github.com/redhat-developer/app-services-cli/pkg/cmd/connector/connectorcmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"

	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	connectormgmtclient "github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1/client"
	"github.com/spf13/cobra"
)

type options struct {
	outputFormat string
	f            *factory.Factory
	connectorID  string
}

// NewStartCommand creates a new command to start a connector
func NewStartCommand(f *factory.Factory) *cobra.Command {

	opts := &options{
		f: f,
	}

	cmd := &cobra.Command{
		Use:     "start",
		Short:   f.Localizer.MustLocalize("connector.start.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("connector.start.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("connector.start.cmd.example"),
		Hidden:  false,
		RunE: func(cmd *cobra.Command, args []string) error {

			validOutputFormats := flagutil.ValidOutputFormats
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flagutil.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			return runUpdateCommand(opts)
		},
	}

	flags := connectorcmdutil.NewFlagSet(cmd, f)
	flags.AddConnectorID(&opts.connectorID)
	flags.AddOutput(&opts.outputFormat)

	return cmd

}

func runUpdateCommand(opts *options) error {
	f := opts.f

	var conn connection.Connection
	conn, err := f.Connection()
	if err != nil {
		return err
	}

	if opts.connectorID == "" {
		connectorInstance, instance_err := contextutil.GetCurrentConnectorInstance(&conn, f)
		if instance_err != nil {
			return instance_err
		}

		opts.connectorID = *connectorInstance.Id
	}

	api := conn.API()

	patch := make(map[string]interface{})
	patch["desired_state"] = connectormgmtclient.CONNECTORDESIREDSTATE_READY
	a := api.ConnectorsMgmt().ConnectorsApi.PatchConnector(f.Context, opts.connectorID).Body(patch)

	response, httpRes, err := a.Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	if err != nil {
		return err
	}

	if err = dump.Formatted(f.IOStreams.Out, opts.outputFormat, response); err != nil {
		return err
	}

	f.Logger.Info(f.Localizer.MustLocalize("connector.update.info.success"))

	return nil
}
