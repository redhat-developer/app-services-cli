package ams

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/redhat-developer/app-services-cli/internal/build"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
)

// Contains specification for terms and condition parameters
// NOTE: Before updating this fallback file
// Please update source at https://github.com/redhat-developer/app-services-ui/blob/main/static/configs/terms-conditions-spec.json
var fallbackTocSpec = TermsAndConditionsSpec{
	Kafka: ServiceTermsSpec{
		EventCode: "register",
		SiteCode:  "ocm",
	},
	ServiceRegistry: ServiceTermsSpec{
		EventCode: "onlineService",
		SiteCode:  "ocm",
	},
}

// GetRemoteTermsSpec fetch event and site code information associated with the services
// Function is used to dynamically download new terms and conditions specifications
// without forcing end users to update their CLI.
func GetRemoteTermsSpec(logger logging.Logger) TermsAndConditionsSpec {

	response, err := http.Get(build.TermsReviewSpecURL)

	if err != nil || response.Body == nil {
		logger.Debug("Fetching remote terms failed with error ", err)
		return fallbackTocSpec
	}
	defer response.Body.Close()

	var specJson []byte
	specJson, err = ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Debug("Reading remote terms failed with error ", err)
		return fallbackTocSpec
	}

	var termsAndConditionsSpec TermsAndConditionsSpec
	err = json.Unmarshal([]byte(specJson), &termsAndConditionsSpec)
	if err != nil {
		logger.Debug("Parsing remote terms failed with error ", err)
		return fallbackTocSpec
	}
	return termsAndConditionsSpec
}
