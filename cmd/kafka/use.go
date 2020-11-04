package kafka

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewUseCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "use [Kafka ID]",
		Short: "Use provided cluster",
		Long:  "Set to work with cluster on current context",
		Args:  cobra.MinimumNArgs(1),
		Run:   runUse,
	}
	return cmd
}

func runUse(cmd *cobra.Command, args []string) {
	id := args[0]

	fmt.Print("Selected Kafka cluster: ", id)
}
