package services

import (
	"github.com/redhat-developer/app-services-cli/pkg/api/kas"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/constants"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/kubeclient"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/services/resources"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/v1alpha"
	"github.com/redhat-developer/app-services-cli/pkg/kafka"
	"github.com/redhat-developer/app-services-cli/pkg/kafka/kafkaerr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// KafkaService contains methods to connect and bind Kafka Service instance to cluster
type KafkaService struct {
	CommandEnvironment *v1alpha.CommandEnvironment
	KubernetesClients  *kubeclient.KubernetesClients
}

func (s KafkaService) BuildServiceDetails(serviceId string, namespace string, ignoreContext bool) (*ServiceDetails, error) {
	cliOpts := s.CommandEnvironment
	if serviceId == "" {
		cfg, err := s.CommandEnvironment.Config.Load()
		if err != nil {
			return nil, err
		}

		if cfg.Services.Kafka == nil || ignoreContext {
			// nolint
			selectedService, err := kafka.InteractiveSelect(cliOpts.Context, cliOpts.Connection, cliOpts.Logger, cliOpts.Localizer)
			if err != nil {
				return nil, err
			}
			if selectedService == nil {
				return nil, nil
			}
			serviceId = selectedService.GetId()
		} else {
			serviceId = cfg.Services.Kafka.ClusterID
		}
	}
	api := cliOpts.Connection.API()
	serviceInstance, _, err := api.Kafka().GetKafkaById(cliOpts.Context, serviceId).Execute()

	if kas.IsErr(err, kas.ErrorCode7) {
		return nil, kafkaerr.NotFoundByIDError(serviceId)
	}

	if err != nil {
		return nil, err
	}

	resourceName := serviceInstance.GetName()

	kafkaConnectionCR := &resources.KafkaConnection{
		ObjectMeta: metav1.ObjectMeta{
			Name:      resourceName,
			Namespace: namespace,
		},
		TypeMeta: resources.AKCRMeta,
		Spec: resources.KafkaConnectionSpec{
			KafkaID:               serviceId,
			AccessTokenSecretName: constants.TokenSecretName,
			Credentials: resources.KafkaCredentialsSpec{
				SecretName: constants.ServiceAccountSecretName,
			},
		},
	}

	serviceDetails := ServiceDetails{
		ID:                 serviceId,
		Name:               resourceName,
		KubernetesResource: kafkaConnectionCR,
		GroupMetadata:      resources.AKCResource,
		Type:               resources.KafkaServiceName,
	}

	return &serviceDetails, nil
}

func (s KafkaService) BuildServiceCustomResourceMetadata() (schema.GroupVersionResource, error) {
	return resources.AKCResource, nil
}
