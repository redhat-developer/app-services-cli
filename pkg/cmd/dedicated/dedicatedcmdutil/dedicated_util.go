package kafkacmdutil

import (
	"fmt"
	"github.com/redhat-developer/app-services-cli/pkg/core/errors"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"strconv"
)

// Validator is a type for validating Kafka configuration values
type Validator struct {
	Localizer  localize.Localizer
	Connection factory.ConnectionFunc
}

func ValidateMachinePoolCount(count int) bool {
	// check if the count is a multiple of 3 and greater than or equal to 3
	if count%3 == 0 && count >= 3 {
		return true
	}
	return false
}

func (v *Validator) ValidatorForMachinePoolNodes(val interface{}) error {
	value := fmt.Sprintf("%v", val)
	if val == "" {
		return errors.NewCastError(val, "emtpy string")
	}
	value1, err := strconv.Atoi(value)
	if err != nil {
		return errors.NewCastError(val, "integer")
	}
	if !ValidateMachinePoolCount(value1) {
		return fmt.Errorf("invalid input, machine pool node count must be greater than or equal to 3 and it " +
			"must be a is a multiple of 3")
	}
	return nil
}
