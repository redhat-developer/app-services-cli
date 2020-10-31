// Package cluster contains commands for interacting with cluster logic of the service directly instead of through the
// REST API exposed via the serve command.
package login

import (
	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

func NewLoginCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Login to Managed Application Services",
		Long:  "Login to Managed Application Services in order to manage your services",
		Run:   runLogin,
	}

	return cmd
}

func runLogin(cmd *cobra.Command, _ []string) {
	glog.V(10).Infof("List")
}
