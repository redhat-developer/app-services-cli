package topic

import (
	"errors"
	"strconv"

	"github.com/redhat-developer/app-services-cli/internal/localizer"
	strimziadminclient "github.com/redhat-developer/app-services-cli/pkg/api/strimzi-admin/client"
)

var RetentionMsKey string = "retention.ms"
var RetentionSizeKey string = "retention.bytes"

// CreateConfigEntries converts a key value map of config entries to an array of config entries
func CreateConfigEntries(entryMap map[string]*string) *[]strimziadminclient.ConfigEntry {
	entries := []strimziadminclient.ConfigEntry{}
	for key, value := range entryMap {
		if value != nil {
			// nolint:scopelint
			entry := strimziadminclient.NewConfigEntry()
			entry.SetKey(key)
			entry.SetValue(*value)
			entries = append(entries, *entry)
		}
	}
	return &entries
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
