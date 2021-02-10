# \DefaultApi

All URIs are relative to *https://api.openshift.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnector**](DefaultApi.md#CreateConnector) | **Post** /api/managed-services-api/v1/kafkas/{id}/connector-deployments | Create a new connector
[**CreateKafka**](DefaultApi.md#CreateKafka) | **Post** /api/managed-services-api/v1/kafkas | Create a new kafka Request
[**CreateServiceAccount**](DefaultApi.md#CreateServiceAccount) | **Post** /api/managed-services-api/v1/serviceaccounts | Create a service account
[**DeleteConnector**](DefaultApi.md#DeleteConnector) | **Delete** /api/managed-services-api/v1/kafkas/{id}/connector-deployments/{cid} | Delete a connector
[**DeleteKafkaById**](DefaultApi.md#DeleteKafkaById) | **Delete** /api/managed-services-api/v1/kafkas/{id} | Delete a kafka request by id
[**DeleteServiceAccount**](DefaultApi.md#DeleteServiceAccount) | **Delete** /api/managed-services-api/v1/serviceaccounts/{id} | Delete service account
[**GetConnector**](DefaultApi.md#GetConnector) | **Get** /api/managed-services-api/v1/kafkas/{id}/connector-deployments/{cid} | Get a connector deployment
[**GetConnectorTypeByID**](DefaultApi.md#GetConnectorTypeByID) | **Get** /api/managed-services-api/v1/connector-types/{id} | Get a connector type by name and version
[**GetKafkaById**](DefaultApi.md#GetKafkaById) | **Get** /api/managed-services-api/v1/kafkas/{id} | Get a kafka request by id
[**GetMetricsByKafkaId**](DefaultApi.md#GetMetricsByKafkaId) | **Get** /api/managed-services-api/v1/kafkas/{id}/metrics | Get metrics by kafka id.
[**GetServiceAccountById**](DefaultApi.md#GetServiceAccountById) | **Get** /api/managed-services-api/v1/serviceaccounts/{id} | get service account by id
[**ListCloudProviderRegions**](DefaultApi.md#ListCloudProviderRegions) | **Get** /api/managed-services-api/v1/cloud_providers/{id}/regions | Retrieves the list of supported regions of the supported cloud provider.
[**ListCloudProviders**](DefaultApi.md#ListCloudProviders) | **Get** /api/managed-services-api/v1/cloud_providers | Retrieves the list of supported cloud providers.
[**ListConnectorTypes**](DefaultApi.md#ListConnectorTypes) | **Get** /api/managed-services-api/v1/connector-types | Returns a list of connector types
[**ListConnectors**](DefaultApi.md#ListConnectors) | **Get** /api/managed-services-api/v1/kafkas/{id}/connector-deployments | Returns a list of connector types
[**ListKafkas**](DefaultApi.md#ListKafkas) | **Get** /api/managed-services-api/v1/kafkas | Returns a list of Kafka requests
[**ListServiceAccounts**](DefaultApi.md#ListServiceAccounts) | **Get** /api/managed-services-api/v1/serviceaccounts | List service accounts
[**ResetServiceAccountCreds**](DefaultApi.md#ResetServiceAccountCreds) | **Post** /api/managed-services-api/v1/serviceaccounts/{id}/reset-credentials | reset credentials for the service account



## CreateConnector

> Connector CreateConnector(ctx, id).Async(async).Connector(connector).Execute()

Create a new connector

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
    id := "id_example" // string | The id of record
    async := true // bool | Perform the action in an asynchronous manner
    connector := *openapiclient.NewConnector() // Connector | Connector data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateConnector(context.Background(), id).Async(async).Connector(connector).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateConnector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConnector`: Connector
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **async** | **bool** | Perform the action in an asynchronous manner | 
 **connector** | [**Connector**](Connector.md) | Connector data | 

### Return type

[**Connector**](Connector.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateKafka

> KafkaRequest CreateKafka(ctx).Async(async).KafkaRequestPayload(kafkaRequestPayload).Execute()

Create a new kafka Request

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
    async := true // bool | Perform the action in an asynchronous manner
    kafkaRequestPayload := *openapiclient.NewKafkaRequestPayload("Name_example") // KafkaRequestPayload | Kafka data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateKafka(context.Background()).Async(async).KafkaRequestPayload(kafkaRequestPayload).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateKafka``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateKafka`: KafkaRequest
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateKafka`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateKafkaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **async** | **bool** | Perform the action in an asynchronous manner | 
 **kafkaRequestPayload** | [**KafkaRequestPayload**](KafkaRequestPayload.md) | Kafka data | 

### Return type

[**KafkaRequest**](KafkaRequest.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateServiceAccount

> ServiceAccount CreateServiceAccount(ctx).ServiceAccountRequest(serviceAccountRequest).Execute()

Create a service account

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
    serviceAccountRequest := *openapiclient.NewServiceAccountRequest("Name_example") // ServiceAccountRequest | service account request

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateServiceAccount(context.Background()).ServiceAccountRequest(serviceAccountRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateServiceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateServiceAccount`: ServiceAccount
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateServiceAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceAccountRequest** | [**ServiceAccountRequest**](ServiceAccountRequest.md) | service account request | 

### Return type

[**ServiceAccount**](ServiceAccount.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConnector

> Error DeleteConnector(ctx, id).Execute()

Delete a connector

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
    id := "id_example" // string | The id of record

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteConnector(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteConnector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteConnector`: Error
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Error**](Error.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKafkaById

> Error DeleteKafkaById(ctx, id).Async(async).Execute()

Delete a kafka request by id

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
    id := "id_example" // string | The id of record
    async := true // bool | Perform the action in an asynchronous manner

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteKafkaById(context.Background(), id).Async(async).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteKafkaById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteKafkaById`: Error
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteKafkaById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKafkaByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **async** | **bool** | Perform the action in an asynchronous manner | 

### Return type

[**Error**](Error.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServiceAccount

> Error DeleteServiceAccount(ctx, id).Execute()

Delete service account

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
    id := "id_example" // string | The id of record

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteServiceAccount(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteServiceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteServiceAccount`: Error
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteServiceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Error**](Error.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnector

> Connector GetConnector(ctx, id, cid).Execute()

Get a connector deployment

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
    id := "id_example" // string | The id of record
    cid := "cid_example" // string | The id of the connector

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetConnector(context.Background(), id, cid).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetConnector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnector`: Connector
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 
**cid** | **string** | The id of the connector | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Connector**](Connector.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectorTypeByID

> ConnectorType GetConnectorTypeByID(ctx, id).Execute()

Get a connector type by name and version

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
    id := "id_example" // string | The id of record

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetConnectorTypeByID(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetConnectorTypeByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnectorTypeByID`: ConnectorType
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetConnectorTypeByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorTypeByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConnectorType**](ConnectorType.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaById

> KafkaRequest GetKafkaById(ctx, id).Execute()

Get a kafka request by id

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
    id := "id_example" // string | The id of record

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetKafkaById(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetKafkaById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaById`: KafkaRequest
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetKafkaById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKafkaByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KafkaRequest**](KafkaRequest.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricsByKafkaId

> MetricsList GetMetricsByKafkaId(ctx, id).Duration(duration).Interval(interval).Filters(filters).Execute()

Get metrics by kafka id.

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
    id := "id_example" // string | The id of record
    duration := int64(5) // int64 | The length of time in minutes over which to return the metrics. (default to 5)
    interval := int64(30) // int64 | The interval in seconds between data points. (default to 30)
    filters := []string{"Inner_example"} // []string | List of metrics to fetch. Fetch all metrics when empty. List entries are kafka internal metric names. (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetMetricsByKafkaId(context.Background(), id).Duration(duration).Interval(interval).Filters(filters).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetMetricsByKafkaId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetricsByKafkaId`: MetricsList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetMetricsByKafkaId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricsByKafkaIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **duration** | **int64** | The length of time in minutes over which to return the metrics. | [default to 5]
 **interval** | **int64** | The interval in seconds between data points. | [default to 30]
 **filters** | **[]string** | List of metrics to fetch. Fetch all metrics when empty. List entries are kafka internal metric names. | [default to []]

### Return type

[**MetricsList**](MetricsList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceAccountById

> ServiceAccount GetServiceAccountById(ctx, id).Execute()

get service account by id

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
    id := "id_example" // string | The id of record

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetServiceAccountById(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetServiceAccountById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceAccountById`: ServiceAccount
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetServiceAccountById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceAccountByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceAccount**](ServiceAccount.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCloudProviderRegions

> CloudRegionList ListCloudProviderRegions(ctx, id).Page(page).Size(size).Execute()

Retrieves the list of supported regions of the supported cloud provider.

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
    id := "id_example" // string | The id of record
    page := "1" // string | Page index (optional)
    size := "100" // string | Number of items in each page (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListCloudProviderRegions(context.Background(), id).Page(page).Size(size).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListCloudProviderRegions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCloudProviderRegions`: CloudRegionList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListCloudProviderRegions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCloudProviderRegionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | Page index | 
 **size** | **string** | Number of items in each page | 

### Return type

[**CloudRegionList**](CloudRegionList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCloudProviders

> CloudProviderList ListCloudProviders(ctx).Page(page).Size(size).Execute()

Retrieves the list of supported cloud providers.

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
    page := "1" // string | Page index (optional)
    size := "100" // string | Number of items in each page (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListCloudProviders(context.Background()).Page(page).Size(size).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListCloudProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCloudProviders`: CloudProviderList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListCloudProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCloudProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string** | Page index | 
 **size** | **string** | Number of items in each page | 

### Return type

[**CloudProviderList**](CloudProviderList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectorTypes

> ConnectorTypeList ListConnectorTypes(ctx).Page(page).Size(size).Execute()

Returns a list of connector types

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
    page := "1" // string | Page index (optional)
    size := "100" // string | Number of items in each page (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListConnectorTypes(context.Background()).Page(page).Size(size).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListConnectorTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConnectorTypes`: ConnectorTypeList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListConnectorTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListConnectorTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string** | Page index | 
 **size** | **string** | Number of items in each page | 

### Return type

[**ConnectorTypeList**](ConnectorTypeList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectors

> ConnectorList ListConnectors(ctx, id).Page(page).Size(size).Execute()

Returns a list of connector types

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
    id := "id_example" // string | The id of record
    page := "1" // string | Page index (optional)
    size := "100" // string | Number of items in each page (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListConnectors(context.Background(), id).Page(page).Size(size).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListConnectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConnectors`: ConnectorList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListConnectors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | Page index | 
 **size** | **string** | Number of items in each page | 

### Return type

[**ConnectorList**](ConnectorList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkas

> KafkaRequestList ListKafkas(ctx).Page(page).Size(size).OrderBy(orderBy).Search(search).Execute()

Returns a list of Kafka requests

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
    page := "1" // string | Page index (optional)
    size := "100" // string | Number of items in each page (optional)
    orderBy := "name asc" // string | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement. Each query can be ordered by any of the kafkaRequests fields. For example, in order to retrieve all kafkas ordered by their name:  ```sql name asc ```  Or in order to retrieve all kafkas ordered by their name _and_ created date:  ```sql name asc, created_at asc ```  If the parameter isn't provided, or if the value is empty, then the results will be ordered by name. (optional)
    search := "name = my-kafka and cloud_provider = aws" // string | Search criteria.  The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement. Allowed fields in the search are: cloud_provider, name, owner, region and status. Allowed comparators are `<>`, `=` or `LIKE`. Allowed joins are `AND` and `OR`, however there is a limit of max 10 joins in the search query.  Examples:  To retrieve kafka request with name equal `my-kafka` and region equal `aws`, the value should be:  ``` name = my-kafka and cloud_provider = aws ```  To retrieve kafka request with its name starting with `my`, the value should be:  ``` name like my%25 ```  If the parameter isn't provided, or if the value is empty, then all the kafkas that the user has permission to see will be returned.  Note. If the query is invalid, an error will be returned  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListKafkas(context.Background()).Page(page).Size(size).OrderBy(orderBy).Search(search).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListKafkas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKafkas`: KafkaRequestList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListKafkas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListKafkasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string** | Page index | 
 **size** | **string** | Number of items in each page | 
 **orderBy** | **string** | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement. Each query can be ordered by any of the kafkaRequests fields. For example, in order to retrieve all kafkas ordered by their name:  &#x60;&#x60;&#x60;sql name asc &#x60;&#x60;&#x60;  Or in order to retrieve all kafkas ordered by their name _and_ created date:  &#x60;&#x60;&#x60;sql name asc, created_at asc &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then the results will be ordered by name. | 
 **search** | **string** | Search criteria.  The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement. Allowed fields in the search are: cloud_provider, name, owner, region and status. Allowed comparators are &#x60;&lt;&gt;&#x60;, &#x60;&#x3D;&#x60; or &#x60;LIKE&#x60;. Allowed joins are &#x60;AND&#x60; and &#x60;OR&#x60;, however there is a limit of max 10 joins in the search query.  Examples:  To retrieve kafka request with name equal &#x60;my-kafka&#x60; and region equal &#x60;aws&#x60;, the value should be:  &#x60;&#x60;&#x60; name &#x3D; my-kafka and cloud_provider &#x3D; aws &#x60;&#x60;&#x60;  To retrieve kafka request with its name starting with &#x60;my&#x60;, the value should be:  &#x60;&#x60;&#x60; name like my%25 &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then all the kafkas that the user has permission to see will be returned.  Note. If the query is invalid, an error will be returned  | 

### Return type

[**KafkaRequestList**](KafkaRequestList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceAccounts

> ServiceAccountList ListServiceAccounts(ctx).Execute()

List service accounts

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
    resp, r, err := api_client.DefaultApi.ListServiceAccounts(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListServiceAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceAccounts`: ServiceAccountList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListServiceAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListServiceAccountsRequest struct via the builder pattern


### Return type

[**ServiceAccountList**](ServiceAccountList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetServiceAccountCreds

> ServiceAccount ResetServiceAccountCreds(ctx, id).Execute()

reset credentials for the service account

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
    id := "id_example" // string | The id of record

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ResetServiceAccountCreds(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ResetServiceAccountCreds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetServiceAccountCreds`: ServiceAccount
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ResetServiceAccountCreds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetServiceAccountCredsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceAccount**](ServiceAccount.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

