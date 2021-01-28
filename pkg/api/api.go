package api

import (
	serviceapiclient "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/serviceapi/client"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/strimzi-admin/client"
)

// API is a type which defines a number of APIs
type API struct {
	Kafka        serviceapiclient.DefaultApi
	StrimziAdmin strimziadminclient.DefaultApi
}
