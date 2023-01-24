package namespaceutil

import (
	connectormgmtclient "github.com/jackdelahunt/app-services-sdk-core/app-services-sdk-go/connectormgmt/apiv1/client"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
)

func GetNamespaceByID(api *connectormgmtclient.APIClient, id string, f *factory.Factory) (*connectormgmtclient.ConnectorNamespace, error) {
	namespaceInstance, _, err := api.ConnectorNamespacesApi.GetConnectorNamespace(f.Context, id).Execute()
	if err != nil {
		return nil, err
	}

	return &namespaceInstance, nil
}

func GetNamespaceByName(api *connectormgmtclient.APIClient, name string, f *factory.Factory) (*connectormgmtclient.ConnectorNamespace, error) {
	list, _, err := api.ConnectorNamespacesApi.ListConnectorNamespaces(f.Context).Execute()
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(list.Items); i++ {
		if list.Items[i].GetName() == name {
			return &list.Items[i], nil
		}
	}

	return nil, f.Localizer.MustLocalizeError("namespace.common.error.nameNotFound", localize.NewEntry("Name", name))
}
