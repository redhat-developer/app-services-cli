package connect

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	// _ "k8s.io/apimachinery/pkg/api/errors"
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

	cmd.Flags().BoolVarP(&localOnly, "local-only", "lo", false, "Provide yaml file containing changes without applying them to the cluster. Developers can use `oc apply -f kafka.yml` to apply it manually")
	cmd.Flags().BoolVarP(&secretOnly, "secret-only", "so", false, "Apply only secret and without CR. Can be used when no binding operator is configured")
	cmd.Flags().StringP(&kubeConfigCustomLocation, "kubeconfig", "", "", "Location of the .kube/config file")
	return cmd
}

func runBind(cmd *cobra.Command, _ []string) {
	connectToCluster()
	if dryrun {
		fmt.Fprintf(os.Stderr, "Generating CR files")
	}
	fmt.Fprintf(os.Stderr, "Successfully bound kafka credentials to your cluster")
}

func connectToCluster() {
	var kubeconfig *string

	if kubeConfigCustomLocation != "" {
		kubeconfig = kubeConfigCustomLocation
	} else if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	}

	if !fileExists(kubeconfig) {
		fmt.Fprint(os.Stderr, `Command uses oc or kubectl login context file. 
		Please make sure that you have configured access to your cluster and selected the right namespace`, err)
		return
	}

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
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
	for {
		pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
