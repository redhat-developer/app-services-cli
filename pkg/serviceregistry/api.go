package serviceregistry

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	srsclient "github.com/redhat-developer/app-services-cli/pkg/api/srs/client"
)

func GetServiceRegistryByID(ctx context.Context, api srsclient.DefaultApi, registryID string) (*srsclient.Registry, *http.Response, error) {
	rgInt, _ := strconv.Atoi(registryID)
	registryInt := int32(rgInt)

	request := api.GetRegistry(ctx, registryInt)
	registry, _, err := request.Execute()
	if err != nil {
		return nil, nil, err
	}
	return &registry, nil, err
}

func GetServiceRegistryByName(ctx context.Context, api srsclient.DefaultApi, name string) (*srsclient.Registry, *http.Response, error) {
	return nil, nil, errors.New("Not implemented. Only ID supported because of missing backend feature")
}
