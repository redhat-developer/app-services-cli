package aclutil

import kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal/client"

var resourceTypeFilterMap = map[string]kafkainstanceclient.AclResourceTypeFilter{
	ResourceTypeANY:              kafkainstanceclient.ACLRESOURCETYPEFILTER_ANY,
	ResourceTypeCLUSTER:          kafkainstanceclient.ACLRESOURCETYPEFILTER_CLUSTER,
	ResourceTypeTOPIC:            kafkainstanceclient.ACLRESOURCETYPEFILTER_TOPIC,
	ResourceTypeGROUP:            kafkainstanceclient.ACLRESOURCETYPEFILTER_GROUP,
	ResourceTypeTRANSACTIONAL_ID: kafkainstanceclient.ACLRESOURCETYPEFILTER_TRANSACTIONAL_ID,
}

var operationFilterMap = map[string]kafkainstanceclient.AclOperationFilter{
	OperationALL:              kafkainstanceclient.ACLOPERATIONFILTER_ALL,
	OperationREAD:             kafkainstanceclient.ACLOPERATIONFILTER_READ,
	OperationWRITE:            kafkainstanceclient.ACLOPERATIONFILTER_WRITE,
	OperationCREATE:           kafkainstanceclient.ACLOPERATIONFILTER_CREATE,
	OperationDELETE:           kafkainstanceclient.ACLOPERATIONFILTER_DELETE,
	OperationALTER:            kafkainstanceclient.ACLOPERATIONFILTER_ALTER,
	OperationDESCRIBE:         kafkainstanceclient.ACLOPERATIONFILTER_DESCRIBE,
	OperationDESCRIBE_CONFIGS: kafkainstanceclient.ACLOPERATIONFILTER_DESCRIBE_CONFIGS,
	OperationALTER_CONFIGS:    kafkainstanceclient.ACLOPERATIONFILTER_ALTER_CONFIGS,
}

var operationMap = map[string]kafkainstanceclient.AclOperation{
	OperationALL:              kafkainstanceclient.ACLOPERATION_ALL,
	OperationREAD:             kafkainstanceclient.ACLOPERATION_READ,
	OperationWRITE:            kafkainstanceclient.ACLOPERATION_WRITE,
	OperationCREATE:           kafkainstanceclient.ACLOPERATION_CREATE,
	OperationDELETE:           kafkainstanceclient.ACLOPERATION_DELETE,
	OperationALTER:            kafkainstanceclient.ACLOPERATION_ALTER,
	OperationDESCRIBE:         kafkainstanceclient.ACLOPERATION_DESCRIBE,
	OperationDESCRIBE_CONFIGS: kafkainstanceclient.ACLOPERATION_DESCRIBE_CONFIGS,
	OperationALTER_CONFIGS:    kafkainstanceclient.ACLOPERATION_ALTER_CONFIGS,
}

var permissionTypeFilterMap = map[string]kafkainstanceclient.AclPermissionTypeFilter{
	PermissionANY:   kafkainstanceclient.ACLPERMISSIONTYPEFILTER_ANY,
	PermissionALLOW: kafkainstanceclient.ACLPERMISSIONTYPEFILTER_ALLOW,
	PermissionDENY:  kafkainstanceclient.ACLPERMISSIONTYPEFILTER_DENY,
}

var permissionTypeMap = map[string]kafkainstanceclient.AclPermissionType{
	PermissionALLOW: kafkainstanceclient.ACLPERMISSIONTYPE_ALLOW,
	PermissionDENY:  kafkainstanceclient.ACLPERMISSIONTYPE_DENY,
}

var patternTypeFilterMap = map[string]kafkainstanceclient.AclPatternTypeFilter{
	PatternTypeANY:     kafkainstanceclient.ACLPATTERNTYPEFILTER_ANY,
	PatternTypeLITERAL: kafkainstanceclient.ACLPATTERNTYPEFILTER_LITERAL,
	PatternTypePREFIX:  kafkainstanceclient.ACLPATTERNTYPEFILTER_PREFIXED,
}

var patternTypeMap = map[string]kafkainstanceclient.AclPatternType{
	PatternTypeLITERAL: kafkainstanceclient.ACLPATTERNTYPE_LITERAL,
	PatternTypePREFIX:  kafkainstanceclient.ACLPATTERNTYPE_PREFIXED,
}

// for backwards-compatibility, two resource types are possible
var resourceTypeOperationKeyMap = map[string][]string{
	ResourceTypeCLUSTER:          {"cluster", string(kafkainstanceclient.ACLRESOURCETYPE_CLUSTER)},
	ResourceTypeTOPIC:            {"topic", string(kafkainstanceclient.ACLRESOURCETYPE_TOPIC)},
	ResourceTypeGROUP:            {"group", string(kafkainstanceclient.ACLRESOURCETYPE_GROUP)},
	ResourceTypeTRANSACTIONAL_ID: {"transactional_id", string(kafkainstanceclient.ACLRESOURCETYPEFILTER_TRANSACTIONAL_ID)},
}

var validOperationsResponseMap = map[string]string{
	"alter_configs":    OperationALTER_CONFIGS,
	"describe_configs": OperationDESCRIBE_CONFIGS,
}

// GetOperationTypeFilterMap gets the mappings for ACL type filters
func GetOperationFilterMap() map[string]kafkainstanceclient.AclOperationFilter {
	return operationFilterMap
}

func GetOperationMap() map[string]kafkainstanceclient.AclOperation {
	return operationMap
}

// GetReversedOperationMap returns a map of operations with the SDK enums as the keys
func GetReversedOperationMap() map[kafkainstanceclient.AclOperation]string {
	reversedMap := make(map[kafkainstanceclient.AclOperation]string)
	for k, v := range operationMap {
		reversedMap[v] = k
	}

	return reversedMap
}

// GetMappedOperationFilterValue gets the mapped operation filter value
func GetMappedOperationFilterValue(operation string) kafkainstanceclient.AclOperationFilter {
	return operationFilterMap[operation]
}

// GetMappedOperationValue gets the mapped operation value
func GetMappedOperationValue(operation string) kafkainstanceclient.AclOperation {
	return operationMap[operation]
}

// GetPatternTypeFilterMap gets the mappings for ACL pattern type filters
func GetPatternTypeFilterMap() map[string]kafkainstanceclient.AclPatternTypeFilter {
	return patternTypeFilterMap
}

// GetPatternTypeMap gets the mappings for ACL pattern type
func GetPatternTypeMap() map[string]kafkainstanceclient.AclPatternType {
	return patternTypeMap
}

// GetMappedPatternTypeFilterValue gets the mapped pattern type filter value
func GetMappedPatternTypeFilterValue(patternType string) kafkainstanceclient.AclPatternTypeFilter {
	return patternTypeFilterMap[patternType]
}

// GetMappedPatternTypeValue gets the mapped pattern type value
func GetMappedPatternTypeValue(patternType string) kafkainstanceclient.AclPatternType {
	return patternTypeMap[patternType]
}

// GetPermissionTypeFilterMap gets the mappings for ACL permission type filters
func GetPermissionTypeFilterMap() map[string]kafkainstanceclient.AclPermissionTypeFilter {
	return permissionTypeFilterMap
}

// GetPermissionTypeMap gets the mappings for ACL permission types
func GetPermissionTypeMap() map[string]kafkainstanceclient.AclPermissionType {
	return permissionTypeMap
}

// GetMappedPermissionTypeFilterValue gets the mapped permission type filter value
func GetMappedPermissionTypeFilterValue(permission string) kafkainstanceclient.AclPermissionTypeFilter {
	return permissionTypeFilterMap[permission]
}

// GetMappedPermissionTypeValue gets the mapped permission type value
func GetMappedPermissionTypeValue(permission string) kafkainstanceclient.AclPermissionType {
	return permissionTypeMap[permission]
}

// GetResourceTypeFilterMap gets the mappings for ACL resource type filters
func GetResourceTypeFilterMap() map[string]kafkainstanceclient.AclResourceTypeFilter {
	return resourceTypeFilterMap
}

// GetMappedResourceTypeFilterValue gets the mapped resource type filter value
func GetMappedResourceTypeFilterValue(resourceType string) kafkainstanceclient.AclResourceTypeFilter {
	return resourceTypeFilterMap[resourceType]
}

// getValidOperationsResponseMap returns a map of the ACL
// resource operations to the values used in this package
// eg: "ALTER": "alter"
func getValidOperationsResponseMap() map[string]string {
	reverseOperationMap := GetReversedOperationMap()
	// The API is returning internal Kafka operations, e.g describe_configs
	// It will change to align with the API enums, eg: DESCRIBE_CONFIGS
	// This will provide backwards compatibility for both and can be removed in a later release
	for op, v := range reverseOperationMap {
		validOperationsResponseMap[string(op)] = string(v)
	}
	return validOperationsResponseMap
}
