package consumergroup

import (
	strimziadminclient "github.com/redhat-developer/app-services-cli/pkg/api/strimzi-admin/client"
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

func GetActiveConsumersCount(consumers []strimziadminclient.Consumer) (count int) {
	for _, c := range consumers {
		if c.Partition != -1 {
			count++
		}
	}
	return count
}
