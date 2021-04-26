package connect

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

	offlineAccessToken      string
	forceCreationWithoutAsk bool
	ignoreContext           bool
	selectedKafka           string
}

func NewConnectCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
	}

	cmd := &cobra.Command{
		Use:     localizer.MustLocalizeFromID("cluster.connect.cmd.use"),
		Short:   localizer.MustLocalizeFromID("cluster.connect.cmd.shortDescription"),
		Long:    localizer.MustLocalizeFromID("cluster.connect.cmd.longDescription"),
		Example: localizer.MustLocalizeFromID("cluster.connect.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if opts.ignoreContext == true && !opts.IO.CanPrompt() {
				return errors.New(localizer.MustLocalize(&localizer.Config{
					MessageID: "flag.error.requiredWhenNonInteractive",
					TemplateData: map[string]interface{}{
						"Flag": "ignore-context",
					},
				}))
			}
			return runConnect(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.kubeconfigLocation, "kubeconfig", "", "", localizer.MustLocalizeFromID("cluster.common.flag.kubeconfig.description"))
	cmd.Flags().StringVarP(&opts.offlineAccessToken, "token", "", "", localizer.MustLocalizeFromID("cluster.common.flag.offline.token.description"))
	cmd.Flags().StringVarP(&opts.namespace, "namespace", "n", "", localizer.MustLocalizeFromID("cluster.common.flag.namespace.description"))
	cmd.Flags().BoolVarP(&opts.forceCreationWithoutAsk, "yes", "y", false, localizer.MustLocalizeFromID("cluster.common.flag.yes.description"))
	cmd.Flags().BoolVarP(&opts.ignoreContext, "ignore-context", "", false, localizer.MustLocalizeFromID("cluster.common.flag.ignoreContext.description"))

	return cmd
}

func runConnect(opts *Options) error {
	connection, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	clusterConn, err := cluster.NewKubernetesClusterConnection(connection, opts.Config, logger, opts.kubeconfigLocation, opts.IO)
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
		selectedKafka, err := kafka.InteractiveSelect(connection, logger)
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

	arguments := &cluster.ConnectArguments{
		OfflineAccessToken:      opts.offlineAccessToken,
		ForceCreationWithoutAsk: opts.forceCreationWithoutAsk,
		IgnoreContext:           opts.ignoreContext,
		SelectedKafka:           opts.selectedKafka,
		Namespace:               opts.namespace,
	}

	err = clusterConn.Connect(context.Background(), arguments)
	if err != nil {
		return err
	}

	return nil
}
