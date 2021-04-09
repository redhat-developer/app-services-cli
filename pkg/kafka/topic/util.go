package topic

import (
	"errors"
	"strconv"

	"github.com/redhat-developer/app-services-cli/internal/localizer"
	strimziadminclient "github.com/redhat-developer/app-services-cli/pkg/api/strimzi-admin/client"
)

var retentionMsKey string = "retention.ms"

// CreateConfig creates a list of topic ConfigEntries
func CreateConfig(retentionMs int) *[]strimziadminclient.ConfigEntry {
	retentionEntry := CreateRetentionConfigEntry(retentionMs)

	return &[]strimziadminclient.ConfigEntry{
		*retentionEntry,
	}
}

func CreateRetentionConfigEntry(retentionMs int) *strimziadminclient.ConfigEntry {
	retentionPeriodF := strconv.FormatInt(int64(retentionMs), 10)

	return &strimziadminclient.ConfigEntry{
		Key:   &retentionMsKey,
		Value: &retentionPeriodF,
	}
}

// ConvertPartitionsToInt converts the value from "partitions" to int32
func ConvertPartitionsToInt(partitionStr string) (int32, error) {

	patitionsInt, err := strconv.ParseInt(partitionStr, 10, 32)

	if err != nil {
		return 0, errors.New(localizer.MustLocalize(&localizer.Config{
			MessageID: "kafka.topic.common.input.partitions.error.invalid",
			TemplateData: map[string]interface{}{
				"Partition": partitionStr,
			},
		}))
	}

	return int32(patitionsInt), nil
}

// ConvertRetentionMsToInt converts the value from "retention-ms" to int
func ConvertRetentionMsToInt(retentionMsStr string) (int, error) {
	retentionMsInt, err := strconv.Atoi(retentionMsStr)

	if err != nil {
		return 0, errors.New(localizer.MustLocalize(&localizer.Config{
			MessageID: "kafka.topic.common.input.retentionMs.error.invalid",
			TemplateData: map[string]interface{}{
				"RetentionMs": retentionMsStr,
			},
		}))
	}

	return retentionMsInt, nil
}
