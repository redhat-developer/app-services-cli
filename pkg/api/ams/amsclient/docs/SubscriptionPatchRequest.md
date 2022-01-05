# SubscriptionPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudAccountId** | Pointer to **string** |  | [optional] 
**CloudProviderId** | Pointer to **string** |  | [optional] 
**ClusterBillingModel** | Pointer to **string** |  | [optional] 
**ClusterId** | Pointer to **string** |  | [optional] 
**ConsoleUrl** | Pointer to **string** |  | [optional] 
**ConsumerUuid** | Pointer to **string** |  | [optional] 
**CpuTotal** | Pointer to **int32** |  | [optional] 
**CreatorId** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ExternalClusterId** | Pointer to **string** |  | [optional] 
**Managed** | Pointer to **bool** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 
**ProductBundle** | Pointer to **string** |  | [optional] 
**Provenance** | Pointer to **string** |  | [optional] 
**RegionId** | Pointer to **string** |  | [optional] 
**Released** | Pointer to **bool** |  | [optional] 
**ServiceLevel** | Pointer to **string** |  | [optional] 
**SocketTotal** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SupportLevel** | Pointer to **string** |  | [optional] 
**SystemUnits** | Pointer to **string** |  | [optional] 
**TrialEndDate** | Pointer to **time.Time** |  | [optional] 
**Usage** | Pointer to **string** |  | [optional] 

## Methods

### NewSubscriptionPatchRequest

`func NewSubscriptionPatchRequest() *SubscriptionPatchRequest`

NewSubscriptionPatchRequest instantiates a new SubscriptionPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPatchRequestWithDefaults

`func NewSubscriptionPatchRequestWithDefaults() *SubscriptionPatchRequest`

NewSubscriptionPatchRequestWithDefaults instantiates a new SubscriptionPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudAccountId

`func (o *SubscriptionPatchRequest) GetCloudAccountId() string`

GetCloudAccountId returns the CloudAccountId field if non-nil, zero value otherwise.

### GetCloudAccountIdOk

`func (o *SubscriptionPatchRequest) GetCloudAccountIdOk() (*string, bool)`

GetCloudAccountIdOk returns a tuple with the CloudAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAccountId

`func (o *SubscriptionPatchRequest) SetCloudAccountId(v string)`

SetCloudAccountId sets CloudAccountId field to given value.

### HasCloudAccountId

`func (o *SubscriptionPatchRequest) HasCloudAccountId() bool`

HasCloudAccountId returns a boolean if a field has been set.

### GetCloudProviderId

`func (o *SubscriptionPatchRequest) GetCloudProviderId() string`

GetCloudProviderId returns the CloudProviderId field if non-nil, zero value otherwise.

### GetCloudProviderIdOk

`func (o *SubscriptionPatchRequest) GetCloudProviderIdOk() (*string, bool)`

GetCloudProviderIdOk returns a tuple with the CloudProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderId

`func (o *SubscriptionPatchRequest) SetCloudProviderId(v string)`

SetCloudProviderId sets CloudProviderId field to given value.

### HasCloudProviderId

`func (o *SubscriptionPatchRequest) HasCloudProviderId() bool`

HasCloudProviderId returns a boolean if a field has been set.

### GetClusterBillingModel

`func (o *SubscriptionPatchRequest) GetClusterBillingModel() string`

GetClusterBillingModel returns the ClusterBillingModel field if non-nil, zero value otherwise.

### GetClusterBillingModelOk

`func (o *SubscriptionPatchRequest) GetClusterBillingModelOk() (*string, bool)`

GetClusterBillingModelOk returns a tuple with the ClusterBillingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterBillingModel

`func (o *SubscriptionPatchRequest) SetClusterBillingModel(v string)`

SetClusterBillingModel sets ClusterBillingModel field to given value.

### HasClusterBillingModel

`func (o *SubscriptionPatchRequest) HasClusterBillingModel() bool`

HasClusterBillingModel returns a boolean if a field has been set.

### GetClusterId

`func (o *SubscriptionPatchRequest) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *SubscriptionPatchRequest) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *SubscriptionPatchRequest) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *SubscriptionPatchRequest) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetConsoleUrl

`func (o *SubscriptionPatchRequest) GetConsoleUrl() string`

GetConsoleUrl returns the ConsoleUrl field if non-nil, zero value otherwise.

### GetConsoleUrlOk

`func (o *SubscriptionPatchRequest) GetConsoleUrlOk() (*string, bool)`

GetConsoleUrlOk returns a tuple with the ConsoleUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleUrl

`func (o *SubscriptionPatchRequest) SetConsoleUrl(v string)`

SetConsoleUrl sets ConsoleUrl field to given value.

### HasConsoleUrl

`func (o *SubscriptionPatchRequest) HasConsoleUrl() bool`

HasConsoleUrl returns a boolean if a field has been set.

### GetConsumerUuid

`func (o *SubscriptionPatchRequest) GetConsumerUuid() string`

GetConsumerUuid returns the ConsumerUuid field if non-nil, zero value otherwise.

### GetConsumerUuidOk

`func (o *SubscriptionPatchRequest) GetConsumerUuidOk() (*string, bool)`

GetConsumerUuidOk returns a tuple with the ConsumerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerUuid

`func (o *SubscriptionPatchRequest) SetConsumerUuid(v string)`

SetConsumerUuid sets ConsumerUuid field to given value.

### HasConsumerUuid

`func (o *SubscriptionPatchRequest) HasConsumerUuid() bool`

HasConsumerUuid returns a boolean if a field has been set.

### GetCpuTotal

`func (o *SubscriptionPatchRequest) GetCpuTotal() int32`

GetCpuTotal returns the CpuTotal field if non-nil, zero value otherwise.

### GetCpuTotalOk

`func (o *SubscriptionPatchRequest) GetCpuTotalOk() (*int32, bool)`

GetCpuTotalOk returns a tuple with the CpuTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuTotal

`func (o *SubscriptionPatchRequest) SetCpuTotal(v int32)`

SetCpuTotal sets CpuTotal field to given value.

### HasCpuTotal

`func (o *SubscriptionPatchRequest) HasCpuTotal() bool`

HasCpuTotal returns a boolean if a field has been set.

### GetCreatorId

`func (o *SubscriptionPatchRequest) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *SubscriptionPatchRequest) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *SubscriptionPatchRequest) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *SubscriptionPatchRequest) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetDisplayName

`func (o *SubscriptionPatchRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SubscriptionPatchRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SubscriptionPatchRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SubscriptionPatchRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExternalClusterId

`func (o *SubscriptionPatchRequest) GetExternalClusterId() string`

GetExternalClusterId returns the ExternalClusterId field if non-nil, zero value otherwise.

### GetExternalClusterIdOk

`func (o *SubscriptionPatchRequest) GetExternalClusterIdOk() (*string, bool)`

GetExternalClusterIdOk returns a tuple with the ExternalClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalClusterId

`func (o *SubscriptionPatchRequest) SetExternalClusterId(v string)`

SetExternalClusterId sets ExternalClusterId field to given value.

### HasExternalClusterId

`func (o *SubscriptionPatchRequest) HasExternalClusterId() bool`

HasExternalClusterId returns a boolean if a field has been set.

### GetManaged

`func (o *SubscriptionPatchRequest) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *SubscriptionPatchRequest) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *SubscriptionPatchRequest) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *SubscriptionPatchRequest) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetOrganizationId

`func (o *SubscriptionPatchRequest) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *SubscriptionPatchRequest) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *SubscriptionPatchRequest) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *SubscriptionPatchRequest) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetPlanId

`func (o *SubscriptionPatchRequest) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *SubscriptionPatchRequest) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *SubscriptionPatchRequest) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *SubscriptionPatchRequest) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetProductBundle

`func (o *SubscriptionPatchRequest) GetProductBundle() string`

GetProductBundle returns the ProductBundle field if non-nil, zero value otherwise.

### GetProductBundleOk

`func (o *SubscriptionPatchRequest) GetProductBundleOk() (*string, bool)`

GetProductBundleOk returns a tuple with the ProductBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductBundle

`func (o *SubscriptionPatchRequest) SetProductBundle(v string)`

SetProductBundle sets ProductBundle field to given value.

### HasProductBundle

`func (o *SubscriptionPatchRequest) HasProductBundle() bool`

HasProductBundle returns a boolean if a field has been set.

### GetProvenance

`func (o *SubscriptionPatchRequest) GetProvenance() string`

GetProvenance returns the Provenance field if non-nil, zero value otherwise.

### GetProvenanceOk

`func (o *SubscriptionPatchRequest) GetProvenanceOk() (*string, bool)`

GetProvenanceOk returns a tuple with the Provenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenance

`func (o *SubscriptionPatchRequest) SetProvenance(v string)`

SetProvenance sets Provenance field to given value.

### HasProvenance

`func (o *SubscriptionPatchRequest) HasProvenance() bool`

HasProvenance returns a boolean if a field has been set.

### GetRegionId

`func (o *SubscriptionPatchRequest) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *SubscriptionPatchRequest) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *SubscriptionPatchRequest) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *SubscriptionPatchRequest) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetReleased

`func (o *SubscriptionPatchRequest) GetReleased() bool`

GetReleased returns the Released field if non-nil, zero value otherwise.

### GetReleasedOk

`func (o *SubscriptionPatchRequest) GetReleasedOk() (*bool, bool)`

GetReleasedOk returns a tuple with the Released field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleased

`func (o *SubscriptionPatchRequest) SetReleased(v bool)`

SetReleased sets Released field to given value.

### HasReleased

`func (o *SubscriptionPatchRequest) HasReleased() bool`

HasReleased returns a boolean if a field has been set.

### GetServiceLevel

`func (o *SubscriptionPatchRequest) GetServiceLevel() string`

GetServiceLevel returns the ServiceLevel field if non-nil, zero value otherwise.

### GetServiceLevelOk

`func (o *SubscriptionPatchRequest) GetServiceLevelOk() (*string, bool)`

GetServiceLevelOk returns a tuple with the ServiceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLevel

`func (o *SubscriptionPatchRequest) SetServiceLevel(v string)`

SetServiceLevel sets ServiceLevel field to given value.

### HasServiceLevel

`func (o *SubscriptionPatchRequest) HasServiceLevel() bool`

HasServiceLevel returns a boolean if a field has been set.

### GetSocketTotal

`func (o *SubscriptionPatchRequest) GetSocketTotal() int32`

GetSocketTotal returns the SocketTotal field if non-nil, zero value otherwise.

### GetSocketTotalOk

`func (o *SubscriptionPatchRequest) GetSocketTotalOk() (*int32, bool)`

GetSocketTotalOk returns a tuple with the SocketTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketTotal

`func (o *SubscriptionPatchRequest) SetSocketTotal(v int32)`

SetSocketTotal sets SocketTotal field to given value.

### HasSocketTotal

`func (o *SubscriptionPatchRequest) HasSocketTotal() bool`

HasSocketTotal returns a boolean if a field has been set.

### GetStatus

`func (o *SubscriptionPatchRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriptionPatchRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriptionPatchRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubscriptionPatchRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSupportLevel

`func (o *SubscriptionPatchRequest) GetSupportLevel() string`

GetSupportLevel returns the SupportLevel field if non-nil, zero value otherwise.

### GetSupportLevelOk

`func (o *SubscriptionPatchRequest) GetSupportLevelOk() (*string, bool)`

GetSupportLevelOk returns a tuple with the SupportLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLevel

`func (o *SubscriptionPatchRequest) SetSupportLevel(v string)`

SetSupportLevel sets SupportLevel field to given value.

### HasSupportLevel

`func (o *SubscriptionPatchRequest) HasSupportLevel() bool`

HasSupportLevel returns a boolean if a field has been set.

### GetSystemUnits

`func (o *SubscriptionPatchRequest) GetSystemUnits() string`

GetSystemUnits returns the SystemUnits field if non-nil, zero value otherwise.

### GetSystemUnitsOk

`func (o *SubscriptionPatchRequest) GetSystemUnitsOk() (*string, bool)`

GetSystemUnitsOk returns a tuple with the SystemUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemUnits

`func (o *SubscriptionPatchRequest) SetSystemUnits(v string)`

SetSystemUnits sets SystemUnits field to given value.

### HasSystemUnits

`func (o *SubscriptionPatchRequest) HasSystemUnits() bool`

HasSystemUnits returns a boolean if a field has been set.

### GetTrialEndDate

`func (o *SubscriptionPatchRequest) GetTrialEndDate() time.Time`

GetTrialEndDate returns the TrialEndDate field if non-nil, zero value otherwise.

### GetTrialEndDateOk

`func (o *SubscriptionPatchRequest) GetTrialEndDateOk() (*time.Time, bool)`

GetTrialEndDateOk returns a tuple with the TrialEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEndDate

`func (o *SubscriptionPatchRequest) SetTrialEndDate(v time.Time)`

SetTrialEndDate sets TrialEndDate field to given value.

### HasTrialEndDate

`func (o *SubscriptionPatchRequest) HasTrialEndDate() bool`

HasTrialEndDate returns a boolean if a field has been set.

### GetUsage

`func (o *SubscriptionPatchRequest) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *SubscriptionPatchRequest) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *SubscriptionPatchRequest) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *SubscriptionPatchRequest) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


