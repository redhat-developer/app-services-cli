# ReservedResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
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

### NewReservedResource

`func NewReservedResource(byoc bool, ) *ReservedResource`

NewReservedResource instantiates a new ReservedResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReservedResourceWithDefaults

`func NewReservedResourceWithDefaults() *ReservedResource`

NewReservedResourceWithDefaults instantiates a new ReservedResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *ReservedResource) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ReservedResource) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ReservedResource) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ReservedResource) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *ReservedResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReservedResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReservedResource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReservedResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *ReservedResource) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ReservedResource) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ReservedResource) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ReservedResource) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetAvailabilityZoneType

`func (o *ReservedResource) GetAvailabilityZoneType() string`

GetAvailabilityZoneType returns the AvailabilityZoneType field if non-nil, zero value otherwise.

### GetAvailabilityZoneTypeOk

`func (o *ReservedResource) GetAvailabilityZoneTypeOk() (*string, bool)`

GetAvailabilityZoneTypeOk returns a tuple with the AvailabilityZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneType

`func (o *ReservedResource) SetAvailabilityZoneType(v string)`

SetAvailabilityZoneType sets AvailabilityZoneType field to given value.

### HasAvailabilityZoneType

`func (o *ReservedResource) HasAvailabilityZoneType() bool`

HasAvailabilityZoneType returns a boolean if a field has been set.

### GetBillingModel

`func (o *ReservedResource) GetBillingModel() string`

GetBillingModel returns the BillingModel field if non-nil, zero value otherwise.

### GetBillingModelOk

`func (o *ReservedResource) GetBillingModelOk() (*string, bool)`

GetBillingModelOk returns a tuple with the BillingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingModel

`func (o *ReservedResource) SetBillingModel(v string)`

SetBillingModel sets BillingModel field to given value.

### HasBillingModel

`func (o *ReservedResource) HasBillingModel() bool`

HasBillingModel returns a boolean if a field has been set.

### GetByoc

`func (o *ReservedResource) GetByoc() bool`

GetByoc returns the Byoc field if non-nil, zero value otherwise.

### GetByocOk

`func (o *ReservedResource) GetByocOk() (*bool, bool)`

GetByocOk returns a tuple with the Byoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByoc

`func (o *ReservedResource) SetByoc(v bool)`

SetByoc sets Byoc field to given value.


### GetCluster

`func (o *ReservedResource) GetCluster() bool`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ReservedResource) GetClusterOk() (*bool, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ReservedResource) SetCluster(v bool)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ReservedResource) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCount

`func (o *ReservedResource) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ReservedResource) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ReservedResource) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ReservedResource) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ReservedResource) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ReservedResource) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ReservedResource) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ReservedResource) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetResourceName

`func (o *ReservedResource) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *ReservedResource) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *ReservedResource) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *ReservedResource) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetResourceType

`func (o *ReservedResource) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ReservedResource) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ReservedResource) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ReservedResource) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetSubscription

`func (o *ReservedResource) GetSubscription() ObjectReference`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *ReservedResource) GetSubscriptionOk() (*ObjectReference, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *ReservedResource) SetSubscription(v ObjectReference)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *ReservedResource) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ReservedResource) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ReservedResource) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ReservedResource) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ReservedResource) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


