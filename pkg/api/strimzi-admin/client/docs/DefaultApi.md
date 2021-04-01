# \DefaultApi

All URIs are relative to *https://api.openshift.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTopic**](DefaultApi.md#CreateTopic) | **Post** /topics | Creates a new topic
[**DeleteConsumerGroupById**](DefaultApi.md#DeleteConsumerGroupById) | **Delete** /consumer-groups/{consumerGroupId} | Delete a consumer group.
[**DeleteTopic**](DefaultApi.md#DeleteTopic) | **Delete** /topics/{topicName} | Deletes a  topic
[**GetConsumerGroupById**](DefaultApi.md#GetConsumerGroupById) | **Get** /consumer-groups/{consumerGroupId} | Get a single consumer group by its unique ID.
[**GetConsumerGroupList**](DefaultApi.md#GetConsumerGroupList) | **Get** /consumer-groups | List of consumer groups in the Kafka instance.
[**GetTopic**](DefaultApi.md#GetTopic) | **Get** /topics/{topicName} | Retrieves the topic with the specified name.
[**GetTopicsList**](DefaultApi.md#GetTopicsList) | **Get** /topics | List of topics
[**Metrics**](DefaultApi.md#Metrics) | **Get** /metrics | Admin server metrics
[**OpenApi**](DefaultApi.md#OpenApi) | **Get** /api | 
[**ResetConsumerGroupOffset**](DefaultApi.md#ResetConsumerGroupOffset) | **Post** /consumer-groups/{consumerGroupId}/reset-offset | Reset the offset for a consumer group.
[**UpdateTopic**](DefaultApi.md#UpdateTopic) | **Patch** /topics/{topicName} | Updates the topic with the specified name.



## CreateTopic

> Topic CreateTopic(ctx).NewTopicInput(newTopicInput).Execute()

Creates a new topic



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
    newTopicInput := *openapiclient.NewNewTopicInput("Name_example", *openapiclient.NewTopicSettings(int32(123))) // NewTopicInput | Topic to create.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateTopic(context.Background()).NewTopicInput(newTopicInput).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTopic`: Topic
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateTopic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newTopicInput** | [**NewTopicInput**](NewTopicInput.md) | Topic to create. | 

### Return type

[**Topic**](Topic.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConsumerGroupById

> DeleteConsumerGroupById(ctx, consumerGroupId).Execute()

Delete a consumer group.



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
    consumerGroupId := "consumerGroupId_example" // string | The unique name of the topic.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteConsumerGroupById(context.Background(), consumerGroupId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteConsumerGroupById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerGroupId** | **string** | The unique name of the topic. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConsumerGroupByIdRequest struct via the builder pattern


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


## DeleteTopic

> DeleteTopic(ctx, topicName).Execute()

Deletes a  topic



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
    topicName := "topicName_example" // string | The topic name to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteTopic(context.Background(), topicName).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topicName** | **string** | The topic name to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConsumerGroupById

> ConsumerGroup GetConsumerGroupById(ctx, consumerGroupId).Topic(topic).Execute()

Get a single consumer group by its unique ID.

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
    consumerGroupId := "consumerGroupId_example" // string | The unique ID of the consumer group
    topic := "topic_example" // string | Filter consumer groups for a specific topic (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetConsumerGroupById(context.Background(), consumerGroupId).Topic(topic).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetConsumerGroupById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConsumerGroupById`: ConsumerGroup
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetConsumerGroupById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerGroupId** | **string** | The unique ID of the consumer group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConsumerGroupByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **topic** | **string** | Filter consumer groups for a specific topic | 

### Return type

[**ConsumerGroup**](ConsumerGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConsumerGroupList

> ConsumerGroupList GetConsumerGroupList(ctx).Limit(limit).Offset(offset).Topic(topic).Execute()

List of consumer groups in the Kafka instance.



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
    limit := int32(56) // int32 | Maximum number of consumer groups to returnd (optional)
    offset := int32(56) // int32 | The page offset when returning the list of consumer groups (optional)
    topic := "topic_example" // string | Filter to apply when returning the list of consumer groups (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetConsumerGroupList(context.Background()).Limit(limit).Offset(offset).Topic(topic).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetConsumerGroupList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConsumerGroupList`: ConsumerGroupList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetConsumerGroupList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConsumerGroupListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of consumer groups to returnd | 
 **offset** | **int32** | The page offset when returning the list of consumer groups | 
 **topic** | **string** | Filter to apply when returning the list of consumer groups | 

### Return type

[**ConsumerGroupList**](ConsumerGroupList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTopic

> Topic GetTopic(ctx, topicName).Execute()

Retrieves the topic with the specified name.



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
    topicName := "topicName_example" // string | The topic name to retrieve.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetTopic(context.Background(), topicName).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTopic`: Topic
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topicName** | **string** | The topic name to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Topic**](Topic.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTopicsList

> TopicsList GetTopicsList(ctx).Limit(limit).Filter(filter).Offset(offset).Order(order).Execute()

List of topics



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
    limit := int32(56) // int32 | Maximum number of topics to return (optional)
    filter := "filter_example" // string | Filter to apply when returning the list of topics (optional)
    offset := int32(56) // int32 | The page offset when returning  the limit of requested topics. (optional)
    order := "order_example" // string | Order of the items sorting. If \"asc\" is set as a value, ascending order is used, descending otherwise. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetTopicsList(context.Background()).Limit(limit).Filter(filter).Offset(offset).Order(order).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetTopicsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTopicsList`: TopicsList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetTopicsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTopicsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of topics to return | 
 **filter** | **string** | Filter to apply when returning the list of topics | 
 **offset** | **int32** | The page offset when returning  the limit of requested topics. | 
 **order** | **string** | Order of the items sorting. If \&quot;asc\&quot; is set as a value, ascending order is used, descending otherwise. | 

### Return type

[**TopicsList**](TopicsList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Metrics

> Metrics(ctx).Execute()

Admin server metrics

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
    resp, r, err := api_client.DefaultApi.Metrics(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Metrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMetricsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenApi

> OpenApi(ctx).Execute()



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
    resp, r, err := api_client.DefaultApi.OpenApi(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OpenApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOpenApiRequest struct via the builder pattern


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


## ResetConsumerGroupOffset

> ConsumerGroup ResetConsumerGroupOffset(ctx, consumerGroupId).Execute()

Reset the offset for a consumer group.



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
    consumerGroupId := "consumerGroupId_example" // string | The ID of the consumer group.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ResetConsumerGroupOffset(context.Background(), consumerGroupId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ResetConsumerGroupOffset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetConsumerGroupOffset`: ConsumerGroup
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ResetConsumerGroupOffset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerGroupId** | **string** | The ID of the consumer group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetConsumerGroupOffsetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConsumerGroup**](ConsumerGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTopic

> Topic UpdateTopic(ctx, topicName).UpdateTopicInput(updateTopicInput).Execute()

Updates the topic with the specified name.



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
    topicName := "topicName_example" // string | The topic name which is its unique id.
    updateTopicInput := *openapiclient.NewUpdateTopicInput() // UpdateTopicInput | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateTopic(context.Background(), topicName).UpdateTopicInput(updateTopicInput).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTopic`: Topic
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topicName** | **string** | The topic name which is its unique id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTopicInput** | [**UpdateTopicInput**](UpdateTopicInput.md) |  | 

### Return type

[**Topic**](Topic.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

