# NamedLogConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Level** | [**LogLevel**](LogLevel.md) |  | 

## Methods

### NewNamedLogConfiguration

`func NewNamedLogConfiguration(name string, level LogLevel, ) *NamedLogConfiguration`

NewNamedLogConfiguration instantiates a new NamedLogConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamedLogConfigurationWithDefaults

`func NewNamedLogConfigurationWithDefaults() *NamedLogConfiguration`

NewNamedLogConfigurationWithDefaults instantiates a new NamedLogConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NamedLogConfiguration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NamedLogConfiguration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NamedLogConfiguration) SetName(v string)`

SetName sets Name field to given value.


### GetLevel

`func (o *NamedLogConfiguration) GetLevel() LogLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *NamedLogConfiguration) GetLevelOk() (*LogLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *NamedLogConfiguration) SetLevel(v LogLevel)`

SetLevel sets Level field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


