package cluster

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	serviceapi "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/serviceapi/client"
	pkgConnection "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/sdk/utils"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/client-go/util/homedir"

	// Get all auth schemes
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

var statusMsg = `
Linking your cluster with Managed Kafka
Kafka instance: %v
Namespace: %v
Secret name: %v
`

var MKCRMeta = metav1.TypeMeta{
	Kind:       "ManagedKafkaConnection",
	APIVersion: "rhoas.redhat.com/v1alpha1",
}

func ConnectToCluster(connection pkgConnection.Connection,
	config config.IConfig,
	secretName string,
	kubeConfigCustomLocation string,
	forceSelect bool) {
	var kubeconfig string

	if kubeConfigCustomLocation != "" {
		kubeconfig = kubeConfigCustomLocation
	} else if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	}

	if !utils.FileExists(kubeconfig) {
		fmt.Fprint(os.Stderr, `
		Command uses oc or kubectl login context file. 
		Please make sure that you have configured access to your cluster and selected the right namespace`)
		return
	}

	kubeClientconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfig},
		&clientcmd.ConfigOverrides{ClusterInfo: clientcmdapi.Cluster{Server: ""}})

	// use the current context in kubeconfig
	restConfig, err := kubeClientconfig.ClientConfig()
	if err != nil {
		fmt.Fprint(os.Stderr, "\nFailed to load kube config file", err)
		return
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(restConfig)

	if err != nil {
		fmt.Fprint(os.Stderr, "\nFailed to load kube config file", err)
		return
	}

	currentNamespace, _, _ := kubeClientconfig.Namespace()
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not load config: %v\n", err)
		return
	}

	if cfg.Services.Kafka == nil || forceSelect {
		cfg = useKafka(cfg, connection)
		if cfg == nil {
			return
		}
	}

	kafkaCfg := cfg.Services.Kafka

	api := connection.API()
	kafkaInstance, _, err := api.Kafka.GetKafkaById(context.TODO(), kafkaCfg.ClusterID).Execute()

	if err.Error() != "" {
		fmt.Fprintf(os.Stderr, "Could not get Kafka instance with ID '%v': %v\n", kafkaCfg.ClusterID, err)
		return
	}

	if kafkaInstance.BootstrapServerHost == nil || *kafkaInstance.BootstrapServerHost == "" {
		fmt.Fprintf(os.Stderr, "Kafka instance is missing required BootstrapServerHost variable")
		return
	}

	fmt.Fprintf(os.Stderr, statusMsg, color.HiGreenString(*kafkaInstance.Name), color.HiGreenString(currentNamespace), color.HiGreenString(secretName))
	if shouldContinue := utils.ShowQuestion("Do you want to continue?"); shouldContinue == false {
		return
	}

	credentials := CreateCredentials(connection)
	if credentials == nil {
		return
	}
	CreateSecret(credentials, currentNamespace, clientset, secretName)
	CreateCR(clientset, &kafkaInstance, currentNamespace, secretName)

}

func CreateCredentials(connection pkgConnection.Connection) *serviceapi.ServiceAccount {
	api := connection.API()

	t := time.Now()
	serviceAcct := &serviceapi.ServiceAccountRequest{Name: fmt.Sprintf("srvc-acct-%v", t.String())}
	a := api.Kafka.CreateServiceAccount(context.Background())
	a = a.ServiceAccountRequest(*serviceAcct)
	res, _, apiErr := a.Execute()

	if apiErr.Error() != "" {
		fmt.Fprintf(os.Stderr, "\nError creating Kafka Credentials: %v\n", apiErr)
		return nil
	}

	jsonResponse, _ := json.Marshal(res)
	var credentials serviceapi.ServiceAccount
	err := json.Unmarshal(jsonResponse, &credentials)
	if err != nil {
		fmt.Fprint(os.Stderr, "Invalid JSON response from server\n", err)
		return nil
	}

	fmt.Fprintf(os.Stderr, "Credentials created\n")
	return &credentials
}

func CreateSecret(credentials *serviceapi.ServiceAccount,
	currentNamespace string,
	clientset *kubernetes.Clientset,
	secretName string) *apiv1.Secret {
	// Create secret
	secret := &apiv1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name: secretName,
		},

		StringData: map[string]string{
			"clientID":     *credentials.ClientID,
			"clientSecret": *credentials.ClientSecret,
		},
	}

	_, err := clientset.CoreV1().Secrets(currentNamespace).Get(context.TODO(), secretName, metav1.GetOptions{})

	if err == nil {
		fmt.Fprint(os.Stderr, "\nSecret exist. Please use --secretName argument to change name\n")
		return nil
	}

	createdSecret, err := clientset.CoreV1().Secrets(currentNamespace).Create(context.TODO(), secret, metav1.CreateOptions{})
	if err != nil {
		fmt.Fprint(os.Stderr, "\nError when creating secret\n", err)
		return nil
	}

	fmt.Fprintf(os.Stderr, "\nSecret %v created\n", createdSecret.Name)

	return secret
}

func CreateCR(clientset *kubernetes.Clientset, kafkaInstance *serviceapi.KafkaRequest, namespace string, secretName string) {
	crName := secretName + "-" + *kafkaInstance.Name
	instanceID := *kafkaInstance.Id

	crInstance := &ManagedKafkaConnection{
		ObjectMeta: metav1.ObjectMeta{
			Name:      crName,
			Namespace: namespace,
		},
		TypeMeta: MKCRMeta,
		Spec: ManagedKafkaConnectionSpec{
			KafkaID: instanceID,
		},
		Status: ManagedKafkaConnectionStatus{
			CreatedBy: "RHOASCLI",
			BootstrapServer: BootstrapServerSpec{
				Host: *kafkaInstance.BootstrapServerHost,
			},
			SecretName: secretName,
		},
	}

	crJSON, err := json.Marshal(crInstance)
	if err != nil {
		fmt.Fprint(os.Stderr, "\nError when parsing ManagedKafkaConnection CR: ", err)
		return
	}

	crAPIURL := "/apis/rhoas.redhat.com/v1alpha1/namespaces/" + namespace + "/managedkafkaconnections"
	data := clientset.RESTClient().
		Post().
		AbsPath(crAPIURL).
		Body(crJSON).
		Do(context.TODO())

	if data.Error() != nil {
		rawData, _ := data.Raw()
		fmt.Fprint(os.Stderr, "\nError when creating ManagedKafkaConnection CR: ", string(rawData))
		return
	}

	fmt.Fprintf(os.Stderr, "\nManagedKafkaConnection resource %v created\n", crName)
}

/**
* Checks if we can fetch managedkafkaconnections
 */
func IsCRDInstalled(clientset *kubernetes.Clientset, namespace string) bool {
	crAPIURL := "/apis/rhoas.redhat.com/v1alpha1/namespaces/" + namespace + "managedkafkaconnections"
	data := clientset.RESTClient().
		Get().
		AbsPath(crAPIURL).
		Do(context.TODO())

	if data.Error() != nil {
		var status int
		if data.StatusCode(&status); status != 404 {
			rawData, _ := data.Raw()
			fmt.Fprint(os.Stderr, "\nCannot verify if cluster has ManagedKafkaConnection", string(rawData))
		}

		return false
	}

	return true
}

func useKafka(cliconfig *config.Config, connection pkgConnection.Connection) *config.Config {
	api := connection.API()

	response, _, apiErr := api.Kafka.ListKafkas(context.Background()).Execute()

	if apiErr.Error() != "" {
		fmt.Fprintf(os.Stderr, "Unable to get Kafka clusters: %v\n", apiErr)
		os.Exit(1)
	}

	if response.Size == 0 {
		fmt.Fprintln(os.Stderr, "No Kafka clusters found.")
		return nil
	}

	kafkas := []string{}
	for index := 0; index < len(response.Items); index++ {
		kafkas = append(kafkas, *response.Items[index].Name)
	}

	prompt := promptui.Select{
		Label: "Select Kafka cluster to connect",
		Items: kafkas,
	}

	index, _, err := prompt.Run()
	if err == nil {
		selectedKafka := response.Items[index]
		var kafkaConfig config.KafkaConfig = config.KafkaConfig{ClusterID: *selectedKafka.Id}
		cliconfig.Services.Kafka = &kafkaConfig

		return cliconfig
	}
	return nil
}
