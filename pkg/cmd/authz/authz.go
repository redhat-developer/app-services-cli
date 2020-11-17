package authz

import (
	"github.com/spf13/cobra"
)

func NewAuthGroupCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "authz",
		Short: "Authorization rules",
		Long:  "Manage Authorization rules",
	}

	// add sub-commands
	cmd.AddCommand(NewAuthzViewCommand())

	return cmd
}
