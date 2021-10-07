package services

import (
	"github.com/redhat-developer/app-services-cli/pkg/api/srs"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/constants"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/kubeclient"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/services/resources"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/v1alpha"
	"github.com/redhat-developer/app-services-cli/pkg/serviceregistry"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// RegistryService contains methods to connect and bind Service registry instance to cluster
type RegistryService struct {
	CommandEnvironment *v1alpha.CommandEnvironment
	KubernetesClients  *kubeclient.KubernetesClients
}

func (s RegistryService) BuildServiceDetails(serviceId string, namespace string, ignoreContext bool) (*ServiceDetails, error) {
	cliOpts := s.CommandEnvironment
	if serviceId == "" {
		cfg, err := s.CommandEnvironment.Config.Load()
		if err != nil {
			return nil, err
		}

		if cfg.Services.ServiceRegistry == nil || ignoreContext {
			// nolint
			selectedService, err := serviceregistry.InteractiveSelect(cliOpts.Context, cliOpts.Connection, cliOpts.Logger)
			if err != nil {
				return nil, err
			}
			if selectedService == nil {
				return nil, nil
			}
			serviceId = selectedService.GetId()
		} else {
			serviceId = cfg.Services.ServiceRegistry.InstanceID
		}
	}
	api := cliOpts.Connection.API()
	serviceInstance, _, err := api.ServiceRegistryMgmt().GetRegistry(cliOpts.Context, serviceId).Execute()

	if srs.IsErr(err, srs.ErrorNotFound) {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	resourceName := serviceInstance.GetName()

	serviceRegistryCR := &resources.ServiceRegistryConnection{
		ObjectMeta: metav1.ObjectMeta{
			Name:      resourceName,
			Namespace: namespace,
		},
		TypeMeta: resources.RegistryResourceMeta,
		Spec: resources.ServiceRegistryConnectionSpec{
			ServiceRegistryId:     serviceId,
			AccessTokenSecretName: constants.TokenSecretName,
			Credentials: resources.RegistryCredentialsSpec{
				SecretName: constants.ServiceAccountSecretName,
			},
		},
	}

	serviceDetails := ServiceDetails{
		ID:                 serviceId,
		Name:               resourceName,
		KubernetesResource: serviceRegistryCR,
		GroupMetadata:      resources.SRCResource,
		Type:               resources.ServiceRegistryServiceName,
	}

	return &serviceDetails, nil
}

func (s RegistryService) BuildServiceCustomResourceMetadata() (schema.GroupVersionResource, error) {
	return resources.SRCResource, nil
}
