package status

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil/flags"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/color"
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

const (
	kafkaSvcName = "kafka"
)

var validServices = []string{kafkaSvcName}

type Options struct {
	IO         *iostreams.IOStreams
	Config     config.IConfig
	Logger     func() (logging.Logger, error)
	Connection func() (connection.Connection, error)

	outputFormat string
	services     []string
}

func NewStatusCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		IO:         f.IOStreams,
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		services:   validServices,
	}

	cmd := &cobra.Command{
		Use:   "status [args]",
		Short: "View the status of all currently used services",
		Long: heredoc.Docf(`
			View status information of your currently used services.
			Choose to view the status of all services with %v or specific services with %v

			To use a different service run %v. Example: %v.
			Services available: %v
		`, color.CodeSnippet("rhoas status"),
			color.CodeSnippet("rhoas status <service>"),
			color.CodeSnippet("rhoas <service> use"),
			color.CodeSnippet("rhoas kafka use --id=1nh3qkcXuBGMlbIPDqhNbswIZCB"),
			color.Info(fmt.Sprintf("%v", validServices))),
		Example: heredoc.Doc(`
			# view the status of all services
			$ rhoas status

			# view the status of the used Kafka
			$ rhoas status kafka

			# view the status of your services in JSON
			$ rhoas status -o json
		`),
		ValidArgs: []string{kafkaSvcName},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				for _, s := range args {
					if !flags.IsValidInput(s, validServices...) {
						return fmt.Errorf("Invalid service '%v' specified", s)
					}
				}

				opts.services = args
			}

			return runStatus(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "", "Format to display the Kafka instance. Choose from: \"json\", \"yaml\", \"yml\".")

	return cmd
}

func runStatus(opts *Options) error {
	pkgOpts := &pkgStatus.Options{
		Config:     opts.Config,
		Connection: opts.Connection,
		Logger:     opts.Logger,
		Services:   opts.services,
	}

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	if len(opts.services) > 0 {
		logger.Debug("Requesting status of the following services:", opts.services)
	}

	status, ok, err := pkgStatus.Get(context.Background(), pkgOpts)
	if err != nil {
		return err
	}

	if !ok {
		logger.Info("\nNo services are currently used.")
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
