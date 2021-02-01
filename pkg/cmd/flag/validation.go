package flag

import (
	"fmt"

	flagutil "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil/flags"
)

// ValidOutput checks if value v is a valid value for --output
func ValidateOutput(v string) error {
	isValid := flagutil.IsValidInput(v, flagutil.ValidOutputFormats...)

	if isValid {
		return nil
	}

	return fmt.Errorf("invalid value '%v' for --output, valid options are: %q", v, flagutil.ValidOutputFormats)
}
