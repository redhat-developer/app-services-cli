package connect

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	ms "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/managedservices"
	msapi "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/managedservices/client"
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
		Long: `Connect will create secret containing kafka credentials that can be co

Connect command will use current Kubernetes context (namespace/project you have selected) using oc or kubectl command line.
Connect command will retrieve credentials for your kafka and mount them as secret into your project.
You can use secret directly or utilize service-binding-operator to automatically bind your instance

https://github.com/bf2fc6cc711aee1a0c2a/operator

If your cluster has binding-operator installed you would be able to bind your application with credentials directly from the console etc.
`,
		Run: runBind,
	}

	cmd.Flags().BoolVarP(&localOnly, "local-only", "l", false, "Provide yaml file containing changes without applying them to the cluster. Developers can use `oc apply -f kafka.yml` to apply it manually")
	cmd.Flags().BoolVarP(&secretOnly, "secret-only", "s", false, "Apply only secret and without CR. Can be used when no binding operator is configured")
	cmd.Flags().StringVarP(&kubeConfigCustomLocation, "kubeconfig", "", "", "Location of the .kube/config file")
	cmd.Flags().StringVarP(&secretName, "secretName", "", "kafka-credentials", "Name of the secret that will be used to hold Kafka credentials")
	return cmd
}

func runBind(cmd *cobra.Command, _ []string) {
	if localOnly {
		fmt.Fprintf(os.Stderr, "Generating CR files locally")
		return
	}
	connectToCluster()
}

func connectToCluster() {
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
		fmt.Fprint(os.Stderr, "\nInvalid configuration file", err)
		return
	}

	fmt.Fprintf(os.Stderr, statusMsg, color.HiGreenString(clicfg.Services.Kafka.ClusterName), color.HiGreenString(currentNamespace), color.HiGreenString(secretName))
	if shouldStop := showQuestion("Do you want to continue?"); shouldStop {
		return
	}
	client := ms.BuildClient()
	response, _, err := client.DefaultApi.CreateServiceAccount(context.Background())

	if err != nil {
		fmt.Fprintf(os.Stderr, "\nError creating Kafka Credentials: %v\n", err)
		return
	}

	jsonResponse, _ := json.Marshal(response)
	var credentials msapi.TokenResponse
	err = json.Unmarshal(jsonResponse, &credentials)
	if err != nil {
		fmt.Fprint(os.Stderr, "\nInvalid JSON response from server", err)
		return
	}

	fmt.Fprintf(os.Stderr, "\nCredentials created")

	secret := &apiv1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name: secretName,
		},
		StringData: map[string]string{
			"clientID":     "test",
			"clientSecret": "testtestetse",
		},
	}

	_, err = clientset.CoreV1().Secrets(currentNamespace).Get(context.TODO(), secretName, metav1.GetOptions{})

	if err == nil {
		fmt.Fprint(os.Stderr, "\nSecret exist. Please use --secretName argument to change name")
		return
	}

	createdSecret, err := clientset.CoreV1().Secrets(currentNamespace).Create(context.TODO(), secret, metav1.CreateOptions{})
	if err != nil {
		fmt.Fprint(os.Stderr, "\nError when creating secret", err)
		return
	}

	fmt.Fprintf(os.Stderr, "\nSecret %v created", createdSecret.Name)

	// TODO CR
	crName := secretName + "-" + clicfg.Services.Kafka.ClusterName
	fmt.Fprintf(os.Stderr, "\nManagedKafkaConnection resource %v created\n", crName)
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
		return errors.New(fmt.Sprintf("Number should be one of the values %v", allowedValues))
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
