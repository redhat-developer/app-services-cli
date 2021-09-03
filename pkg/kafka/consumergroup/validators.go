package consumergroup

import (
	"errors"
	"regexp"
	"strconv"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
)

var timestampOffsetRegExp = regexp.MustCompile(`^(\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}-\d{2}:\d{2})$`)

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
// value for absolute offset should be integer and timestamp should be in format "yyyy-MM-dd'T'HH:mm:ssz"
func (v *Validator) ValidateOffsetValue(offset string, value string) error {
	offsetValueTmplPair := localize.NewEntry("Value", value)
	switch offset {
	case OffsetTimestamp:
		matched := timestampOffsetRegExp.MatchString(value)
		if !matched {
			return errors.New(v.Localizer.MustLocalize("kafka.consumerGroup.resetOffset.error.invalidTimestampOffset", offsetValueTmplPair))
		}
	case OffsetAbsolute:
		if _, parseErr := strconv.Atoi(value); parseErr != nil {
			offsetValueTmplPair := localize.NewEntry("Value", value)
			return errors.New(v.Localizer.MustLocalize("kafka.consumerGroup.resetOffset.error.invalidAbsoluteOffset", offsetValueTmplPair))
		}
	}
	return nil
}
