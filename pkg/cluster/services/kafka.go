package services

import (
	"github.com/redhat-developer/app-services-cli/pkg/cluster/constants"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/kubeclient"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/services/resources"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/v1alpha"
	"github.com/redhat-developer/app-services-cli/pkg/kafka"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KafkaService contains methods to connect and bind Kafka Service instance to cluster
type KafkaService struct {
	CommandEnvironment *v1alpha.CommandEnvironment
	KubernetesClients  *kubeclient.KubernetesClients
}

func (s KafkaService) BuildServiceDetails(serviceName string, namespace string, ignoreContext bool) (*ServiceDetails, error) {
	cliOpts := s.CommandEnvironment
	cfg, err := cliOpts.Config.Load()
	if err != nil {
		return nil, err
	}

	api := cliOpts.Connection.API()
	var serviceId string

	if serviceName == "" {
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
			serviceName = selectedService.GetName()
		} else {
			serviceId = cfg.Services.Kafka.ClusterID
			selectedService, _, err := kafka.GetKafkaByID(cliOpts.Context, api.Kafka(), serviceId)
			if err != nil {
				return nil, err
			}
			serviceName = selectedService.GetName()
		}
	} else {
		selectedService, _, err := kafka.GetKafkaByName(cliOpts.Context, api.Kafka(), serviceName)
		if err != nil {
			return nil, err
		}
		serviceId = selectedService.GetId()
	}

	kafkaConnectionCR := &resources.KafkaConnection{
		ObjectMeta: metav1.ObjectMeta{
			Name:      serviceName,
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
		Name:               serviceName,
		KubernetesResource: kafkaConnectionCR,
		GroupMetadata:      resources.AKCResource,
		Type:               resources.KafkaServiceName,
	}

	return &serviceDetails, nil
}

// PrintAccessCommands prints command to grant service account acccess to the Kafka instance
func (s KafkaService) PrintAccessCommands(clientID string) {
	cliOpts := s.CommandEnvironment
	cliOpts.Logger.Info(cliOpts.Localizer.MustLocalize("cluster.kubernetes.printKafkaAccessCommands", localize.NewEntry("ClientID", clientID)))
}
