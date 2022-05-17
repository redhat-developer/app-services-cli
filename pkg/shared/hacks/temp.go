package hacks

// Temporary hack package
// Nothing to see here

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
)

// Temporary hack that we use to determine if
// Our CLI needs to use mas-sso token
func ShouldUseMasSSO(logger logging.Logger, apiUrl string) bool {
	req, err := http.NewRequest("GET", apiUrl+"/api/kafkas_mgmt/v1/sso_providers", nil)
	if err != nil {
		logger.Debug("Error when fetching auth config", err)
		return true
	}

	req = req.WithContext(context.Background())

	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Debug("Error when fetching auth config", err)
		return true
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Debug("Error when fetching auth config", err)
		return true
	}

	response := string(b)

	// defining a struct instance
	var provider *kafkamgmtclient.SsoProvider

	responseBytes := []byte(fmt.Sprintf("%v", response))
	err = json.Unmarshal(responseBytes, &provider)
	if err != nil {
		logger.Debug("Error when fetching auth config", err)
		return true
	}

	if provider.GetBaseUrl() == "" {
		logger.Debug("Error when fetching auth config", err)
		return true
	}

	return provider.GetName() == "mas_sso"
}
