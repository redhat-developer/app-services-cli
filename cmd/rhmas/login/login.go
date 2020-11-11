// Package cluster contains commands for interacting with cluster logic of the service directly instead of through the
// REST API exposed via the serve command.
package login

import (
	"fmt"
	"os"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/config"
	"github.com/spf13/cobra"
)

// NewLoginCommand gets the command that's log the user in
func NewLoginCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Login to Managed Application Services",
		Long:  "Login to Managed Application Services in order to manage your services",
		Run:   runLogin,
	}

	cmd.Flags().StringVar(&args.token, "token", "", "access token that can be used for login")
	rootCmd.MarkFlagRequired("token")
	cmd.Flags().StringVar(&args.tokenURL, "token-url", "", "OpenID token URL")
	cmd.Flags().StringVar(&args.url, "url", "staging", "URL of the API gateway. The value can be the complete URL or an alias. The valid aliases are 'production', 'staging', 'integration', 'development' and their shorthands.")

	return cmd
}

func runLogin(cmd *cobra.Command, _ []string) {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't load config file: %v\n", err)
		os.Exit(1)
	}
	if cfg == nil {
		cfg = new(config.Config)
	}

	// login(&cfg)
}
