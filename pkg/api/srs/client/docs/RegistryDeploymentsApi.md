# \RegistryDeploymentsApi

All URIs are relative to *https://api.openshift.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRegistryDeployment**](RegistryDeploymentsApi.md#CreateRegistryDeployment) | **Post** /api/v1/registryDeployments | Create a registry deployment.
[**DeleteRegistryDeployment**](RegistryDeploymentsApi.md#DeleteRegistryDeployment) | **Delete** /api/v1/registryDeployments/{registryDeploymentId} | Delete a specific Registry Deployment.
[**GetRegistryDeployment**](RegistryDeploymentsApi.md#GetRegistryDeployment) | **Get** /api/v1/registryDeployments/{registryDeploymentId} | Get a specific registry deployment.
[**GetRegistryDeployments**](RegistryDeploymentsApi.md#GetRegistryDeployments) | **Get** /api/v1/registryDeployments | Get the list of all registry deployments.



## CreateRegistryDeployment

> RegistryDeployment CreateRegistryDeployment(ctx).RegistryDeploymentCreate(registryDeploymentCreate).Execute()

Create a registry deployment.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    registryDeploymentCreate := *openapiclient.NewRegistryDeploymentCreate("RegistryDeploymentUrl_example", "TenantManagerUrl_example") // RegistryDeploymentCreate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RegistryDeploymentsApi.CreateRegistryDeployment(context.Background()).RegistryDeploymentCreate(registryDeploymentCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryDeploymentsApi.CreateRegistryDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRegistryDeployment`: RegistryDeployment
    fmt.Fprintf(os.Stdout, "Response from `RegistryDeploymentsApi.CreateRegistryDeployment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRegistryDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registryDeploymentCreate** | [**RegistryDeploymentCreate**](RegistryDeploymentCreate.md) |  | 

### Return type

[**RegistryDeployment**](RegistryDeployment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRegistryDeployment

> DeleteRegistryDeployment(ctx, registryDeploymentId).Execute()

Delete a specific Registry Deployment.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    registryDeploymentId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RegistryDeploymentsApi.DeleteRegistryDeployment(context.Background(), registryDeploymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryDeploymentsApi.DeleteRegistryDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryDeploymentId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRegistryDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegistryDeployment

> RegistryDeployment GetRegistryDeployment(ctx, registryDeploymentId).Execute()

Get a specific registry deployment.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    registryDeploymentId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RegistryDeploymentsApi.GetRegistryDeployment(context.Background(), registryDeploymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryDeploymentsApi.GetRegistryDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegistryDeployment`: RegistryDeployment
    fmt.Fprintf(os.Stdout, "Response from `RegistryDeploymentsApi.GetRegistryDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryDeploymentId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegistryDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RegistryDeployment**](RegistryDeployment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegistryDeployments

> []RegistryDeployment GetRegistryDeployments(ctx).Execute()

Get the list of all registry deployments.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RegistryDeploymentsApi.GetRegistryDeployments(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryDeploymentsApi.GetRegistryDeployments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegistryDeployments`: []RegistryDeployment
    fmt.Fprintf(os.Stdout, "Response from `RegistryDeploymentsApi.GetRegistryDeployments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegistryDeploymentsRequest struct via the builder pattern


### Return type

[**[]RegistryDeployment**](RegistryDeployment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

