# ArtifactSearchResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artifacts** | [**[]SearchedArtifact**](SearchedArtifact.md) | The artifacts returned in the result set. | 
**Count** | **int32** | The total number of artifacts that matched the query that produced the result set (may be  more than the number of artifacts in the result set). | 

## Methods

### NewArtifactSearchResults

`func NewArtifactSearchResults(artifacts []SearchedArtifact, count int32, ) *ArtifactSearchResults`

NewArtifactSearchResults instantiates a new ArtifactSearchResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactSearchResultsWithDefaults

`func NewArtifactSearchResultsWithDefaults() *ArtifactSearchResults`

NewArtifactSearchResultsWithDefaults instantiates a new ArtifactSearchResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifacts

`func (o *ArtifactSearchResults) GetArtifacts() []SearchedArtifact`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *ArtifactSearchResults) GetArtifactsOk() (*[]SearchedArtifact, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *ArtifactSearchResults) SetArtifacts(v []SearchedArtifact)`

SetArtifacts sets Artifacts field to given value.


### GetCount

`func (o *ArtifactSearchResults) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ArtifactSearchResults) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ArtifactSearchResults) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


