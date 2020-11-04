package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	mas "gitlab.cee.redhat.com/mas-dx/rhmas/client/mas"
	"gitlab.cee.redhat.com/mas-dx/rhmas/cmd/flags"
)

// NewCreateCommand creates a new command for creating kafkas.
func NewCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create OpenShift Kafka instance",
		Long:  "Create OpenShift Kafka instance",
		Run:   runCreate,
	}

	cmd.Flags().String(FlagName, "", "Kafka request name")
	cmd.Flags().String(FlagRegion, "eu-west-1", "Region ID")
	cmd.Flags().String(FlagProvider, "aws", "OCM provider ID")
	cmd.Flags().Bool(FlagMultiAZ, false, "Whether Kafka request should be Multi AZ or not")

	return cmd
}

func runCreate(cmd *cobra.Command, _ []string) {
	name := flags.MustGetDefinedString(FlagName, cmd.Flags())
	region := flags.MustGetDefinedString(FlagRegion, cmd.Flags())
	provider := flags.MustGetDefinedString(FlagProvider, cmd.Flags())
	multiAZ := flags.MustGetBool(FlagMultiAZ, cmd.Flags())

	client := BuildMasClient()

	kafkaRequest := mas.KafkaRequest{Name: name, Region: region, CloudProvider: provider, MultiAz: multiAZ}
	// data, _ := json.Marshal(kafkaRequest)
	// fmt.Print("kafkaRequest API", string(data))
	response, status, err := client.DefaultApi.ApiManagedServicesApiV1KafkasPost(context.Background(), false, kafkaRequest)

	if err != nil {
		glog.Fatalf("Error while requesting new Kafka instance: %v", err)
	}
	if status.StatusCode == 200 {
		jsonResponse, _ := json.MarshalIndent(response, "", "  ")
		fmt.Print("Created Cluster \n ", string(jsonResponse))
	} else {
		fmt.Print("Creation failed", response, status)
	}
}
