# Consumer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** | Unique identifier for the consumer group to which this consumer belongs. | 
**Topic** | **string** | The unique topic name to which this consumer belongs | 
**Partition** | **int32** | The partition number to which this consumer group is assigned to. | 
**Offset** | **float32** | Offset denotes the position of the consumer in a partition. | 
**LogEndOffset** | Pointer to **float32** | The log end offset is the offset of the last message written to a log. | [optional] 
**Lag** | **int32** | Offset Lag is the delta between the last produced message and the last consumer&#39;s committed offset. | 
**MemberId** | Pointer to **string** | The member ID is a unique identifier given to a consumer by the coordinator upon initially joining the group. | [optional] 

## Methods

### NewConsumer

`func NewConsumer(groupId string, topic string, partition int32, offset float32, lag int32, ) *Consumer`

NewConsumer instantiates a new Consumer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerWithDefaults

`func NewConsumerWithDefaults() *Consumer`

NewConsumerWithDefaults instantiates a new Consumer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *Consumer) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *Consumer) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *Consumer) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetTopic

`func (o *Consumer) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *Consumer) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *Consumer) SetTopic(v string)`

SetTopic sets Topic field to given value.


### GetPartition

`func (o *Consumer) GetPartition() int32`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *Consumer) GetPartitionOk() (*int32, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *Consumer) SetPartition(v int32)`

SetPartition sets Partition field to given value.


### GetOffset

`func (o *Consumer) GetOffset() float32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *Consumer) GetOffsetOk() (*float32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *Consumer) SetOffset(v float32)`

SetOffset sets Offset field to given value.


### GetLogEndOffset

`func (o *Consumer) GetLogEndOffset() float32`

GetLogEndOffset returns the LogEndOffset field if non-nil, zero value otherwise.

### GetLogEndOffsetOk

`func (o *Consumer) GetLogEndOffsetOk() (*float32, bool)`

GetLogEndOffsetOk returns a tuple with the LogEndOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogEndOffset

`func (o *Consumer) SetLogEndOffset(v float32)`

SetLogEndOffset sets LogEndOffset field to given value.

### HasLogEndOffset

`func (o *Consumer) HasLogEndOffset() bool`

HasLogEndOffset returns a boolean if a field has been set.

### GetLag

`func (o *Consumer) GetLag() int32`

GetLag returns the Lag field if non-nil, zero value otherwise.

### GetLagOk

`func (o *Consumer) GetLagOk() (*int32, bool)`

GetLagOk returns a tuple with the Lag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLag

`func (o *Consumer) SetLag(v int32)`

SetLag sets Lag field to given value.


### GetMemberId

`func (o *Consumer) GetMemberId() string`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *Consumer) GetMemberIdOk() (*string, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *Consumer) SetMemberId(v string)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *Consumer) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


