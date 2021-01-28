# MetricsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]Metric**](Metric.md) |  | [optional] 

## Methods

### NewMetricsList

`func NewMetricsList() *MetricsList`

NewMetricsList instantiates a new MetricsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsListWithDefaults

`func NewMetricsListWithDefaults() *MetricsList`

NewMetricsListWithDefaults instantiates a new MetricsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *MetricsList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MetricsList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MetricsList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *MetricsList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *MetricsList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetricsList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetricsList) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MetricsList) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItems

`func (o *MetricsList) GetItems() []Metric`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *MetricsList) GetItemsOk() (*[]Metric, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *MetricsList) SetItems(v []Metric)`

SetItems sets Items field to given value.

### HasItems

`func (o *MetricsList) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


