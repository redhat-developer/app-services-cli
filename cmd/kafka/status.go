package kafka

import (
	"context"
	"fmt"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/config"
	"github.com/golang/glog"

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
		glog.Fatal(err)
	}

	id := cfg.Services.Kafka.ClusterID

	if id == "" {
		fmt.Println("No Kafka cluster is being used. To use a cluster run `rhmas kafka use {clusterId}`")
		return
	}

	client := BuildMasClient()

	res, status, err := client.DefaultApi.ApiManagedServicesApiV1KafkasIdGet(context.Background(), id)
	if err != nil {
		fmt.Printf("Error retrieving Kafka cluster \"%v\": %v", id, err)
		return
	}

	if status.StatusCode != 200 {
		fmt.Printf("Unable to retrieve selected Kafka cluster \"%v\": %v", id, err)
		return
	}

	//jsonCluster, _ := json.MarshalIndent(res, "", "  ")

	fmt.Printf("Using Kafka cluster \"%v\":\n", res.Id)
	// fmt.Print(string(jsonCluster))
}
