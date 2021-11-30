# SummaryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | [**[]SummaryMetrics**](SummaryMetrics.md) |  | 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewSummaryAllOf

`func NewSummaryAllOf(metrics []SummaryMetrics, ) *SummaryAllOf`

NewSummaryAllOf instantiates a new SummaryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryAllOfWithDefaults

`func NewSummaryAllOfWithDefaults() *SummaryAllOf`

NewSummaryAllOfWithDefaults instantiates a new SummaryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *SummaryAllOf) GetMetrics() []SummaryMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *SummaryAllOf) GetMetricsOk() (*[]SummaryMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *SummaryAllOf) SetMetrics(v []SummaryMetrics)`

SetMetrics sets Metrics field to given value.


### GetName

`func (o *SummaryAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SummaryAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SummaryAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SummaryAllOf) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


