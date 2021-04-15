package cluster

import (
	"context"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	sboContext "github.com/redhat-developer/service-binding-operator/pkg/reconcile/pipeline/context"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/rest"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"os"
	"path/filepath"
	"time"

	"github.com/redhat-developer/app-services-cli/internal/localizer"
	"github.com/redhat-developer/service-binding-operator/api/v1alpha1"
	"github.com/redhat-developer/service-binding-operator/pkg/reconcile/pipeline/builder"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
)

type KubernetesClients struct {
	dynamicClient dynamic.Interface
	restConfig    *rest.Config
	clientConfig  *clientcmd.ClientConfig
}

var deploymentResource = schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}

func ExecuteServiceBinding(logger logging.Logger, serviceName string, ns string, appName string) error {
	clients, err := client()
	if err != nil {
		return err
	}

	if ns == "" {
		ns, _, err = (*clients.clientConfig).Namespace()
		if err != nil {
			return err
		}
		logger.Info("Namespace not provided. Using ", ns)
	}
	logger.Debug("Binding arguments", serviceName, ns, appName)

	// Get proper deployment
	if appName == "" {
		logger.Debug("AppName missing. Looking for all apps on server")
		list, err := clients.dynamicClient.Resource(deploymentResource).Namespace(ns).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			logger.Debug("Cannot list")
			return err
		}
		var appNames []string
		for _, d := range list.Items {
			name, found, err := unstructured.NestedString(d.Object, "metadata", "name")
			if err != nil || !found {
				continue
			}
			appNames = append(appNames, name)
		}

		if len(appNames) == 0{
			return fmt.Errorf("Selected namespace has no deployments ")
		}

		prompt := &survey.Select{
			Message:  "Please select application you want to connect with",
			Options:  appNames,
			PageSize: 10,
		}

		var selectedAppIndex int
		err = survey.AskOne(prompt, &selectedAppIndex)
		if err != nil {
			return err
		}

		appName = appNames[selectedAppIndex]
	} else {
		_, err := clients.dynamicClient.Resource(deploymentResource).Namespace(ns).Get(context.TODO(), appName, metav1.GetOptions{})
		if err != nil {
			return err
		}
	}

	// Status and confirmation
	fmt.Printf("Binding '%v' with '%v' app \n", serviceName, appName)

	var shouldContinue bool
	confirm := &survey.Confirm{
		Message: "Do you want to continue?",
	}
	err = survey.AskOne(confirm, &shouldContinue)
	if err != nil {
		return err
	}

	if !shouldContinue {
		return nil
	}

	// Check KafkaConnection
	_, err = clients.dynamicClient.Resource(AKCResource).Namespace(ns).Get(context.TODO(), serviceName, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("Selected Kafka doesn't exist on the server. Please execute rhoas cluster connect first ")
	}

	// Execute binding
	serviceRef := v1alpha1.Service{
		NamespacedRef: v1alpha1.NamespacedRef{
			Ref: v1alpha1.Ref{
				Group:    AKCResource.Group,
				Version:  AKCResource.Version,
				Resource: AKCResource.Resource,
				Name:     serviceName,
			},
		},
	}

	appRef := v1alpha1.Application{
		Ref: v1alpha1.Ref{
			Group:    deploymentResource.Group,
			Version:  deploymentResource.Version,
			Resource: deploymentResource.Resource,
			Name:     appName,
		},
	}

	now := time.Now()
	sb := &v1alpha1.ServiceBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%v-%v", serviceName, now.Unix()),
			Namespace: ns,
		},
		Spec: v1alpha1.ServiceBindingSpec{
			BindAsFiles: true,
			Services:    []v1alpha1.Service{serviceRef},
			Application: &appRef,
		},
	}
	sb.SetGroupVersionKind(v1alpha1.GroupVersionKind)

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

	if err != nil {
		return err
	}
	fmt.Printf("Binding %v with %v app succeeded\n", serviceName, appName)
	return nil
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
