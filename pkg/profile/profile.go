// This file contains functions that help to manage visibility of early stage commands
package profile

import (
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var DevPreviewEnv string = "RHOAS_DEV"

// ApplyDevPreviewLabel adds visual element displayed in help
func ApplyDevPreviewLabel(cmd *cobra.Command) {
	cmd.Short = "[alpha] " + cmd.Short
	cmd.Long += "\nThis command is available as part of the developer preview\n"

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
