# TermsReviewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**OrganizationId** | **string** |  | 
**RedirectUrl** | Pointer to **string** |  | [optional] 
**TermsAvailable** | **bool** |  | 
**TermsRequired** | **bool** |  | 

## Methods

### NewTermsReviewResponse

`func NewTermsReviewResponse(accountId string, organizationId string, termsAvailable bool, termsRequired bool, ) *TermsReviewResponse`

NewTermsReviewResponse instantiates a new TermsReviewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTermsReviewResponseWithDefaults

`func NewTermsReviewResponseWithDefaults() *TermsReviewResponse`

NewTermsReviewResponseWithDefaults instantiates a new TermsReviewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *TermsReviewResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TermsReviewResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TermsReviewResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetOrganizationId

`func (o *TermsReviewResponse) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *TermsReviewResponse) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *TermsReviewResponse) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetRedirectUrl

`func (o *TermsReviewResponse) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *TermsReviewResponse) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *TermsReviewResponse) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *TermsReviewResponse) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetTermsAvailable

`func (o *TermsReviewResponse) GetTermsAvailable() bool`

GetTermsAvailable returns the TermsAvailable field if non-nil, zero value otherwise.

### GetTermsAvailableOk

`func (o *TermsReviewResponse) GetTermsAvailableOk() (*bool, bool)`

GetTermsAvailableOk returns a tuple with the TermsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsAvailable

`func (o *TermsReviewResponse) SetTermsAvailable(v bool)`

SetTermsAvailable sets TermsAvailable field to given value.


### GetTermsRequired

`func (o *TermsReviewResponse) GetTermsRequired() bool`

GetTermsRequired returns the TermsRequired field if non-nil, zero value otherwise.

### GetTermsRequiredOk

`func (o *TermsReviewResponse) GetTermsRequiredOk() (*bool, bool)`

GetTermsRequiredOk returns a tuple with the TermsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsRequired

`func (o *TermsReviewResponse) SetTermsRequired(v bool)`

SetTermsRequired sets TermsRequired field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


