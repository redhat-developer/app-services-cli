package kubeclient

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/redhat-developer/app-services-cli/pkg/cluster/v1alpha"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

// KubernetesClients - contains all sorts of kubernetes clients that can be used to contact with cluster
type KubernetesClients struct {
	Clientset          *kubernetes.Clientset
	RestConfig         *rest.Config
	DynamicClient      dynamic.Interface
	ClientConfig       clientcmd.ClientConfig
	kubeconfigLocation string

	CommandEnvironment *v1alpha.CommandEnvironment
}

// NewKubernetesClusterClients configures and returns clients for kubernetes cluster
func NewKubernetesClusterClients(env *v1alpha.CommandEnvironment, kubeconfig string) (*KubernetesClients, error) {
	if kubeconfig == "" {
		kubeconfig = os.Getenv("KUBECONFIG")
	}

	if kubeconfig == "" {
		home, _ := os.UserHomeDir()
		kubeconfig = filepath.Join(home, ".kube", "config")
	}

	_, err := os.Stat(kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", env.Localizer.MustLocalize("cluster.kubernetes.error.configNotFoundError"), err)
	}

	kubeClientConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", env.Localizer.MustLocalize("cluster.kubernetes.error.loadConfigError"), err)
	}

	// create the clientset for using Rest Client
	clientset, err := kubernetes.NewForConfig(kubeClientConfig)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", env.Localizer.MustLocalize("cluster.kubernetes.error.loadConfigError"), err)
	}

	// Used for namespaces and general queries
	clientConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfig},
		&clientcmd.ConfigOverrides{ClusterInfo: clientcmdapi.Cluster{Server: ""}})

	dynamicClient, err := dynamic.NewForConfig(kubeClientConfig)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", env.Localizer.MustLocalize("cluster.kubernetes.error.loadConfigError"), err)
	}

	restConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", env.Localizer.MustLocalize("cluster.kubernetes.error.loadConfigError"), err)
	}

	k8sCluster := &KubernetesClients{
		clientset,
		restConfig,
		dynamicClient,
		clientConfig,
		kubeconfig,
		env,
	}

	return k8sCluster, nil
}

// CurrentNamespace returns the currently set namespace
func (c *KubernetesClients) CurrentNamespace() (string, error) {
	namespace, _, err := c.ClientConfig.Namespace()
	return namespace, err
}

// IsResourceAvailableOnCluster checks the cluster to see if specified CRD is installed
func (c *KubernetesClients) IsResourceAvailableOnCluster(resource *schema.GroupVersionResource, namespace string) (bool, error) {
	env := c.CommandEnvironment

	_, err := c.DynamicClient.Resource(*resource).Namespace(namespace).
		List(env.Context, metav1.ListOptions{Limit: 1})

	if err == nil {
		return true, nil
	}

	if errors.IsNotFound(err) || errors.IsForbidden(err) {
		return false, nil
	}

	if errors.IsServiceUnavailable(err) || errors.IsServerTimeout(err) || errors.IsTimeout(err) {
		return false, env.Localizer.MustLocalizeError("cluster.common.kube.unauthorized")
	}

	return true, err
}
