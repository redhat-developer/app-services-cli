package consumergroup

import (
	"errors"
	"regexp"

	"github.com/redhat-developer/app-services-cli/pkg/localize"
	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal/client"
)

// valid values for consumer group reset offset operaion
const (
	AbsoluteOffset  = "absolute"
	EarliestOffset  = "earliest"
	TimestampOffset = "timestamp"
	LatestOffset    = "latest"
)

var timestampOffsetRegExp = regexp.MustCompile(`^(\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}-\d{2}:\d{2})$`)

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

// ValidateTimestampValue validates the value for timestamp offset
// value should be in format "yyyy-MM-dd'T'HH:mm:ss"
func ValidateTimestampValue(localizer localize.Localizer, time string) error {
	offsetValueTmplPair := localize.NewEntry("Value", time)
	matched := timestampOffsetRegExp.MatchString(time)

	if matched {
		return nil
	}

	return errors.New(localizer.MustLocalize("kafka.consumerGroup.resetOffset.error.invalidTimestampOffset", offsetValueTmplPair))
}
