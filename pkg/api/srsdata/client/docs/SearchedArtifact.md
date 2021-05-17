# SearchedArtifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of a single artifact. | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CreatedOn** | **time.Time** |  | 
**CreatedBy** | **string** |  | 
**Type** | [**ArtifactType**](ArtifactType.md) |  | 
**Labels** | Pointer to **[]string** |  | [optional] 
**State** | [**ArtifactState**](ArtifactState.md) |  | 
**ModifiedOn** | Pointer to **time.Time** |  | [optional] 
**ModifiedBy** | Pointer to **string** |  | [optional] 
**GroupId** | Pointer to **string** | An ID of a single artifact group. | [optional] 

## Methods

### NewSearchedArtifact

`func NewSearchedArtifact(id string, createdOn time.Time, createdBy string, type_ ArtifactType, state ArtifactState, ) *SearchedArtifact`

NewSearchedArtifact instantiates a new SearchedArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchedArtifactWithDefaults

`func NewSearchedArtifactWithDefaults() *SearchedArtifact`

NewSearchedArtifactWithDefaults instantiates a new SearchedArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SearchedArtifact) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchedArtifact) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchedArtifact) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SearchedArtifact) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchedArtifact) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchedArtifact) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SearchedArtifact) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *SearchedArtifact) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SearchedArtifact) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SearchedArtifact) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SearchedArtifact) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedOn

`func (o *SearchedArtifact) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *SearchedArtifact) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *SearchedArtifact) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.


### GetCreatedBy

`func (o *SearchedArtifact) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SearchedArtifact) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SearchedArtifact) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetType

`func (o *SearchedArtifact) GetType() ArtifactType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SearchedArtifact) GetTypeOk() (*ArtifactType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SearchedArtifact) SetType(v ArtifactType)`

SetType sets Type field to given value.


### GetLabels

`func (o *SearchedArtifact) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *SearchedArtifact) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *SearchedArtifact) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *SearchedArtifact) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetState

`func (o *SearchedArtifact) GetState() ArtifactState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SearchedArtifact) GetStateOk() (*ArtifactState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SearchedArtifact) SetState(v ArtifactState)`

SetState sets State field to given value.


### GetModifiedOn

`func (o *SearchedArtifact) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *SearchedArtifact) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *SearchedArtifact) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *SearchedArtifact) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetModifiedBy

`func (o *SearchedArtifact) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *SearchedArtifact) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *SearchedArtifact) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *SearchedArtifact) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetGroupId

`func (o *SearchedArtifact) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *SearchedArtifact) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *SearchedArtifact) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *SearchedArtifact) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


