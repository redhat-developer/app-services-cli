package cluster

import (
	"context"
	"crypto/rand"
	"fmt"
	"os"
	"path/filepath"

	"github.com/redhat-developer/app-services-cli/pkg/icon"
	"github.com/redhat-developer/service-binding-operator/apis/binding/v1alpha1"

	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/pkg/color"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

type KubernetesClients struct {
	dynamicClient dynamic.Interface
	restConfig    *rest.Config
	clientConfig  *clientcmd.ClientConfig
}

var (
	deploymentResource       = schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}
	deploymentConfigResource = schema.GroupVersionResource{Group: "apps.openshift.io", Version: "v1", Resource: "deploymentconfigs"}
)

type ServiceBindingOptions struct {
	ServiceName             string
	Namespace               string
	AppName                 string
	ForceCreationWithoutAsk bool
	BindingName             string
	BindAsFiles             bool
	DeploymentConfig        bool
}

func ExecuteServiceBinding(ctx context.Context, logger logging.Logger, localizer localize.Localizer, options *ServiceBindingOptions) error {
	clients, err := client(localizer)
	if err != nil {
		return err
	}
	ns := options.Namespace
	if ns == "" {
		ns, _, err = (*clients.clientConfig).Namespace()
		if err != nil {
			return err
		}
		logger.Info(localizer.MustLocalize("cluster.serviceBinding.namespaceInfo", localize.NewEntry("Namespace", color.Info(ns))))
	}

	var clusterResource schema.GroupVersionResource
	if options.DeploymentConfig {
		clusterResource = deploymentConfigResource
	} else {
		clusterResource = deploymentResource
	}
	// Get proper deployment
	if options.AppName == "" {
		options.AppName, err = fetchAppNameFromCluster(ctx, clusterResource, clients, localizer, ns)
		if err != nil {
			return err
		}
	} else {
		_, err = clients.dynamicClient.Resource(clusterResource).Namespace(ns).Get(ctx, options.AppName, metav1.GetOptions{})
		if err != nil {
			return err
		}
	}

	// Print desired action
	logger.Info(fmt.Sprintf(localizer.MustLocalize("cluster.serviceBinding.status.message"), options.ServiceName, options.AppName))

	if !options.ForceCreationWithoutAsk {
		var shouldContinue bool
		confirm := &survey.Confirm{
			Message: localizer.MustLocalize("cluster.serviceBinding.confirm.message"),
		}
		err = survey.AskOne(confirm, &shouldContinue)
		if err != nil {
			return err
		}

		if !shouldContinue {
			return nil
		}
	}

	// Check KafkaConnection
	_, err = clients.dynamicClient.Resource(AKCResource).Namespace(ns).Get(ctx, options.ServiceName, metav1.GetOptions{})
	if err != nil {
		return localizer.MustLocalizeError("cluster.serviceBinding.serviceMissing.message")
	}

	// Execute binding
	err = performBinding(ctx, options, ns, clients, logger, localizer)
	if err != nil {
		return err
	}

	logger.Info(icon.SuccessPrefix(), fmt.Sprintf(localizer.MustLocalize("cluster.serviceBinding.bindingSuccess"), options.ServiceName, options.AppName))
	return nil
}

func performBinding(ctx context.Context, options *ServiceBindingOptions, ns string, clients *KubernetesClients, logger logging.Logger, localizer localize.Localizer) error {
	serviceRef := v1alpha1.Service{
		NamespacedRef: v1alpha1.NamespacedRef{
			Ref: v1alpha1.Ref{
				Group:    AKCResource.Group,
				Version:  AKCResource.Version,
				Resource: AKCResource.Resource,
				Name:     options.ServiceName,
			},
		},
	}

	appRef := v1alpha1.Application{
		Ref: v1alpha1.Ref{
			Group:    deploymentResource.Group,
			Version:  deploymentResource.Version,
			Resource: deploymentResource.Resource,
			Name:     options.AppName,
		},
	}

	if options.BindingName == "" {
		randomValue := make([]byte, 2)
		_, err := rand.Read(randomValue)
		if err != nil {
			return err
		}
		options.BindingName = fmt.Sprintf("%v-%x", options.ServiceName, randomValue)
	}

	sb := &v1alpha1.ServiceBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      options.BindingName,
			Namespace: ns,
		},
		Spec: v1alpha1.ServiceBindingSpec{
			BindAsFiles: true,
			Services:    []v1alpha1.Service{serviceRef},
			Application: appRef,
		},
	}
	sb.SetGroupVersionKind(v1alpha1.GroupVersionKind)

	// Check of operator is installed
	_, err := clients.dynamicClient.Resource(v1alpha1.GroupVersionResource).Namespace(ns).
		List(ctx, metav1.ListOptions{Limit: 1})
	if err != nil {
		return fmt.Errorf("%s: %w", localizer.MustLocalizeError("cluster.serviceBinding.operatorMissing"), err)
	}

	return useOperatorForBinding(ctx, logger, localizer, sb, clients, ns)
}

func useOperatorForBinding(ctx context.Context, logger logging.Logger, localizer localize.Localizer, sb *v1alpha1.ServiceBinding, clients *KubernetesClients, ns string) error {
	logger.Info(localizer.MustLocalize("cluster.serviceBinding.usingOperator"))
	sbData, err := runtime.DefaultUnstructuredConverter.ToUnstructured(sb)
	if err != nil {
		return err
	}

	unstructuredSB := unstructured.Unstructured{Object: sbData}
	_, err = clients.dynamicClient.Resource(v1alpha1.GroupVersionResource).Namespace(ns).
		Create(ctx, &unstructuredSB, metav1.CreateOptions{})

	return err
}

func fetchAppNameFromCluster(ctx context.Context, resource schema.GroupVersionResource, clients *KubernetesClients, localizer localize.Localizer, ns string) (string, error) {
	list, err := clients.dynamicClient.Resource(resource).Namespace(ns).List(ctx, metav1.ListOptions{})
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
