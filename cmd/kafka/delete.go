package kafka

import (
	"context"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"gitlab.cee.redhat.com/mas-dx/rhmas/cmd/flags"
)

// NewDeleteCommand command for deleting kafkas.
func NewDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete kafka cluster",
		Long:  "Request deletion of the kafka cluster",
		Run:   runDelete,
	}

	cmd.Flags().String(FlagID, "", "Kafka id")

	return cmd
}

func runDelete(cmd *cobra.Command, _ []string) {
	id := flags.MustGetDefinedString(FlagID, cmd.Flags())

	client := BuildMasClient()

	response, status, err := client.DefaultApi.ApiManagedServicesApiV1KafkasIdDelete(context.Background(), id)

	if err != nil {
		glog.Fatalf("Error while deleting Kafka instance: %v", err)
	}
	if status.StatusCode == 200 {
		glog.Info("Deleted Kafka \n ", id)
	} else {
		glog.Info("Deletion failed", response, status)
	}
}
