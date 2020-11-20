package connect

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

var localOnly bool
var secretOnly bool
var kubeConfigCustomLocation string

func NewConnectCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "connect",
		Short: "connect currently selected Kafka to your OpenShift cluster",
		Long: `Connect will create secret containing kafka credentials that can be co
		
		Connect command will use current Kubernetes context (namespace/project you have selected) using oc or kubectl command line.
		Connect command will retrieve credentials for your kafka and mount them as secret into your project.
		You can use secret directly or utilize service-binding-operator to automatically bind your instance

		https://github.com/bf2fc6cc711aee1a0c2a/binding-operator

		If your cluster has binding-operator installed you would be able to bind your application with credentials directly from the console etc.
		`,
		Run: runBind,
	}

	cmd.Flags().BoolVarP(&localOnly, "local-only", "l", false, "Provide yaml file containing changes without applying them to the cluster. Developers can use `oc apply -f kafka.yml` to apply it manually")
	cmd.Flags().BoolVarP(&secretOnly, "secret-only", "s", false, "Apply only secret and without CR. Can be used when no binding operator is configured")
	cmd.Flags().StringVarP(&kubeConfigCustomLocation, "kubeconfig", "", "", "Location of the .kube/config file")
	return cmd
}

func runBind(cmd *cobra.Command, _ []string) {
	connectToCluster()
	if localOnly {
		fmt.Fprintf(os.Stderr, "Generating CR files")
	}
	fmt.Fprintf(os.Stderr, "Successfully bound kafka credentials to your cluster")
}

func connectToCluster() {
	var kubeconfig string

	if kubeConfigCustomLocation != "" {
		kubeconfig = kubeConfigCustomLocation
	} else if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	}

	if !fileExists(kubeconfig) {
		fmt.Fprint(os.Stderr, `Command uses oc or kubectl login context file. 
		Please make sure that you have configured access to your cluster and selected the right namespace`)
		return
	}

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		fmt.Fprint(os.Stderr, "Failed to load kube config file", err)
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Fprint(os.Stderr, "Failed to load kube config file", err)
		panic(err.Error())
	}

	/**

	  apiVersion: v1
	  kind: Secret
	  metadata:
	    name: mysecret
	  type: Opaque
	  stringData:
	    config.yaml: |
	      apiUrl: "https://my.api.com/api/v1"
	      username: <user>
		  password: <password>

	  **/

	secret := &apiv1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name: "kafka-credentials2",
		},
		StringData: map[string]string{
			"clientID":     "test",
			"clientSecret": "testtestetse",
		},
	}

	createdSecret, err := clientset.CoreV1().Secrets(metav1.NamespaceDefault).Create(context.TODO(), secret, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Print("There is secret created", createdSecret)

	secretJsonData, _ := json.Marshal(createdSecret)
	fmt.Printf(string(secretJsonData))

	// secrets, err := clientset.CoreV1().Secrets(metav1.NamespaceDefault).List(context.TODO(), metav1.ListOptions{})
	// if err != nil {
	// 	panic(err.Error())
	// }
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
