package aclutil

const (
	Wildcard     = "*"
	KafkaCluster = "kafka-cluster"
)

const (
	ResourceTypeFilterANY              = "any"
	ResourceTypeFilterTOPIC            = "topic"
	ResourceTypeFilterCLUSTER          = "cluster"
	ResourceTypeFilterGROUP            = "group"
	ResourceTypeFilterTRANSACTIONAL_ID = "transactional-id"
)

const (
	PermissionALLOW = "allow"
	PermissionDENY  = "deny"
	PermissionANY   = "any"
)

const (
	OperationFilterALL              = "all"
	OperationFilterREAD             = "read"
	OperationFilterWRITE            = "write"
	OperationFilterCREATE           = "create"
	OperationFilterDELETE           = "delete"
	OperationFilterALTER            = "alter"
	OperationFilterDESCRIBE         = "describe"
	OperationFilterDESCRIBE_CONFIGS = "describe-configs"
	OperationFilterALTER_CONFIGS    = "alter-configs"
)

const (
	PatternTypeFilterLITERAL = "literal"
	PatternTypeFilterPREFIX  = "prefix"
	PatternTypeFilterANY     = "any"
)
