package validation

import (
	"errors"
	"regexp"

	"github.com/redhat-developer/app-services-cli/pkg/common/commonerr"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
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

// Validator is interface for validation object
type Validator struct {
	Localizer localize.Localizer
}

func (v *Validator) ValidateName(val interface{}) error {
	name, ok := val.(string)
	if !ok {
		return commonerr.NewCastError(val, "string")
	}

	if len(name) < minNameLength {
		return errors.New(v.Localizer.MustLocalize("serviceAccount.common.validation.name.error.required"))
	} else if len(name) > maxNameLength {
		return errors.New(v.Localizer.MustLocalize("serviceAccount.common.validation.name.error.lengthError", localize.NewEntry("MaxNameLen", maxNameLength)))
	}

	matched, _ := regexp.Match(legalNameChars, []byte(name))

	if matched {
		return nil
	}

	return errors.New(v.Localizer.MustLocalize("serviceAccount.common.validation.name.error.invalidChars", localize.NewEntry("Name", name)))
}

// ValidateDescription validates the service account description text
func (v *Validator) ValidateDescription(val interface{}) error {
	description, ok := val.(string)
	if !ok {
		return commonerr.NewCastError(val, "string")
	}

	if description == "" {
		return nil
	}

	if len(description) > maxDescriptionLength {
		return errors.New(v.Localizer.MustLocalize("serviceAccount.common.validation.description.error.lengthError", localize.NewEntry("MaxLen", maxDescriptionLength)))
	}

	matched, _ := regexp.Match(legalDescriptionChars, []byte(description))

	if matched {
		return nil
	}

	return errors.New(v.Localizer.MustLocalize("serviceAccount.common.validation.description.error.invalidChars"))
}
