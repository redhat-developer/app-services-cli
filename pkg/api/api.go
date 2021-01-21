package api

import (
	msclient "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/managedservices/client"
)

type API struct {
	Kafka msclient.DefaultApi
}
