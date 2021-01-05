package token

import (
	"errors"
	"fmt"
)

var (
	// ErrTokenExpired defines when a user's token is expired
	ErrTokenExpired = errors.New("token expired")
)

// Error records an error with a JWT token
type Error struct {
	Err    error
	Reason string
}

func (e *Error) Error() string {
	var reason string
	if e.Reason != "" {
		reason = ": " + e.Reason
	}
	return fmt.Sprintf("TokenError: %v%v", e.Err, reason)
}

func (e *Error) Unwrap() error {
	return e.Err
}

func Errorf(format string, a ...interface{}) *Error {
	err := fmt.Errorf(format, a...)
	return &Error{err, ""}
}
