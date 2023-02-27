package listclusters

import (
	"context"
	"fmt"
	clustersmgmtv1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection/api/clustermgmt"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
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
	outputFormat            string
	search                  string
	kfmClusterList          *kafkamgmtclient.EnterpriseClusterList
	clustermgmtClusterList  []*clustersmgmtv1.Cluster
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
		Short:   f.Localizer.MustLocalize("dedicated.list.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("dedicated.list.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("dedicated.list.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runListClusters(opts, f)
		},
	}

	return cmd
}

func runListClusters(opts *options, f *factory.Factory) error {
	err := listEnterpriseClusters(opts, f)
	if err != nil {
		return err
	}
	clist, err := getPaginatedClusterList(opts)
	if err != nil {
		return err
	}
	opts.clustermgmtClusterList = clist.Slice()
	opts.registeredClusters = kfmListToClusterRowList(opts)
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
	opts.kfmClusterList = &clist
	f.Logger.Info(response)
	if len(opts.kfmClusterList.Items) == 0 {
		return f.Localizer.MustLocalizeError("dedicated.listClusters.log.info.noClusters")
	}
	return nil
}

func createSearchString(opts *options) string {
	searchString := ""
	for idx, kfmcluster := range opts.kfmClusterList.Items {
		searchString += fmt.Sprintf("id = '%s'", kfmcluster.Id)
		if idx > 0 {
			searchString += " or "
		}
	}
	return searchString
}

func getPaginatedClusterList(opts *options) (*clustersmgmtv1.ClusterList, error) {
	// get ids of clusters and create an ocm call filtering by those ids
	ocmcl, err := clustermgmt.GetClusterListByIds(opts.f, opts.accessToken, opts.clusterManagementApiUrl, createSearchString(opts), len(opts.kfmClusterList.Items))
	if err != nil {
		return nil, err
	}
	return ocmcl, nil
}

func kfmListToClusterRowList(opts *options) []clusterRow {
	var crl []clusterRow
	for _, kfmcluster := range opts.kfmClusterList.Items {
		for _, ocmCluster := range opts.clustermgmtClusterList {
			if kfmcluster.Id == ocmCluster.ID() {
				crl = append(crl, clusterRow{
					Name:          ocmCluster.Name(),
					ID:            kfmcluster.Id,
					Status:        *kfmcluster.Status,
					CloudProvider: ocmCluster.CloudProvider().ID(),
					Region:        ocmCluster.Region().ID(),
				})
			}
		}
	}
	return crl
}

func displayRegisteredClusters(opts *options) error {
	if len(opts.registeredClusters) == 0 {
		return opts.f.Localizer.MustLocalizeError("dedicated.list.cmd.errorNoRegisteredClusters")
	}
	dump.Table(opts.f.IOStreams.Out, opts.registeredClusters)
	opts.f.Logger.Info("")

	return nil
}
