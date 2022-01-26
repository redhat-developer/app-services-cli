package remote

import (
	"context"
	_ "embed"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/redhat-developer/app-services-cli/internal/build"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
)

//go:embed service-constants.json
var serviceConstants []byte

// Fetch service constants that can be reused by multiple commands
func GetRemoteServiceConstants(context context.Context, logger logging.Logger) (error, *DynamicServiceConstants) {
	var embeddedConstants DynamicServiceConstants
	err := json.Unmarshal(serviceConstants, &embeddedConstants)
	if err != nil {
		return errors.New("unable to unmarshal embedded service constants"), nil
	}

	if build.DynamicConfigURL == "" {
		return nil, &embeddedConstants
	}

	client := &http.Client{}
	req, err := http.NewRequestWithContext(context, http.MethodGet, build.DynamicConfigURL, nil)
	if err != nil {
		logger.Debug("Fetching remote constants failed with error", err)
		return nil, &embeddedConstants
	}

	response, err := client.Do(req)
	if err != nil || response == nil {
		logger.Debug("Fetching remote constants failed with error ", err)
		return nil, &embeddedConstants
	}
	defer response.Body.Close()

	specJson, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Debug("Reading remote constants failed with error ", err)
		return nil, &embeddedConstants
	}

	logger.Debug("Service Constants: ", string(specJson))

	var dynamicServiceConstants DynamicServiceConstants
	err = json.Unmarshal([]byte(specJson), &dynamicServiceConstants)
	if err != nil {
		logger.Debug("Parsing remote constants failed with error ", err)
		return err, &embeddedConstants
	}

	return nil, &dynamicServiceConstants
}
