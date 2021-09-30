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
)

// IsKCInstalledOnCluster checks the cluster to see if a KafkaConnection CRD is installed
func IsKCInstalledOnCluster(ctx context.Context, c *KubernetesCluster) (bool, error) {
	namespace, err := c.CurrentNamespace()
	if err != nil {
		return false, err
	}

	data := c.Clientset.
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

// IsSRCInstalledOnCluster checks the cluster to see if a ServiceRegistry CRD is installed
func IsSRCInstalledOnCluster(ctx context.Context, c *KubernetesCluster) (bool, error) {
	namespace, err := c.CurrentNamespace()
	if err != nil {
		return false, err
	}

	data := c.Clientset.
		RESTClient().
		Get().
		AbsPath(serviceregistry.GetServiceRegistryAPIURL(namespace)).
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

// IsSBOInstalledOnCluster checks the cluster to see if ServiceBinding CRD is installed
func IsSBOInstalledOnCluster(ctx context.Context, c *KubernetesCluster) (bool, error) {

	namespace, err := c.CurrentNamespace()
	if err != nil {
		return false, err
	}

	data := c.Clientset.
		RESTClient().
		Get().
		AbsPath(fmt.Sprintf("/apis/binding.operators.coreos.com/v1alpha1/namespaces/%v/servicebindings", namespace)).
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
