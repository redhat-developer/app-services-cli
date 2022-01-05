# Registry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**CloudAlias** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OrgName** | Pointer to **string** |  | [optional] 
**TeamName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewRegistry

`func NewRegistry() *Registry`

NewRegistry instantiates a new Registry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryWithDefaults

`func NewRegistryWithDefaults() *Registry`

NewRegistryWithDefaults instantiates a new Registry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *Registry) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Registry) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Registry) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Registry) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *Registry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Registry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Registry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Registry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *Registry) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Registry) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Registry) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Registry) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetCloudAlias

`func (o *Registry) GetCloudAlias() bool`

GetCloudAlias returns the CloudAlias field if non-nil, zero value otherwise.

### GetCloudAliasOk

`func (o *Registry) GetCloudAliasOk() (*bool, bool)`

GetCloudAliasOk returns a tuple with the CloudAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAlias

`func (o *Registry) SetCloudAlias(v bool)`

SetCloudAlias sets CloudAlias field to given value.

### HasCloudAlias

`func (o *Registry) HasCloudAlias() bool`

HasCloudAlias returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Registry) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Registry) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Registry) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Registry) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetName

`func (o *Registry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Registry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Registry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Registry) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgName

`func (o *Registry) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *Registry) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *Registry) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *Registry) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetTeamName

`func (o *Registry) GetTeamName() string`

GetTeamName returns the TeamName field if non-nil, zero value otherwise.

### GetTeamNameOk

`func (o *Registry) GetTeamNameOk() (*string, bool)`

GetTeamNameOk returns a tuple with the TeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamName

`func (o *Registry) SetTeamName(v string)`

SetTeamName sets TeamName field to given value.

### HasTeamName

`func (o *Registry) HasTeamName() bool`

HasTeamName returns a boolean if a field has been set.

### GetType

`func (o *Registry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Registry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Registry) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Registry) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Registry) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Registry) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Registry) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Registry) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *Registry) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Registry) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Registry) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Registry) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


