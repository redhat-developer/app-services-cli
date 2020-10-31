package kafka

import (
	"context"
	"encoding/json"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"gitlab.cee.redhat.com/mas-dx/rhmas/cmd/flags"
)

// NewGetCommand gets a new command for getting kafkas.
func NewGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a managed-services-api kafka request",
		Long:  "Get a managed-services-api kafka request.",
		Run:   runGet,
	}
	cmd.Flags().String(FlagID, "", "Kafka id")
	return cmd
}

func runGet(cmd *cobra.Command, _ []string) {
	id := flags.MustGetDefinedString(FlagID, cmd.Flags())

	client := BuildMasClient()

	response, status, err := client.DefaultApi.ApiManagedServicesApiV1KafkasIdGet(context.Background(), id)

	if err != nil {
		glog.Fatalf("Error while fetching Kafka instance: %v", err)
	}
	if status.StatusCode == 200 {
		jsonResponse, _ := json.MarshalIndent(response, "", "  ")
		glog.Info("Kafka instance \n ", string(jsonResponse))
	} else {
		glog.Info("Get failed", response, status)
	}

	glog.V(10).Infof("get")
}
