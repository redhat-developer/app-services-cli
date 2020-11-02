package kafka

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.cee.redhat.com/mas-dx/rhmas/cmd/flags"
)

func NewUseCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "use",
		Short: "Use provided cluster",
		Long:  "Set to work with cluster on current context",
		Run:   runUse,
	}
	cmd.Flags().String(FlagID, "", "Kafka id")
	return cmd
}

func runUse(cmd *cobra.Command, _ []string) {
	id := flags.MustGetDefinedString(FlagID, cmd.Flags())

	fmt.Print("Selected kafka cluster with ", id)
}
