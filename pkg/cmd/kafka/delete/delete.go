package delete

import (
	"context"
	"fmt"
	"os"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/config"
	"github.com/spf13/cobra"
)

// NewDeleteCommand command for deleting kafkas.
func NewDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete [Kafka ID or name]",
		Short: "Delete Kafka cluster",
		Long:  "Delete Kafka cluster",
		Run:   runDelete,
	}

	return cmd
}

func runDelete(cmd *cobra.Command, args []string) {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v", err)
		os.Exit(1)
	}

	connection, err := cfg.Connection()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't create connection: %v\n", err)
		os.Exit(1)
	}

	client := connection.NewMASClient()

	id := ""

	if len(args) > 0 {
		// TODO: Determine if it is an ID or name
		id = args[0]
	} else {
		// TODO: Get ID of current cluster
		fmt.Fprintln(os.Stderr, "No Kafka cluster selected")
		os.Exit(1)
	}

	response, status, err := client.DefaultApi.DeleteKafkaById(context.Background(), id)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while deleting Kafka cluster: %v", err)
	}

	if status.StatusCode == 204 {
		fmt.Fprint(os.Stderr, "Deleted Kafka cluster with ID ", id)
	} else {
		fmt.Fprintln(os.Stderr, "Deletion failed", response, status)
	}
}
