package commonerr

import "fmt"

var (
	CastErr error
)

func NewCastError(v interface{}, t string) error {
	CastErr = fmt.Errorf(`could not cast %v, to type "%v"`, v, t)
	return CastErr
}
