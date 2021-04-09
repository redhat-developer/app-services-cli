# OneMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | **string** |  | 
**ClusterType** | **string** |  | 
**ComputeNodesCpu** | [**ClusterResource**](ClusterResource.md) |  | 
**ComputeNodesMemory** | [**ClusterResource**](ClusterResource.md) |  | 
**ComputeNodesSockets** | [**ClusterResource**](ClusterResource.md) |  | 
**ConsoleUrl** | **string** |  | 
**Cpu** | [**ClusterResource**](ClusterResource.md) |  | 
**CriticalAlertsFiring** | **float64** |  | 
**HealthState** | Pointer to **string** |  | [optional] 
**Memory** | [**ClusterResource**](ClusterResource.md) |  | 
**Nodes** | [**ClusterMetricsNodes**](ClusterMetricsNodes.md) |  | 
**OpenshiftVersion** | **string** |  | 
**OperatingSystem** | **string** |  | 
**OperatorsConditionFailing** | **float64** |  | 
**Region** | **string** |  | 
**Sockets** | [**ClusterResource**](ClusterResource.md) |  | 
**State** | **string** |  | 
**StateDescription** | **string** |  | 
**Storage** | [**ClusterResource**](ClusterResource.md) |  | 
**SubscriptionCpuTotal** | **float64** |  | 
**SubscriptionObligationExists** | **float64** |  | 
**SubscriptionSocketTotal** | **float64** |  | 
**Upgrade** | [**ClusterUpgrade**](ClusterUpgrade.md) |  | 

## Methods

### NewOneMetric

`func NewOneMetric(cloudProvider string, clusterType string, computeNodesCpu ClusterResource, computeNodesMemory ClusterResource, computeNodesSockets ClusterResource, consoleUrl string, cpu ClusterResource, criticalAlertsFiring float64, memory ClusterResource, nodes ClusterMetricsNodes, openshiftVersion string, operatingSystem string, operatorsConditionFailing float64, region string, sockets ClusterResource, state string, stateDescription string, storage ClusterResource, subscriptionCpuTotal float64, subscriptionObligationExists float64, subscriptionSocketTotal float64, upgrade ClusterUpgrade, ) *OneMetric`

NewOneMetric instantiates a new OneMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOneMetricWithDefaults

`func NewOneMetricWithDefaults() *OneMetric`

NewOneMetricWithDefaults instantiates a new OneMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *OneMetric) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *OneMetric) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *OneMetric) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetClusterType

`func (o *OneMetric) GetClusterType() string`

GetClusterType returns the ClusterType field if non-nil, zero value otherwise.

### GetClusterTypeOk

`func (o *OneMetric) GetClusterTypeOk() (*string, bool)`

GetClusterTypeOk returns a tuple with the ClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterType

`func (o *OneMetric) SetClusterType(v string)`

SetClusterType sets ClusterType field to given value.


### GetComputeNodesCpu

`func (o *OneMetric) GetComputeNodesCpu() ClusterResource`

GetComputeNodesCpu returns the ComputeNodesCpu field if non-nil, zero value otherwise.

### GetComputeNodesCpuOk

`func (o *OneMetric) GetComputeNodesCpuOk() (*ClusterResource, bool)`

GetComputeNodesCpuOk returns a tuple with the ComputeNodesCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeNodesCpu

`func (o *OneMetric) SetComputeNodesCpu(v ClusterResource)`

SetComputeNodesCpu sets ComputeNodesCpu field to given value.


### GetComputeNodesMemory

`func (o *OneMetric) GetComputeNodesMemory() ClusterResource`

GetComputeNodesMemory returns the ComputeNodesMemory field if non-nil, zero value otherwise.

### GetComputeNodesMemoryOk

`func (o *OneMetric) GetComputeNodesMemoryOk() (*ClusterResource, bool)`

GetComputeNodesMemoryOk returns a tuple with the ComputeNodesMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeNodesMemory

`func (o *OneMetric) SetComputeNodesMemory(v ClusterResource)`

SetComputeNodesMemory sets ComputeNodesMemory field to given value.


### GetComputeNodesSockets

`func (o *OneMetric) GetComputeNodesSockets() ClusterResource`

GetComputeNodesSockets returns the ComputeNodesSockets field if non-nil, zero value otherwise.

### GetComputeNodesSocketsOk

`func (o *OneMetric) GetComputeNodesSocketsOk() (*ClusterResource, bool)`

GetComputeNodesSocketsOk returns a tuple with the ComputeNodesSockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeNodesSockets

`func (o *OneMetric) SetComputeNodesSockets(v ClusterResource)`

SetComputeNodesSockets sets ComputeNodesSockets field to given value.


### GetConsoleUrl

`func (o *OneMetric) GetConsoleUrl() string`

GetConsoleUrl returns the ConsoleUrl field if non-nil, zero value otherwise.

### GetConsoleUrlOk

`func (o *OneMetric) GetConsoleUrlOk() (*string, bool)`

GetConsoleUrlOk returns a tuple with the ConsoleUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleUrl

`func (o *OneMetric) SetConsoleUrl(v string)`

SetConsoleUrl sets ConsoleUrl field to given value.


### GetCpu

`func (o *OneMetric) GetCpu() ClusterResource`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *OneMetric) GetCpuOk() (*ClusterResource, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *OneMetric) SetCpu(v ClusterResource)`

SetCpu sets Cpu field to given value.


### GetCriticalAlertsFiring

`func (o *OneMetric) GetCriticalAlertsFiring() float64`

GetCriticalAlertsFiring returns the CriticalAlertsFiring field if non-nil, zero value otherwise.

### GetCriticalAlertsFiringOk

`func (o *OneMetric) GetCriticalAlertsFiringOk() (*float64, bool)`

GetCriticalAlertsFiringOk returns a tuple with the CriticalAlertsFiring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalAlertsFiring

`func (o *OneMetric) SetCriticalAlertsFiring(v float64)`

SetCriticalAlertsFiring sets CriticalAlertsFiring field to given value.


### GetHealthState

`func (o *OneMetric) GetHealthState() string`

GetHealthState returns the HealthState field if non-nil, zero value otherwise.

### GetHealthStateOk

`func (o *OneMetric) GetHealthStateOk() (*string, bool)`

GetHealthStateOk returns a tuple with the HealthState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthState

`func (o *OneMetric) SetHealthState(v string)`

SetHealthState sets HealthState field to given value.

### HasHealthState

`func (o *OneMetric) HasHealthState() bool`

HasHealthState returns a boolean if a field has been set.

### GetMemory

`func (o *OneMetric) GetMemory() ClusterResource`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *OneMetric) GetMemoryOk() (*ClusterResource, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *OneMetric) SetMemory(v ClusterResource)`

SetMemory sets Memory field to given value.


### GetNodes

`func (o *OneMetric) GetNodes() ClusterMetricsNodes`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *OneMetric) GetNodesOk() (*ClusterMetricsNodes, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *OneMetric) SetNodes(v ClusterMetricsNodes)`

SetNodes sets Nodes field to given value.


### GetOpenshiftVersion

`func (o *OneMetric) GetOpenshiftVersion() string`

GetOpenshiftVersion returns the OpenshiftVersion field if non-nil, zero value otherwise.

### GetOpenshiftVersionOk

`func (o *OneMetric) GetOpenshiftVersionOk() (*string, bool)`

GetOpenshiftVersionOk returns a tuple with the OpenshiftVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenshiftVersion

`func (o *OneMetric) SetOpenshiftVersion(v string)`

SetOpenshiftVersion sets OpenshiftVersion field to given value.


### GetOperatingSystem

`func (o *OneMetric) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *OneMetric) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *OneMetric) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.


### GetOperatorsConditionFailing

`func (o *OneMetric) GetOperatorsConditionFailing() float64`

GetOperatorsConditionFailing returns the OperatorsConditionFailing field if non-nil, zero value otherwise.

### GetOperatorsConditionFailingOk

`func (o *OneMetric) GetOperatorsConditionFailingOk() (*float64, bool)`

GetOperatorsConditionFailingOk returns a tuple with the OperatorsConditionFailing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorsConditionFailing

`func (o *OneMetric) SetOperatorsConditionFailing(v float64)`

SetOperatorsConditionFailing sets OperatorsConditionFailing field to given value.


### GetRegion

`func (o *OneMetric) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *OneMetric) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *OneMetric) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetSockets

`func (o *OneMetric) GetSockets() ClusterResource`

GetSockets returns the Sockets field if non-nil, zero value otherwise.

### GetSocketsOk

`func (o *OneMetric) GetSocketsOk() (*ClusterResource, bool)`

GetSocketsOk returns a tuple with the Sockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockets

`func (o *OneMetric) SetSockets(v ClusterResource)`

SetSockets sets Sockets field to given value.


### GetState

`func (o *OneMetric) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OneMetric) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OneMetric) SetState(v string)`

SetState sets State field to given value.


### GetStateDescription

`func (o *OneMetric) GetStateDescription() string`

GetStateDescription returns the StateDescription field if non-nil, zero value otherwise.

### GetStateDescriptionOk

`func (o *OneMetric) GetStateDescriptionOk() (*string, bool)`

GetStateDescriptionOk returns a tuple with the StateDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateDescription

`func (o *OneMetric) SetStateDescription(v string)`

SetStateDescription sets StateDescription field to given value.


### GetStorage

`func (o *OneMetric) GetStorage() ClusterResource`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *OneMetric) GetStorageOk() (*ClusterResource, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *OneMetric) SetStorage(v ClusterResource)`

SetStorage sets Storage field to given value.


### GetSubscriptionCpuTotal

`func (o *OneMetric) GetSubscriptionCpuTotal() float64`

GetSubscriptionCpuTotal returns the SubscriptionCpuTotal field if non-nil, zero value otherwise.

### GetSubscriptionCpuTotalOk

`func (o *OneMetric) GetSubscriptionCpuTotalOk() (*float64, bool)`

GetSubscriptionCpuTotalOk returns a tuple with the SubscriptionCpuTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionCpuTotal

`func (o *OneMetric) SetSubscriptionCpuTotal(v float64)`

SetSubscriptionCpuTotal sets SubscriptionCpuTotal field to given value.


### GetSubscriptionObligationExists

`func (o *OneMetric) GetSubscriptionObligationExists() float64`

GetSubscriptionObligationExists returns the SubscriptionObligationExists field if non-nil, zero value otherwise.

### GetSubscriptionObligationExistsOk

`func (o *OneMetric) GetSubscriptionObligationExistsOk() (*float64, bool)`

GetSubscriptionObligationExistsOk returns a tuple with the SubscriptionObligationExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionObligationExists

`func (o *OneMetric) SetSubscriptionObligationExists(v float64)`

SetSubscriptionObligationExists sets SubscriptionObligationExists field to given value.


### GetSubscriptionSocketTotal

`func (o *OneMetric) GetSubscriptionSocketTotal() float64`

GetSubscriptionSocketTotal returns the SubscriptionSocketTotal field if non-nil, zero value otherwise.

### GetSubscriptionSocketTotalOk

`func (o *OneMetric) GetSubscriptionSocketTotalOk() (*float64, bool)`

GetSubscriptionSocketTotalOk returns a tuple with the SubscriptionSocketTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionSocketTotal

`func (o *OneMetric) SetSubscriptionSocketTotal(v float64)`

SetSubscriptionSocketTotal sets SubscriptionSocketTotal field to given value.


### GetUpgrade

`func (o *OneMetric) GetUpgrade() ClusterUpgrade`

GetUpgrade returns the Upgrade field if non-nil, zero value otherwise.

### GetUpgradeOk

`func (o *OneMetric) GetUpgradeOk() (*ClusterUpgrade, bool)`

GetUpgradeOk returns a tuple with the Upgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgrade

`func (o *OneMetric) SetUpgrade(v ClusterUpgrade)`

SetUpgrade sets Upgrade field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


