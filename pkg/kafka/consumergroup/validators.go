package consumergroup

import (
	"errors"
	"regexp"

	"github.com/redhat-developer/app-services-cli/pkg/common/commonerr"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
)

const (
	legalSearchChars = "^[a-zA-Z0-9-]+$"
)

// Validator is a type for validating Kafka consumer group configuration values
type Validator struct {
	Localizer localize.Localizer
}

func (v *Validator) ValidateSearchInput(val interface{}) error {
	search, ok := val.(string)
	if !ok {
		return commonerr.NewCastError(val, "string")
	}

	matched, _ := regexp.Match(legalSearchChars, []byte(search))

	if matched {
		return nil
	}

	return errors.New(v.Localizer.MustLocalize("kafka.consumerGroup.list.error.illegalSearchValue", localize.NewEntry("Search", search)))
}
