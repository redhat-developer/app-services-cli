package browser

import (
	"errors"
	"os/exec"
	"runtime"

	"github.com/redhat-developer/app-services-cli/internal/localizer"
)

// Open opens the URL in the default browser
func Open(url string) error {
	switch runtime.GOOS {
	case "linux":
		return exec.Command("xdg-open", url).Run()
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Run()
	case "darwin":
		return exec.Command("open", url).Run()
	default:
		return errors.New(localizer.MustLocalize(&localizer.Config{
			MessageID: "browser.getOpenBrowserCommand.error.unsupportedOperatingSystem",
			TemplateData: map[string]interface{}{
				"OS": runtime.GOOS,
			},
		}))
	}
}
