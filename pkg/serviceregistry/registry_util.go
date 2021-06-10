package serviceregistry

import (
	"errors"
	"regexp"

	"github.com/redhat-developer/app-services-cli/pkg/common/commonerr"
	"github.com/redhat-developer/app-services-cli/pkg/serviceregistry/registryerr"
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
		return errors.New("ServiceRegistry instance name must be between 1 and 32 characters")
	}

	matched := validNameRegexp.MatchString(name)

	if matched {
		return nil
	}

	return registryerr.InvalidNameError(name)
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

	return registryerr.InvalidSearchValueError(search)
}
