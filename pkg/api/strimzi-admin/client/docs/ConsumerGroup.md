# ConsumerGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the consumer group | 
**Consumers** | [**[]Consumer**](Consumer.md) | The list of consumers associated with this consumer group | 

## Methods

### NewConsumerGroup

`func NewConsumerGroup(id string, consumers []Consumer, ) *ConsumerGroup`

NewConsumerGroup instantiates a new ConsumerGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerGroupWithDefaults

`func NewConsumerGroupWithDefaults() *ConsumerGroup`

NewConsumerGroupWithDefaults instantiates a new ConsumerGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConsumerGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsumerGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsumerGroup) SetId(v string)`

SetId sets Id field to given value.


### GetConsumers

`func (o *ConsumerGroup) GetConsumers() []Consumer`

GetConsumers returns the Consumers field if non-nil, zero value otherwise.

### GetConsumersOk

`func (o *ConsumerGroup) GetConsumersOk() (*[]Consumer, bool)`

GetConsumersOk returns a tuple with the Consumers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumers

`func (o *ConsumerGroup) SetConsumers(v []Consumer)`

SetConsumers sets Consumers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


