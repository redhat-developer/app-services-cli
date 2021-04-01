# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
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

### NewAccount

`func NewAccount(username string, ) *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *Account) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Account) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Account) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Account) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *Account) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Account) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Account) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Account) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *Account) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Account) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Account) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Account) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetBanCode

`func (o *Account) GetBanCode() string`

GetBanCode returns the BanCode field if non-nil, zero value otherwise.

### GetBanCodeOk

`func (o *Account) GetBanCodeOk() (*string, bool)`

GetBanCodeOk returns a tuple with the BanCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanCode

`func (o *Account) SetBanCode(v string)`

SetBanCode sets BanCode field to given value.

### HasBanCode

`func (o *Account) HasBanCode() bool`

HasBanCode returns a boolean if a field has been set.

### GetBanDescription

`func (o *Account) GetBanDescription() string`

GetBanDescription returns the BanDescription field if non-nil, zero value otherwise.

### GetBanDescriptionOk

`func (o *Account) GetBanDescriptionOk() (*string, bool)`

GetBanDescriptionOk returns a tuple with the BanDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanDescription

`func (o *Account) SetBanDescription(v string)`

SetBanDescription sets BanDescription field to given value.

### HasBanDescription

`func (o *Account) HasBanDescription() bool`

HasBanDescription returns a boolean if a field has been set.

### GetBanned

`func (o *Account) GetBanned() bool`

GetBanned returns the Banned field if non-nil, zero value otherwise.

### GetBannedOk

`func (o *Account) GetBannedOk() (*bool, bool)`

GetBannedOk returns a tuple with the Banned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanned

`func (o *Account) SetBanned(v bool)`

SetBanned sets Banned field to given value.

### HasBanned

`func (o *Account) HasBanned() bool`

HasBanned returns a boolean if a field has been set.

### GetCapabilities

`func (o *Account) GetCapabilities() []Capability`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *Account) GetCapabilitiesOk() (*[]Capability, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *Account) SetCapabilities(v []Capability)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *Account) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Account) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Account) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Account) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Account) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEmail

`func (o *Account) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Account) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Account) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Account) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *Account) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Account) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Account) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *Account) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLabels

`func (o *Account) GetLabels() []Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Account) GetLabelsOk() (*[]Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Account) SetLabels(v []Label)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Account) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLastName

`func (o *Account) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Account) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Account) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *Account) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetOrganization

`func (o *Account) GetOrganization() Organization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *Account) GetOrganizationOk() (*Organization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *Account) SetOrganization(v Organization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *Account) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOrganizationId

`func (o *Account) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Account) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Account) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *Account) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetServiceAccount

`func (o *Account) GetServiceAccount() bool`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *Account) GetServiceAccountOk() (*bool, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *Account) SetServiceAccount(v bool)`

SetServiceAccount sets ServiceAccount field to given value.

### HasServiceAccount

`func (o *Account) HasServiceAccount() bool`

HasServiceAccount returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Account) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Account) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Account) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Account) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUsername

`func (o *Account) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Account) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Account) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


