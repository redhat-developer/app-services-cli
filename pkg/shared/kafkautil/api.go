package kafkautil

import (
	"context"
	"fmt"
	"net/http"

	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
)

func GetKafkaByID(ctx context.Context, api kafkamgmtclient.DefaultApi, id string) (*kafkamgmtclient.KafkaRequest, *http.Response, error) {
	r := api.GetKafkaById(ctx, id)

	kafkaReq, httpResponse, err := r.Execute()
	if IsErr(err, ErrorCode7) {
		return nil, httpResponse, NotFoundByIDError(id)
	}

	return &kafkaReq, httpResponse, err
}

func GetKafkaByName(ctx context.Context, api kafkamgmtclient.DefaultApi, name string) (*kafkamgmtclient.KafkaRequest, *http.Response, error) {
	r := api.GetKafkas(ctx)
	r = r.Search(fmt.Sprintf("name = %v", name))
	kafkaList, httpResponse, err := r.Execute()
	if err != nil {
		return nil, httpResponse, err
	}

	if kafkaList.GetTotal() == 0 {
		return nil, nil, NotFoundByNameError(name)
	}

	items := kafkaList.GetItems()
	kafkaReq := items[0]

	return &kafkaReq, httpResponse, err
}
