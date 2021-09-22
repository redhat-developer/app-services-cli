package cluster

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redhat-developer/app-services-cli/pkg/api/kas"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/kafka"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/serviceregistry"
	"github.com/redhat-developer/app-services-cli/pkg/kafka/kafkaerr"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	registryPkg "github.com/redhat-developer/app-services-cli/pkg/serviceregistry"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
)

type Service interface {
	ResourceExists(ctx context.Context, c *KubernetesCluster, namespace string, serviceName string, opts Options) error
	CreateResource(ctx context.Context, c *KubernetesCluster, namespace string, serviceID string, opts Options) error
}

type Kafka struct{}

func (k *Kafka) ResourceExists(ctx context.Context, c *KubernetesCluster, namespace string, serviceName string, opts Options) error {

	path := kafka.GetKafkaConnectionsAPIURL(namespace)

	err := c.makeKubernetesGetRequest(ctx, path, serviceName, opts.Localizer)

	return err

}

func (k *Kafka) CreateResource(ctx context.Context, c *KubernetesCluster, namespace string, serviceID string, opts Options) error {

	api := opts.Connection.API()

	path := kafka.GetKafkaConnectionsAPIURL(namespace)

	kafkaInstance, _, err := api.Kafka().GetKafkaById(ctx, serviceID).Execute()
	if kas.IsErr(err, kas.ErrorNotFound) {
		return kafkaerr.NotFoundByIDError(serviceID)
	}

	serviceName := kafkaInstance.GetName()

	kafkaConnectionCR := kafka.CreateKCObject(serviceName, namespace, serviceID)

	CRJson, err := json.Marshal(kafkaConnectionCR)
	if err != nil {
		return fmt.Errorf("%v: %w", opts.Localizer.MustLocalize("cluster.kubernetes.createKafkaCR.error.marshalError"), err)
	}

	opts.Logger.Info(opts.Localizer.MustLocalize("cluster.kubernetes.createKafkaCR.log.info.customResourceCreated", localize.NewEntry("Name", serviceName)))

	err = c.makeKubernetesPostRequest(ctx, path, serviceName, CRJson)

	if err != nil {
		return err
	}

	resource := kafka.AKCResource

	w, err := c.dynamicClient.Resource(resource).Namespace(namespace).Watch(context.TODO(), metav1.ListOptions{
		FieldSelector: fields.OneTermEqualSelector("metadata.name", serviceName).String(),
	})
	if err != nil {
		return err
	}

	return watchCustomResourceStatus(w, namespace, serviceID, opts)

}

type ServiceRegistry struct{}

func (r *ServiceRegistry) ResourceExists(ctx context.Context, c *KubernetesCluster, namespace string, serviceName string, opts Options) error {

	path := serviceregistry.GetServiceRegistryAPIURL(namespace)

	err := c.makeKubernetesGetRequest(ctx, path, serviceName, opts.Localizer)

	return err

}

func (r *ServiceRegistry) CreateResource(ctx context.Context, c *KubernetesCluster, namespace string, serviceID string, opts Options) error {

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

	opts.Logger.Info(opts.Localizer.MustLocalize("cluster.kubernetes.createKafkaCR.log.info.customResourceCreated", localize.NewEntry("Name", serviceName)))

	err = c.makeKubernetesPostRequest(ctx, path, namespace, crJSON)
	if err != nil {
		return err
	}

	opts.Logger.Info(opts.Localizer.MustLocalize("cluster.kubernetes.watchForRegistryStatus.log.info.wait"))
	resource := serviceregistry.SRCResource

	w, err := c.dynamicClient.Resource(resource).Namespace(namespace).Watch(context.TODO(), metav1.ListOptions{
		FieldSelector: fields.OneTermEqualSelector("metadata.name", serviceName).String(),
	})
	if err != nil {
		return err
	}

	return watchCustomResourceStatus(w, namespace, serviceName, opts)
}
