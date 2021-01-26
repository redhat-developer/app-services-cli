package status

import (
	"context"
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cluster"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"

	"github.com/fatih/color"

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
			View status of the current Kubernetes or OpenShift cluster using your kubeconfig file.

			The information shown is useful for connecting your service to the OpenShift cluster.

			For this command to work you must be logged into a Kubernetes or OpenShift cluster. The command
			uses the kubeconfig file to identify the cluster context.
		`),
		Example: heredoc.Doc(`
			# print status of the current cluster
			$ rhoas cluster status
		`),
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runStatus(opts)
		},
	}

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

	clusterConn, err := cluster.NewKubernetesClusterConnection(connection, opts.Config, logger, "")
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
		operatorStatus = color.HiGreenString("Installed")
	} else {
		operatorStatus = color.HiRedString("Not installed")
	}

	currentNamespace, err := clusterConn.CurrentNamespace()
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf(statusMsg, color.HiGreenString(currentNamespace), operatorStatus))

	return nil
}
