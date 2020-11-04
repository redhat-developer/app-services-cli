// Package cluster contains commands for interacting with cluster logic of the service directly instead of through the
// REST API exposed via the serve command.
package login

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"

	"github.com/golang/glog"

	"github.com/spf13/cobra"
	"gitlab.cee.redhat.com/mas-dx/rhmas/cmd/flags"
)

func NewLoginCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Login to Managed Application Services",
		Long:  "Login to Managed Application Services in order to manage your services",
		Run:   runLogin,
	}

	cmd.Flags().String("token", "", "access token that can be used for login")
	return cmd
}

func runLogin(cmd *cobra.Command, _ []string) {
	token := flags.GetString("token", cmd.Flags())
	if len(token) > 0 {
		fmt.Printf("Successfully logged in using token")
	} else {
		glog.Infof("Redirecting to login page")
		cmd, err := getOpenBrowserCommand("https://sso.redhat.com/auth/realms/redhat-external/protocol/openid-connect/auth?client_id=cloud-services&redirect_uri=https%3A%2F%2Fcloud.redhat.com%2F&state=d8b10b88-8699-4c9b-80fd-665c39343e53&response_mode=fragment&response_type=code&scope=openid&nonce=7ba8050f-5f7b-4a1c-80dd-0392c922f5f8")
		if err != nil {
			glog.Fatal(err)
		} else {
			cmd.Start()
			time.Sleep(30 * time.Second)
		}
	}
}

func getOpenBrowserCommand(url string) (*exec.Cmd, error) {
	switch runtime.GOOS {
	case "linux":
		return exec.Command("xdg-open", url), nil
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", url), nil
	case "darwin":
		return exec.Command("open", url), nil
	// TODO: Add more operating systems
	default:
		return nil, fmt.Errorf("Unsupported operating system: %s", runtime.GOOS)
	}
}
