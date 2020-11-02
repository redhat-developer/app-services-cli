package kafka

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewGetCommand gets a new command for getting kafkas.
func NewCredentialsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "credentials",
		Short: "Get config",
		Long:  "Get get configuration for connecting to managed kafka",
		Run:   runCredentials,
	}
	cmd.Flags().String("format", "", "Format of the config (java, json, yml")
	return cmd
}

func runCredentials(cmd *cobra.Command, _ []string) {
	fmt.Print(`Credentials for Streams instance: 'serviceapi' 
	----------------------------------------------
	user=wtrocki@redhat.com
	password=d0b8122f-8dfb-46b7-b68a-f5cc4e25d000
	----------------------------------------------

`)
}
