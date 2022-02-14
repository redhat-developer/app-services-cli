package rulecmdutil

import (
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
)

// Validator is a struct to validate inputs for service registry rule operations
type Validator struct {
	Localizer localize.Localizer
}

// ValidateRuleType checks if value v is a valid value for --offset
func (v *Validator) ValidateRuleType(ruleType string) error {
	isValid := flagutil.IsValidInput(ruleType, ValidRules...)

	if isValid {
		return nil
	}

	return flagutil.InvalidValueError("rule-type", v, ValidRules...)
}

func (v *Validator) IsValidRuleConfig(ruleType string, config string) (bool, []string) {

	validConfigs := ValidRuleConfigs[ruleType]

	for _, validConfig := range validConfigs {
		if validConfig == config {
			return true, []string{}
		}
	}

	return false, validConfigs
}
