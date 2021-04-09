# RoleBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Account** | Pointer to [**ObjectReference**](ObjectReference.md) |  | [optional] 
**ConfigManaged** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Organization** | Pointer to [**ObjectReference**](ObjectReference.md) |  | [optional] 
**Role** | Pointer to [**ObjectReference**](ObjectReference.md) |  | [optional] 
**Subscription** | Pointer to [**ObjectReference**](ObjectReference.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRoleBinding

`func NewRoleBinding() *RoleBinding`

NewRoleBinding instantiates a new RoleBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleBindingWithDefaults

`func NewRoleBindingWithDefaults() *RoleBinding`

NewRoleBindingWithDefaults instantiates a new RoleBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *RoleBinding) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RoleBinding) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RoleBinding) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RoleBinding) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *RoleBinding) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoleBinding) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoleBinding) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RoleBinding) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *RoleBinding) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RoleBinding) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RoleBinding) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *RoleBinding) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetAccount

`func (o *RoleBinding) GetAccount() ObjectReference`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *RoleBinding) GetAccountOk() (*ObjectReference, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *RoleBinding) SetAccount(v ObjectReference)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *RoleBinding) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetConfigManaged

`func (o *RoleBinding) GetConfigManaged() bool`

GetConfigManaged returns the ConfigManaged field if non-nil, zero value otherwise.

### GetConfigManagedOk

`func (o *RoleBinding) GetConfigManagedOk() (*bool, bool)`

GetConfigManagedOk returns a tuple with the ConfigManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigManaged

`func (o *RoleBinding) SetConfigManaged(v bool)`

SetConfigManaged sets ConfigManaged field to given value.

### HasConfigManaged

`func (o *RoleBinding) HasConfigManaged() bool`

HasConfigManaged returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RoleBinding) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RoleBinding) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RoleBinding) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RoleBinding) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetOrganization

`func (o *RoleBinding) GetOrganization() ObjectReference`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *RoleBinding) GetOrganizationOk() (*ObjectReference, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *RoleBinding) SetOrganization(v ObjectReference)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *RoleBinding) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetRole

`func (o *RoleBinding) GetRole() ObjectReference`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *RoleBinding) GetRoleOk() (*ObjectReference, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *RoleBinding) SetRole(v ObjectReference)`

SetRole sets Role field to given value.

### HasRole

`func (o *RoleBinding) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSubscription

`func (o *RoleBinding) GetSubscription() ObjectReference`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *RoleBinding) GetSubscriptionOk() (*ObjectReference, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *RoleBinding) SetSubscription(v ObjectReference)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *RoleBinding) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetType

`func (o *RoleBinding) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoleBinding) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoleBinding) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RoleBinding) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RoleBinding) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RoleBinding) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RoleBinding) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RoleBinding) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


