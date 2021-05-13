# Registry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Status** | [**RegistryStatus**](RegistryStatus.md) |  | 
**RegistryUrl** | **string** |  | 
**Name** | Pointer to **string** | User-defined Registry name. Does not have to be unique. | [optional] 
**RegistryDeploymentId** | Pointer to **int32** | Identifier of a multi-tenant deployment, where this Service Registry instance resides. | [optional] 

## Methods

### NewRegistry

`func NewRegistry(id int32, status RegistryStatus, registryUrl string, ) *Registry`

NewRegistry instantiates a new Registry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryWithDefaults

`func NewRegistryWithDefaults() *Registry`

NewRegistryWithDefaults instantiates a new Registry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Registry) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Registry) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Registry) SetId(v int32)`

SetId sets Id field to given value.


### GetStatus

`func (o *Registry) GetStatus() RegistryStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Registry) GetStatusOk() (*RegistryStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Registry) SetStatus(v RegistryStatus)`

SetStatus sets Status field to given value.


### GetRegistryUrl

`func (o *Registry) GetRegistryUrl() string`

GetRegistryUrl returns the RegistryUrl field if non-nil, zero value otherwise.

### GetRegistryUrlOk

`func (o *Registry) GetRegistryUrlOk() (*string, bool)`

GetRegistryUrlOk returns a tuple with the RegistryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryUrl

`func (o *Registry) SetRegistryUrl(v string)`

SetRegistryUrl sets RegistryUrl field to given value.


### GetName

`func (o *Registry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Registry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Registry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Registry) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegistryDeploymentId

`func (o *Registry) GetRegistryDeploymentId() int32`

GetRegistryDeploymentId returns the RegistryDeploymentId field if non-nil, zero value otherwise.

### GetRegistryDeploymentIdOk

`func (o *Registry) GetRegistryDeploymentIdOk() (*int32, bool)`

GetRegistryDeploymentIdOk returns a tuple with the RegistryDeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryDeploymentId

`func (o *Registry) SetRegistryDeploymentId(v int32)`

SetRegistryDeploymentId sets RegistryDeploymentId field to given value.

### HasRegistryDeploymentId

`func (o *Registry) HasRegistryDeploymentId() bool`

HasRegistryDeploymentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


