package groupcmdutil

import (
	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1/client"
)

// valid values for consumer group reset offset operaion
const (
	OffsetAbsolute  = string(kafkainstanceclient.OFFSETTYPE_ABSOLUTE)
	OffsetEarliest  = string(kafkainstanceclient.OFFSETTYPE_EARLIEST)
	OffsetTimestamp = string(kafkainstanceclient.OFFSETTYPE_TIMESTAMP)
	OffsetLatest    = string(kafkainstanceclient.OFFSETTYPE_LATEST)
)

var ValidOffsets = []string{OffsetAbsolute, OffsetEarliest, OffsetTimestamp, OffsetLatest}
