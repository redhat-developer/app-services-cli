# DataPlaneClusterUpdateStatusRequestResizeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeDelta** | Pointer to **int32** |  | [optional] 
**Delta** | Pointer to [**DataPlaneClusterUpdateStatusRequestResizeInfoDelta**](DataPlaneClusterUpdateStatusRequest_resizeInfo_delta.md) |  | [optional] 

## Methods

### NewDataPlaneClusterUpdateStatusRequestResizeInfo

`func NewDataPlaneClusterUpdateStatusRequestResizeInfo() *DataPlaneClusterUpdateStatusRequestResizeInfo`

NewDataPlaneClusterUpdateStatusRequestResizeInfo instantiates a new DataPlaneClusterUpdateStatusRequestResizeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataPlaneClusterUpdateStatusRequestResizeInfoWithDefaults

`func NewDataPlaneClusterUpdateStatusRequestResizeInfoWithDefaults() *DataPlaneClusterUpdateStatusRequestResizeInfo`

NewDataPlaneClusterUpdateStatusRequestResizeInfoWithDefaults instantiates a new DataPlaneClusterUpdateStatusRequestResizeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeDelta

`func (o *DataPlaneClusterUpdateStatusRequestResizeInfo) GetNodeDelta() int32`

GetNodeDelta returns the NodeDelta field if non-nil, zero value otherwise.

### GetNodeDeltaOk

`func (o *DataPlaneClusterUpdateStatusRequestResizeInfo) GetNodeDeltaOk() (*int32, bool)`

GetNodeDeltaOk returns a tuple with the NodeDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeDelta

`func (o *DataPlaneClusterUpdateStatusRequestResizeInfo) SetNodeDelta(v int32)`

SetNodeDelta sets NodeDelta field to given value.

### HasNodeDelta

`func (o *DataPlaneClusterUpdateStatusRequestResizeInfo) HasNodeDelta() bool`

HasNodeDelta returns a boolean if a field has been set.

### GetDelta

`func (o *DataPlaneClusterUpdateStatusRequestResizeInfo) GetDelta() DataPlaneClusterUpdateStatusRequestResizeInfoDelta`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *DataPlaneClusterUpdateStatusRequestResizeInfo) GetDeltaOk() (*DataPlaneClusterUpdateStatusRequestResizeInfoDelta, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *DataPlaneClusterUpdateStatusRequestResizeInfo) SetDelta(v DataPlaneClusterUpdateStatusRequestResizeInfoDelta)`

SetDelta sets Delta field to given value.

### HasDelta

`func (o *DataPlaneClusterUpdateStatusRequestResizeInfo) HasDelta() bool`

HasDelta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


