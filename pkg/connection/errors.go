package connection

import (
	"fmt"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/color"
)

var (
	loginCmd = color.CodeSnippet("rhoas login")
	// ErrNotLoggedIn defines when a user is not authenticated
	ErrNotLoggedIn = fmt.Errorf("Not logged in. Run %v to authenticate", loginCmd)
	// ErrSessionExpired defines when a user's session has expired
	ErrSessionExpired = fmt.Errorf("Session expired. Run %v to authenticate", loginCmd)
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
