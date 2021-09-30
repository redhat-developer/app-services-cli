package cluster

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"

	"github.com/redhat-developer/app-services-cli/pkg/cluster/constants"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/localize"

	"k8s.io/client-go/dynamic"

	"github.com/golang-jwt/jwt/v4"

	"github.com/redhat-developer/app-services-cli/internal/build"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"

	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/pkg/color"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
)

// Options object
type Options struct {
	Connection connection.Connection
	Config     config.IConfig
	Logger     logging.Logger
	IO         *iostreams.IOStreams
	Localizer  localize.Localizer
}

type KubernetesCluster struct {
	Clientset          *kubernetes.Clientset
	clientconfig       clientcmd.ClientConfig
	DynamicClient      dynamic.Interface
	kubeconfigLocation string
}

// CustomResourceOptions object contains the data required to create a custom connection
type CustomResourceOptions struct {
	Path        string
	CRName      string
	ServiceName string
	CRJSON      []byte
	Resource    schema.GroupVersionResource
}

// NewKubernetesClusterConnection configures and connects to a Kubernetes cluster
func NewKubernetesClusterConnection(connection connection.Connection,
	config config.IConfig,
	logger logging.Logger,
	kubeconfig string,
	io *iostreams.IOStreams, localizer localize.Localizer) (Cluster, error) {
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

// nolint:funlen
// Connect connects a remote Kafka instance to the Kubernetes cluster
func (c *KubernetesCluster) Connect(ctx context.Context, cmdOptions *ConnectArguments, connection CustomConnection, opts Options) error {

	var currentNamespace string
	var err error
	if cmdOptions.Namespace != "" {
		currentNamespace = cmdOptions.Namespace
	} else {
		currentNamespace, err = c.CurrentNamespace()
		if err != nil {
			return err
		}
	}

	// print status
	opts.Logger.Info(opts.Localizer.MustLocalize("cluster.kubernetes.log.info.statusMessage"))

	opts.Logger.Info(opts.Localizer.MustLocalize("cluster.kubernetes.statusInfo",
		localize.NewEntry("ServiceType", color.Info(cmdOptions.SelectedService)),
		localize.NewEntry("ServiceID", color.Info(cmdOptions.SelectedServiceID)),
		localize.NewEntry("Namespace", color.Info(currentNamespace)),
		localize.NewEntry("ServiceAccountSecretName", color.Info(constants.ServiceAccountSecretName))))

	if cmdOptions.ForceCreationWithoutAsk == false {
		var shouldContinue bool
		confirm := &survey.Confirm{
			Message: opts.Localizer.MustLocalize("cluster.kubernetes.connect.input.confirm.message"),
		}
		err = survey.AskOne(confirm, &shouldContinue)
		if err != nil {
			return err
		}

		if !shouldContinue {
			opts.Logger.Debug(opts.Localizer.MustLocalize("cluster.kubernetes.connect.log.debug.cancellingConnect"))
			return nil
		}
	}

	// Token with auth for operator to pick
	err = c.createTokenSecretIfNeeded(ctx, currentNamespace, cmdOptions, opts)
	if err != nil {
		return err
	}

	err = c.createServiceAccountSecretIfNeeded(ctx, currentNamespace, opts)
	if err != nil {
		return err
	}

	err = connection.CustomResourceExists(ctx, c, cmdOptions.SelectedServiceID)
	if err != nil {
		return err
	}
	err = connection.CreateCustomResource(ctx, c, cmdOptions.SelectedServiceID)
	if err != nil {
		return err
	}

	return nil
}

// IsRhoasOperatorAvailableOnCluster checks the cluster for availability of RHOAS operator in cluster
func (c *KubernetesCluster) IsRhoasOperatorAvailableOnCluster(ctx context.Context) (bool, error) {
	installed, err := IsKCInstalledOnCluster(ctx, c)
	if !installed {
		return installed, err
	}

	installed, err = IsSRCInstalledOnCluster(ctx, c)
	if !installed {
		return installed, err
	}

	installed, err = IsSBOInstalledOnCluster(ctx, c)

	return installed, err
}

func (c *KubernetesCluster) createTokenSecretIfNeeded(ctx context.Context, namespace string, connectOpts *ConnectArguments, opts Options) error {
	_, err := c.Clientset.CoreV1().Secrets(namespace).Get(context.TODO(), constants.TokenSecretName, metav1.GetOptions{})
	if err == nil {
		opts.Logger.Info(opts.Localizer.MustLocalize("cluster.kubernetes.tokensecret.log.info.found"), constants.TokenSecretName)
		return nil
	}

	if connectOpts.OfflineAccessToken == "" && !opts.IO.CanPrompt() {
		return opts.Localizer.MustLocalizeError("flag.error.requiredWhenNonInteractive", localize.NewEntry("Flag", "token"))
	}

	if connectOpts.OfflineAccessToken == "" {
		apiTokenInput := &survey.Input{
			Message: opts.Localizer.MustLocalize("cluster.common.flag.offline.token.description", localize.NewEntry("OfflineTokenURL", build.OfflineTokenURL)),
		}
		surveyErr := survey.AskOne(apiTokenInput, &connectOpts.OfflineAccessToken)
		if surveyErr != nil {
			return err
		}
	}
	parser := new(jwt.Parser)
	_, _, err = parser.ParseUnverified(connectOpts.OfflineAccessToken, jwt.MapClaims{})
	if err != nil {
		return err
	}

	// Create secret type
	secret := &apiv1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      constants.TokenSecretName,
			Namespace: namespace,
		},
		StringData: map[string]string{
			"value": connectOpts.OfflineAccessToken,
		},
	}

	_, err = c.Clientset.CoreV1().Secrets(namespace).Create(ctx, secret, metav1.CreateOptions{})
	tokenSecretNameTmplEntry := localize.NewEntry("Name", constants.TokenSecretName)
	if err != nil {
		return fmt.Errorf("%v: %w", opts.Localizer.MustLocalize("cluster.kubernetes.createTokenSecret.log.info.createFailed", tokenSecretNameTmplEntry), err)
	}

	opts.Logger.Info(opts.Localizer.MustLocalize("cluster.kubernetes.createTokenSecret.log.info.createSuccess", tokenSecretNameTmplEntry))

	return nil
}

// createSecret creates a new secret to store the SASL/PLAIN credentials from the service account
func (c *KubernetesCluster) createServiceAccountSecretIfNeeded(ctx context.Context, namespace string, opts Options) error {
	_, err := c.Clientset.CoreV1().Secrets(namespace).Get(context.TODO(), constants.ServiceAccountSecretName, metav1.GetOptions{})
	if err == nil {
		opts.Logger.Info(opts.Localizer.MustLocalize("cluster.kubernetes.serviceaccountsecret.log.info.exist"))
		return nil
	}

	serviceAcct, err := c.createServiceAccount(ctx, opts)
	if err != nil {
		return err
	}

	secret := &apiv1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      constants.ServiceAccountSecretName,
			Namespace: namespace,
		},
		StringData: map[string]string{
			"client-id":     serviceAcct.GetClientId(),
			"client-secret": serviceAcct.GetClientSecret(),
		},
	}

	createdSecret, err := c.Clientset.CoreV1().Secrets(namespace).Create(ctx, secret, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("%v: %w", opts.Localizer.MustLocalize("cluster.kubernetes.serviceaccountsecret.error.createError"), err)
	}

	opts.Logger.Info(opts.Localizer.MustLocalize("cluster.kubernetes.createSASecret.log.info.createSuccess", localize.NewEntry("Name", createdSecret.Name)))

	return nil
}

// createServiceAccount creates a service account
func (c *KubernetesCluster) createServiceAccount(ctx context.Context, opts Options) (*kafkamgmtclient.ServiceAccount, error) {
	t := time.Now()

	api := opts.Connection.API()
	serviceAcct := &kafkamgmtclient.ServiceAccountRequest{Name: fmt.Sprintf("rhoascli-%v", t.Unix())}
	req := api.ServiceAccount().CreateServiceAccount(ctx)
	req = req.ServiceAccountRequest(*serviceAcct)
	serviceAcctRes, httpRes, err := req.Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}
	if err != nil {
		return nil, fmt.Errorf("%v: %w", opts.Localizer.MustLocalize("cluster.kubernetes.createServiceAccount.error.createError"), err)
	}

	return &serviceAcctRes, nil
}

func (c *KubernetesCluster) MakeKubernetesGetRequest(ctx context.Context, path string, serviceName string, localizer localize.Localizer) error {
	var status int

	data := c.Clientset.
		RESTClient().
		Get().
		AbsPath(path, serviceName).
		Do(ctx)

	if data.StatusCode(&status); status == http.StatusNotFound {
		return nil
	}

	if data.Error() == nil {
		return fmt.Errorf("%v: %s", localizer.MustLocalize("cluster.kubernetes.checkIfConnectionExist.existError"), serviceName)
	}

	return nil
}

func (c *KubernetesCluster) MakeKubernetesPostRequest(ctx context.Context, path string, serviceName string, crJSON []byte) error {

	data := c.Clientset.
		RESTClient().
		Post().
		AbsPath(path, serviceName).
		Body(crJSON).
		Do(ctx)

	if data.Error() != nil {
		return data.Error()
	}

	return nil
}

// CreateResource creates a CustomResource connection in the cluster
func (c *KubernetesCluster) CreateResource(ctx context.Context, resourceOpts *CustomResourceOptions, opts Options) error {

	namespace, err := c.CurrentNamespace()
	if err != nil {
		return err
	}

	err = c.MakeKubernetesPostRequest(ctx, resourceOpts.Path, resourceOpts.ServiceName, resourceOpts.CRJSON)

	if err != nil {
		return err
	}

	opts.Logger.Info(opts.Localizer.MustLocalize("cluster.kubernetes.createCR.log.info.customResourceCreated", localize.NewEntry("Resource", resourceOpts.CRName), localize.NewEntry("Name", resourceOpts.ServiceName)))

	w, err := c.DynamicClient.Resource(resourceOpts.Resource).Namespace(namespace).Watch(context.TODO(), metav1.ListOptions{
		FieldSelector: fields.OneTermEqualSelector("metadata.name", resourceOpts.ServiceName).String(),
	})
	if err != nil {
		return err
	}

	return watchCustomResourceStatus(w, opts, resourceOpts.CRName)
}

// ResourceExists checks if a CustomResource connection already exists in the cluster
func (c *KubernetesCluster) ResourceExists(ctx context.Context, path string, serviceName string, opts Options) error {

	err := c.MakeKubernetesGetRequest(ctx, path, serviceName, opts.Localizer)

	return err
}

func watchCustomResourceStatus(w watch.Interface, opts Options, crName string) error {
	for {
		select {
		case event := <-w.ResultChan():
			if event.Type == watch.Modified {
				unstructuredObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(event.Object)
				if err != nil {
					return err
				}
				conditions, found, err := unstructured.NestedSlice(unstructuredObj, "status", "conditions")
				if err != nil {
					return err
				}

				if found {
					for _, condition := range conditions {
						typedCondition, ok := condition.(map[string]interface{})
						if !ok {
							return fmt.Errorf(opts.Localizer.MustLocalize("cluster.kubernetes.watchForConnectionStatus.error.format"), typedCondition)
						}
						if typedCondition["type"].(string) == "Finished" {
							if typedCondition["status"].(string) == "False" {
								w.Stop()
								return fmt.Errorf(opts.Localizer.MustLocalize("cluster.kubernetes.watchForResourceStatus.error.status", localize.NewEntry("Resource", crName)), typedCondition["message"])
							}
							if typedCondition["status"].(string) == "True" {
								opts.Logger.Info(opts.Localizer.MustLocalize("cluster.kubernetes.watchForResourceStatus.log.info.success", localize.NewEntry("Resource", crName)))

								w.Stop()
								return nil
							}
						}
					}
					w.Stop()
				}
			}

		case <-time.After(60 * time.Second):
			w.Stop()
			return fmt.Errorf(opts.Localizer.MustLocalize("cluster.kubernetes.watchForResourceStatus.error.timeout", localize.NewEntry("Resource", crName)))
		}
	}
}
