# ArtifactMetaData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CreatedBy** | **string** |  | 
**CreatedOn** | **time.Time** |  | 
**ModifiedBy** | **string** |  | 
**ModifiedOn** | **time.Time** |  | 
**Id** | **string** | The ID of a single artifact. | 
**Version** | **string** |  | 
**Type** | [**ArtifactType**](ArtifactType.md) |  | 
**GlobalId** | **int64** |  | 
**State** | [**ArtifactState**](ArtifactState.md) |  | 
**Labels** | Pointer to **[]string** |  | [optional] 
**Properties** | Pointer to **map[string]string** | User-defined name-value pairs. Name and value must be strings. | [optional] 
**GroupId** | Pointer to **string** | An ID of a single artifact group. | [optional] 
**ContentId** | **int64** |  | 

## Methods

### NewArtifactMetaData

`func NewArtifactMetaData(createdBy string, createdOn time.Time, modifiedBy string, modifiedOn time.Time, id string, version string, type_ ArtifactType, globalId int64, state ArtifactState, contentId int64, ) *ArtifactMetaData`

NewArtifactMetaData instantiates a new ArtifactMetaData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactMetaDataWithDefaults

`func NewArtifactMetaDataWithDefaults() *ArtifactMetaData`

NewArtifactMetaDataWithDefaults instantiates a new ArtifactMetaData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ArtifactMetaData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArtifactMetaData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArtifactMetaData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ArtifactMetaData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ArtifactMetaData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ArtifactMetaData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ArtifactMetaData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ArtifactMetaData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ArtifactMetaData) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ArtifactMetaData) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ArtifactMetaData) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedOn

`func (o *ArtifactMetaData) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *ArtifactMetaData) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *ArtifactMetaData) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.


### GetModifiedBy

`func (o *ArtifactMetaData) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *ArtifactMetaData) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *ArtifactMetaData) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetModifiedOn

`func (o *ArtifactMetaData) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *ArtifactMetaData) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *ArtifactMetaData) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.


### GetId

`func (o *ArtifactMetaData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArtifactMetaData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArtifactMetaData) SetId(v string)`

SetId sets Id field to given value.


### GetVersion

`func (o *ArtifactMetaData) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ArtifactMetaData) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ArtifactMetaData) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetType

`func (o *ArtifactMetaData) GetType() ArtifactType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArtifactMetaData) GetTypeOk() (*ArtifactType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArtifactMetaData) SetType(v ArtifactType)`

SetType sets Type field to given value.


### GetGlobalId

`func (o *ArtifactMetaData) GetGlobalId() int64`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *ArtifactMetaData) GetGlobalIdOk() (*int64, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *ArtifactMetaData) SetGlobalId(v int64)`

SetGlobalId sets GlobalId field to given value.


### GetState

`func (o *ArtifactMetaData) GetState() ArtifactState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ArtifactMetaData) GetStateOk() (*ArtifactState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ArtifactMetaData) SetState(v ArtifactState)`

SetState sets State field to given value.


### GetLabels

`func (o *ArtifactMetaData) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ArtifactMetaData) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ArtifactMetaData) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ArtifactMetaData) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetProperties

`func (o *ArtifactMetaData) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ArtifactMetaData) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ArtifactMetaData) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ArtifactMetaData) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetGroupId

`func (o *ArtifactMetaData) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ArtifactMetaData) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ArtifactMetaData) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ArtifactMetaData) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetContentId

`func (o *ArtifactMetaData) GetContentId() int64`

GetContentId returns the ContentId field if non-nil, zero value otherwise.

### GetContentIdOk

`func (o *ArtifactMetaData) GetContentIdOk() (*int64, bool)`

GetContentIdOk returns a tuple with the ContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentId

`func (o *ArtifactMetaData) SetContentId(v int64)`

SetContentId sets ContentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


