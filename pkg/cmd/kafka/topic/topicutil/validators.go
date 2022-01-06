package topicutil

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/factory"
	"github.com/redhat-developer/app-services-cli/pkg/core/connection"
	"github.com/redhat-developer/app-services-cli/pkg/core/errors"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
)

const (
	legalNameChars = "^[a-zA-Z0-9._-]+$"
	maxNameLength  = 249
	minPartitions  = 1
	maxPartitions  = 100
)

// Validator is a type for validating Kafka topic configuration values
type Validator struct {
	Localizer     localize.Localizer
	InstanceID    string
	Connection    factory.ConnectionFunc
	CurPartitions int
}

// ValidateName validates the name of the topic
func (v *Validator) ValidateName(val interface{}) error {
	name, ok := val.(string)
	if !ok {
		return errors.NewCastError(val, "string")
	}

	if len(name) < 1 {
		return v.Localizer.MustLocalizeError("kafka.topic.common.validation.name.error.required")
	} else if len(name) > maxNameLength {
		return v.Localizer.MustLocalizeError("kafka.topic.common.validation.name.error.lengthError", localize.NewEntry("MaxNameLen", maxNameLength))
	}

	if (name == ".") || (name == "..") {
		return v.Localizer.MustLocalizeError("kafka.topic.common.validation.name.error.dotsError")
	}

	matched, _ := regexp.Match(legalNameChars, []byte(name))

	if matched {
		return nil
	}

	return v.Localizer.MustLocalizeError("kafka.topic.common.validation.name.error.invalidChars", localize.NewEntry("Name", name))
}

func (v *Validator) ValidateSearchInput(val interface{}) error {
	search, ok := val.(string)
	if !ok {
		return errors.NewCastError(val, "string")
	}

	matched, _ := regexp.Match(legalNameChars, []byte(search))

	if matched {
		return nil
	}

	return v.Localizer.MustLocalizeError("kafka.topic.list.error.illegalSearchValue", localize.NewEntry("Search", search))
}

// ValidatePartitionsN performs validation on the number of partitions v
func (v *Validator) ValidatePartitionsN(val interface{}) error {
	partitionsStr := fmt.Sprintf("%v", val)

	if partitionsStr == "" {
		return nil
	}

	partitions, err := strconv.Atoi(partitionsStr)
	if err != nil {
		return errors.NewCastError(val, "int32")
	}

	if partitions < minPartitions {
		return v.Localizer.MustLocalizeError("kafka.topic.common.validation.partitions.error.invalid.minValue", localize.NewEntry("Partitions", partitions), localize.NewEntry("Min", minPartitions))
	}

	if partitions < v.CurPartitions {
		return v.Localizer.MustLocalizeError("kafka.topic.common.validation.partitions.error.invalid.lesserValue", localize.NewEntry("CurrPartitions", v.CurPartitions), localize.NewEntry("Partitions", partitions))
	}

	if partitions > maxPartitions {
		return v.Localizer.MustLocalizeError("kafka.topic.common.validation.partitions.error.invalid.maxValue", localize.NewEntry("Partitions", partitions), localize.NewEntry("Max", maxPartitions))
	}

	return nil
}

// ValidateMessageRetentionPeriod validates the value (ms) of the retention period
// the valid values can range from [-1,...]
func (v *Validator) ValidateMessageRetentionPeriod(val interface{}) error {
	retentionPeriodMsStr := fmt.Sprintf("%v", val)

	if retentionPeriodMsStr == "" {
		return nil
	}

	retentionPeriodMs, err := strconv.Atoi(retentionPeriodMsStr)
	if err != nil {
		return errors.NewCastError(val, "int")
	}

	if retentionPeriodMs < -1 {
		return v.Localizer.MustLocalizeError("kafka.topic.common.validation.retentionPeriod.error.invalid", localize.NewEntry("RetentionPeriod", retentionPeriodMs))
	}

	return nil
}

// ValidateMessageRetentionSize validates the value (bytes) of the retention size
// the valid values can range from [-1,...]
func (v *Validator) ValidateMessageRetentionSize(val interface{}) error {
	retentionSizeStr := fmt.Sprintf("%v", val)

	if retentionSizeStr == "" {
		return nil
	}

	retentionPeriodBytes, err := strconv.Atoi(retentionSizeStr)
	if err != nil {
		return errors.NewCastError(val, "int")
	}

	if retentionPeriodBytes < -1 {
		return v.Localizer.MustLocalizeError("kafka.topic.common.validation.retentionSize.error.invalid", localize.NewEntry("RetentionSize", retentionPeriodBytes))
	}

	return nil
}

// ValidateNameIsAvailable checks if a topic with the given name already exists
func (v *Validator) ValidateNameIsAvailable(val interface{}) error {
	name := fmt.Sprintf("%v", val)

	conn, err := v.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	api, kafkaInstance, err := conn.API().KafkaAdmin(v.InstanceID)
	if err != nil {
		return err
	}

	_, httpRes, _ := api.TopicsApi.GetTopic(context.Background(), name).Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
		if httpRes.StatusCode == http.StatusOK {
			return v.Localizer.MustLocalizeError("kafka.topic.create.error.conflictError", localize.NewEntry("TopicName", name), localize.NewEntry("InstanceName", kafkaInstance.GetName()))
		}
	}

	return nil
}
