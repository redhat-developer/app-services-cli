# ConsumerGroupList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]ConsumerGroup**](ConsumerGroup.md) | Consumer group list items | 
**Count** | **float32** | The total number of consumer groups. | 
**Limit** | **float32** | The number of consumer groups per page. | 
**Offset** | **int32** | The page offset | 

## Methods

### NewConsumerGroupList

`func NewConsumerGroupList(items []ConsumerGroup, count float32, limit float32, offset int32, ) *ConsumerGroupList`

NewConsumerGroupList instantiates a new ConsumerGroupList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerGroupListWithDefaults

`func NewConsumerGroupListWithDefaults() *ConsumerGroupList`

NewConsumerGroupListWithDefaults instantiates a new ConsumerGroupList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ConsumerGroupList) GetItems() []ConsumerGroup`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ConsumerGroupList) GetItemsOk() (*[]ConsumerGroup, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ConsumerGroupList) SetItems(v []ConsumerGroup)`

SetItems sets Items field to given value.


### GetCount

`func (o *ConsumerGroupList) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ConsumerGroupList) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ConsumerGroupList) SetCount(v float32)`

SetCount sets Count field to given value.


### GetLimit

`func (o *ConsumerGroupList) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ConsumerGroupList) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ConsumerGroupList) SetLimit(v float32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *ConsumerGroupList) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ConsumerGroupList) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ConsumerGroupList) SetOffset(v int32)`

SetOffset sets Offset field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


