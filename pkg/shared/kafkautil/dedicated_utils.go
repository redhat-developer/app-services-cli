package kafkautil

import (
	"context"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/kafkamgmt/apiv1/client"
)

func ListEnterpriseClusters(f *factory.Factory) (*kafkamgmtclient.EnterpriseClusterList, error) {
	conn, err := f.Connection()
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	api := conn.API()
	cl := api.KafkaMgmtEnterprise().GetEnterpriseOsdClusters(ctx)
	clist, response, err := cl.Execute()
	if err != nil {
		return nil, err
	}
	if len(clist.Items) == 0 {
		return &clist, nil
	}
	f.Logger.Debug(response)
	return &clist, nil
}
