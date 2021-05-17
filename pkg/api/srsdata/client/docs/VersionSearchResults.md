# VersionSearchResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | The total number of versions that matched the query (may be more than the number of versions returned in the result set). | 
**Versions** | [**[]SearchedVersion**](SearchedVersion.md) | The collection of artifact versions returned in the result set. | 

## Methods

### NewVersionSearchResults

`func NewVersionSearchResults(count int32, versions []SearchedVersion, ) *VersionSearchResults`

NewVersionSearchResults instantiates a new VersionSearchResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionSearchResultsWithDefaults

`func NewVersionSearchResultsWithDefaults() *VersionSearchResults`

NewVersionSearchResultsWithDefaults instantiates a new VersionSearchResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *VersionSearchResults) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *VersionSearchResults) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *VersionSearchResults) SetCount(v int32)`

SetCount sets Count field to given value.


### GetVersions

`func (o *VersionSearchResults) GetVersions() []SearchedVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *VersionSearchResults) GetVersionsOk() (*[]SearchedVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *VersionSearchResults) SetVersions(v []SearchedVersion)`

SetVersions sets Versions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


