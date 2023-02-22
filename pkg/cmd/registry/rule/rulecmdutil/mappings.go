package rulecmdutil

import (
	registryinstanceclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal/client"
)

var ruleTypeMap = map[string]registryinstanceclient.RuleType{
	CompatibilityRule: registryinstanceclient.RULETYPE_COMPATIBILITY,
	ValidityRule:      registryinstanceclient.RULETYPE_VALIDITY,
}

var configMap = map[string]string{
	ConfigFULL:                "FULL",
	ConfigSYNTAX_ONLY:         "SYNTAX_ONLY",
	ConfigBACKWARD:            "BACKWARD",
	ConfigBACKWARD_TRANSITIVE: "BACKWARD_TRANSITIVE",
	ConfigFORWARD:             "FORWARD",
	ConfigFORWARD_TRANSITIVE:  "FORWARD_TRANSITIVE",
	ConfigFULL_TRANSITIVE:     "FULL_TRANSITIVE",
	ConfigNONE:                "NONE",
}

var ValidRules []string = []string{ValidityRule, CompatibilityRule}

// GetRuleTypeMap returns the mappings for rule types
func GetRuleTypeMap() map[string]registryinstanceclient.RuleType {
	return ruleTypeMap
}

// GetConfigMap returns the mappings for rule configurations
func GetConfigMap() map[string]string {
	return configMap
}

// GetMappedRuleType gets the mapped rule type value
func GetMappedRuleType(ruleType string) *registryinstanceclient.RuleType {
	r := ruleTypeMap[ruleType]
	return &r
}

// GetConfigMap gets the mapped configuration value
func GetMappedConfigValue(config string) string {
	return configMap[config]
}

var ValidRuleConfigs = map[string][]string{
	CompatibilityRule: {ConfigBACKWARD, ConfigBACKWARD_TRANSITIVE, ConfigFORWARD, ConfigFORWARD_TRANSITIVE, ConfigFULL, ConfigFULL_TRANSITIVE, ConfigNONE},
	ValidityRule:      {ConfigFULL, ConfigSYNTAX_ONLY, ConfigNONE},
}
