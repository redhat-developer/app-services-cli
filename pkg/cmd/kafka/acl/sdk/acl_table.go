package sdk

import (
	"fmt"
	"strings"

	"github.com/redhat-developer/app-services-cli/pkg/core/localize"

	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal/client"
)

type permissionsRow struct {
	Principal   string `json:"principal,omitempty" header:"Principal"`
	Permission  string `json:"permission,omitempty" header:"Permission"`
	Operation   string `json:"operation,omitempty" header:"Operation"`
	Description string `json:"description,omitempty" header:"description"`
}

// MapACLsToTableRows converts a list of ACL bindings into a formatted table for printing
func MapACLsToTableRows(bindings []kafkainstanceclient.AclBinding, localizer localize.Localizer) []permissionsRow {
	rows := make([]permissionsRow, len(bindings))

	// get the SDK => CLI key mappings
	permissionMap := GetPermissionTypeMap()
	reversedPermissionMap := make(map[kafkainstanceclient.AclPermissionType]string)
	for k, v := range permissionMap {
		reversedPermissionMap[v] = k
	}

	// get the SDK => CLI key mappings
	operationMap := GetOperationMap()
	reversedOperationMap := make(map[kafkainstanceclient.AclOperation]string)
	for k, v := range operationMap {
		reversedOperationMap[v] = k
	}

	// get the SDK => CLI key mappings
	resourceTypeMap := GetResourceTypeMap()
	reversedResourceTypeMap := make(map[kafkainstanceclient.AclResourceType]string)
	for k, v := range resourceTypeMap {
		reversedResourceTypeMap[v] = k
	}

	for i, p := range bindings {
		description := formatTablePatternType(p.PatternType, localizer)
		row := permissionsRow{
			Principal:   formatTablePrincipal(p.GetPrincipal(), localizer),
			Permission:  reversedPermissionMap[p.GetPermission()],
			Operation:   reversedOperationMap[p.GetOperation()],
			Description: fmt.Sprintf("%s %s \"%s\"", reversedResourceTypeMap[p.GetResourceType()], description, p.GetResourceName()),
		}
		rows[i] = row
	}
	return rows
}

func formatTablePatternType(patternType kafkainstanceclient.AclPatternType, localizer localize.Localizer) string {
	if patternType == kafkainstanceclient.ACLPATTERNTYPE_LITERAL {
		return localizer.MustLocalize("kafka.acl.list.is")
	}

	return localizer.MustLocalize("kafka.acl.list.startsWith")
}

func formatTablePrincipal(principal string, localizer localize.Localizer) string {
	s := strings.Split(principal, ":")[1]

	if s == Wildcard {
		return localizer.MustLocalize("kafka.acl.list.allAccounts")
	}

	return s
}
