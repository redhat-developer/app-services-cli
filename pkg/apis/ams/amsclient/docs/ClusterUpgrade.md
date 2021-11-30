# ClusterUpgrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**UpdatedTimestamp** | Pointer to **time.Time** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewClusterUpgrade

`func NewClusterUpgrade() *ClusterUpgrade`

NewClusterUpgrade instantiates a new ClusterUpgrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterUpgradeWithDefaults

`func NewClusterUpgradeWithDefaults() *ClusterUpgrade`

NewClusterUpgradeWithDefaults instantiates a new ClusterUpgrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *ClusterUpgrade) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *ClusterUpgrade) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *ClusterUpgrade) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *ClusterUpgrade) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetState

`func (o *ClusterUpgrade) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ClusterUpgrade) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ClusterUpgrade) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ClusterUpgrade) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *ClusterUpgrade) GetUpdatedTimestamp() time.Time`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *ClusterUpgrade) GetUpdatedTimestampOk() (*time.Time, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *ClusterUpgrade) SetUpdatedTimestamp(v time.Time)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *ClusterUpgrade) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.

### GetVersion

`func (o *ClusterUpgrade) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ClusterUpgrade) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ClusterUpgrade) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ClusterUpgrade) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


