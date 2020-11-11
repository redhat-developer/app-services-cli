package authz

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func NewAuthzViewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "view",
		Short: "View authorization rules",
		Long:  "View authorization rules",
		Run:   runAuthzView,
	}

	return cmd
}

func runAuthzView(cmd *cobra.Command, _ []string) {
	fmt.Fprintln(os.Stderr, "No available authorization rules")
}
