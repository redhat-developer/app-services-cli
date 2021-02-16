package connect

import (
	"context"
	"errors"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cluster"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/iostreams"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
	"github.com/spf13/cobra"
)

type Options struct {
	Config     config.IConfig
	Connection func() (connection.Connection, error)
	Logger     func() (logging.Logger, error)
	IO         *iostreams.IOStreams

	secretOnly         bool
	kubeconfigLocation string
	secretName         string
	interactiveSelect  bool
}

func NewConnectCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
	}

	localizer.LoadMessageFiles("cmd/cluster/common", "cmd/cluster/connect", "cmd/common/flags")

	cmd := &cobra.Command{
		Use:     localizer.MustLocalizeFromID("cluster.connect.cmd.use"),
		Short:   localizer.MustLocalizeFromID("cluster.connect.cmd.shortDescription"),
		Long:    localizer.MustLocalizeFromID("cluster.connect.cmd.longDescription"),
		Example: localizer.MustLocalizeFromID("cluster.connect.cmd.example"),
		RunE: func(cmd *cobra.Command, _ []string) error {
			if !opts.IO.CanPrompt() && opts.interactiveSelect {
				return errors.New(localizer.MustLocalize(&localizer.Config{
					MessageID: "flag.error.requiredWhenNonInteractive",
					TemplateData: map[string]interface{}{
						"Flag": "interactive-select",
					},
				}))
			}

			return runBind(opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.secretOnly, "secret-only", "", false, localizer.MustLocalizeFromID("cluster.connect.flag.secretOnly.description"))
	cmd.Flags().BoolVarP(&opts.interactiveSelect, "interactive-select", "", false, localizer.MustLocalizeFromID("cluster.connect.flag.interactiveSelect.description"))
	cmd.Flags().StringVarP(&opts.secretName, "secret-name", "", "kafka-credentials", localizer.MustLocalizeFromID("cluster.connect.flag.secretName.description"))
	cmd.Flags().StringVarP(&opts.kubeconfigLocation, "kubeconfig", "", "", localizer.MustLocalizeFromID("cluster.common.flag.kubeconfig.description"))

	return cmd
}

func runBind(opts *Options) error {
	connection, err := opts.Connection()
	if err != nil {
		return err
	}

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	clusterConn, err := cluster.NewKubernetesClusterConnection(connection, opts.Config, logger, opts.kubeconfigLocation)
	if err != nil {
		return err
	}

	err = clusterConn.Connect(context.Background(), opts.secretName, opts.interactiveSelect)
	if err != nil {
		return err
	}

	return nil
}
