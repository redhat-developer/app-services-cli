package clean

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/shared/cluster"
	"github.com/redhat-developer/app-services-cli/pkg/shared/cluster/kubeclient"
	"github.com/redhat-developer/app-services-cli/pkg/shared/cluster/v1alpha"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"

	// Get all auth schemes

	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

type options struct {
	Config                  config.IConfig
	Connection              factory.ConnectionFunc
	Logger                  logging.Logger
	IO                      *iostreams.IOStreams
	localizer               localize.Localizer
	Context                 context.Context
	kubeconfig              string
	namespace               string
	forceCreationWithoutAsk bool
}

func NewCleanCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "clean",
		Short:   opts.localizer.MustLocalize("cluster.clean.shortDescription"),
		Long:    opts.localizer.MustLocalize("cluster.clean.longDescription"),
		Example: opts.localizer.MustLocalize("cluster.clean.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runStatus(opts)
		},
	}

	cmd.Flags().StringVar(&opts.kubeconfig, "kubeconfig", "", opts.localizer.MustLocalize("cluster.common.flag.kubeconfig.description"))
	cmd.Flags().StringVarP(&opts.namespace, "namespace", "n", "", opts.localizer.MustLocalize("cluster.common.flag.namespace.description"))
	cmd.Flags().BoolVarP(&opts.forceCreationWithoutAsk, "yes", "y", false, opts.localizer.MustLocalize("cluster.common.flag.yes.description"))
	return cmd
}

func runStatus(opts *options) error {
	conn, err := opts.Connection()
	if err != nil {
		return err
	}

	env := v1alpha.CommandEnvironment{
		IO:         opts.IO,
		Logger:     opts.Logger,
		Localizer:  opts.localizer,
		Config:     opts.Config,
		Connection: conn,
		Context:    opts.Context,
	}

	kubeClients, err := kubeclient.NewKubernetesClusterClients(&env, opts.kubeconfig)
	if err != nil {
		return err
	}
	clusterAPI := cluster.KubernetesClusterAPIImpl{
		KubernetesClients:  kubeClients,
		CommandEnvironment: &env,
	}
	return clusterAPI.ExecuteClean(&v1alpha.CleanOperationOptions{
		ForceDeleteWithoutAsk: opts.forceCreationWithoutAsk,
		Namespace:             opts.namespace,
	})
}
