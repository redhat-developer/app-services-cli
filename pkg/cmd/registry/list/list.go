package list

import (
	"context"
	"encoding/json"
	"fmt"

	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	srsmgmtv1 "github.com/redhat-developer/app-services-sdk-go/registrymgmt/apiv1/client"

	"github.com/redhat-developer/app-services-cli/pkg/dump"

	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	"github.com/redhat-developer/app-services-cli/pkg/logging"

	"gopkg.in/yaml.v2"
)

// row is the details of a Service Registry instance needed to print to a table
type RegistryRow struct {
	ID     string `json:"id" header:"ID"`
	Name   string `json:"name" header:"Name"`
	URL    string `json:"registryUrl" header:"Registry URL"`
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
	Logger     func() (logging.Logger, error)
	localizer  localize.Localizer
}

// NewListCommand creates a new command for listing service registries.
func NewListCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
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

			return runList(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "", opts.localizer.MustLocalize("registry.cmd.flag.output.description"))
	cmd.Flags().Int32VarP(&opts.page, "page", "", 0, opts.localizer.MustLocalize("registry.list.flag.page"))
	cmd.Flags().Int32VarP(&opts.limit, "limit", "", 100, opts.localizer.MustLocalize("registry.list.flag.limit"))
	cmd.Flags().StringVarP(&opts.search, "search", "", "", opts.localizer.MustLocalize("registry.list.flag.search"))

	return cmd
}

func runList(opts *options) error {
	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	connection, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	api := connection.API()

	a := api.ServiceRegistryMgmt().GetRegistries(context.Background())
	a = a.Page(opts.page)
	a = a.Size(opts.limit)

	if opts.search != "" {
		query := buildQuery(opts.search)
		logger.Debug("Filtering Service Registries with query", query)
		a = a.Search(query)
	}

	response, _, err := a.Execute()

	if err != nil {
		return err
	}

	if len(response.Items) == 0 && opts.outputFormat == "" {
		logger.Info(opts.localizer.MustLocalize("registry.common.log.info.noInstances"))
		return nil
	}

	switch opts.outputFormat {
	case "json":
		data, _ := json.Marshal(response)
		_ = dump.JSON(opts.IO.Out, data)
	case "yaml", "yml":
		data, _ := yaml.Marshal(response)
		_ = dump.YAML(opts.IO.Out, data)
	default:
		rows := mapResponseItemsToRows(&response.Items)
		dump.Table(opts.IO.Out, rows)
		logger.Info("")
	}

	return nil
}

func mapResponseItemsToRows(registries *[]srsmgmtv1.RegistryRest) []RegistryRow {
	rows := []RegistryRow{}

	for i := range *registries {
		k := (*registries)[i]
		row := RegistryRow{
			ID:     fmt.Sprint(k.Id),
			Name:   k.GetName(),
			URL:    k.GetRegistryUrl(),
			Status: string(k.GetStatus()),
		}

		rows = append(rows, row)
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
