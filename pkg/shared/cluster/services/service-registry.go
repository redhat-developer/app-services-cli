package services

import (
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/shared/cluster/constants"
	"github.com/redhat-developer/app-services-cli/pkg/shared/cluster/kubeclient"
	"github.com/redhat-developer/app-services-cli/pkg/shared/cluster/services/resources"
	"github.com/redhat-developer/app-services-cli/pkg/shared/cluster/v1alpha"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/serviceregistryutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/servicespec"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RegistryService contains methods to connect and bind Service registry instance to cluster
type RegistryService struct {
	CommandEnvironment *v1alpha.CommandEnvironment
	KubernetesClients  *kubeclient.KubernetesClients
}

func (s RegistryService) BuildServiceDetails(serviceName string, namespace string, ignoreContext bool) (*ServiceDetails, error) {
	cliOpts := s.CommandEnvironment
	svcContext, err := cliOpts.ServiceContext.Load()
	if err != nil {
		return nil, err
	}

	currCtx, err := contextutil.GetCurrentContext(svcContext, cliOpts.Localizer)
	if err != nil {
		return nil, err
	}

	api := cliOpts.Connection.API()
	var serviceId string

	if serviceName == "" {
		if currCtx.ServiceRegistryID == "" || ignoreContext {
			// nolint
			selectedService, err := serviceregistryutil.InteractiveSelect(cliOpts.Context, cliOpts.Connection, cliOpts.Logger)
			if err != nil {
				return nil, err
			}
			if selectedService == nil {
				return nil, nil
			}
			serviceId = selectedService.GetId()
			serviceName = selectedService.GetName()
		} else {
			serviceId = currCtx.ServiceRegistryID
			selectedService, _, err := serviceregistryutil.GetServiceRegistryByID(
				cliOpts.Context, api.ServiceRegistryMgmt(), serviceId)
			if err != nil {
				return nil, err
			}
			serviceName = selectedService.GetName()
		}
	} else {
		selectedService, _, err := serviceregistryutil.GetServiceRegistryByName(cliOpts.Context, api.ServiceRegistryMgmt(), serviceName)
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
		Type:               servicespec.ServiceRegistryServiceName,
	}

	return &serviceDetails, nil
}

// PrintAccessCommands prints command to assign service account roles in the service registry instance
func (s RegistryService) PrintAccessCommands(clientID string) {
	cliOpts := s.CommandEnvironment
	cliOpts.Logger.Info(cliOpts.Localizer.MustLocalize("cluster.kubernetes.printRegistryAccessCommands", localize.NewEntry("ClientID", clientID)))
}
