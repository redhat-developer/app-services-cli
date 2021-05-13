# RegistryDeploymentCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegistryDeploymentUrl** | **string** |  | 
**TenantManagerUrl** | **string** |  | 
**Name** | Pointer to **string** | User-defined Registry Deployment name. Does not have to be unique. | [optional] 

## Methods

### NewRegistryDeploymentCreate

`func NewRegistryDeploymentCreate(registryDeploymentUrl string, tenantManagerUrl string, ) *RegistryDeploymentCreate`

NewRegistryDeploymentCreate instantiates a new RegistryDeploymentCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryDeploymentCreateWithDefaults

`func NewRegistryDeploymentCreateWithDefaults() *RegistryDeploymentCreate`

NewRegistryDeploymentCreateWithDefaults instantiates a new RegistryDeploymentCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistryDeploymentUrl

`func (o *RegistryDeploymentCreate) GetRegistryDeploymentUrl() string`

GetRegistryDeploymentUrl returns the RegistryDeploymentUrl field if non-nil, zero value otherwise.

### GetRegistryDeploymentUrlOk

`func (o *RegistryDeploymentCreate) GetRegistryDeploymentUrlOk() (*string, bool)`

GetRegistryDeploymentUrlOk returns a tuple with the RegistryDeploymentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryDeploymentUrl

`func (o *RegistryDeploymentCreate) SetRegistryDeploymentUrl(v string)`

SetRegistryDeploymentUrl sets RegistryDeploymentUrl field to given value.


### GetTenantManagerUrl

`func (o *RegistryDeploymentCreate) GetTenantManagerUrl() string`

GetTenantManagerUrl returns the TenantManagerUrl field if non-nil, zero value otherwise.

### GetTenantManagerUrlOk

`func (o *RegistryDeploymentCreate) GetTenantManagerUrlOk() (*string, bool)`

GetTenantManagerUrlOk returns a tuple with the TenantManagerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantManagerUrl

`func (o *RegistryDeploymentCreate) SetTenantManagerUrl(v string)`

SetTenantManagerUrl sets TenantManagerUrl field to given value.


### GetName

`func (o *RegistryDeploymentCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegistryDeploymentCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegistryDeploymentCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RegistryDeploymentCreate) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


