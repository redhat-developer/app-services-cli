package clustermgmt

import (
	clustersmgmtv1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	v1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/kafkamgmt/apiv1/client"
)

func clustermgmtConnection(f *factory.Factory, accessToken string, clustermgmturl string) (*v1.Client, func(), error) {
	conn, err := f.Connection()
	if err != nil {
		return nil, nil, err
	}
	client, closeConnection, err := conn.API().OCMClustermgmt(clustermgmturl, accessToken)
	if err != nil {
		return nil, nil, err
	}
	return client, closeConnection, nil
}

func GetClusterList(f *factory.Factory, accessToken string, clustermgmturl string, pageNumber int, pageLimit int) (*v1.ClusterList, error) {
	client, closeConnection, err := clustermgmtConnection(f, accessToken, clustermgmturl)
	if err != nil {
		return nil, err
	}
	defer closeConnection()

	resource := client.Clusters().List()
	resource = resource.Page(pageNumber)
	resource = resource.Size(pageLimit)
	response, err := resource.Send()
	if err != nil {
		return nil, err
	}
	clusters := response.Items()
	return clusters, nil
}

func GetMachinePoolList(f *factory.Factory, clustermgmturl string, accessToken string, clusterId string) (*v1.MachinePoolsListResponse, error) {
	client, closeConnection, err := clustermgmtConnection(f, accessToken, clustermgmturl)
	if err != nil {
		return nil, err
	}
	defer closeConnection()
	resource := client.Clusters().Cluster(clusterId).MachinePools().List()
	response, err := resource.Send()
	if err != nil {
		return nil, err
	}
	return response, nil
}

func GetClusterListByIds(f *factory.Factory, clustermgmturl string, accessToken string, params string, size int) (*v1.ClusterList, error) {
	client, closeConnection, err := clustermgmtConnection(f, accessToken, clustermgmturl)
	if err != nil {
		return nil, err
	}
	defer closeConnection()
	resource := client.Clusters().List().Search(params).Size(size)
	response, err := resource.Send()
	if err != nil {
		return nil, err
	}
	return response.Items(), nil
}

func CreateAddonWithParams(f *factory.Factory, clustermgmturl string, accessToken string, addonId string, params *[]kafkamgmtclient.FleetshardParameter, clusterId string) error {
	client, closeConnection, err := clustermgmtConnection(f, accessToken, clustermgmturl)
	if err != nil {
		return err
	}
	defer closeConnection()
	addon := v1.NewAddOn().ID(addonId)
	addonParameters := newAddonParameterListBuilder(params)
	addonInstallationBuilder := v1.NewAddOnInstallation().Addon(addon)
	if addonParameters != nil {
		addonInstallationBuilder = addonInstallationBuilder.Parameters(addonParameters)
	}
	addonInstallation, err := addonInstallationBuilder.Build()
	if err != nil {
		return err
	}
	_, err = client.Clusters().Cluster(clusterId).Addons().Add().Body(addonInstallation).Send()
	if err != nil {
		return err
	}

	return nil
}

func newAddonParameterListBuilder(params *[]kafkamgmtclient.FleetshardParameter) *v1.AddOnInstallationParameterListBuilder {
	if params == nil {
		return nil
	}
	var items []*v1.AddOnInstallationParameterBuilder
	for _, p := range *params {
		pb := v1.NewAddOnInstallationParameter().ID(*p.Id).Value(*p.Value)
		items = append(items, pb)
	}
	return v1.NewAddOnInstallationParameterList().Items(items...)
}

func CreateMachinePool(f *factory.Factory, clustermgmturl string, accessToken string, mprequest *v1.MachinePool, clusterId string) (*v1.MachinePool, error) {
	client, closeConnection, err := clustermgmtConnection(f, accessToken, clustermgmturl)
	if err != nil {
		return nil, err
	}
	defer closeConnection()
	response, err := client.Clusters().Cluster(clusterId).MachinePools().Add().Body(mprequest).Send()
	if err != nil {
		return nil, err
	}
	return response.Body(), nil
}

func GetMachinePoolNodeCount(machinePool *v1.MachinePool) int {
	var nodeCount int
	replicas, ok := machinePool.GetReplicas()
	if ok {
		nodeCount = replicas
	} else {
		autoscaledReplicas, ok := machinePool.GetAutoscaling()
		if ok {
			nodeCount = autoscaledReplicas.MaxReplicas()
		}
	}
	return nodeCount
}

func RemoveAddonsFromCluster(f *factory.Factory, clusterManagementApiUrl string, accessToken string, cluster *clustersmgmtv1.Cluster, addonList []string) error {
	// create a new addon via ocm
	conn, err := f.Connection()
	if err != nil {
		return err
	}
	client, cc, err := conn.API().OCMClustermgmt(clusterManagementApiUrl, accessToken)
	if err != nil {
		return err
	}
	defer cc()

	addons, err := client.Clusters().Cluster(cluster.ID()).Addons().List().Send()
	if err != nil {
		return err
	}

	for _, addonToDelete := range addonList {
		for i := 0; i < addons.Size(); i++ {
			addon := addons.Items().Get(i)

			if addon.ID() == addonToDelete {
				f.Logger.Info(f.Localizer.MustLocalize("dedicated.common.addons.deleting.message", localize.NewEntry("Id", addon.ID())))
				_, err = client.Clusters().Cluster(cluster.ID()).Addons().Addoninstallation(addon.ID()).Delete().Send()
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
