# RuleViolationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Causes** | [**[]RuleViolationCause**](RuleViolationCause.md) | List of rule violation causes. | 
**Message** | Pointer to **string** | The short error message. | [optional] 
**ErrorCode** | Pointer to **int32** | The server-side error code. | [optional] 
**Detail** | Pointer to **string** | Full details about the error.  This might contain a server stack trace, for example. | [optional] 
**Name** | Pointer to **string** | The error name - typically the classname of the exception thrown by the server. | [optional] 

## Methods

### NewRuleViolationError

`func NewRuleViolationError(causes []RuleViolationCause, ) *RuleViolationError`

NewRuleViolationError instantiates a new RuleViolationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleViolationErrorWithDefaults

`func NewRuleViolationErrorWithDefaults() *RuleViolationError`

NewRuleViolationErrorWithDefaults instantiates a new RuleViolationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCauses

`func (o *RuleViolationError) GetCauses() []RuleViolationCause`

GetCauses returns the Causes field if non-nil, zero value otherwise.

### GetCausesOk

`func (o *RuleViolationError) GetCausesOk() (*[]RuleViolationCause, bool)`

GetCausesOk returns a tuple with the Causes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCauses

`func (o *RuleViolationError) SetCauses(v []RuleViolationCause)`

SetCauses sets Causes field to given value.


### GetMessage

`func (o *RuleViolationError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RuleViolationError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RuleViolationError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RuleViolationError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetErrorCode

`func (o *RuleViolationError) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *RuleViolationError) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *RuleViolationError) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *RuleViolationError) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetDetail

`func (o *RuleViolationError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *RuleViolationError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *RuleViolationError) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *RuleViolationError) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetName

`func (o *RuleViolationError) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleViolationError) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleViolationError) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RuleViolationError) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


