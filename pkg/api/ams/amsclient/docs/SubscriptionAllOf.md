# SubscriptionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | Pointer to [**[]Capability**](Capability.md) |  | [optional] 
**CloudAccountId** | Pointer to **string** |  | [optional] 
**CloudProviderId** | Pointer to **string** |  | [optional] 
**ClusterBillingModel** | Pointer to **string** |  | [optional] 
**ClusterId** | Pointer to **string** |  | [optional] 
**ConsoleUrl** | Pointer to **string** |  | [optional] 
**ConsumerUuid** | Pointer to **string** |  | [optional] 
**CpuTotal** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Creator** | Pointer to [**AccountReference**](AccountReference.md) |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ExternalClusterId** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to [**[]Label**](Label.md) |  | [optional] 
**LastReconcileDate** | Pointer to **time.Time** | Last time this subscription were reconciled about cluster usage | [optional] 
**LastReleasedAt** | Pointer to **time.Time** | Last time status was set to Released for this cluster/subscription in Unix time | [optional] 
**LastTelemetryDate** | Pointer to **time.Time** | Last telemetry authorization request for this cluster/subscription in Unix time | [optional] 
**Managed** | **bool** |  | 
**Metrics** | Pointer to [**[]OneMetric**](OneMetric.md) |  | [optional] 
**NotificationContacts** | Pointer to [**[]Account**](Account.md) |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**Plan** | Pointer to [**Plan**](Plan.md) |  | [optional] 
**ProductBundle** | Pointer to **string** |  | [optional] 
**Provenance** | Pointer to **string** |  | [optional] 
**RegionId** | Pointer to **string** |  | [optional] 
**Released** | Pointer to **bool** |  | [optional] 
**ServiceLevel** | Pointer to **string** |  | [optional] 
**SocketTotal** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SupportLevel** | Pointer to **string** |  | [optional] 
**SystemUnits** | Pointer to **string** |  | [optional] 
**TrialEndDate** | Pointer to **time.Time** | If the subscription is a trial, date the trial ends | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Usage** | Pointer to **string** |  | [optional] 

## Methods

### NewSubscriptionAllOf

`func NewSubscriptionAllOf(managed bool, ) *SubscriptionAllOf`

NewSubscriptionAllOf instantiates a new SubscriptionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionAllOfWithDefaults

`func NewSubscriptionAllOfWithDefaults() *SubscriptionAllOf`

NewSubscriptionAllOfWithDefaults instantiates a new SubscriptionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *SubscriptionAllOf) GetCapabilities() []Capability`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *SubscriptionAllOf) GetCapabilitiesOk() (*[]Capability, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *SubscriptionAllOf) SetCapabilities(v []Capability)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *SubscriptionAllOf) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetCloudAccountId

`func (o *SubscriptionAllOf) GetCloudAccountId() string`

GetCloudAccountId returns the CloudAccountId field if non-nil, zero value otherwise.

### GetCloudAccountIdOk

`func (o *SubscriptionAllOf) GetCloudAccountIdOk() (*string, bool)`

GetCloudAccountIdOk returns a tuple with the CloudAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAccountId

`func (o *SubscriptionAllOf) SetCloudAccountId(v string)`

SetCloudAccountId sets CloudAccountId field to given value.

### HasCloudAccountId

`func (o *SubscriptionAllOf) HasCloudAccountId() bool`

HasCloudAccountId returns a boolean if a field has been set.

### GetCloudProviderId

`func (o *SubscriptionAllOf) GetCloudProviderId() string`

GetCloudProviderId returns the CloudProviderId field if non-nil, zero value otherwise.

### GetCloudProviderIdOk

`func (o *SubscriptionAllOf) GetCloudProviderIdOk() (*string, bool)`

GetCloudProviderIdOk returns a tuple with the CloudProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderId

`func (o *SubscriptionAllOf) SetCloudProviderId(v string)`

SetCloudProviderId sets CloudProviderId field to given value.

### HasCloudProviderId

`func (o *SubscriptionAllOf) HasCloudProviderId() bool`

HasCloudProviderId returns a boolean if a field has been set.

### GetClusterBillingModel

`func (o *SubscriptionAllOf) GetClusterBillingModel() string`

GetClusterBillingModel returns the ClusterBillingModel field if non-nil, zero value otherwise.

### GetClusterBillingModelOk

`func (o *SubscriptionAllOf) GetClusterBillingModelOk() (*string, bool)`

GetClusterBillingModelOk returns a tuple with the ClusterBillingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterBillingModel

`func (o *SubscriptionAllOf) SetClusterBillingModel(v string)`

SetClusterBillingModel sets ClusterBillingModel field to given value.

### HasClusterBillingModel

`func (o *SubscriptionAllOf) HasClusterBillingModel() bool`

HasClusterBillingModel returns a boolean if a field has been set.

### GetClusterId

`func (o *SubscriptionAllOf) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *SubscriptionAllOf) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *SubscriptionAllOf) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *SubscriptionAllOf) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetConsoleUrl

`func (o *SubscriptionAllOf) GetConsoleUrl() string`

GetConsoleUrl returns the ConsoleUrl field if non-nil, zero value otherwise.

### GetConsoleUrlOk

`func (o *SubscriptionAllOf) GetConsoleUrlOk() (*string, bool)`

GetConsoleUrlOk returns a tuple with the ConsoleUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleUrl

`func (o *SubscriptionAllOf) SetConsoleUrl(v string)`

SetConsoleUrl sets ConsoleUrl field to given value.

### HasConsoleUrl

`func (o *SubscriptionAllOf) HasConsoleUrl() bool`

HasConsoleUrl returns a boolean if a field has been set.

### GetConsumerUuid

`func (o *SubscriptionAllOf) GetConsumerUuid() string`

GetConsumerUuid returns the ConsumerUuid field if non-nil, zero value otherwise.

### GetConsumerUuidOk

`func (o *SubscriptionAllOf) GetConsumerUuidOk() (*string, bool)`

GetConsumerUuidOk returns a tuple with the ConsumerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerUuid

`func (o *SubscriptionAllOf) SetConsumerUuid(v string)`

SetConsumerUuid sets ConsumerUuid field to given value.

### HasConsumerUuid

`func (o *SubscriptionAllOf) HasConsumerUuid() bool`

HasConsumerUuid returns a boolean if a field has been set.

### GetCpuTotal

`func (o *SubscriptionAllOf) GetCpuTotal() int32`

GetCpuTotal returns the CpuTotal field if non-nil, zero value otherwise.

### GetCpuTotalOk

`func (o *SubscriptionAllOf) GetCpuTotalOk() (*int32, bool)`

GetCpuTotalOk returns a tuple with the CpuTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuTotal

`func (o *SubscriptionAllOf) SetCpuTotal(v int32)`

SetCpuTotal sets CpuTotal field to given value.

### HasCpuTotal

`func (o *SubscriptionAllOf) HasCpuTotal() bool`

HasCpuTotal returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SubscriptionAllOf) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubscriptionAllOf) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubscriptionAllOf) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SubscriptionAllOf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *SubscriptionAllOf) GetCreator() AccountReference`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *SubscriptionAllOf) GetCreatorOk() (*AccountReference, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *SubscriptionAllOf) SetCreator(v AccountReference)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *SubscriptionAllOf) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDisplayName

`func (o *SubscriptionAllOf) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SubscriptionAllOf) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SubscriptionAllOf) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SubscriptionAllOf) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExternalClusterId

`func (o *SubscriptionAllOf) GetExternalClusterId() string`

GetExternalClusterId returns the ExternalClusterId field if non-nil, zero value otherwise.

### GetExternalClusterIdOk

`func (o *SubscriptionAllOf) GetExternalClusterIdOk() (*string, bool)`

GetExternalClusterIdOk returns a tuple with the ExternalClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalClusterId

`func (o *SubscriptionAllOf) SetExternalClusterId(v string)`

SetExternalClusterId sets ExternalClusterId field to given value.

### HasExternalClusterId

`func (o *SubscriptionAllOf) HasExternalClusterId() bool`

HasExternalClusterId returns a boolean if a field has been set.

### GetLabels

`func (o *SubscriptionAllOf) GetLabels() []Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *SubscriptionAllOf) GetLabelsOk() (*[]Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *SubscriptionAllOf) SetLabels(v []Label)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *SubscriptionAllOf) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLastReconcileDate

`func (o *SubscriptionAllOf) GetLastReconcileDate() time.Time`

GetLastReconcileDate returns the LastReconcileDate field if non-nil, zero value otherwise.

### GetLastReconcileDateOk

`func (o *SubscriptionAllOf) GetLastReconcileDateOk() (*time.Time, bool)`

GetLastReconcileDateOk returns a tuple with the LastReconcileDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReconcileDate

`func (o *SubscriptionAllOf) SetLastReconcileDate(v time.Time)`

SetLastReconcileDate sets LastReconcileDate field to given value.

### HasLastReconcileDate

`func (o *SubscriptionAllOf) HasLastReconcileDate() bool`

HasLastReconcileDate returns a boolean if a field has been set.

### GetLastReleasedAt

`func (o *SubscriptionAllOf) GetLastReleasedAt() time.Time`

GetLastReleasedAt returns the LastReleasedAt field if non-nil, zero value otherwise.

### GetLastReleasedAtOk

`func (o *SubscriptionAllOf) GetLastReleasedAtOk() (*time.Time, bool)`

GetLastReleasedAtOk returns a tuple with the LastReleasedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReleasedAt

`func (o *SubscriptionAllOf) SetLastReleasedAt(v time.Time)`

SetLastReleasedAt sets LastReleasedAt field to given value.

### HasLastReleasedAt

`func (o *SubscriptionAllOf) HasLastReleasedAt() bool`

HasLastReleasedAt returns a boolean if a field has been set.

### GetLastTelemetryDate

`func (o *SubscriptionAllOf) GetLastTelemetryDate() time.Time`

GetLastTelemetryDate returns the LastTelemetryDate field if non-nil, zero value otherwise.

### GetLastTelemetryDateOk

`func (o *SubscriptionAllOf) GetLastTelemetryDateOk() (*time.Time, bool)`

GetLastTelemetryDateOk returns a tuple with the LastTelemetryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTelemetryDate

`func (o *SubscriptionAllOf) SetLastTelemetryDate(v time.Time)`

SetLastTelemetryDate sets LastTelemetryDate field to given value.

### HasLastTelemetryDate

`func (o *SubscriptionAllOf) HasLastTelemetryDate() bool`

HasLastTelemetryDate returns a boolean if a field has been set.

### GetManaged

`func (o *SubscriptionAllOf) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *SubscriptionAllOf) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *SubscriptionAllOf) SetManaged(v bool)`

SetManaged sets Managed field to given value.


### GetMetrics

`func (o *SubscriptionAllOf) GetMetrics() []OneMetric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *SubscriptionAllOf) GetMetricsOk() (*[]OneMetric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *SubscriptionAllOf) SetMetrics(v []OneMetric)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *SubscriptionAllOf) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetNotificationContacts

`func (o *SubscriptionAllOf) GetNotificationContacts() []Account`

GetNotificationContacts returns the NotificationContacts field if non-nil, zero value otherwise.

### GetNotificationContactsOk

`func (o *SubscriptionAllOf) GetNotificationContactsOk() (*[]Account, bool)`

GetNotificationContactsOk returns a tuple with the NotificationContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationContacts

`func (o *SubscriptionAllOf) SetNotificationContacts(v []Account)`

SetNotificationContacts sets NotificationContacts field to given value.

### HasNotificationContacts

`func (o *SubscriptionAllOf) HasNotificationContacts() bool`

HasNotificationContacts returns a boolean if a field has been set.

### GetOrganizationId

`func (o *SubscriptionAllOf) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *SubscriptionAllOf) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *SubscriptionAllOf) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *SubscriptionAllOf) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetPlan

`func (o *SubscriptionAllOf) GetPlan() Plan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *SubscriptionAllOf) GetPlanOk() (*Plan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *SubscriptionAllOf) SetPlan(v Plan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *SubscriptionAllOf) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetProductBundle

`func (o *SubscriptionAllOf) GetProductBundle() string`

GetProductBundle returns the ProductBundle field if non-nil, zero value otherwise.

### GetProductBundleOk

`func (o *SubscriptionAllOf) GetProductBundleOk() (*string, bool)`

GetProductBundleOk returns a tuple with the ProductBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductBundle

`func (o *SubscriptionAllOf) SetProductBundle(v string)`

SetProductBundle sets ProductBundle field to given value.

### HasProductBundle

`func (o *SubscriptionAllOf) HasProductBundle() bool`

HasProductBundle returns a boolean if a field has been set.

### GetProvenance

`func (o *SubscriptionAllOf) GetProvenance() string`

GetProvenance returns the Provenance field if non-nil, zero value otherwise.

### GetProvenanceOk

`func (o *SubscriptionAllOf) GetProvenanceOk() (*string, bool)`

GetProvenanceOk returns a tuple with the Provenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenance

`func (o *SubscriptionAllOf) SetProvenance(v string)`

SetProvenance sets Provenance field to given value.

### HasProvenance

`func (o *SubscriptionAllOf) HasProvenance() bool`

HasProvenance returns a boolean if a field has been set.

### GetRegionId

`func (o *SubscriptionAllOf) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *SubscriptionAllOf) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *SubscriptionAllOf) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *SubscriptionAllOf) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetReleased

`func (o *SubscriptionAllOf) GetReleased() bool`

GetReleased returns the Released field if non-nil, zero value otherwise.

### GetReleasedOk

`func (o *SubscriptionAllOf) GetReleasedOk() (*bool, bool)`

GetReleasedOk returns a tuple with the Released field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleased

`func (o *SubscriptionAllOf) SetReleased(v bool)`

SetReleased sets Released field to given value.

### HasReleased

`func (o *SubscriptionAllOf) HasReleased() bool`

HasReleased returns a boolean if a field has been set.

### GetServiceLevel

`func (o *SubscriptionAllOf) GetServiceLevel() string`

GetServiceLevel returns the ServiceLevel field if non-nil, zero value otherwise.

### GetServiceLevelOk

`func (o *SubscriptionAllOf) GetServiceLevelOk() (*string, bool)`

GetServiceLevelOk returns a tuple with the ServiceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLevel

`func (o *SubscriptionAllOf) SetServiceLevel(v string)`

SetServiceLevel sets ServiceLevel field to given value.

### HasServiceLevel

`func (o *SubscriptionAllOf) HasServiceLevel() bool`

HasServiceLevel returns a boolean if a field has been set.

### GetSocketTotal

`func (o *SubscriptionAllOf) GetSocketTotal() int32`

GetSocketTotal returns the SocketTotal field if non-nil, zero value otherwise.

### GetSocketTotalOk

`func (o *SubscriptionAllOf) GetSocketTotalOk() (*int32, bool)`

GetSocketTotalOk returns a tuple with the SocketTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketTotal

`func (o *SubscriptionAllOf) SetSocketTotal(v int32)`

SetSocketTotal sets SocketTotal field to given value.

### HasSocketTotal

`func (o *SubscriptionAllOf) HasSocketTotal() bool`

HasSocketTotal returns a boolean if a field has been set.

### GetStatus

`func (o *SubscriptionAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriptionAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriptionAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubscriptionAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSupportLevel

`func (o *SubscriptionAllOf) GetSupportLevel() string`

GetSupportLevel returns the SupportLevel field if non-nil, zero value otherwise.

### GetSupportLevelOk

`func (o *SubscriptionAllOf) GetSupportLevelOk() (*string, bool)`

GetSupportLevelOk returns a tuple with the SupportLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLevel

`func (o *SubscriptionAllOf) SetSupportLevel(v string)`

SetSupportLevel sets SupportLevel field to given value.

### HasSupportLevel

`func (o *SubscriptionAllOf) HasSupportLevel() bool`

HasSupportLevel returns a boolean if a field has been set.

### GetSystemUnits

`func (o *SubscriptionAllOf) GetSystemUnits() string`

GetSystemUnits returns the SystemUnits field if non-nil, zero value otherwise.

### GetSystemUnitsOk

`func (o *SubscriptionAllOf) GetSystemUnitsOk() (*string, bool)`

GetSystemUnitsOk returns a tuple with the SystemUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemUnits

`func (o *SubscriptionAllOf) SetSystemUnits(v string)`

SetSystemUnits sets SystemUnits field to given value.

### HasSystemUnits

`func (o *SubscriptionAllOf) HasSystemUnits() bool`

HasSystemUnits returns a boolean if a field has been set.

### GetTrialEndDate

`func (o *SubscriptionAllOf) GetTrialEndDate() time.Time`

GetTrialEndDate returns the TrialEndDate field if non-nil, zero value otherwise.

### GetTrialEndDateOk

`func (o *SubscriptionAllOf) GetTrialEndDateOk() (*time.Time, bool)`

GetTrialEndDateOk returns a tuple with the TrialEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEndDate

`func (o *SubscriptionAllOf) SetTrialEndDate(v time.Time)`

SetTrialEndDate sets TrialEndDate field to given value.

### HasTrialEndDate

`func (o *SubscriptionAllOf) HasTrialEndDate() bool`

HasTrialEndDate returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SubscriptionAllOf) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SubscriptionAllOf) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SubscriptionAllOf) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SubscriptionAllOf) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUsage

`func (o *SubscriptionAllOf) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *SubscriptionAllOf) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *SubscriptionAllOf) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *SubscriptionAllOf) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


