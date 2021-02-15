package connection

import (
	"errors"
	"fmt"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
)

func init() {
	localizer.LoadMessageFiles("connection")

	ErrNotLoggedIn = errors.New(localizer.MustLocalizeFromID("connection.error.notLoggedInError"))
	ErrSessionExpired = errors.New(localizer.MustLocalizeFromID("connection.error.sessionExpiredError"))
}

var (
	// ErrNotLoggedIn defines when a user is not authenticated
	ErrNotLoggedIn error
	// ErrSessionExpired defines when a user's session has expired
	ErrSessionExpired error
)

// AuthError defines an Authentication error
type AuthError struct {
	Err    error
	Reason string
}

func (e *AuthError) Error() string {
	var reason string
	if e.Reason != "" {
		reason = ": " + e.Reason
	}
	return fmt.Sprintf("%v%v", e.Err, reason)
}

func (e *AuthError) Unwrap() error {
	return e.Err
}

func AuthErrorf(format string, a ...interface{}) *AuthError {
	err := fmt.Errorf(format, a...)
	return &AuthError{err, ""}
}

func notLoggedInError() *AuthError {
	return &AuthError{ErrNotLoggedIn, ""}
}

func sessionExpiredError() *AuthError {
	return &AuthError{ErrSessionExpired, ""}
}
