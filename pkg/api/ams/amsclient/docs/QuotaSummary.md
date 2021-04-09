# QuotaSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Allowed** | **int32** |  | 
**AvailabilityZoneType** | **string** |  | 
**Byoc** | **bool** |  | 
**OrganizationId** | Pointer to **string** |  | [optional] 
**Reserved** | **int32** |  | 
**ResourceName** | **string** |  | 
**ResourceType** | **string** |  | 

## Methods

### NewQuotaSummary

`func NewQuotaSummary(allowed int32, availabilityZoneType string, byoc bool, reserved int32, resourceName string, resourceType string, ) *QuotaSummary`

NewQuotaSummary instantiates a new QuotaSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaSummaryWithDefaults

`func NewQuotaSummaryWithDefaults() *QuotaSummary`

NewQuotaSummaryWithDefaults instantiates a new QuotaSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *QuotaSummary) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *QuotaSummary) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *QuotaSummary) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *QuotaSummary) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *QuotaSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuotaSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuotaSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *QuotaSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *QuotaSummary) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *QuotaSummary) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *QuotaSummary) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *QuotaSummary) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetAllowed

`func (o *QuotaSummary) GetAllowed() int32`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *QuotaSummary) GetAllowedOk() (*int32, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *QuotaSummary) SetAllowed(v int32)`

SetAllowed sets Allowed field to given value.


### GetAvailabilityZoneType

`func (o *QuotaSummary) GetAvailabilityZoneType() string`

GetAvailabilityZoneType returns the AvailabilityZoneType field if non-nil, zero value otherwise.

### GetAvailabilityZoneTypeOk

`func (o *QuotaSummary) GetAvailabilityZoneTypeOk() (*string, bool)`

GetAvailabilityZoneTypeOk returns a tuple with the AvailabilityZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneType

`func (o *QuotaSummary) SetAvailabilityZoneType(v string)`

SetAvailabilityZoneType sets AvailabilityZoneType field to given value.


### GetByoc

`func (o *QuotaSummary) GetByoc() bool`

GetByoc returns the Byoc field if non-nil, zero value otherwise.

### GetByocOk

`func (o *QuotaSummary) GetByocOk() (*bool, bool)`

GetByocOk returns a tuple with the Byoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByoc

`func (o *QuotaSummary) SetByoc(v bool)`

SetByoc sets Byoc field to given value.


### GetOrganizationId

`func (o *QuotaSummary) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *QuotaSummary) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *QuotaSummary) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *QuotaSummary) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetReserved

`func (o *QuotaSummary) GetReserved() int32`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *QuotaSummary) GetReservedOk() (*int32, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *QuotaSummary) SetReserved(v int32)`

SetReserved sets Reserved field to given value.


### GetResourceName

`func (o *QuotaSummary) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *QuotaSummary) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *QuotaSummary) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetResourceType

`func (o *QuotaSummary) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *QuotaSummary) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *QuotaSummary) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


