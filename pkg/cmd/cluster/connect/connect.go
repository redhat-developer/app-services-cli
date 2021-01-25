package connect

import (
	"context"
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cluster"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/iostreams"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/spf13/cobra"
)

type Options struct {
	Config     config.IConfig
	Connection func() (connection.Connection, error)
	Logger     func() (logging.Logger, error)
	IO         *iostreams.IOStreams

	secretOnly         bool
	kubeconfigLocation string
	secretName         string
	interactiveSelect  bool
}

func NewConnectCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
	}

	cmd := &cobra.Command{
		Use:   "connect",
		Short: "Connect your services to a Kubernetes or OpenShift",
		Long: heredoc.Doc(`
			Connect your services to your Kubernetes or OpenShift cluster.
			The kubeconfig file is used to connect with the cluster and identify the context.

			A service account is created and mounted as a secret into your cluster. 
			This gives you the ability to mount credentials directly to your application.

			Command work in two modes:

			1) Using RHOAS operator installed on cluster.
			You can  or utilize service-binding-operator to automatically bind your instance.
			For more details please visit:
			https://github.com/bf2fc6cc711aee1a0c2a/operator
			2) Secret only (--secret-only) creates only secret (no extra operator installation is required)

			Using --interactive-select will ignore current command context make interactive prompt for selecting service instance you want to use.
		`),
		Example: heredoc.Doc(`
			# connect the current Kafka instance to your cluster
			$ rhoas cluster connect
		`),
		RunE: func(cmd *cobra.Command, _ []string) error {
			if !opts.IO.CanPrompt() && opts.interactiveSelect {
				return fmt.Errorf("Cannot use --interactive-select when not running interactively")
			}

			return runBind(opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.secretOnly, "secret-only", "", false, "Apply only secret and without CR. Can be used without installing RHOAS operator on cluster")
	cmd.Flags().BoolVarP(&opts.interactiveSelect, "interactive-select", "", false, "Allows to select services before performing binding")
	cmd.Flags().StringVarP(&opts.secretName, "secret-name", "", "kafka-credentials", "Name of the secret that will be used to hold Kafka credentials")
	cmd.Flags().StringVarP(&opts.kubeconfigLocation, "kubeconfig", "", "", "Location of the .kube/config file")

	return cmd
}

func runBind(opts *Options) error {
	connection, err := opts.Connection()
	if err != nil {
		return err
	}

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	clusterConn, err := cluster.NewKubernetesClusterConnection(connection, opts.Config, logger, opts.kubeconfigLocation)
	if err != nil {
		return err
	}

	err = clusterConn.Connect(context.Background(), opts.secretName, opts.interactiveSelect)
	if err != nil {
		return err
	}

	return nil
}
