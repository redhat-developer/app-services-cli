# QuotaSummaryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allowed** | **int32** |  | 
**AvailabilityZoneType** | **string** |  | 
**Byoc** | **bool** |  | 
**OrganizationId** | Pointer to **string** |  | [optional] 
**Reserved** | **int32** |  | 
**ResourceName** | **string** |  | 
**ResourceType** | **string** |  | 

## Methods

### NewQuotaSummaryAllOf

`func NewQuotaSummaryAllOf(allowed int32, availabilityZoneType string, byoc bool, reserved int32, resourceName string, resourceType string, ) *QuotaSummaryAllOf`

NewQuotaSummaryAllOf instantiates a new QuotaSummaryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaSummaryAllOfWithDefaults

`func NewQuotaSummaryAllOfWithDefaults() *QuotaSummaryAllOf`

NewQuotaSummaryAllOfWithDefaults instantiates a new QuotaSummaryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowed

`func (o *QuotaSummaryAllOf) GetAllowed() int32`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *QuotaSummaryAllOf) GetAllowedOk() (*int32, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *QuotaSummaryAllOf) SetAllowed(v int32)`

SetAllowed sets Allowed field to given value.


### GetAvailabilityZoneType

`func (o *QuotaSummaryAllOf) GetAvailabilityZoneType() string`

GetAvailabilityZoneType returns the AvailabilityZoneType field if non-nil, zero value otherwise.

### GetAvailabilityZoneTypeOk

`func (o *QuotaSummaryAllOf) GetAvailabilityZoneTypeOk() (*string, bool)`

GetAvailabilityZoneTypeOk returns a tuple with the AvailabilityZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneType

`func (o *QuotaSummaryAllOf) SetAvailabilityZoneType(v string)`

SetAvailabilityZoneType sets AvailabilityZoneType field to given value.


### GetByoc

`func (o *QuotaSummaryAllOf) GetByoc() bool`

GetByoc returns the Byoc field if non-nil, zero value otherwise.

### GetByocOk

`func (o *QuotaSummaryAllOf) GetByocOk() (*bool, bool)`

GetByocOk returns a tuple with the Byoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByoc

`func (o *QuotaSummaryAllOf) SetByoc(v bool)`

SetByoc sets Byoc field to given value.


### GetOrganizationId

`func (o *QuotaSummaryAllOf) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *QuotaSummaryAllOf) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *QuotaSummaryAllOf) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *QuotaSummaryAllOf) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetReserved

`func (o *QuotaSummaryAllOf) GetReserved() int32`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *QuotaSummaryAllOf) GetReservedOk() (*int32, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *QuotaSummaryAllOf) SetReserved(v int32)`

SetReserved sets Reserved field to given value.


### GetResourceName

`func (o *QuotaSummaryAllOf) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *QuotaSummaryAllOf) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *QuotaSummaryAllOf) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetResourceType

`func (o *QuotaSummaryAllOf) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *QuotaSummaryAllOf) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *QuotaSummaryAllOf) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


