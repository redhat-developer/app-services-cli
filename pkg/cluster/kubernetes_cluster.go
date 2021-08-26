package cluster

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redhat-developer/app-services-cli/pkg/icon"
	"os"
	"path/filepath"
	"time"

	"github.com/redhat-developer/app-services-cli/pkg/api/kas"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/kafka"
	registryPkg "github.com/redhat-developer/app-services-cli/pkg/cluster/serviceregistry"
	kafkaUtil "github.com/redhat-developer/app-services-cli/pkg/kafka"
	"github.com/redhat-developer/app-services-cli/pkg/kafka/kafkaerr"
	"github.com/redhat-developer/app-services-cli/pkg/serviceregistry"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
	srsmgmtv1 "github.com/redhat-developer/app-services-sdk-go/registrymgmt/apiv1/client"

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

// KubernetesCluster is a type which represents a Kubernetes cluster
type KubernetesCluster struct {
	connection connection.Connection
	config     config.IConfig
	logger     logging.Logger

	clientset          *kubernetes.Clientset
	clientconfig       clientcmd.ClientConfig
	dynamicClient      dynamic.Interface
	io                 *iostreams.IOStreams
	localizer          localize.Localizer
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
		connection,
		config,
		logger,
		clientset,
		clientconfig,
		dynamicClient,
		io,
		localizer,
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
func (c *KubernetesCluster) Connect(ctx context.Context, cmdOptions *ConnectArguments) error {

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
	c.logger.Info(c.localizer.MustLocalize("cluster.kubernetes.log.info.statusMessage"))

	c.logger.Info(c.localizer.MustLocalize("cluster.kubernetes.statusInfo",
		localize.NewEntry("InstanceName", color.Info("random name")),
		localize.NewEntry("Namespace", color.Info(currentNamespace)),
		localize.NewEntry("RegistryInstanceName", color.Info("random-name")),
		localize.NewEntry("ServiceAccountSecretName", color.Info(serviceAccountSecretName))))

	if cmdOptions.ForceCreationWithoutAsk == false {
		var shouldContinue bool
		confirm := &survey.Confirm{
			Message: c.localizer.MustLocalize("cluster.kubernetes.connect.input.confirm.message"),
		}
		err = survey.AskOne(confirm, &shouldContinue)
		if err != nil {
			return err
		}

		if !shouldContinue {
			c.logger.Debug(c.localizer.MustLocalize("cluster.kubernetes.connect.log.debug.cancellingConnect"))
			return nil
		}
	}

	// Token with auth for operator to pick
	err = c.createTokenSecretIfNeeded(ctx, currentNamespace, cmdOptions)
	if err != nil {
		return err
	}

	err = c.createServiceAccountSecretIfNeeded(ctx, currentNamespace)
	if err != nil {
		return err
	}

	switch cmdOptions.SelectedService {
	case "kafka":
		err = c.checkAndCreateKafkaConnectionCustomResource(ctx, currentNamespace, cmdOptions.SelectedServiceID)
		if err != nil {
			return err
		}
	case "service-registry":
		err = c.checkAndCreateServiceRegistryConnectionCustomResource(ctx, currentNamespace, cmdOptions.SelectedServiceID)
		if err != nil {
			return err
		}
	case "":
		var selectedKafkaInstance string
		var selectedRegistryInstance string

		cfg, err := c.config.Load()
		if err != nil {
			return err
		}

		if cfg.Services.Kafka == nil || cmdOptions.IgnoreContext {
			// nolint
			selectedKafka, err := kafkaUtil.InteractiveSelect(c.connection, c.logger)
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

		err = c.checkAndCreateKafkaConnectionCustomResource(ctx, currentNamespace, selectedKafkaInstance)
		if err != nil {
			return err
		}

		if cfg.Services.ServiceRegistry == nil || cmdOptions.IgnoreContext {
			// nolint
			selectedServiceRegistry, err := serviceregistry.InteractiveSelect(c.connection, c.logger)
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

		err = c.checkAndCreateServiceRegistryConnectionCustomResource(ctx, currentNamespace, selectedRegistryInstance)
		if err != nil {
			return err
		}

	}

	return nil
}

func (c *KubernetesCluster) checkAndCreateKafkaConnectionCustomResource(ctx context.Context, namespace string, kafkaID string) error {

	api := c.connection.API()

	kafkaInstance, _, err := api.Kafka().GetKafkaById(ctx, kafkaID).Execute()
	if kas.IsErr(err, kas.ErrorNotFound) {
		return kafkaerr.NotFoundByIDError(kafkaID)
	}

	if err != nil {
		return err
	}
	err = CheckIfKafkaConnectionExists(ctx, c, namespace, kafkaInstance.GetName())
	if err != nil {
		return err
	}

	err = c.createKafkaConnectionCustomResource(ctx, namespace, &kafkaInstance)
	if err != nil {
		return err
	}

	return nil
}

// createKafkaConnectionCustomResource creates a new "KafkaConnection" CR
func (c *KubernetesCluster) createKafkaConnectionCustomResource(ctx context.Context, namespace string, kafkaInstance *kafkamgmtclient.KafkaRequest) error {
	crName := kafkaInstance.GetName()
	kafkaID := kafkaInstance.GetId()

	kafkaConnectionCR := kafka.CreateKCObject(crName, namespace, kafkaID)

	crJSON, err := json.Marshal(kafkaConnectionCR)
	if err != nil {
		return fmt.Errorf("%v: %w", c.localizer.MustLocalize("cluster.kubernetes.createKafkaCR.error.marshalError"), err)
	}

	data := c.clientset.RESTClient().
		Post().
		AbsPath(kafka.GetKafkaConnectionsAPIURL(namespace)).
		Body(crJSON).
		Do(ctx)

	if data.Error() != nil {
		return data.Error()
	}

	c.logger.Info(c.localizer.MustLocalize("cluster.kubernetes.createKafkaCR.log.info.customResourceCreated", localize.NewEntry("Name", crName)))

	return watchForKafkaStatus(ctx, c, crName, namespace)
}

func (c *KubernetesCluster) checkAndCreateServiceRegistryConnectionCustomResource(ctx context.Context, namespace string, registryID string) error {

	api := c.connection.API()

	registryInstance, _, err := serviceregistry.GetServiceRegistryByID(ctx, api.ServiceRegistryMgmt(), registryID)
	if err != nil {
		return err
	}

	err = CheckIfRegistryConnectionExists(ctx, c, namespace, registryInstance.GetName())
	if err != nil {
		return err
	}

	err = c.createServiceRegistryCustomResource(ctx, namespace, registryInstance)
	if err != nil {
		return err
	}

	return nil
}

// createServiceRegistryCustomResource creates a new "ServiceRegistryConnection" CR
func (c *KubernetesCluster) createServiceRegistryCustomResource(ctx context.Context, namespace string, registryInstance *srsmgmtv1.RegistryRest) error {
	crName := registryInstance.GetName()
	registryId := registryInstance.GetId()

	serviceRegistryCR := registryPkg.CreateSRObject(crName, namespace, registryId)

	crJSON, err := json.Marshal(serviceRegistryCR)
	if err != nil {
		return fmt.Errorf("%v: %w", c.localizer.MustLocalize("cluster.kubernetes.createRegistryCR.error.marshalError"), err)
	}

	data := c.clientset.RESTClient().
		Post().
		AbsPath(registryPkg.GetServiceRegistryAPIURL(namespace)).
		Body(crJSON).
		Do(ctx)

	if data.Error() != nil {
		return data.Error()
	}

	c.logger.Info(c.localizer.MustLocalize("cluster.kubernetes.createRegistryCR.log.info.customResourceCreated", localize.NewEntry("Name", crName)))
	// c.logger.Info("KafkaConnection resource some-registry-name has been created'")

	return watchForServiceRegistryStatus(c, crName, namespace)
}

// IsRhoasOperatorAvailableOnCluster checks the cluster to see if a KafkaConnection CRD is installed
func (c *KubernetesCluster) IsRhoasOperatorAvailableOnCluster(ctx context.Context) (bool, error) {
	return IsKCInstalledOnCluster(ctx, c)
}

func (c *KubernetesCluster) createTokenSecretIfNeeded(ctx context.Context, namespace string, opts *ConnectArguments) error {
	_, err := c.clientset.CoreV1().Secrets(namespace).Get(ctx, tokenSecretName, metav1.GetOptions{})
	if err == nil {
		c.logger.Info(c.localizer.MustLocalize("cluster.kubernetes.tokensecret.log.info.found"), tokenSecretName)
		return nil
	}

	if opts.OfflineAccessToken == "" && !c.io.CanPrompt() {
		return c.localizer.MustLocalizeError("flag.error.requiredWhenNonInteractive", localize.NewEntry("Flag", "token"))
	}

	if opts.OfflineAccessToken == "" {
		apiTokenInput := &survey.Input{
			Message: c.localizer.MustLocalize("cluster.common.flag.offline.token.description", localize.NewEntry("OfflineTokenURL", build.OfflineTokenURL)),
		}
		surveyErr := survey.AskOne(apiTokenInput, &opts.OfflineAccessToken)
		if surveyErr != nil {
			return err
		}
	}
	parser := new(jwt.Parser)
	_, _, err = parser.ParseUnverified(opts.OfflineAccessToken, jwt.MapClaims{})
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
			"value": opts.OfflineAccessToken,
		},
	}

	_, err = c.clientset.CoreV1().Secrets(namespace).Create(ctx, secret, metav1.CreateOptions{})
	tokenSecretNameTmplEntry := localize.NewEntry("Name", tokenSecretName)
	if err != nil {
		return fmt.Errorf("%v: %w", c.localizer.MustLocalize("cluster.kubernetes.createTokenSecret.log.info.createFailed", tokenSecretNameTmplEntry), err)
	}

	c.logger.Info(icon.SuccessPrefix(), c.localizer.MustLocalize("cluster.kubernetes.createTokenSecret.log.info.createSuccess", tokenSecretNameTmplEntry))

	return nil
}

// createSecret creates a new secret to store the SASL/PLAIN credentials from the service account
func (c *KubernetesCluster) createServiceAccountSecretIfNeeded(ctx context.Context, namespace string) error {
	_, err := c.clientset.CoreV1().Secrets(namespace).Get(ctx, serviceAccountSecretName, metav1.GetOptions{})
	if err == nil {
		c.logger.Info(c.localizer.MustLocalize("cluster.kubernetes.serviceaccountsecret.log.info.exist"))
		return nil
	}

	serviceAcct, err := c.createServiceAccount(ctx)
	if err != nil {
		return err
	}

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
		return fmt.Errorf("%v: %w", c.localizer.MustLocalize("cluster.kubernetes.serviceaccountsecret.error.createError"), err)
	}

	c.logger.Info(icon.SuccessPrefix(), c.localizer.MustLocalize("cluster.kubernetes.createSASecret.log.info.createSuccess", localize.NewEntry("Name", createdSecret.Name)))

	return nil
}

// createServiceAccount creates a service account
func (c *KubernetesCluster) createServiceAccount(ctx context.Context) (*kafkamgmtclient.ServiceAccount, error) {
	t := time.Now()

	api := c.connection.API()
	serviceAcct := &kafkamgmtclient.ServiceAccountRequest{Name: fmt.Sprintf("rhoascli-%v", t.Unix())}
	req := api.ServiceAccount().CreateServiceAccount(ctx)
	req = req.ServiceAccountRequest(*serviceAcct)
	res, _, err := req.Execute()
	if err != nil {
		return nil, fmt.Errorf("%v: %w", c.localizer.MustLocalize("cluster.kubernetes.createServiceAccount.error.createError"), err)
	}

	return &res, nil
}
