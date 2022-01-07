package sdk

import (
	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal/client"
)

// valid values for consumer group reset offset operaion
const (
	OffsetAbsolute  = "absolute"
	OffsetEarliest  = "earliest"
	OffsetTimestamp = "timestamp"
	OffsetLatest    = "latest"
)

var ValidOffsets = []string{OffsetAbsolute, OffsetEarliest, OffsetTimestamp, OffsetLatest}

// GetPartitionsWithLag returns the number of partitions having lag for a consumer group
func GetPartitionsWithLag(consumers []kafkainstanceclient.Consumer) (partitionsWithLag int) {
	for _, consumer := range consumers {
		if consumer.Lag > 0 {
			partitionsWithLag++
		}
	}

	return partitionsWithLag
}

func GetActiveConsumersCount(consumers []kafkainstanceclient.Consumer) (count int) {
	for _, c := range consumers {
		if c.Partition != -1 {
			count++
		}
	}
	return count
}

func GetUnassignedPartitions(consumers []kafkainstanceclient.Consumer) (unassignedPartitions int) {
	for _, c := range consumers {
		if c.GetMemberId() == "" {
			unassignedPartitions++
		}
	}
	return unassignedPartitions
}
