package cluster

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/pkg/core/cluster/kubeclient"
	"github.com/redhat-developer/app-services-cli/pkg/core/cluster/services"
	"github.com/redhat-developer/app-services-cli/pkg/core/cluster/services/resources"
	"github.com/redhat-developer/app-services-cli/pkg/core/cluster/v1alpha"
)

// KubernetesClusterAPIImpl	implements KubernetesClusterAPI
type KubernetesClusterAPIImpl struct {
	KubernetesClients  *kubeclient.KubernetesClients
	CommandEnvironment *v1alpha.CommandEnvironment
}

// see bind.go cluster.go status.go for interface implementations

func (c *KubernetesClusterAPIImpl) createServiceInstance(serviceType string) (services.RHOASKubernetesService, error) {
	var service services.RHOASKubernetesService
	if serviceType == "" {
		serviceTypeInput := &survey.Select{
			Message: c.CommandEnvironment.Localizer.MustLocalize("cluster.common.input.servicetype"),
			Options: resources.AllServiceLabels,
		}
		surveyErr := survey.AskOne(serviceTypeInput, &serviceType)
		if surveyErr != nil {
			return nil, surveyErr
		}
	}

	switch serviceType {
	case resources.KafkaServiceName:
		service = &services.KafkaService{
			CommandEnvironment: c.CommandEnvironment,
			KubernetesClients:  c.KubernetesClients,
		}
	case resources.ServiceRegistryServiceName:
		service = &services.RegistryService{
			CommandEnvironment: c.CommandEnvironment,
			KubernetesClients:  c.KubernetesClients,
		}
	default:
		return nil, c.CommandEnvironment.Localizer.MustLocalizeError("cluster.common.error.servicetype")
	}

	return service, nil
}
