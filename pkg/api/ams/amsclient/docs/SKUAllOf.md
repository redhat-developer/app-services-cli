# SKUAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityZoneType** | Pointer to **string** |  | [optional] 
**Byoc** | **bool** |  | 
**Id** | Pointer to **string** |  | [optional] 
**ResourceName** | Pointer to **string** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**Resources** | Pointer to [**[]EphemeralResourceQuota**](EphemeralResourceQuota.md) |  | [optional] 

## Methods

### NewSKUAllOf

`func NewSKUAllOf(byoc bool, ) *SKUAllOf`

NewSKUAllOf instantiates a new SKUAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSKUAllOfWithDefaults

`func NewSKUAllOfWithDefaults() *SKUAllOf`

NewSKUAllOfWithDefaults instantiates a new SKUAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityZoneType

`func (o *SKUAllOf) GetAvailabilityZoneType() string`

GetAvailabilityZoneType returns the AvailabilityZoneType field if non-nil, zero value otherwise.

### GetAvailabilityZoneTypeOk

`func (o *SKUAllOf) GetAvailabilityZoneTypeOk() (*string, bool)`

GetAvailabilityZoneTypeOk returns a tuple with the AvailabilityZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneType

`func (o *SKUAllOf) SetAvailabilityZoneType(v string)`

SetAvailabilityZoneType sets AvailabilityZoneType field to given value.

### HasAvailabilityZoneType

`func (o *SKUAllOf) HasAvailabilityZoneType() bool`

HasAvailabilityZoneType returns a boolean if a field has been set.

### GetByoc

`func (o *SKUAllOf) GetByoc() bool`

GetByoc returns the Byoc field if non-nil, zero value otherwise.

### GetByocOk

`func (o *SKUAllOf) GetByocOk() (*bool, bool)`

GetByocOk returns a tuple with the Byoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByoc

`func (o *SKUAllOf) SetByoc(v bool)`

SetByoc sets Byoc field to given value.


### GetId

`func (o *SKUAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SKUAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SKUAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SKUAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResourceName

`func (o *SKUAllOf) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *SKUAllOf) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *SKUAllOf) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *SKUAllOf) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetResourceType

`func (o *SKUAllOf) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *SKUAllOf) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *SKUAllOf) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *SKUAllOf) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetResources

`func (o *SKUAllOf) GetResources() []EphemeralResourceQuota`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *SKUAllOf) GetResourcesOk() (*[]EphemeralResourceQuota, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *SKUAllOf) SetResources(v []EphemeralResourceQuota)`

SetResources sets Resources field to given value.

### HasResources

`func (o *SKUAllOf) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


