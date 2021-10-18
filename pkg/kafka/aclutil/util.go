package aclutil

import (
	"fmt"
	"net/http"

	"github.com/redhat-developer/app-services-cli/pkg/localize"
	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal/client"
)

// When the value of the `--topic`, `--group`, `user` or `service-account` option is one of
// the keys of this map, it will be replaced by the corresponding value.
var commonArgAliases = map[string]string{
	"all": Wildcard,
}

// ExecuteACLRuleCreate makes request to create an ACL rule
func ExecuteACLRuleCreate(req kafkainstanceclient.ApiCreateAclRequest, localizer localize.Localizer, kafkaInstanceName string) error {

	httpRes, err := req.Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	return ValidateAPIError(httpRes, localizer, err, "create", kafkaInstanceName)
}

// FormatPrincipal formats the provided principal ID to "User:principal"
func FormatPrincipal(userID string) string {
	return fmt.Sprintf("User:%s", userID)
}

// GetResourceName returns the name of the resource
// transformed into a server recognized format
func GetResourceName(resourceName string) string {
	if commonArgAliases[resourceName] == Wildcard {
		return Wildcard
	}
	return resourceName
}

// IsValidResourceOperation checks if the operation is valid, and returns the list valid operations when invalid
func IsValidResourceOperation(resourceType string, operation string, resourceOperationsMap map[string][]string) (bool, []string) {
	resourceTypeMapped := resourceTypeOperationKeyMap[resourceType]
	resourceOperations := resourceOperationsMap[resourceTypeMapped]

	for i, op := range resourceOperations {
		if operationMapped, ok := validOperationsResponseMap[op]; ok {
			resourceOperations[i] = operationMapped
		} else {
			resourceOperations[i] = op
		}
		if resourceOperations[i] == operation {
			return true, nil
		}
	}

	return false, resourceOperations
}

// ValidateAPIError checks for a HTTP error and maps it to a user friendly error
func ValidateAPIError(httpRes *http.Response, localizer localize.Localizer, err error, operation string, instanceName string) error {
	if err == nil {
		return nil
	}

	if httpRes == nil {
		return err
	}

	operationTmplPair := localize.NewEntry("Operation", operation)

	switch httpRes.StatusCode {
	case http.StatusUnauthorized:
		return localizer.MustLocalizeError("kafka.acl.common.error.unauthorized", operationTmplPair)
	case http.StatusForbidden:
		return localizer.MustLocalizeError("kafka.acl.common.error.forbidden", operationTmplPair)
	case http.StatusInternalServerError:
		return localizer.MustLocalizeError("kafka.acl.common.error.internalServerError")
	case http.StatusServiceUnavailable:
		return localizer.MustLocalizeError("kafka.acl.common.error.unableToConnectToKafka", localize.NewEntry("Name", instanceName))
	default:
		return err
	}
}
