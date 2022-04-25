package list

import (
	"context"
	"fmt"
	"strconv"

	kafkaFlagutil "github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/kafkacmdutil"

	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"

	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"

	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/internal/build"
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

	IO             *iostreams.IOStreams
	Connection     factory.ConnectionFunc
	Logger         logging.Logger
	localizer      localize.Localizer
	Context        context.Context
	ServiceContext servicecontext.IContext
}

// NewListCommand creates a new command for listing kafkas.
func NewListCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		page:           0,
		limit:          100,
		search:         "",
		Connection:     f.Connection,
		Logger:         f.Logger,
		IO:             f.IOStreams,
		localizer:      f.Localizer,
		Context:        f.Context,
		ServiceContext: f.ServiceContext,
	}

	cmd := &cobra.Command{
		Use:     "list",
		Short:   opts.localizer.MustLocalize("kafka.list.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("kafka.list.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("kafka.list.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, flagutil.ValidOutputFormats...) {
				return flagutil.InvalidValueError("output", opts.outputFormat, flagutil.ValidOutputFormats...)
			}

			validator := &kafkacmdutil.Validator{
				Localizer: opts.localizer,
			}

			if err := validator.ValidateSearchInput(opts.search); err != nil {
				return err
			}

			return runList(opts)
		},
	}

	flags := kafkaFlagutil.NewFlagSet(cmd, opts.localizer)

	flags.AddOutput(&opts.outputFormat)
	flags.IntVar(&opts.page, "page", int(cmdutil.ConvertPageValueToInt32(build.DefaultPageNumber)), opts.localizer.MustLocalize("kafka.list.flag.page"))
	flags.IntVar(&opts.limit, "limit", 100, opts.localizer.MustLocalize("kafka.list.flag.limit"))
	flags.StringVar(&opts.search, "search", "", opts.localizer.MustLocalize("kafka.list.flag.search"))

	return cmd
}

func runList(opts *options) error {
	conn, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	api := conn.API()

	a := api.KafkaMgmt().GetKafkas(opts.Context)
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
		svcContext, err := opts.ServiceContext.Load()
		if err != nil {
			return err
		}

		currCtx, err := contextutil.GetCurrentContext(svcContext, opts.localizer)
		if err != nil {
			return err
		}

		if currCtx.KafkaID != "" {
			rows = mapResponseItemsToRows(response.GetItems(), currCtx.KafkaID)
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

	for i := range kafkas {
		k := kafkas[i]
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
