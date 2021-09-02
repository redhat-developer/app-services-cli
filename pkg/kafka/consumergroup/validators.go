package consumergroup

import (
	"errors"
	"strconv"

	"github.com/redhat-developer/app-services-cli/pkg/localize"
)

type Validator struct {
	Localizer localize.Localizer
}

// ValidateOffsetValue validates value for timestamp and absolute offset
// value for absolute offset should be integer and timestamp should be in format "yyyy-MM-dd'T'HH:mm:ss"
func (v *Validator) ValidateOffsetValue(offset string, value string) error {
	offsetValueTmplPair := localize.NewEntry("Value", value)
	switch offset {
	case "timestamp":
		matched := timestampOffsetRegExp.MatchString(value)
		if !matched {
			return errors.New(v.Localizer.MustLocalize("kafka.consumerGroup.resetOffset.error.invalidTimestampOffset", offsetValueTmplPair))
		}
	case OffsetAbssolute:
		if _, parseErr := strconv.Atoi(value); parseErr != nil {
			offsetValueTmplPair := localize.NewEntry("Value", value)
			return errors.New(v.Localizer.MustLocalize("kafka.consumerGroup.resetOffset.error.invalidAbsoluteOffset", offsetValueTmplPair))
		}
	}
	return nil
}
