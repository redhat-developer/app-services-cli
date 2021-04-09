# ClusterAuthorizationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountUsername** | **string** |  | 
**AvailabilityZone** | Pointer to **string** |  | [optional] 
**Byoc** | Pointer to **bool** |  | [optional] 
**CloudAccountId** | Pointer to **string** |  | [optional] 
**CloudProviderId** | Pointer to **string** |  | [optional] 
**ClusterId** | **string** |  | 
**Disconnected** | Pointer to **bool** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ExternalClusterId** | Pointer to **string** |  | [optional] 
**Managed** | Pointer to **bool** |  | [optional] 
**ProductCategory** | Pointer to **string** |  | [optional] 
**ProductId** | Pointer to **string** |  | [optional] [default to "osd"]
**Reserve** | Pointer to **bool** |  | [optional] 
**Resources** | Pointer to [**[]ReservedResource**](ReservedResource.md) |  | [optional] 

## Methods

### NewClusterAuthorizationRequest

`func NewClusterAuthorizationRequest(accountUsername string, clusterId string, ) *ClusterAuthorizationRequest`

NewClusterAuthorizationRequest instantiates a new ClusterAuthorizationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterAuthorizationRequestWithDefaults

`func NewClusterAuthorizationRequestWithDefaults() *ClusterAuthorizationRequest`

NewClusterAuthorizationRequestWithDefaults instantiates a new ClusterAuthorizationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountUsername

`func (o *ClusterAuthorizationRequest) GetAccountUsername() string`

GetAccountUsername returns the AccountUsername field if non-nil, zero value otherwise.

### GetAccountUsernameOk

`func (o *ClusterAuthorizationRequest) GetAccountUsernameOk() (*string, bool)`

GetAccountUsernameOk returns a tuple with the AccountUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUsername

`func (o *ClusterAuthorizationRequest) SetAccountUsername(v string)`

SetAccountUsername sets AccountUsername field to given value.


### GetAvailabilityZone

`func (o *ClusterAuthorizationRequest) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *ClusterAuthorizationRequest) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *ClusterAuthorizationRequest) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *ClusterAuthorizationRequest) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetByoc

`func (o *ClusterAuthorizationRequest) GetByoc() bool`

GetByoc returns the Byoc field if non-nil, zero value otherwise.

### GetByocOk

`func (o *ClusterAuthorizationRequest) GetByocOk() (*bool, bool)`

GetByocOk returns a tuple with the Byoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByoc

`func (o *ClusterAuthorizationRequest) SetByoc(v bool)`

SetByoc sets Byoc field to given value.

### HasByoc

`func (o *ClusterAuthorizationRequest) HasByoc() bool`

HasByoc returns a boolean if a field has been set.

### GetCloudAccountId

`func (o *ClusterAuthorizationRequest) GetCloudAccountId() string`

GetCloudAccountId returns the CloudAccountId field if non-nil, zero value otherwise.

### GetCloudAccountIdOk

`func (o *ClusterAuthorizationRequest) GetCloudAccountIdOk() (*string, bool)`

GetCloudAccountIdOk returns a tuple with the CloudAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAccountId

`func (o *ClusterAuthorizationRequest) SetCloudAccountId(v string)`

SetCloudAccountId sets CloudAccountId field to given value.

### HasCloudAccountId

`func (o *ClusterAuthorizationRequest) HasCloudAccountId() bool`

HasCloudAccountId returns a boolean if a field has been set.

### GetCloudProviderId

`func (o *ClusterAuthorizationRequest) GetCloudProviderId() string`

GetCloudProviderId returns the CloudProviderId field if non-nil, zero value otherwise.

### GetCloudProviderIdOk

`func (o *ClusterAuthorizationRequest) GetCloudProviderIdOk() (*string, bool)`

GetCloudProviderIdOk returns a tuple with the CloudProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderId

`func (o *ClusterAuthorizationRequest) SetCloudProviderId(v string)`

SetCloudProviderId sets CloudProviderId field to given value.

### HasCloudProviderId

`func (o *ClusterAuthorizationRequest) HasCloudProviderId() bool`

HasCloudProviderId returns a boolean if a field has been set.

### GetClusterId

`func (o *ClusterAuthorizationRequest) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ClusterAuthorizationRequest) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ClusterAuthorizationRequest) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetDisconnected

`func (o *ClusterAuthorizationRequest) GetDisconnected() bool`

GetDisconnected returns the Disconnected field if non-nil, zero value otherwise.

### GetDisconnectedOk

`func (o *ClusterAuthorizationRequest) GetDisconnectedOk() (*bool, bool)`

GetDisconnectedOk returns a tuple with the Disconnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnected

`func (o *ClusterAuthorizationRequest) SetDisconnected(v bool)`

SetDisconnected sets Disconnected field to given value.

### HasDisconnected

`func (o *ClusterAuthorizationRequest) HasDisconnected() bool`

HasDisconnected returns a boolean if a field has been set.

### GetDisplayName

`func (o *ClusterAuthorizationRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ClusterAuthorizationRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ClusterAuthorizationRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ClusterAuthorizationRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExternalClusterId

`func (o *ClusterAuthorizationRequest) GetExternalClusterId() string`

GetExternalClusterId returns the ExternalClusterId field if non-nil, zero value otherwise.

### GetExternalClusterIdOk

`func (o *ClusterAuthorizationRequest) GetExternalClusterIdOk() (*string, bool)`

GetExternalClusterIdOk returns a tuple with the ExternalClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalClusterId

`func (o *ClusterAuthorizationRequest) SetExternalClusterId(v string)`

SetExternalClusterId sets ExternalClusterId field to given value.

### HasExternalClusterId

`func (o *ClusterAuthorizationRequest) HasExternalClusterId() bool`

HasExternalClusterId returns a boolean if a field has been set.

### GetManaged

`func (o *ClusterAuthorizationRequest) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *ClusterAuthorizationRequest) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *ClusterAuthorizationRequest) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *ClusterAuthorizationRequest) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetProductCategory

`func (o *ClusterAuthorizationRequest) GetProductCategory() string`

GetProductCategory returns the ProductCategory field if non-nil, zero value otherwise.

### GetProductCategoryOk

`func (o *ClusterAuthorizationRequest) GetProductCategoryOk() (*string, bool)`

GetProductCategoryOk returns a tuple with the ProductCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCategory

`func (o *ClusterAuthorizationRequest) SetProductCategory(v string)`

SetProductCategory sets ProductCategory field to given value.

### HasProductCategory

`func (o *ClusterAuthorizationRequest) HasProductCategory() bool`

HasProductCategory returns a boolean if a field has been set.

### GetProductId

`func (o *ClusterAuthorizationRequest) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ClusterAuthorizationRequest) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ClusterAuthorizationRequest) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *ClusterAuthorizationRequest) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetReserve

`func (o *ClusterAuthorizationRequest) GetReserve() bool`

GetReserve returns the Reserve field if non-nil, zero value otherwise.

### GetReserveOk

`func (o *ClusterAuthorizationRequest) GetReserveOk() (*bool, bool)`

GetReserveOk returns a tuple with the Reserve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserve

`func (o *ClusterAuthorizationRequest) SetReserve(v bool)`

SetReserve sets Reserve field to given value.

### HasReserve

`func (o *ClusterAuthorizationRequest) HasReserve() bool`

HasReserve returns a boolean if a field has been set.

### GetResources

`func (o *ClusterAuthorizationRequest) GetResources() []ReservedResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ClusterAuthorizationRequest) GetResourcesOk() (*[]ReservedResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ClusterAuthorizationRequest) SetResources(v []ReservedResource)`

SetResources sets Resources field to given value.

### HasResources

`func (o *ClusterAuthorizationRequest) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


