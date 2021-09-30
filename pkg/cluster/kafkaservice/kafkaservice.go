package kafkaservice

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"

	"github.com/redhat-developer/app-services-cli/pkg/api/kas"
	"github.com/redhat-developer/app-services-cli/pkg/cluster"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/constants"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/constants/kafka"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/utils"
	"github.com/redhat-developer/app-services-cli/pkg/kafka/kafkaerr"
	"github.com/redhat-developer/service-binding-operator/apis/binding/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/dynamic"
)

// KafkaService contains methods to connect and bind Kafka instance to cluster
type KafkaService struct {
	Opts cluster.Options
}

func (k *KafkaService) CustomResourceExists(ctx context.Context, c *cluster.KubernetesCluster, serviceName string) error {

	ns, err := c.CurrentNamespace()
	if err != nil {
		return err
	}

	path := kafka.GetKafkaConnectionsAPIURL(ns)

	err = c.ResourceExists(ctx, path, serviceName, k.Opts)

	return err
}

func (k *KafkaService) CreateCustomResource(ctx context.Context, c *cluster.KubernetesCluster, serviceID string) error {

	ns, err := c.CurrentNamespace()
	if err != nil {
		return err
	}

	api := k.Opts.Connection.API()

	// path := kafka.GetKafkaConnectionsAPIURL(ns)

	kafkaInstance, _, err := api.Kafka().GetKafkaById(ctx, serviceID).Execute()
	if kas.IsErr(err, kas.ErrorNotFound) {
		return kafkaerr.NotFoundByIDError(serviceID)
	}

	serviceName := kafkaInstance.GetName()

	kafkaConnectionCR := createKCObject(serviceName, ns, serviceID)

	crJSON, err := json.Marshal(kafkaConnectionCR)
	if err != nil {
		return fmt.Errorf("%v: %w", k.Opts.Localizer.MustLocalize("cluster.kubernetes.createKafkaCR.error.marshalError"), err)
	}

	resourceOpts := &cluster.CustomResourceOptions{
		CRName:      kafka.AKCRMeta.Kind,
		Resource:    kafka.AKCResource,
		CRJSON:      crJSON,
		ServiceName: serviceName,
		Path:        kafka.GetKafkaConnectionsAPIURL(ns),
	}

	err = c.CreateResource(ctx, resourceOpts, k.Opts)

	return err
}

func (k *KafkaService) CustomConnectionExists(ctx context.Context, dynamicClient dynamic.Interface, serviceName string, ns string) error {
	_, err := dynamicClient.Resource(kafka.AKCResource).Namespace(ns).Get(ctx, serviceName, metav1.GetOptions{})
	if err != nil {
		return k.Opts.Localizer.MustLocalizeError("cluster.serviceBinding.serviceMissing.message")
	}
	return nil
}

func (k *KafkaService) BindCustomConnection(ctx context.Context, serviceName string, options cluster.ServiceBindingOptions, clients *cluster.KubernetesClients) error {

	serviceRef := createKCServiceRef(serviceName)

	appRef := constants.CreateAppRef(options.AppName)

	if options.BindingName == "" {
		randomValue := make([]byte, 2)
		_, err := rand.Read(randomValue)
		if err != nil {
			return err
		}
		options.BindingName = fmt.Sprintf("%v-%x", serviceName, randomValue)
	}

	sb := constants.CreateSBObject(options.BindingName, options.Namespace, &serviceRef, &appRef)

	err := utils.CheckIfOperatorIsInstalled(ctx, clients.DynamicClient, options.Namespace)
	if err != nil {
		return fmt.Errorf("%s: %w", k.Opts.Localizer.MustLocalizeError("cluster.serviceBinding.operatorMissing"), err)
	}

	return utils.UseOperatorForBinding(ctx, k.Opts, sb, clients.DynamicClient, options.Namespace)
}

func createKCObject(crName string, namespace string, kafkaID string) *kafka.KafkaConnection {
	kafkaConnectionCR := &kafka.KafkaConnection{
		ObjectMeta: metav1.ObjectMeta{
			Name:      crName,
			Namespace: namespace,
		},
		TypeMeta: kafka.AKCRMeta,
		Spec: kafka.KafkaConnectionSpec{
			KafkaID:               kafkaID,
			AccessTokenSecretName: constants.TokenSecretName,
			Credentials: kafka.CredentialsSpec{
				SecretName: constants.ServiceAccountSecretName,
			},
		},
	}

	return kafkaConnectionCR
}

func createKCServiceRef(serviceName string) v1alpha1.Service {
	serviceRef := v1alpha1.Service{
		NamespacedRef: v1alpha1.NamespacedRef{
			Ref: v1alpha1.Ref{
				Group:    kafka.AKCResource.Group,
				Version:  kafka.AKCResource.Version,
				Resource: kafka.AKCResource.Resource,
				Name:     serviceName,
			},
		},
	}
	return serviceRef
}
