package kafka

import (
	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

const (
	FlagPage = "page"
	FlagSize = "size"
)

// NewListCommand creates a new command for listing kafkas.
func NewListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "lists all managed kafka requests",
		Long:  "lists all managed kafka requests",
		Run:   runList,
	}

	cmd.Flags().String(FlagOwner, "test-user", "Username")
	cmd.Flags().String(FlagPage, "1", "Page index")
	cmd.Flags().String(FlagSize, "100", "Number of kafka requests per page")

	return cmd
}

func runList(cmd *cobra.Command, _ []string) {
	// owner := flags.MustGetDefinedString(FlagOwner, cmd.Flags())
	// page := flags.MustGetString(FlagPage, cmd.Flags())
	// size := flags.MustGetString(FlagSize, cmd.Flags())

	glog.V(10).Infof("List")
}
