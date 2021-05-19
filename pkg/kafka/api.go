package kafka

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redhat-developer/app-services-cli/pkg/api/kas"
	"github.com/redhat-developer/app-services-cli/pkg/kafka/kafkaerr"
	kafkamgmtv1 "github.com/redhat-developer/app-services-sdk-go/kafka/mgmt/apiv1"
)

func GetKafkaByID(ctx context.Context, api kafkamgmtv1.DefaultApi, id string) (*kafkamgmtv1.KafkaRequest, *http.Response, error) {
	r := api.GetKafkaById(ctx, id)

	kafkaReq, httpResponse, err := r.Execute()
	if kas.IsErr(err, kas.ErrorNotFound) {
		return nil, httpResponse, kafkaerr.NotFoundByIDError(id)
	}

	return &kafkaReq, httpResponse, err
}

func GetKafkaByName(ctx context.Context, api kafkamgmtv1.DefaultApi, name string) (*kafkamgmtv1.KafkaRequest, *http.Response, error) {
	r := api.ListKafkas(ctx)
	r = r.Search(fmt.Sprintf("name = %v", name))
	kafkaList, httpResponse, err := r.Execute()
	if err != nil {
		return nil, httpResponse, err
	}

	if kafkaList.GetTotal() == 0 {
		return nil, nil, kafkaerr.NotFoundByNameError(name)
	}

	items := kafkaList.GetItems()
	kafkaReq := items[0]

	return &kafkaReq, httpResponse, err
}
