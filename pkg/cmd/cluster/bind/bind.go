package bind

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/cluster"
	"github.com/redhat-developer/app-services-cli/pkg/shared/cluster/kubeclient"
	"github.com/redhat-developer/app-services-cli/pkg/shared/cluster/v1alpha"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"

	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/redhat-developer/app-services-cli/pkg/shared/servicespec"

	"github.com/spf13/cobra"
)

type options struct {
	Connection     func() (connection.Connection, error)
	Logger         logging.Logger
	IO             *iostreams.IOStreams
	localizer      localize.Localizer
	Context        context.Context
	ServiceContext servicecontext.IContext

	kubeconfigLocation string
	namespace          string

	forceCreationWithoutAsk bool
	ignoreContext           bool
	appName                 string
	serviceType             string
	serviceName             string

	deploymentConfigEnabled bool
	bindAsEnv               bool
	bindingName             string
}

func NewBindCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Connection:     f.Connection,
		Logger:         f.Logger,
		IO:             f.IOStreams,
		localizer:      f.Localizer,
		Context:        f.Context,
		ServiceContext: f.ServiceContext,
	}

	cmd := &cobra.Command{
		Use:     "bind",
		Short:   opts.localizer.MustLocalize("cluster.bind.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("cluster.bind.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("cluster.bind.cmd.example"),
		RunE: func(cmd *cobra.Command, _ []string) error {
			if opts.ignoreContext == true && !opts.IO.CanPrompt() {
				return opts.localizer.MustLocalizeError("flag.error.requiredWhenNonInteractive", localize.NewEntry("Flag", "ignore-context"))
			}

			if opts.appName == "" && !opts.IO.CanPrompt() {
				return opts.localizer.MustLocalizeError("flag.error.requiredWhenNonInteractive", localize.NewEntry("Flag", "app-name"))
			}

			if opts.serviceType == "" && !opts.IO.CanPrompt() {
				return opts.localizer.MustLocalizeError("flag.error.requiredWhenNonInteractive", localize.NewEntry("Flag", "service-type"))
			}

			if opts.serviceName == "" && !opts.IO.CanPrompt() {
				return opts.localizer.MustLocalizeError("flag.error.requiredWhenNonInteractive", localize.NewEntry("Flag", "service-name"))
			}
			return runBind(opts)
		},
	}

	cmd.Flags().StringVar(&opts.kubeconfigLocation, "kubeconfig", "", opts.localizer.MustLocalize("cluster.common.flag.kubeconfig.description"))
	cmd.Flags().StringVar(&opts.appName, "app-name", "", opts.localizer.MustLocalize("cluster.bind.flag.appName"))
	cmd.Flags().StringVar(&opts.bindingName, "binding-name", "", opts.localizer.MustLocalize("cluster.bind.flag.bindName"))
	cmd.Flags().BoolVarP(&opts.forceCreationWithoutAsk, "yes", "y", false, opts.localizer.MustLocalize("cluster.common.flag.yes.description"))
	cmd.Flags().StringVarP(&opts.namespace, "namespace", "n", "", opts.localizer.MustLocalize("cluster.common.flag.namespace.description"))
	cmd.Flags().BoolVar(&opts.ignoreContext, "ignore-context", false, opts.localizer.MustLocalize("cluster.common.flag.ignoreContext.description"))
	cmd.Flags().BoolVar(&opts.deploymentConfigEnabled, "deployment-config", false, opts.localizer.MustLocalize("cluster.bind.flag.deploymentConfig.description"))
	cmd.Flags().BoolVar(&opts.bindAsEnv, "bind-env", false, opts.localizer.MustLocalize("cluster.bind.flag.bindenv.description"))
	cmd.Flags().StringVar(&opts.serviceType, "service-type", "", opts.localizer.MustLocalize("cluster.common.flag.serviceType.description"))
	cmd.Flags().StringVar(&opts.serviceName, "service-name", "", opts.localizer.MustLocalize("cluster.common.flag.serviceName.description"))

	_ = cmd.RegisterFlagCompletionFunc("service-type", func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return servicespec.AllServiceLabels, cobra.ShellCompDirectiveNoSpace
	})

	return cmd
}

func runBind(opts *options) error {
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

	kubeClients, err := kubeclient.NewKubernetesClusterClients(&cliProperties, opts.kubeconfigLocation)
	if err != nil {
		return err
	}

	clusterAPI := cluster.KubernetesClusterAPIImpl{
		KubernetesClients:  kubeClients,
		CommandEnvironment: &cliProperties,
	}

	err = clusterAPI.ExecuteServiceBinding(&v1alpha.BindOperationOptions{
		Namespace:               opts.namespace,
		ServiceName:             opts.serviceName,
		ServiceType:             opts.serviceType,
		AppName:                 opts.appName,
		ForceCreationWithoutAsk: opts.forceCreationWithoutAsk,
		BindingName:             opts.bindingName,
		BindAsFiles:             !opts.bindAsEnv,
		DeploymentConfigEnabled: opts.deploymentConfigEnabled,
		IgnoreContext:           opts.ignoreContext,
	})

	return err
}
