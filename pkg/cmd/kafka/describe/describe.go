package describe

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type options struct {
	id           string
	outputFormat string

	cfg *config.Config
}

// NewDescribeCommand describes a Kafka cluster, either by passing an `--id flag`
// or by using the kafka cluster set in the config, if any
func NewDescribeCommand() *cobra.Command {
	var cfg *config.Config
	opts := &options{
		cfg: cfg,
	}

	cmd := &cobra.Command{
		Use:   "describe",
		Short: "Get details of single Kafka cluster",
		Long:  "Get details of single Kafka cluster, either by loading the currently selected Kafka cluster or a custom one with the `--id` flag",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.outputFormat != "json" && opts.outputFormat != "yaml" && opts.outputFormat != "yml" {
				return fmt.Errorf("Invalid output format '%v'", opts.outputFormat)
			}

			cfg, err := config.Load()
			if err != nil {
				return fmt.Errorf("Error loading config: %w", err)
			}
			opts.cfg = cfg

			if opts.id != "" {
				return runDescribe(opts)
			}

			var kafkaConfig *config.KafkaConfig
			if cfg.Services.Kafka == kafkaConfig || cfg.Services.Kafka.ClusterID == "" {
				return fmt.Errorf("Please select a Kafka cluster to describe with the '--id' flag, or set one with the 'use' command")
			}

			opts.id = cfg.Services.Kafka.ClusterID

			return runDescribe(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "json", "Format to display the Kafka clusters. Choose from: \"json\", \"yaml\", \"yml\"")
	cmd.Flags().StringVar(&opts.id, "id", "", "ID of the Kafka cluster you want to describe. If not set, the currently selected Kafka cluster will be used")

	return cmd
}

func runDescribe(opts *options) error {
	connection, err := opts.cfg.Connection()
	if err != nil {
		return fmt.Errorf("Can't create connection: %w", err)
	}

	client := connection.NewMASClient()

	response, status, err := client.DefaultApi.GetKafkaById(context.Background(), opts.id)

	if err != nil {
		return fmt.Errorf("Error getting Kafka cluster: %w", err)
	}

	if status.StatusCode != 200 {
		fmt.Fprintf(os.Stderr, "Unable to describe Kafka cluster with ID '%v': %v\n", opts.id, response)
		return nil
	}

	switch opts.outputFormat {
	case "json":
		data, _ := json.MarshalIndent(response, "", cmdutil.DefaultJSONIndent)
		fmt.Print(string(data))
	case "yaml", "yml":
		data, _ := yaml.Marshal(response)
		fmt.Print(string(data))
	}

	return nil
}
