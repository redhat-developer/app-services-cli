package connect

import (
	"fmt"
	"os"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/sdk/cluster"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/config"
	"github.com/spf13/cobra"
)

var secretOnly bool
var kubeConfigCustomLocation string
var secretName string
var forceSelect bool

func NewConnectCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "connect",
		Short: "connect currently selected Kafka to your OpenShift cluster",
		Long: `Connect command links your own OpenShift cluster with Managed Services.

Connect command will use current Kubernetes context (namespace/project you have selected) created by oc or kubectl command line.
Command will create new service account and mount it as secret into your cluster, giving you ability to mount credentials directly
to your application. 

Command work in two modes:

1) Using RHOAS operator installed on cluster.
You can  or utilize service-binding-operator to automatically bind your instance.
For more details please visit:
https://github.com/bf2fc6cc711aee1a0c2a/operator
2) Secret only (--secret-only) creates only secret (no extra operator installation is required)

Using --forceSelect will ignore current command context make interactive prompt for selecting service instance you want to use.

`,
		Run: runBind,
	}

	cmd.Flags().BoolVarP(&secretOnly, "secret-only", "", false, "Apply only secret and without CR. Can be used without installing RHOAS operator on cluster")
	cmd.Flags().BoolVarP(&forceSelect, "skip-context", "", false, "Allows to select services before performing binding")
	cmd.Flags().StringVarP(&secretName, "secretName", "", "kafka-credentials", "Name of the secret that will be used to hold Kafka credentials")
	cmd.Flags().StringVarP(&kubeConfigCustomLocation, "kubeconfig", "", "", "Location of the .kube/config file")
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

	cluster.ConnectToCluster(connection, secretName, kubeConfigCustomLocation, forceSelect)
}
