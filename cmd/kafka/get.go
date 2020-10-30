package kafka

import (
	"github.com/golang/glog"
	"github.com/spf13/cobra"
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
	// id := flags.MustGetDefinedString(FlagID, cmd.Flags())

	glog.V(10).Infof("get")
}
