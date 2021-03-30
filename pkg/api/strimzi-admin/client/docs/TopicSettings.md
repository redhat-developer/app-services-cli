# TopicSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumPartitions** | **int32** | Number of partitions for this topic. | 
**Config** | Pointer to [**[]ConfigEntry**](ConfigEntry.md) | Topic configuration entry. | [optional] 

## Methods

### NewTopicSettings

`func NewTopicSettings(numPartitions int32, ) *TopicSettings`

NewTopicSettings instantiates a new TopicSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopicSettingsWithDefaults

`func NewTopicSettingsWithDefaults() *TopicSettings`

NewTopicSettingsWithDefaults instantiates a new TopicSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumPartitions

`func (o *TopicSettings) GetNumPartitions() int32`

GetNumPartitions returns the NumPartitions field if non-nil, zero value otherwise.

### GetNumPartitionsOk

`func (o *TopicSettings) GetNumPartitionsOk() (*int32, bool)`

GetNumPartitionsOk returns a tuple with the NumPartitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPartitions

`func (o *TopicSettings) SetNumPartitions(v int32)`

SetNumPartitions sets NumPartitions field to given value.


### GetConfig

`func (o *TopicSettings) GetConfig() []ConfigEntry`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *TopicSettings) GetConfigOk() (*[]ConfigEntry, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *TopicSettings) SetConfig(v []ConfigEntry)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *TopicSettings) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


