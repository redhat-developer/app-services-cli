package listclusters

import (
	"context"
	"fmt"
	clustersmgmtv1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	"github.com/redhat-developer/app-services-cli/internal/build"
	kafkaFlagutil "github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection/api/clustermgmt"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
	"github.com/spf13/cobra"
)

type clusterRow struct {
	ID            string `json:"id" header:"ID"`
	Name          string `json:"name" header:"Name"`
	Status        string `json:"status" header:"Status"`
	CloudProvider string `json:"cloud_provider" header:"Cloud Provider"`
	Region        string `json:"region" header:"Region"`
}

type options struct {
	outputFormat            string
	page                    int
	limit                   int
	search                  string
	kfmClusterList          kafkamgmtclient.EnterpriseClusterList
	selectedCluster         clustersmgmtv1.Cluster
	clustermgmtClusterList  []clustersmgmtv1.Cluster
	registeredClusters      []clusterRow
	accessToken             string
	clusterManagementApiUrl string

	f *factory.Factory
}

const clusterReadyState = "ready"

func NewListClusterCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		f: f,
	}

	cmd := &cobra.Command{
		Use:     "list",
		Short:   f.Localizer.MustLocalize("dedicated.listClusters.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("dedicated.listClusters.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("dedicated.listClusters.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runListClusters(opts, f)
		},
	}

	flags := kafkaFlagutil.NewFlagSet(cmd, f.Localizer)
	flags.IntVar(&opts.page, "page", int(cmdutil.ConvertPageValueToInt32(build.DefaultPageNumber)), f.Localizer.MustLocalize("dedicated.listClusters.flag.page"))
	flags.IntVar(&opts.limit, "size", int(cmdutil.ConvertSizeValueToInt32(build.DefaultPageSize)), f.Localizer.MustLocalize("dedicated.listClusters.flag.limit"))

	return cmd
}

func runListClusters(opts *options, f *factory.Factory) error {
	err := listEnterpriseClusters(opts, f)
	if err != nil {
		return err
	}
	err = getPaginatedClusterList(opts)
	if err != nil {
		return err
	}
	err = mapRegisteredClusters(opts)
	if err != nil {
		return err
	}
	err = displayRegisteredClusters(opts)
	if err != nil {
		return err
	}
	return nil
}

func listEnterpriseClusters(opts *options, f *factory.Factory) error {
	conn, err := f.Connection()
	if err != nil {
		return err
	}
	ctx := context.Background()
	api := conn.API()
	cl := api.KafkaMgmtEnterprise().GetEnterpriseOsdClusters(ctx)
	clist, response, err := cl.Execute()
	if err != nil {
		return err
	}
	opts.kfmClusterList = clist
	f.Logger.Info(response)
	if len(opts.kfmClusterList.Items) == 0 {
		return f.Localizer.MustLocalizeError("dedicated.listClusters.log.info.noClusters")
	}
	return nil
}

func getPaginatedClusterList(opts *options) error {
	cl, err := clustermgmt.GetClusterList(opts.f, opts.accessToken, opts.clusterManagementApiUrl, opts.page, opts.limit)
	if err != nil {
		opts.f.Localizer.MustLocalizeError("dedicated.listClusters.cmd.errorGettingClusterList")
		return err
	}
	opts.clustermgmtClusterList = validateClustermgmtClusters(cl, opts.clustermgmtClusterList)
	return nil
}

func validateClustermgmtClusters(clusters *clustersmgmtv1.ClusterList, cls []clustersmgmtv1.Cluster) []clustersmgmtv1.Cluster {
	for _, cluster := range clusters.Slice() {
		if cluster.State() == clusterReadyState && cluster.MultiAZ() == true {
			cls = append(cls, *cluster)
		}
	}
	return cls
}

func mapRegisteredClusters(opts *options) error {
	for _, cluster := range opts.clustermgmtClusterList {
		// if cluster is registered, add it to the list of registered clusters
		if isRegistered, kfmcluster := isClusterRegistered(&cluster, &opts.kfmClusterList); isRegistered {
			opts.registeredClusters = append(opts.registeredClusters, clusterRow{
				ID:            cluster.ID(),
				Name:          cluster.Name(),
				Status:        kfmcluster.GetStatus(),
				CloudProvider: cluster.CloudProvider().ID(),
				Region:        cluster.Region().ID(),
			})
		}
	}
	if len(opts.registeredClusters) == 0 {
		return opts.f.Localizer.MustLocalizeError("dedicated.listClusters.cmd.errorNoRegisteredClusters")
	}
	return nil
}

// This is needed as KFM doesn't return the cluster
func isClusterRegistered(cluster *clustersmgmtv1.Cluster, kfmClusterList *kafkamgmtclient.EnterpriseClusterList) (bool, *kafkamgmtclient.EnterpriseCluster) {
	for _, kfmCluster := range kfmClusterList.Items {
		if cluster.ID() == kfmCluster.GetId() {
			return true, &kfmCluster
		}
	}
	return false, &kafkamgmtclient.EnterpriseCluster{}
}

func displayRegisteredClusters(opts *options) error {
	if len(opts.registeredClusters) == 0 {
		return opts.f.Localizer.MustLocalizeError("dedicated.listClusters.cmd.errorNoRegisteredClusters")

	}
	for _, cluster := range opts.registeredClusters {
		dump.Formatted(opts.f.IOStreams.Out, "", cluster)
		opts.f.Logger.Info("")

	}

	lmt := localize.NewEntry("CCClusterLimit", opts.limit)
	totalClusters := localize.NewEntry("CCClusterTotal", len(opts.registeredClusters))
	// show limit or total
	if opts.limit > len(opts.registeredClusters) {
		lmt = localize.NewEntry("CCClusterLimit", len(opts.registeredClusters))
	}
	opts.f.Logger.Info(opts.f.Localizer.MustLocalize("dedicated.listClusters.common.limitTotal", lmt, totalClusters))
	start := (opts.page - 1) * opts.limit
	end := start + len(opts.registeredClusters)
	opts.f.Logger.Info(fmt.Sprintf("[%v %v : %v - %v]", opts.f.Localizer.MustLocalize("dedicated.listClusters.common.page"), opts.page, start, end))
	return nil
}
