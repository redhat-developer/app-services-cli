// This file contains functions that help to manage visibility of early stage commands
package profile

import (
	"os"
)

var DevPreviewEnv string = "RHOAS_DEV"

// Visual element displayed in help
func ApplyDevPreviewLabel(message string) string {
	return "[Preview] " + message
}

// Annotation used in templates and tools like documentation generation
func DevPreviewAnnotation() map[string]string {
	return map[string]string{"channel": "preview"}
}

// Check if preview is enabled
func DevPreviewEnabled() bool {
	return os.Getenv(DevPreviewEnv) != ""
}
