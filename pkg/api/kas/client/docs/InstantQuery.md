# InstantQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | Pointer to **map[string]string** |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 
**Value** | **float64** |  | 

## Methods

### NewInstantQuery

`func NewInstantQuery(value float64, ) *InstantQuery`

NewInstantQuery instantiates a new InstantQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstantQueryWithDefaults

`func NewInstantQueryWithDefaults() *InstantQuery`

NewInstantQueryWithDefaults instantiates a new InstantQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *InstantQuery) GetMetric() map[string]string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *InstantQuery) GetMetricOk() (*map[string]string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *InstantQuery) SetMetric(v map[string]string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *InstantQuery) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetTimestamp

`func (o *InstantQuery) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *InstantQuery) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *InstantQuery) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *InstantQuery) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetValue

`func (o *InstantQuery) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InstantQuery) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InstantQuery) SetValue(v float64)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


