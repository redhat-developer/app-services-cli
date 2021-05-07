package kafka

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redhat-developer/app-services-cli/pkg/api/kas"
	kasclient "github.com/redhat-developer/app-services-cli/pkg/api/kas/client"
	"github.com/redhat-developer/app-services-cli/pkg/kafka/kafkaerr"
)

func GetKafkaByID(ctx context.Context, api kasclient.DefaultApi, id string) (*kasclient.KafkaRequest, *http.Response, error) {
	r := api.GetKafkaById(ctx, id)

	kafkaReq, httpResponse, err := r.Execute()
	if kas.IsErr(err, kas.ErrorNotFound) {
		return nil, httpResponse, kafkaerr.NotFoundByIDError(id)
	}

	return &kafkaReq, httpResponse, err
}

func GetKafkaByName(ctx context.Context, api kasclient.DefaultApi, name string) (*kasclient.KafkaRequest, *http.Response, error) {
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
