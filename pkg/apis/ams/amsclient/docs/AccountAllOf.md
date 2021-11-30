# AccountAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BanCode** | Pointer to **string** |  | [optional] 
**BanDescription** | Pointer to **string** |  | [optional] 
**Banned** | Pointer to **bool** |  | [optional] [default to false]
**Capabilities** | Pointer to [**[]Capability**](Capability.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to [**[]Label**](Label.md) |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to [**Organization**](Organization.md) |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**ServiceAccount** | Pointer to **bool** |  | [optional] [default to false]
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Username** | **string** |  | 

## Methods

### NewAccountAllOf

`func NewAccountAllOf(username string, ) *AccountAllOf`

NewAccountAllOf instantiates a new AccountAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountAllOfWithDefaults

`func NewAccountAllOfWithDefaults() *AccountAllOf`

NewAccountAllOfWithDefaults instantiates a new AccountAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBanCode

`func (o *AccountAllOf) GetBanCode() string`

GetBanCode returns the BanCode field if non-nil, zero value otherwise.

### GetBanCodeOk

`func (o *AccountAllOf) GetBanCodeOk() (*string, bool)`

GetBanCodeOk returns a tuple with the BanCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanCode

`func (o *AccountAllOf) SetBanCode(v string)`

SetBanCode sets BanCode field to given value.

### HasBanCode

`func (o *AccountAllOf) HasBanCode() bool`

HasBanCode returns a boolean if a field has been set.

### GetBanDescription

`func (o *AccountAllOf) GetBanDescription() string`

GetBanDescription returns the BanDescription field if non-nil, zero value otherwise.

### GetBanDescriptionOk

`func (o *AccountAllOf) GetBanDescriptionOk() (*string, bool)`

GetBanDescriptionOk returns a tuple with the BanDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanDescription

`func (o *AccountAllOf) SetBanDescription(v string)`

SetBanDescription sets BanDescription field to given value.

### HasBanDescription

`func (o *AccountAllOf) HasBanDescription() bool`

HasBanDescription returns a boolean if a field has been set.

### GetBanned

`func (o *AccountAllOf) GetBanned() bool`

GetBanned returns the Banned field if non-nil, zero value otherwise.

### GetBannedOk

`func (o *AccountAllOf) GetBannedOk() (*bool, bool)`

GetBannedOk returns a tuple with the Banned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanned

`func (o *AccountAllOf) SetBanned(v bool)`

SetBanned sets Banned field to given value.

### HasBanned

`func (o *AccountAllOf) HasBanned() bool`

HasBanned returns a boolean if a field has been set.

### GetCapabilities

`func (o *AccountAllOf) GetCapabilities() []Capability`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *AccountAllOf) GetCapabilitiesOk() (*[]Capability, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *AccountAllOf) SetCapabilities(v []Capability)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *AccountAllOf) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AccountAllOf) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccountAllOf) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccountAllOf) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AccountAllOf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEmail

`func (o *AccountAllOf) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AccountAllOf) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AccountAllOf) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AccountAllOf) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *AccountAllOf) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *AccountAllOf) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *AccountAllOf) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *AccountAllOf) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLabels

`func (o *AccountAllOf) GetLabels() []Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AccountAllOf) GetLabelsOk() (*[]Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AccountAllOf) SetLabels(v []Label)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AccountAllOf) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLastName

`func (o *AccountAllOf) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *AccountAllOf) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *AccountAllOf) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *AccountAllOf) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetOrganization

`func (o *AccountAllOf) GetOrganization() Organization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *AccountAllOf) GetOrganizationOk() (*Organization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *AccountAllOf) SetOrganization(v Organization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *AccountAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOrganizationId

`func (o *AccountAllOf) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *AccountAllOf) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *AccountAllOf) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *AccountAllOf) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetServiceAccount

`func (o *AccountAllOf) GetServiceAccount() bool`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *AccountAllOf) GetServiceAccountOk() (*bool, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *AccountAllOf) SetServiceAccount(v bool)`

SetServiceAccount sets ServiceAccount field to given value.

### HasServiceAccount

`func (o *AccountAllOf) HasServiceAccount() bool`

HasServiceAccount returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AccountAllOf) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AccountAllOf) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AccountAllOf) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AccountAllOf) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUsername

`func (o *AccountAllOf) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AccountAllOf) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AccountAllOf) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


