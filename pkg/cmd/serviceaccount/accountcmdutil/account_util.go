package accountcmdutil

import "github.com/redhat-developer/app-services-cli/pkg/cmd/serviceaccount/accountcmdutil/credentials"

var (
	CredentialsOutputFormats = []string{credentials.EnvFormat, credentials.JSONFormat, credentials.PropertiesFormat}
)
