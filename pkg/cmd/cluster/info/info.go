package info

import (
	"context"
	"fmt"

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

func NewInfoCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
	}

	cmd := &cobra.Command{
		Use:   "info",
		Short: "Prints information about your OpenShift cluster connection",
		Long:  `Prints information about your OpenShift cluster connection`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runInfo(opts)
		},
	}

	return cmd
}

func runInfo(opts *Options) error {
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
