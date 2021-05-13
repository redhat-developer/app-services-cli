# ErrorInfo1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | **int32** |  | 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewErrorInfo1

`func NewErrorInfo1(errorCode int32, ) *ErrorInfo1`

NewErrorInfo1 instantiates a new ErrorInfo1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorInfo1WithDefaults

`func NewErrorInfo1WithDefaults() *ErrorInfo1`

NewErrorInfo1WithDefaults instantiates a new ErrorInfo1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *ErrorInfo1) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ErrorInfo1) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ErrorInfo1) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.


### GetMessage

`func (o *ErrorInfo1) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorInfo1) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorInfo1) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ErrorInfo1) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


