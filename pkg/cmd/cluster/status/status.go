package status

import (
	"context"
	"fmt"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/color"

	"github.com/MakeNowJust/heredoc"
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cluster"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"

	"github.com/spf13/cobra"

	// Get all auth schemes
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

var statusMsg = `
Namespace: %v
Managed Application Services Operator: %v 
`

type Options struct {
	Config     config.IConfig
	Connection func() (connection.Connection, error)
	Logger     func() (logging.Logger, error)

	kubeconfig string
}

func NewStatusCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
	}

	cmd := &cobra.Command{
		Use:   "status",
		Short: "View status of the current Kubernetes or OpenShift cluster.",
		Long: heredoc.Doc(`
			View information about the current Kubernetes or OpenShift cluster. 
			You can use this information to connect your application service to the cluster.

			Before using this command, you must be logged into a Kubernetes or OpenShift 
			cluster. The command uses your kubeconfig file to identify the cluster context.
		`),
		Example: heredoc.Doc(`
			# print status of the current cluster
			$ rhoas cluster status
		`),
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runStatus(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.kubeconfig, "kubeconfig", "", "", "Location of the kubeconfig file.")

	return cmd
}

func runStatus(opts *Options) error {
	connection, err := opts.Connection()
	if err != nil {
		return err
	}

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	clusterConn, err := cluster.NewKubernetesClusterConnection(connection, opts.Config, logger, opts.kubeconfig)
	if err != nil {
		return err
	}

	var operatorStatus string
	// Add versioning in future
	isCRDInstalled, err := clusterConn.IsKafkaConnectionCRDInstalled(context.Background())
	if isCRDInstalled && err != nil {
		logger.Debug(err)
	}

	if isCRDInstalled {
		operatorStatus = color.Success("Installed")
	} else {
		operatorStatus = color.Error("Not installed")
	}

	currentNamespace, err := clusterConn.CurrentNamespace()
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf(statusMsg, color.Info(currentNamespace), operatorStatus))

	return nil
}
