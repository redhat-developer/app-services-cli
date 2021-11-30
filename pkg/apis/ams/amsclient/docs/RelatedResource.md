# RelatedResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**AvailabilityZoneType** | **string** |  | 
**BillingModel** | **string** |  | 
**Byoc** | **string** |  | 
**CloudProvider** | **string** |  | 
**Cost** | **int32** |  | 
**Product** | **string** |  | 
**ProductId** | Pointer to **string** |  | [optional] 
**ResourceName** | Pointer to **string** |  | [optional] 
**ResourceType** | **string** |  | 

## Methods

### NewRelatedResource

`func NewRelatedResource(availabilityZoneType string, billingModel string, byoc string, cloudProvider string, cost int32, product string, resourceType string, ) *RelatedResource`

NewRelatedResource instantiates a new RelatedResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelatedResourceWithDefaults

`func NewRelatedResourceWithDefaults() *RelatedResource`

NewRelatedResourceWithDefaults instantiates a new RelatedResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *RelatedResource) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RelatedResource) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RelatedResource) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RelatedResource) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *RelatedResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RelatedResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RelatedResource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RelatedResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *RelatedResource) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RelatedResource) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RelatedResource) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *RelatedResource) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetAvailabilityZoneType

`func (o *RelatedResource) GetAvailabilityZoneType() string`

GetAvailabilityZoneType returns the AvailabilityZoneType field if non-nil, zero value otherwise.

### GetAvailabilityZoneTypeOk

`func (o *RelatedResource) GetAvailabilityZoneTypeOk() (*string, bool)`

GetAvailabilityZoneTypeOk returns a tuple with the AvailabilityZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneType

`func (o *RelatedResource) SetAvailabilityZoneType(v string)`

SetAvailabilityZoneType sets AvailabilityZoneType field to given value.


### GetBillingModel

`func (o *RelatedResource) GetBillingModel() string`

GetBillingModel returns the BillingModel field if non-nil, zero value otherwise.

### GetBillingModelOk

`func (o *RelatedResource) GetBillingModelOk() (*string, bool)`

GetBillingModelOk returns a tuple with the BillingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingModel

`func (o *RelatedResource) SetBillingModel(v string)`

SetBillingModel sets BillingModel field to given value.


### GetByoc

`func (o *RelatedResource) GetByoc() string`

GetByoc returns the Byoc field if non-nil, zero value otherwise.

### GetByocOk

`func (o *RelatedResource) GetByocOk() (*string, bool)`

GetByocOk returns a tuple with the Byoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByoc

`func (o *RelatedResource) SetByoc(v string)`

SetByoc sets Byoc field to given value.


### GetCloudProvider

`func (o *RelatedResource) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *RelatedResource) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *RelatedResource) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetCost

`func (o *RelatedResource) GetCost() int32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *RelatedResource) GetCostOk() (*int32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *RelatedResource) SetCost(v int32)`

SetCost sets Cost field to given value.


### GetProduct

`func (o *RelatedResource) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *RelatedResource) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *RelatedResource) SetProduct(v string)`

SetProduct sets Product field to given value.


### GetProductId

`func (o *RelatedResource) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *RelatedResource) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *RelatedResource) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *RelatedResource) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetResourceName

`func (o *RelatedResource) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *RelatedResource) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *RelatedResource) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *RelatedResource) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetResourceType

`func (o *RelatedResource) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *RelatedResource) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *RelatedResource) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


