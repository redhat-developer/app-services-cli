package list

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/config"

	"gopkg.in/yaml.v2"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/managedservices"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/kafka"
)

type options struct {
	outputFormat string
}

// NewListCommand creates a new command for listing kafkas.
func NewListCommand() *cobra.Command {
	opts := &options{}

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all Kafka clusters",
		Long:  "List all Kafka clusters",
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.outputFormat != "json" && opts.outputFormat != "yaml" && opts.outputFormat != "yml" && opts.outputFormat != "table" {
				return fmt.Errorf("Invalid output format '%v'", opts.outputFormat)
			}

			return runList(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "table", "Format to display the Kafka clusters. Choose from: \"json\", \"yaml\", \"yml\", \"table\"")

	return cmd
}

func runList(opts *options) error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("Error loading config: %w", err)
	}

	connection, err := cfg.Connection()
	if err != nil {
		return fmt.Errorf("Can't create connection: %w", err)
	}

	client := connection.NewMASClient()

	options := managedservices.ListKafkasOpts{}
	response, _, err := client.DefaultApi.ListKafkas(context.Background(), &options)

	if err != nil {
		return fmt.Errorf("Error retrieving Kafka clusters: %w", err)
	}

	if response.Size == 0 {
		fmt.Fprintln(os.Stderr, "No Kafka clusters found.")
		return nil
	}

	jsonResponse, _ := json.Marshal(response)

	var kafkaList kafka.ClusterList

	outputFormat := opts.outputFormat

	if err = json.Unmarshal(jsonResponse, &kafkaList); err != nil {
		fmt.Fprintf(os.Stderr, "Could not unmarshal Kakfa items into table, defaulting to JSON: %v", err)
		outputFormat = "json"
	}

	switch outputFormat {
	case "json":
		data, _ := json.MarshalIndent(response, "", "  ")
		fmt.Print(string(data))
	case "yaml", "yml":
		data, _ := yaml.Marshal(response)
		fmt.Print(string(data))
	default:
		kafka.PrintToTable(kafkaList.Items)
	}

	return nil
}
