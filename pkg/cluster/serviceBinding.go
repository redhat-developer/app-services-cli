package cluster

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/internal/localizer"
	"github.com/redhat-developer/app-services-cli/pkg/color"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	"github.com/redhat-developer/service-binding-operator/api/v1alpha1"
	"github.com/redhat-developer/service-binding-operator/pkg/reconcile/pipeline/builder"
	sboContext "github.com/redhat-developer/service-binding-operator/pkg/reconcile/pipeline/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"os"
	"path/filepath"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
)

type KubernetesClients struct {
	dynamicClient dynamic.Interface
	restConfig    *rest.Config
	clientConfig  *clientcmd.ClientConfig
}

var deploymentResource = schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}

type ServiceBindingOptions struct {
	ServiceName             string
	Namespace               string
	AppName                 string
	ForceCreationWithoutAsk bool
	ForceUseOperator        bool
	ForceUseSDK             bool
	BindingName             string
}

func ExecuteServiceBinding(logger logging.Logger, options *ServiceBindingOptions) error {
	clients, err := client()
	if err != nil {
		return err
	}
	ns := options.Namespace
	if ns == "" {
		ns, _, err = (*clients.clientConfig).Namespace()
		if err != nil {
			return err
		}
		logger.Info(localizer.Config{
			MessageID: "cluster.serviceBinding.namespaceInfo",
			TemplateData: map[string]interface{}{
				"Namespace": color.Info(ns),
			},
		})
	}

	// Get proper deployment
	if options.AppName == "" {
		options.AppName, err = fetchAppNameFromCluster(clients, ns)
		if err != nil {
			return err
		}
	} else {
		_, err = clients.dynamicClient.Resource(deploymentResource).Namespace(ns).Get(context.TODO(), options.AppName, metav1.GetOptions{})
		if err != nil {
			return err
		}
	}

	// Print desired action
	logger.Info(fmt.Sprintf(localizer.MustLocalizeFromID("cluster.serviceBinding.status.message"), options.ServiceName, options.AppName))

	if !options.ForceCreationWithoutAsk {
		var shouldContinue bool
		confirm := &survey.Confirm{
			Message: localizer.MustLocalizeFromID("cluster.serviceBinding.confirm.message"),
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
	_, err = clients.dynamicClient.Resource(AKCResource).Namespace(ns).Get(context.TODO(), options.ServiceName, metav1.GetOptions{})
	if err != nil {
		return errors.New(localizer.MustLocalizeFromID("cluster.serviceBinding.serviceMissing.message"))
	}

	// Execute binding
	err = performBinding(options, ns, clients, logger)
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf(localizer.MustLocalizeFromID("cluster.serviceBinding.bindingSuccess"), options.ServiceName, options.AppName))
	return nil
}

func performBinding(options *ServiceBindingOptions, ns string, clients *KubernetesClients, logger logging.Logger) error {
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
		randomValue := make([]byte, 4)
		_, err := rand.Read(randomValue)
		if err != nil {
			return err
		}
		options.BindingName = fmt.Sprintf("%v-%x", options.ServiceName, randomValue)
	}

	sb := &v1alpha1.ServiceBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%v-%v", options.ServiceName, options.BindingName),
			Namespace: options.Namespace,
		},
		Spec: v1alpha1.ServiceBindingSpec{
			BindAsFiles: true,
			Services:    []v1alpha1.Service{serviceRef},
			Application: &appRef,
		},
	}
	sb.SetGroupVersionKind(v1alpha1.GroupVersionKind)

	if options.ForceUseSDK {
		return useSDKForBinding(clients, sb)
	}

	// Check of operator is installed
	_, err := clients.dynamicClient.Resource(v1alpha1.GroupVersionResource).Namespace(ns).
		List(context.TODO(), metav1.ListOptions{Limit: 1})

	if err != nil {
		if options.ForceUseOperator {
			return errors.New(localizer.MustLocalizeFromID("cluster.serviceBinding.operatorMissing") + err.Error())
		}
		logger.Debug("Service binding Operator not available. Will use SDK option for binding")
		return useSDKForBinding(clients, sb)
	}

	logger.Debug("Using Operator to perform binding")
	return useOperatorForBinding(logger, sb, clients, ns)

}

func useOperatorForBinding(logger logging.Logger, sb *v1alpha1.ServiceBinding, clients *KubernetesClients, ns string) error {
	logger.Info(localizer.MustLocalizeFromID("cluster.serviceBinding.usingOperator"))
	sbData, err := runtime.DefaultUnstructuredConverter.ToUnstructured(sb)
	if err != nil {
		return err
	}

	unstructuredSB := unstructured.Unstructured{Object: sbData}
	_, err = clients.dynamicClient.Resource(v1alpha1.GroupVersionResource).Namespace(ns).
		Create(context.TODO(), &unstructuredSB, metav1.CreateOptions{})

	return err
}

func useSDKForBinding(clients *KubernetesClients, sb *v1alpha1.ServiceBinding) error {
	restMapper, err := apiutil.NewDynamicRESTMapper(clients.restConfig)
	if err != nil {
		return err
	}
	typeLookup := sboContext.ResourceLookup(restMapper)

	p := builder.DefaultBuilder.WithContextProvider(sboContext.Provider(clients.dynamicClient, typeLookup)).Build()

	retry, err := p.Process(sb)

	if retry {
		_, err = p.Process(sb)
	}
	return err
}

func fetchAppNameFromCluster(clients *KubernetesClients, ns string) (string, error) {
	list, err := clients.dynamicClient.Resource(deploymentResource).Namespace(ns).List(context.TODO(), metav1.ListOptions{})
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
		return "", fmt.Errorf("Selected namespace has no deployments ")
	}

	prompt := &survey.Select{
		Message:  localizer.MustLocalizeFromID("cluster.serviceBinding.connect.survey.message"),
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

func client() (*KubernetesClients, error) {
	kubeconfig := os.Getenv("KUBECONFIG")

	if kubeconfig == "" {
		home, _ := os.UserHomeDir()
		kubeconfig = filepath.Join(home, ".kube", "config")
	}

	_, err := os.Stat(kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", localizer.MustLocalizeFromID("cluster.kubernetes.error.configNotFoundError"), err)
	}

	restConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", localizer.MustLocalizeFromID("cluster.kubernetes.error.loadConfigError"), err)
	}

	if err != nil {
		return nil, fmt.Errorf("%v: %w", localizer.MustLocalizeFromID("cluster.kubernetes.error.loadConfigError"), err)
	}

	dynamicClient, err := dynamic.NewForConfig(restConfig)

	if err != nil {
		return nil, fmt.Errorf("%v: %w", localizer.MustLocalizeFromID("cluster.kubernetes.error.loadConfigError"), err)
	}

	// Used for namespaces and general queries
	clientconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfig},
		&clientcmd.ConfigOverrides{ClusterInfo: clientcmdapi.Cluster{Server: ""}})

	return &KubernetesClients{dynamicClient, restConfig, &clientconfig}, nil
}
