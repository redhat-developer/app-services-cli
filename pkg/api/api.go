package api

import (
	kasclient "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/kas/client"
	strimziadminclient "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/strimzi-admin/client"
)

// API is a type which defines a number of APIs
type API struct {
	Kafka        kasclient.DefaultApi
	StrimziAdmin strimziadminclient.DefaultApi
}
