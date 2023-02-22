package topiccmdutil

import (
	"fmt"
	"strconv"

	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/kafkainstance/apiv1/client"
)

var (
	RetentionMsKey   = "retention.ms"
	RetentionSizeKey = "retention.bytes"
	CleanupPolicy    = "cleanup.policy"
)

var ValidCleanupPolicies = []string{"delete", "compact", "compact,delete"}

// CreateConfigEntries converts a key value map of config entries to an array of config entries
func CreateConfigEntries(entryMap map[string]*string) *[]kafkainstanceclient.ConfigEntry {
	var entries []kafkainstanceclient.ConfigEntry
	for key, value := range entryMap {
		if value != nil {
			// nolint:scopelint
			entry := kafkainstanceclient.NewConfigEntry(key, *value)
			entries = append(entries, *entry)
		}
	}
	return &entries
}

// ConvertPartitionsToInt converts the value from "partitions" to int32
func ConvertPartitionsToInt(partitionStr string) (int32, error) {
	patitionsInt, err := strconv.ParseInt(partitionStr, 10, 32)
	if err != nil {
		err = fmt.Errorf("invalid value for partitions: %v", partitionStr)
		return 0, err
	}

	return int32(patitionsInt), nil
}

// ConvertRetentionMsToInt converts the value from "retention-ms" to int
func ConvertRetentionMsToInt(retentionMsStr string) (int, error) {
	retentionMsInt, err := strconv.Atoi(retentionMsStr)
	if err != nil {
		return 0, fmt.Errorf("invalid value for retention period: %v", retentionMsInt)
	}

	return retentionMsInt, nil
}

// ConvertRetentionBytesToInt converts the value from "retention-bytes" to int
func ConvertRetentionBytesToInt(retentionBytesStr string) (int, error) {
	retentionMsInt, err := strconv.Atoi(retentionBytesStr)
	if err != nil {
		return 0, fmt.Errorf("invalid value for retention size: %v", retentionMsInt)
	}

	return retentionMsInt, nil
}

func GetConfigValue(configEntries []kafkainstanceclient.ConfigEntry, keyName string) (val string) {
	for _, config := range configEntries {

		if config.Key == keyName {
			val = config.GetValue()
		}
	}

	return val
}
