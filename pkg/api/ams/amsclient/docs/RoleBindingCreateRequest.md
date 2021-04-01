# RoleBindingCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**ConfigManaged** | Pointer to **bool** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**RoleId** | **string** |  | 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewRoleBindingCreateRequest

`func NewRoleBindingCreateRequest(accountId string, roleId string, type_ string, ) *RoleBindingCreateRequest`

NewRoleBindingCreateRequest instantiates a new RoleBindingCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleBindingCreateRequestWithDefaults

`func NewRoleBindingCreateRequestWithDefaults() *RoleBindingCreateRequest`

NewRoleBindingCreateRequestWithDefaults instantiates a new RoleBindingCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *RoleBindingCreateRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *RoleBindingCreateRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *RoleBindingCreateRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetConfigManaged

`func (o *RoleBindingCreateRequest) GetConfigManaged() bool`

GetConfigManaged returns the ConfigManaged field if non-nil, zero value otherwise.

### GetConfigManagedOk

`func (o *RoleBindingCreateRequest) GetConfigManagedOk() (*bool, bool)`

GetConfigManagedOk returns a tuple with the ConfigManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigManaged

`func (o *RoleBindingCreateRequest) SetConfigManaged(v bool)`

SetConfigManaged sets ConfigManaged field to given value.

### HasConfigManaged

`func (o *RoleBindingCreateRequest) HasConfigManaged() bool`

HasConfigManaged returns a boolean if a field has been set.

### GetOrganizationId

`func (o *RoleBindingCreateRequest) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *RoleBindingCreateRequest) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *RoleBindingCreateRequest) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *RoleBindingCreateRequest) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetRoleId

`func (o *RoleBindingCreateRequest) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *RoleBindingCreateRequest) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *RoleBindingCreateRequest) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.


### GetSubscriptionId

`func (o *RoleBindingCreateRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *RoleBindingCreateRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *RoleBindingCreateRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *RoleBindingCreateRequest) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetType

`func (o *RoleBindingCreateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoleBindingCreateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoleBindingCreateRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


