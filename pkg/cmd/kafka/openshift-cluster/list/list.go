package list

import (
	"fmt"
	clustersmgmtv1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	"github.com/redhat-developer/app-services-cli/internal/build"
	kafkaFlagutil "github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/openshift-cluster/openshiftclustercmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection/api/clustermgmt"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/redhat-developer/app-services-cli/pkg/shared/kafkautil"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/kafkamgmt/apiv1/client"
	"github.com/spf13/cobra"
)

type clusterRow struct {
	Name          string `json:"name" header:"Name"`
	ID            string `json:"id" header:"ID"`
	Status        string `json:"status" header:"Status"`
	CloudProvider string `json:"cloud_provider" header:"Cloud Provider"`
	Region        string `json:"region" header:"Region"`
}

type options struct {
	kfmClusterList          *kafkamgmtclient.EnterpriseClusterList
	clustermgmtClusterList  *clustersmgmtv1.ClusterList
	registeredClusters      []clusterRow
	accessToken             string
	clusterManagementApiUrl string

	f *factory.Factory
}

func NewListClusterCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		f: f,
	}

	cmd := &cobra.Command{
		Use:     "list",
		Short:   f.Localizer.MustLocalize("kafka.openshiftCluster.list.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("kafka.openshiftCluster.list.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("kafka.openshiftCluster.list.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runListClusters(opts, f)
		},
	}

	flags := kafkaFlagutil.NewFlagSet(cmd, f.Localizer)

	flags.StringVar(&opts.clusterManagementApiUrl, "cluster-mgmt-api-url", "", f.Localizer.MustLocalize("kafka.openshiftCluster.registerCluster.flag.clusterMgmtApiUrl.description"))
	flags.StringVar(&opts.accessToken, "access-token", "", f.Localizer.MustLocalize("kafka.openshiftCluster.registercluster.flag.accessToken.description"))

	openshiftclustercmdutil.HideClusterMgmtFlags(flags)

	return cmd

}

func runListClusters(opts *options, f *factory.Factory) error {
	kfmClusterList, response, err := kafkautil.ListEnterpriseClusters(f)
	if err != nil {

		if response != nil {
			if response.StatusCode == 403 {
				return opts.f.Localizer.MustLocalizeError("kafka.openshiftCluster.list.error.permissionDenied")
			}

			return fmt.Errorf("%v, %w", response.Status, err)
		}

		return err
	}

	opts.kfmClusterList = kfmClusterList

	clist, err := clustermgmt.GetClusterListWithSearchParams(opts.f, opts.clusterManagementApiUrl, opts.accessToken, kafkautil.CreateClusterSearchStringFromKafkaList(opts.kfmClusterList), int(cmdutil.ConvertPageValueToInt32(build.DefaultPageNumber)), len(opts.kfmClusterList.Items))
	if err != nil {
		return err
	}
	if clist == nil {
		return opts.f.Localizer.MustLocalizeError("kafka.openshiftCluster.list.error.noRegisteredClustersFound")
	}
	opts.clustermgmtClusterList = clist
	opts.registeredClusters = kfmListToClusterRowList(opts)
	displayRegisteredClusters(opts)
	return nil
}

func kfmListToClusterRowList(opts *options) []clusterRow {
	var crl []clusterRow

	clusterMap := make(map[string]*clustersmgmtv1.Cluster, len(opts.clustermgmtClusterList.Slice()))
	// create a map of cluster ids to cluster objects
	for _, ocmCluster := range opts.clustermgmtClusterList.Slice() {
		clusterMap[ocmCluster.ID()] = ocmCluster
	}
	for _, kfmcluster := range opts.kfmClusterList.Items {
		ocmCluster := clusterMap[kfmcluster.Id]
		crl = append(crl, clusterRow{
			Name:          ocmCluster.Name(),
			ID:            kfmcluster.Id,
			Status:        kafkautil.MapClusterStatus(kfmcluster.GetStatus()),
			CloudProvider: ocmCluster.CloudProvider().ID(),
			Region:        ocmCluster.Region().ID(),
		})
	}
	return crl
}

func displayRegisteredClusters(opts *options) {
	dump.Table(opts.f.IOStreams.Out, opts.registeredClusters)
	opts.f.Logger.Info("")
}
