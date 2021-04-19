package bind

import (
	"context"
	"errors"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/internal/localizer"
	"github.com/redhat-developer/app-services-cli/pkg/cluster"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/kafka"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	"github.com/spf13/cobra"
)

type Options struct {
	Config     config.IConfig
	Connection func(connectionCfg *connection.Config) (connection.Connection, error)
	Logger     func() (logging.Logger, error)
	IO         *iostreams.IOStreams

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
	}

	cmd := &cobra.Command{
		Use:     "bind",
		Short:   localizer.MustLocalizeFromID("cluster.bind.cmd.shortDescription"),
		Long:    localizer.MustLocalizeFromID("cluster.bind.cmd.longDescription"),
		Example: localizer.MustLocalizeFromID("cluster.bind.cmd.example"),
		RunE: func(cmd *cobra.Command, _ []string) error {
			if opts.ignoreContext == true && !opts.IO.CanPrompt() {
				return errors.New(localizer.MustLocalize(&localizer.Config{
					MessageID: "flag.error.requiredWhenNonInteractive",
					TemplateData: map[string]interface{}{
						"Flag": "ignore-context",
					},
				}))
			}
			if opts.appName == "" && !opts.IO.CanPrompt() {
				return errors.New(localizer.MustLocalize(&localizer.Config{
					MessageID: "flag.error.requiredWhenNonInteractive",
					TemplateData: map[string]interface{}{
						"Flag": "appName",
					},
				}))
			}
			return runBind(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.kubeconfigLocation, "kubeconfig", "", "", localizer.MustLocalizeFromID("cluster.common.flag.kubeconfig.description"))
	cmd.Flags().StringVarP(&opts.appName, "app-name", "", "", localizer.MustLocalizeFromID("cluster.bind.flag.appName"))
	cmd.Flags().BoolVarP(&opts.forceCreationWithoutAsk, "yes", "y", false, localizer.MustLocalizeFromID("cluster.common.flag.yes.description"))
	cmd.Flags().StringVarP(&opts.namespace, "namespace", "n", "", localizer.MustLocalizeFromID("cluster.common.flag.namespace.description"))
	cmd.Flags().BoolVarP(&opts.ignoreContext, "ignore-context", "", false, localizer.MustLocalizeFromID("cluster.common.flag.ignoreContext.description"))
	cmd.Flags().BoolVarP(&opts.forceOperator, "force-operator", "", false, localizer.MustLocalizeFromID("cluster.bind.flag.forceOperator.description"))
	cmd.Flags().BoolVarP(&opts.forceSDK, "force-sdk", "", false, localizer.MustLocalizeFromID("cluster.bind.flag.forceSDK.description"))
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
		// nolint
		selectedKafka, err := kafka.InteractiveSelect(apiConnection, logger)
		if err != nil {

			return err
		}
		opts.selectedKafka = *selectedKafka.Id
	} else {
		opts.selectedKafka = cfg.Services.Kafka.ClusterID
	}

	api := apiConnection.API()
	kafkaInstance, _, err2 := api.Kafka().GetKafkaById(context.Background(), opts.selectedKafka).Execute()

	if err2 != nil {
		return err2
	}

	if kafkaInstance.Name == nil {
		return errors.New(localizer.MustLocalizeFromID("cluster.bind.error.emptyResponse"))
	}

	err = cluster.ExecuteServiceBinding(logger, &cluster.ServiceBindingOptions{
		ServiceName:             *kafkaInstance.Name,
		Namespace:               opts.namespace,
		AppName:                 opts.appName,
		ForceCreationWithoutAsk: opts.forceCreationWithoutAsk,
		ForceUseOperator:        opts.forceOperator,
		ForceUseSDK:             opts.forceSDK,
		BindingName:             opts.bindingName,
	})

	return err
}
