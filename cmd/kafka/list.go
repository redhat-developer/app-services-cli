package kafka

import (
	"context"
	"encoding/json"

	"github.com/antihax/optional"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"gitlab.cee.redhat.com/mas-dx/rhmas/cmd/flags"

	mas "gitlab.cee.redhat.com/mas-dx/rhmas/client/mas"
)

const (
	FlagPage = "page"
	FlagSize = "size"
)

// NewListCommand creates a new command for listing kafkas.
func NewListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "lists all managed kafka requests",
		Long:  "lists all managed kafka requests",
		Run:   runList,
	}

	cmd.Flags().String(FlagPage, "1", "Page index")
	cmd.Flags().String(FlagSize, "100", "Number of kafka requests per page")

	return cmd
}

func runList(cmd *cobra.Command, _ []string) {
	page := flags.GetString(FlagPage, cmd.Flags())
	size := flags.GetString(FlagSize, cmd.Flags())

	client := BuildMasClient()
	options := mas.ApiManagedServicesApiV1KafkasGetOpts{Page: optional.NewString(page), Size: optional.NewString(size)}
	response, status, err := client.DefaultApi.ApiManagedServicesApiV1KafkasGet(context.Background(), &options)

	if err != nil {
		glog.Fatalf("Error while fetching Kafka instance: %v", err)
	}
	if status.StatusCode == 200 {
		jsonResponse, _ := json.MarshalIndent(response, "", "  ")
		glog.Info("Kafka instance \n ", string(jsonResponse))
	} else {
		glog.Info("Get failed", response, status)
	}

	glog.V(10).Infof("List")
}
