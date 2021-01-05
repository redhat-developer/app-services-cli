package connection

import (
	"errors"
	"fmt"
)

var (
	// ErrNotLoggedIn defines when a user is not authenticated
	ErrNotLoggedIn = errors.New("Not logged in. Run `rhoas login` to authenticate")
	// ErrSessionExpired defines when a user's session has expired
	ErrSessionExpired = errors.New("Session expired. Run `rhoas login` to authenticate")
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
	return fmt.Sprintf("AuthError: %v%v", e.Err, reason)
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
