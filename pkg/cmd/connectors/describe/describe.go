package describe

import (
	"github.com/redhat-developer/app-services-cli/pkg/cmd/connectors/connectorcmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"

	"github.com/spf13/cobra"
)

type options struct {
	id           string
	outputFormat string

	f *factory.Factory
}

func NewDescribeCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		f: f,
	}

	cmd := &cobra.Command{
		Use:     "describe",
		Short:   f.Localizer.MustLocalize("connector.describe.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("connector.describe.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("connector.describe.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {

			validOutputFormats := flagutil.ValidOutputFormats
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flagutil.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			return runDescribe(opts)
		},
	}
	flags := connectorcmdutil.NewFlagSet(cmd, f)
	_ = flags.AddConnectorID(&opts.id).Required()
	flags.AddOutput(&opts.outputFormat)

	return cmd
}

func runDescribe(opts *options) error {
	f := opts.f

	var conn connection.Connection
	conn, err := f.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	api := conn.API()

	a := api.ConnectorsMgmt().ConnectorsApi.GetConnector(f.Context, opts.id)

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

	f.Logger.Info(f.Localizer.MustLocalize("connectors.describe.info.success"))

	return nil
}
