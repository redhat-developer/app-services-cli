package browser

import (
	"fmt"
	"os/exec"
	"runtime"
)

func GetOpenBrowserCommand(url string) (*exec.Cmd, error) {
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
