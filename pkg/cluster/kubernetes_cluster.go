package cluster

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"path/filepath"
	"time"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/kas"
	kasclient "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/kas/client"

	"github.com/AlecAivazis/survey/v2"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/color"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/kafka"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"
)

// KubernetesCluster is a type which represents a Kubernetes cluster
type KubernetesCluster struct {
	connection connection.Connection
	config     config.IConfig
	logger     logging.Logger

	clientset    *kubernetes.Clientset
	clientconfig clientcmd.ClientConfig
	// restConfig   *rest.Config
}

var MKCGroup = "rhoas.redhat.com"
var MKCVersion = "v1alpha1"

var MKCRMeta = metav1.TypeMeta{
	Kind:       "ManagedKafkaConnection",
	APIVersion: MKCGroup + "/" + MKCVersion,
}

/* #nosec */
var tokenSecretName = "rh-managed-services-accesstoken-cli"

/* #nosec */
var serviceAccountSecretName = "rh-managed-services-service-account"

// NewKubernetesClusterConnection configures and connects to a Kubernetes cluster
func NewKubernetesClusterConnection(connection connection.Connection, config config.IConfig, logger logging.Logger, kubeconfig string) (Cluster, error) {
	if kubeconfig == "" {
		kubeconfig = os.Getenv("KUBECONFIG")
	}

	if kubeconfig == "" {
		home, _ := os.UserHomeDir()
		kubeconfig = filepath.Join(home, ".kube", "config")
	}

	_, err := os.Stat(kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", localizer.MustLocalizeFromID("cluster.kubernetes.error.configNotFoundError"), err)
	}

	kubeClientConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", localizer.MustLocalizeFromID("cluster.kubernetes.error.loadConfigError"), err)
	}

	// create the clientset for using Rest Client
	clientset, err := kubernetes.NewForConfig(kubeClientConfig)

	// Used for namespaces and general queries
	clientconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfig},
		&clientcmd.ConfigOverrides{ClusterInfo: clientcmdapi.Cluster{Server: ""}})

	if err != nil {
		return nil, fmt.Errorf("%v: %w", localizer.MustLocalizeFromID("cluster.kubernetes.error.loadConfigError"), err)
	}

	k8sCluster := &KubernetesCluster{
		connection,
		config,
		logger,
		clientset,
		clientconfig,
	}

	return k8sCluster, nil
}

// CurrentNamespace returns the currently set namespace
func (c *KubernetesCluster) CurrentNamespace() (string, error) {
	namespace, _, err := c.clientconfig.Namespace()

	return namespace, err
}

// Connect connects a remote Kafka instance to the Kubernetes cluster
func (c *KubernetesCluster) Connect(ctx context.Context, cmdOptions *CommandConnectOptions) error {
	api := c.connection.API()
	kafkaInstance, _, apiError := api.Kafka().GetKafkaById(ctx, cmdOptions.SelectedKafka).Execute()
	if kas.IsErr(apiError, kas.ErrorNotFound) {
		return kafka.ErrorNotFound(cmdOptions.SelectedKafka)
	}

	if apiError.Error() != "" {
		return errors.New(apiError.Error())
	}

	currentNamespace, err := c.getCurrentNamespace(cmdOptions)
	if err != nil {
		return err
	}

	// print status
	c.logger.Info(localizer.MustLocalize(&localizer.Config{MessageID: "cluster.kubernetes.log.info.statusMessage"}))

	c.logger.Info(localizer.MustLocalize(&localizer.Config{
		MessageID: "cluster.kubernetes.statusInfo",
		TemplateData: map[string]interface{}{
			"InstanceName":             color.Info(kafkaInstance.GetName()),
			"Namespace":                color.Info(currentNamespace),
			"ServiceAccountSecretName": color.Info(serviceAccountSecretName),
		},
	}))

	if cmdOptions.Force == false {
		var shouldContinue bool
		confirm := &survey.Confirm{
			Message: localizer.MustLocalizeFromID("cluster.kubernetes.connect.input.confirm.message"),
		}
		err = survey.AskOne(confirm, &shouldContinue)
		if err != nil {
			return err
		}

		if !shouldContinue {
			c.logger.Debug(localizer.MustLocalizeFromID("cluster.kubernetes.connect.log.debug.cancellingConnect"))
			return nil
		}
	}

	err = c.checkIfConnectionsExist(ctx, currentNamespace, &kafkaInstance)
	if err != nil {
		return err
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

	err = c.createKafkaConnectionCustomResource(ctx, currentNamespace, &kafkaInstance)
	if err != nil {
		return err
	}

	return nil
}

func (c *KubernetesCluster) getCurrentNamespace(cmdOptions *CommandConnectOptions) (string, error) {

	if cmdOptions.Namespace != "" {
		return cmdOptions.Namespace, nil
	} else {
		return c.CurrentNamespace()
	}
}

// IsKafkaConnectionCRDInstalled checks the cluster to see if a ManagedKafkaConnection CRD is installed
func (c *KubernetesCluster) IsKafkaConnectionCRDInstalled(ctx context.Context) (bool, error) {
	namespace, err := c.CurrentNamespace()
	if err != nil {
		return false, err
	}

	data := c.clientset.
		RESTClient().
		Get().
		AbsPath(c.getKafkaConnectionsAPIURL(namespace)).
		Do(ctx)

	if data.Error() == nil {
		return true, nil
	}

	var status int
	if data.StatusCode(&status); status == 404 {
		return false, nil
	}

	return true, data.Error()
}

// createKafkaConnectionCustomResource creates a new "ManagedKafkaConnection" CR
func (c *KubernetesCluster) createKafkaConnectionCustomResource(ctx context.Context, namespace string, kafkaInstance *kasclient.KafkaRequest) error {
	crName := kafkaInstance.GetName()
	kafkaID := kafkaInstance.GetId()

	kafkaConnectionCR := &ManagedKafkaConnection{
		ObjectMeta: metav1.ObjectMeta{
			Name:      crName,
			Namespace: namespace,
		},
		TypeMeta: MKCRMeta,
		Spec: ManagedKafkaConnectionSpec{
			KafkaID:               kafkaID,
			AccessTokenSecretName: tokenSecretName,
			Credentials: CredentialsSpec{
				SecretName: serviceAccountSecretName,
			},
		},
	}

	crJSON, err := json.Marshal(kafkaConnectionCR)
	if err != nil {
		return fmt.Errorf("%v: %w", "cluster.kubernetes.createKafkaCR.error.marshalError", err)
	}

	data := c.clientset.RESTClient().
		Post().
		AbsPath(c.getKafkaConnectionsAPIURL(namespace)).
		Body(crJSON).
		Do(ctx)

	if data.Error() != nil {
		return data.Error()
	}

	c.logger.Info(localizer.MustLocalize(&localizer.Config{
		MessageID: "cluster.kubernetes.createKafkaCR.log.info.customResourceCreated",
		TemplateData: map[string]interface{}{
			"Name": crName,
		},
	}))

	c.logger.Info(localizer.MustLocalize(&localizer.Config{
		MessageID: "cluster.kubernetes.createKafkaCR.log.info.wait",
		TemplateData: map[string]interface{}{
			"Name":      crName,
			"Namespace": namespace,
			"Group":     MKCGroup,
			"Version":   MKCVersion,
			"Kind":      MKCRMeta.Kind,
		},
	}))

	return nil
}

func (c *KubernetesCluster) createTokenSecretIfNeeded(ctx context.Context, namespace string, opts *CommandConnectOptions) error {
	_, err := c.clientset.CoreV1().Secrets(namespace).Get(context.TODO(), tokenSecretName, metav1.GetOptions{})
	if err == nil {
		c.logger.Info(localizer.MustLocalizeFromID("cluster.kubernetes.tokensecret.log.info.found"), tokenSecretName)
		return nil
	}

	if opts.OfflineAccessToken == "" && !opts.IO.CanPrompt() {
		return errors.New(localizer.MustLocalize(&localizer.Config{
			MessageID: "flag.error.requiredWhenNonInteractive",
			TemplateData: map[string]interface{}{
				"Flag": "token",
			},
		}))
	}

	if opts.OfflineAccessToken == "" {
		apiTokenInput := &survey.Input{
			Message: localizer.MustLocalizeFromID("cluster.common.flag.offline.token.description"),
		}
		err := survey.AskOne(apiTokenInput, &opts.OfflineAccessToken)
		if err != nil {
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
	if err != nil {
		return fmt.Errorf("%v: %w", localizer.MustLocalize(&localizer.Config{
			MessageID: "cluster.kubernetes.createTokenSecret.log.info.createFailed",
			TemplateData: map[string]interface{}{
				"Name": tokenSecretName,
			},
		}), err)
	}

	c.logger.Info(localizer.MustLocalize(&localizer.Config{
		MessageID: "cluster.kubernetes.createTokenSecret.log.info.createSuccess",
		TemplateData: map[string]interface{}{
			"Name": tokenSecretName,
		},
	}))

	return nil
}

// createSecret creates a new secret to store the SASL/PLAIN credentials from the service account
func (c *KubernetesCluster) createServiceAccountSecretIfNeeded(ctx context.Context, namespace string) error {
	_, err := c.clientset.CoreV1().Secrets(namespace).Get(context.TODO(), serviceAccountSecretName, metav1.GetOptions{})
	if err == nil {
		c.logger.Info(localizer.MustLocalizeFromID("cluster.kubernetes.serviceaccountsecret.log.info.exist"))
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
			"client-id":     serviceAcct.GetClientID(),
			"client-secret": serviceAcct.GetClientSecret(),
		},
	}

	createdSecret, err := c.clientset.CoreV1().Secrets(namespace).Create(context.TODO(), secret, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("%v: %w", localizer.MustLocalizeFromID("cluster.kubernetes.serviceaccountsecret.error.createError"), err)
	}

	c.logger.Info(localizer.MustLocalize(&localizer.Config{
		MessageID: "cluster.kubernetes.createSASecret.log.info.createSuccess",
		TemplateData: map[string]interface{}{
			"Name": createdSecret.Name,
		},
	}))

	return nil
}

// createServiceAccount creates a service account with a random name
func (c *KubernetesCluster) createServiceAccount(ctx context.Context) (*kasclient.ServiceAccount, error) {
	t := time.Now()

	api := c.connection.API()
	serviceAcct := &kasclient.ServiceAccountRequest{Name: fmt.Sprintf("svc-acct-%v", t.String())}
	req := api.Kafka().CreateServiceAccount(ctx)
	req = req.ServiceAccountRequest(*serviceAcct)
	res, _, apiErr := req.Execute()

	if apiErr.Error() != "" {
		return nil, fmt.Errorf("%v: %w", localizer.MustLocalizeFromID("cluster.kubernetes.createServiceAccount.error.createError"), apiErr)
	}

	return &res, nil
}

func (c *KubernetesCluster) getKafkaConnectionsAPIURL(namespace string) string {
	return fmt.Sprintf("/apis/rhoas.redhat.com/v1alpha1/namespaces/%v/managedkafkaconnections", namespace)
}

func (c *KubernetesCluster) checkIfConnectionsExist(ctx context.Context, namespace string, k *kasclient.KafkaRequest) error {
	data := c.clientset.
		RESTClient().
		Get().
		AbsPath(c.getKafkaConnectionsAPIURL(namespace), k.GetName()).
		Do(ctx)

	var status int
	if data.StatusCode(&status); status == 404 {
		return nil
	}

	if data.Error() == nil {
		return fmt.Errorf("%v: %s", localizer.MustLocalizeFromID("cluster.kubernetes.checkIfConnectionExist.existError"), k.GetName())
	}

	return nil
}
