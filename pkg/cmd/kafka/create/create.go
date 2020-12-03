package create

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/managedservices"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/flags"
	cmdflags "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil/flags"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/config"
)

// NewCreateCommand creates a new command for creating kafkas.
func NewCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create Kafka cluster",
		Long:  "Create Kafka cluster",
		Run:   runCreate,
	}

	cmd.Flags().String(flags.FlagName, "", "Name of Kafka cluster")
	cmd.Flags().String(flags.FlagProvider, "aws", "Cloud provider ID [aws]")
	cmd.Flags().String(flags.FlagRegion, "us-east-1", "Cloud Provider Region ID (us-east-1)")
	cmd.Flags().Bool(flags.FlagMultiAZ, false, "Determines if cluster should be provisioned across multiple Availability Zones")

	return cmd
}

func runCreate(cmd *cobra.Command, _ []string) {
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

	name := cmdflags.MustGetDefinedString(flags.FlagName, cmd.Flags())
	region := cmdflags.MustGetDefinedString(flags.FlagRegion, cmd.Flags())
	provider := cmdflags.MustGetDefinedString(flags.FlagProvider, cmd.Flags())
	multiAZ := cmdflags.MustGetBool(flags.FlagMultiAZ, cmd.Flags())

	kafkaRequest := managedservices.KafkaRequest{Name: name, Region: region, CloudProvider: provider, MultiAz: multiAZ}
	response, status, err := client.DefaultApi.CreateKafka(context.Background(), true, kafkaRequest)

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
