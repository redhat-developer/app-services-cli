package status

import (
	"context"
	"fmt"

	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/color"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/cluster"
	"github.com/redhat-developer/app-services-cli/pkg/shared/cluster/kubeclient"
	"github.com/redhat-developer/app-services-cli/pkg/shared/cluster/v1alpha"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"

	// Get all auth schemes
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

type options struct {
	Connection     factory.ConnectionFunc
	Logger         logging.Logger
	IO             *iostreams.IOStreams
	localizer      localize.Localizer
	Context        context.Context
	ServiceContext servicecontext.IContext
	kubeconfig     string
}

func NewStatusCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Connection:     f.Connection,
		Logger:         f.Logger,
		IO:             f.IOStreams,
		localizer:      f.Localizer,
		Context:        f.Context,
		ServiceContext: f.ServiceContext,
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

func runStatus(opts *options) error {
	conn, err := opts.Connection()
	if err != nil {
		return err
	}

	cliProperties := v1alpha.CommandEnvironment{
		IO:             opts.IO,
		Logger:         opts.Logger,
		Localizer:      opts.localizer,
		Connection:     conn,
		Context:        opts.Context,
		ServiceContext: opts.ServiceContext,
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
