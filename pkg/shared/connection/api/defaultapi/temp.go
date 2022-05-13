package defaultapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
)

// Temporary hack that we use to determine if
// Our CLI needs to use mas-sso token
func ShouldUseMasSSO() bool {
	req, err := http.NewRequest("GET", "https://api.openshift.com/api/kafkas_mgmt/v1/sso_providers", nil)
	if err != nil {
		return true
	}

	req = req.WithContext(context.Background())

	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return true
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return true
	}

	response := string(b)

	// defining a struct instance
	var provider *kafkamgmtclient.SsoProvider

	responseBytes := []byte(fmt.Sprintf("%v", response))
	err = json.Unmarshal(responseBytes, &provider)
	if err != nil {
		return true
	}

	return provider.GetBaseUrl() == "https://identity.api.redhat.com"
}
