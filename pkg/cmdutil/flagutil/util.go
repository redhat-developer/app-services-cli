package flagutil

import (
	"fmt"
	"sort"

	"github.com/redhat-developer/app-services-cli/pkg/localize"
)

// IsValidInput checks if the input value is in the range of valid values
func IsValidInput(input string, validValues ...string) bool {
	for _, b := range validValues {
		if input == b {
			return true
		}
	}

	return false
}

// FlagDescription creates a flag description and adds a list of valid options (if any)
func FlagDescription(localizer localize.Localizer, messageID string, validOptions ...string) string {
	// ensure consistent order
	sort.Strings(validOptions)

	description := localizer.MustLocalize(messageID)
	if description[len(description)-1:] != "." {
		description += "."
	}

	var chooseFrom string
	if len(validOptions) > 0 {
		chooseFrom = localizer.MustLocalize("flag.common.chooseFrom")

		for i, val := range validOptions {
			chooseFrom += fmt.Sprintf("\"%v\"", val)
			if i < len(validOptions)-1 {
				chooseFrom += ", "
			}
		}
	}

	return fmt.Sprintf("%v %v", description, chooseFrom)
}
