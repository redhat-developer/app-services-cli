package connectorcmdutil

import (
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/spf13/cobra"
)

func FilterValidTypesArgs(f *factory.Factory, toComplete string) ([]string, cobra.ShellCompDirective) {
	validTypes := []string{}
	directive := cobra.ShellCompDirectiveNoSpace
	
	svcContext, err := f.ServiceContext.Load()
	if err != nil {
		return validTypes, directive
	}

	currCtx, err := contextutil.GetCurrentContext(svcContext, f.Localizer)
	if err != nil {
		return validTypes, directive
	}

	instanceID := currCtx.KafkaID
	if instanceID == "" {
		return validTypes, directive
	}

	conn, err := f.Connection()
	if err != nil {
		return validTypes, directive
	}

	api := conn.API().ConnectorsMgmt().ConnectorTypesApi
	if err != nil {
		return validTypes, directive
	}

	req := api.GetConnectorTypes(f.Context)

	req = req.Size("500")
	req = req.Page("1")

	types, httpRes, err := req.Execute()
	if httpRes != nil {
		defer func() {
			_ = httpRes.Body.Close()
		}()
	}
	if err != nil {
		return validTypes, directive
	}

	items := types.GetItems()
	for _, connector_type := range items {
		validTypes = append(validTypes, connector_type.GetId())
	}

	return validTypes, directive
}