package describe

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/serviceapi/client"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil/flags"

	"github.com/MakeNowJust/heredoc"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/dump"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/iostreams"
	pkgKafka "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/kafka"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type options struct {
	id           string
	outputFormat string
	fieldName    string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection func() (connection.Connection, error)
	Logger     func() (logging.Logger, error)
}

// possible values for the field argument
const (
	idFieldCmd           = "id"
	bootstrapURLFieldCmd = "bootstrap-server-host"
	statusFieldCmd       = "status"
	nameFieldCmd         = "name"
)

// NewDescribeCommand describes a Kafka instance, either by passing an `--id flag`
// or by using the kafka instance set in the config, if any
func NewDescribeCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
	}

	cmd := &cobra.Command{
		Use:   "describe",
		Short: "View all configuration values of a Kafka instance",
		Long: heredoc.Docf(`
			View all configuration fields and their values for a Kafka instance.

			Pass the --id flag to specify which instance you would like to view.
			If the --id flag is not passed then the selected Kafka instance will be used, if available.
			The result can be output either as JSON or YAML.

			Pass one of "%v", "%v", "%v" or "%v" as an argument to request 
			that single field instead of printing the full instance object.
		`, nameFieldCmd, idFieldCmd, bootstrapURLFieldCmd, statusFieldCmd),
		Example: heredoc.Doc(`
			# view the current Kafka instance instance
			$ rhoas kafka describe

			# view a specific instance by passing the --id flag
			$ rhoas kafka describe --id=1iSY6RQ3JKI8Q0OTmjQFd3ocFRg

			# customise the output format
			$ rhoas kafka describe -o yaml

			# retrieve the bootstrap server url
			$ rhoas kafka describe bootstrap-server-host
			`,
		),
		Args:      cobra.RangeArgs(0, 1),
		ValidArgs: []string{idFieldCmd, bootstrapURLFieldCmd, statusFieldCmd, nameFieldCmd},
		RunE: func(cmd *cobra.Command, args []string) error {
			logger, err := opts.Logger()
			if err != nil {
				return err
			}

			// check if an argument has been passed
			// and perform validation checks
			if len(args) > 0 {
				opts.fieldName = args[0]

				if !flags.IsValidInput(opts.fieldName, cmd.ValidArgs...) {
					return fmt.Errorf("Invalid argument '%v'. Valid values: %q", opts.fieldName, cmd.ValidArgs)
				}

				if opts.outputFormat != "" {
					logger.Debugf("--output=%v has no effect when a field argument is passed", opts.outputFormat)
				}
			}

			if opts.fieldName != "" && opts.outputFormat != "json" && opts.outputFormat != "yaml" && opts.outputFormat != "yml" {
				return fmt.Errorf("Invalid output format '%v'", opts.outputFormat)
			}

			if opts.id != "" {
				return runDescribe(opts)
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			var kafkaConfig *config.KafkaConfig
			if cfg.Services.Kafka == kafkaConfig || cfg.Services.Kafka.ClusterID == "" {
				return fmt.Errorf("No Kafka instance selected. Use the '--id' flag or set one in context with the 'use' command")
			}

			opts.id = cfg.Services.Kafka.ClusterID

			return runDescribe(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "json", "Format to display the Kafka instance. Choose from: \"json\", \"yaml\", \"yml\"")
	cmd.Flags().StringVar(&opts.id, "id", "", "ID of the Kafka instance you want to describe. If not set, the current Kafka instance will be used")

	return cmd
}

func runDescribe(opts *options) error {
	connection, err := opts.Connection()
	if err != nil {
		return err
	}

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	api := connection.API()

	response, _, apiErr := api.Kafka.GetKafkaById(context.Background(), opts.id).Execute()

	if apiErr.Error() != "" {
		return fmt.Errorf("Unable to get Kafka instance: %w", apiErr)
	}

	pkgKafka.TransformKafkaRequest(&response)

	if opts.fieldName != "" {
		return printField(opts.IO.Out, logger, opts.fieldName, &response)
	}

	switch opts.outputFormat {
	case "json":
		data, _ := json.MarshalIndent(response, "", cmdutil.DefaultJSONIndent)
		_ = dump.JSON(os.Stdout, data)
	case "yaml", "yml":
		data, _ := yaml.Marshal(response)
		_ = dump.YAML(os.Stdout, data)
	}

	return nil
}

func printField(w io.Writer, logger logging.Logger, fieldName string, kafka *client.KafkaRequest) error {
	logger.Debugf("Printing '%v' field requested from Kafka instance '%v'", fieldName, kafka.GetName())
	logger.Info("")
	switch fieldName {
	case idFieldCmd:
		fmt.Fprintln(w, kafka.GetId())
	case nameFieldCmd:
		fmt.Fprintln(w, kafka.GetName())
	case statusFieldCmd:
		fmt.Fprintln(w, kafka.GetStatus())
	case bootstrapURLFieldCmd:
		fmt.Fprintln(w, kafka.GetBootstrapServerHost())
	}

	return nil
}
