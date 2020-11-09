package kafka

import (
	"os"
	"context"
	"encoding/json"
	"fmt"

	mas "github.com/bf2fc6cc711aee1a0c2a/cli/client/mas"
	"github.com/bf2fc6cc711aee1a0c2a/cli/cmd/flags"
	"github.com/spf13/cobra"
)

// NewCreateCommand creates a new command for creating kafkas.
func NewCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create Kafka cluster",
		Long:  "Create Kafka cluster",
		Run:   runCreate,
	}

	cmd.Flags().String(FlagName, "", "Name of Kafka cluster")
	cmd.Flags().String(FlagProvider, "aws", "Cloud provider ID [aws]")
	cmd.Flags().String(FlagRegion, "eu-west-1", "Cloud Provider Region ID (eu-west-1)")
	cmd.Flags().Bool(FlagMultiAZ, false, "Determines if cluster should be provisioned across multiple Availability Zones")

	return cmd
}

func runCreate(cmd *cobra.Command, _ []string) {
	name := flags.MustGetDefinedString(FlagName, cmd.Flags())
	region := flags.MustGetDefinedString(FlagRegion, cmd.Flags())
	provider := flags.MustGetDefinedString(FlagProvider, cmd.Flags())
	multiAZ := flags.MustGetBool(FlagMultiAZ, cmd.Flags())

	client := BuildMasClient()

	kafkaRequest := mas.KafkaRequest{Name: name, Region: region, CloudProvider: provider, MultiAz: multiAZ}
	response, status, err := client.DefaultApi.ApiManagedServicesApiV1KafkasPost(context.Background(), true, kafkaRequest)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while requesting new Kafka cluster: %v", err)
		return
	}
	if status.StatusCode == 202 {
		jsonResponse, _ := json.MarshalIndent(response, "", "  ")
		fmt.Fprintf(os.Stderr, "Created new Kakfa cluster \"%v\"\n", response.Name)
		fmt.Print(string(jsonResponse))
	} else {
		fmt.Fprintf(os.Stderr, "Failed to create Kafka cluster \"%v\"", name)
	}
}
