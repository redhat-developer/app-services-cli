# VersionMetaData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CreatedBy** | **string** |  | 
**CreatedOn** | **time.Time** |  | 
**Type** | [**ArtifactType**](ArtifactType.md) |  | 
**GlobalId** | **int64** |  | 
**State** | Pointer to [**ArtifactState**](ArtifactState.md) |  | [optional] 
**Id** | **string** | The ID of a single artifact. | 
**Labels** | Pointer to **[]string** |  | [optional] 
**Properties** | Pointer to **map[string]string** | User-defined name-value pairs. Name and value must be strings. | [optional] 
**GroupId** | Pointer to **string** | An ID of a single artifact group. | [optional] 
**ContentId** | **int64** |  | 

## Methods

### NewVersionMetaData

`func NewVersionMetaData(version string, createdBy string, createdOn time.Time, type_ ArtifactType, globalId int64, id string, contentId int64, ) *VersionMetaData`

NewVersionMetaData instantiates a new VersionMetaData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionMetaDataWithDefaults

`func NewVersionMetaDataWithDefaults() *VersionMetaData`

NewVersionMetaDataWithDefaults instantiates a new VersionMetaData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *VersionMetaData) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VersionMetaData) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VersionMetaData) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetName

`func (o *VersionMetaData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VersionMetaData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VersionMetaData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VersionMetaData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *VersionMetaData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VersionMetaData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VersionMetaData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VersionMetaData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedBy

`func (o *VersionMetaData) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *VersionMetaData) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *VersionMetaData) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedOn

`func (o *VersionMetaData) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *VersionMetaData) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *VersionMetaData) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.


### GetType

`func (o *VersionMetaData) GetType() ArtifactType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VersionMetaData) GetTypeOk() (*ArtifactType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VersionMetaData) SetType(v ArtifactType)`

SetType sets Type field to given value.


### GetGlobalId

`func (o *VersionMetaData) GetGlobalId() int64`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *VersionMetaData) GetGlobalIdOk() (*int64, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *VersionMetaData) SetGlobalId(v int64)`

SetGlobalId sets GlobalId field to given value.


### GetState

`func (o *VersionMetaData) GetState() ArtifactState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VersionMetaData) GetStateOk() (*ArtifactState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VersionMetaData) SetState(v ArtifactState)`

SetState sets State field to given value.

### HasState

`func (o *VersionMetaData) HasState() bool`

HasState returns a boolean if a field has been set.

### GetId

`func (o *VersionMetaData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VersionMetaData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VersionMetaData) SetId(v string)`

SetId sets Id field to given value.


### GetLabels

`func (o *VersionMetaData) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *VersionMetaData) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *VersionMetaData) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *VersionMetaData) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetProperties

`func (o *VersionMetaData) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *VersionMetaData) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *VersionMetaData) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *VersionMetaData) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetGroupId

`func (o *VersionMetaData) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *VersionMetaData) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *VersionMetaData) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *VersionMetaData) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetContentId

`func (o *VersionMetaData) GetContentId() int64`

GetContentId returns the ContentId field if non-nil, zero value otherwise.

### GetContentIdOk

`func (o *VersionMetaData) GetContentIdOk() (*int64, bool)`

GetContentIdOk returns a tuple with the ContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentId

`func (o *VersionMetaData) SetContentId(v int64)`

SetContentId sets ContentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


