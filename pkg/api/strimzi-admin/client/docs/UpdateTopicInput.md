# UpdateTopicInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**[]ConfigEntry**](ConfigEntry.md) | Topic configuration entry. | [optional] 

## Methods

### NewUpdateTopicInput

`func NewUpdateTopicInput() *UpdateTopicInput`

NewUpdateTopicInput instantiates a new UpdateTopicInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTopicInputWithDefaults

`func NewUpdateTopicInputWithDefaults() *UpdateTopicInput`

NewUpdateTopicInputWithDefaults instantiates a new UpdateTopicInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *UpdateTopicInput) GetConfig() []ConfigEntry`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateTopicInput) GetConfigOk() (*[]ConfigEntry, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateTopicInput) SetConfig(v []ConfigEntry)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateTopicInput) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


