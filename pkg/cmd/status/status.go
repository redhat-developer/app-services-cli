package status

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/flag"
	flagutil "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil/flags"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil/flags"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/dump"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/iostreams"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"
	pkgStatus "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/status"

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
	Connection factory.ConnectionFunc

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
		Use:       localizer.MustLocalizeFromID("status.cmd.use"),
		Short:     localizer.MustLocalizeFromID("status.cmd.shortDescription"),
		Long:      localizer.MustLocalizeFromID("status.cmd.longDescription"),
		Example:   localizer.MustLocalizeFromID("status.cmd.example"),
		ValidArgs: []string{kafkaSvcName},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				for _, s := range args {
					if !flags.IsValidInput(s, validServices...) {
						return fmt.Errorf(localizer.MustLocalize(&localizer.Config{
							MessageID: "status.error.args.error.unknownServiceError",
							TemplateData: map[string]interface{}{
								"ServiceName": s,
							},
						}))
					}
				}

				opts.services = args
			}

			validOutputFormats := flagutil.ValidOutputFormats
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flag.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			return runStatus(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "", localizer.MustLocalizeFromID("status.flag.output.description"))

	return cmd
}

func runStatus(opts *Options) error {
	connection, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return nil
	}

	pkgOpts := &pkgStatus.Options{
		Config:     opts.Config,
		Connection: connection,
		Logger:     opts.Logger,
		Services:   opts.services,
	}

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	if len(opts.services) > 0 {
		logger.Debug(localizer.MustLocalizeFromID("status.log.debug.requestingStatusOfServices"), opts.services)
	}

	status, ok, err := pkgStatus.Get(context.Background(), pkgOpts)
	if err != nil {
		return err
	}

	if !ok {
		logger.Info("")
		logger.Info(localizer.MustLocalizeFromID("status.log.info.noStatusesAreUsed"))
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
		_ = dump.YAML(stdout, data)
		return nil
	}

	pkgStatus.Print(stdout, status)

	return nil
}
