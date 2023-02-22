package error

import (
	"errors"

	svcacctmgmtclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/serviceaccountmgmt/apiv1/client"
)

// GetAPIError gets a strongly typed error from an error
func GetAPIError(err error) *svcacctmgmtclient.RedHatErrorRepresentation {
	var openapiError svcacctmgmtclient.GenericOpenAPIError

	if ok := errors.As(err, &openapiError); ok {
		errModel := openapiError.Model()

		svcAcctError, ok := errModel.(svcacctmgmtclient.RedHatErrorRepresentation)
		if !ok {
			return nil
		}
		return &svcAcctError
	}

	return nil
}

// IsAPIError returns true if the error contains the errCode
func IsAPIError(err error, code string) bool {
	mappedErr := GetAPIError(err)
	if mappedErr == nil {
		return false
	}

	return mappedErr.GetError() == string(code)
}
