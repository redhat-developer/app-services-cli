# CloudProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | Indicates the type of this object. Will be &#39;CloudProvider&#39; link. | [optional] 
**Id** | Pointer to **string** | Unique identifier of the object. | [optional] 
**DisplayName** | Pointer to **string** | Name of the cloud provider for display purposes. | [optional] 
**Name** | Pointer to **string** | Human friendly identifier of the cloud provider, for example &#x60;aws&#x60;. | [optional] 
**Enabled** | **bool** | Whether the cloud provider is enabled for deploying an OSD cluster. | 

## Methods

### NewCloudProvider

`func NewCloudProvider(enabled bool, ) *CloudProvider`

NewCloudProvider instantiates a new CloudProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProviderWithDefaults

`func NewCloudProviderWithDefaults() *CloudProvider`

NewCloudProviderWithDefaults instantiates a new CloudProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CloudProvider) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudProvider) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudProvider) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CloudProvider) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *CloudProvider) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudProvider) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudProvider) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudProvider) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *CloudProvider) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CloudProvider) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CloudProvider) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CloudProvider) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetName

`func (o *CloudProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudProvider) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudProvider) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *CloudProvider) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CloudProvider) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CloudProvider) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


