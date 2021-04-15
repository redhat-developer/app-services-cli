# VersionMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Collections** | Pointer to [**[]ObjectReference**](ObjectReference.md) |  | [optional] 

## Methods

### NewVersionMetadata

`func NewVersionMetadata() *VersionMetadata`

NewVersionMetadata instantiates a new VersionMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionMetadataWithDefaults

`func NewVersionMetadataWithDefaults() *VersionMetadata`

NewVersionMetadataWithDefaults instantiates a new VersionMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VersionMetadata) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VersionMetadata) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VersionMetadata) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VersionMetadata) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *VersionMetadata) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *VersionMetadata) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *VersionMetadata) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *VersionMetadata) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetHref

`func (o *VersionMetadata) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *VersionMetadata) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *VersionMetadata) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *VersionMetadata) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetCollections

`func (o *VersionMetadata) GetCollections() []ObjectReference`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *VersionMetadata) GetCollectionsOk() (*[]ObjectReference, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *VersionMetadata) SetCollections(v []ObjectReference)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *VersionMetadata) HasCollections() bool`

HasCollections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


