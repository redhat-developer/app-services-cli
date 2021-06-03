package kafka

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/redhat-developer/app-services-cli/pkg/common/commonerr"
	"github.com/redhat-developer/app-services-cli/pkg/kafka/kafkaerr"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	kafkamgmtv1 "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1"
)

var (
	validNameRegexp   = regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`)
	validSearchRegexp = regexp.MustCompile(`^([a-zA-Z0-9-_%]*[a-zA-Z0-9-_%])?$`)
)

// ValidateName validates the proposed name of a Kafka instance
func ValidateName(val interface{}) error {
	name, ok := val.(string)

	if !ok {
		return commonerr.NewCastError(val, "string")
	}

	if len(name) < 1 || len(name) > 32 {
		return errors.New("Kafka instance name must be between 1 and 32 characters")
	}

	matched := validNameRegexp.MatchString(name)

	if matched {
		return nil
	}

	return kafkaerr.InvalidNameError(name)
}

// TransformKafkaRequestListItems modifies fields fields from a list of kafka instances
// The main transformation is appending ":443" to the Bootstrap Server URL
func TransformKafkaRequestListItems(items []kafkamgmtv1.KafkaRequest) []kafkamgmtv1.KafkaRequest {
	for i := range items {
		kafka := items[i]
		kafka = *TransformKafkaRequest(&kafka)
		items[i] = kafka
	}

	return items
}

// TransformKafkaRequest modifies fields from the KafkaRequest payload object
// The main transformation is appending ":443" to the Bootstrap Server URL
func TransformKafkaRequest(kafka *kafkamgmtv1.KafkaRequest) *kafkamgmtv1.KafkaRequest {
	bootstrapHost := kafka.GetBootstrapServerHost()

	if bootstrapHost == "" {
		return kafka
	}

	if !strings.HasSuffix(bootstrapHost, ":443") {
		hostURL := fmt.Sprintf("%v:443", bootstrapHost)
		kafka.SetBootstrapServerHost(hostURL)
	}

	return kafka
}

// ValidateSearchInput validates the text provided to filter the Kafka instances
func ValidateSearchInput(val interface{}) error {
	search, ok := val.(string)

	if !ok {
		return commonerr.NewCastError(val, "string")
	}

	matched := validSearchRegexp.MatchString(search)

	if matched {
		return nil
	}

	return kafkaerr.InvalidSearchValueError(search)
}

// ValidateNameIsAvailable checks if a kafka instance with the given name already exists
func ValidateNameIsAvailable(api kafkamgmtv1.DefaultApi, localizer localize.Localizer) func(v interface{}) error {
	return func(v interface{}) error {
		name, _ := v.(string)

		_, httpRes, _ := GetKafkaByName(context.Background(), api, name)

		if httpRes != nil && httpRes.StatusCode == 200 {
			return errors.New(localizer.MustLocalize("kafka.create.error.conflictError", localize.NewEntry("Name", name)))
		}

		if httpRes != nil && httpRes.Body != nil {
			httpRes.Body.Close()
		}

		return nil
	}
}
