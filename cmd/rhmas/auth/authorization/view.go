// Package cluster contains commands for interacting with cluster logic of the service directly instead of through the
// REST API exposed via the serve command.
package authorization

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
		Run:   runAuthz,
	}

	return cmd
}

func runAuthz(cmd *cobra.Command, _ []string) {
	fmt.Fprintln(os.Stderr, "No available authorization rules")
}
