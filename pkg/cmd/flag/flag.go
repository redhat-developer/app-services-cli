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

func InvalidValueError(flag string, val interface{}, validOptions ...string) *Error {
	var chooseFromStr string
	if len(validOptions) > 0 {
		chooseFromStr = ", valid options are "
		for i, option := range validOptions {
			chooseFromStr += fmt.Sprintf(`"%v"`, option)
			if (i + 1) < len(validOptions) {
				chooseFromStr += ", "
			}
		}
	}
	return &Error{Err: fmt.Errorf(`invalid value "%v" for --%v%v`, val, flag, chooseFromStr)}
}
