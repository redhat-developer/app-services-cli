package topic

import (
	"errors"
	"fmt"
	"regexp"

	"strconv"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
)

const (
	legalNameChars       = "^[a-zA-Z0-9\\_\\-]+$"
	maxNameLength        = 249
	minReplicationFactor = 1
	minPartitions        = 1
)

// ValidateName validates the name of the topic
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

	if len(name) < 1 {
		return errors.New(localizer.MustLocalizeFromID("kafka.topic.common.validation.name.error.required"))
	} else if len(name) > maxNameLength {
		return errors.New(localizer.MustLocalize(&localizer.Config{
			MessageID: "kafka.topic.common.validation.name.error.lengthError",
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
		MessageID: "kafka.topic.common.validation.name.error.invalidChars",
		TemplateData: map[string]interface{}{
			"Name": name,
		},
	}))
}

// ValidatePartitionsN performs validation on the number of partitions v
func ValidatePartitionsN(v interface{}) error {
	partitionsStr := fmt.Sprintf("%v", v)

	partitions, err := strconv.Atoi(partitionsStr)
	if err != nil {
		return errors.New(localizer.MustLocalize(&localizer.Config{
			MessageID: "common.error.castError",
			TemplateData: map[string]interface{}{
				"Value": v,
				"Type":  "int32",
			},
		}))
	}

	if partitions < minPartitions {
		return errors.New(localizer.MustLocalize(&localizer.Config{
			MessageID: "kafka.topic.common.validation.partitions.error.invalid",
			TemplateData: map[string]interface{}{
				"Partitions":    partitions,
				"MinPartitions": minPartitions,
			},
		}))
	}

	return nil
}

// ValidationReplicationFactorN performs validation on the number of replicas v
func ValidateReplicationFactorN(v interface{}) error {
	replicas, ok := v.(int32)
	if !ok {
		return errors.New(localizer.MustLocalize(&localizer.Config{
			MessageID: "common.error.castError",
			TemplateData: map[string]interface{}{
				"Value": v,
				"Type":  "int32",
			},
		}))
	}

	if replicas < minReplicationFactor {
		return errors.New(localizer.MustLocalize(&localizer.Config{
			MessageID: "kafka.topic.common.validation.replicationFactor.error.invalid",
			TemplateData: map[string]interface{}{
				"ReplicationFactor":    replicas,
				"MinReplicationFactor": minReplicationFactor,
			},
		}))
	}

	return nil
}

// ValidateMessageRetentionPeriod validates the value (ms) of the retention period
// the valid values can range from [-1,...]
func ValidateMessageRetentionPeriod(v interface{}) error {
	retentionPeriodMsStr := fmt.Sprintf("%v", v)

	retentionPeriodMs, err := strconv.Atoi(retentionPeriodMsStr)
	if err != nil {
		return errors.New(localizer.MustLocalize(&localizer.Config{
			MessageID: "common.error.castError",
			TemplateData: map[string]interface{}{
				"Value": v,
				"Type":  "int",
			},
		}))
	}

	if retentionPeriodMs < -1 {
		return errors.New(localizer.MustLocalize(&localizer.Config{
			MessageID: "kafka.topic.common.validation.retentionPeriod.error.invalid",
			TemplateData: map[string]interface{}{
				"RetentionPeriod": retentionPeriodMs,
			},
		}))
	}

	return nil
}
