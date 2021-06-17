package topic

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"strconv"

	"github.com/redhat-developer/app-services-cli/pkg/common/commonerr"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal/client"
)

const (
	legalNameChars = "^[a-zA-Z0-9._-]+$"
	maxNameLength  = 249
	minPartitions  = 1
	maxPartitions  = 100
)

type Validator struct {
	Localizer localize.Localizer
}

// ValidateName validates the name of the topic
func (v Validator) ValidateName(val interface{}) error {
	name, ok := val.(string)
	if !ok {
		return commonerr.NewCastError(val, "string")
	}

	if len(name) < 1 {
		return errors.New(v.Localizer.MustLocalize("kafka.topic.common.validation.name.error.required"))
	} else if len(name) > maxNameLength {
		return errors.New(v.Localizer.MustLocalize("kafka.topic.common.validation.name.error.lengthError", localize.NewEntry("MaxNameLen", maxNameLength)))
	}

	if (name == ".") || (name == "..") {
		return errors.New(v.Localizer.MustLocalize("kafka.topic.common.validation.name.error.dotsError"))
	}

	matched, _ := regexp.Match(legalNameChars, []byte(name))

	if matched {
		return nil
	}

	return errors.New(v.Localizer.MustLocalize("kafka.topic.common.validation.name.error.invalidChars", localize.NewEntry("Name", name)))
}

func ValidateSearchInput(val interface{}, localizer localize.Localizer) error {

	search, ok := val.(string)
	if !ok {
		return commonerr.NewCastError(val, "string")
	}

	matched, _ := regexp.Match(legalNameChars, []byte(search))

	if matched {
		return nil
	}

	return errors.New(localizer.MustLocalize("kafka.topic.list.error.illegalSearchValue", localize.NewEntry("Search", search)))

}

// ValidatePartitionsN performs validation on the number of partitions v
func (v Validator) ValidatePartitionsN(val interface{}) error {
	partitionsStr := fmt.Sprintf("%v", v)

	partitions, err := strconv.Atoi(partitionsStr)
	if err != nil {
		return commonerr.NewCastError(v, "int32")
	}

	if partitions < minPartitions {
		return errors.New(v.Localizer.MustLocalize("kafka.topic.common.validation.partitions.error.invalid.minValue", localize.NewEntry("Partitions", partitions), localize.NewEntry("Min", minPartitions)))
	}

	if partitions > maxPartitions {
		return errors.New(v.Localizer.MustLocalize("kafka.topic.common.validation.partitions.error.invalid.maxValue", localize.NewEntry("Partitions", partitions), localize.NewEntry("Max", maxPartitions)))
	}

	return nil
}

// ValidateMessageRetentionPeriod validates the value (ms) of the retention period
// the valid values can range from [-1,...]
func (v Validator) ValidateMessageRetentionPeriod(val interface{}) error {
	retentionPeriodMsStr := fmt.Sprintf("%v", val)

	if retentionPeriodMsStr == "" {
		return nil
	}

	retentionPeriodMs, err := strconv.Atoi(retentionPeriodMsStr)
	if err != nil {
		return commonerr.NewCastError(val, "int")
	}

	if retentionPeriodMs < -1 {
		return errors.New(v.Localizer.MustLocalize("kafka.topic.common.validation.retentionPeriod.error.invalid", localize.NewEntry("RetentionPeriod", retentionPeriodMs)))
	}

	return nil
}

// ValidateMessageRetentionSize validates the value (bytes) of the retention size
// the valid values can range from [-1,...]
func (v Validator) ValidateMessageRetentionSize(val interface{}) error {
	retentionSizeStr := fmt.Sprintf("%v", val)

	if retentionSizeStr == "" {
		return nil
	}

	retentionPeriodBytes, err := strconv.Atoi(retentionSizeStr)
	if err != nil {
		return commonerr.NewCastError(val, "int")
	}

	if retentionPeriodBytes < -1 {
		return errors.New(v.Localizer.MustLocalize("kafka.topic.common.validation.retentionSize.error.invalid", localize.NewEntry("RetentionSize", retentionPeriodBytes)))
	}

	return nil
}

// ValidateNameIsAvailable checks if a topic with the given name already exists
func ValidateNameIsAvailable(api kafkainstanceclient.DefaultApi, instance string, localizer localize.Localizer) func(v interface{}) error {
	return func(v interface{}) error {
		name, _ := v.(string)

		_, httpRes, _ := api.GetTopic(context.Background(), name).Execute()

		if httpRes != nil && httpRes.StatusCode == 200 {
			return errors.New(localizer.MustLocalize("kafka.topic.create.error.conflictError", localize.NewEntry("TopicName", name), localize.NewEntry("InstanceName", instance)))
		}

		return nil
	}
}
