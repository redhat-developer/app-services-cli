# Capability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Inherited** | **bool** |  | 
**Name** | **string** |  | 
**Value** | **string** |  | 

## Methods

### NewCapability

`func NewCapability(inherited bool, name string, value string, ) *Capability`

NewCapability instantiates a new Capability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityWithDefaults

`func NewCapabilityWithDefaults() *Capability`

NewCapabilityWithDefaults instantiates a new Capability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *Capability) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Capability) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Capability) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Capability) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *Capability) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Capability) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Capability) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Capability) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *Capability) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Capability) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Capability) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Capability) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetInherited

`func (o *Capability) GetInherited() bool`

GetInherited returns the Inherited field if non-nil, zero value otherwise.

### GetInheritedOk

`func (o *Capability) GetInheritedOk() (*bool, bool)`

GetInheritedOk returns a tuple with the Inherited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInherited

`func (o *Capability) SetInherited(v bool)`

SetInherited sets Inherited field to given value.


### GetName

`func (o *Capability) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Capability) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Capability) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *Capability) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Capability) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Capability) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


