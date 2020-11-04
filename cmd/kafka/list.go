package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/antihax/optional"
	"github.com/spf13/cobra"
	"gitlab.cee.redhat.com/mas-dx/rhmas/cmd/flags"
	"gitlab.cee.redhat.com/mas-dx/rhmas/pkg/kafka"

	mas "gitlab.cee.redhat.com/mas-dx/rhmas/client/mas"
)

const (
	FlagFormat = "format"
	FlagPage   = "page"
	FlagSize   = "size"
)

// NewListCommand creates a new command for listing kafkas.
func NewListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "lists all managed kafka requests",
		Long:  "lists all managed kafka requests",
		Run:   runList,
	}

	cmd.PersistentFlags().String(FlagFormat, "table", "Format to display the Kafka instances. Choose from \"json\" or \"table\"")
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
		fmt.Printf("Error retrieving Kafka instances: %v", err)
		return
	}

	if status.StatusCode == 200 {
		displayFormat := flags.GetString("format", cmd.Flags())
		jsonResponse, _ := json.Marshal(response)

		var kafkaList kafka.List
		if err = json.Unmarshal(jsonResponse, &kafkaList); err != nil {
			fmt.Print(err)
			displayFormat = "json"
		}


		switch displayFormat {
		case "json":
			data, _ := json.MarshalIndent(kafkaList.Items, "", "  ")
			fmt.Print(string(data))
		default:
			kafka.PrintInstances(kafkaList.Items)
		}
	}
}