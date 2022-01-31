package connect

import (
	"context"

	"github.com/redhat-developer/app-services-cli/internal/build"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/shared/cluster"
	"github.com/redhat-developer/app-services-cli/pkg/shared/cluster/kubeclient"
	"github.com/redhat-developer/app-services-cli/pkg/shared/cluster/v1alpha"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/redhat-developer/app-services-cli/pkg/shared/servicespec"

	"github.com/spf13/cobra"
)

type options struct {
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	IO         *iostreams.IOStreams
	localizer  localize.Localizer
	Context    context.Context

	kubeconfigLocation string
	namespace          string

	offlineAccessToken      string
	forceCreationWithoutAsk bool
	serviceName             string
	serviceType             string
	ignoreContext           bool
}

func NewConnectCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "connect",
		Short:   opts.localizer.MustLocalize("cluster.connect.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("cluster.connect.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("cluster.connect.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runConnect(opts)
		},
	}

	cmd.Flags().StringVar(&opts.kubeconfigLocation, "kubeconfig", "", opts.localizer.MustLocalize("cluster.common.flag.kubeconfig.description"))
	cmd.Flags().StringVar(&opts.offlineAccessToken, "token", "", opts.localizer.MustLocalize("cluster.common.flag.offline.token.description", localize.NewEntry("OfflineTokenURL", build.OfflineTokenURL)))
	cmd.Flags().StringVarP(&opts.namespace, "namespace", "n", "", opts.localizer.MustLocalize("cluster.common.flag.namespace.description"))
	cmd.Flags().BoolVarP(&opts.forceCreationWithoutAsk, "yes", "y", false, opts.localizer.MustLocalize("cluster.common.flag.yes.description"))
	cmd.Flags().StringVar(&opts.serviceName, "service-name", "", opts.localizer.MustLocalize("cluster.common.flag.serviceName.description"))
	cmd.Flags().StringVar(&opts.serviceType, "service-type", "", opts.localizer.MustLocalize("cluster.common.flag.serviceType.description"))
	cmd.Flags().BoolVar(&opts.ignoreContext, "ignore-context", false, opts.localizer.MustLocalize("cluster.common.flag.ignoreContext.description"))

	_ = cmd.RegisterFlagCompletionFunc("service-type", func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return servicespec.AllServiceLabels, cobra.ShellCompDirectiveNoSpace
	})

	return cmd
}

func runConnect(opts *options) error {
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

	kubeClients, err := kubeclient.NewKubernetesClusterClients(&cliProperties, opts.kubeconfigLocation)
	if err != nil {
		return err
	}

	clusterAPI := cluster.KubernetesClusterAPIImpl{
		KubernetesClients:  kubeClients,
		CommandEnvironment: &cliProperties,
	}

	arguments := &v1alpha.ConnectOperationOptions{
		OfflineAccessToken:      opts.offlineAccessToken,
		ForceCreationWithoutAsk: opts.forceCreationWithoutAsk,
		Namespace:               opts.namespace,
		ServiceType:             opts.serviceType,
		ServiceName:             opts.serviceName,
		IgnoreContext:           opts.ignoreContext,
	}

	err = clusterAPI.ExecuteConnect(arguments)
	if err != nil {
		return err
	}

	return nil
}
