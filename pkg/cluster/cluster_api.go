package cluster

import (
	"github.com/redhat-developer/app-services-cli/pkg/cluster/kubeclient"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/v1alpha"
)

// KubernetesClusterAPIImpl	implements KubernetesClusterAPI
type KubernetesClusterAPIImpl struct {
	KubernetesClients  *kubeclient.KubernetesClients
	CommandEnvironment *v1alpha.CommandEnvironment
}

func (KubernetesClusterAPIImpl) ExecuteConnect(connectOpts *v1alpha.ConnectOperationOptions) error {
	return nil
}

func (KubernetesClusterAPIImpl) ExecuteServiceBinding(bindinOptions *v1alpha.BindOperationOptions) error {
	return nil
}

func (KubernetesClusterAPIImpl) IsRhoasOperatorAvailableOnCluster() (bool, error) {
	return true, nil
}

func (KubernetesClusterAPIImpl) IsSBOOperatorAvailableOnCluster() (bool, error) {
	return true, nil
}
