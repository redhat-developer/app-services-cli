package svcaccountcmdutil

import "github.com/redhat-developer/app-services-cli/pkg/cmd/serviceaccount/svcaccountcmdutil/credentials"

var (
	CredentialsOutputFormats = []string{credentials.EnvFormat, credentials.JSONFormat, credentials.PropertiesFormat}
)
