package cluster

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/color"
	"os"
	"path/filepath"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/kafka"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/mitchellh/go-homedir"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"

	serviceapiclient "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/serviceapi/client"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"
)

// Kubernetes is a type which represents a Kubernetes cluster
type Kubernetes struct {
	connection connection.Connection
	config     config.IConfig
	logger     logging.Logger

	clientset    *kubernetes.Clientset
	clientconfig clientcmd.ClientConfig
}

const statusMsg = `
Linking your cluster with Managed Kafka
Kafka instance: %v
Namespace: %v
Secret name: %v
`

var MKCRMeta = metav1.TypeMeta{
	Kind:       "ManagedKafkaConnection",
	APIVersion: "rhoas.redhat.com/v1alpha1",
}

// NewKubernetesClusterConnection configures and connects to a Kubernetes cluster
func NewKubernetesClusterConnection(connection connection.Connection, config config.IConfig, logger logging.Logger, kubeconfig string) (Cluster, error) {
	if kubeconfig == "" {
		home, _ := homedir.Dir()
		kubeconfig = filepath.Join(home, ".kube", "config")
	}

	_, err := os.Stat(kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("Missing kubeconfig file: %w", err)
	}

	clientconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfig},
		&clientcmd.ConfigOverrides{ClusterInfo: clientcmdapi.Cluster{Server: ""}})

	// use the current context in kubeconfig
	restConfig, err := clientconfig.ClientConfig()
	if err != nil {
		return nil, fmt.Errorf("Failed to load kubeconfig: %w", err)
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(restConfig)

	if err != nil {
		return nil, fmt.Errorf("Failed to load kubeconfig: %w", err)
	}

	k8sCluster := &Kubernetes{
		connection,
		config,
		logger,
		clientset,
		clientconfig,
	}

	return k8sCluster, nil
}

// CurrentNamespace returns the currently set namespace
func (c *Kubernetes) CurrentNamespace() (string, error) {
	namespace, _, err := c.clientconfig.Namespace()

	return namespace, err
}

// Connect connects a remote Kafka instance to the Kubernetes cluster
func (c *Kubernetes) Connect(ctx context.Context, secretName string, forceSelect bool) error {
	cfg, err := c.config.Load()
	if err != nil {
		return err
	}

	if cfg.Services.Kafka == nil || forceSelect {
		// nolint
		selectedKafka, err := kafka.InteractiveSelect(c.connection, c.logger)
		if err != nil {
			return err
		}
		cfg.Services.Kafka.ClusterID = selectedKafka.GetId()
		_ = c.config.Save(cfg)
	}

	api := c.connection.API()
	kafkaInstance, _, err := api.Kafka.GetKafkaById(ctx, cfg.Services.Kafka.ClusterID).Execute()
	if err.Error() != "" {
		return err
	}

	currentNamespace, err := c.CurrentNamespace()
	if err != nil {
		return err
	}

	// print status
	c.logger.Infof(statusMsg, color.Info(kafkaInstance.GetName()), color.Info(currentNamespace), color.Info(secretName))

	var shouldContinue bool
	confirm := &survey.Confirm{
		Message: "Do you want to continue?",
	}
	err = survey.AskOne(confirm, &shouldContinue)
	if err != nil {
		return err
	}

	if !shouldContinue {
		c.logger.Debug("Cancelling connection to cluster")
		return nil
	}

	serviceAcct, err := c.createServiceAccount(ctx)
	if err != nil {
		return err
	}

	err = c.createSecret(ctx, serviceAcct, secretName)
	if err != nil {
		return err
	}

	err = c.createKafkaConnectionCustomResource(ctx, &kafkaInstance, secretName)
	if err != nil {
		return err
	}

	return nil
}

// IsKafkaConnectionCRDInstalled checks the cluster to see if a ManagedKafkaConnection CRD is installed
func (c *Kubernetes) IsKafkaConnectionCRDInstalled(ctx context.Context) (bool, error) {
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
func (c *Kubernetes) createKafkaConnectionCustomResource(ctx context.Context, kafkaInstance *serviceapiclient.KafkaRequest, secretName string) error {
	crName := fmt.Sprintf("%v-%v", secretName, kafkaInstance.GetName())
	kafkaID := kafkaInstance.GetId()

	namespace, err := c.CurrentNamespace()
	if err != nil {
		return err
	}

	kafkaConnectionCR := &ManagedKafkaConnection{
		ObjectMeta: metav1.ObjectMeta{
			Name:      crName,
			Namespace: namespace,
		},
		TypeMeta: MKCRMeta,
		Spec: ManagedKafkaConnectionSpec{
			KafkaID: kafkaID,
		},
		Status: ManagedKafkaConnectionStatus{
			CreatedBy: "RHOASCLI",
			BootstrapServer: BootstrapServerSpec{
				Host: *kafkaInstance.BootstrapServerHost,
			},
			SecretName: secretName,
		},
	}

	crJSON, err := json.Marshal(kafkaConnectionCR)
	if err != nil {
		return fmt.Errorf("Could not marshal ManagedKafkaConnection to JSON object: %w", err)
	}

	data := c.clientset.RESTClient().
		Post().
		AbsPath(c.getKafkaConnectionsAPIURL(namespace)).
		Body(crJSON).
		Do(ctx)

	if data.Error() != nil {
		return data.Error()
	}

	c.logger.Infof("ManagedKafkaConnection resource '%v' has been created", crName)

	return nil
}

// createSecret creates a new secret to store the SASL/PLAIN credentials from the service account
func (c *Kubernetes) createSecret(ctx context.Context, serviceAcct *serviceapiclient.ServiceAccount, secretName string) error {
	// Create secret type
	secret := &apiv1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name: secretName,
		},

		StringData: map[string]string{
			"clientID":     serviceAcct.GetClientID(),
			"clientSecret": serviceAcct.GetClientSecret(),
		},
	}

	namespace, err := c.CurrentNamespace()
	if err != nil {
		return err
	}

	_, err = c.clientset.CoreV1().Secrets(namespace).Get(ctx, secretName, metav1.GetOptions{})
	if err == nil {
		return fmt.Errorf("Secret %v already exists. Please choose a different name with --secret-name", color.Info(secretName))
	}

	createdSecret, err := c.clientset.CoreV1().Secrets(namespace).Create(context.TODO(), secret, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("Could not create secret: %w", err)
	}

	c.logger.Infof("Secret %v created", color.Info(createdSecret.Name))

	return nil
}

// createServiceAccount creates a service account with a random name
func (c *Kubernetes) createServiceAccount(ctx context.Context) (*serviceapiclient.ServiceAccount, error) {
	t := time.Now()

	api := c.connection.API()
	serviceAcct := &serviceapiclient.ServiceAccountRequest{Name: fmt.Sprintf("svc-acct-%v", t.String())}
	req := api.Kafka.CreateServiceAccount(ctx)
	req = req.ServiceAccountRequest(*serviceAcct)
	res, _, apiErr := req.Execute()

	if apiErr.Error() != "" {
		return nil, fmt.Errorf("Could not create service account: %w", apiErr)
	}

	return &res, nil
}

func (c *Kubernetes) getKafkaConnectionsAPIURL(namespace string) string {
	return fmt.Sprintf("/apis/rhoas.redhat.com/v1alpha1/namespaces/%v/managedkafkaconnections", namespace)
}
