package cluster

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	kafkaUtil "github.com/redhat-developer/app-services-cli/pkg/kafka"
	"github.com/redhat-developer/app-services-cli/pkg/serviceregistry"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"

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
	clientset          *kubernetes.Clientset
	clientconfig       clientcmd.ClientConfig
	dynamicClient      dynamic.Interface
	kubeconfigLocation string
}

/*  #nosec */
var tokenSecretName = "rh-cloud-services-accesstoken-cli"

/*  #nosec */
var serviceAccountSecretName = "rh-cloud-services-service-account"

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
func (c *KubernetesCluster) Connect(ctx context.Context, cmdOptions *ConnectArguments, opts Options) error {

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
		localize.NewEntry("InstanceName", color.Info("random name")),
		localize.NewEntry("Namespace", color.Info(currentNamespace)),
		localize.NewEntry("RegistryInstanceName", color.Info("random-name")),
		localize.NewEntry("ServiceAccountSecretName", color.Info(serviceAccountSecretName))))

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

	var service Service

	switch cmdOptions.SelectedService {

	case "kafka":
		service = &Kafka{}
		err = c.createConnectionCustomResource(ctx, currentNamespace, cmdOptions.SelectedServiceID, service, opts)
		if err != nil {
			return err
		}
	case "service-registry":
		service = &ServiceRegistry{}
		err = c.createConnectionCustomResource(ctx, currentNamespace, cmdOptions.SelectedServiceID, service, opts)
		if err != nil {
			return err
		}
	case "":
		var selectedKafkaInstance string
		var selectedRegistryInstance string

		cfg, err := opts.Config.Load()
		if err != nil {
			return err
		}

		if cfg.Services.Kafka == nil || cmdOptions.IgnoreContext {
			// nolint
			selectedKafka, err := kafkaUtil.InteractiveSelect(ctx, opts.Connection, opts.Logger, opts.Localizer)
			if err != nil {
				return err
			}
			if selectedKafka == nil {
				return nil
			}
			selectedKafkaInstance = selectedKafka.GetId()
		} else {
			selectedKafkaInstance = cfg.Services.Kafka.ClusterID
		}

		service = &Kafka{}

		err = c.createConnectionCustomResource(ctx, currentNamespace, selectedKafkaInstance, service, opts)
		if err != nil {
			return err
		}

		if cfg.Services.ServiceRegistry == nil || cmdOptions.IgnoreContext {
			// nolint
			selectedServiceRegistry, err := serviceregistry.InteractiveSelect(ctx, opts.Connection, opts.Logger)
			if err != nil {
				return err
			}
			if selectedServiceRegistry == nil {
				return nil
			}
			selectedRegistryInstance = selectedServiceRegistry.GetId()
		} else {
			selectedRegistryInstance = cfg.Services.ServiceRegistry.InstanceID
		}

		service = &ServiceRegistry{}

		err = c.createConnectionCustomResource(ctx, currentNamespace, selectedRegistryInstance, service, opts)
		if err != nil {
			return err
		}

	}

	return nil
}

// createConnectionCustomResource
func (c *KubernetesCluster) createConnectionCustomResource(ctx context.Context, namespace string, serviceID string, service Service, opts Options) error {

	err := service.ResourceExists(ctx, c, namespace, serviceID, opts)
	if err != nil {
		return err
	}

	err = service.CreateResource(ctx, c, namespace, serviceID, opts)

	if err != nil {
		return err
	}

	return nil
}

// IsRhoasOperatorAvailableOnCluster checks the cluster to see if a KafkaConnection CRD is installed
func (c *KubernetesCluster) IsRhoasOperatorAvailableOnCluster(ctx context.Context) (bool, error) {
	return IsKCInstalledOnCluster(ctx, c)
}

func (c *KubernetesCluster) createTokenSecretIfNeeded(ctx context.Context, namespace string, connectOpts *ConnectArguments, opts Options) error {
	_, err := c.clientset.CoreV1().Secrets(namespace).Get(context.TODO(), tokenSecretName, metav1.GetOptions{})
	if err == nil {
		opts.Logger.Info(opts.Localizer.MustLocalize("cluster.kubernetes.tokensecret.log.info.found"), tokenSecretName)
		return nil
	}

	if connectOpts.OfflineAccessToken == "" && !opts.IO.CanPrompt() {
		return errors.New(opts.Localizer.MustLocalize("flag.error.requiredWhenNonInteractive", localize.NewEntry("Flag", "token")))
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
			Name:      tokenSecretName,
			Namespace: namespace,
		},
		StringData: map[string]string{
			"value": connectOpts.OfflineAccessToken,
		},
	}

	_, err = c.clientset.CoreV1().Secrets(namespace).Create(ctx, secret, metav1.CreateOptions{})
	tokenSecretNameTmplEntry := localize.NewEntry("Name", tokenSecretName)
	if err != nil {
		return fmt.Errorf("%v: %w", opts.Localizer.MustLocalize("cluster.kubernetes.createTokenSecret.log.info.createFailed", tokenSecretNameTmplEntry), err)
	}

	opts.Logger.Info(opts.Localizer.MustLocalize("cluster.kubernetes.createTokenSecret.log.info.createSuccess", tokenSecretNameTmplEntry))

	return nil
}

// createSecret creates a new secret to store the SASL/PLAIN credentials from the service account
func (c *KubernetesCluster) createServiceAccountSecretIfNeeded(ctx context.Context, namespace string, opts Options) error {
	_, err := c.clientset.CoreV1().Secrets(namespace).Get(context.TODO(), serviceAccountSecretName, metav1.GetOptions{})
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
			Name:      serviceAccountSecretName,
			Namespace: namespace,
		},
		StringData: map[string]string{
			"client-id":     serviceAcct.GetClientId(),
			"client-secret": serviceAcct.GetClientSecret(),
		},
	}

	createdSecret, err := c.clientset.CoreV1().Secrets(namespace).Create(ctx, secret, metav1.CreateOptions{})
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

func (c *KubernetesCluster) makeKubernetesGetRequest(ctx context.Context, path string, serviceName string, localizer localize.Localizer) error {
	var status int

	data := c.clientset.
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

func (c *KubernetesCluster) makeKubernetesPostRequest(ctx context.Context, path string, serviceName string, crJSON []byte) error {

	data := c.clientset.
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
