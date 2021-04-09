# QuotaCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Allowed** | **int32** |  | 
**Consumed** | **int32** |  | 
**OrganizationId** | Pointer to **string** |  | [optional] 
**QuotaId** | **string** |  | 
**RelatedResources** | Pointer to [**[]RelatedResource**](RelatedResource.md) |  | [optional] 

## Methods

### NewQuotaCost

`func NewQuotaCost(allowed int32, consumed int32, quotaId string, ) *QuotaCost`

NewQuotaCost instantiates a new QuotaCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaCostWithDefaults

`func NewQuotaCostWithDefaults() *QuotaCost`

NewQuotaCostWithDefaults instantiates a new QuotaCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *QuotaCost) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *QuotaCost) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *QuotaCost) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *QuotaCost) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *QuotaCost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuotaCost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuotaCost) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *QuotaCost) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *QuotaCost) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *QuotaCost) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *QuotaCost) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *QuotaCost) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetAllowed

`func (o *QuotaCost) GetAllowed() int32`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *QuotaCost) GetAllowedOk() (*int32, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *QuotaCost) SetAllowed(v int32)`

SetAllowed sets Allowed field to given value.


### GetConsumed

`func (o *QuotaCost) GetConsumed() int32`

GetConsumed returns the Consumed field if non-nil, zero value otherwise.

### GetConsumedOk

`func (o *QuotaCost) GetConsumedOk() (*int32, bool)`

GetConsumedOk returns a tuple with the Consumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumed

`func (o *QuotaCost) SetConsumed(v int32)`

SetConsumed sets Consumed field to given value.


### GetOrganizationId

`func (o *QuotaCost) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *QuotaCost) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *QuotaCost) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *QuotaCost) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetQuotaId

`func (o *QuotaCost) GetQuotaId() string`

GetQuotaId returns the QuotaId field if non-nil, zero value otherwise.

### GetQuotaIdOk

`func (o *QuotaCost) GetQuotaIdOk() (*string, bool)`

GetQuotaIdOk returns a tuple with the QuotaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaId

`func (o *QuotaCost) SetQuotaId(v string)`

SetQuotaId sets QuotaId field to given value.


### GetRelatedResources

`func (o *QuotaCost) GetRelatedResources() []RelatedResource`

GetRelatedResources returns the RelatedResources field if non-nil, zero value otherwise.

### GetRelatedResourcesOk

`func (o *QuotaCost) GetRelatedResourcesOk() (*[]RelatedResource, bool)`

GetRelatedResourcesOk returns a tuple with the RelatedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResources

`func (o *QuotaCost) SetRelatedResources(v []RelatedResource)`

SetRelatedResources sets RelatedResources field to given value.

### HasRelatedResources

`func (o *QuotaCost) HasRelatedResources() bool`

HasRelatedResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


