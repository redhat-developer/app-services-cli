package kafka

import (
	"os"
	"encoding/json"
	"context"
	"fmt"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/config"
	"github.com/spf13/cobra"
)

func NewStatusCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Get status of current Kafka cluster",
		Long:  "Gets the status of the current Kafka cluster context",
		Run:   runStatus,
	}
	return cmd
}

func runStatus(cmd *cobra.Command, args []string) {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}

	id := cfg.Services.Kafka.ClusterID

	if id == "" {
		fmt.Fprint(os.Stderr, "No Kafka cluster is being used. To use a cluster run `rhmas kafka use {clusterId}`")
		return
	}

	client := BuildMasClient()

	res, status, err := client.DefaultApi.ApiManagedServicesApiV1KafkasIdGet(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error retrieving Kafka cluster \"%v\": %v", id, err)
		return
	}

	if status.StatusCode != 200 {
		fmt.Fprintf(os.Stderr, "Unable to retrieve selected Kafka cluster \"%v\": %v", id, err)
		return
	}

	jsonCluster, _ := json.MarshalIndent(res, "", "  ")

	fmt.Fprintf(os.Stderr, "Using Kafka cluster \"%v\":\n", res.Id)
	fmt.Print(string(jsonCluster))
}
