package cluster

import (
	"github.com/redhat-developer/app-services-cli/pkg/cluster/kubeclient"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/utils"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/v1alpha"
)

// KubernetesClusterAPIImpl	implements KubernetesClusterAPI
type KubernetesClusterAPIImpl struct {
	KubernetesClients  *kubeclient.KubernetesClients
	CommandEnvironment *v1alpha.CommandEnvironment
}

func (api *KubernetesClusterAPIImpl) IsRhoasOperatorAvailableOnCluster() (bool, error) {
	installed, err := utils.IsKCInstalledOnCluster(api.CommandEnvironment.Context, api.KubernetesClients)
	if !installed {
		return installed, err
	}

	// TODO replace boolean and return v1 and v2 versions for user
	return utils.IsSRCInstalledOnCluster(api.CommandEnvironment.Context, api.KubernetesClients)

}

func (api *KubernetesClusterAPIImpl) IsSBOOperatorAvailableOnCluster() (bool, error) {
	return utils.IsSBOInstalledOnCluster(api.CommandEnvironment.Context, api.KubernetesClients)
}
