package kafka

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	managedservices "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/managedservices/client"

	"github.com/MakeNowJust/heredoc"
)

var (
	validNameRegexp = regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`)
	errInvalidName  = errors.New(heredoc.Doc(`
	Invalid Kafka instance name. Valid names must satisfy the following conditions:

	- must be between 1 and 32 characters
	- must only consist of lower case, alphanumeric characters and '-'
	- must start with an alphabetic character
	- must end with an alphanumeric character
	`))
)

// ValidateName validates the proposed name of a Kafka instance
func ValidateName(val interface{}) error {
	name, ok := val.(string)
	if !ok {
		return fmt.Errorf("could not case %v to string", val)
	}

	if len(name) < 1 || len(name) > 32 {
		return fmt.Errorf("Kafka instance name must be between 1 and 32 characters")
	}

	matched := validNameRegexp.MatchString(name)

	if matched {
		return nil
	}

	return errInvalidName
}

// TransformKafkaRequestListItems modifies fields fields from a list of kafka instances
// The main transformation is appending ":443" to the Bootstrap Server URL
func TransformKafkaRequestListItems(items []managedservices.KafkaRequest) []managedservices.KafkaRequest {
	for i := range items {
		kafka := items[i]
		kafka = *TransformKafkaRequest(&kafka)
		items[i] = kafka
	}

	return items
}

// TransformKafkaRequest modifies fields from the KafkaRequest payload object
// The main transformation is appending ":443" to the Bootstrap Server URL
func TransformKafkaRequest(kafka *managedservices.KafkaRequest) *managedservices.KafkaRequest {
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
