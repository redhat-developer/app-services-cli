package status

import (
	"context"
	"fmt"

	"github.com/redhat-developer/app-services-cli/pkg/color"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/localize"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cluster"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/logging"

	"github.com/spf13/cobra"

	// Get all auth schemes
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

type Options struct {
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     func() (logging.Logger, error)
	IO         *iostreams.IOStreams
	localizer  localize.Localizer

	kubeconfig string
}

func NewStatusCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
	}

	cmd := &cobra.Command{
		Use:     opts.localizer.LoadMessage("cluster.status.cmd.use"),
		Short:   opts.localizer.LoadMessage("cluster.status.cmd.shortDescription"),
		Long:    opts.localizer.LoadMessage("cluster.status.cmd.longDescription"),
		Example: opts.localizer.LoadMessage("cluster.status.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runStatus(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.kubeconfig, "kubeconfig", "", "", opts.localizer.LoadMessage("cluster.common.flag.kubeconfig.description"))

	return cmd
}

func runStatus(opts *Options) error {
	connection, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	clusterConn, err := cluster.NewKubernetesClusterConnection(connection, opts.Config, logger, opts.kubeconfig, opts.IO, opts.localizer)
	if err != nil {
		return err
	}

	var operatorStatus string
	// Add versioning in future
	isCRDInstalled, err := clusterConn.IsRhoasOperatorAvailableOnCluster(context.Background())
	if isCRDInstalled && err != nil {
		logger.Debug(err)
	}

	if isCRDInstalled {
		operatorStatus = color.Success(opts.localizer.LoadMessage("cluster.common.operatorInstalledMessage"))
	} else {
		operatorStatus = color.Error(opts.localizer.LoadMessage("cluster.common.operatorNotInstalledMessage"))
	}

	currentNamespace, err := clusterConn.CurrentNamespace()
	if err != nil {
		return err
	}

	fmt.Fprintln(
		opts.IO.Out,
		opts.localizer.LoadMessage("cluster.status.statusMessage",
			localize.NewEntry("Namespace", color.Info(currentNamespace)),
			localize.NewEntry("OperatorStatus", operatorStatus)),
	)

	return nil
}
