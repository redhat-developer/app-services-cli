# Summary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Metrics** | [**[]SummaryMetrics**](SummaryMetrics.md) |  | 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewSummary

`func NewSummary(metrics []SummaryMetrics, ) *Summary`

NewSummary instantiates a new Summary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryWithDefaults

`func NewSummaryWithDefaults() *Summary`

NewSummaryWithDefaults instantiates a new Summary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *Summary) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Summary) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Summary) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Summary) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *Summary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Summary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Summary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Summary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *Summary) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Summary) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Summary) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Summary) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetrics

`func (o *Summary) GetMetrics() []SummaryMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *Summary) GetMetricsOk() (*[]SummaryMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *Summary) SetMetrics(v []SummaryMetrics)`

SetMetrics sets Metrics field to given value.


### GetName

`func (o *Summary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Summary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Summary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Summary) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


