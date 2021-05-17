# SearchedVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CreatedOn** | **time.Time** |  | 
**CreatedBy** | **string** |  | 
**Type** | [**ArtifactType**](ArtifactType.md) |  | 
**Labels** | Pointer to **[]string** |  | [optional] 
**State** | [**ArtifactState**](ArtifactState.md) |  | 
**GlobalId** | **int64** |  | 
**Version** | **string** |  | 
**Properties** | Pointer to **map[string]string** | User-defined name-value pairs. Name and value must be strings. | [optional] 
**ContentId** | **int64** |  | 

## Methods

### NewSearchedVersion

`func NewSearchedVersion(createdOn time.Time, createdBy string, type_ ArtifactType, state ArtifactState, globalId int64, version string, contentId int64, ) *SearchedVersion`

NewSearchedVersion instantiates a new SearchedVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchedVersionWithDefaults

`func NewSearchedVersionWithDefaults() *SearchedVersion`

NewSearchedVersionWithDefaults instantiates a new SearchedVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SearchedVersion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchedVersion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchedVersion) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SearchedVersion) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *SearchedVersion) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SearchedVersion) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SearchedVersion) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SearchedVersion) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedOn

`func (o *SearchedVersion) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *SearchedVersion) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *SearchedVersion) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.


### GetCreatedBy

`func (o *SearchedVersion) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SearchedVersion) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SearchedVersion) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetType

`func (o *SearchedVersion) GetType() ArtifactType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SearchedVersion) GetTypeOk() (*ArtifactType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SearchedVersion) SetType(v ArtifactType)`

SetType sets Type field to given value.


### GetLabels

`func (o *SearchedVersion) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *SearchedVersion) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *SearchedVersion) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *SearchedVersion) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetState

`func (o *SearchedVersion) GetState() ArtifactState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SearchedVersion) GetStateOk() (*ArtifactState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SearchedVersion) SetState(v ArtifactState)`

SetState sets State field to given value.


### GetGlobalId

`func (o *SearchedVersion) GetGlobalId() int64`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *SearchedVersion) GetGlobalIdOk() (*int64, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *SearchedVersion) SetGlobalId(v int64)`

SetGlobalId sets GlobalId field to given value.


### GetVersion

`func (o *SearchedVersion) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SearchedVersion) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SearchedVersion) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetProperties

`func (o *SearchedVersion) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SearchedVersion) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SearchedVersion) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *SearchedVersion) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetContentId

`func (o *SearchedVersion) GetContentId() int64`

GetContentId returns the ContentId field if non-nil, zero value otherwise.

### GetContentIdOk

`func (o *SearchedVersion) GetContentIdOk() (*int64, bool)`

GetContentIdOk returns a tuple with the ContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentId

`func (o *SearchedVersion) SetContentId(v int64)`

SetContentId sets ContentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


