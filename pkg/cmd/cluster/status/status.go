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

type args struct {
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	IO         *iostreams.IOStreams
	localizer  localize.Localizer
	Context    context.Context
	kubeconfig string
}

func NewStatusCommand(f *factory.Factory) *cobra.Command {
	opts := &args{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "status",
		Short:   opts.localizer.MustLocalize("cluster.status.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("cluster.status.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("cluster.status.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runStatus(opts)
		},
	}

	cmd.Flags().StringVar(&opts.kubeconfig, "kubeconfig", "", opts.localizer.MustLocalize("cluster.common.flag.kubeconfig.description"))

	return cmd
}

func runStatus(opts *args) error {
	conn, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	clusterConn, err := cluster.NewKubernetesClusterConnection(conn, opts.Config, opts.Logger, opts.kubeconfig, opts.IO, opts.localizer)
	if err != nil {
		return err
	}

	var operatorStatus string
	// Add versioning in future
	isCRDInstalled, err := clusterConn.IsRhoasOperatorAvailableOnCluster(opts.Context)
	if isCRDInstalled && err != nil {
		opts.Logger.Debug(err)
	}

	if isCRDInstalled {
		operatorStatus = color.Success(opts.localizer.MustLocalize("cluster.common.operatorInstalledMessage"))
	} else {
		operatorStatus = color.Error(opts.localizer.MustLocalize("cluster.common.operatorNotInstalledMessage"))
	}

	currentNamespace, err := clusterConn.CurrentNamespace()
	if err != nil {
		return err
	}

	fmt.Fprintln(
		opts.IO.Out,
		opts.localizer.MustLocalize("cluster.status.statusMessage",
			localize.NewEntry("Namespace", color.Info(currentNamespace)),
			localize.NewEntry("OperatorStatus", operatorStatus)),
	)

	return nil
}
