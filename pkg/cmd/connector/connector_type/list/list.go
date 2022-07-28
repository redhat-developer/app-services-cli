package list

import (
	"strconv"

	"github.com/redhat-developer/app-services-cli/internal/build"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/connector/connectorcmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"

	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	connectormgmtclient "github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1/client"
	connectorerror "github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1/error"
	"github.com/spf13/cobra"
)

const (
	DefaultSearch = ""
)

type options struct {
	search       string
	limit        int
	page         int
	outputFormat string

	f *factory.Factory
}

type connectorOutput struct {
	Name        string `json:"name,omitempty"`
	Id          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
}

// NewListCommand creates a new command to list connector types
func NewListCommand(f *factory.Factory) *cobra.Command {

	opts := &options{
		f: f,
	}

	cmd := &cobra.Command{
		Use:     "list",
		Short:   f.Localizer.MustLocalize("connector.type.list.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("connector.type.list.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("connector.type.list.cmd.example"),
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
	flags.StringVar(&opts.search, "search", DefaultSearch, f.Localizer.MustLocalize("connector.type.list.flag.search.description"))
	flags.IntVar(&opts.limit, "limit", int(cmdutil.ConvertPageValueToInt32(build.DefaultPageSize)), f.Localizer.MustLocalize("connector.type.list.flag.page.description"))
	flags.IntVar(&opts.page, "page", int(cmdutil.ConvertPageValueToInt32(build.DefaultPageNumber)), f.Localizer.MustLocalize("connector.type.list.flag.page.description"))
	return cmd

}

func runUpdateCommand(opts *options) error {
	f := opts.f

	var conn connection.Connection
	conn, err := f.Connection()
	if err != nil {
		return err
	}

	api := conn.API()

	request := api.ConnectorsMgmt().ConnectorTypesApi.GetConnectorTypes(f.Context)
	request = request.Page(strconv.Itoa(opts.page))
	request = request.Size(strconv.Itoa(opts.limit))

	if opts.search != DefaultSearch {
		query := connectorcmdutil.NewSearchQuery(opts.search).Filter("name").Filter("description").Build()
		request = request.Search(query)
	}

	types, httpRes, err := request.Execute()

	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	if apiErr := connectorerror.GetAPIError(err); apiErr != nil {
		switch apiErr.GetCode() {
		case connectorerror.ERROR_11:
			return opts.f.Localizer.MustLocalizeError("connector.common.error.unauthorized")
		case connectorerror.ERROR_23:
			return opts.f.Localizer.MustLocalizeError("connector.common.error.parse.search")

		default:
			return err
		}
	}

	rows := mapResponseToConnectorTypes(&types)
	for i := 0; i < len(rows); i++ {
		if err = dump.Formatted(f.IOStreams.Out, opts.outputFormat, rows[i]); err != nil {
			return err
		}
	}

	return nil
}

func mapResponseToConnectorTypes(list *connectormgmtclient.ConnectorTypeList) []connectorOutput {
	types := make([]connectorOutput, len(list.Items))

	for i := 0; i < len(list.Items); i++ {
		item := &list.Items[i]
		types[i] = connectorOutput{
			Name:        item.GetName(),
			Id:          item.GetId(),
			Description: item.GetDescription(),
		}
	}

	return types
}
