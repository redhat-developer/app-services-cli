package acl

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/redhat-developer/app-services-cli/pkg/localize"
	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal/client"
)

type permissionsRow struct {
	Principal   string `json:"principal,omitempty" header:"Principal"`
	Permission  string `json:"permission,omitempty" header:"permission"`
	Description string `json:"description,omitempty" header:"description"`
}

// ExecuteACLRuleCreate makes request to create an ACL rule
func ExecuteACLRuleCreate(req kafkainstanceclient.ApiCreateAclRequest, localizer localize.Localizer, kafkaInstanceName string) error {

	httpRes, err := req.Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	if err != nil {
		if httpRes == nil {
			return err
		}

		operationTmplPair := localize.NewEntry("Operation", "create")

		switch httpRes.StatusCode {
		case http.StatusUnauthorized:
			return localizer.MustLocalizeError("kafka.acl.common.error.unauthorized", operationTmplPair)
		case http.StatusForbidden:
			return localizer.MustLocalizeError("kafka.acl.common.error.forbidden", operationTmplPair)
		case http.StatusInternalServerError:
			return localizer.MustLocalizeError("kafka.acl.common.error.internalServerError")
		case http.StatusServiceUnavailable:
			return localizer.MustLocalizeError("kafka.acl.common.error.unableToConnectToKafka", localize.NewEntry("Name", kafkaInstanceName))
		default:
			return err
		}
	}

	return nil
}

// MapPermissionListToTableFormat displays list of ACL binding rules in tabular format
func MapPermissionListToTableFormat(permissions []kafkainstanceclient.AclBinding, localizer localize.Localizer) []permissionsRow {

	rows := make([]permissionsRow, len(permissions))

	for i, p := range permissions {

		description := buildDescription(p.PatternType, localizer)
		row := permissionsRow{
			Principal:   formatPrincipal(p.GetPrincipal(), localizer),
			Permission:  fmt.Sprintf("%s | %s", p.GetPermission(), p.GetOperation()),
			Description: fmt.Sprintf("%s %s \"%s\"", p.GetResourceType(), description, p.GetResourceName()),
		}
		rows[i] = row
	}
	return rows
}

func formatPrincipal(principal string, localizer localize.Localizer) string {
	s := strings.Split(principal, ":")[1]

	if s == Wildcard {
		return localizer.MustLocalize("kafka.acl.common.allAccounts")
	}

	return s
}

func buildDescription(patternType kafkainstanceclient.AclPatternType, localizer localize.Localizer) string {
	if patternType == kafkainstanceclient.ACLPATTERNTYPE_LITERAL {
		return localizer.MustLocalize("kafka.acl.common.is")
	}

	return localizer.MustLocalize("kafka.acl.common.startsWith")
}
