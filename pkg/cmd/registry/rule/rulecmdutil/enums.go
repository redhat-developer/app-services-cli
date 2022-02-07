package rulecmdutil

import (
	registryinstanceclient "github.com/redhat-developer/app-services-sdk-go/registryinstance/apiv1internal/client"
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

func GetRuleTypeMap() map[string]registryinstanceclient.RuleType {
	return ruleTypeMap
}

func GetConfigMap() map[string]string {
	return configMap
}

func GetMappedRuleType(ruleType string) *registryinstanceclient.RuleType {
	r := ruleTypeMap[ruleType]
	return &r
}

func GetMappedConfigValue(config string) string {
	return configMap[config]
}

var validRuleConfigs = map[string][]string{
	CompatibilityRule: {ConfigBACKWARD, ConfigBACKWARD_TRANSITIVE, ConfigFORWARD, ConfigFORWARD_TRANSITIVE, ConfigFULL, ConfigFULL_TRANSITIVE, ConfigNONE},
	ValidityRule:      {ConfigFULL, ConfigSYNTAX_ONLY},
}
