package connection

import (
	"errors"
	"fmt"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
)

// AuthError defines an Authentication error
type AuthError struct {
	Err error
}

type MasAuthError struct {
	Err error
}

func (e *AuthError) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

func (e *MasAuthError) Unwrap() error {
	return e.Err
}

func (e *MasAuthError) Error() string {
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
	return errors.New(localizer.MustLocalizeFromID("connection.error.notLoggedInError"))
}

func notLoggedInMASError() error {
	return errors.New(localizer.MustLocalizeFromID("connection.error.notLoggedInMASError"))
}

func sessionExpiredError() error {
	return errors.New(localizer.MustLocalizeFromID("connection.error.sessionExpiredError"))
}
