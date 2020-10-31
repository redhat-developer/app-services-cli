package kafka

import (
	mas "gitlab.cee.redhat.com/mas-dx/rhmas/client/mas"
)

// TODO refactor into separate config class

func BuildMasClient() *mas.APIClient {
	// TODO config abstraction
	testHost := "localhost:8000"
	testScheme := "http"
	// Based on https://github.com/OpenAPITools/openapi-generator/blob/master/samples/client/petstore/go/pet_api_test.go

	cfg := mas.NewConfiguration()
	// TODO read flag from config
	cfg.AddDefaultHeader("Authorization", "Bearer 9f4068b1c2cc720dd44dc2c6157569ae")
	cfg.Host = testHost
	cfg.Scheme = testScheme
	client := mas.NewAPIClient(cfg)

	return client
}
