package kafkautil

import (
	"context"
	"fmt"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/kafkamgmt/apiv1/client"
	"net/http"
)

func CreateClusterSearchStringFromKafkaList(kfmClusterList *kafkamgmtclient.EnterpriseClusterList) string {
	searchString := ""
	for idx, kfmcluster := range kfmClusterList.Items {
		if idx > 0 {
			searchString += " or "
		}
		searchString += fmt.Sprintf("id = '%s'", kfmcluster.Id)
	}
	return searchString
}

func ListEnterpriseClusters(f *factory.Factory) (*kafkamgmtclient.EnterpriseClusterList, *http.Response, error) {
	conn, err := f.Connection()
	if err != nil {
		return nil, nil, err
	}

	ctx := context.Background()
	api := conn.API()
	cl := api.KafkaMgmtEnterprise().GetEnterpriseOsdClusters(ctx)
	clist, response, err := cl.Execute()
	if err != nil {
		return nil, response, err
	}

	f.Logger.Debug(response)

	return &clist, response, nil
}
