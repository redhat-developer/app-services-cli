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