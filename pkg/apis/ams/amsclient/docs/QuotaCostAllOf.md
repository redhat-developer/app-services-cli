# QuotaCostAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allowed** | **int32** |  | 
**Consumed** | **int32** |  | 
**OrganizationId** | Pointer to **string** |  | [optional] 
**QuotaId** | **string** |  | 
**RelatedResources** | Pointer to [**[]RelatedResource**](RelatedResource.md) |  | [optional] 

## Methods

### NewQuotaCostAllOf

`func NewQuotaCostAllOf(allowed int32, consumed int32, quotaId string, ) *QuotaCostAllOf`

NewQuotaCostAllOf instantiates a new QuotaCostAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaCostAllOfWithDefaults

`func NewQuotaCostAllOfWithDefaults() *QuotaCostAllOf`

NewQuotaCostAllOfWithDefaults instantiates a new QuotaCostAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowed

`func (o *QuotaCostAllOf) GetAllowed() int32`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *QuotaCostAllOf) GetAllowedOk() (*int32, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *QuotaCostAllOf) SetAllowed(v int32)`

SetAllowed sets Allowed field to given value.


### GetConsumed

`func (o *QuotaCostAllOf) GetConsumed() int32`

GetConsumed returns the Consumed field if non-nil, zero value otherwise.

### GetConsumedOk

`func (o *QuotaCostAllOf) GetConsumedOk() (*int32, bool)`

GetConsumedOk returns a tuple with the Consumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumed

`func (o *QuotaCostAllOf) SetConsumed(v int32)`

SetConsumed sets Consumed field to given value.


### GetOrganizationId

`func (o *QuotaCostAllOf) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *QuotaCostAllOf) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *QuotaCostAllOf) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *QuotaCostAllOf) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetQuotaId

`func (o *QuotaCostAllOf) GetQuotaId() string`

GetQuotaId returns the QuotaId field if non-nil, zero value otherwise.

### GetQuotaIdOk

`func (o *QuotaCostAllOf) GetQuotaIdOk() (*string, bool)`

GetQuotaIdOk returns a tuple with the QuotaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaId

`func (o *QuotaCostAllOf) SetQuotaId(v string)`

SetQuotaId sets QuotaId field to given value.


### GetRelatedResources

`func (o *QuotaCostAllOf) GetRelatedResources() []RelatedResource`

GetRelatedResources returns the RelatedResources field if non-nil, zero value otherwise.

### GetRelatedResourcesOk

`func (o *QuotaCostAllOf) GetRelatedResourcesOk() (*[]RelatedResource, bool)`

GetRelatedResourcesOk returns a tuple with the RelatedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResources

`func (o *QuotaCostAllOf) SetRelatedResources(v []RelatedResource)`

SetRelatedResources sets RelatedResources field to given value.

### HasRelatedResources

`func (o *QuotaCostAllOf) HasRelatedResources() bool`

HasRelatedResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


