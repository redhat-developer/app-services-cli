package connect

import (
	"fmt"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/sdk/cluster"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/spf13/cobra"
)

type Options struct {
	Config func() (config.Config, error)

	secretOnly         bool
	kubeconfigLocation string
	secretName         string
	// TODO: Rename to interactive
	forceSelect bool
}

func NewConnectCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config: f.Config,
	}

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
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runBind(opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.secretOnly, "secret-only", "", false, "Apply only secret and without CR. Can be used without installing RHOAS operator on cluster")
	cmd.Flags().BoolVarP(&opts.forceSelect, "skip-context", "", false, "Allows to select services before performing binding")
	cmd.Flags().StringVarP(&opts.secretName, "secret-name", "", "kafka-credentials", "Name of the secret that will be used to hold Kafka credentials")
	cmd.Flags().StringVarP(&opts.kubeconfigLocation, "kubeconfig", "", "", "Location of the .kube/config file")

	return cmd
}

func runBind(opts *Options) error {
	cfg, err := opts.Config()
	if err != nil {
		return fmt.Errorf("Error loading config: %w", err)
	}

	connection, err := cfg.Connection()
	if err != nil {
		return fmt.Errorf("Can't create connection: %w", err)
	}

	cluster.ConnectToCluster(connection, opts.secretName, opts.kubeconfigLocation, opts.forceSelect)

	return nil
}
