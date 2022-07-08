package connectorutil

import (
	"context"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	connectormgmtclient "github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1/client"
)

func GetConnectorByID(ctx context.Context, api *connectormgmtclient.APIClient, id string) (*connectormgmtclient.Connector, error) {
	connectorInstance, _, err := api.ConnectorsApi.GetConnector(ctx, id).Execute()
	if err != nil {
		return nil, err
	}

	return &connectorInstance, nil
}

func GetConnectorByName(ctx context.Context, api *connectormgmtclient.APIClient, name string, localizer localize.Localizer) (*connectormgmtclient.Connector, error) {
	list, _, err := api.ConnectorsApi.ListConnectors(ctx).Execute()
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(list.Items); i++ {
		if list.Items[i].Name == name {
			return &list.Items[i], nil
		}
	}

	return nil, localizer.MustLocalizeError("connector.common.error.nameNotFound", localize.NewEntry("Name", name))
}
