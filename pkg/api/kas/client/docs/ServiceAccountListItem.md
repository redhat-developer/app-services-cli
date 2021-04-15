# ServiceAccountListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | server generated unique id of the service account | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**ClientID** | Pointer to **string** | client id of the service account | [optional] 
**Name** | Pointer to **string** | name of the service account | [optional] 
**Owner** | Pointer to **string** | owner of the service account | [optional] 
**CreatedAt** | Pointer to **time.Time** | service account creation timestamp | [optional] 
**Description** | Pointer to **string** | description of the service account | [optional] 

## Methods

### NewServiceAccountListItem

`func NewServiceAccountListItem() *ServiceAccountListItem`

NewServiceAccountListItem instantiates a new ServiceAccountListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountListItemWithDefaults

`func NewServiceAccountListItemWithDefaults() *ServiceAccountListItem`

NewServiceAccountListItemWithDefaults instantiates a new ServiceAccountListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceAccountListItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceAccountListItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceAccountListItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceAccountListItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *ServiceAccountListItem) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ServiceAccountListItem) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ServiceAccountListItem) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ServiceAccountListItem) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetHref

`func (o *ServiceAccountListItem) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ServiceAccountListItem) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ServiceAccountListItem) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ServiceAccountListItem) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetClientID

`func (o *ServiceAccountListItem) GetClientID() string`

GetClientID returns the ClientID field if non-nil, zero value otherwise.

### GetClientIDOk

`func (o *ServiceAccountListItem) GetClientIDOk() (*string, bool)`

GetClientIDOk returns a tuple with the ClientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientID

`func (o *ServiceAccountListItem) SetClientID(v string)`

SetClientID sets ClientID field to given value.

### HasClientID

`func (o *ServiceAccountListItem) HasClientID() bool`

HasClientID returns a boolean if a field has been set.

### GetName

`func (o *ServiceAccountListItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceAccountListItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceAccountListItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceAccountListItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *ServiceAccountListItem) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ServiceAccountListItem) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ServiceAccountListItem) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ServiceAccountListItem) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ServiceAccountListItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceAccountListItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceAccountListItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ServiceAccountListItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *ServiceAccountListItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceAccountListItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceAccountListItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceAccountListItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


