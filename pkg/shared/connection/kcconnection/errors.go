package kcconnection

import (
	"errors"
	"fmt"
)

// AuthError defines an Authentication error
type AuthError struct {
	Err error
}

func (e *AuthError) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

func (e *AuthError) Unwrap() error {
	return e.Err
}

func AuthErrorf(format string, a ...interface{}) *AuthError {
	err := fmt.Errorf(format, a...)
	return &AuthError{err}
}

func notLoggedInError() error {
	return errors.New(`not logged in. Run "rhoas login" to authenticate`)
}

func sessionExpiredError() error {
	return errors.New(`session expired. Run "rhoas login" to authenticate`)
}
