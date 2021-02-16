package topic

import (
	"fmt"
	"regexp"
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
		return fmt.Errorf("could not cast %v to string", val)
	}

	if len(name) < 1 {
		return fmt.Errorf("topic name is required")
	} else if len(name) > maxNameLength {
		return fmt.Errorf("topic name cannot exceed %v characters", maxNameLength)
	}

	matched, _ := regexp.Match(legalNameChars, []byte(name))

	if matched {
		return nil
	}

	return fmt.Errorf("invalid topic name \"%v\", only letters (Aa-Zz), numbers, '_' and '-' are accepted", name)
}

// ValidatePartitionsN performs validation on the number of partitions v
func ValidatePartitionsN(v interface{}) error {
	partitions, ok := v.(int32)
	if !ok {
		return fmt.Errorf("could not cast %v to int32", v)
	}

	if partitions < minPartitions {
		return fmt.Errorf("invalid partition count %v, minimum partition count is %v", partitions, minPartitions)
	}

	return nil
}

// ValidationReplicationFactorN performs validation on the number of replicas v
func ValidateReplicationFactorN(v interface{}) error {
	replicas, ok := v.(int32)
	if !ok {
		return fmt.Errorf("could not cast %v to int32", v)
	}

	if replicas < minReplicationFactor {
		return fmt.Errorf("invalid replication factor %v, minimum replication factor is %v", replicas, minReplicationFactor)
	}

	return nil
}

// ValidateMessageRetentionPeriod validates the value (ms) of the retention period
// the valid values can range from [-1,...]
func ValidateMessageRetentionPeriod(v interface{}) error {
	retentionPeriodMs, ok := v.(int)
	if !ok {
		return fmt.Errorf("could not cast %v to int", v)
	}

	if retentionPeriodMs < -1 {
		return fmt.Errorf("invalid retention period %v, minimum value is -1", retentionPeriodMs)
	}

	return nil
}
