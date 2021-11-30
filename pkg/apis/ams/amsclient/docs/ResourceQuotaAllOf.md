# ResourceQuotaAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewResourceQuotaAllOf

`func NewResourceQuotaAllOf(allowed int32, byoc bool, resourceName string, resourceType string, ) *ResourceQuotaAllOf`

NewResourceQuotaAllOf instantiates a new ResourceQuotaAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceQuotaAllOfWithDefaults

`func NewResourceQuotaAllOfWithDefaults() *ResourceQuotaAllOf`

NewResourceQuotaAllOfWithDefaults instantiates a new ResourceQuotaAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowed

`func (o *ResourceQuotaAllOf) GetAllowed() int32`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *ResourceQuotaAllOf) GetAllowedOk() (*int32, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *ResourceQuotaAllOf) SetAllowed(v int32)`

SetAllowed sets Allowed field to given value.


### GetAvailabilityZoneType

`func (o *ResourceQuotaAllOf) GetAvailabilityZoneType() string`

GetAvailabilityZoneType returns the AvailabilityZoneType field if non-nil, zero value otherwise.

### GetAvailabilityZoneTypeOk

`func (o *ResourceQuotaAllOf) GetAvailabilityZoneTypeOk() (*string, bool)`

GetAvailabilityZoneTypeOk returns a tuple with the AvailabilityZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneType

`func (o *ResourceQuotaAllOf) SetAvailabilityZoneType(v string)`

SetAvailabilityZoneType sets AvailabilityZoneType field to given value.

### HasAvailabilityZoneType

`func (o *ResourceQuotaAllOf) HasAvailabilityZoneType() bool`

HasAvailabilityZoneType returns a boolean if a field has been set.

### GetByoc

`func (o *ResourceQuotaAllOf) GetByoc() bool`

GetByoc returns the Byoc field if non-nil, zero value otherwise.

### GetByocOk

`func (o *ResourceQuotaAllOf) GetByocOk() (*bool, bool)`

GetByocOk returns a tuple with the Byoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByoc

`func (o *ResourceQuotaAllOf) SetByoc(v bool)`

SetByoc sets Byoc field to given value.


### GetCreatedAt

`func (o *ResourceQuotaAllOf) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResourceQuotaAllOf) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResourceQuotaAllOf) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ResourceQuotaAllOf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetOrganizationId

`func (o *ResourceQuotaAllOf) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ResourceQuotaAllOf) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ResourceQuotaAllOf) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ResourceQuotaAllOf) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetResourceName

`func (o *ResourceQuotaAllOf) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *ResourceQuotaAllOf) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *ResourceQuotaAllOf) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetResourceType

`func (o *ResourceQuotaAllOf) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ResourceQuotaAllOf) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ResourceQuotaAllOf) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetSku

`func (o *ResourceQuotaAllOf) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *ResourceQuotaAllOf) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *ResourceQuotaAllOf) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *ResourceQuotaAllOf) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetSkuCount

`func (o *ResourceQuotaAllOf) GetSkuCount() int32`

GetSkuCount returns the SkuCount field if non-nil, zero value otherwise.

### GetSkuCountOk

`func (o *ResourceQuotaAllOf) GetSkuCountOk() (*int32, bool)`

GetSkuCountOk returns a tuple with the SkuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuCount

`func (o *ResourceQuotaAllOf) SetSkuCount(v int32)`

SetSkuCount sets SkuCount field to given value.

### HasSkuCount

`func (o *ResourceQuotaAllOf) HasSkuCount() bool`

HasSkuCount returns a boolean if a field has been set.

### GetType

`func (o *ResourceQuotaAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourceQuotaAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourceQuotaAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResourceQuotaAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ResourceQuotaAllOf) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ResourceQuotaAllOf) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ResourceQuotaAllOf) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ResourceQuotaAllOf) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


