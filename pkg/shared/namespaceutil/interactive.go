package namespaceutil

import (
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"

	"github.com/AlecAivazis/survey/v2"
	connectormgmtclient "github.com/jackdelahunt/app-services-sdk-core/app-services-sdk-go/connectormgmt/apiv1/client"
)

func InteractiveSelect(connection connection.Connection, f *factory.Factory) (*connectormgmtclient.ConnectorNamespace, error) {
	api := connection.API().ConnectorsMgmt()

	list, _, err := api.ConnectorNamespacesApi.ListConnectorNamespaces(f.Context).Execute()
	if err != nil {
		return nil, err
	}

	if len(list.Items) == 0 {
		return nil, f.Localizer.MustLocalizeError("connector.error.interactive.noNamespaces")
	}

	namespaceNames := make([]string, len(list.Items))
	for index := 0; index < len(list.Items); index++ {
		namespaceNames[index] = list.Items[index].Name
	}

	prompt := &survey.Select{
		Message:  f.Localizer.MustLocalize("connector.common.input.instanceName.message"),
		Options:  namespaceNames,
		PageSize: 10,
	}

	var selectedIndex int
	err = survey.AskOne(prompt, &selectedIndex)
	if err != nil {
		return nil, err
	}

	return &list.Items[selectedIndex], nil
}
