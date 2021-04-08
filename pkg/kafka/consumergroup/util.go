package consumergroup

import (
	strimziadminclient "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/strimzi-admin/client"
)

// GetPartitionsWithLag returns the number of partitions having lag for a consumer group
func GetPartitionsWithLag(consumers []strimziadminclient.Consumer) (partitionsWithLag int) {
	for _, consumer := range consumers {
		if consumer.Lag > 0 {
			partitionsWithLag++
		}
	}

	return partitionsWithLag
}
