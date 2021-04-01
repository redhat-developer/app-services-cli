# FeatureToggle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Enabled** | **bool** |  | [default to false]

## Methods

### NewFeatureToggle

`func NewFeatureToggle(enabled bool, ) *FeatureToggle`

NewFeatureToggle instantiates a new FeatureToggle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureToggleWithDefaults

`func NewFeatureToggleWithDefaults() *FeatureToggle`

NewFeatureToggleWithDefaults instantiates a new FeatureToggle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *FeatureToggle) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FeatureToggle) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FeatureToggle) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FeatureToggle) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *FeatureToggle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FeatureToggle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FeatureToggle) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FeatureToggle) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *FeatureToggle) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FeatureToggle) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FeatureToggle) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *FeatureToggle) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetEnabled

`func (o *FeatureToggle) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FeatureToggle) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FeatureToggle) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


