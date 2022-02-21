package rulecmdutil

const (
	ValidityRule      = "validity"
	CompatibilityRule = "compatibility"
)

const (
	ConfigFULL                = "full"
	ConfigSYNTAX_ONLY         = "syntax-only"
	ConfigFULL_TRANSITIVE     = "full-transitive"
	ConfigBACKWARD            = "backward"
	ConfigBACKWARD_TRANSITIVE = "backward-transitive"
	ConfigFORWARD             = "forward"
	ConfigFORWARD_TRANSITIVE  = "forward-transitive"
	ConfigNONE                = "none"
)

var ValidRuleTypes = []string{ValidityRule, CompatibilityRule}
