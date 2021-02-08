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
		Short: "Connect your services to Kubernetes or OpenShift",
		Long: heredoc.Doc(`
			Connect your application services to your Kubernetes or OpenShift cluster.
			The kubeconfig file is used to connect to the cluster and identify the context.

			A service account is created and mounted as a secret into your cluster. 
			This enables you to mount credentials directly to your application.

			This command works in two modes:

				* If the RHOAS Operator is installed in the cluster, you can use it to bind your instance automatically.

				* Create the secret only. This mode does not require the Operator to be installed.

			You can interactively select the service instance by using the "--interactive-select" flag.
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

	cmd.Flags().BoolVarP(&opts.secretOnly, "secret-only", "", false, "Creates the secret, but doesn't bind the instance. Use this flag if the RHOAS Operator is not installed in the Kubernetes or OpenShift cluster.")
	cmd.Flags().BoolVarP(&opts.interactiveSelect, "interactive-select", "", false, "Interactively select the service instance that will be bound to your Kubernetes or OpenShift cluster.")
	cmd.Flags().StringVarP(&opts.secretName, "secret-name", "", "kafka-credentials", "Name of the secret that holds the Kafka credentials.")
	cmd.Flags().StringVarP(&opts.kubeconfigLocation, "kubeconfig", "", "", "Location of the kubeconfig file.")

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
