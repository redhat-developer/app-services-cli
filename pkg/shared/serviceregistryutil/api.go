package serviceregistryutil

import (
	"context"
	"fmt"
	"net/http"

	srsmgmtv1 "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registrymgmt/apiv1/client"
	srsmgmtv1errors "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registrymgmt/apiv1/error"
)

func GetServiceRegistryByID(ctx context.Context, api srsmgmtv1.RegistriesApi, registryID string) (*srsmgmtv1.Registry, *http.Response, error) {
	request := api.GetRegistry(ctx, registryID)
	registry, _, err := request.Execute()

	if srsmgmtv1errors.IsAPIError(err, srsmgmtv1errors.ERROR_2) {
		return nil, nil, NotFoundByIDError(registryID)
	}

	return &registry, nil, err
}

func GetServiceRegistryByName(ctx context.Context, api srsmgmtv1.RegistriesApi, name string) (*srsmgmtv1.Registry, *http.Response, error) {
	r := api.GetRegistries(ctx)
	r = r.Search(fmt.Sprintf("name=%v", name))
	registryList, httpResponse, err := r.Execute()
	if registryList.GetTotal() == 0 {
		return nil, nil, NotFoundByNameError(name)
	}

	items := registryList.GetItems()
	registryReq := items[0]

	return &registryReq, httpResponse, err
}
