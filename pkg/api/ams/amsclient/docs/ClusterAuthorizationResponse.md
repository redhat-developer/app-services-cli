# ClusterAuthorizationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allowed** | **bool** |  | 
**ExcessResources** | [**[]ExcessResource**](ExcessResource.md) |  | 
**OrganizationId** | Pointer to **string** |  | [optional] 
**Subscription** | Pointer to [**ObjectReference**](ObjectReference.md) |  | [optional] 

## Methods

### NewClusterAuthorizationResponse

`func NewClusterAuthorizationResponse(allowed bool, excessResources []ExcessResource, ) *ClusterAuthorizationResponse`

NewClusterAuthorizationResponse instantiates a new ClusterAuthorizationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterAuthorizationResponseWithDefaults

`func NewClusterAuthorizationResponseWithDefaults() *ClusterAuthorizationResponse`

NewClusterAuthorizationResponseWithDefaults instantiates a new ClusterAuthorizationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowed

`func (o *ClusterAuthorizationResponse) GetAllowed() bool`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *ClusterAuthorizationResponse) GetAllowedOk() (*bool, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *ClusterAuthorizationResponse) SetAllowed(v bool)`

SetAllowed sets Allowed field to given value.


### GetExcessResources

`func (o *ClusterAuthorizationResponse) GetExcessResources() []ExcessResource`

GetExcessResources returns the ExcessResources field if non-nil, zero value otherwise.

### GetExcessResourcesOk

`func (o *ClusterAuthorizationResponse) GetExcessResourcesOk() (*[]ExcessResource, bool)`

GetExcessResourcesOk returns a tuple with the ExcessResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcessResources

`func (o *ClusterAuthorizationResponse) SetExcessResources(v []ExcessResource)`

SetExcessResources sets ExcessResources field to given value.


### GetOrganizationId

`func (o *ClusterAuthorizationResponse) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ClusterAuthorizationResponse) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ClusterAuthorizationResponse) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ClusterAuthorizationResponse) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetSubscription

`func (o *ClusterAuthorizationResponse) GetSubscription() ObjectReference`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *ClusterAuthorizationResponse) GetSubscriptionOk() (*ObjectReference, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *ClusterAuthorizationResponse) SetSubscription(v ObjectReference)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *ClusterAuthorizationResponse) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


