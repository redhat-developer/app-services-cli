package kafka

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewGetCommand gets a new command for getting kafkas.
func NewTopicsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "topics",
		Short: "Manage topics",
		Long:  "Manage topics",
		Run:   runTopics,
	}
	cmd.Flags().String("name", "", "topic name")
	cmd.Flags().String("operation", "", "create, delete, update")
	return cmd
}

func runTopics(cmd *cobra.Command, _ []string) {
	fmt.Print("TODO")
}
