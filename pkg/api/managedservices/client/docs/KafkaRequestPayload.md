# KafkaRequestPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to **string** | The cloud provider where the Kafka cluster will be created in | [optional] 
**MultiAz** | Pointer to **bool** | Set this to true to configure the Kafka cluster to be multiAZ | [optional] 
**Name** | **string** | The name of the Kafka cluster. It must consist of lower-case alphanumeric characters or &#39;-&#39;, start with an alphabetic character, and end with an alphanumeric character, and can not be longer than 32 characters. | 
**Region** | Pointer to **string** | The region where the Kafka cluster will be created in | [optional] 

## Methods

### NewKafkaRequestPayload

`func NewKafkaRequestPayload(name string, ) *KafkaRequestPayload`

NewKafkaRequestPayload instantiates a new KafkaRequestPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaRequestPayloadWithDefaults

`func NewKafkaRequestPayloadWithDefaults() *KafkaRequestPayload`

NewKafkaRequestPayloadWithDefaults instantiates a new KafkaRequestPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *KafkaRequestPayload) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *KafkaRequestPayload) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *KafkaRequestPayload) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *KafkaRequestPayload) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetMultiAz

`func (o *KafkaRequestPayload) GetMultiAz() bool`

GetMultiAz returns the MultiAz field if non-nil, zero value otherwise.

### GetMultiAzOk

`func (o *KafkaRequestPayload) GetMultiAzOk() (*bool, bool)`

GetMultiAzOk returns a tuple with the MultiAz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiAz

`func (o *KafkaRequestPayload) SetMultiAz(v bool)`

SetMultiAz sets MultiAz field to given value.

### HasMultiAz

`func (o *KafkaRequestPayload) HasMultiAz() bool`

HasMultiAz returns a boolean if a field has been set.

### GetName

`func (o *KafkaRequestPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KafkaRequestPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KafkaRequestPayload) SetName(v string)`

SetName sets Name field to given value.


### GetRegion

`func (o *KafkaRequestPayload) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *KafkaRequestPayload) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *KafkaRequestPayload) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *KafkaRequestPayload) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


