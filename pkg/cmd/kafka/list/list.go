package list

import (
	"context"
	"fmt"
	"strconv"

	cmdFlagUtil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/icon"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/kafka"
	"github.com/redhat-developer/app-services-cli/pkg/localize"

	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/internal/build"
	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
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
	Logger     logging.Logger
	localizer  localize.Localizer
	Context    context.Context
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
		localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "list",
		Short:   opts.localizer.MustLocalize("kafka.list.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("kafka.list.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("kafka.list.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.outputFormat != "" && !cmdFlagUtil.IsValidInput(opts.outputFormat, cmdFlagUtil.ValidOutputFormats...) {
				return flag.InvalidValueError("output", opts.outputFormat, cmdFlagUtil.ValidOutputFormats...)
			}

			validator := &kafka.Validator{
				Localizer: opts.localizer,
			}

			if err := validator.ValidateSearchInput(opts.search); err != nil {
				return err
			}

			return runList(opts)
		},
	}

	flagSet := flagutil.NewFlagSet(cmd, opts.localizer)
	flagSet.AddOutput(&opts.outputFormat)

	cmd.Flags().IntVar(&opts.page, "page", int(cmdutil.ConvertPageValueToInt32(build.DefaultPageNumber)), opts.localizer.MustLocalize("kafka.list.flag.page"))
	cmd.Flags().IntVar(&opts.limit, "limit", 100, opts.localizer.MustLocalize("kafka.list.flag.limit"))
	cmd.Flags().StringVar(&opts.search, "search", "", opts.localizer.MustLocalize("kafka.list.flag.search"))

	return cmd
}

func runList(opts *options) error {
	conn, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	api := conn.API()

	a := api.Kafka().GetKafkas(opts.Context)
	a = a.Page(strconv.Itoa(opts.page))
	a = a.Size(strconv.Itoa(opts.limit))

	if opts.search != "" {
		query := buildQuery(opts.search)
		opts.Logger.Debug(opts.localizer.MustLocalize("kafka.list.log.debug.filteringKafkaList", localize.NewEntry("Search", query)))
		a = a.Search(query)
	}

	response, _, err := a.Execute()
	if err != nil {
		return err
	}

	if response.Size == 0 && opts.outputFormat == "" {
		opts.Logger.Info(opts.localizer.MustLocalize("kafka.common.log.info.noKafkaInstances"))
		return nil
	}

	switch opts.outputFormat {
	case dump.EmptyFormat:
		var rows []kafkaRow
		serviceConfig, _ := opts.Config.Load()
		if serviceConfig != nil && serviceConfig.Services.Kafka != nil {
			rows = mapResponseItemsToRows(response.GetItems(), serviceConfig.Services.Kafka.ClusterID)
		} else {
			rows = mapResponseItemsToRows(response.GetItems(), "-")
		}
		dump.Table(opts.IO.Out, rows)
		opts.Logger.Info("")
	default:
		return dump.Formatted(opts.IO.Out, opts.outputFormat, response)
	}
	return nil
}

func mapResponseItemsToRows(kafkas []kafkamgmtclient.KafkaRequest, selectedId string) []kafkaRow {
	rows := make([]kafkaRow, len(kafkas))

	for i, k := range kafkas {
		name := k.GetName()
		if k.GetId() == selectedId {
			name = fmt.Sprintf("%s %s", name, icon.Emoji("✔", "(current)"))
		}
		row := kafkaRow{
			ID:            k.GetId(),
			Name:          name,
			Owner:         k.GetOwner(),
			Status:        k.GetStatus(),
			CloudProvider: k.GetCloudProvider(),
			Region:        k.GetRegion(),
		}

		rows[i] = row
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
