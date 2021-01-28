# CloudRegion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | Indicates the type of this object. Will be &#39;CloudRegion&#39;. | [optional] 
**Id** | Pointer to **string** | Unique identifier of the object. | [optional] 
**DisplayName** | Pointer to **string** | Name of the region for display purposes, for example &#x60;N. Virginia&#x60;. | [optional] 
**Enabled** | **bool** | Whether the region is enabled for deploying an OSD cluster. | [default to false]

## Methods

### NewCloudRegion

`func NewCloudRegion(enabled bool, ) *CloudRegion`

NewCloudRegion instantiates a new CloudRegion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRegionWithDefaults

`func NewCloudRegionWithDefaults() *CloudRegion`

NewCloudRegionWithDefaults instantiates a new CloudRegion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CloudRegion) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudRegion) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudRegion) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CloudRegion) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *CloudRegion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudRegion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudRegion) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudRegion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *CloudRegion) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CloudRegion) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CloudRegion) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CloudRegion) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnabled

`func (o *CloudRegion) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CloudRegion) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CloudRegion) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


