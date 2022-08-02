package connectorcmdutil

import (
	"regexp"

	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
)

const (
	legalNamespaceChars = "^(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])?$"
)

// Validator is a type for validating connector attributes
type Validator struct {
	Localizer localize.Localizer
}

// ValidateNamespace validates the name of the namespace
func (v *Validator) ValidateNamespace(name string) error {
	// TODO check for empty string
	matched, _ := regexp.Match(legalNamespaceChars, []byte(name))

	if matched {
		return nil
	}

	return v.Localizer.MustLocalizeError("connector.common.validation.namespace.error.invalidChars", localize.NewEntry("Namespace", name))
}
