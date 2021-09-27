// This file contains functions that help to manage visibility of early stage commands
package profile

import (
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

const DevPreviewEnv = "RHOAS_DEV"

// ApplyDevPreviewLabel adds visual element displayed in help
func ApplyDevPreviewLabel(cmd *cobra.Command) {
	cmd.Short = "[beta] " + cmd.Short
	cmd.Long += "\nThis command is available as part of the Development Preview release.\n"

	for _, child := range cmd.Commands() {
		ApplyDevPreviewLabel(child)
	}
}

// DevPreviewAnnotation used in templates and tools like documentation generation
func DevPreviewAnnotation() map[string]string {
	return map[string]string{"channel": "preview"}
}

// DevModeEnabled Check if development mode is enabled
func DevModeEnabled() bool {
	rawEnvVal := os.Getenv(DevPreviewEnv)

	boolVal, err := strconv.ParseBool(rawEnvVal)
	if err != nil {
		return false
	}

	return boolVal
}
