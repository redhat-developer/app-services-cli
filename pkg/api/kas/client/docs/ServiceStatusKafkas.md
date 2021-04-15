# ServiceStatusKafkas

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxCapacityReached** | **bool** | Indicates whether we have reached kafka maximum capacity | 

## Methods

### NewServiceStatusKafkas

`func NewServiceStatusKafkas(maxCapacityReached bool, ) *ServiceStatusKafkas`

NewServiceStatusKafkas instantiates a new ServiceStatusKafkas object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceStatusKafkasWithDefaults

`func NewServiceStatusKafkasWithDefaults() *ServiceStatusKafkas`

NewServiceStatusKafkasWithDefaults instantiates a new ServiceStatusKafkas object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxCapacityReached

`func (o *ServiceStatusKafkas) GetMaxCapacityReached() bool`

GetMaxCapacityReached returns the MaxCapacityReached field if non-nil, zero value otherwise.

### GetMaxCapacityReachedOk

`func (o *ServiceStatusKafkas) GetMaxCapacityReachedOk() (*bool, bool)`

GetMaxCapacityReachedOk returns a tuple with the MaxCapacityReached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCapacityReached

`func (o *ServiceStatusKafkas) SetMaxCapacityReached(v bool)`

SetMaxCapacityReached sets MaxCapacityReached field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


