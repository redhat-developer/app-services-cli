# KafkaRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**CloudProvider** | Pointer to **string** |  | [optional] 
**MultiAz** | Pointer to **bool** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**BootstrapServerHost** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**FailedReason** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewKafkaRequestAllOf

`func NewKafkaRequestAllOf() *KafkaRequestAllOf`

NewKafkaRequestAllOf instantiates a new KafkaRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaRequestAllOfWithDefaults

`func NewKafkaRequestAllOfWithDefaults() *KafkaRequestAllOf`

NewKafkaRequestAllOfWithDefaults instantiates a new KafkaRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *KafkaRequestAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KafkaRequestAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KafkaRequestAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KafkaRequestAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCloudProvider

`func (o *KafkaRequestAllOf) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *KafkaRequestAllOf) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *KafkaRequestAllOf) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *KafkaRequestAllOf) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetMultiAz

`func (o *KafkaRequestAllOf) GetMultiAz() bool`

GetMultiAz returns the MultiAz field if non-nil, zero value otherwise.

### GetMultiAzOk

`func (o *KafkaRequestAllOf) GetMultiAzOk() (*bool, bool)`

GetMultiAzOk returns a tuple with the MultiAz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiAz

`func (o *KafkaRequestAllOf) SetMultiAz(v bool)`

SetMultiAz sets MultiAz field to given value.

### HasMultiAz

`func (o *KafkaRequestAllOf) HasMultiAz() bool`

HasMultiAz returns a boolean if a field has been set.

### GetRegion

`func (o *KafkaRequestAllOf) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *KafkaRequestAllOf) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *KafkaRequestAllOf) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *KafkaRequestAllOf) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetOwner

`func (o *KafkaRequestAllOf) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *KafkaRequestAllOf) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *KafkaRequestAllOf) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *KafkaRequestAllOf) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetName

`func (o *KafkaRequestAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KafkaRequestAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KafkaRequestAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KafkaRequestAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBootstrapServerHost

`func (o *KafkaRequestAllOf) GetBootstrapServerHost() string`

GetBootstrapServerHost returns the BootstrapServerHost field if non-nil, zero value otherwise.

### GetBootstrapServerHostOk

`func (o *KafkaRequestAllOf) GetBootstrapServerHostOk() (*string, bool)`

GetBootstrapServerHostOk returns a tuple with the BootstrapServerHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapServerHost

`func (o *KafkaRequestAllOf) SetBootstrapServerHost(v string)`

SetBootstrapServerHost sets BootstrapServerHost field to given value.

### HasBootstrapServerHost

`func (o *KafkaRequestAllOf) HasBootstrapServerHost() bool`

HasBootstrapServerHost returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KafkaRequestAllOf) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KafkaRequestAllOf) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KafkaRequestAllOf) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KafkaRequestAllOf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KafkaRequestAllOf) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KafkaRequestAllOf) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KafkaRequestAllOf) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KafkaRequestAllOf) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetFailedReason

`func (o *KafkaRequestAllOf) GetFailedReason() string`

GetFailedReason returns the FailedReason field if non-nil, zero value otherwise.

### GetFailedReasonOk

`func (o *KafkaRequestAllOf) GetFailedReasonOk() (*string, bool)`

GetFailedReasonOk returns a tuple with the FailedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedReason

`func (o *KafkaRequestAllOf) SetFailedReason(v string)`

SetFailedReason sets FailedReason field to given value.

### HasFailedReason

`func (o *KafkaRequestAllOf) HasFailedReason() bool`

HasFailedReason returns a boolean if a field has been set.

### GetVersion

`func (o *KafkaRequestAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KafkaRequestAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KafkaRequestAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *KafkaRequestAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


