package connect

import (
	"context"
	"errors"

	"github.com/redhat-developer/app-services-cli/internal/build"
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
		localizer:  f.Localizer,
	}

	cmd := &cobra.Command{
		Use:     opts.localizer.LoadMessage("cluster.connect.cmd.use"),
		Short:   opts.localizer.LoadMessage("cluster.connect.cmd.shortDescription"),
		Long:    opts.localizer.LoadMessage("cluster.connect.cmd.longDescription"),
		Example: opts.localizer.LoadMessage("cluster.connect.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if opts.ignoreContext == true && !opts.IO.CanPrompt() {
				return errors.New(opts.localizer.LoadMessage("flag.error.requiredWhenNonInteractive", localize.NewEntry("Flag", "ignore-context")))
			}
			return runConnect(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.kubeconfigLocation, "kubeconfig", "", "", opts.localizer.LoadMessage("cluster.common.flag.kubeconfig.description"))
	cmd.Flags().StringVarP(&opts.offlineAccessToken, "token", "", "", opts.localizer.LoadMessage("cluster.common.flag.offline.token.description", localize.NewEntry("OfflineTokenURL", build.OfflineTokenURL)))
	cmd.Flags().StringVarP(&opts.namespace, "namespace", "n", "", opts.localizer.LoadMessage("cluster.common.flag.namespace.description"))
	cmd.Flags().BoolVarP(&opts.forceCreationWithoutAsk, "yes", "y", false, opts.localizer.LoadMessage("cluster.common.flag.yes.description"))
	cmd.Flags().BoolVarP(&opts.ignoreContext, "ignore-context", "", false, opts.localizer.LoadMessage("cluster.common.flag.ignoreContext.description"))

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

	clusterConn, err := cluster.NewKubernetesClusterConnection(connection, opts.Config, logger, opts.kubeconfigLocation, opts.IO, opts.localizer)
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
