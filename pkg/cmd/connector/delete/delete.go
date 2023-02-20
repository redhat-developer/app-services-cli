package delete

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/connector/connectorcmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"

	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connectorutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	connectormgmtclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/connectormgmt/apiv1/client"

	"github.com/spf13/cobra"
)

type options struct {
	id           string
	name         string
	outputFormat string

	f           *factory.Factory
	skipConfirm bool
}

// NewDeleteCommand creates a new command to delete a connector
func NewDeleteCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		f: f,
	}

	cmd := &cobra.Command{
		Use:     "delete",
		Short:   f.Localizer.MustLocalize("connector.delete.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("connector.delete.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("connector.delete.cmd.example"),
		Hidden:  false,
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			validOutputFormats := flagutil.ValidOutputFormats
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flagutil.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			if opts.name != "" && opts.id != "" {
				return opts.f.Localizer.MustLocalizeError("service.error.idAndNameCannotBeUsed")
			}

			return runDelete(opts)
		},
	}

	flags := connectorcmdutil.NewFlagSet(cmd, f)
	flags.AddOutput(&opts.outputFormat)
	_ = flags.AddConnectorID(&opts.id)
	_ = flags.AddConnectorName(&opts.name)
	flags.AddYes(&opts.skipConfirm)

	return cmd
}

func runDelete(opts *options) error {
	f := opts.f

	var conn connection.Connection
	conn, err := f.Connection()
	if err != nil {
		return err
	}

	api := conn.API()
	connectorMgmt := api.ConnectorsMgmt()

	var connector *connectormgmtclient.Connector

	if opts.id != "" {
		connector, err = connectorutil.GetConnectorByID(&connectorMgmt, opts.id, f)
		if err != nil {
			return err
		}
	}

	if opts.name != "" {
		connector, err = connectorutil.GetConnectorByName(&connectorMgmt, opts.name, f)
		if err != nil {
			return err
		}
	}

	if connector == nil {
		connector, err = contextutil.GetCurrentConnectorInstance(&conn, f)
		if err != nil {
			return err
		}
	}

	opts.id = connector.GetId()

	if !opts.skipConfirm {
		confirm, promptErr := promptConfirmDelete(opts)
		if promptErr != nil {
			return promptErr
		}
		if !confirm {
			opts.f.Logger.Debug("User has chosen to not delete connector cluster")
			return nil
		}
	}

	a := connectorMgmt.ConnectorsApi.DeleteConnector(f.Context, opts.id)

	_, httpRes, err := a.Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	if err != nil {
		return err
	}

	f.Logger.Info(icon.SuccessPrefix(), f.Localizer.MustLocalize("connector.delete.info.success"))

	svcContext, err := f.ServiceContext.Load()
	if err != nil {
		return err
	}

	currCtx, err := contextutil.GetCurrentContext(svcContext, f.Localizer)
	if err != nil {
		return err
	}

	// this is not the current instance, our work here is done
	if currCtx.ConnectorID != connector.GetId() {
		return nil
	}

	// the connector that was deleted is set as the user's current cluster
	// since it was deleted it should be removed from the context
	currCtx.ConnectorID = ""
	svcContext.Contexts[svcContext.CurrentContext] = *currCtx

	if err := opts.f.ServiceContext.Save(svcContext); err != nil {
		return err
	}

	return nil
}

func promptConfirmDelete(opts *options) (bool, error) {
	promptConfirm := survey.Confirm{
		Message: opts.f.Localizer.MustLocalize("connector.delete.confirmDialog.message", localize.NewEntry("ID", opts.id)),
	}

	var confirmDelete bool
	if err := survey.AskOne(&promptConfirm, &confirmDelete); err != nil {
		return false, err
	}
	return confirmDelete, nil
}
