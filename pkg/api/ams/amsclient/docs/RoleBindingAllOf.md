# RoleBindingAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**ObjectReference**](ObjectReference.md) |  | [optional] 
**ConfigManaged** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Organization** | Pointer to [**ObjectReference**](ObjectReference.md) |  | [optional] 
**Role** | Pointer to [**ObjectReference**](ObjectReference.md) |  | [optional] 
**Subscription** | Pointer to [**ObjectReference**](ObjectReference.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRoleBindingAllOf

`func NewRoleBindingAllOf() *RoleBindingAllOf`

NewRoleBindingAllOf instantiates a new RoleBindingAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleBindingAllOfWithDefaults

`func NewRoleBindingAllOfWithDefaults() *RoleBindingAllOf`

NewRoleBindingAllOfWithDefaults instantiates a new RoleBindingAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *RoleBindingAllOf) GetAccount() ObjectReference`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *RoleBindingAllOf) GetAccountOk() (*ObjectReference, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *RoleBindingAllOf) SetAccount(v ObjectReference)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *RoleBindingAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetConfigManaged

`func (o *RoleBindingAllOf) GetConfigManaged() bool`

GetConfigManaged returns the ConfigManaged field if non-nil, zero value otherwise.

### GetConfigManagedOk

`func (o *RoleBindingAllOf) GetConfigManagedOk() (*bool, bool)`

GetConfigManagedOk returns a tuple with the ConfigManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigManaged

`func (o *RoleBindingAllOf) SetConfigManaged(v bool)`

SetConfigManaged sets ConfigManaged field to given value.

### HasConfigManaged

`func (o *RoleBindingAllOf) HasConfigManaged() bool`

HasConfigManaged returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RoleBindingAllOf) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RoleBindingAllOf) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RoleBindingAllOf) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RoleBindingAllOf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetOrganization

`func (o *RoleBindingAllOf) GetOrganization() ObjectReference`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *RoleBindingAllOf) GetOrganizationOk() (*ObjectReference, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *RoleBindingAllOf) SetOrganization(v ObjectReference)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *RoleBindingAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetRole

`func (o *RoleBindingAllOf) GetRole() ObjectReference`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *RoleBindingAllOf) GetRoleOk() (*ObjectReference, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *RoleBindingAllOf) SetRole(v ObjectReference)`

SetRole sets Role field to given value.

### HasRole

`func (o *RoleBindingAllOf) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSubscription

`func (o *RoleBindingAllOf) GetSubscription() ObjectReference`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *RoleBindingAllOf) GetSubscriptionOk() (*ObjectReference, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *RoleBindingAllOf) SetSubscription(v ObjectReference)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *RoleBindingAllOf) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetType

`func (o *RoleBindingAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoleBindingAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoleBindingAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RoleBindingAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RoleBindingAllOf) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RoleBindingAllOf) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RoleBindingAllOf) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RoleBindingAllOf) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


