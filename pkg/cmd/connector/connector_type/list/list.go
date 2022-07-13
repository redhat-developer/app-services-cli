package list

import (
	"strconv"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/connector/connectorcmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	connectormgmtclient "github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1/client"
	"github.com/spf13/cobra"
)

type options struct {
	search string
	limt int
	page int
	outputFormat string

	f            *factory.Factory
}

type connectorType struct {
	Id string `json:"id,omitempty"`
	Kind string `json:"kind,omitempty"`
	Href string `json:"href,omitempty"`
}

// NewListCommand creates a new command to list connector types
func NewListCommand(f *factory.Factory) *cobra.Command {

	opts := &options{
		f: f,
	}

	cmd := &cobra.Command{
		Use:     "list",
		Short:   f.Localizer.MustLocalize("connector.start.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("connector.start.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("connector.start.cmd.example"),
		RunE: func(cmd *cobra.Command, args []string) error {

			validOutputFormats := flagutil.ValidOutputFormats
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flagutil.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			return runUpdateCommand(opts)
		},
	}

	flags := connectorcmdutil.NewFlagSet(cmd, f)
	flags.AddOutput(&opts.outputFormat)
	flags.StringVar(&opts.search, "search", "", "search description")
	flags.IntVar(&opts.limt, "limit", 20, "limit description")
	flags.IntVar(&opts.page, "page", 1, "page description")
	return cmd

}

func runUpdateCommand(opts *options) error {
	f := opts.f

	var conn connection.Connection
	conn, err := f.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	api := conn.API()

	types, httpRes, err := api.ConnectorsMgmt().ConnectorTypesApi.GetConnectorTypes(f.Context).Page(strconv.Itoa(opts.page)).Size("100").Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	f.Logger.Info(len(types.Items))
	if err != nil {
		return err
	}

	rows := mapResponseToConnectorTypes(&types)
	for i := 0; i < len(rows); i++ {
		if err = dump.Formatted(f.IOStreams.Out, opts.outputFormat, rows[i]); err != nil {
			return err
		}
	}

	return nil
}

func mapResponseToConnectorTypes(list *connectormgmtclient.ConnectorTypeList) []connectorType {
	types := make([]connectorType, len(list.Items))

	for i := 0; i < len(list.Items); i++ {
		item := &list.Items[i]
		types[i] = connectorType{
			Id: *item.ObjectReference.Id,
			Href: *item.ObjectReference.Href,
			Kind: *item.ObjectReference.Kind,
		}
	}

	return types
}