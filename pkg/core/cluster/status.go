package cluster

import (
	"github.com/redhat-developer/app-services-cli/pkg/core/cluster/services/resources"
	"github.com/redhat-developer/app-services-cli/pkg/core/cluster/v1alpha"
	bindv1alpha1 "github.com/redhat-developer/service-binding-operator/apis/binding/v1alpha1"
)

// ExecuteStatus executes status command by checking availability of operators and resources
// When unexpected error happens (non 404 error) function returns only error to the user
func (api *KubernetesClusterAPIImpl) ExecuteStatus() (*v1alpha.OperatorStatus, error) {
	// Start with all resources unavailable
	operatorStatus := v1alpha.OperatorStatus{
		ServiceBindingOperatorAvailable: false,
		RHOASOperatorAvailable:          false,
		LatestRHOASVersionAvailable:     false,
	}
	namespace, err := api.KubernetesClients.CurrentNamespace()
	if err != nil {
		return nil, err
	}
	installed, err := api.KubernetesClients.IsResourceAvailableOnCluster(&resources.SRCResource, namespace)
	// If unhandled error return instantly
	if err != nil {
		return nil, err
	}

	// Assign respective status otherwise
	operatorStatus.LatestRHOASVersionAvailable = installed

	installed, err = api.KubernetesClients.IsResourceAvailableOnCluster(&resources.AKCResource, namespace)
	// If unhandled error return instantly
	if err != nil {
		return nil, err
	}

	operatorStatus.RHOASOperatorAvailable = installed

	installed, err = api.KubernetesClients.IsResourceAvailableOnCluster(&bindv1alpha1.GroupVersionResource, namespace)
	// If unhandled error return instantly
	if err != nil {
		return nil, err
	}

	operatorStatus.ServiceBindingOperatorAvailable = installed

	return &operatorStatus, err
}
