package serviceregistry

import (
	"errors"
	"regexp"

	"github.com/redhat-developer/app-services-cli/pkg/common/commonerr"
)

var validNameRegexp = regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`)

// ValidateName validates the proposed name of a Kafka instance
func ValidateName(val interface{}) error {
	name, ok := val.(string)

	if !ok {
		return commonerr.NewCastError(val, "string")
	}

	if len(name) < 1 || len(name) > 32 {
		return errors.New("ServiceRegistry instance name must be between 1 and 32 characters")
	}

	matched := validNameRegexp.MatchString(name)

	if matched {
		return nil
	}

	return errors.New("Invalid service registry name: " + name)
}
