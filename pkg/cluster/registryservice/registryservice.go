package registryservice

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redhat-developer/app-services-cli/pkg/cluster"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/constants/serviceregistry"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/utils"
	registryPkg "github.com/redhat-developer/app-services-cli/pkg/serviceregistry"
)

type ServiceRegistryConnection struct {
}

func (r *ServiceRegistryConnection) CustomResourceExists(ctx context.Context, c *cluster.KubernetesCluster, namespace string, serviceName string, opts cluster.Options) error {

	path := serviceregistry.GetServiceRegistryAPIURL(namespace)

	err := utils.ResourceExists(ctx, c, path, serviceName, opts)

	return err
}

func (r *ServiceRegistryConnection) CreateCustomResource(ctx context.Context, c *cluster.KubernetesCluster, serviceID string, namespace string, opts cluster.Options) error {

	api := opts.Connection.API()

	path := serviceregistry.GetServiceRegistryAPIURL(namespace)

	registryInstance, _, err := registryPkg.GetServiceRegistryByID(ctx, api.ServiceRegistryMgmt(), serviceID)
	if err != nil {
		return err
	}

	serviceName := registryInstance.GetName()

	serviceRegistryCR := serviceregistry.CreateSRObject(serviceName, namespace, serviceID)

	crJSON, err := json.Marshal(serviceRegistryCR)
	if err != nil {
		return fmt.Errorf("%v: %w", opts.Localizer.MustLocalize("cluster.kubernetes.createKafkaCR.error.marshalError"), err)
	}

	resource := serviceregistry.SRCResource

	err = utils.CreateResource(ctx, c, path, serviceName, namespace, crJSON, resource, opts, getWatchErrorMessages())

	return err
}

func getWatchErrorMessages() map[string]string {

	errorMessages := make(map[string]string)

	errorMessages["statusError"] = "cluster.kubernetes.watchForRegistryStatus.error.status"
	errorMessages["timeoutError"] = "cluster.kubernetes.watchForRegistryStatus.error.timeout"
	errorMessages["awaitStatus"] = "cluster.kubernetes.watchForRegistryStatus.log.info.wait"
	errorMessages["successfullyCreated"] = "cluster.kubernetes.watchForRegistryStatus.log.info.success"
	errorMessages["customResourceCreated"] = "cluster.kubernetes.createRegistryCR.log.info.customResourceCreated"

	return errorMessages
}
