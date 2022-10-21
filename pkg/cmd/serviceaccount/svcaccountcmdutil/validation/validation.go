package validation

import (
	"regexp"

	"github.com/redhat-developer/app-services-cli/pkg/core/errors"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
)

const (
	// name validation rules
	legalShortDescriptionChars = "^[a-z]([-a-z0-9]*[a-z0-9])?$"
	maxNameLength              = 50
	minNameLength              = 1
	legalUUID                  = "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"
)

// Validator is a type for validating service account configuration values
type Validator struct {
	Localizer localize.Localizer
}

// ValidateShortDescription validates the short description value
func (v *Validator) ValidateShortDescription(val interface{}) error {
	name, ok := val.(string)
	if !ok {
		return errors.NewCastError(val, "string")
	}

	if len(name) < minNameLength {
		return v.Localizer.MustLocalizeError("serviceAccount.common.validation.shortDescription.error.required")
	} else if len(name) > maxNameLength {
		return v.Localizer.MustLocalizeError("serviceAccount.common.validation.shortDescription.error.lengthError", localize.NewEntry("MaxNameLen", maxNameLength))
	}

	matched, _ := regexp.Match(legalShortDescriptionChars, []byte(name))

	if matched {
		return nil
	}

	return v.Localizer.MustLocalizeError("serviceAccount.common.validation.shortDescription.error.invalidChars", localize.NewEntry("Name", name))
}

// ValidateUUID validates if ID is a valid UUID
func (v *Validator) ValidateUUID(val interface{}) error {
	id, ok := val.(string)
	if !ok {
		return errors.NewCastError(val, "string")
	}

	matched, _ := regexp.Match(legalUUID, []byte(id))

	if matched {
		return nil
	}

	return v.Localizer.MustLocalizeError("serviceAccount.common.validation.id.error.invalidID", localize.NewEntry("ID", id))
}

// ValidatePage validates if page number is valid
func (v *Validator) ValidatePage(val interface{}) error {
	page, ok := val.(int32)
	if !ok {
		return errors.NewCastError(val, "int32")
	}

	if page < 1 {
		return v.Localizer.MustLocalizeError("serviceAccount.list.validation.page.error", localize.NewEntry("Page", page))
	}

	return nil

}

// ValidateSize validates if a valid size is provided
func (v *Validator) ValidateSize(val interface{}) error {
	size, ok := val.(int32)
	if !ok {
		return errors.NewCastError(val, "int32")
	}

	if size > 100 {
		return v.Localizer.MustLocalizeError("serviceAccount.list.validation.size.error.max", localize.NewEntry("Size", size))
	}

	if size < 1 {
		return v.Localizer.MustLocalizeError("serviceAccount.list.validation.size.error.min", localize.NewEntry("Size", size))
	}

	return nil

}
