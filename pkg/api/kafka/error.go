package kafka

import (
	"encoding/json"
	"errors"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
)

// GetAPIError gets a strongly typed error from an error
func GetAPIError(err error) (e kafkamgmtclient.Error, ok bool) {
	var apiError kafkamgmtclient.GenericOpenAPIError
	var kafkaError kafkamgmtclient.Error

	if ok = errors.As(err, &apiError); ok {
		kafkaError = kafkamgmtclient.Error{}
		err = json.Unmarshal(apiError.Body(), &kafkaError)
		if err != nil {
			return kafkaError, false
		}
	}

	return kafkaError, ok
}

// TransformError code contains message that can be returned to the user
func TransformError(err error) error {
	mappedErr, ok := GetAPIError(err)
	if !ok {
		return err
	}
	return errors.New(mappedErr.GetCode() + ": " + mappedErr.GetReason())
}
