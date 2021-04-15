# ServiceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kafkas** | Pointer to [**ServiceStatusKafkas**](ServiceStatusKafkas.md) |  | [optional] 

## Methods

### NewServiceStatus

`func NewServiceStatus() *ServiceStatus`

NewServiceStatus instantiates a new ServiceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceStatusWithDefaults

`func NewServiceStatusWithDefaults() *ServiceStatus`

NewServiceStatusWithDefaults instantiates a new ServiceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKafkas

`func (o *ServiceStatus) GetKafkas() ServiceStatusKafkas`

GetKafkas returns the Kafkas field if non-nil, zero value otherwise.

### GetKafkasOk

`func (o *ServiceStatus) GetKafkasOk() (*ServiceStatusKafkas, bool)`

GetKafkasOk returns a tuple with the Kafkas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkas

`func (o *ServiceStatus) SetKafkas(v ServiceStatusKafkas)`

SetKafkas sets Kafkas field to given value.

### HasKafkas

`func (o *ServiceStatus) HasKafkas() bool`

HasKafkas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


