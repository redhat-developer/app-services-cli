package services

import (
	"github.com/redhat-developer/app-services-cli/pkg/cluster/constants"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/kubeclient"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/services/resources"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/v1alpha"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/serviceregistry"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RegistryService contains methods to connect and bind Service registry instance to cluster
type RegistryService struct {
	CommandEnvironment *v1alpha.CommandEnvironment
	KubernetesClients  *kubeclient.KubernetesClients
}

func (s RegistryService) BuildServiceDetails(serviceName string, namespace string, ignoreContext bool) (*ServiceDetails, error) {
	cliOpts := s.CommandEnvironment
	cfg, err := cliOpts.Config.Load()
	if err != nil {
		return nil, err
	}

	api := cliOpts.Connection.API()
	var serviceId string

	if serviceName == "" {
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
			serviceName = selectedService.GetName()
		} else {
			serviceId = cfg.Services.ServiceRegistry.InstanceID
			selectedService, _, err := serviceregistry.GetServiceRegistryByID(
				cliOpts.Context, api.ServiceRegistryMgmt(), serviceId)
			if err != nil {
				return nil, err
			}
			serviceName = selectedService.GetName()
		}
	} else {
		selectedService, _, err := serviceregistry.GetServiceRegistryByName(cliOpts.Context, api.ServiceRegistryMgmt(), serviceName)
		if err != nil {
			return nil, err
		}
		serviceId = selectedService.GetId()
	}

	serviceRegistryCR := &resources.ServiceRegistryConnection{
		ObjectMeta: metav1.ObjectMeta{
			Name:      serviceName,
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
		Name:               serviceName,
		KubernetesResource: serviceRegistryCR,
		GroupMetadata:      resources.SRCResource,
		Type:               resources.ServiceRegistryServiceName,
	}

	return &serviceDetails, nil
}

func (s RegistryService) PrintAccessCommands(clientID string) {
	cliOpts := s.CommandEnvironment
	cliOpts.Logger.Info(cliOpts.Localizer.MustLocalize("cluster.kubernetes.printRegistryAccessCommands", localize.NewEntry("ClientID", clientID)))
}
