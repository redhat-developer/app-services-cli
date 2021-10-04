package kas

import (
	"errors"
	"fmt"

	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
)

const (
	// ErrorCode7 Resource not found
	ErrorCode7 = "KAFKAS-MGMT-7"
	// ErrorCode24 The maximum number of allowed kafka instances has been reached
	ErrorCode24 = "KAFKAS-MGMT-24"
	// ErrorCode21 Bad Request
	ErrorCode21 = "KAFKAS-MGMT-21"
	// ErrorCode36 Kafka cluster name is already used
	ErrorCode36 = "KAFKAS-MGMT-36"
)

type ServiceErrorCode string

type Error struct {
	Err error
}

func (e *Error) Error() string {
	return fmt.Sprint(e.Err)
}

func (e *Error) Unwrap() error {
	return e.Err
}

// GetAPIError gets a strongly typed error from an error
func GetAPIError(err error) *kafkamgmtclient.Error {
	var openapiError kafkamgmtclient.GenericOpenAPIError

	if ok := errors.As(err, &openapiError); ok {
		errModel := openapiError.Model()

		kafkaMgmtError, ok := errModel.(kafkamgmtclient.Error)
		if !ok {
			return nil
		}
		return &kafkaMgmtError
	}

	return nil
}

// IsErr returns true if the error contains the errCode
func IsErr(err error, code ServiceErrorCode) bool {
	mappedErr := GetAPIError(err)
	if mappedErr == nil {
		return false
	}

	return mappedErr.GetCode() == string(code)
}
