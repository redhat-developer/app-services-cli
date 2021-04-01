# ReservedResourceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityZoneType** | Pointer to **string** |  | [optional] 
**BillingModel** | Pointer to **string** |  | [optional] 
**Byoc** | **bool** |  | 
**Cluster** | Pointer to **bool** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ResourceName** | Pointer to **string** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**Subscription** | Pointer to [**ObjectReference**](ObjectReference.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewReservedResourceAllOf

`func NewReservedResourceAllOf(byoc bool, ) *ReservedResourceAllOf`

NewReservedResourceAllOf instantiates a new ReservedResourceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReservedResourceAllOfWithDefaults

`func NewReservedResourceAllOfWithDefaults() *ReservedResourceAllOf`

NewReservedResourceAllOfWithDefaults instantiates a new ReservedResourceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityZoneType

`func (o *ReservedResourceAllOf) GetAvailabilityZoneType() string`

GetAvailabilityZoneType returns the AvailabilityZoneType field if non-nil, zero value otherwise.

### GetAvailabilityZoneTypeOk

`func (o *ReservedResourceAllOf) GetAvailabilityZoneTypeOk() (*string, bool)`

GetAvailabilityZoneTypeOk returns a tuple with the AvailabilityZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneType

`func (o *ReservedResourceAllOf) SetAvailabilityZoneType(v string)`

SetAvailabilityZoneType sets AvailabilityZoneType field to given value.

### HasAvailabilityZoneType

`func (o *ReservedResourceAllOf) HasAvailabilityZoneType() bool`

HasAvailabilityZoneType returns a boolean if a field has been set.

### GetBillingModel

`func (o *ReservedResourceAllOf) GetBillingModel() string`

GetBillingModel returns the BillingModel field if non-nil, zero value otherwise.

### GetBillingModelOk

`func (o *ReservedResourceAllOf) GetBillingModelOk() (*string, bool)`

GetBillingModelOk returns a tuple with the BillingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingModel

`func (o *ReservedResourceAllOf) SetBillingModel(v string)`

SetBillingModel sets BillingModel field to given value.

### HasBillingModel

`func (o *ReservedResourceAllOf) HasBillingModel() bool`

HasBillingModel returns a boolean if a field has been set.

### GetByoc

`func (o *ReservedResourceAllOf) GetByoc() bool`

GetByoc returns the Byoc field if non-nil, zero value otherwise.

### GetByocOk

`func (o *ReservedResourceAllOf) GetByocOk() (*bool, bool)`

GetByocOk returns a tuple with the Byoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByoc

`func (o *ReservedResourceAllOf) SetByoc(v bool)`

SetByoc sets Byoc field to given value.


### GetCluster

`func (o *ReservedResourceAllOf) GetCluster() bool`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ReservedResourceAllOf) GetClusterOk() (*bool, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ReservedResourceAllOf) SetCluster(v bool)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ReservedResourceAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCount

`func (o *ReservedResourceAllOf) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ReservedResourceAllOf) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ReservedResourceAllOf) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ReservedResourceAllOf) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ReservedResourceAllOf) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ReservedResourceAllOf) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ReservedResourceAllOf) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ReservedResourceAllOf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetResourceName

`func (o *ReservedResourceAllOf) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *ReservedResourceAllOf) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *ReservedResourceAllOf) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *ReservedResourceAllOf) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetResourceType

`func (o *ReservedResourceAllOf) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ReservedResourceAllOf) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ReservedResourceAllOf) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ReservedResourceAllOf) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetSubscription

`func (o *ReservedResourceAllOf) GetSubscription() ObjectReference`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *ReservedResourceAllOf) GetSubscriptionOk() (*ObjectReference, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *ReservedResourceAllOf) SetSubscription(v ObjectReference)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *ReservedResourceAllOf) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ReservedResourceAllOf) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ReservedResourceAllOf) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ReservedResourceAllOf) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ReservedResourceAllOf) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


