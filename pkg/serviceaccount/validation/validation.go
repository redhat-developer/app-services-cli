package validation

import (
	"errors"
	"regexp"

	"github.com/redhat-developer/app-services-cli/internal/localizer"
)

const (
	// name validation rules
	legalNameChars = "^[a-z]([-a-z0-9]*[a-z0-9])?$"
	maxNameLength  = 50
	minNameLength  = 1
	// description validation rules
	legalDescriptionChars = "^[a-z0-9\\s]*$"
	maxDescriptionLength  = 255
)

// ValidateName validates the name of the service account
func ValidateName(val interface{}) error {
	name, ok := val.(string)
	if !ok {
		return errors.New(localizer.MustLocalize(&localizer.Config{
			MessageID: "common.error.castError",
			TemplateData: map[string]interface{}{
				"Value": val,
				"Type":  "string",
			},
		}))
	}

	if len(name) < minNameLength {
		return errors.New(localizer.MustLocalizeFromID("serviceAccount.common.validation.name.error.required"))
	} else if len(name) > maxNameLength {
		return errors.New(localizer.MustLocalize(&localizer.Config{
			MessageID: "serviceAccount.common.validation.name.error.lengthError",
			TemplateData: map[string]interface{}{
				"MaxNameLen": maxNameLength,
			},
		}))
	}

	matched, _ := regexp.Match(legalNameChars, []byte(name))

	if matched {
		return nil
	}

	return errors.New(localizer.MustLocalize(&localizer.Config{
		MessageID: "serviceAccount.common.validation.name.error.invalidChars",
		TemplateData: map[string]interface{}{
			"Name": name,
		},
	}))
}

// ValidateDescription validates the service account description text
func ValidateDescription(val interface{}) error {
	description, ok := val.(string)
	if !ok {
		return errors.New(localizer.MustLocalize(&localizer.Config{
			MessageID: "common.error.castError",
			TemplateData: map[string]interface{}{
				"Value": val,
				"Type":  "string",
			},
		}))
	}

	if description == "" {
		return nil
	}

	if len(description) > maxDescriptionLength {
		return errors.New(localizer.MustLocalize(&localizer.Config{
			MessageID: "serviceAccount.common.validation.description.error.lengthError",
			TemplateData: map[string]interface{}{
				"MaxNameLen": maxDescriptionLength,
			},
		}))
	}

	matched, _ := regexp.Match(legalDescriptionChars, []byte(description))

	if matched {
		return nil
	}

	return errors.New(localizer.MustLocalize(&localizer.Config{
		MessageID: "serviceAccount.common.validation.description.error.invalidChars",
		TemplateData: map[string]interface{}{
			"Description": description,
		},
	}))
}
