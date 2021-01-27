package status

import (
	"context"
	"encoding/json"
	"os"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/dump"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/iostreams"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"
	pkgStatus "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/status"

	"github.com/MakeNowJust/heredoc"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type Options struct {
	IO         *iostreams.IOStreams
	Config     config.IConfig
	Logger     func() (logging.Logger, error)
	Connection func() (connection.Connection, error)

	outputFormat string
}

func NewStatusCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		IO:         f.IOStreams,
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
	}

	cmd := &cobra.Command{
		Use:   "status",
		Short: "Create, view, use and manage your Kafka instances",
		Long: heredoc.Doc(`
			Perform various operations on your Kafka instances.
		`),
		Example: heredoc.Doc(`
			# create a Kafka instance
			$ rhoas kafka create

			# list Kafka instances
			$ rhoas kafka list

			# create a Kafka topic
			$ rhoas kafka topics create --name "my-kafka-topic"
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runStatus(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "", "Format to display the Kafka instance. Choose from: \"json\", \"yaml\", \"yml\"")

	return cmd
}

func runStatus(opts *Options) error {
	pkgOpts := &pkgStatus.Options{
		Config:     opts.Config,
		Connection: opts.Connection,
		Logger:     opts.Logger,
	}

	status, ok, err := pkgStatus.Get(context.Background(), pkgOpts)
	if err != nil {
		return err
	}

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	if !ok {
		logger.Info("No services are currently being used")
		return nil
	}

	stdout := opts.IO.Out
	switch opts.outputFormat {
	case "json":
		data, _ := json.Marshal(status)
		_ = dump.JSON(stdout, data)
		return nil
	case "yaml", "yml":
		data, _ := yaml.Marshal(status)
		_ = dump.YAML(os.Stdout, data)
		return nil
	}

	pkgStatus.Print(stdout, status)

	return nil
}
