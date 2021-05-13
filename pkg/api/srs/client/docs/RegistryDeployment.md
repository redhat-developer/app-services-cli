# RegistryDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**TenantManagerUrl** | **string** |  | 
**RegistryDeploymentUrl** | **string** |  | 
**Status** | [**RegistryDeploymentStatus**](RegistryDeploymentStatus.md) |  | 
**Name** | Pointer to **string** | User-defined Registry Deployment name. Does not have to be unique. | [optional] 

## Methods

### NewRegistryDeployment

`func NewRegistryDeployment(id int32, tenantManagerUrl string, registryDeploymentUrl string, status RegistryDeploymentStatus, ) *RegistryDeployment`

NewRegistryDeployment instantiates a new RegistryDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryDeploymentWithDefaults

`func NewRegistryDeploymentWithDefaults() *RegistryDeployment`

NewRegistryDeploymentWithDefaults instantiates a new RegistryDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RegistryDeployment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegistryDeployment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegistryDeployment) SetId(v int32)`

SetId sets Id field to given value.


### GetTenantManagerUrl

`func (o *RegistryDeployment) GetTenantManagerUrl() string`

GetTenantManagerUrl returns the TenantManagerUrl field if non-nil, zero value otherwise.

### GetTenantManagerUrlOk

`func (o *RegistryDeployment) GetTenantManagerUrlOk() (*string, bool)`

GetTenantManagerUrlOk returns a tuple with the TenantManagerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantManagerUrl

`func (o *RegistryDeployment) SetTenantManagerUrl(v string)`

SetTenantManagerUrl sets TenantManagerUrl field to given value.


### GetRegistryDeploymentUrl

`func (o *RegistryDeployment) GetRegistryDeploymentUrl() string`

GetRegistryDeploymentUrl returns the RegistryDeploymentUrl field if non-nil, zero value otherwise.

### GetRegistryDeploymentUrlOk

`func (o *RegistryDeployment) GetRegistryDeploymentUrlOk() (*string, bool)`

GetRegistryDeploymentUrlOk returns a tuple with the RegistryDeploymentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryDeploymentUrl

`func (o *RegistryDeployment) SetRegistryDeploymentUrl(v string)`

SetRegistryDeploymentUrl sets RegistryDeploymentUrl field to given value.


### GetStatus

`func (o *RegistryDeployment) GetStatus() RegistryDeploymentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RegistryDeployment) GetStatusOk() (*RegistryDeploymentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RegistryDeployment) SetStatus(v RegistryDeploymentStatus)`

SetStatus sets Status field to given value.


### GetName

`func (o *RegistryDeployment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegistryDeployment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegistryDeployment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RegistryDeployment) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


