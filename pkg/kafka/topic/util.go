package topic

import (
	"strconv"

	strimziadminclient "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/strimzi-admin/client"
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
