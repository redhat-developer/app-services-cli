package bind

import (
	"context"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cluster"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/kafka"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	"github.com/spf13/cobra"
)

type options struct {
	Config     config.IConfig
	Connection func(connectionCfg *connection.Config) (connection.Connection, error)
	Logger     logging.Logger
	IO         *iostreams.IOStreams
	localizer  localize.Localizer
	Context    context.Context

	kubeconfigLocation string
	namespace          string

	forceCreationWithoutAsk bool
	ignoreContext           bool
	appName                 string
	selectedKafka           string

	deploymentConfigEnabled bool
	bindAsEnv               bool
	bindingName             string
}

func NewBindCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
		Context:    f.Context,
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
				return opts.localizer.MustLocalizeError("flag.error.requiredWhenNonInteractive", localize.NewEntry("Flag", "appName"))
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
	return cmd
}

func runBind(opts *options) error {
	apiConnection, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	cfg, err := opts.Config.Load()
	if err != nil {
		return err
	}

	// Multiservice support: use cluster CR's instead of all kafkas
	if cfg.Services.Kafka == nil || opts.ignoreContext {
		// nolint:govet
		selectedKafka, err := kafka.InteractiveSelect(opts.Context, apiConnection, opts.Logger)
		if err != nil {
			return err
		}
		if selectedKafka == nil {
			return nil
		}
		opts.selectedKafka = selectedKafka.GetId()
	} else {
		opts.selectedKafka = cfg.Services.Kafka.ClusterID
	}

	api := apiConnection.API()
	kafkaInstance, _, err := api.Kafka().GetKafkaById(opts.Context, opts.selectedKafka).Execute()
	if err != nil {
		return err
	}

	if kafkaInstance.Name == nil {
		return opts.localizer.MustLocalizeError("cluster.bind.error.emptyResponse")
	}

	err = cluster.ExecuteServiceBinding(opts.Context, opts.Logger, opts.localizer, &cluster.ServiceBindingOptions{
		ServiceName:             kafkaInstance.GetName(),
		Namespace:               opts.namespace,
		AppName:                 opts.appName,
		ForceCreationWithoutAsk: opts.forceCreationWithoutAsk,
		BindingName:             opts.bindingName,
		BindAsFiles:             !opts.bindAsEnv,
		DeploymentConfigEnabled: opts.deploymentConfigEnabled,
	})

	return err
}
