package status

import (
	"context"
	"fmt"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cluster"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/kubeclient"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/v1alpha"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/color"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/logging"

	"github.com/spf13/cobra"

	// Get all auth schemes
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

type options struct {
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	IO         *iostreams.IOStreams
	localizer  localize.Localizer
	Context    context.Context
	kubeconfig string
}

func NewStatusCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
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

	flags := flagutil.NewFlagSet(cmd, opts.localizer)
	flags.StringVar(&opts.kubeconfig, "kubeconfig", "", opts.localizer.MustLocalize("cluster.common.flag.kubeconfig.description"))

	return cmd
}

func runStatus(opts *options) error {
	conn, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	cliProperties := v1alpha.CommandEnvironment{
		IO:         opts.IO,
		Logger:     opts.Logger,
		Localizer:  opts.localizer,
		Config:     opts.Config,
		Connection: conn,
		Context:    opts.Context,
	}

	kubeClients, err := kubeclient.NewKubernetesClusterClients(&cliProperties, opts.kubeconfig)
	if err != nil {
		return err
	}
	clusterAPI := cluster.KubernetesClusterAPIImpl{
		KubernetesClients:  kubeClients,
		CommandEnvironment: &cliProperties,
	}

	status, err := clusterAPI.ExecuteStatus()
	if err != nil {
		return err
	}
	var rhoasStatus string
	//nolint:gocritic
	if status.RHOASOperatorAvailable {
		if status.LatestRHOASVersionAvailable {
			rhoasStatus = color.Success(opts.localizer.MustLocalize("cluster.common.operatorInstalledMessage"))
		} else {
			rhoasStatus = color.Info(opts.localizer.MustLocalize("cluster.common.operatorOutdatedMessage"))
		}
	} else {
		rhoasStatus = color.Error(opts.localizer.MustLocalize("cluster.common.operatorNotInstalledMessage"))
	}

	var sboStatus string
	if status.ServiceBindingOperatorAvailable {
		sboStatus = color.Success(opts.localizer.MustLocalize("cluster.common.operatorInstalledMessage"))
	} else {
		sboStatus = color.Error(opts.localizer.MustLocalize("cluster.common.operatorNotInstalledMessage"))
	}

	fmt.Fprintln(
		opts.IO.Out,
		opts.localizer.MustLocalize("cluster.status.statusMessage",
			localize.NewEntry("RHOASOperatorStatus", rhoasStatus),
			localize.NewEntry("SBOOperatorStatus", sboStatus)),
	)

	return nil
}
