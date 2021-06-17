package serviceregistry

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redhat-developer/app-services-cli/pkg/kafka/kafkaerr"
	srsmgmtv1 "github.com/redhat-developer/app-services-sdk-go/registrymgmt/apiv1/client"
)

func GetServiceRegistryByID(ctx context.Context, api srsmgmtv1.RegistriesApi, registryID string) (*srsmgmtv1.RegistryRest, *http.Response, error) {
	request := api.GetRegistry(ctx, registryID)
	registry, _, err := request.Execute()
	if err != nil {
		return nil, nil, err
	}
	return &registry, nil, err
}

func GetServiceRegistryByName(ctx context.Context, api srsmgmtv1.RegistriesApi, name string) (*srsmgmtv1.RegistryRest, *http.Response, error) {
	r := api.GetRegistries(ctx)
	r = r.Search(fmt.Sprintf("name=%v", name))
	registryList, httpResponse, err := r.Execute()
	if err != nil {
		return nil, httpResponse, err
	}

	if registryList.GetTotal() == 0 {
		return nil, nil, kafkaerr.NotFoundByNameError(name)
	}

	items := registryList.GetItems()
	kafkaReq := items[0]

	return &kafkaReq, httpResponse, err
}
