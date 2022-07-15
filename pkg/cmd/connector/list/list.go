package list

import (
	"strconv"

	connectormgmtclient "github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1/client"

	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"

	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/internal/build"
)

// row is the details of a Kafka instance needed to print to a table
type itemRow struct {
	ID     string `json:"id" header:"ID"`
	Name   string `json:"name" header:"Name"`
	Owner  string `json:"owner" header:"Owner"`
	Status string `json:"status" header:"Status"`
}

type options struct {
	outputFormat string
	page         int
	limit        int

	f *factory.Factory
}

// NewListCommand creates a new command to list connectors
func NewListCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		page:  0,
		limit: 100,
		f:     f,
	}

	cmd := &cobra.Command{
		Use:     "list",
		Short:   f.Localizer.MustLocalize("connector.list.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("connector.list.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("connector.list.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, flagutil.ValidOutputFormats...) {
				return flagutil.InvalidValueError("output", opts.outputFormat, flagutil.ValidOutputFormats...)
			}

			return runList(opts)
		},
	}

	flags := flagutil.NewFlagSet(cmd, f.Localizer)

	flags.AddOutput(&opts.outputFormat)
	flags.IntVar(&opts.page, "page", int(cmdutil.ConvertPageValueToInt32(build.DefaultPageNumber)), f.Localizer.MustLocalize("connector.common.list.flag.page"))
	flags.IntVar(&opts.limit, "limit", int(cmdutil.ConvertPageValueToInt32(build.DefaultPageSize)), f.Localizer.MustLocalize("connector.common.list.flag.limit"))

	return cmd
}

func runList(opts *options) error {
	f := opts.f
	conn, err := f.Connection()
	if err != nil {
		return err
	}

	api := conn.API()

	a := api.ConnectorsMgmt().ConnectorsApi.ListConnectors(f.Context)
	a = a.Page(strconv.Itoa(opts.page))
	a = a.Size(strconv.Itoa(opts.limit))

	response, _, err := a.Execute()
	if err != nil {
		return err
	}

	if response.Items == nil && opts.outputFormat == "" {
		f.Logger.Info(f.Localizer.MustLocalize("connector.common.log.info.noResults"))
		return nil
	}

	if response.Size == 0  && opts.outputFormat == "" {
		f.Logger.Info(f.Localizer.MustLocalize("connector.common.log.info.noResults"))
		return nil
	}

	switch opts.outputFormat {
	case dump.EmptyFormat:
		rows := mapResponseItemsToRows(response.Items)

		dump.Table(f.IOStreams.Out, rows)
		f.Logger.Info("")
	default:
		return dump.Formatted(f.IOStreams.Out, opts.outputFormat, response)
	}
	return nil
}

func mapResponseItemsToRows(items []connectormgmtclient.Connector) []itemRow {
	rows := make([]itemRow, len(items))

	for i := range items {
		k := items[i]
		name := k.GetName()

		row := itemRow{
			ID:     k.GetId(),
			Name:   name,
			Owner:  k.GetOwner(),
			Status: string(*k.Status.State),
		}

		rows[i] = row
	}

	return rows
}
