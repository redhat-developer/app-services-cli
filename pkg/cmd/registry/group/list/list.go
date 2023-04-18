package list

import (
	"github.com/redhat-developer/app-services-cli/internal/build"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/registrycmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/rule/rulecmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	registryinstanceclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal/client"
	"github.com/spf13/cobra"
	"time"

	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
)

// groupRow is the details of a group needed to print to a table
type groupRow struct {
	Id string `json:"id" header:"Id"`

	Description string `json:"description" header:"Description"`

	CreatedOn time.Time `json:"createdOn" header:"Created on"`

	CreatedBy string `json:"createdBy" header:"Created by"`

	ModifiedOn time.Time `json:"modifiedOn" header:"Modified on"`

	ModifiedBy string `json:"modifiedBy" header:"Modified by"`
}

type options struct {
	registryID string

	outputFormat string
	page         int32
	limit        int32
	orderBy      string
	descending   bool

	f *factory.Factory
}

// NewListCommand creates a new command to view a list of groups
func NewListCommand(f *factory.Factory) *cobra.Command {

	opts := &options{
		f: f,
	}

	cmd := &cobra.Command{
		Use:     "list",
		Short:   f.Localizer.MustLocalize("group.list.cmd.description.short"),
		Long:    f.Localizer.MustLocalize("group.list.cmd.description.long"),
		Example: f.Localizer.MustLocalize("group.list.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) (err error) {

			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, flagutil.ValidOutputFormats...) {
				return flagutil.InvalidValueError("output", opts.outputFormat, flagutil.ValidOutputFormats...)
			}
			if opts.page < 1 {
				return opts.f.Localizer.MustLocalizeError("common.validation.page.error.invalid.minValue", localize.NewEntry("Page", opts.page))
			}

			if opts.limit < 1 {
				return opts.f.Localizer.MustLocalizeError("common.validation.limit.error.invalid.minValue", localize.NewEntry("Limit", opts.limit))
			}

			if opts.registryID != "" {
				return runList(opts)
			}

			registryInstance, err := contextutil.GetCurrentRegistryInstance(f)
			if err != nil {
				return err
			}

			opts.registryID = registryInstance.GetId()

			return runList(opts)
		},
	}

	flags := rulecmdutil.NewFlagSet(cmd, f)

	flags.Int32VarP(&opts.page, "page", "", cmdutil.ConvertPageValueToInt32(build.DefaultPageNumber), opts.f.Localizer.MustLocalize("group.cmd.list.flag.page"))
	flags.Int32VarP(&opts.limit, "limit", "", 100, opts.f.Localizer.MustLocalize("group.cmd.list.flag.limit"))
	flags.BoolVar(&opts.descending, "descending", false, opts.f.Localizer.MustLocalize("group.cmd.list.flag.descending"))
	flags.StringVar(&opts.orderBy, "order-by", "name", opts.f.Localizer.MustLocalize("group.cmd.list.flag.orderby"))

	flags.AddOutput(&opts.outputFormat)

	flags.AddRegistryInstance(&opts.registryID)

	return cmd

}

func runList(opts *options) error {
	conn, err := opts.f.Connection()
	if err != nil {
		return err
	}

	api := conn.API()

	a, _, err := api.ServiceRegistryInstance(opts.registryID)
	if err != nil {
		return err
	}
	request := a.GroupsApi.ListGroups(opts.f.Context)

	request = request.Limit(opts.limit)
	request = request.Offset((opts.page - 1) * opts.limit)

	if opts.descending {
		request = request.Order(registryinstanceclient.SORTORDER_DESC)
	} else {
		request = request.Order(registryinstanceclient.SORTORDER_ASC)
	}

	sortBy, err := registryinstanceclient.NewSortByFromValue(opts.orderBy)
	if err != nil {
		return err
	}
	request = request.Orderby(*sortBy)

	response, _, err := request.Execute()
	if err != nil {
		return registrycmdutil.TransformInstanceError(err)
	}

	if opts.outputFormat == dump.EmptyFormat {
		rows := mapResponseItemsToRows(response)
		opts.f.Logger.Info("")
		dump.Table(opts.f.IOStreams.Out, rows)
		opts.f.Logger.Info("")
	} else {
		return dump.Formatted(opts.f.IOStreams.Out, opts.outputFormat, response)
	}

	return nil
}

func mapResponseItemsToRows(groupResult registryinstanceclient.GroupSearchResults) []groupRow {
	rows := make([]groupRow, groupResult.Count)

	for i, k := range groupResult.Groups {
		row := groupRow{
			Id:          k.GetId(),
			Description: k.GetDescription(),
			CreatedOn:   k.GetCreatedOn(),
			CreatedBy:   k.GetCreatedBy(),
			ModifiedOn:  k.GetModifiedOn(),
			ModifiedBy:  k.GetModifiedBy(),
		}

		rows[i] = row
	}

	return rows
}
