package flag

import (
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
	flagutil "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil/flags"
)

func init() {
	localizer.LoadMessageFiles("cmd/common/flags")
}

// ValidOutput checks if value v is a valid value for --output
func ValidateOutput(v string) error {
	isValid := flagutil.IsValidInput(v, flagutil.ValidOutputFormats...)

	if isValid {
		return nil
	}

	return InvalidValueError("output", v, flagutil.ValidOutputFormats...)
}
