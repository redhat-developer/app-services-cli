# Metric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | Pointer to **map[string]string** |  | [optional] 
**Values** | Pointer to [**[]Values**](Values.md) |  | [optional] 

## Methods

### NewMetric

`func NewMetric() *Metric`

NewMetric instantiates a new Metric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricWithDefaults

`func NewMetricWithDefaults() *Metric`

NewMetricWithDefaults instantiates a new Metric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *Metric) GetMetric() map[string]string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *Metric) GetMetricOk() (*map[string]string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *Metric) SetMetric(v map[string]string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *Metric) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetValues

`func (o *Metric) GetValues() []Values`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Metric) GetValuesOk() (*[]Values, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *Metric) SetValues(v []Values)`

SetValues sets Values field to given value.

### HasValues

`func (o *Metric) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


