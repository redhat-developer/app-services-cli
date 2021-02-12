package browser

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
)

func GetOpenBrowserCommand(url string) (*exec.Cmd, error) {
	switch runtime.GOOS {
	case "linux":
		return exec.Command("xdg-open", url), nil
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", url), nil
	case "darwin":
		return exec.Command("open", url), nil
	default:
		localizer.LoadMessageFiles("browser")
		return nil, fmt.Errorf(localizer.MustLocalize(&localizer.Config{
			MessageID: "browser.getOpenBrowserCommand.error.unsupportedOperatingSystem",
			TemplateData: map[string]interface{}{
				"OS": runtime.GOOS,
			},
		}))
	}
}
