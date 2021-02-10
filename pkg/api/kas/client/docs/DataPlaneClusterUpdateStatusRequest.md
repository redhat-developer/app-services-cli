# DataPlaneClusterUpdateStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]DataPlaneClusterUpdateStatusRequestConditions**](DataPlaneClusterUpdateStatusRequestConditions.md) | The cluster data plane conditions | [optional] 
**Total** | Pointer to [**DataPlaneClusterUpdateStatusRequestTotal**](DataPlaneClusterUpdateStatusRequest_total.md) |  | [optional] 
**Remaining** | Pointer to [**DataPlaneClusterUpdateStatusRequestTotal**](DataPlaneClusterUpdateStatusRequest_total.md) |  | [optional] 
**NodeInfo** | Pointer to [**DataPlaneClusterUpdateStatusRequestNodeInfo**](DataPlaneClusterUpdateStatusRequest_nodeInfo.md) |  | [optional] 
**ResizeInfo** | Pointer to [**DataPlaneClusterUpdateStatusRequestResizeInfo**](DataPlaneClusterUpdateStatusRequest_resizeInfo.md) |  | [optional] 

## Methods

### NewDataPlaneClusterUpdateStatusRequest

`func NewDataPlaneClusterUpdateStatusRequest() *DataPlaneClusterUpdateStatusRequest`

NewDataPlaneClusterUpdateStatusRequest instantiates a new DataPlaneClusterUpdateStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataPlaneClusterUpdateStatusRequestWithDefaults

`func NewDataPlaneClusterUpdateStatusRequestWithDefaults() *DataPlaneClusterUpdateStatusRequest`

NewDataPlaneClusterUpdateStatusRequestWithDefaults instantiates a new DataPlaneClusterUpdateStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *DataPlaneClusterUpdateStatusRequest) GetConditions() []DataPlaneClusterUpdateStatusRequestConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *DataPlaneClusterUpdateStatusRequest) GetConditionsOk() (*[]DataPlaneClusterUpdateStatusRequestConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *DataPlaneClusterUpdateStatusRequest) SetConditions(v []DataPlaneClusterUpdateStatusRequestConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *DataPlaneClusterUpdateStatusRequest) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetTotal

`func (o *DataPlaneClusterUpdateStatusRequest) GetTotal() DataPlaneClusterUpdateStatusRequestTotal`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *DataPlaneClusterUpdateStatusRequest) GetTotalOk() (*DataPlaneClusterUpdateStatusRequestTotal, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *DataPlaneClusterUpdateStatusRequest) SetTotal(v DataPlaneClusterUpdateStatusRequestTotal)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *DataPlaneClusterUpdateStatusRequest) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetRemaining

`func (o *DataPlaneClusterUpdateStatusRequest) GetRemaining() DataPlaneClusterUpdateStatusRequestTotal`

GetRemaining returns the Remaining field if non-nil, zero value otherwise.

### GetRemainingOk

`func (o *DataPlaneClusterUpdateStatusRequest) GetRemainingOk() (*DataPlaneClusterUpdateStatusRequestTotal, bool)`

GetRemainingOk returns a tuple with the Remaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemaining

`func (o *DataPlaneClusterUpdateStatusRequest) SetRemaining(v DataPlaneClusterUpdateStatusRequestTotal)`

SetRemaining sets Remaining field to given value.

### HasRemaining

`func (o *DataPlaneClusterUpdateStatusRequest) HasRemaining() bool`

HasRemaining returns a boolean if a field has been set.

### GetNodeInfo

`func (o *DataPlaneClusterUpdateStatusRequest) GetNodeInfo() DataPlaneClusterUpdateStatusRequestNodeInfo`

GetNodeInfo returns the NodeInfo field if non-nil, zero value otherwise.

### GetNodeInfoOk

`func (o *DataPlaneClusterUpdateStatusRequest) GetNodeInfoOk() (*DataPlaneClusterUpdateStatusRequestNodeInfo, bool)`

GetNodeInfoOk returns a tuple with the NodeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeInfo

`func (o *DataPlaneClusterUpdateStatusRequest) SetNodeInfo(v DataPlaneClusterUpdateStatusRequestNodeInfo)`

SetNodeInfo sets NodeInfo field to given value.

### HasNodeInfo

`func (o *DataPlaneClusterUpdateStatusRequest) HasNodeInfo() bool`

HasNodeInfo returns a boolean if a field has been set.

### GetResizeInfo

`func (o *DataPlaneClusterUpdateStatusRequest) GetResizeInfo() DataPlaneClusterUpdateStatusRequestResizeInfo`

GetResizeInfo returns the ResizeInfo field if non-nil, zero value otherwise.

### GetResizeInfoOk

`func (o *DataPlaneClusterUpdateStatusRequest) GetResizeInfoOk() (*DataPlaneClusterUpdateStatusRequestResizeInfo, bool)`

GetResizeInfoOk returns a tuple with the ResizeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizeInfo

`func (o *DataPlaneClusterUpdateStatusRequest) SetResizeInfo(v DataPlaneClusterUpdateStatusRequestResizeInfo)`

SetResizeInfo sets ResizeInfo field to given value.

### HasResizeInfo

`func (o *DataPlaneClusterUpdateStatusRequest) HasResizeInfo() bool`

HasResizeInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


