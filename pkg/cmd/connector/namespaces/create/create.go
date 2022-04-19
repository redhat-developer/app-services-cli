package create

import (
	"net/http"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"

	connectormgmtclient "github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1/client"
)

type options struct {
	f *factory.Factory

	name         string
	eval         bool
	outputFormat string
}

// NewCreateCommand creates a new namespace
func NewCreateCommand(f *factory.Factory) *cobra.Command {

	opts := &options{
		f: f,
	}

	cmd := &cobra.Command{
		Use:     "create",
		Short:   f.Localizer.MustLocalize("connector.namespace.create.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("connector.namespace.create.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("connector.namespace.create.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCreate(opts)
		},
	}

	flags := flagutil.NewFlagSet(cmd, f.Localizer)
	flags.StringVar(&opts.name, "name", "", f.Localizer.MustLocalize("connector.namespace.create.flag.name.description"))
	flags.BoolVar(&opts.eval, "eval", false, f.Localizer.MustLocalize("connector.namespace.create.flag.eval.description"))
	flags.AddOutput(&opts.outputFormat)

	return cmd
}

func runCreate(opts *options) error {

	f := opts.f
	conn, err := f.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	api := conn.API()

	a := api.ConnectorsMgmt().ConnectorNamespacesApi.CreateEvaluationNamespace(f.Context)

	connectorNameSpaceEvalReq := connectormgmtclient.ConnectorNamespaceEvalRequest{
		Name: &opts.name,
	}

	var connector connectormgmtclient.ConnectorNamespace
	var newErr error
	var httpRes *http.Response

	if opts.eval {

		a := api.ConnectorsMgmt().ConnectorNamespacesApi.CreateEvaluationNamespace(f.Context)

		a = a.ConnectorNamespaceEvalRequest(connectorNameSpaceEvalReq)
		connector, httpRes, newErr = a.Execute()

	} else {
		a = a.ConnectorNamespaceEvalRequest(connectorNameSpaceEvalReq)
		connector, httpRes, newErr = a.Execute()
	}

	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	if newErr != nil {
		return newErr
	}

	if newErr = dump.Formatted(f.IOStreams.Out, opts.outputFormat, connector); newErr != nil {
		return newErr
	}

	f.Logger.Info(f.Localizer.MustLocalize("connector.namespace.create.info.success"))

	return nil

}
