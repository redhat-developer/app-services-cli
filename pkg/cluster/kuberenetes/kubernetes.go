package kubernetes

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/v1alpha"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

// TODO missing comment
type KubernetesCluster struct {
	Clientset          *kubernetes.Clientset
	clientconfig       clientcmd.ClientConfig
	DynamicClient      dynamic.Interface
	kubeconfigLocation string
}

// TODO unify with KubernetesCluster this is used in binding
type KubernetesClients struct {
	DynamicClient dynamic.Interface
	restConfig    *rest.Config
	clientConfig  *clientcmd.ClientConfig
}

// NewKubernetesClusterConnection configures and connects to a Kubernetes cluster
func NewKubernetesClusterConnection(connection connection.Connection,
	config config.IConfig,
	logger logging.Logger,
	kubeconfig string,
	io *iostreams.IOStreams, localizer localize.Localizer) (v1alpha.Cluster, error) {
	if kubeconfig == "" {
		kubeconfig = os.Getenv("KUBECONFIG")
	}

	if kubeconfig == "" {
		home, _ := os.UserHomeDir()
		kubeconfig = filepath.Join(home, ".kube", "config")
	}

	_, err := os.Stat(kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", localizer.MustLocalize("cluster.kubernetes.error.configNotFoundError"), err)
	}

	kubeClientConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", localizer.MustLocalize("cluster.kubernetes.error.loadConfigError"), err)
	}

	// create the clientset for using Rest Client
	clientset, err := kubernetes.NewForConfig(kubeClientConfig)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", localizer.MustLocalize("cluster.kubernetes.error.loadConfigError"), err)
	}

	// Used for namespaces and general queries
	clientconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfig},
		&clientcmd.ConfigOverrides{ClusterInfo: clientcmdapi.Cluster{Server: ""}})

	dynamicClient, err := dynamic.NewForConfig(kubeClientConfig)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", localizer.MustLocalize("cluster.kubernetes.error.loadConfigError"), err)
	}

	k8sCluster := &KubernetesCluster{
		clientset,
		clientconfig,
		dynamicClient,
		kubeconfig,
	}

	return k8sCluster, nil
}

// CurrentNamespace returns the currently set namespace
func (c *KubernetesCluster) CurrentNamespace() (string, error) {
	namespace, _, err := c.clientconfig.Namespace()

	return namespace, err
}
