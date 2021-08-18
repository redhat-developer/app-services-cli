package bind

import (
	"context"
	"errors"

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

type Options struct {
	Config     config.IConfig
	Connection func(connectionCfg *connection.Config) (connection.Connection, error)
	Logger     func() (logging.Logger, error)
	IO         *iostreams.IOStreams
	localizer  localize.Localizer

	kubeconfigLocation string
	namespace          string

	forceCreationWithoutAsk bool
	ignoreContext           bool
	appName                 string
	selectedKafka           string

	forceOperator bool
	forceSDK      bool
	bindingName   string
}

func NewBindCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
	}

	cmd := &cobra.Command{
		Use:     "bind",
		Short:   opts.localizer.MustLocalize("cluster.bind.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("cluster.bind.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("cluster.bind.cmd.example"),
		RunE: func(cmd *cobra.Command, _ []string) error {
			if opts.ignoreContext == true && !opts.IO.CanPrompt() {
				return errors.New(opts.localizer.MustLocalize("flag.error.requiredWhenNonInteractive", localize.NewEntry("Flag", "ignore-context")))
			}
			if opts.appName == "" && !opts.IO.CanPrompt() {
				return errors.New(opts.localizer.MustLocalize("flag.error.requiredWhenNonInteractive", localize.NewEntry("Flag", "appName")))
			}
			return runBind(opts)
		},
	}

	cmd.Flags().StringVar(&opts.kubeconfigLocation, "kubeconfig", "", opts.localizer.MustLocalize("cluster.common.flag.kubeconfig.description"))
	cmd.Flags().StringVar(&opts.appName, "app-name", "", opts.localizer.MustLocalize("cluster.bind.flag.appName"))
	cmd.Flags().StringVarP(&opts.bindingName, "binding-name", "", opts.localizer.MustLocalize("cluster.bind.flag.bindName"))
	cmd.Flags().BoolVarP(&opts.forceCreationWithoutAsk, "yes", "y", false, opts.localizer.MustLocalize("cluster.common.flag.yes.description"))
	cmd.Flags().StringVarP(&opts.namespace, "namespace", "n", "", opts.localizer.MustLocalize("cluster.common.flag.namespace.description"))
	cmd.Flags().BoolVarP(&opts.ignoreContext, "ignore-context", "", false, opts.localizer.MustLocalize("cluster.common.flag.ignoreContext.description"))
	cmd.Flags().BoolVarP(&opts.forceOperator, "force-operator", "", false, opts.localizer.MustLocalize("cluster.bind.flag.forceOperator.description"))
	cmd.Flags().BoolVarP(&opts.forceSDK, "force-sdk", "", false, opts.localizer.MustLocalize("cluster.bind.flag.forceSDK.description"))
	return cmd
}

func runBind(opts *Options) error {
	apiConnection, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	cfg, err := opts.Config.Load()
	if err != nil {
		return err
	}

	// In future config will include Id's of other services
	if cfg.Services.Kafka == nil || opts.ignoreContext {
		// nolint:govet
		selectedKafka, err := kafka.InteractiveSelect(apiConnection, logger)
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
	kafkaInstance, _, err := api.Kafka().GetKafkaById(context.Background(), opts.selectedKafka).Execute()
	if err != nil {
		return err
	}

	if kafkaInstance.Name == nil {
		return errors.New(opts.localizer.MustLocalize("cluster.bind.error.emptyResponse"))
	}

	err = cluster.ExecuteServiceBinding(logger, opts.localizer, &cluster.ServiceBindingOptions{
		ServiceName:             kafkaInstance.GetName(),
		Namespace:               opts.namespace,
		AppName:                 opts.appName,
		ForceCreationWithoutAsk: opts.forceCreationWithoutAsk,
		ForceUseOperator:        opts.forceOperator,
		ForceUseSDK:             opts.forceSDK,
		BindingName:             opts.bindingName,
		BindAsFiles:             true,
	})

	return err
}
