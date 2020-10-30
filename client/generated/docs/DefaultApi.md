# \DefaultApi

All URIs are relative to *https://api.openshift.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiManagedServicesApiV1KafkasGet**](DefaultApi.md#ApiManagedServicesApiV1KafkasGet) | **Get** /api/managed-services-api/v1/kafkas | Returns a list of Kafka requests
[**ApiManagedServicesApiV1KafkasIdDelete**](DefaultApi.md#ApiManagedServicesApiV1KafkasIdDelete) | **Delete** /api/managed-services-api/v1/kafkas/{id} | Delete a kafka request by id
[**ApiManagedServicesApiV1KafkasIdGet**](DefaultApi.md#ApiManagedServicesApiV1KafkasIdGet) | **Get** /api/managed-services-api/v1/kafkas/{id} | Get a kafka request by id
[**ApiManagedServicesApiV1KafkasPost**](DefaultApi.md#ApiManagedServicesApiV1KafkasPost) | **Post** /api/managed-services-api/v1/kafkas | Create a new kafka Request



## ApiManagedServicesApiV1KafkasGet

> KafkaRequestList ApiManagedServicesApiV1KafkasGet(ctx, optional)

Returns a list of Kafka requests

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApiManagedServicesApiV1KafkasGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApiManagedServicesApiV1KafkasGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**| Page index | 
 **size** | **optional.String**| Number of items in each page | 

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


## ApiManagedServicesApiV1KafkasIdDelete

> Error ApiManagedServicesApiV1KafkasIdDelete(ctx, id)

Delete a kafka request by id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of record | 

### Return type

[**Error**](Error.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiManagedServicesApiV1KafkasIdGet

> KafkaRequest ApiManagedServicesApiV1KafkasIdGet(ctx, id)

Get a kafka request by id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of record | 

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


## ApiManagedServicesApiV1KafkasPost

> KafkaRequest ApiManagedServicesApiV1KafkasPost(ctx, async, kafkaRequest)

Create a new kafka Request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**async** | **bool**| Perform the action in an asynchronous manner | 
**kafkaRequest** | [**KafkaRequest**](KafkaRequest.md)| Kafka data | 

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

