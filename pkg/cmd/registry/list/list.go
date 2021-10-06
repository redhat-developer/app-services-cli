package list

import (
	"context"
	"fmt"
	"github.com/redhat-developer/app-services-cli/pkg/icon"
	srsmgmtv1 "github.com/redhat-developer/app-services-sdk-go/registrymgmt/apiv1/client"

	"github.com/redhat-developer/app-services-cli/pkg/cmdutil"
	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/localize"

	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/internal/build"
	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
)

// RegistryRow is the details of a Service Registry instance needed to print to a table
type RegistryRow struct {
	ID     string `json:"id" header:"ID"`
	Name   string `json:"name" header:"Name"`
	Owner  string `json:"owner" header:"Owner"`
	Status string `json:"status" header:"Status"`
}

type options struct {
	outputFormat string
	page         int32
	limit        int32
	search       string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	localizer  localize.Localizer
	Context    context.Context
}

// NewListCommand creates a new command for listing service registries.
func NewListCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "list",
		Short:   f.Localizer.MustLocalize("registry.cmd.list.shortDescription"),
		Long:    f.Localizer.MustLocalize("registry.cmd.list.longDescription"),
		Example: f.Localizer.MustLocalize("registry.cmd.list.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, flagutil.ValidOutputFormats...) {
				return flag.InvalidValueError("output", opts.outputFormat, flagutil.ValidOutputFormats...)
			}
			if opts.page < 1 {
				return opts.localizer.MustLocalizeError("common.validation.page.error.invalid.minValue", localize.NewEntry("Page", opts.page))
			}

			if opts.limit < 1 {
				return opts.localizer.MustLocalizeError("common.validation.limit.error.invalid.minValue", localize.NewEntry("Limit", opts.limit))
			}

			return runList(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "", opts.localizer.MustLocalize("registry.cmd.flag.output.description"))
	cmd.Flags().Int32VarP(&opts.page, "page", "", cmdutil.ConvertPageValueToInt32(build.DefaultPageNumber), opts.localizer.MustLocalize("registry.list.flag.page"))
	cmd.Flags().Int32VarP(&opts.limit, "limit", "", 100, opts.localizer.MustLocalize("registry.list.flag.limit"))

	flagutil.EnableOutputFlagCompletion(cmd)

	return cmd
}

func runList(opts *options) error {
	conn, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	api := conn.API()

	a := api.ServiceRegistryMgmt().GetRegistries(opts.Context)
	a = a.Page(opts.page)
	a = a.Size(opts.limit)

	if opts.search != "" {
		query := buildQuery(opts.search)
		opts.Logger.Debug("Filtering Service Registries with query", query)
		a = a.Search(query)
	}

	response, _, err := a.Execute()
	if err != nil {
		return err
	}

	if len(response.Items) == 0 && opts.outputFormat == "" {
		opts.Logger.Info(opts.localizer.MustLocalize("registry.common.log.info.noInstances"))
		return nil
	}

	switch opts.outputFormat {
	case dump.EmptyFormat:
		var rows []RegistryRow
		serviceConfig, _ := opts.Config.Load()
		if serviceConfig != nil && serviceConfig.Services.ServiceRegistry != nil {
			rows = mapResponseItemsToRows(&response.Items, serviceConfig.Services.ServiceRegistry.InstanceID)
		} else {
			rows = mapResponseItemsToRows(&response.Items, "-")
		}
		dump.Table(opts.IO.Out, rows)
		opts.Logger.Info("")
	default:
		return dump.Formatted(opts.IO.Out, opts.outputFormat, response)
	}

	return nil
}

func mapResponseItemsToRows(registries *[]srsmgmtv1.Registry, selectedId string) []RegistryRow {
	rows := make([]RegistryRow, len(*registries))

	for i := range *registries {
		k := (*registries)[i]
		name := k.GetName()
		if k.Id == selectedId {
			name = fmt.Sprintf("%s %s", name, icon.Emoji("✔", "(current)"))
		}
		row := RegistryRow{
			ID:     fmt.Sprint(k.Id),
			Name:   name,
			Status: string(k.GetStatus()),
			Owner:  k.GetOwner(),
		}

		rows[i] = row
	}

	return rows
}

func buildQuery(search string) string {
	queryString := fmt.Sprintf(
		"name=%v",
		search,
	)

	return queryString
}
