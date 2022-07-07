package connectorutil

import (
	"context"
	connectormgmtclient "github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1/client"
)

func GetConnectorByID(ctx context.Context, api connectormgmtclient.APIClient, id string) (*connectormgmtclient.Connector, error) {
	connectorInstance, _, err := api.ConnectorsApi.GetConnector(ctx, id).Execute()
	if err != nil {
		return nil, err
	}

	return &connectorInstance, nil
}

func GetConnectorByName(ctx context.Context, api connectormgmtclient.APIClient, name string) (*connectormgmtclient.Connector, error) {
	list, _, err := api.ConnectorsApi.ListConnectors(ctx).Execute()
	if err != nil {
		return nil, err
	}

	for i, connector := range list.Items {
		if connector.Name == name {
			return &list.Items[i], nil
		}
	}

	return nil, NotFoundByNameError(name)
}
