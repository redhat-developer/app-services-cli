package common

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
	OperationFilterANY:              kafkainstanceclient.ACLOPERATIONFILTER_ANY,
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

func GetResourceTypeFilter(resourceType string) kafkainstanceclient.AclResourceTypeFilter {
	return resourceTypeFilterMap[resourceType]
}

func GetOperationFilter(operation string) kafkainstanceclient.AclOperationFilter {
	return operationFilterMap[operation]
}

func GetPatternTypeFilter(patternType string) kafkainstanceclient.AclPatternTypeFilter {
	return patternTypeFilterMap[patternType]
}

func GetPermissionFilter(permission string) kafkainstanceclient.AclPermissionTypeFilter {
	return permissionTypeFilterMap[permission]
}

func GetValidResourceOperations(resourceType string, resourceOperationsMap map[string][]string) []string {
	resourceTypeMapped := resourceTypeOperationKeyMap[resourceType]
	resourceOperations := resourceOperationsMap[resourceTypeMapped]

	for i, operation := range resourceOperations {
		if operationMapped, ok := validOperationsResponseMap[operation]; ok {
			resourceOperations[i] = operationMapped
		}
	}

	return resourceOperations
}

func IsValidOperation(operation string, validOperations []string) bool {
	for _, op := range validOperations {
		if operation == op {
			return true
		}
	}
	return false
}
