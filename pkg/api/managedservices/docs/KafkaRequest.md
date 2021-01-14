# KafkaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
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

## Methods

### NewKafkaRequest

`func NewKafkaRequest() *KafkaRequest`

NewKafkaRequest instantiates a new KafkaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaRequestWithDefaults

`func NewKafkaRequestWithDefaults() *KafkaRequest`

NewKafkaRequestWithDefaults instantiates a new KafkaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KafkaRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KafkaRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KafkaRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KafkaRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *KafkaRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *KafkaRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *KafkaRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *KafkaRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetHref

`func (o *KafkaRequest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *KafkaRequest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *KafkaRequest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *KafkaRequest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetStatus

`func (o *KafkaRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KafkaRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KafkaRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KafkaRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCloudProvider

`func (o *KafkaRequest) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *KafkaRequest) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *KafkaRequest) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *KafkaRequest) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetMultiAz

`func (o *KafkaRequest) GetMultiAz() bool`

GetMultiAz returns the MultiAz field if non-nil, zero value otherwise.

### GetMultiAzOk

`func (o *KafkaRequest) GetMultiAzOk() (*bool, bool)`

GetMultiAzOk returns a tuple with the MultiAz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiAz

`func (o *KafkaRequest) SetMultiAz(v bool)`

SetMultiAz sets MultiAz field to given value.

### HasMultiAz

`func (o *KafkaRequest) HasMultiAz() bool`

HasMultiAz returns a boolean if a field has been set.

### GetRegion

`func (o *KafkaRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *KafkaRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *KafkaRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *KafkaRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetOwner

`func (o *KafkaRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *KafkaRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *KafkaRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *KafkaRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetName

`func (o *KafkaRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KafkaRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KafkaRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KafkaRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBootstrapServerHost

`func (o *KafkaRequest) GetBootstrapServerHost() string`

GetBootstrapServerHost returns the BootstrapServerHost field if non-nil, zero value otherwise.

### GetBootstrapServerHostOk

`func (o *KafkaRequest) GetBootstrapServerHostOk() (*string, bool)`

GetBootstrapServerHostOk returns a tuple with the BootstrapServerHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapServerHost

`func (o *KafkaRequest) SetBootstrapServerHost(v string)`

SetBootstrapServerHost sets BootstrapServerHost field to given value.

### HasBootstrapServerHost

`func (o *KafkaRequest) HasBootstrapServerHost() bool`

HasBootstrapServerHost returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KafkaRequest) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KafkaRequest) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KafkaRequest) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KafkaRequest) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KafkaRequest) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KafkaRequest) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KafkaRequest) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KafkaRequest) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetFailedReason

`func (o *KafkaRequest) GetFailedReason() string`

GetFailedReason returns the FailedReason field if non-nil, zero value otherwise.

### GetFailedReasonOk

`func (o *KafkaRequest) GetFailedReasonOk() (*string, bool)`

GetFailedReasonOk returns a tuple with the FailedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedReason

`func (o *KafkaRequest) SetFailedReason(v string)`

SetFailedReason sets FailedReason field to given value.

### HasFailedReason

`func (o *KafkaRequest) HasFailedReason() bool`

HasFailedReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


