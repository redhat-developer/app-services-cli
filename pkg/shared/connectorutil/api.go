package connectorutil

import (
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	connectormgmtclient "github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1/client"
)

func GetConnectorByID(api *connectormgmtclient.APIClient, id string, f *factory.Factory) (*connectormgmtclient.Connector, error) {
	connectorInstance, _, err := api.ConnectorsApi.GetConnector(f.Context, id).Execute()
	if err != nil {
		return nil, err
	}

	return &connectorInstance, nil
}

func GetConnectorByName(api *connectormgmtclient.APIClient, name string, f *factory.Factory) (*connectormgmtclient.Connector, error) {
	list, _, err := api.ConnectorsApi.ListConnectors(f.Context).Execute()
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(list.Items); i++ {
		if list.Items[i].Name == name {
			return &list.Items[i], nil
		}
	}

	return nil, f.Localizer.MustLocalizeError("connector.common.error.nameNotFound", localize.NewEntry("Name", name))
}
