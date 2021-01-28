package describe

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/MakeNowJust/heredoc"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/kas"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/dump"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/kafka"
	pkgKafka "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/kafka"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type options struct {
	id           string
	outputFormat string

	Config     config.IConfig
	Connection func() (connection.Connection, error)
}

// NewDescribeCommand describes a Kafka instance, either by passing an `--id flag`
// or by using the kafka instance set in the config, if any
func NewDescribeCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Config:     f.Config,
		Connection: f.Connection,
	}

	cmd := &cobra.Command{
		Use:   "describe",
		Short: "View all configuration values of a Kafka instance",
		Long: heredoc.Doc(`
			View all configuration fields and their values for a Kafka instance.

			Pass the --id flag to specify which instance you would like to view.

			If the --id flag is not passed then the selected Kafka instance will be used, if available.

			The result can be output either as JSON or YAML.
		`),
		Example: heredoc.Doc(`
			# view the current Kafka instance instance
			$ rhoas kafka describe

			# view a specific instance by passing the --id flag
			$ rhoas kafka describe --id=1iSY6RQ3JKI8Q0OTmjQFd3ocFRg

			# customise the output format
			$ rhoas kafka describe -o yaml
			`,
		),
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.outputFormat != "json" && opts.outputFormat != "yaml" && opts.outputFormat != "yml" {
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

	api := connection.API()

	response, _, apiErr := api.Kafka.GetKafkaById(context.Background(), opts.id).Execute()
	if kas.IsErr(apiErr, kas.ErrorNotFound) {
		return kafka.ErrorNotFound(opts.id)
	}

	if apiErr.Error() != "" {
		return fmt.Errorf("Unable to get Kafka instance: %w", apiErr)
	}

	pkgKafka.TransformKafkaRequest(&response)

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
