# ResourceQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Allowed** | **int32** |  | 
**AvailabilityZoneType** | Pointer to **string** |  | [optional] 
**Byoc** | **bool** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**ResourceName** | **string** |  | 
**ResourceType** | **string** |  | 
**Sku** | Pointer to **string** |  | [optional] 
**SkuCount** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewResourceQuota

`func NewResourceQuota(allowed int32, byoc bool, resourceName string, resourceType string, ) *ResourceQuota`

NewResourceQuota instantiates a new ResourceQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceQuotaWithDefaults

`func NewResourceQuotaWithDefaults() *ResourceQuota`

NewResourceQuotaWithDefaults instantiates a new ResourceQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *ResourceQuota) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ResourceQuota) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ResourceQuota) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ResourceQuota) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *ResourceQuota) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceQuota) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceQuota) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResourceQuota) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *ResourceQuota) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ResourceQuota) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ResourceQuota) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ResourceQuota) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetAllowed

`func (o *ResourceQuota) GetAllowed() int32`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *ResourceQuota) GetAllowedOk() (*int32, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *ResourceQuota) SetAllowed(v int32)`

SetAllowed sets Allowed field to given value.


### GetAvailabilityZoneType

`func (o *ResourceQuota) GetAvailabilityZoneType() string`

GetAvailabilityZoneType returns the AvailabilityZoneType field if non-nil, zero value otherwise.

### GetAvailabilityZoneTypeOk

`func (o *ResourceQuota) GetAvailabilityZoneTypeOk() (*string, bool)`

GetAvailabilityZoneTypeOk returns a tuple with the AvailabilityZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneType

`func (o *ResourceQuota) SetAvailabilityZoneType(v string)`

SetAvailabilityZoneType sets AvailabilityZoneType field to given value.

### HasAvailabilityZoneType

`func (o *ResourceQuota) HasAvailabilityZoneType() bool`

HasAvailabilityZoneType returns a boolean if a field has been set.

### GetByoc

`func (o *ResourceQuota) GetByoc() bool`

GetByoc returns the Byoc field if non-nil, zero value otherwise.

### GetByocOk

`func (o *ResourceQuota) GetByocOk() (*bool, bool)`

GetByocOk returns a tuple with the Byoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByoc

`func (o *ResourceQuota) SetByoc(v bool)`

SetByoc sets Byoc field to given value.


### GetCreatedAt

`func (o *ResourceQuota) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResourceQuota) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResourceQuota) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ResourceQuota) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetOrganizationId

`func (o *ResourceQuota) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ResourceQuota) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ResourceQuota) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ResourceQuota) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetResourceName

`func (o *ResourceQuota) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *ResourceQuota) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *ResourceQuota) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetResourceType

`func (o *ResourceQuota) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ResourceQuota) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ResourceQuota) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetSku

`func (o *ResourceQuota) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *ResourceQuota) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *ResourceQuota) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *ResourceQuota) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetSkuCount

`func (o *ResourceQuota) GetSkuCount() int32`

GetSkuCount returns the SkuCount field if non-nil, zero value otherwise.

### GetSkuCountOk

`func (o *ResourceQuota) GetSkuCountOk() (*int32, bool)`

GetSkuCountOk returns a tuple with the SkuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuCount

`func (o *ResourceQuota) SetSkuCount(v int32)`

SetSkuCount sets SkuCount field to given value.

### HasSkuCount

`func (o *ResourceQuota) HasSkuCount() bool`

HasSkuCount returns a boolean if a field has been set.

### GetType

`func (o *ResourceQuota) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourceQuota) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourceQuota) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResourceQuota) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ResourceQuota) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ResourceQuota) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ResourceQuota) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ResourceQuota) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


