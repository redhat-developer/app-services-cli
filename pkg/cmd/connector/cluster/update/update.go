package update

import (
	"github.com/redhat-developer/app-services-cli/pkg/cmd/connector/cluster/clustercmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"

	connectormgmtclient "github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1/client"
	connectorerror "github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1/error"
)

type options struct {
	id          string
	name        string
	annotations []string

	f           *factory.Factory
	skipConfirm bool
}

// NewUpdateCommand creates a new command to update a connector cluster
func NewUpdateCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		f: f,
	}

	cmd := &cobra.Command{
		Use:     "update",
		Short:   f.Localizer.MustLocalize("connector.cluster.update.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("connector.cluster.update.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("connector.cluster.update.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {

			if !f.IOStreams.CanPrompt() && !opts.skipConfirm {
				return flagutil.RequiredWhenNonInteractiveError("yes")
			}

			return runUpdate(opts)
		},
	}

	cmd.Flags().StringVar(&opts.id, "id", "", f.Localizer.MustLocalize("connector.cluster.update.flag.id.description"))
	cmd.Flags().StringVar(&opts.name, "name", "", f.Localizer.MustLocalize("connector.cluster.update.flag.name.description"))
	cmd.Flags().StringSliceVar(&opts.annotations, "annotations", []string{}, f.Localizer.MustLocalize("connector.cluster.update.flag.annotations.description"))

	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func runUpdate(opts *options) error {
	f := opts.f

	conn, err := f.Connection()
	if err != nil {
		return err
	}

	api := conn.API()

	a := api.ConnectorsMgmt().ConnectorClustersApi.UpdateConnectorClusterById(opts.f.Context, opts.id)

	data := connectormgmtclient.ConnectorClusterRequest{}

	clusterChanged := false
	if opts.name != "" {
		data.SetName(opts.name)
		clusterChanged = true
	}

	if len(opts.annotations) > 0 {
		annotationMap, annotationErr := clustercmdutil.BuildAnnotationsMap(opts.annotations)
		if annotationErr != nil {
			return annotationErr
		}
		data.SetAnnotations(annotationMap)
		clusterChanged = true
	}

	if !clusterChanged {
		opts.f.Logger.Info(opts.f.Localizer.MustLocalize("connector.cluster.update.log.info.nothingToUpdate"))
		return nil
	}

	a = a.ConnectorClusterRequest(data)

	httpRes, err := a.Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	if apiErr := connectorerror.GetAPIError(err); apiErr != nil {
		if apiErr.GetCode() == connectorerror.ERROR_7 {
			return opts.f.Localizer.MustLocalizeError("connector.cluster.update.error.notFound", localize.NewEntry("ID", opts.id))
		}

		return opts.f.Localizer.MustLocalizeError("connector.type.create.error.other", localize.NewEntry("Error", apiErr.GetReason()))
	}
	if err != nil {
		return err
	}

	return nil

}
