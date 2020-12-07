package connect

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	pkgConnection "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/operator/connection"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/managedservices"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/config"
	"github.com/spf13/cobra"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/client-go/util/homedir"

	// Get all auth schemes
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

var localOnly bool
var secretOnly bool
var kubeConfigCustomLocation string
var secretName string
var forceKafkaSelect bool

var statusMsg = `
Linking your cluster with Managed Kafka
Kafka instance: %v
Namespace: %v
Secret name: %v

`

func NewConnectCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "connect",
		Short: "connect currently selected Kafka to your OpenShift cluster",
		Long: `Connect will create secret containing Kafka credentials.

Connect command will use current Kubernetes context (namespace/project you have selected) created by oc or kubectl command line.
Connect command will retrieve credentials for your kafka and mount them as secret into your project.
You can use secret directly or utilize service-binding-operator to automatically bind your instance
For more details please visit:

https://github.com/bf2fc6cc711aee1a0c2a/operator

If your cluster has binding-operator installed you would be able to bind your application with credentials directly from the console etc.
`,
		Run: runBind,
	}

	cmd.Flags().BoolVarP(&localOnly, "local-only", "l", false, "Provide yaml file containing changes without applying them to the cluster. Developers can use `oc apply -f kafka.yml` to apply it manually")
	cmd.Flags().BoolVarP(&secretOnly, "secret-only", "s", false, "Apply only secret and without CR. Can be used when no binding operator is configured")
	cmd.Flags().StringVarP(&kubeConfigCustomLocation, "kubeconfig", "", "", "Location of the .kube/config file")
	cmd.Flags().StringVarP(&secretName, "secretName", "", "kafka-credentials", "Name of the secret that will be used to hold Kafka credentials")
	cmd.Flags().BoolVarP(&forceKafkaSelect, "forceKafkaSelection", "", false, "Name of the secret that will be used to hold Kafka credentials")
	return cmd
}

func runBind(cmd *cobra.Command, _ []string) {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v", err)
		os.Exit(1)
	}

	connection, err := cfg.Connection()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't create connection: %v\n", err)
		os.Exit(1)
	}

	if localOnly {
		// TODO
		fmt.Fprintf(os.Stderr, "Generating CR files locally")
		return
	}
	connectToCluster(connection)
}

func connectToCluster(connection *pkgConnection.Connection) {
	var kubeconfig string

	if kubeConfigCustomLocation != "" {
		kubeconfig = kubeConfigCustomLocation
	} else if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	}

	if !fileExists(kubeconfig) {
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
	clicfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not load config: %v\n", err)
		return
	}

	if !clicfg.HasKafka() || forceKafkaSelect {
		clicfg = useKafka(clicfg, connection)
		if clicfg == nil {
			return
		}
	}

	kafkaCfg := clicfg.Services.Kafka

	managedservices := connection.NewMASClient()
	kafkaInstance, _, err := managedservices.DefaultApi.GetKafkaById(context.TODO(), kafkaCfg.ClusterID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get Kafka instance with ID '%v': %v", kafkaCfg.ClusterID, err)
		return
	}

	if err != nil {
		fmt.Fprint(os.Stderr, "\nInvalid configuration file", err)
		return
	}

	fmt.Fprintf(os.Stderr, statusMsg, color.HiGreenString(kafkaInstance.Name), color.HiGreenString(currentNamespace), color.HiGreenString(secretName))
	if shouldContinue := showQuestion("Do you want to continue?"); shouldContinue == false {
		return
	}

	credentials := createCredentials(connection)
	if credentials == nil {
		return
	}
	createSecret(credentials, currentNamespace, clientset)
	createCR(clientset, &kafkaInstance, currentNamespace)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func showQuestion(message string) bool {
	allowedValues := [...]string{"y", "yes", "no", "n"}

	validate := func(input string) error {
		for _, value := range allowedValues {
			if strings.ToLower(input) == value {
				return nil
			}
		}
		return fmt.Errorf("Number should be one of the values %v", allowedValues)
	}

	prompt := promptui.Prompt{
		Label:    message,
		Validate: validate,
		Default:  "y",
	}

	result, err := prompt.Run()
	if err != nil {
		return showQuestion(message)
	}

	result = strings.ToLower(result)

	return result == "y" || result == "yes"
}

func createCredentials(connection *pkgConnection.Connection) *managedservices.TokenResponse {
	client := connection.NewMASClient()

	t := time.Now()
	serviceAcct := &managedservices.ServiceAccountRequest{Name: fmt.Sprintf("srvc-acct-%v", t.String())}
	response, _, err := client.DefaultApi.CreateServiceAccount(context.Background(), *serviceAcct)

	if err != nil {
		fmt.Fprintf(os.Stderr, "\nError creating Kafka Credentials: %v\n", err)
		return nil
	}

	jsonResponse, _ := json.Marshal(response)
	var credentials managedservices.TokenResponse
	err = json.Unmarshal(jsonResponse, &credentials)
	if err != nil {
		fmt.Fprint(os.Stderr, "\nInvalid JSON response from server", err)
		return nil
	}

	fmt.Fprintf(os.Stderr, "\nCredentials created")
	return &credentials
}

func createSecret(credentials *managedservices.TokenResponse, currentNamespace string, clientset *kubernetes.Clientset) *apiv1.Secret {
	// Create secret
	secret := &apiv1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name: secretName,
		},
		// Type of CredentialsSecret
		StringData: map[string]string{
			"clientID":     credentials.ClientID,
			"clientSecret": credentials.ClientSecret,
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

	fmt.Fprintf(os.Stderr, "\nSecret %v created", createdSecret.Name)

	return secret
}

func createCR(clientset *kubernetes.Clientset, kafkaInstance *managedservices.KafkaRequest, namespace string) {
	crName := secretName + "-" + kafkaInstance.Name
	crInstance := &connection.ManagedKafkaConnection{
		ObjectMeta: metav1.ObjectMeta{
			Name:      crName,
			Namespace: namespace,
		},
		TypeMeta: metav1.TypeMeta{
			Kind:       "ManagedKafkaConnection",
			APIVersion: "rhoas.redhat.com/v1",
		},
		Spec: connection.ManagedKafkaConnectionSpec{
			BootstrapServer: connection.BootstrapServerSpec{
				Host: kafkaInstance.BootstrapServerHost,
			},
			Credentials: connection.CredentialsSpec{
				Kind:       connection.ClientCredentials,
				SecretName: secretName,
			},
		},
	}

	crJSON, err := json.Marshal(crInstance)
	if err != nil {
		fmt.Fprint(os.Stderr, "\nError when parsing ManagedKafkaConnection CR: ", err)
		return
	}

	crAPIURL := "/apis/rhoas.redhat.com/v1/namespaces/" + namespace + "/managedkafkaconnections"
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

func useKafka(cliconfig *config.Config, connection *pkgConnection.Connection) *config.Config {
	client := connection.NewMASClient()
	options := managedservices.ListKafkasOpts{}
	response, _, err := client.DefaultApi.ListKafkas(context.Background(), &options)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error retrieving Kafka instances: %v\n", err)
		os.Exit(1)
	}

	if response.Size == 0 {
		fmt.Fprintln(os.Stderr, "No Kafka instances found.")
		return nil
	}

	kafkas := []string{}
	for index := 0; index < len(response.Items); index++ {
		kafkas = append(kafkas, response.Items[index].Name)
	}

	prompt := promptui.Select{
		Label: "Select Kafka instance to connect",
		Items: kafkas,
	}

	index, _, err := prompt.Run()
	if err == nil {
		selectedKafka := response.Items[index]
		var kafkaConfig config.KafkaConfig = config.KafkaConfig{ClusterID: selectedKafka.Id}
		cliconfig.Services.SetKafka(&kafkaConfig)

		return cliconfig
	}
	return nil
}
