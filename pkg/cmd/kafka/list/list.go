package list

import (
	"fmt"
	v1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	"strconv"

	kafkaFlagutil "github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/kafkacmdutil"

	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"

	clustermgmt "github.com/redhat-developer/app-services-cli/pkg/shared/connection/api/clustermgmt"

	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/kafkamgmt/apiv1/client"

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
	CustomerCloud string `json:"customer_cloud" header:"Customer Cloud"`
}

type options struct {
	outputFormat            string
	page                    int
	limit                   int
	search                  string
	accessToken             string
	clusterManagementApiUrl string

	f *factory.Factory
}

// NewListCommand creates a new command for listing kafkas.
func NewListCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		page:   0,
		limit:  100,
		search: "",
		f:      f,
	}

	cmd := &cobra.Command{
		Use:     "list",
		Short:   opts.f.Localizer.MustLocalize("kafka.list.cmd.shortDescription"),
		Long:    opts.f.Localizer.MustLocalize("kafka.list.cmd.longDescription"),
		Example: opts.f.Localizer.MustLocalize("kafka.list.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, flagutil.ValidOutputFormats...) {
				return flagutil.InvalidValueError("output", opts.outputFormat, flagutil.ValidOutputFormats...)
			}

			validator := &kafkacmdutil.Validator{
				Localizer: opts.f.Localizer,
			}

			if err := validator.ValidateSearchInput(opts.search); err != nil {
				return err
			}

			return runList(opts)
		},
	}

	flags := kafkaFlagutil.NewFlagSet(cmd, opts.f.Localizer)

	flags.AddOutput(&opts.outputFormat)
	flags.IntVar(&opts.page, "page", int(cmdutil.ConvertPageValueToInt32(build.DefaultPageNumber)), opts.f.Localizer.MustLocalize("kafka.list.flag.page"))
	flags.IntVar(&opts.limit, "limit", 100, opts.f.Localizer.MustLocalize("kafka.list.flag.limit"))
	flags.StringVar(&opts.search, "search", "", opts.f.Localizer.MustLocalize("kafka.list.flag.search"))
	flags.StringVar(&opts.clusterManagementApiUrl, "cluster-mgmt-api-url", "", f.Localizer.MustLocalize("dedicated.registerCluster.flag.clusterMgmtApiUrl.description"))
	flags.StringVar(&opts.accessToken, "access-token", "", f.Localizer.MustLocalize("dedicated.registercluster.flag.accessToken.description"))

	_ = flags.MarkHidden("cluster-mgmt-api-url")
	_ = flags.MarkHidden("access-token")

	return cmd
}

func runList(opts *options) error {
	conn, err := opts.f.Connection()
	if err != nil {
		return err
	}

	api := conn.API()

	a := api.KafkaMgmt().GetKafkas(opts.f.Context)
	a = a.Page(strconv.Itoa(opts.page))
	a = a.Size(strconv.Itoa(opts.limit))

	if opts.search != "" {
		query := buildQuery(opts.search)
		opts.f.Logger.Debug(opts.f.Localizer.MustLocalize("kafka.list.log.debug.filteringKafkaList", localize.NewEntry("Search", query)))
		a = a.Search(query)
	}

	response, _, err := a.Execute()
	if err != nil {
		return err
	}

	if response.Size == 0 && opts.outputFormat == "" {
		opts.f.Logger.Info(opts.f.Localizer.MustLocalize("kafka.common.log.info.noKafkaInstances"))
		return nil
	}

	clusterIdMap, err := getClusterIdMapFromKafkas(opts, response.GetItems())
	if err != nil {
		return err
	}

	switch opts.outputFormat {
	case dump.EmptyFormat:
		var rows []kafkaRow
		svcContext, err := opts.f.ServiceContext.Load()
		if err != nil {
			return err
		}

		currCtx, err := contextutil.GetCurrentContext(svcContext, opts.f.Localizer)
		if err != nil {
			return err
		}

		if currCtx.KafkaID != "" {
			rows = mapResponseItemsToRows(response.GetItems(), currCtx.KafkaID, &clusterIdMap)
		} else {
			rows = mapResponseItemsToRows(response.GetItems(), "-", &clusterIdMap)
		}
		dump.Table(opts.f.IOStreams.Out, rows)
		opts.f.Logger.Info("")
	default:
		return dump.Formatted(opts.f.IOStreams.Out, opts.outputFormat, response)
	}
	return nil
}

func mapResponseItemsToRows(kafkas []kafkamgmtclient.KafkaRequest, selectedId string, clusterIdMap *map[string]*v1.Cluster) []kafkaRow {
	rows := make([]kafkaRow, len(kafkas))

	for i := range kafkas {
		k := kafkas[i]
		name := k.GetName()
		if k.GetId() == selectedId {
			name = fmt.Sprintf("%s %s", name, icon.Emoji("✔", "(current)"))
		}

		var customerCloud string
		if id, ok := k.GetClusterIdOk(); ok {
			cluster := (*clusterIdMap)[*id]
			customerCloud = fmt.Sprintf("%v (%v)", cluster.Name(), cluster.ID())
		} else {
			customerCloud = "Red Hat Infrastructure"
		}

		row := kafkaRow{
			ID:            k.GetId(),
			Name:          name,
			Owner:         k.GetOwner(),
			Status:        k.GetStatus(),
			CloudProvider: k.GetCloudProvider(),
			Region:        k.GetRegion(),
			CustomerCloud: customerCloud,
		}

		rows[i] = row
	}

	return rows
}

func getClusterIdMapFromKafkas(opts *options, kafkas []kafkamgmtclient.KafkaRequest) (map[string]*v1.Cluster, error) {
	// map[string]struct{} is used remove duplicated ids from being added to the request
	kafkaClusterIds := make(map[string]struct{}, len(kafkas))
	for _, kafka := range kafkas {
		if kafka.GetClusterId() != "" {
			kafkaClusterIds[kafka.GetClusterId()] = struct{}{}
		}
	}

	idToCluster := make(map[string]*v1.Cluster)

	// if no kafkas have a cluster id assigned then we can skip the call to get
	// the clusters as we dont need their info
	if len(kafkaClusterIds) > 0 {
		clusterList, err := clustermgmt.GetClusterListByIds(opts.f, opts.clusterManagementApiUrl, opts.accessToken, createSearchString(&kafkaClusterIds), len(kafkaClusterIds))
		if err != nil {
			return nil, err
		}

		for _, cluster := range clusterList.Slice() {
			idToCluster[cluster.ID()] = cluster
		}
	}

	return idToCluster, nil
}

func createSearchString(idSet *map[string]struct{}) string {
	searchString := ""
	index := 0
	for id := range *idSet {
		if index > 0 {
			searchString += " or "
		}
		searchString += fmt.Sprintf("id = '%s'", id)
		index += 1
	}
	return searchString
}

func buildQuery(search string) string {
	queryString := fmt.Sprintf(
		"name like %%%[1]v%% or owner like %%%[1]v%% or cloud_provider like %%%[1]v%% or region like %%%[1]v%% or status like %%%[1]v%%",
		search,
	)

	return queryString
}
