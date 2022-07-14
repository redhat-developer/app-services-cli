package kafkacmdutil

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/redhat-developer/app-services-cli/pkg/core/errors"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/redhat-developer/app-services-cli/pkg/shared/kafkautil"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
)

var (
	validNameRegexp   = regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`)
	validSearchRegexp = regexp.MustCompile(`^([a-zA-Z0-9-_%]*[a-zA-Z0-9-_%])?$`)
)

// Validator is a type for validating Kafka configuration values
type Validator struct {
	Localizer  localize.Localizer
	Connection factory.ConnectionFunc
}

// ValidateName validates the proposed name of a Kafka instance
func (v *Validator) ValidateName(val interface{}) error {
	name, ok := val.(string)

	if !ok {
		return errors.NewCastError(val, "string")
	}

	if len(name) < 1 || len(name) > 32 {
		return v.Localizer.MustLocalizeError("kafka.validation.name.error.lengthError", localize.NewEntry("MaxLength", 32))
	}

	matched := validNameRegexp.MatchString(name)

	if matched {
		return nil
	}

	return kafkautil.InvalidNameError(name)
}

// TransformKafkaRequestListItems modifies fields fields from a list of kafka instances
// The main transformation is appending ":443" to the Bootstrap Server URL
func TransformKafkaRequestListItems(items []kafkamgmtclient.KafkaRequest) []kafkamgmtclient.KafkaRequest {
	for i := range items {
		kafka := items[i]
		kafka = *TransformKafkaRequest(&kafka)
		items[i] = kafka
	}

	return items
}

// TransformKafkaRequest modifies fields from the KafkaRequest payload object
// The main transformation is appending ":443" to the Bootstrap Server URL
func TransformKafkaRequest(kafka *kafkamgmtclient.KafkaRequest) *kafkamgmtclient.KafkaRequest {
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
func (v *Validator) ValidateSearchInput(val interface{}) error {
	search, ok := val.(string)

	if !ok {
		return errors.NewCastError(val, "string")
	}

	matched := validSearchRegexp.MatchString(search)

	if matched {
		return nil
	}

	return kafkautil.InvalidSearchValueError(search)
}

// ValidateNameIsAvailable checks if a kafka instance with the given name already exists
func (v *Validator) ValidateNameIsAvailable(val interface{}) error {
	name, _ := val.(string)

	conn, err := v.Connection()
	if err != nil {
		return err
	}

	api := conn.API()

	_, httpRes, _ := kafkautil.GetKafkaByName(context.Background(), api.KafkaMgmt(), name)

	if httpRes != nil {
		defer httpRes.Body.Close()
		if httpRes.StatusCode == http.StatusOK {
			return v.Localizer.MustLocalizeError("kafka.create.error.conflictError", localize.NewEntry("Name", name))
		}
	}

	return nil
}
