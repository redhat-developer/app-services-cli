package clustermgmt

import (
	"github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
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