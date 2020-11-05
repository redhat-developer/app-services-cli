package kafka

import (
	"context"
	"fmt"

	"github.com/golang/glog"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/config"

	"github.com/spf13/cobra"
)

func NewUseCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "use [Kafka ID]",
		Short: "Use provided cluster",
		Long:  "Set to work with cluster on current context",
		Args:  cobra.MinimumNArgs(1),
		Run:   runUse,
	}
	return cmd
}

func runUse(cmd *cobra.Command, args []string) {
	cfg, err := config.Load()
	if err != nil {
		glog.Fatal(err)
	}

	id := args[0]

	client := BuildMasClient()

	res, status, err := client.DefaultApi.ApiManagedServicesApiV1KafkasIdGet(context.Background(), id)
	if err != nil {
		fmt.Printf("Error retrieving Kafka cluster \"%v\": %v", id, err)
		return
	}

	if status.StatusCode != 200 {
		fmt.Printf("Could not use cluster \"%v\": %v", id, err)
		return
	}

	var kafkaConfig config.KafkaConfig = config.KafkaConfig{ClusterID: res.Id}
	cfg.Services.SetKafka(&kafkaConfig)
	config.Save(cfg)
	fmt.Printf("Using Kafka cluster \"%v\"", res.Id)
}
