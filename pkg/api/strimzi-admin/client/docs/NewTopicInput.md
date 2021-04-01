# NewTopicInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The topic name, this value must be unique. | 
**Settings** | [**TopicSettings**](TopicSettings.md) |  | 

## Methods

### NewNewTopicInput

`func NewNewTopicInput(name string, settings TopicSettings, ) *NewTopicInput`

NewNewTopicInput instantiates a new NewTopicInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewTopicInputWithDefaults

`func NewNewTopicInputWithDefaults() *NewTopicInput`

NewNewTopicInputWithDefaults instantiates a new NewTopicInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NewTopicInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewTopicInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewTopicInput) SetName(v string)`

SetName sets Name field to given value.


### GetSettings

`func (o *NewTopicInput) GetSettings() TopicSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *NewTopicInput) GetSettingsOk() (*TopicSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *NewTopicInput) SetSettings(v TopicSettings)`

SetSettings sets Settings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


