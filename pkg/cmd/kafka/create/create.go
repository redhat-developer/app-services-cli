package create

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"

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

			return runCreate(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "json", "Format to display the Kafka clusters. Choose from: \"json\", \"yaml\", \"yml\"")
	cmd.Flags().StringVar(&opts.name, flags.FlagName, "", "Name of the new Kafka cluster")
	cmd.Flags().StringVar(&opts.provider, flags.FlagProvider, "aws", "Cloud provider ID")
	cmd.Flags().StringVar(&opts.region, flags.FlagRegion, "us-east-1", "Cloud Provider Region ID")
	cmd.Flags().BoolVar(&opts.multiAZ, flags.FlagMultiAZ, false, "Determines if cluster should be provisioned across multiple Availability Zones")

	_ = cmd.MarkFlagRequired(flags.FlagName)

	return cmd
}

func runCreate(opts *options) error {
	connection, err := opts.cfg.Connection()
	if err != nil {
		return fmt.Errorf("Can't create connection: %w", err)
	}

	client := connection.NewMASClient()

	kafkaRequest := managedservices.KafkaRequest{Name: opts.name, Region: opts.region, CloudProvider: opts.provider, MultiAz: opts.multiAZ}
	response, status, err := client.DefaultApi.CreateKafka(context.Background(), true, kafkaRequest)

	if err != nil {
		return fmt.Errorf("Error while requesting new Kafka cluster: %w", err)
	}

	if status.StatusCode != 202 {
		fmt.Fprintf(os.Stderr, "Could not create Kafka cluster: %v", response)
	}

	jsonResponse, _ := json.MarshalIndent(response, "", cmdutil.DefaultJSONIndent)
	fmt.Fprintf(os.Stderr, "Created new Kafka cluster:\n")
	fmt.Print(string(jsonResponse))

	return nil
}
