package kas

import (
	"errors"
	"fmt"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/kas/client"
)

const (
	NotFoundErrCode = "MGD-SERV-API-7"
)

type Error struct {
	Err error
}

func (e *Error) Error() string {
	return fmt.Sprint(e.Err)
}

func (e *Error) Unwrap() error {
	return e.Err
}

func GetAPIError(err error) (e kasclient.Error, ok bool) {
	var apiError kasclient.GenericOpenAPIError
	if ok = errors.As(err, &apiError); ok {
		errModel := apiError.Model()

		e, ok = errModel.(kasclient.Error)
	}

	return e, ok
}

func IsNotFoundError(err error) bool {
	mappedErr, ok := GetAPIError(err)
	if !ok {
		return false
	}

	return mappedErr.GetCode() == NotFoundErrCode
}
