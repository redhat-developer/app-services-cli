package consumergroup

import (
	"strconv"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"
	"github.com/redhat-developer/app-services-cli/pkg/localize"

	"github.com/relvacode/iso8601"
)

type Validator struct {
	Localizer localize.Localizer
}

// ValidateOffset checks if value v is a valid value for --offset
func (v *Validator) ValidateOffset(offset string) error {
	isValid := flagutil.IsValidInput(offset, ValidOffsets...)

	if isValid {
		return nil
	}

	return flag.InvalidValueError("output", v, ValidOffsets...)
}

// ValidateOffsetValue validates value for timestamp and absolute offset
// value for absolute offset should be integer and timestamp should be in ISO 8601 format
func (v *Validator) ValidateOffsetValue(offset string, value string) error {
	offsetValueTmplPair := localize.NewEntry("Value", value)
	switch offset {
	case OffsetTimestamp:
		if _, parseErr := iso8601.ParseString(value); parseErr != nil {
			return v.Localizer.MustLocalizeError("kafka.consumerGroup.resetOffset.error.invalidTimestampOffset", offsetValueTmplPair)
		}
	case OffsetAbsolute:
		if _, parseErr := strconv.Atoi(value); parseErr != nil {
			return v.Localizer.MustLocalizeError("kafka.consumerGroup.resetOffset.error.invalidAbsoluteOffset", offsetValueTmplPair)
		}
	}
	return nil
}
