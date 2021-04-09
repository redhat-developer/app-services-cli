package kafka

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redhat-developer/app-services-cli/pkg/api/kas"
	kasclient "github.com/redhat-developer/app-services-cli/pkg/api/kas/client"
)

func GetKafkaByID(ctx context.Context, api kasclient.DefaultApi, id string) (*kasclient.KafkaRequest, *http.Response, error) {
	r := api.GetKafkaById(ctx, id)

	kafkaReq, httpResponse, apiErr := r.Execute()
	if kas.IsErr(apiErr, kas.ErrorNotFound) {
		return nil, httpResponse, ErrorNotFound(id)
	}

	return &kafkaReq, httpResponse, apiErr
}

func GetKafkaByName(ctx context.Context, api kasclient.DefaultApi, name string) (*kasclient.KafkaRequest, *http.Response, error) {
	r := api.ListKafkas(ctx)
	r = r.Search(fmt.Sprintf("name = %v", name))
	kafkaList, httpResponse, apiErr := r.Execute()
	if apiErr.Error() != "" {
		return nil, httpResponse, apiErr
	}

	if kafkaList.GetTotal() == 0 {
		return nil, nil, ErrorNotFoundByName(name)
	}

	items := kafkaList.GetItems()
	kafkaReq := items[0]

	return &kafkaReq, httpResponse, apiErr
}
