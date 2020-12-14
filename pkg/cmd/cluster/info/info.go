package info

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/sdk/cluster"

	"github.com/fatih/color"

	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/client-go/util/homedir"

	// Get all auth schemes
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

var statusMsg = `
Namespace: %v
Managed Application Services Operator: %v 
`

func NewInfoCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info",
		Short: "Prints information about your OpenShift cluster connection",
		Long:  `Prints information about your OpenShift cluster connection`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runInfo()
		},
	}

	return cmd
}

func runInfo() error {
	var kubeconfig string

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	}

	if !fileExists(kubeconfig) {
		return fmt.Errorf(`
		Command uses oc or kubectl login context file. 
		Please make sure that you have configured access to your cluster and selected the right namespace`)
	}

	kubeClientconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfig},
		&clientcmd.ConfigOverrides{ClusterInfo: clientcmdapi.Cluster{Server: ""}})

	// use the current context in kubeconfig
	restConfig, err := kubeClientconfig.ClientConfig()
	if err != nil {
		return fmt.Errorf("Failed to load kube config file: %w", err)
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(restConfig)

	if err != nil {
		return fmt.Errorf("Failed to load kube config file: %w", err)
	}

	currentNamespace, _, _ := kubeClientconfig.Namespace()

	var operatorStatus string
	// Add versioning in future
	if cluster.IsCRDInstalled(clientset, currentNamespace) {
		operatorStatus = color.HiGreenString("Installed")
	} else {
		operatorStatus = color.HiRedString("Not installed")
	}
	fmt.Fprintf(os.Stderr, statusMsg, color.HiGreenString(currentNamespace), operatorStatus)

	return nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
