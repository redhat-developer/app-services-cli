# ClusterResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | [**ClusterResourceTotal**](ClusterResourceTotal.md) |  | 
**UpdatedTimestamp** | **time.Time** |  | 
**Used** | [**ClusterResourceTotal**](ClusterResourceTotal.md) |  | 

## Methods

### NewClusterResource

`func NewClusterResource(total ClusterResourceTotal, updatedTimestamp time.Time, used ClusterResourceTotal, ) *ClusterResource`

NewClusterResource instantiates a new ClusterResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterResourceWithDefaults

`func NewClusterResourceWithDefaults() *ClusterResource`

NewClusterResourceWithDefaults instantiates a new ClusterResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *ClusterResource) GetTotal() ClusterResourceTotal`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ClusterResource) GetTotalOk() (*ClusterResourceTotal, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ClusterResource) SetTotal(v ClusterResourceTotal)`

SetTotal sets Total field to given value.


### GetUpdatedTimestamp

`func (o *ClusterResource) GetUpdatedTimestamp() time.Time`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *ClusterResource) GetUpdatedTimestampOk() (*time.Time, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *ClusterResource) SetUpdatedTimestamp(v time.Time)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetUsed

`func (o *ClusterResource) GetUsed() ClusterResourceTotal`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *ClusterResource) GetUsedOk() (*ClusterResourceTotal, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *ClusterResource) SetUsed(v ClusterResourceTotal)`

SetUsed sets Used field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


