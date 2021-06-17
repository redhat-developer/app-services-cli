package srs

import (
	"errors"
	"fmt"

	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/registrymgmt/apiv1/client"
)

type ServiceErrorCode int

const (
	ErrCodePrefix = "KAFKAS-MGMT"

	// Forbidden occurs when a user is not allowed to access the service
	ErrorForbidden ServiceErrorCode = 4

	// Forbidden occurs when a user or organisation has reached maximum number of allowed instances
	ErrorMaxAllowedInstanceReached ServiceErrorCode = 5

	// Conflict occurs when a database constraint is violated
	ErrorConflict ServiceErrorCode = 6

	// NotFound occurs when a record is not found in the database
	ErrorNotFound ServiceErrorCode = 7

	// Validation occurs when an object fails validation
	ErrorValidation ServiceErrorCode = 8

	// General occurs when an error fails to match any other error code
	ErrorGeneral ServiceErrorCode = 9

	// NotImplemented occurs when an API REST method is not implemented in a handler
	ErrorNotImplemented ServiceErrorCode = 10

	// Unauthorized occurs when the requester is not authorized to perform the specified action
	ErrorUnauthorized ServiceErrorCode = 11

	// Unauthenticated occurs when the provided credentials cannot be validated
	ErrorUnauthenticated ServiceErrorCode = 15

	// MalformedRequest occurs when the request body cannot be read
	ErrorMalformedRequest ServiceErrorCode = 17

	// Bad Request
	ErrorBadRequest ServiceErrorCode = 21

	// Invalid Search Query
	ErrorFailedToParseSearch ServiceErrorCode = 23

	// Failed to create service account
	ErrorFailedToCreateServiceAccount ServiceErrorCode = 110

	// Failed to get service account
	ErrorFailedToGetServiceAccount ServiceErrorCode = 111

	// Failed to delete service account
	ErrorFailedToDeleteServiceAccount ServiceErrorCode = 112

	// Provider not supported
	ErrorProviderNotSupported ServiceErrorCode = 30

	// Region not supported
	ErrorRegionNotSupported ServiceErrorCode = 31

	// Invalid kafka cluster name
	ErrorMalformedKafkaClusterName ServiceErrorCode = 32

	// Minimum field length validation
	ErrorMinimumFieldLength ServiceErrorCode = 33

	// Maximum field length validation
	ErrorMaximumFieldLength ServiceErrorCode = 34

	// Only MultiAZ is supported
	ErrorOnlyMultiAZSupported ServiceErrorCode = 35

	// Kafka cluster name must be unique
	ErrorDuplicateKafkaClusterName ServiceErrorCode = 36

	// Failure to send an error response (i.e. unable to send error response as the error can't be converted to JSON.)
	ErrorUnableToSendErrorResponse ServiceErrorCode = 1000
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

// GetAPIError gets a strongly typed error from an error
func GetAPIError(err error) (e kafkamgmtclient.Error, ok bool) {
	var apiError kafkamgmtclient.GenericOpenAPIError

	if ok = errors.As(err, &apiError); ok {
		errModel := apiError.Model()

		e, ok = errModel.(kafkamgmtclient.Error)
	}

	return e, ok
}

// IsErr returns true if the error contains the errCode
func IsErr(err error, errCode ServiceErrorCode) bool {
	mappedErr, ok := GetAPIError(err)
	if !ok {
		return false
	}

	return mappedErr.GetCode() == getCode(errCode)
}

func getCode(code ServiceErrorCode) string {
	return fmt.Sprintf("%v-%v", ErrCodePrefix, code)
}
