package kafkautil

import (
	"context"
	"fmt"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/kafkamgmt/apiv1/client"
	"net/http"
)

const (
	statusClusterAccepted         = "cluster_accepted"
	statusClusterProvisioning     = "cluster_provisioning"
	statusClusterProvisioned      = "cluster_provisioned"
	statusWaitingForKasFleetshard = "waiting_for_kas_fleetshard_operator"
	statusReady                   = "ready"
	statusFailed                  = "failed"
	statusDeprovisioning          = "deprovisioning"
	statusCleanup                 = "cleanup"

	customStatusAccepted      = "Cluster Accepted"
	customStatusRegistering   = "Registering"
	customStatusUnregistering = "Unregistering"
	customStatusReady         = "Ready"
	customStatusFailed        = "Failed"
)

func MapClusterStatus(clusterStatus string) string {
	switch clusterStatus {
	case statusClusterAccepted:
		return customStatusAccepted
	case statusClusterProvisioning:
	case statusClusterProvisioned:
	case statusWaitingForKasFleetshard:
		return customStatusRegistering
	case statusCleanup:
	case statusDeprovisioning:
		return customStatusUnregistering
	case statusReady:
		return customStatusReady
	case statusFailed:
		return customStatusFailed
	}

	return clusterStatus
}

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
