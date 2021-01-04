package create

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/sdk/kafka"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/managedservices"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/flags"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil"
)

type Options struct {
	name     string
	provider string
	region   string
	multiAZ  bool

	outputFormat string

	Config     config.IConfig
	Connection func() (connection.IConnection, error)
}

// NewCreateCommand creates a new command for creating kafkas.
func NewCreateCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,

		multiAZ: true,
	}

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new Kafka instance",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			if opts.outputFormat != "json" && opts.outputFormat != "yaml" && opts.outputFormat != "yml" {
				return fmt.Errorf("Invalid output format '%v'", opts.outputFormat)
			}

			if err := kafka.ValidateName(opts.name); err != nil {
				return err
			}

			return runCreate(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.name, flags.FlagName, "n", "", "Name of the new Kafka instance")
	cmd.Flags().StringVar(&opts.provider, flags.FlagProvider, "aws", "Cloud provider ID")
	cmd.Flags().StringVar(&opts.region, flags.FlagRegion, "us-east-1", "Cloud Provider Region ID")
	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "json", "Format to display the Kafka instance. Choose from: \"json\", \"yaml\", \"yml\"")

	_ = cmd.MarkFlagRequired(flags.FlagName)

	return cmd
}

func runCreate(opts *Options) error {
	cfg, err := opts.Config.Load()
	if err != nil {
		return fmt.Errorf("Error loading config: %w", err)
	}

	connection, err := opts.Connection()
	if err != nil {
		return fmt.Errorf("Can't create connection: %w", err)
	}

	client := connection.NewMASClient()

	fmt.Fprintln(os.Stderr, "Creating Kafka instance")

	kafkaRequest := managedservices.KafkaRequestPayload{Name: opts.name, Region: &opts.region, CloudProvider: &opts.provider, MultiAz: &opts.multiAZ}
	a := client.DefaultApi.CreateKafka(context.Background())
	a = a.KafkaRequestPayload(kafkaRequest)
	a = a.Async(true)
	response, _, apiErr := a.Execute()

	if apiErr.Error() != "" {
		return fmt.Errorf("Error while requesting new Kafka instance: %w", apiErr)
	}

	fmt.Fprintf(os.Stderr, "Created new Kafka instance:\n")

	switch opts.outputFormat {
	case "json":
		data, _ := json.MarshalIndent(response, "", cmdutil.DefaultJSONIndent)
		fmt.Println(string(data))
	case "yaml", "yml":
		data, _ := yaml.Marshal(response)
		fmt.Println(string(data))
	}

	kafkaCfg := &config.KafkaConfig{
		ClusterID: *response.Id,
	}

	cfg.Services.Kafka = kafkaCfg
	if err := opts.Config.Save(cfg); err != nil {
		return fmt.Errorf("Unable to automatically use Kafka instance: %w", err)
	}

	return nil
}
