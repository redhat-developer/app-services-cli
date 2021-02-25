package connect

import (
	"context"
	"errors"
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cluster"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/kafka"
	"github.com/spf13/cobra"
)

func NewConnectCommand(f *factory.Factory) *cobra.Command {
	opts := &cluster.CommandConnectOptions{
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
		RunE: func(cmd *cobra.Command, _ []string) error {
			if opts.IgnoreConfig == true && !opts.IO.CanPrompt() {
				return errors.New(localizer.MustLocalize(&localizer.Config{
					MessageID: "flag.error.requiredWhenNonInteractive",
					TemplateData: map[string]interface{}{
						"Flag": "ignore-config",
					},
				}))
			}
			return runBind(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.KubeconfigLocation, "kubeconfig", "", "", localizer.MustLocalizeFromID("cluster.common.flag.kubeconfig.description"))
	cmd.Flags().StringVarP(&opts.OfflineAccessToken, "token", "", "", localizer.MustLocalizeFromID("cluster.common.flag.offline.token.description"))
	cmd.Flags().StringVarP(&opts.Namespace, "namespace", "n", "", localizer.MustLocalizeFromID("cluster.common.flag.namespace.description"))
	cmd.Flags().BoolVarP(&opts.Force, "force", "f", false, localizer.MustLocalizeFromID("cluster.common.flag.force.description"))
	cmd.Flags().BoolVarP(&opts.IgnoreConfig, "ignore-config", "", false, localizer.MustLocalizeFromID("cluster.common.flag.ignore.config.description"))

	return cmd
}

func runBind(opts *cluster.CommandConnectOptions) error {
	connection, err := opts.Connection()
	if err != nil {
		return err
	}

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	clusterConn, err := cluster.NewKubernetesClusterConnection(connection, opts.Config, logger, opts.KubeconfigLocation)
	if err != nil {
		return err
	}

	cfg, err := opts.Config.Load()
	if err != nil {
		return err
	}

	// In future config will include Id's of other services
	if cfg.Services.Kafka == nil || opts.IgnoreConfig {
		// nolint
		selectedKafka, err := kafka.InteractiveSelect(connection, logger)
		if err != nil {

			return err
		}
		opts.SelectedKafka = *selectedKafka.Id
	} else {
		opts.SelectedKafka = cfg.Services.Kafka.ClusterID
	}

	err = clusterConn.Connect(context.Background(), opts)
	if err != nil {
		return err
	}

	return nil
}
