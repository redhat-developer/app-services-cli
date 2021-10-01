package flag

import (
	"fmt"
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
		chooseFromStr = ", valid options are: "
		for i, option := range validOptions {
			chooseFromStr += fmt.Sprintf(`"%v"`, option)
			if (i + 1) < len(validOptions) {
				chooseFromStr += ", "
			}
		}
	}
	return &Error{Err: fmt.Errorf(`invalid value "%v" for --%v%v`, val, flag, chooseFromStr)}
}

func RequiredWhenNonInteractiveError(flags ...string) error {
	var flagsF string
	for i := 0; i < len(flags); i++ {
		delimiter := ","
		switch i {
		case len(flags) - 1:
			delimiter = ""
		case len(flags) - 2:
			delimiter = " and "
		}
		flagsF += fmt.Sprintf("--%v%v", flags[i], delimiter)
	}
	flagTitle := "flag"
	if len(flags) > 1 {
		flagTitle = "flags"
	}
	return fmt.Errorf("%v %v required when not running interactively", flagsF, flagTitle)
}
