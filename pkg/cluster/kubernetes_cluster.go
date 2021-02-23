package cluster

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/dgrijalva/jwt-go"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/kas"
	kasclient "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/kas/client"

	"github.com/AlecAivazis/survey/v2"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/color"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/kafka"
	"github.com/gofrs/uuid"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/mitchellh/go-homedir"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"
)

// Kubernetes is a type which represents a Kubernetes cluster
type KubernetesCluster struct {
	connection   connection.Connection
	config       config.IConfig
	logger       logging.Logger
	clientset    *kubernetes.Clientset
	clientconfig clientcmd.ClientConfig
	restConfig   *rest.Config
}

var MKCRMeta = metav1.TypeMeta{
	Kind:       "ManagedKafkaConnection",
	APIVersion: "rhoas.redhat.com/v1alpha1",
}

/* #nosec */
var tokenSecretName = "rhoas-cli-api-token"

/* #nosec */
var serviceAccountSecretName = "rhoas-cli-serviceaccounts"

func init() {
	localizer.LoadMessageFiles("cluster/kubernetes")
}

var (
	runtimeMSARGVR = schema.GroupVersionResource{
		Group:    "rhoas.redhat.com",
		Version:  "v1alpha1",
		Resource: "managedserviceaccountrequests",
	}
)

// NewKubernetesClusterConnection configures and connects to a Kubernetes cluster
func NewKubernetesClusterConnection(connection connection.Connection, config config.IConfig, logger logging.Logger, kubeconfig string) (Cluster, error) {
	if kubeconfig == "" {
		home, _ := homedir.Dir()
		kubeconfig = filepath.Join(home, ".kube", "config")
	}
	if kubeconfig == "" {
		kubeconfig = os.Getenv("KUBECONFIG")
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
		kubeClientConfig,
	}

	return k8sCluster, nil
}

// CurrentNamespace returns the currently set namespace
func (c *KubernetesCluster) CurrentNamespace() (string, error) {
	namespace, _, err := c.clientconfig.Namespace()

	return namespace, err
}

// Connect connects a remote Kafka instance to the Kubernetes cluster
func (c *KubernetesCluster) Connect(ctx context.Context, forceSelect bool, apiToken string) error {
	cfg, err := c.config.Load()
	if err != nil {
		return err
	}

	err = c.useInteractiveMode(cfg, forceSelect, apiToken)
	if err != nil {
		return err
	}

	api := c.connection.API()
	kafkaInstance, _, err := api.Kafka().GetKafkaById(ctx, cfg.Services.Kafka.ClusterID).Execute()
	if kas.IsErr(err, kas.ErrorNotFound) {
		return kafka.ErrorNotFound(cfg.Services.Kafka.ClusterID)
	}

	if err.Error() != "" {
		return err
	}

	currentNamespace, err := c.CurrentNamespace()
	if err != nil {
		return err
	}

	// print status
	c.logger.Info(localizer.MustLocalize(&localizer.Config{MessageID: "cluster.kubernetes.log.info.statusMessage"}))

	c.logger.Info(localizer.MustLocalize(&localizer.Config{
		MessageID: "cluster.kubernetes.statusInfo",
		TemplateData: map[string]interface{}{
			"InstanceName": color.Info(kafkaInstance.GetName()),
			"Namespace":    color.Info(currentNamespace),
			"SecretName":   color.Info(serviceAccountSecretName),
		},
	}))

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

	// Token with auth for operator to pick
	err = c.createTokenSecret(ctx, apiToken)
	if err != nil {
		return err
	}

	_, err = c.createServiceAccount(ctx)
	if err != nil {
		return err
	}

	// err = c.createServiceAccountSecret(ctx, serviceAcct)
	// if err != nil {
	// 	return err
	// }

	err = c.createKafkaConnectionCustomResource(ctx, &kafkaInstance)
	if err != nil {
		return err
	}

	return nil
}

func (c *KubernetesCluster) useInteractiveMode(cfg *config.Config, forceSelect bool, apiToken string) error {
	if cfg.Services.Kafka == nil || forceSelect {
		// nolint
		selectedKafka, err := kafka.InteractiveSelect(c.connection, c.logger)
		if err != nil {
			return err
		}
		cfg.Services.Kafka = &config.KafkaConfig{
			ClusterID: selectedKafka.GetId(),
		}
		_ = c.config.Save(cfg)
	}
	if apiToken == "" || forceSelect {
		apiTokenInput := &survey.Input{
			Message: localizer.MustLocalizeFromID("cluster.common.flag.offline.token.description"),
		}
		err := survey.AskOne(apiTokenInput, &apiToken)
		if err != nil {
			return err
		}

	}
	parser := new(jwt.Parser)
	_, _, err := parser.ParseUnverified(apiToken, jwt.MapClaims{})

	return err
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
func (c *KubernetesCluster) createKafkaConnectionCustomResource(ctx context.Context, kafkaInstance *kasclient.KafkaRequest) error {
	id, _ := uuid.NewV1()
	crName := fmt.Sprintf("%v-%v", kafkaInstance.GetName(), id.String())
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

	return nil
}

func (c *KubernetesCluster) createTokenSecret(ctx context.Context, apiToken string) error {
	namespace, err := c.CurrentNamespace()
	if err != nil {
		return err
	}

	err = c.clientset.CoreV1().Secrets(namespace).Delete(context.TODO(), tokenSecretName, metav1.DeleteOptions{})
	if err == nil {
		c.logger.Info(localizer.MustLocalizeFromID("cluster.kubernetes.tokensecret.removed"))
	}

	// Create secret type
	secret := &apiv1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      tokenSecretName,
			Namespace: namespace,
		},
		StringData: map[string]string{
			"value": apiToken,
		},
	}

	_, err = c.clientset.CoreV1().Secrets(namespace).Create(ctx, secret, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("%v: %w", localizer.MustLocalizeFromID("cluster.kubernetes.creatTokenSecret.createError"), err)
	}

	c.logger.Info(localizer.MustLocalize(&localizer.Config{
		MessageID: "cluster.kubernetes.createSecret.log.info.createSuccess",
		TemplateData: map[string]interface{}{
			"Name": tokenSecretName,
		},
	}))

	return nil
}

// createSecret creates a new secret to store the SASL/PLAIN credentials from the service account
func (c *KubernetesCluster) createServiceAccountSecret(ctx context.Context, serviceAcct *kasclient.ServiceAccount) error {
	namespace, err := c.CurrentNamespace()
	if err != nil {
		return err
	}
	// Create secret type
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

	err = c.clientset.CoreV1().Secrets(namespace).Delete(context.TODO(), serviceAccountSecretName, metav1.DeleteOptions{})
	if err == nil {
		c.logger.Info(localizer.MustLocalizeFromID("cluster.kubernetes.serviceaccountsecret.removed"))
	}

	_, err = c.clientset.CoreV1().Secrets(namespace).Get(ctx, serviceAccountSecretName, metav1.GetOptions{})
	if err == nil {
		c.logger.Debug(localizer.MustLocalizeFromID("cluster.kubernetes.debug.serviceaccount.secret.removed"))
	}

	createdSecret, err := c.clientset.CoreV1().Secrets(namespace).Create(context.TODO(), secret, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("%v: %w", localizer.MustLocalizeFromID("cluster.kubernetes.createKafkaCR.error.createError"), err)
	}

	c.logger.Info(localizer.MustLocalize(&localizer.Config{
		MessageID: "cluster.kubernetes.createSecret.log.info.createSuccess",
		TemplateData: map[string]interface{}{
			"Name": createdSecret.Name,
		},
	}))

	return nil
}

// createServiceAccount creates a service account with a random name
func (c *KubernetesCluster) createServiceAccount(ctx context.Context) (*kasclient.ServiceAccount, error) {
	t := time.Now()
	client, err := dynamic.NewForConfig(c.restConfig)
	if nil != err {
		return nil, err
	}

	namespace, err := c.CurrentNamespace()

	if nil != err {
		return nil, err
	}

	name := fmt.Sprintf("svc-acct-%v", t.Nanosecond())
	res := client.Resource(runtimeMSARGVR)
	msar := NewMSAR(name)
	_, err = res.Namespace(namespace).Create(context.TODO(), msar, metav1.CreateOptions{})

	if nil != err {
		return nil, err
	}

	w, err := res.Namespace(namespace).Watch(context.TODO(), metav1.ListOptions{})

	if nil != err {
		fmt.Printf("ERROR!!! %v", err)
		return nil, fmt.Errorf("%+v", err)
	}

	for {
		select {
		case event := <-w.ResultChan():
			if event.Type == watch.Modified {
				unstructuredObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(event.Object)

				objName, _, _ := unstructured.NestedString(unstructuredObj, "spec", "serviceAccountName")
				if objName != name {
					continue
				}
				if err != nil {
					return nil, err
				}

				conditions, found, err := unstructured.NestedSlice(unstructuredObj, "status", "conditions")

				if err != nil {
					return nil, err
				}

				if found {
					for _, condition := range conditions {
						typedCondition := condition.(map[string]interface{})
						if typedCondition["type"].(string) == "Finished" {
							if typedCondition["status"].(string) == "False" {
								fmt.Printf("False %v", typedCondition)
								w.Stop()
								return nil, fmt.Errorf("Error Creating: %v", typedCondition["message"])
							}
							if typedCondition["status"].(string) == "True" {
								fmt.Printf("True %v", typedCondition)
								w.Stop()
								return nil, nil
							}
						}

					}
					w.Stop()
				}

			}

		case <-time.After(30 * time.Second):
			w.Stop()
			return nil, fmt.Errorf("timeout: %v", name)
		}
	}

}

func NewMSAR(name string) *unstructured.Unstructured {
	return &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "ManagedServiceAccountRequest",
			"apiVersion": "rhoas.redhat.com/v1alpha1",
			"metadata": map[string]interface{}{
				"name": name,
			},
			"spec": map[string]interface{}{
				"accessTokenSecretName":     tokenSecretName,
				"serviceAccountDescription": "test-description",
				"serviceAccountName":        name,
				"serviceAccountSecretName":  serviceAccountSecretName,
			},
		},
	}
}

func (c *KubernetesCluster) getKafkaConnectionsAPIURL(namespace string) string {
	return fmt.Sprintf("/apis/rhoas.redhat.com/v1alpha1/namespaces/%v/managedkafkaconnections", namespace)
}
