package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.cee.redhat.com/mas-dx/rhmas/cmd/flags"
	"gitlab.cee.redhat.com/mas-dx/rhmas/pkg/kafka"
)

// NewGetCommand gets a new command for getting kafkas.
func NewGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get [Kafka ID or name]",
		Short: "Get details of a managed Kafka cluster",
		Long:  "Get details of a managed Kafka cluster.",
		Run:   runGet,
	}

	cmd.PersistentFlags().String(FlagFormat, "table", "Format to display the Kafka instances. Choose from \"json\" or \"table\"")

	return cmd
}

func runGet(cmd *cobra.Command, args []string) {
	id := ""

	if len(args) > 0 {
		// TODO: Determine if it is an ID or name
		id = args[0]
	} else {
		// TODO: Get ID of current cluster
		fmt.Println("No Kafka instance selected")
		return
	}

	client := BuildMasClient()

	response, status, err := client.DefaultApi.ApiManagedServicesApiV1KafkasIdGet(context.Background(), id)

	if err != nil {
		fmt.Printf("Error retrieving Kafka instances: %v", err)
		return
	}

	if status.StatusCode == 200 {
		jsonResponse, _ := json.MarshalIndent(response, "", "  ")
		var instance kafka.Instance
		json.Unmarshal(jsonResponse, &instance)

		displayFormat := flags.GetString("format", cmd.Flags())

		switch displayFormat {
		case "json":
			fmt.Print(string(jsonResponse))
		default:
			kafka.PrintInstances([]kafka.Instance{instance})
		}
	} else {
		fmt.Print("Get failed", response, status)
	}
}
