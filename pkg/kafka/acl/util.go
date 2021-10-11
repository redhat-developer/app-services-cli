package acl

import (
	"net/http"

	"github.com/redhat-developer/app-services-cli/pkg/localize"
	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal/client"
)

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
