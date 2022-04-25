package browser

import (
	"fmt"
	"os/exec"
	"runtime"
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
		return fmt.Errorf("unsupported operating system: %v", runtime.GOOS)
	}
}
