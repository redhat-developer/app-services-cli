package flag

import (
	"fmt"

	"github.com/redhat-developer/app-services-cli/internal/localizer"
)

type Error struct {
	Err error
}

func (e *Error) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

func (e *Error) Unwrap() error {
	return e.Err
}

// InvalidValueError returns an error when an invalid flag value is provided
func InvalidValueError(flag string, val interface{}, validOptions ...string) *Error {
	var chooseFromStr string
	if len(validOptions) > 0 {
		chooseFromStr = localizer.MustLocalizeFromID("flag.error.invalidValue.options")
		for i, option := range validOptions {
			chooseFromStr += fmt.Sprintf(`"%v"`, option)
			if (i + 1) < len(validOptions) {
				chooseFromStr += ", "
			}
		}
	}
	return &Error{Err: fmt.Errorf("%v%v", localizer.MustLocalize(&localizer.Config{
		MessageID: "flag.error.invalidValue.base",
		TemplateData: map[string]interface{}{
			"Flag":  flag,
			"Value": val,
		},
	}), chooseFromStr)}
}
