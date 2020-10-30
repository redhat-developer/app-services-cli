package kafka

import (
	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

// NewDeleteCommand command for deleting kafkas.
func NewDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a managed-services-api kafka request",
		Long:  "Delete a managed-services-api kafka request.",
		Run:   runDelete,
	}

	cmd.Flags().String(FlagID, "", "Kafka id")

	return cmd
}

func runDelete(cmd *cobra.Command, _ []string) {
	//id := flags.MustGetDefinedString(FlagID, cmd.Flags())

	glog.V(10).Infof("Deleted kafka request with id %s", 1)
}
