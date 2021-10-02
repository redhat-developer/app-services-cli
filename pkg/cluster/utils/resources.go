package utils

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redhat-developer/app-services-cli/pkg/cluster/constants/kafka"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/constants/serviceregistry"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/kubeclient"
	"k8s.io/apimachinery/pkg/api/errors"
)

// THIS IS WRONG - should be part of the service

// IsKCInstalledOnCluster checks the cluster to see if a KafkaConnection CRD is installed
func IsKCInstalledOnCluster(ctx context.Context, c *kubeclient.KubernetesClients) (bool, error) {
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
func IsSRCInstalledOnCluster(ctx context.Context, c *kubeclient.KubernetesClients) (bool, error) {
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

	// TODO verify if this handling works
	if errors.IsNotFound(err) {
		return false, nil
	}

	var status int
	if data.StatusCode(&status); status == http.StatusNotFound {
		return false, nil
	}

	return true, data.Error()
}

// IsSBOInstalledOnCluster checks the cluster to see if ServiceBinding CRD is installed
func IsSBOInstalledOnCluster(ctx context.Context, c *kubeclient.KubernetesClients) (bool, error) {
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
	// TODO verify if this handling works
	if errors.IsNotFound(err) {
		return false, nil
	}

	var status int
	if data.StatusCode(&status); status == http.StatusNotFound {
		return false, nil
	}

	return true, data.Error()

}
