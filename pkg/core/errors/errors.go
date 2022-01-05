package errors

import "fmt"

var CastErr error

// NewCastError returns a new error for when a type could not be cast
func NewCastError(v interface{}, t string) error {
	CastErr = fmt.Errorf(`could not cast %v, to type "%v"`, v, t)
	return CastErr
}
