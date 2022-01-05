# Subscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
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

### NewSubscription

`func NewSubscription(managed bool, ) *Subscription`

NewSubscription instantiates a new Subscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionWithDefaults

`func NewSubscriptionWithDefaults() *Subscription`

NewSubscriptionWithDefaults instantiates a new Subscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *Subscription) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Subscription) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Subscription) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Subscription) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *Subscription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Subscription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Subscription) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Subscription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *Subscription) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Subscription) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Subscription) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Subscription) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetCapabilities

`func (o *Subscription) GetCapabilities() []Capability`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *Subscription) GetCapabilitiesOk() (*[]Capability, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *Subscription) SetCapabilities(v []Capability)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *Subscription) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetCloudAccountId

`func (o *Subscription) GetCloudAccountId() string`

GetCloudAccountId returns the CloudAccountId field if non-nil, zero value otherwise.

### GetCloudAccountIdOk

`func (o *Subscription) GetCloudAccountIdOk() (*string, bool)`

GetCloudAccountIdOk returns a tuple with the CloudAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAccountId

`func (o *Subscription) SetCloudAccountId(v string)`

SetCloudAccountId sets CloudAccountId field to given value.

### HasCloudAccountId

`func (o *Subscription) HasCloudAccountId() bool`

HasCloudAccountId returns a boolean if a field has been set.

### GetCloudProviderId

`func (o *Subscription) GetCloudProviderId() string`

GetCloudProviderId returns the CloudProviderId field if non-nil, zero value otherwise.

### GetCloudProviderIdOk

`func (o *Subscription) GetCloudProviderIdOk() (*string, bool)`

GetCloudProviderIdOk returns a tuple with the CloudProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderId

`func (o *Subscription) SetCloudProviderId(v string)`

SetCloudProviderId sets CloudProviderId field to given value.

### HasCloudProviderId

`func (o *Subscription) HasCloudProviderId() bool`

HasCloudProviderId returns a boolean if a field has been set.

### GetClusterBillingModel

`func (o *Subscription) GetClusterBillingModel() string`

GetClusterBillingModel returns the ClusterBillingModel field if non-nil, zero value otherwise.

### GetClusterBillingModelOk

`func (o *Subscription) GetClusterBillingModelOk() (*string, bool)`

GetClusterBillingModelOk returns a tuple with the ClusterBillingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterBillingModel

`func (o *Subscription) SetClusterBillingModel(v string)`

SetClusterBillingModel sets ClusterBillingModel field to given value.

### HasClusterBillingModel

`func (o *Subscription) HasClusterBillingModel() bool`

HasClusterBillingModel returns a boolean if a field has been set.

### GetClusterId

`func (o *Subscription) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *Subscription) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *Subscription) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *Subscription) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetConsoleUrl

`func (o *Subscription) GetConsoleUrl() string`

GetConsoleUrl returns the ConsoleUrl field if non-nil, zero value otherwise.

### GetConsoleUrlOk

`func (o *Subscription) GetConsoleUrlOk() (*string, bool)`

GetConsoleUrlOk returns a tuple with the ConsoleUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleUrl

`func (o *Subscription) SetConsoleUrl(v string)`

SetConsoleUrl sets ConsoleUrl field to given value.

### HasConsoleUrl

`func (o *Subscription) HasConsoleUrl() bool`

HasConsoleUrl returns a boolean if a field has been set.

### GetConsumerUuid

`func (o *Subscription) GetConsumerUuid() string`

GetConsumerUuid returns the ConsumerUuid field if non-nil, zero value otherwise.

### GetConsumerUuidOk

`func (o *Subscription) GetConsumerUuidOk() (*string, bool)`

GetConsumerUuidOk returns a tuple with the ConsumerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerUuid

`func (o *Subscription) SetConsumerUuid(v string)`

SetConsumerUuid sets ConsumerUuid field to given value.

### HasConsumerUuid

`func (o *Subscription) HasConsumerUuid() bool`

HasConsumerUuid returns a boolean if a field has been set.

### GetCpuTotal

`func (o *Subscription) GetCpuTotal() int32`

GetCpuTotal returns the CpuTotal field if non-nil, zero value otherwise.

### GetCpuTotalOk

`func (o *Subscription) GetCpuTotalOk() (*int32, bool)`

GetCpuTotalOk returns a tuple with the CpuTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuTotal

`func (o *Subscription) SetCpuTotal(v int32)`

SetCpuTotal sets CpuTotal field to given value.

### HasCpuTotal

`func (o *Subscription) HasCpuTotal() bool`

HasCpuTotal returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Subscription) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Subscription) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Subscription) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Subscription) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *Subscription) GetCreator() AccountReference`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *Subscription) GetCreatorOk() (*AccountReference, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *Subscription) SetCreator(v AccountReference)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *Subscription) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDisplayName

`func (o *Subscription) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Subscription) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Subscription) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Subscription) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExternalClusterId

`func (o *Subscription) GetExternalClusterId() string`

GetExternalClusterId returns the ExternalClusterId field if non-nil, zero value otherwise.

### GetExternalClusterIdOk

`func (o *Subscription) GetExternalClusterIdOk() (*string, bool)`

GetExternalClusterIdOk returns a tuple with the ExternalClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalClusterId

`func (o *Subscription) SetExternalClusterId(v string)`

SetExternalClusterId sets ExternalClusterId field to given value.

### HasExternalClusterId

`func (o *Subscription) HasExternalClusterId() bool`

HasExternalClusterId returns a boolean if a field has been set.

### GetLabels

`func (o *Subscription) GetLabels() []Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Subscription) GetLabelsOk() (*[]Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Subscription) SetLabels(v []Label)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Subscription) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLastReconcileDate

`func (o *Subscription) GetLastReconcileDate() time.Time`

GetLastReconcileDate returns the LastReconcileDate field if non-nil, zero value otherwise.

### GetLastReconcileDateOk

`func (o *Subscription) GetLastReconcileDateOk() (*time.Time, bool)`

GetLastReconcileDateOk returns a tuple with the LastReconcileDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReconcileDate

`func (o *Subscription) SetLastReconcileDate(v time.Time)`

SetLastReconcileDate sets LastReconcileDate field to given value.

### HasLastReconcileDate

`func (o *Subscription) HasLastReconcileDate() bool`

HasLastReconcileDate returns a boolean if a field has been set.

### GetLastReleasedAt

`func (o *Subscription) GetLastReleasedAt() time.Time`

GetLastReleasedAt returns the LastReleasedAt field if non-nil, zero value otherwise.

### GetLastReleasedAtOk

`func (o *Subscription) GetLastReleasedAtOk() (*time.Time, bool)`

GetLastReleasedAtOk returns a tuple with the LastReleasedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReleasedAt

`func (o *Subscription) SetLastReleasedAt(v time.Time)`

SetLastReleasedAt sets LastReleasedAt field to given value.

### HasLastReleasedAt

`func (o *Subscription) HasLastReleasedAt() bool`

HasLastReleasedAt returns a boolean if a field has been set.

### GetLastTelemetryDate

`func (o *Subscription) GetLastTelemetryDate() time.Time`

GetLastTelemetryDate returns the LastTelemetryDate field if non-nil, zero value otherwise.

### GetLastTelemetryDateOk

`func (o *Subscription) GetLastTelemetryDateOk() (*time.Time, bool)`

GetLastTelemetryDateOk returns a tuple with the LastTelemetryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTelemetryDate

`func (o *Subscription) SetLastTelemetryDate(v time.Time)`

SetLastTelemetryDate sets LastTelemetryDate field to given value.

### HasLastTelemetryDate

`func (o *Subscription) HasLastTelemetryDate() bool`

HasLastTelemetryDate returns a boolean if a field has been set.

### GetManaged

`func (o *Subscription) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *Subscription) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *Subscription) SetManaged(v bool)`

SetManaged sets Managed field to given value.


### GetMetrics

`func (o *Subscription) GetMetrics() []OneMetric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *Subscription) GetMetricsOk() (*[]OneMetric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *Subscription) SetMetrics(v []OneMetric)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *Subscription) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetNotificationContacts

`func (o *Subscription) GetNotificationContacts() []Account`

GetNotificationContacts returns the NotificationContacts field if non-nil, zero value otherwise.

### GetNotificationContactsOk

`func (o *Subscription) GetNotificationContactsOk() (*[]Account, bool)`

GetNotificationContactsOk returns a tuple with the NotificationContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationContacts

`func (o *Subscription) SetNotificationContacts(v []Account)`

SetNotificationContacts sets NotificationContacts field to given value.

### HasNotificationContacts

`func (o *Subscription) HasNotificationContacts() bool`

HasNotificationContacts returns a boolean if a field has been set.

### GetOrganizationId

`func (o *Subscription) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Subscription) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Subscription) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *Subscription) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetPlan

`func (o *Subscription) GetPlan() Plan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *Subscription) GetPlanOk() (*Plan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *Subscription) SetPlan(v Plan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *Subscription) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetProductBundle

`func (o *Subscription) GetProductBundle() string`

GetProductBundle returns the ProductBundle field if non-nil, zero value otherwise.

### GetProductBundleOk

`func (o *Subscription) GetProductBundleOk() (*string, bool)`

GetProductBundleOk returns a tuple with the ProductBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductBundle

`func (o *Subscription) SetProductBundle(v string)`

SetProductBundle sets ProductBundle field to given value.

### HasProductBundle

`func (o *Subscription) HasProductBundle() bool`

HasProductBundle returns a boolean if a field has been set.

### GetProvenance

`func (o *Subscription) GetProvenance() string`

GetProvenance returns the Provenance field if non-nil, zero value otherwise.

### GetProvenanceOk

`func (o *Subscription) GetProvenanceOk() (*string, bool)`

GetProvenanceOk returns a tuple with the Provenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenance

`func (o *Subscription) SetProvenance(v string)`

SetProvenance sets Provenance field to given value.

### HasProvenance

`func (o *Subscription) HasProvenance() bool`

HasProvenance returns a boolean if a field has been set.

### GetRegionId

`func (o *Subscription) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *Subscription) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *Subscription) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *Subscription) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetReleased

`func (o *Subscription) GetReleased() bool`

GetReleased returns the Released field if non-nil, zero value otherwise.

### GetReleasedOk

`func (o *Subscription) GetReleasedOk() (*bool, bool)`

GetReleasedOk returns a tuple with the Released field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleased

`func (o *Subscription) SetReleased(v bool)`

SetReleased sets Released field to given value.

### HasReleased

`func (o *Subscription) HasReleased() bool`

HasReleased returns a boolean if a field has been set.

### GetServiceLevel

`func (o *Subscription) GetServiceLevel() string`

GetServiceLevel returns the ServiceLevel field if non-nil, zero value otherwise.

### GetServiceLevelOk

`func (o *Subscription) GetServiceLevelOk() (*string, bool)`

GetServiceLevelOk returns a tuple with the ServiceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLevel

`func (o *Subscription) SetServiceLevel(v string)`

SetServiceLevel sets ServiceLevel field to given value.

### HasServiceLevel

`func (o *Subscription) HasServiceLevel() bool`

HasServiceLevel returns a boolean if a field has been set.

### GetSocketTotal

`func (o *Subscription) GetSocketTotal() int32`

GetSocketTotal returns the SocketTotal field if non-nil, zero value otherwise.

### GetSocketTotalOk

`func (o *Subscription) GetSocketTotalOk() (*int32, bool)`

GetSocketTotalOk returns a tuple with the SocketTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketTotal

`func (o *Subscription) SetSocketTotal(v int32)`

SetSocketTotal sets SocketTotal field to given value.

### HasSocketTotal

`func (o *Subscription) HasSocketTotal() bool`

HasSocketTotal returns a boolean if a field has been set.

### GetStatus

`func (o *Subscription) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Subscription) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Subscription) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Subscription) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSupportLevel

`func (o *Subscription) GetSupportLevel() string`

GetSupportLevel returns the SupportLevel field if non-nil, zero value otherwise.

### GetSupportLevelOk

`func (o *Subscription) GetSupportLevelOk() (*string, bool)`

GetSupportLevelOk returns a tuple with the SupportLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLevel

`func (o *Subscription) SetSupportLevel(v string)`

SetSupportLevel sets SupportLevel field to given value.

### HasSupportLevel

`func (o *Subscription) HasSupportLevel() bool`

HasSupportLevel returns a boolean if a field has been set.

### GetSystemUnits

`func (o *Subscription) GetSystemUnits() string`

GetSystemUnits returns the SystemUnits field if non-nil, zero value otherwise.

### GetSystemUnitsOk

`func (o *Subscription) GetSystemUnitsOk() (*string, bool)`

GetSystemUnitsOk returns a tuple with the SystemUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemUnits

`func (o *Subscription) SetSystemUnits(v string)`

SetSystemUnits sets SystemUnits field to given value.

### HasSystemUnits

`func (o *Subscription) HasSystemUnits() bool`

HasSystemUnits returns a boolean if a field has been set.

### GetTrialEndDate

`func (o *Subscription) GetTrialEndDate() time.Time`

GetTrialEndDate returns the TrialEndDate field if non-nil, zero value otherwise.

### GetTrialEndDateOk

`func (o *Subscription) GetTrialEndDateOk() (*time.Time, bool)`

GetTrialEndDateOk returns a tuple with the TrialEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEndDate

`func (o *Subscription) SetTrialEndDate(v time.Time)`

SetTrialEndDate sets TrialEndDate field to given value.

### HasTrialEndDate

`func (o *Subscription) HasTrialEndDate() bool`

HasTrialEndDate returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Subscription) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Subscription) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Subscription) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Subscription) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUsage

`func (o *Subscription) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *Subscription) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *Subscription) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *Subscription) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


