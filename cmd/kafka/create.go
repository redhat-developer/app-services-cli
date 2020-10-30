package kafka

import (
	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

// NewCreateCommand creates a new command for creating kafkas.
func NewCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new managed-services-api kafka request",
		Long:  "Create a new managed-services-api kafka request.",
		Run:   runCreate,
	}

	cmd.Flags().String(FlagName, "", "Kafka request name")
	cmd.Flags().String(FlagRegion, "eu-west-1", "OCM region ID")
	cmd.Flags().String(FlagProvider, "aws", "OCM provider ID")
	cmd.Flags().String(FlagOwner, "test-user", "User name")
	cmd.Flags().String(FlagClusterID, "000", "Kafka  request cluster ID")
	cmd.Flags().Bool(FlagMultiAZ, false, "Whether Kafka request should be Multi AZ or not")

	return cmd
}

func runCreate(cmd *cobra.Command, _ []string) {
	// name := flags.MustGetDefinedString(FlagName, cmd.Flags())
	// region := flags.MustGetDefinedString(FlagRegion, cmd.Flags())
	// provider := flags.MustGetDefinedString(FlagProvider, cmd.Flags())
	// owner := flags.MustGetDefinedString(FlagOwner, cmd.Flags())
	// multiAZ := flags.MustGetBool(FlagMultiAZ, cmd.Flags())
	// clusterID := flags.MustGetDefinedString(FlagClusterID, cmd.Flags())

	glog.V(10).Infof("Done")
}
