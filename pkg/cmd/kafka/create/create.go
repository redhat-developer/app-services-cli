package create

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/sdk/kafka"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/managedservices"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/flags"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/config"
)

type options struct {
	name     string
	provider string
	region   string
	multiAZ  bool

	outputFormat string

	cfg *config.Config
}

// NewCreateCommand creates a new command for creating kafkas.
func NewCreateCommand() *cobra.Command {
	opts := &options{}

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new Kafka cluster",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			if opts.outputFormat != "json" && opts.outputFormat != "yaml" && opts.outputFormat != "yml" {
				return fmt.Errorf("Invalid output format '%v'", opts.outputFormat)
			}

			cfg, err := config.Load()
			if err != nil {
				return fmt.Errorf("Error loading config: %w", err)
			}
			opts.cfg = cfg

			err = kafka.ValidateName(opts.name)
			if err != nil {
				return err
			}

			return runCreate(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.name, flags.FlagName, "n", "", "Name of the new Kafka cluster")
	cmd.Flags().StringVar(&opts.provider, flags.FlagProvider, "aws", "Cloud provider ID")
	cmd.Flags().StringVar(&opts.region, flags.FlagRegion, "us-east-1", "Cloud Provider Region ID")
	// Hardcoded as only true is possible
	// cmd.Flags().BoolVar(&opts.multiAZ, flags.FlagMultiAZ, true, "Determines if cluster should be provisioned across multiple Availability Zones")
	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "json", "Format to display the Kafka cluster. Choose from: \"json\", \"yaml\", \"yml\"")

	_ = cmd.MarkFlagRequired(flags.FlagName)

	return cmd
}

func runCreate(opts *options) error {
	connection, err := opts.cfg.Connection()
	if err != nil {
		return fmt.Errorf("Can't create connection: %w", err)
	}

	client := connection.NewMASClient()

	kafkaRequest := managedservices.KafkaRequestPayload{Name: opts.name, Region: opts.region, CloudProvider: opts.provider, MultiAz: true}
	response, _, err := client.DefaultApi.CreateKafka(context.Background(), true, kafkaRequest)

	if err != nil {
		return fmt.Errorf("Error while requesting new Kafka cluster: %w", err)
	}

	fmt.Fprintf(os.Stderr, "Created new Kafka cluster:\n")

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
