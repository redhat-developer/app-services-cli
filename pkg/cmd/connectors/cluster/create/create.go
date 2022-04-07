package create

import (
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	connectormgmtclient "github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1/client"

	"github.com/spf13/cobra"
)

type options struct {
	name string

	outputFormat string
	f            *factory.Factory
	interactive  bool
}

func NewCreateCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		f: f,
	}

	cmd := &cobra.Command{
		Use:     "create",
		Short:   f.Localizer.MustLocalize("connector.cluster.create.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("connector.cluster.create.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("connector.cluster.create.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.name != "" {
				// TODO
			}

			if !f.IOStreams.CanPrompt() && opts.name == "" {
				return f.Localizer.MustLocalizeError("connector.cluster.create.argument.name.error.requiredWhenNonInteractive")
			} else if opts.name == "" {
				opts.interactive = true
			}

			validOutputFormats := flagutil.ValidOutputFormats
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flagutil.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			return runCreate(opts)
		},
	}
	flags := flagutil.NewFlagSet(cmd, f.Localizer)
	flags.StringVar(&opts.name, "name", "", f.Localizer.MustLocalize("connector.cluster.name.flag.id.description"))
	flags.AddOutput(&opts.outputFormat)

	return cmd
}

func runCreate(opts *options) error {
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
		Name: &opts.name,
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

	f.Logger.Info(f.Localizer.MustLocalize("connectors.cluster.create.info.success"))

	return nil
}