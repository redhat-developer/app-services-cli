package kafka

import (
	"os"
	"context"
	"encoding/json"
	"fmt"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/kafka"
	"github.com/spf13/cobra"
)

// NewGetCommand gets a new command for getting kafkas.
func NewGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get [Kafka cluster ID]",
		Short: "Get details of single Kafka cluster",
		Long:  "Get details of single Kafka cluster",
		Run:   runGet,
	}

	return cmd
}

func runGet(cmd *cobra.Command, args []string) {
	id := ""

	if len(args) > 0 {
		// TODO: Determine if it is an ID or name
		id = args[0]
	} else {
		// TODO: Get ID of current cluster
		fmt.Fprintln(os.Stderr, "No Kafka cluster selected")
		return
	}

	client := BuildMasClient()

	response, status, err := client.DefaultApi.ApiManagedServicesApiV1KafkasIdGet(context.Background(), id)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error retrieving Kafka clusters: %v", err)
		return
	}

	if status.StatusCode == 200 {
		jsonResponse, _ := json.MarshalIndent(response, "", "  ")
		var kafkaCluster kafka.Cluster
		json.Unmarshal(jsonResponse, &kafkaCluster)
		fmt.Print(string(jsonResponse))
	} else {
		fmt.Fprintln(os.Stderr, "Get failed", response, status)
	}
}
