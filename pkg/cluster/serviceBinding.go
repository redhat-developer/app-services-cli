package cluster

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/redhat-developer/app-services-cli/pkg/cluster/constants"
	"github.com/redhat-developer/app-services-cli/pkg/icon"

	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/pkg/color"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

type KubernetesClients struct {
	DynamicClient dynamic.Interface
	restConfig    *rest.Config
	clientConfig  *clientcmd.ClientConfig
}

var (
	deploymentConfigResource = schema.GroupVersionResource{Group: "apps.openshift.io", Version: "v1", Resource: "deploymentconfigs"}
)

type ServiceBindingOptions struct {
	ServiceName             string
	Namespace               string
	AppName                 string
	ForceCreationWithoutAsk bool
	BindingName             string
	BindAsFiles             bool
	DeploymentConfigEnabled bool
}

func (c *KubernetesCluster) ExecuteServiceBinding(ctx context.Context, service CustomConnection, opts Options, options *ServiceBindingOptions) error {
	clients, err := client(opts.Localizer)
	if err != nil {
		return err
	}
	ns := options.Namespace
	if ns == "" {
		ns, _, err = (*clients.clientConfig).Namespace()
		if err != nil {
			return err
		}
		opts.Logger.Info(opts.Localizer.MustLocalize("cluster.serviceBinding.namespaceInfo", localize.NewEntry("Namespace", color.Info(ns))))
	}

	var clusterResource schema.GroupVersionResource
	if options.DeploymentConfigEnabled {
		clusterResource = deploymentConfigResource
	} else {
		clusterResource = constants.DeploymentResource
	}
	// Get proper deployment
	if options.AppName == "" {
		options.AppName, err = fetchAppNameFromCluster(ctx, clusterResource, clients, opts.Localizer, ns)
		if err != nil {
			return err
		}
	} else {
		_, err = clients.DynamicClient.Resource(clusterResource).Namespace(ns).Get(ctx, options.AppName, metav1.GetOptions{})
		if err != nil {
			return err
		}
	}

	// Print desired action
	opts.Logger.Info(fmt.Sprintf(opts.Localizer.MustLocalize("cluster.serviceBinding.status.message"), options.ServiceName, options.AppName))

	if !options.ForceCreationWithoutAsk {
		var shouldContinue bool
		confirm := &survey.Confirm{
			Message: opts.Localizer.MustLocalize("cluster.serviceBinding.confirm.message"),
		}
		err = survey.AskOne(confirm, &shouldContinue)
		if err != nil {
			return err
		}

		if !shouldContinue {
			return nil
		}
	}

	// Check if connection exists
	status, err := service.CustomResourceExists(ctx, c, options.ServiceName)
	if status != http.StatusOK && err != nil {
		return opts.Localizer.MustLocalizeError("cluster.serviceBinding.serviceMissing.message")
	}

	// Execute binding
	err = service.BindCustomConnection(ctx, options.ServiceName, *options, clients)
	if err != nil {
		return err
	}

	opts.Logger.Info(icon.SuccessPrefix(), fmt.Sprintf(opts.Localizer.MustLocalize("cluster.serviceBinding.bindingSuccess"), options.ServiceName, options.AppName))
	return nil
}

func fetchAppNameFromCluster(ctx context.Context, resource schema.GroupVersionResource, clients *KubernetesClients, localizer localize.Localizer, ns string) (string, error) {
	list, err := clients.DynamicClient.Resource(resource).Namespace(ns).List(ctx, metav1.ListOptions{})
	if err != nil {
		return "", err
	}
	var appNames []string
	for _, d := range list.Items {
		name, found, err2 := unstructured.NestedString(d.Object, "metadata", "name")
		if err2 != nil || !found {
			continue
		}
		appNames = append(appNames, name)
	}

	if len(appNames) == 0 {
		return "", localizer.MustLocalizeError("cluster.serviceBinding.error.noapps")
	}

	prompt := &survey.Select{
		Message:  localizer.MustLocalize("cluster.serviceBinding.connect.survey.message"),
		Options:  appNames,
		PageSize: 10,
	}

	var selectedAppIndex int
	err = survey.AskOne(prompt, &selectedAppIndex)
	if err != nil {
		return "", err
	}
	return appNames[selectedAppIndex], nil
}

func client(localizer localize.Localizer) (*KubernetesClients, error) {
	kubeconfig := os.Getenv("KUBECONFIG")

	if kubeconfig == "" {
		home, _ := os.UserHomeDir()
		kubeconfig = filepath.Join(home, ".kube", "config")
	}

	_, err := os.Stat(kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", localizer.MustLocalize("cluster.kubernetes.error.configNotFoundError"), err)
	}

	restConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", localizer.MustLocalize("cluster.kubernetes.error.loadConfigError"), err)
	}

	if err != nil {
		return nil, fmt.Errorf("%v: %w", localizer.MustLocalize("cluster.kubernetes.error.loadConfigError"), err)
	}

	dynamicClient, err := dynamic.NewForConfig(restConfig)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", localizer.MustLocalize("cluster.kubernetes.error.loadConfigError"), err)
	}

	// Used for namespaces and general queries
	clientconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfig},
		&clientcmd.ConfigOverrides{ClusterInfo: clientcmdapi.Cluster{Server: ""}})

	return &KubernetesClients{dynamicClient, restConfig, &clientconfig}, nil
}
