package kafka

import (
	"context"
	"fmt"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

// NewDeleteCommand command for deleting kafkas.
func NewDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete [Kafka ID or name]",
		Short: "Delete Kafka cluster",
		Long:  "Request deletion of a Kafka cluster",
		Run:   runDelete,
	}

	return cmd
}

func runDelete(cmd *cobra.Command, args []string) {
	id := ""

	if (len(args) > 0) {
		// TODO: Determine if it is an ID or name
		id = args[0]
	} else {
		// TODO: Get ID of current cluster
		glog.Fatalf("No Kafka instance selected")
	}

	client := BuildMasClient()

	response, status, err := client.DefaultApi.ApiManagedServicesApiV1KafkasIdDelete(context.Background(), id)

	if err != nil {
		glog.Fatalf("Error while deleting Kafka instance: %v", err)
	}
	if status.StatusCode == 200 {
		fmt.Print("Deleted Kafka cluster with ID ", id)
	} else {
		fmt.Print("Deletion failed", response, status)
	}
}
