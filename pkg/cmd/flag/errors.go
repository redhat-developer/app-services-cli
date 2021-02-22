package flag

import (
	"fmt"
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
)

func InvalidArgumentError(flag string, value string, err error) error {
	return fmt.Errorf("%v: %w", localizer.MustLocalize(&localizer.Config{
		MessageID: "flag.error.invalidArgumentError",
		TemplateData: map[string]interface{}{
			"Argument": flag,
			"Value":    value,
		},
	}), err)
}
