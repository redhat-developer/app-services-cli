# SystemInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**BuiltOn** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSystemInfo

`func NewSystemInfo() *SystemInfo`

NewSystemInfo instantiates a new SystemInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemInfoWithDefaults

`func NewSystemInfoWithDefaults() *SystemInfo`

NewSystemInfoWithDefaults instantiates a new SystemInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SystemInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SystemInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SystemInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SystemInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *SystemInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SystemInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SystemInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SystemInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersion

`func (o *SystemInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SystemInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SystemInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SystemInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBuiltOn

`func (o *SystemInfo) GetBuiltOn() time.Time`

GetBuiltOn returns the BuiltOn field if non-nil, zero value otherwise.

### GetBuiltOnOk

`func (o *SystemInfo) GetBuiltOnOk() (*time.Time, bool)`

GetBuiltOnOk returns a tuple with the BuiltOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuiltOn

`func (o *SystemInfo) SetBuiltOn(v time.Time)`

SetBuiltOn sets BuiltOn field to given value.

### HasBuiltOn

`func (o *SystemInfo) HasBuiltOn() bool`

HasBuiltOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


