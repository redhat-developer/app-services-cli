package registryutil

import (
	"errors"
	"regexp"

	coreErrors "github.com/redhat-developer/app-services-cli/pkg/core/errors"
)

var validNameRegexp = regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`)

// ValidateName validates the proposed name of a Kafka instance
func ValidateName(val interface{}) error {
	name, ok := val.(string)

	if !ok {
		return coreErrors.NewCastError(val, "string")
	}

	if len(name) < 1 || len(name) > 32 {
		return errors.New("service registry instance name must be between 1 and 32 characters")
	}

	matched := validNameRegexp.MatchString(name)

	if matched {
		return nil
	}

	return errors.New("invalid service registry name: " + name)
}
