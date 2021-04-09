package list

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	kasclient "github.com/redhat-developer/app-services-cli/pkg/api/kas/client"
	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/kafka"

	"github.com/redhat-developer/app-services-cli/pkg/dump"

	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/internal/localizer"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	"github.com/redhat-developer/app-services-cli/pkg/logging"

	"gopkg.in/yaml.v2"
)

// row is the details of a Kafka instance needed to print to a table
type kafkaRow struct {
	ID            string `json:"id" header:"ID"`
	Name          string `json:"name" header:"Name"`
	Owner         string `json:"owner" header:"Owner"`
	Status        string `json:"status" header:"Status"`
	CloudProvider string `json:"cloud_provider" header:"Cloud Provider"`
	Region        string `json:"region" header:"Region"`
}

type options struct {
	outputFormat string
	page         int
	limit        int
	search       string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     func() (logging.Logger, error)
}

// NewListCommand creates a new command for listing kafkas.
func NewListCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		page:       0,
		limit:      100,
		search:     "",
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
	}

	cmd := &cobra.Command{
		Use:   localizer.MustLocalizeFromID("kafka.list.cmd.use"),
		Short: localizer.MustLocalizeFromID("kafka.list.cmd.shortDescription"),
		Long:  localizer.MustLocalizeFromID("kafka.list.cmd.longDescription"),
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, flagutil.ValidOutputFormats...) {
				return flag.InvalidValueError("output", opts.outputFormat, flagutil.ValidOutputFormats...)
			}

			if err := kafka.ValidateSearchInput(opts.search); err != nil {
				return err
			}

			return runList(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "", localizer.MustLocalize(&localizer.Config{
		MessageID:   "kafka.common.flag.output.description",
		PluralCount: 2,
	}))
	cmd.Flags().IntVarP(&opts.page, "page", "", 0, localizer.MustLocalizeFromID("kafka.list.flag.page"))
	cmd.Flags().IntVarP(&opts.limit, "limit", "", 100, localizer.MustLocalizeFromID("kafka.list.flag.limit"))
	cmd.Flags().StringVarP(&opts.search, "search", "", "", localizer.MustLocalizeFromID("kafka.list.flag.search"))

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

	a := api.Kafka().ListKafkas(context.Background())
	a = a.Page(strconv.Itoa(opts.page))
	a = a.Size(strconv.Itoa(opts.limit))

	if opts.search != "" {

		logger.Debug(localizer.MustLocalize(&localizer.Config{
			MessageID: "kafka.list.log.debug.filteringKafkaList",
			TemplateData: map[string]interface{}{
				"Search": buildQuery(opts.search),
			},
		}))
		a = a.Search(buildQuery(opts.search))
	}

	response, _, apiErr := a.Execute()

	if apiErr.Error() != "" {
		return apiErr
	}

	if response.Size == 0 {
		logger.Info(localizer.MustLocalizeFromID("kafka.common.log.info.noKafkaInstances"))
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
		rows := mapResponseItemsToRows(response.GetItems())
		dump.Table(opts.IO.Out, rows)
		logger.Info("")
	}

	return nil
}

func mapResponseItemsToRows(kafkas []kasclient.KafkaRequest) []kafkaRow {
	rows := []kafkaRow{}

	for _, k := range kafkas {
		row := kafkaRow{
			ID:            k.GetId(),
			Name:          k.GetName(),
			Owner:         k.GetOwner(),
			Status:        k.GetStatus(),
			CloudProvider: k.GetCloudProvider(),
			Region:        k.GetRegion(),
		}

		rows = append(rows, row)
	}

	return rows
}

func buildQuery(search string) string {

	queryString := fmt.Sprintf(
		"name like %%%[1]v%% or owner like %%%[1]v%% or cloud_provider like %%%[1]v%% or region like %%%[1]v%% or status like %%%[1]v%%",
		search,
	)

	return queryString

}
