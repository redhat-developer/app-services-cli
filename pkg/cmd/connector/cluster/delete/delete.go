package delete

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	connectormgmtclient "github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1/client"

	"github.com/spf13/cobra"
)

type options struct {
	id           string
	outputFormat string

	f           *factory.Factory
	skipConfirm bool
}

// NewDeleteCommand creates a new command to delete a connector cluster
func NewDeleteCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		f: f,
	}

	cmd := &cobra.Command{
		Use:     "delete",
		Short:   f.Localizer.MustLocalize("connector.cluster.delete.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("connector.cluster.delete.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("connector.cluster.delete.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {

			if !f.IOStreams.CanPrompt() && !opts.skipConfirm {
				return flagutil.RequiredWhenNonInteractiveError("yes")
			}

			validOutputFormats := flagutil.ValidOutputFormats
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flagutil.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			return runDelete(opts)
		},
	}
	flags := flagutil.NewFlagSet(cmd, f.Localizer)
	flags.StringVar(&opts.id, "id", "", f.Localizer.MustLocalize("connector.cluster.delete.id.flag.description"))
	flags.AddOutput(&opts.outputFormat)
	flags.AddYes(&opts.skipConfirm)

	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func runDelete(opts *options) error {
	f := opts.f

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

	var conn connection.Connection
	conn, err := f.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	api := conn.API()

	a := api.ConnectorsMgmt().ConnectorClustersApi.CreateConnectorCluster(f.Context)
	a = a.ConnectorClusterRequest(connectormgmtclient.ConnectorClusterRequest{
		Name: &opts.id,
	})
	a = a.Async(true)

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

	f.Logger.Info(f.Localizer.MustLocalize("connector.cluster.delete.info.success"))

	return nil
}

func promptConfirmDelete(opts *options) (bool, error) {
	promptConfirm := survey.Confirm{
		Message: opts.f.Localizer.MustLocalize("connector.cluster.delete.confirmDialog.message", localize.NewEntry("ID", opts.id)),
	}

	var confirmUpdate bool
	if err := survey.AskOne(&promptConfirm, &confirmUpdate); err != nil {
		return false, err
	}
	return confirmUpdate, nil
}
