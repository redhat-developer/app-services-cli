package validation

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/redhat-developer/app-services-cli/pkg/common/commonerr"
)

const (
	// name validation rules
	legalNameChars = "^[a-z]([-a-z0-9]*[a-z0-9])?$"
	maxNameLength  = 50
	minNameLength  = 1
	// description validation rules
	legalDescriptionChars = "^[a-zA-Z0-9.,\\-\\s]*$"
	maxDescriptionLength  = 255
)

// ValidateName validates the name of the service account
func ValidateName(val interface{}) error {
	name, ok := val.(string)
	if !ok {
		return commonerr.NewCastError(val, "string")
	}

	if len(name) < minNameLength {
		return errors.New("service account name is required")
	} else if len(name) > maxNameLength {
		return fmt.Errorf("service account name cannot exceed %v characters", maxNameLength)
	}

	matched, _ := regexp.Match(legalNameChars, []byte(name))

	if matched {
		return nil
	}

	return fmt.Errorf(`invalid service account name "%v"; only lowercase letters (a-z), numbers, and "-" are accepted`, name)
}

// ValidateDescription validates the service account description text
func ValidateDescription(val interface{}) error {
	description, ok := val.(string)
	if !ok {
		return commonerr.NewCastError(val, "string")
	}

	if description == "" {
		return nil
	}

	if len(description) > maxDescriptionLength {
		return fmt.Errorf("service account description cannot exceed %v characters", maxDescriptionLength)
	}

	matched, _ := regexp.Match(legalDescriptionChars, []byte(description))

	if matched {
		return nil
	}

	return errors.New(`invalid service account description; only alphanumeric characters and "-", ".", "," are accepted.`)
}
