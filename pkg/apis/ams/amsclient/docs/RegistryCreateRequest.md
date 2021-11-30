# RegistryCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudAlias** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | 
**OrgName** | Pointer to **string** |  | [optional] 
**TeamName** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**Url** | **string** |  | 

## Methods

### NewRegistryCreateRequest

`func NewRegistryCreateRequest(name string, type_ string, url string, ) *RegistryCreateRequest`

NewRegistryCreateRequest instantiates a new RegistryCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryCreateRequestWithDefaults

`func NewRegistryCreateRequestWithDefaults() *RegistryCreateRequest`

NewRegistryCreateRequestWithDefaults instantiates a new RegistryCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudAlias

`func (o *RegistryCreateRequest) GetCloudAlias() bool`

GetCloudAlias returns the CloudAlias field if non-nil, zero value otherwise.

### GetCloudAliasOk

`func (o *RegistryCreateRequest) GetCloudAliasOk() (*bool, bool)`

GetCloudAliasOk returns a tuple with the CloudAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAlias

`func (o *RegistryCreateRequest) SetCloudAlias(v bool)`

SetCloudAlias sets CloudAlias field to given value.

### HasCloudAlias

`func (o *RegistryCreateRequest) HasCloudAlias() bool`

HasCloudAlias returns a boolean if a field has been set.

### GetName

`func (o *RegistryCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegistryCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegistryCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOrgName

`func (o *RegistryCreateRequest) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *RegistryCreateRequest) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *RegistryCreateRequest) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *RegistryCreateRequest) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetTeamName

`func (o *RegistryCreateRequest) GetTeamName() string`

GetTeamName returns the TeamName field if non-nil, zero value otherwise.

### GetTeamNameOk

`func (o *RegistryCreateRequest) GetTeamNameOk() (*string, bool)`

GetTeamNameOk returns a tuple with the TeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamName

`func (o *RegistryCreateRequest) SetTeamName(v string)`

SetTeamName sets TeamName field to given value.

### HasTeamName

`func (o *RegistryCreateRequest) HasTeamName() bool`

HasTeamName returns a boolean if a field has been set.

### GetType

`func (o *RegistryCreateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RegistryCreateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RegistryCreateRequest) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *RegistryCreateRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RegistryCreateRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RegistryCreateRequest) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


