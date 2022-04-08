package delete

import (
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	connectormgmtclient "github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1/client"

	"github.com/spf13/cobra"
)

type options struct {
	id           string
	outputFormat string

	f           *factory.Factory
	interactive bool
}

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
			if opts.id != "" {
				// TODO
			}

			if !f.IOStreams.CanPrompt() && opts.id == "" {
				return f.Localizer.MustLocalizeError("connector.common.error.requiredWhenNonInteractive")
			} else if opts.id == "" {
				opts.interactive = true
			}

			validOutputFormats := flagutil.ValidOutputFormats
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flagutil.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			return runDelete(opts)
		},
	}
	flags := flagutil.NewFlagSet(cmd, f.Localizer)
	flags.StringVar(&opts.id, "name", "", f.Localizer.MustLocalize("connector.cluster.create.flag.name.description"))
	flags.AddOutput(&opts.outputFormat)

	return cmd
}

func runDelete(opts *options) error {
	f := opts.f

	var conn connection.Connection
	conn, err := f.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	if opts.interactive {
		// TODO
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

	f.Logger.Info(f.Localizer.MustLocalize("connectors.cluster.delete.info.success"))

	return nil
}
