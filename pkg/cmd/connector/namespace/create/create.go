package create

import (
	"github.com/redhat-developer/app-services-cli/pkg/cmd/connector/connectorcmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"

	connectormgmtclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/connectormgmt/apiv1/client"
)

type options struct {
	f *factory.Factory

	name         string
	outputFormat string
}

// NewCreateCommand a new command to create a new namespace
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

			validator := connectorcmdutil.Validator{
				Localizer: f.Localizer,
			}

			if err := validator.ValidateNamespace(opts.name); err != nil {
				return err
			}

			return runCreate(opts)
		},
	}

	flags := flagutil.NewFlagSet(cmd, f.Localizer)
	flags.StringVar(&opts.name, "name", "", f.Localizer.MustLocalize("connector.namespace.create.flag.name.description"))
	flags.AddOutput(&opts.outputFormat)

	return cmd
}

func runCreate(opts *options) error {

	f := opts.f
	conn, err := f.Connection()
	if err != nil {
		return err
	}

	api := conn.API()

	a := api.ConnectorsMgmt().ConnectorNamespacesApi.CreateEvaluationNamespace(f.Context)

	connectorNameSpaceEvalReq := connectormgmtclient.ConnectorNamespaceEvalRequest{
		Name: &opts.name,
	}

	var newErr error

	a = a.ConnectorNamespaceEvalRequest(connectorNameSpaceEvalReq)
	namespace, _, newErr := a.Execute()

	if newErr != nil {
		return newErr
	}

	if newErr = dump.Formatted(f.IOStreams.Out, opts.outputFormat, namespace); newErr != nil {
		return newErr
	}

	if err = contextutil.SetCurrentNamespaceInstance(&namespace, &conn, f); err != nil {
		return err
	}

	f.Logger.Info(f.Localizer.MustLocalize("connector.namespace.create.info.success", localize.NewEntry("Name", namespace.Name)))

	return nil

}
