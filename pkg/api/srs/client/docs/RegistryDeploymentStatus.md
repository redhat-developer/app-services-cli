# RegistryDeploymentStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastUpdated** | **time.Time** | ISO 8601 UTC timestamp. | 
**Value** | [**RegistryDeploymentStatusValue**](RegistryDeploymentStatusValue.md) |  | 

## Methods

### NewRegistryDeploymentStatus

`func NewRegistryDeploymentStatus(lastUpdated time.Time, value RegistryDeploymentStatusValue, ) *RegistryDeploymentStatus`

NewRegistryDeploymentStatus instantiates a new RegistryDeploymentStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryDeploymentStatusWithDefaults

`func NewRegistryDeploymentStatusWithDefaults() *RegistryDeploymentStatus`

NewRegistryDeploymentStatusWithDefaults instantiates a new RegistryDeploymentStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastUpdated

`func (o *RegistryDeploymentStatus) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *RegistryDeploymentStatus) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *RegistryDeploymentStatus) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetValue

`func (o *RegistryDeploymentStatus) GetValue() RegistryDeploymentStatusValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RegistryDeploymentStatus) GetValueOk() (*RegistryDeploymentStatusValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RegistryDeploymentStatus) SetValue(v RegistryDeploymentStatusValue)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


