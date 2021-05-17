# EditableMetaData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Properties** | Pointer to **map[string]string** | User-defined name-value pairs. Name and value must be strings. | [optional] 

## Methods

### NewEditableMetaData

`func NewEditableMetaData() *EditableMetaData`

NewEditableMetaData instantiates a new EditableMetaData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditableMetaDataWithDefaults

`func NewEditableMetaDataWithDefaults() *EditableMetaData`

NewEditableMetaDataWithDefaults instantiates a new EditableMetaData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EditableMetaData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditableMetaData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditableMetaData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EditableMetaData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *EditableMetaData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EditableMetaData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EditableMetaData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EditableMetaData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *EditableMetaData) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *EditableMetaData) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *EditableMetaData) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *EditableMetaData) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetProperties

`func (o *EditableMetaData) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *EditableMetaData) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *EditableMetaData) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *EditableMetaData) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


