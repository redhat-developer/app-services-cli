package svcaccountcmdutil

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/serviceaccount/svcaccountcmdutil/credentials"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
)

var (
	CredentialsOutputFormats = []string{credentials.EnvFormat, credentials.JSONFormat, credentials.PropertiesFormat, credentials.SecretFormat}
)

// Method fetches authentication details for providers
func GetProvidersDetails(conn connection.Connection, context context.Context) (*kafkamgmtclient.SsoProvider, error) {
	providers, httpRes, err := conn.API().
		ServiceAccountMgmt().GetSsoProviders(context).Execute()

	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	return &providers, err
}
