# SummaryMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Vector** | Pointer to [**[]SummaryVector**](SummaryVector.md) |  | [optional] 

## Methods

### NewSummaryMetrics

`func NewSummaryMetrics() *SummaryMetrics`

NewSummaryMetrics instantiates a new SummaryMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryMetricsWithDefaults

`func NewSummaryMetricsWithDefaults() *SummaryMetrics`

NewSummaryMetricsWithDefaults instantiates a new SummaryMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SummaryMetrics) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SummaryMetrics) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SummaryMetrics) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SummaryMetrics) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVector

`func (o *SummaryMetrics) GetVector() []SummaryVector`

GetVector returns the Vector field if non-nil, zero value otherwise.

### GetVectorOk

`func (o *SummaryMetrics) GetVectorOk() (*[]SummaryVector, bool)`

GetVectorOk returns a tuple with the Vector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVector

`func (o *SummaryMetrics) SetVector(v []SummaryVector)`

SetVector sets Vector field to given value.

### HasVector

`func (o *SummaryMetrics) HasVector() bool`

HasVector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


