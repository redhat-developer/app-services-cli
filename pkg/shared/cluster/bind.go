package cluster

import (
	"crypto/rand"
	"fmt"

	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/color"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/shared/cluster/constants"
	"github.com/redhat-developer/app-services-cli/pkg/shared/cluster/kubeclient"
	"github.com/redhat-developer/app-services-cli/pkg/shared/cluster/v1alpha"

	bindv1alpha1 "github.com/redhat-developer/service-binding-operator/apis/binding/v1alpha1"

	"github.com/AlecAivazis/survey/v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func (c *KubernetesClusterAPIImpl) ExecuteServiceBinding(options *v1alpha.BindOperationOptions) error {
	clients := c.KubernetesClients
	cliEnv := c.CommandEnvironment

	if options.Namespace == "" {
		newNamespace, _, err := clients.ClientConfig.Namespace()
		if err != nil {
			return kubeclient.TranslatedKubernetesErrors(cliEnv, err)
		}
		options.Namespace = newNamespace
		cliEnv.Logger.Info(cliEnv.Localizer.MustLocalize("cluster.serviceBinding.namespaceInfo",
			localize.NewEntry("Namespace", color.Info(options.Namespace))))
	}

	var appResource schema.GroupVersionResource
	if options.DeploymentConfigEnabled {
		appResource = constants.DeploymentConfigResource
	} else {
		cliEnv.Logger.Info(cliEnv.Localizer.MustLocalize("cluster.serviceBinding.using.deployment"))
		appResource = constants.DeploymentResource
	}
	// Get deployment/app name if needed
	if options.AppName == "" {
		appName, err := fetchAppNameFromCluster(c, appResource, options.Namespace)
		if err != nil {
			return kubeclient.TranslatedKubernetesErrors(cliEnv, err)
		}
		options.AppName = appName
	} else {
		_, err := clients.DynamicClient.Resource(appResource).Namespace(options.Namespace).Get(cliEnv.Context, options.AppName, metav1.GetOptions{})
		if err != nil {
			return kubeclient.TranslatedKubernetesErrors(cliEnv, err)
		}
	}

	// Execute binding
	err := c.performBinding(options, options.Namespace, appResource)
	if err != nil {
		// Important to handle errors from API and handle missing prerequisites
		return kubeclient.TranslatedKubernetesErrors(cliEnv, err)
	}

	cliEnv.Logger.Info(icon.SuccessPrefix(), fmt.Sprintf(cliEnv.Localizer.MustLocalize("cluster.serviceBinding.bindingSuccess"), options.ServiceName, options.AppName))
	return nil
}

func (c *KubernetesClusterAPIImpl) performBinding(
	options *v1alpha.BindOperationOptions, ns string, appResource schema.GroupVersionResource) error {

	clients := c.KubernetesClients
	cliEnv := c.CommandEnvironment
	logger := cliEnv.Logger
	localizer := cliEnv.Localizer

	// Check of operator is installed
	installed, err := c.KubernetesClients.IsResourceAvailableOnCluster(&bindv1alpha1.GroupVersionResource, ns)
	if !installed {
		return fmt.Errorf("%s: %w", localizer.MustLocalizeError("cluster.serviceBinding.operatorMissing"), err)
	}
	if err != nil {
		return err
	}

	// Create service instance (it can be kafka/registry etc.)
	service, err := c.createServiceInstance(options.ServiceType)
	if err != nil {
		return err
	}

	// Build service metadata
	serviceMetadata, err := service.BuildServiceDetails(options.ServiceName, ns, options.IgnoreContext)
	if err != nil {
		return err
	}

	// Validate if service exist on the cluster
	_, err = clients.DynamicClient.Resource(serviceMetadata.GroupMetadata).Namespace(ns).Get(cliEnv.Context,
		serviceMetadata.Name, metav1.GetOptions{})
	if err != nil {
		return cliEnv.Localizer.MustLocalizeError("cluster.serviceBinding.serviceMissing.message")
	}

	cliEnv.Logger.Info(fmt.Sprintf(cliEnv.Localizer.MustLocalize("cluster.serviceBinding.status.message"),
		serviceMetadata.Name, options.AppName))

	if !options.ForceCreationWithoutAsk {
		var shouldContinue bool
		confirm := &survey.Confirm{
			Message: cliEnv.Localizer.MustLocalize("cluster.serviceBinding.confirm.message"),
		}
		errAsk := survey.AskOne(confirm, &shouldContinue)
		if errAsk != nil {
			return err
		}

		if !shouldContinue {
			return nil
		}
	}

	if options.ServiceName == "" {
		options.ServiceName = serviceMetadata.Name
	}

	sb, err := createBindingCR(options, serviceMetadata.GroupMetadata, appResource, ns)
	if err != nil {
		return err
	}
	logger.Info(localizer.MustLocalize("cluster.serviceBinding.usingOperator"))
	sbData, err := runtime.DefaultUnstructuredConverter.ToUnstructured(sb)
	if err != nil {
		return err
	}

	unstructuredSB := unstructured.Unstructured{Object: sbData}
	_, err = clients.DynamicClient.Resource(bindv1alpha1.GroupVersionResource).Namespace(ns).
		Create(cliEnv.Context, &unstructuredSB, metav1.CreateOptions{})

	return err
}

func createBindingCR(options *v1alpha.BindOperationOptions, serviceMetadata schema.GroupVersionResource,
	appResource schema.GroupVersionResource, namespace string) (*bindv1alpha1.ServiceBinding, error) {
	serviceRef := bindv1alpha1.Service{
		NamespacedRef: bindv1alpha1.NamespacedRef{
			Ref: bindv1alpha1.Ref{
				Group:    serviceMetadata.Group,
				Version:  serviceMetadata.Version,
				Resource: serviceMetadata.Resource,
				Name:     options.ServiceName,
			},
		},
	}
	appRef := bindv1alpha1.Application{
		Ref: bindv1alpha1.Ref{
			Group:    appResource.Group,
			Version:  appResource.Version,
			Resource: appResource.Resource,
			Name:     options.AppName,
		},
	}

	if options.BindingName == "" {
		randomValue := make([]byte, 2)
		_, errRand := rand.Read(randomValue)
		if errRand != nil {
			return nil, errRand
		}
		options.BindingName = fmt.Sprintf("%v-%x", options.ServiceName, randomValue)
	}

	sb := &bindv1alpha1.ServiceBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      options.BindingName,
			Namespace: namespace,
		},
		Spec: bindv1alpha1.ServiceBindingSpec{
			BindAsFiles: true,
			Services:    []bindv1alpha1.Service{serviceRef},
			Application: appRef,
		},
	}
	sb.SetGroupVersionKind(bindv1alpha1.GroupVersionKind)
	return sb, nil
}

func fetchAppNameFromCluster(c *KubernetesClusterAPIImpl, resource schema.GroupVersionResource, ns string) (string, error) {
	clients := c.KubernetesClients
	opts := c.CommandEnvironment
	list, err := clients.DynamicClient.Resource(resource).Namespace(ns).List(opts.Context, metav1.ListOptions{})
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
		return "", opts.Localizer.MustLocalizeError("cluster.serviceBinding.error.noapps")
	}

	prompt := &survey.Select{
		Message:  opts.Localizer.MustLocalize("cluster.serviceBinding.connect.survey.message"),
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
