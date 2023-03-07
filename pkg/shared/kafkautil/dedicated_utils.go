package kafkautil

import (
	"context"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/kafkamgmt/apiv1/client"
	"net/http"
)

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
