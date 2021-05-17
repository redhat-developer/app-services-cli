# Rule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | **string** |  | 
**Type** | Pointer to [**RuleType**](RuleType.md) |  | [optional] 

## Methods

### NewRule

`func NewRule(config string, ) *Rule`

NewRule instantiates a new Rule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleWithDefaults

`func NewRuleWithDefaults() *Rule`

NewRuleWithDefaults instantiates a new Rule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *Rule) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Rule) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Rule) SetConfig(v string)`

SetConfig sets Config field to given value.


### GetType

`func (o *Rule) GetType() RuleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Rule) GetTypeOk() (*RuleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Rule) SetType(v RuleType)`

SetType sets Type field to given value.

### HasType

`func (o *Rule) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


