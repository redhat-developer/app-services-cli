# RegistryCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Account** | Pointer to [**ObjectReference**](ObjectReference.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ExternalResourceId** | Pointer to **string** |  | [optional] 
**Registry** | Pointer to [**ObjectReference**](ObjectReference.md) |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewRegistryCredential

`func NewRegistryCredential() *RegistryCredential`

NewRegistryCredential instantiates a new RegistryCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryCredentialWithDefaults

`func NewRegistryCredentialWithDefaults() *RegistryCredential`

NewRegistryCredentialWithDefaults instantiates a new RegistryCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *RegistryCredential) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RegistryCredential) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RegistryCredential) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RegistryCredential) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *RegistryCredential) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegistryCredential) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegistryCredential) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RegistryCredential) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *RegistryCredential) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RegistryCredential) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RegistryCredential) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *RegistryCredential) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetAccount

`func (o *RegistryCredential) GetAccount() ObjectReference`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *RegistryCredential) GetAccountOk() (*ObjectReference, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *RegistryCredential) SetAccount(v ObjectReference)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *RegistryCredential) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RegistryCredential) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RegistryCredential) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RegistryCredential) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RegistryCredential) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExternalResourceId

`func (o *RegistryCredential) GetExternalResourceId() string`

GetExternalResourceId returns the ExternalResourceId field if non-nil, zero value otherwise.

### GetExternalResourceIdOk

`func (o *RegistryCredential) GetExternalResourceIdOk() (*string, bool)`

GetExternalResourceIdOk returns a tuple with the ExternalResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalResourceId

`func (o *RegistryCredential) SetExternalResourceId(v string)`

SetExternalResourceId sets ExternalResourceId field to given value.

### HasExternalResourceId

`func (o *RegistryCredential) HasExternalResourceId() bool`

HasExternalResourceId returns a boolean if a field has been set.

### GetRegistry

`func (o *RegistryCredential) GetRegistry() ObjectReference`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *RegistryCredential) GetRegistryOk() (*ObjectReference, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *RegistryCredential) SetRegistry(v ObjectReference)`

SetRegistry sets Registry field to given value.

### HasRegistry

`func (o *RegistryCredential) HasRegistry() bool`

HasRegistry returns a boolean if a field has been set.

### GetToken

`func (o *RegistryCredential) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RegistryCredential) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RegistryCredential) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RegistryCredential) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RegistryCredential) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RegistryCredential) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RegistryCredential) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RegistryCredential) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUsername

`func (o *RegistryCredential) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RegistryCredential) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RegistryCredential) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *RegistryCredential) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


