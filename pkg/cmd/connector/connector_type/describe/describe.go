package describe

import (
	"github.com/redhat-developer/app-services-cli/pkg/cmd/connector/connectorcmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"

	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	connectorerror "github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1/error"

	"github.com/spf13/cobra"
)

type options struct {
	type_id      string
	outputFormat string

	f *factory.Factory
}

// NewDescribeCommand creates a new command to view a connector type
func NewDescribeCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		f: f,
	}

	cmd := &cobra.Command{
		Use:     "describe",
		Short:   f.Localizer.MustLocalize("connector.type.describe.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("connector.type.describe.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("connector.type.describe.cmd.example"),
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
	flags.StringVar(&opts.type_id, "type", "", f.Localizer.MustLocalize("connector.type.describe.flag.id"))
	flags.AddOutput(&opts.outputFormat)

	_ = cmd.MarkFlagRequired("type")

	_ = cmd.RegisterFlagCompletionFunc("type", func(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return connectorcmdutil.FilterValidTypesArgs(f, toComplete)
	})

	return cmd
}

func runDescribe(opts *options) error {

	if opts.type_id == "" {
		return opts.f.Localizer.MustLocalizeError("connector.type.error.noType")
	}

	conn, err := opts.f.Connection()
	if err != nil {
		return err
	}

	api := conn.API()

	request := api.ConnectorsMgmt().ConnectorTypesApi.GetConnectorTypeByID(opts.f.Context, opts.type_id)

	response, httpRes, err := request.Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	if apiErr := connectorerror.GetAPIError(err); apiErr != nil {
		switch apiErr.GetCode() {
		case connectorerror.ERROR_7:
			return opts.f.Localizer.MustLocalizeError("connector.type.error.notFound", localize.NewEntry("Id", opts.type_id))
		default:
			return err
		}
	}

	if err != nil {
		return err
	}

	if err = dump.Formatted(opts.f.IOStreams.Out, opts.outputFormat, response); err != nil {
		return err
	}

	opts.f.Logger.Info(opts.f.Localizer.MustLocalize("connector.type.describe.info.success"))

	return nil
}
