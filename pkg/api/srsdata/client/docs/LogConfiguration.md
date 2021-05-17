# LogConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | [**LogLevel**](LogLevel.md) |  | 

## Methods

### NewLogConfiguration

`func NewLogConfiguration(level LogLevel, ) *LogConfiguration`

NewLogConfiguration instantiates a new LogConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogConfigurationWithDefaults

`func NewLogConfigurationWithDefaults() *LogConfiguration`

NewLogConfigurationWithDefaults instantiates a new LogConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *LogConfiguration) GetLevel() LogLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *LogConfiguration) GetLevelOk() (*LogLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *LogConfiguration) SetLevel(v LogLevel)`

SetLevel sets Level field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


