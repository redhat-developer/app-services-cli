# SKU

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**AvailabilityZoneType** | Pointer to **string** |  | [optional] 
**Byoc** | **bool** |  | 
**ResourceName** | Pointer to **string** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**Resources** | Pointer to [**[]EphemeralResourceQuota**](EphemeralResourceQuota.md) |  | [optional] 

## Methods

### NewSKU

`func NewSKU(byoc bool, ) *SKU`

NewSKU instantiates a new SKU object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSKUWithDefaults

`func NewSKUWithDefaults() *SKU`

NewSKUWithDefaults instantiates a new SKU object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *SKU) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SKU) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SKU) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *SKU) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *SKU) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SKU) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SKU) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SKU) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *SKU) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SKU) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SKU) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SKU) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetAvailabilityZoneType

`func (o *SKU) GetAvailabilityZoneType() string`

GetAvailabilityZoneType returns the AvailabilityZoneType field if non-nil, zero value otherwise.

### GetAvailabilityZoneTypeOk

`func (o *SKU) GetAvailabilityZoneTypeOk() (*string, bool)`

GetAvailabilityZoneTypeOk returns a tuple with the AvailabilityZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneType

`func (o *SKU) SetAvailabilityZoneType(v string)`

SetAvailabilityZoneType sets AvailabilityZoneType field to given value.

### HasAvailabilityZoneType

`func (o *SKU) HasAvailabilityZoneType() bool`

HasAvailabilityZoneType returns a boolean if a field has been set.

### GetByoc

`func (o *SKU) GetByoc() bool`

GetByoc returns the Byoc field if non-nil, zero value otherwise.

### GetByocOk

`func (o *SKU) GetByocOk() (*bool, bool)`

GetByocOk returns a tuple with the Byoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByoc

`func (o *SKU) SetByoc(v bool)`

SetByoc sets Byoc field to given value.


### GetResourceName

`func (o *SKU) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *SKU) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *SKU) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *SKU) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetResourceType

`func (o *SKU) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *SKU) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *SKU) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *SKU) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetResources

`func (o *SKU) GetResources() []EphemeralResourceQuota`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *SKU) GetResourcesOk() (*[]EphemeralResourceQuota, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *SKU) SetResources(v []EphemeralResourceQuota)`

SetResources sets Resources field to given value.

### HasResources

`func (o *SKU) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


