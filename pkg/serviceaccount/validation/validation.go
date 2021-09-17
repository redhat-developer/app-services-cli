package validation

import (
	"regexp"

	"github.com/redhat-developer/app-services-cli/pkg/common/commonerr"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
)

const (
	// name validation rules
	legalNameChars = "^[a-z]([-a-z0-9]*[a-z0-9])?$"
	maxNameLength  = 50
	minNameLength  = 1
	legalUUID      = "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"
	// description validation rules
	legalDescriptionChars = "^[a-zA-Z0-9.,\\-\\s]*$"
	maxDescriptionLength  = 255
)

// Validator is a type for validating service account configuration values
type Validator struct {
	Localizer localize.Localizer
}

func (v *Validator) ValidateName(val interface{}) error {
	name, ok := val.(string)
	if !ok {
		return commonerr.NewCastError(val, "string")
	}

	if len(name) < minNameLength {
		return v.Localizer.MustLocalizeError("serviceAccount.common.validation.name.error.required")
	} else if len(name) > maxNameLength {
		return v.Localizer.MustLocalizeError("serviceAccount.common.validation.name.error.lengthError", localize.NewEntry("MaxNameLen", maxNameLength))
	}

	matched, _ := regexp.Match(legalNameChars, []byte(name))

	if matched {
		return nil
	}

	return v.Localizer.MustLocalizeError("serviceAccount.common.validation.name.error.invalidChars", localize.NewEntry("Name", name))
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
		return v.Localizer.MustLocalizeError("serviceAccount.common.validation.description.error.lengthError", localize.NewEntry("MaxLen", maxDescriptionLength))
	}

	matched, _ := regexp.Match(legalDescriptionChars, []byte(description))

	if matched {
		return nil
	}

	return v.Localizer.MustLocalizeError("serviceAccount.common.validation.description.error.invalidChars")
}

// ValidateUUID validates if ID is a valid UUID
func (v *Validator) ValidateUUID(val interface{}) error {
	id, ok := val.(string)
	if !ok {
		return commonerr.NewCastError(val, "string")
	}

	matched, _ := regexp.Match(legalUUID, []byte(id))

	if matched {
		return nil
	}

	return v.Localizer.MustLocalizeError("serviceAccount.common.validation.id.error.invalidID", localize.NewEntry("ID", id))
}
