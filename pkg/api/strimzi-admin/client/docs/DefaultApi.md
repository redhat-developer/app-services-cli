# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTopic**](DefaultApi.md#CreateTopic) | **Post** /topics | Creates a new topic
[**DeleteTopic**](DefaultApi.md#DeleteTopic) | **Delete** /topics/{topicId} | Deletes a  topic
[**GetTopic**](DefaultApi.md#GetTopic) | **Get** /topics/{topicId} | Topic associated with the topic id
[**GetTopicsList**](DefaultApi.md#GetTopicsList) | **Get** /topics | List of topics
[**UpdateTopic**](DefaultApi.md#UpdateTopic) | **Patch** /topics/{topicId} | Updates the topic with the specified id.



## CreateTopic

> Topic CreateTopic(ctx).TopicSettings(topicSettings).Execute()

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
    topicSettings := *openapiclient.NewTopicSettings() // TopicSettings | Topic to create.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateTopic(context.Background()).TopicSettings(topicSettings).Execute()
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
 **topicSettings** | [**TopicSettings**](TopicSettings.md) | Topic to create. | 

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


## DeleteTopic

> DeleteTopic(ctx, topicId).Execute()

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
    topicId := "topicId_example" // string | Topic id to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteTopic(context.Background(), topicId).Execute()
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
**topicId** | **string** | Topic id to delete | 

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


## GetTopic

> Topic GetTopic(ctx, topicId).Execute()

Topic associated with the topic id



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
    topicId := "topicId_example" // string | The id of the topic

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetTopic(context.Background(), topicId).Execute()
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
**topicId** | **string** | The id of the topic | 

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

> TopicsList GetTopicsList(ctx).Limit(limit).Filter(filter).Offset(offset).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetTopicsList(context.Background()).Limit(limit).Filter(filter).Offset(offset).Execute()
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


## UpdateTopic

> Topic UpdateTopic(ctx, topicId).TopicSettings(topicSettings).Execute()

Updates the topic with the specified id.



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
    topicId := "topicId_example" // string | topic id to update
    topicSettings := *openapiclient.NewTopicSettings() // TopicSettings | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateTopic(context.Background(), topicId).TopicSettings(topicSettings).Execute()
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
**topicId** | **string** | topic id to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **topicSettings** | [**TopicSettings**](TopicSettings.md) |  | 

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

