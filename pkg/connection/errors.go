package connection

import (
	"errors"
	"fmt"
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
	return errors.New(`not logged in. Run "rhoas login" to authenticate`)
}

func notLoggedInMASError() error {
	return errors.New(`not logged in to identity.api.openshift.com. Run "rhoas login" to authenticate. Note: token-based login is not supported by the Kafka "topic" and "consumergroup" subcommands`)
}

func sessionExpiredError() error {
	return errors.New(`session expired. Run "rhoas login" to authenticate`)
}
