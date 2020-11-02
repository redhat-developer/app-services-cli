package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

// NewGetCommand gets a new command for getting kafkas.
func NewGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get [Kafka ID or name]",
		Short: "Get a managed-services-api kafka request",
		Long:  "Get a managed-services-api kafka request.",
		Run:   runGet,
	}

	return cmd
}

func runGet(cmd *cobra.Command, args []string) {
	id := ""

	if (len(args) > 0) {
		// TODO: Determine if it is an ID or name
		id = args[0]
	} else {
		// TODO: Get ID of current cluster
		glog.Fatalf("No Kafka instance selected")
	}

	client := BuildMasClient()

	response, status, err := client.DefaultApi.ApiManagedServicesApiV1KafkasIdGet(context.Background(), id)

	if err != nil {
		glog.Fatalf("Error while fetching Kafka instance: %v", err)
	}
	if status.StatusCode == 200 {
		jsonResponse, _ := json.MarshalIndent(response, "", "  ")
		fmt.Print("Kafka instance \n ", string(jsonResponse))
	} else {
		fmt.Print("Get failed", response, status)
	}
}
