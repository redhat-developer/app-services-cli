package aclutil

import kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal/client"

var resourceTypeFilterMap = map[string]kafkainstanceclient.AclResourceTypeFilter{
	ResourceTypeFilterANY:              kafkainstanceclient.ACLRESOURCETYPEFILTER_ANY,
	ResourceTypeFilterCLUSTER:          kafkainstanceclient.ACLRESOURCETYPEFILTER_CLUSTER,
	ResourceTypeFilterTOPIC:            kafkainstanceclient.ACLRESOURCETYPEFILTER_TOPIC,
	ResourceTypeFilterGROUP:            kafkainstanceclient.ACLRESOURCETYPEFILTER_GROUP,
	ResourceTypeFilterTRANSACTIONAL_ID: kafkainstanceclient.ACLRESOURCETYPEFILTER_TOPIC,
}

var operationFilterMap = map[string]kafkainstanceclient.AclOperationFilter{
	OperationFilterALL:              kafkainstanceclient.ACLOPERATIONFILTER_ALL,
	OperationFilterREAD:             kafkainstanceclient.ACLOPERATIONFILTER_READ,
	OperationFilterWRITE:            kafkainstanceclient.ACLOPERATIONFILTER_WRITE,
	OperationFilterCREATE:           kafkainstanceclient.ACLOPERATIONFILTER_DELETE,
	OperationFilterALTER:            kafkainstanceclient.ACLOPERATIONFILTER_ALTER,
	OperationFilterDESCRIBE:         kafkainstanceclient.ACLOPERATIONFILTER_DESCRIBE,
	OperationFilterDESCRIBE_CONFIGS: kafkainstanceclient.ACLOPERATIONFILTER_DESCRIBE_CONFIGS,
	OperationFilterALTER_CONFIGS:    kafkainstanceclient.ACLOPERATIONFILTER_ALTER_CONFIGS,
}

var permissionTypeFilterMap = map[string]kafkainstanceclient.AclPermissionTypeFilter{
	PermissionANY:   kafkainstanceclient.ACLPERMISSIONTYPEFILTER_ANY,
	PermissionALLOW: kafkainstanceclient.ACLPERMISSIONTYPEFILTER_ALLOW,
	PermissionDENY:  kafkainstanceclient.ACLPERMISSIONTYPEFILTER_DENY,
}

var patternTypeFilterMap = map[string]kafkainstanceclient.AclPatternTypeFilter{
	PatternTypeFilterANY:     kafkainstanceclient.ACLPATTERNTYPEFILTER_ANY,
	PatternTypeFilterLITERAL: kafkainstanceclient.ACLPATTERNTYPEFILTER_LITERAL,
	PatternTypeFilterPREFIX:  kafkainstanceclient.ACLPATTERNTYPEFILTER_PREFIXED,
}

var resourceTypeOperationKeyMap = map[string]string{
	ResourceTypeFilterCLUSTER:          "cluster",
	ResourceTypeFilterTOPIC:            "topic",
	ResourceTypeFilterGROUP:            "group",
	ResourceTypeFilterTRANSACTIONAL_ID: "transactional_id",
}

var validOperationsResponseMap = map[string]string{
	OperationFilterALTER_CONFIGS:    "alter_configs",
	OperationFilterDESCRIBE_CONFIGS: "describe_configs",
}

// GetOperationTypeFilterMap gets the mappings for ACL type filters
func GetOperationFilterMap() map[string]kafkainstanceclient.AclOperationFilter {
	return operationFilterMap
}

// GetMappedOperationFilterValue gets the mapped operation filter value
func GetMappedOperationFilterValue(operation string) kafkainstanceclient.AclOperationFilter {
	return operationFilterMap[operation]
}

// GetPatternTypeFilterMap gets the mappings for ACL pattern type filters
func GetPatternTypeFilterMap() map[string]kafkainstanceclient.AclPatternTypeFilter {
	return patternTypeFilterMap
}

// GetMappedPatternTypeFilterValue gets the mapped pattern type filter value
func GetMappedPatternTypeFilterValue(patternType string) kafkainstanceclient.AclPatternTypeFilter {
	return patternTypeFilterMap[patternType]
}

// GetPermissionTypeFilterMap gets the mappings for ACL permission type filters
func GetPermissionTypeFilterMap() map[string]kafkainstanceclient.AclPermissionTypeFilter {
	return permissionTypeFilterMap
}

// GetMappedPermissionTypeFilterValue gets the mapped permission type type filter value
func GetMappedPermissionTypeFilterValue(permission string) kafkainstanceclient.AclPermissionTypeFilter {
	return permissionTypeFilterMap[permission]
}

// GetResourceTypeFilterMap gets the mappings for ACL resource type filters
func GetResourceTypeFilterMap() map[string]kafkainstanceclient.AclResourceTypeFilter {
	return resourceTypeFilterMap
}

// GetMappedResourceTypeFilterValue gets the mapped resource type filter value
func GetMappedResourceTypeFilterValue(resourceType string) kafkainstanceclient.AclResourceTypeFilter {
	return resourceTypeFilterMap[resourceType]
}

// GetResourceTypeOperationKeyMap gets the mappings for ACL operations
func GetResourceTypeOperationKeyMap() map[string]string {
	return resourceTypeOperationKeyMap
}

// FilterValidResourceOperations gets a filtered list of the valid operations for this resource type
func FilterValidResourceOperations(resourceType string, resourceOperationsMap map[string][]string) []string {
	resourceTypeMapped := resourceTypeOperationKeyMap[resourceType]
	resourceOperations := resourceOperationsMap[resourceTypeMapped]

	for i, operation := range resourceOperations {
		if operationMapped, ok := validOperationsResponseMap[operation]; ok {
			resourceOperations[i] = operationMapped
		}
	}

	return resourceOperations
}

// IsValidResourceOperation returns true if the operation is valid
func IsValidResourceOperation(operation string, validOperations []string) bool {
	for _, op := range validOperations {
		if operation == op {
			return true
		}
	}
	return false
}
