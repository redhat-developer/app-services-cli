package kafkaservice

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redhat-developer/app-services-cli/pkg/api/kas"
	"github.com/redhat-developer/app-services-cli/pkg/cluster"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/constants/kafka"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/utils"
	"github.com/redhat-developer/app-services-cli/pkg/kafka/kafkaerr"
)

type KafkaConnection struct {
}

func (k *KafkaConnection) CustomResourceExists(ctx context.Context, c *cluster.KubernetesCluster, namespace string, serviceName string, opts cluster.Options) error {

	path := kafka.GetKafkaConnectionsAPIURL(namespace)

	err := utils.ResourceExists(ctx, c, path, serviceName, opts)

	return err
}

func (k *KafkaConnection) CreateCustomResource(ctx context.Context, c *cluster.KubernetesCluster, serviceID string, namespace string, opts cluster.Options) error {

	api := opts.Connection.API()

	path := kafka.GetKafkaConnectionsAPIURL(namespace)

	kafkaInstance, _, err := api.Kafka().GetKafkaById(ctx, serviceID).Execute()
	if kas.IsErr(err, kas.ErrorNotFound) {
		return kafkaerr.NotFoundByIDError(serviceID)
	}

	serviceName := kafkaInstance.GetName()

	kafkaConnectionCR := kafka.CreateKCObject(serviceName, namespace, serviceID)

	crJSON, err := json.Marshal(kafkaConnectionCR)
	if err != nil {
		return fmt.Errorf("%v: %w", opts.Localizer.MustLocalize("cluster.kubernetes.createKafkaCR.error.marshalError"), err)
	}

	resource := kafka.AKCResource

	err = utils.CreateResource(ctx, c, path, serviceName, namespace, crJSON, resource, opts, GetWatchErrorMessages())

	return err
}

func GetWatchErrorMessages() map[string]string {

	errorMessages := make(map[string]string)

	errorMessages["statusError"] = "cluster.kubernetes.watchForKafkaStatus.error.status"
	errorMessages["timeoutError"] = "cluster.kubernetes.watchForKafkaStatus.error.timeout"
	errorMessages["awaitStatus"] = "cluster.kubernetes.watchForKafkaStatus.log.info.wait"
	errorMessages["successfullyCreated"] = "cluster.kubernetes.watchForKafkaStatus.log.info.success"
	errorMessages["customResourceCreated"] = "cluster.kubernetes.createKafkaCR.log.info.customResourceCreated"

	return errorMessages
}
