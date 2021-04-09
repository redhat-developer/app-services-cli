# ExportControlReviewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountUsername** | **string** |  | 
**IgnoreCache** | Pointer to **bool** |  | [optional] 

## Methods

### NewExportControlReviewRequest

`func NewExportControlReviewRequest(accountUsername string, ) *ExportControlReviewRequest`

NewExportControlReviewRequest instantiates a new ExportControlReviewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportControlReviewRequestWithDefaults

`func NewExportControlReviewRequestWithDefaults() *ExportControlReviewRequest`

NewExportControlReviewRequestWithDefaults instantiates a new ExportControlReviewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountUsername

`func (o *ExportControlReviewRequest) GetAccountUsername() string`

GetAccountUsername returns the AccountUsername field if non-nil, zero value otherwise.

### GetAccountUsernameOk

`func (o *ExportControlReviewRequest) GetAccountUsernameOk() (*string, bool)`

GetAccountUsernameOk returns a tuple with the AccountUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUsername

`func (o *ExportControlReviewRequest) SetAccountUsername(v string)`

SetAccountUsername sets AccountUsername field to given value.


### GetIgnoreCache

`func (o *ExportControlReviewRequest) GetIgnoreCache() bool`

GetIgnoreCache returns the IgnoreCache field if non-nil, zero value otherwise.

### GetIgnoreCacheOk

`func (o *ExportControlReviewRequest) GetIgnoreCacheOk() (*bool, bool)`

GetIgnoreCacheOk returns a tuple with the IgnoreCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreCache

`func (o *ExportControlReviewRequest) SetIgnoreCache(v bool)`

SetIgnoreCache sets IgnoreCache field to given value.

### HasIgnoreCache

`func (o *ExportControlReviewRequest) HasIgnoreCache() bool`

HasIgnoreCache returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


