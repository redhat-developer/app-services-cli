package error

import (
	"errors"

	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/kafkamgmt/apiv1/client"
)

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

// IsAPIError returns true if the error contains the errCode
func IsAPIError(err error, code string) bool {
	mappedErr := GetAPIError(err)
	if mappedErr == nil {
		return false
	}

	return mappedErr.GetCode() == string(code)
}
