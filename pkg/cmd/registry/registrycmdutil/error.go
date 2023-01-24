package registrycmdutil

import (
	"errors"
	"fmt"

	registryinstanceclient "github.com/jackdelahunt/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal/client"
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

// GetInstanceAPIError gets a strongly typed error from an error
func GetInstanceAPIError(err error) (e registryinstanceclient.Error, ok bool) {
	var apiError registryinstanceclient.GenericOpenAPIError

	if ok = errors.As(err, &apiError); ok {
		errModel := apiError.Model()

		e, ok = errModel.(registryinstanceclient.Error)
	}

	return e, ok
}

// TransformInstanceError code contains message that can be returned to the user
func TransformInstanceError(err error) error {
	mappedErr, ok := GetInstanceAPIError(err)
	if !ok {
		return err
	}

	return errors.New(mappedErr.GetName() + ": " + mappedErr.GetMessage())
}
