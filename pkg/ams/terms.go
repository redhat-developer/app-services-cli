package ams

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/redhat-developer/app-services-cli/internal/build"
)

var fallbackTocSpec = []byte(`{
    "kafka":{
       "EventCode":"register",
       "SiteCode":"ocm",
       "StopOnTermsChange": true
    },
    "service-registry":{
       "EventCode":"onlineService",
       "SiteCode":"ocm",
       "StopOnTermsChange": true
    }
  }
`)

func GetRemoteTermsSpec() TermsAndConditionsSpec {
	response, err := http.Get(build.TermsReviewSpecURL)

	var specJson []byte
	if err != nil {
		// TODO log error?
		specJson = fallbackTocSpec
	} else {
		specJson, err = ioutil.ReadAll(response.Body)
		if err != nil {
			// TODO log error?
			specJson = fallbackTocSpec
		}
	}

	var termsAndConditionsSpec TermsAndConditionsSpec
	json.Unmarshal([]byte(specJson), &termsAndConditionsSpec)
	return termsAndConditionsSpec
}
