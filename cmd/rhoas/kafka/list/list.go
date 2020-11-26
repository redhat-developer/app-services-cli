package list

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/antihax/optional"
	"github.com/spf13/cobra"

	"github.com/bf2fc6cc711aee1a0c2a/cli/cmd/rhoas/cmdutil/flags"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/builders"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/managedservices"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/kafka"
)

const (
	FlagPage = "page"
	FlagSize = "size"
)

var outputFormat string

// NewListCommand creates a new command for listing kafkas.
func NewListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all Kafka clusters",
		Long:  "List all Kafka clusters",
		Run:   runList,
	}

	cmd.Flags().StringVarP(&outputFormat, "output", "o", "table", "Format to display the Kafka clusters. Choose from \"json\" or \"table\"")
	cmd.Flags().String(FlagPage, "1", "Page index")
	cmd.Flags().String(FlagSize, "100", "Number of kafka requests per page")

	return cmd
}

func runList(cmd *cobra.Command, _ []string) {
	page := flags.GetString(FlagPage, cmd.Flags())
	size := flags.GetString(FlagSize, cmd.Flags())

	client := builders.BuildClient()
	options := managedservices.ListKafkasOpts{Page: optional.NewString(page), Size: optional.NewString(size)}
	response, _, err := client.DefaultApi.ListKafkas(context.Background(), &options)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error retrieving Kafka clusters: %v\n", err)
		os.Exit(1)
	}

	if response.Size == 0 {
		fmt.Fprintln(os.Stderr, "No Kafka clusters found.")
		return
	}

	jsonResponse, _ := json.Marshal(response)

	var kafkaList kafka.ClusterList

	if err = json.Unmarshal(jsonResponse, &kafkaList); err != nil {
		fmt.Fprintf(os.Stderr, "Could not unmarshal Kakfa items into table: %v", err)
		outputFormat = "json"
	}

	switch outputFormat {
	case "json":
		data, _ := json.MarshalIndent(kafkaList.Items, "", "  ")
		fmt.Print(string(data))
	default:
		kafka.PrintToTable(kafkaList.Items)
	}
}
