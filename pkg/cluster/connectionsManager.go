/**
 * Handles specific operations for Kafka Connection resource
 */
package cluster

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redhat-developer/app-services-cli/pkg/cluster/constants/kafka"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/constants/serviceregistry"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
	srsmgmtv1 "github.com/redhat-developer/app-services-sdk-go/registrymgmt/apiv1/client"
)

// IsKCInstalledOnCluster checks the cluster to see if a KafkaConnection CRD is installed
func IsKCInstalledOnCluster(ctx context.Context, c *KubernetesCluster) (bool, error) {
	namespace, err := c.CurrentNamespace()
	if err != nil {
		return false, err
	}

	data := c.clientset.
		RESTClient().
		Get().
		AbsPath(kafka.GetKafkaConnectionsAPIURL(namespace)).
		Do(ctx)

	if data.Error() == nil {
		return true, nil
	}

	var status int
	if data.StatusCode(&status); status == http.StatusNotFound {
		return false, nil
	}

	return true, data.Error()
}

func CheckIfConnectionExists(ctx context.Context, c *KubernetesCluster, namespace string, service interface{}, serviceName string, opts Options) error {
	var path string
	var status int

	switch service.(type) {
	case *srsmgmtv1.Registry:
		path = serviceregistry.GetServiceRegistryAPIURL(namespace)
	case kafkamgmtclient.KafkaRequest:
		path = kafka.GetKafkaConnectionsAPIURL(namespace)
	}

	data := c.clientset.
		RESTClient().
		Get().
		AbsPath(path, serviceName).
		Do(ctx)

	if data.StatusCode(&status); status == http.StatusNotFound {
		return nil
	}

	if data.Error() == nil {
		return fmt.Errorf("%v: %s", opts.Localizer.MustLocalize("cluster.kubernetes.checkIfConnectionExist.existError"), serviceName)
	}

	return nil
}
