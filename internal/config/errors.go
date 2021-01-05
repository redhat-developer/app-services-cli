package config

import (
	"fmt"
)

// Error records a Config file error
type Error struct {
	Err    error
	Reason string
}

func (e *Error) Error() string {
	var reason string
	if e.Reason != "" {
		reason = ": " + e.Reason
	}
	return fmt.Sprintf("ConfigError: %v%v", e.Err.Error(), reason)
}

func (e *Error) Unwrap() error {
	return e.Err
}

func Errorf(format string, a ...interface{}) *Error {
	err := fmt.Errorf(format, a...)
	return &Error{err, ""}
}
