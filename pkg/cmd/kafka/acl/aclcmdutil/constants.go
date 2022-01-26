package aclcmdutil

const (
	Wildcard     = "*"
	KafkaCluster = "kafka-cluster"
	AllAlias     = "all"
)

const (
	ResourceTypeANY              = "any"
	ResourceTypeTOPIC            = "topic"
	ResourceTypeCLUSTER          = "cluster"
	ResourceTypeGROUP            = "group"
	ResourceTypeTRANSACTIONAL_ID = "transactional-id"
)

const (
	PermissionALLOW = "allow"
	PermissionDENY  = "deny"
	PermissionANY   = "any"
)

const (
	OperationALL              = "all"
	OperationREAD             = "read"
	OperationWRITE            = "write"
	OperationCREATE           = "create"
	OperationDELETE           = "delete"
	OperationALTER            = "alter"
	OperationDESCRIBE         = "describe"
	OperationDESCRIBE_CONFIGS = "describe-configs"
	OperationALTER_CONFIGS    = "alter-configs"
)

const (
	PatternTypeLITERAL = "literal"
	PatternTypePREFIX  = "prefix"
	PatternTypeANY     = "any"
)
