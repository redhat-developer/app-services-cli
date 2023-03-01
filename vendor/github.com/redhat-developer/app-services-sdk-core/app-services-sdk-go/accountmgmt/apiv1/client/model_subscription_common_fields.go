/*
 * Account Management Service API
 *
 * Manage user subscriptions and clusters
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package accountmgmtclient

import (
	"encoding/json"
	"time"
)

// SubscriptionCommonFields struct for SubscriptionCommonFields
type SubscriptionCommonFields struct {
	Href *string `json:"href,omitempty"`
	Id *string `json:"id,omitempty"`
	Kind *string `json:"kind,omitempty"`
	// If set, the date the subscription expires based on the billing model
	BillingExpirationDate *time.Time `json:"billing_expiration_date,omitempty"`
	BillingMarketplaceAccount *string `json:"billing_marketplace_account,omitempty"`
	CloudAccountId *string `json:"cloud_account_id,omitempty"`
	CloudProviderId *string `json:"cloud_provider_id,omitempty"`
	ClusterBillingModel *string `json:"cluster_billing_model,omitempty"`
	ClusterId *string `json:"cluster_id,omitempty"`
	ConsoleUrl *string `json:"console_url,omitempty"`
	ConsumerUuid *string `json:"consumer_uuid,omitempty"`
	CpuTotal *int32 `json:"cpu_total,omitempty"`
	CreatorId *string `json:"creator_id,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	ExternalClusterId *string `json:"external_cluster_id,omitempty"`
	// Last time this subscription were reconciled about cluster usage
	LastReconcileDate *time.Time `json:"last_reconcile_date,omitempty"`
	// Last time status was set to Released for this cluster/subscription in Unix time
	LastReleasedAt *time.Time `json:"last_released_at,omitempty"`
	// Last telemetry authorization request for this cluster/subscription in Unix time
	LastTelemetryDate *time.Time `json:"last_telemetry_date,omitempty"`
	Managed bool `json:"managed"`
	OrganizationId *string `json:"organization_id,omitempty"`
	PlanId *string `json:"plan_id,omitempty"`
	ProductBundle *string `json:"product_bundle,omitempty"`
	Provenance *string `json:"provenance,omitempty"`
	RegionId *string `json:"region_id,omitempty"`
	Released *bool `json:"released,omitempty"`
	ServiceLevel *string `json:"service_level,omitempty"`
	SocketTotal *int32 `json:"socket_total,omitempty"`
	Status *string `json:"status,omitempty"`
	SupportLevel *string `json:"support_level,omitempty"`
	SystemUnits *string `json:"system_units,omitempty"`
	// If the subscription is a trial, date the trial ends
	TrialEndDate *time.Time `json:"trial_end_date,omitempty"`
	Usage *string `json:"usage,omitempty"`
}

// NewSubscriptionCommonFields instantiates a new SubscriptionCommonFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionCommonFields(managed bool) *SubscriptionCommonFields {
	this := SubscriptionCommonFields{}
	this.Managed = managed
	return &this
}

// NewSubscriptionCommonFieldsWithDefaults instantiates a new SubscriptionCommonFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionCommonFieldsWithDefaults() *SubscriptionCommonFields {
	this := SubscriptionCommonFields{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *SubscriptionCommonFields) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCommonFields) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *SubscriptionCommonFields) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *SubscriptionCommonFields) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SubscriptionCommonFields) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCommonFields) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SubscriptionCommonFields) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SubscriptionCommonFields) SetId(v string) {
	o.Id = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *SubscriptionCommonFields) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCommonFields) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *SubscriptionCommonFields) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *SubscriptionCommonFields) SetKind(v string) {
	o.Kind = &v
}

// GetBillingExpirationDate returns the BillingExpirationDate field value if set, zero value otherwise.
func (o *SubscriptionCommonFields) GetBillingExpirationDate() time.Time {
	if o == nil || o.BillingExpirationDate == nil {
		var ret time.Time
		return ret
	}
	return *o.BillingExpirationDate
}

// GetBillingExpirationDateOk returns a tuple with the BillingExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCommonFields) GetBillingExpirationDateOk() (*time.Time, bool) {
	if o == nil || o.BillingExpirationDate == nil {
		return nil, false
	}
	return o.BillingExpirationDate, true
}

// HasBillingExpirationDate returns a boolean if a field has been set.
func (o *SubscriptionCommonFields) HasBillingExpirationDate() bool {
	if o != nil && o.BillingExpirationDate != nil {
		return true
	}

	return false
}

// SetBillingExpirationDate gets a reference to the given time.Time and assigns it to the BillingExpirationDate field.
func (o *SubscriptionCommonFields) SetBillingExpirationDate(v time.Time) {
	o.BillingExpirationDate = &v
}

// GetBillingMarketplaceAccount returns the BillingMarketplaceAccount field value if set, zero value otherwise.
func (o *SubscriptionCommonFields) GetBillingMarketplaceAccount() string {
	if o == nil || o.BillingMarketplaceAccount == nil {
		var ret string
		return ret
	}
	return *o.BillingMarketplaceAccount
}

// GetBillingMarketplaceAccountOk returns a tuple with the BillingMarketplaceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCommonFields) GetBillingMarketplaceAccountOk() (*string, bool) {
	if o == nil || o.BillingMarketplaceAccount == nil {
		return nil, false
	}
	return o.BillingMarketplaceAccount, true
}

// HasBillingMarketplaceAccount returns a boolean if a field has been set.
func (o *SubscriptionCommonFields) HasBillingMarketplaceAccount() bool {
	if o != nil && o.BillingMarketplaceAccount != nil {
		return true
	}

	return false
}

// SetBillingMarketplaceAccount gets a reference to the given string and assigns it to the BillingMarketplaceAccount field.
func (o *SubscriptionCommonFields) SetBillingMarketplaceAccount(v string) {
	o.BillingMarketplaceAccount = &v
}

// GetCloudAccountId returns the CloudAccountId field value if set, zero value otherwise.
func (o *SubscriptionCommonFields) GetCloudAccountId() string {
	if o == nil || o.CloudAccountId == nil {
		var ret string
		return ret
	}
	return *o.CloudAccountId
}

// GetCloudAccountIdOk returns a tuple with the CloudAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCommonFields) GetCloudAccountIdOk() (*string, bool) {
	if o == nil || o.CloudAccountId == nil {
		return nil, false
	}
	return o.CloudAccountId, true
}

// HasCloudAccountId returns a boolean if a field has been set.
func (o *SubscriptionCommonFields) HasCloudAccountId() bool {
	if o != nil && o.CloudAccountId != nil {
		return true
	}

	return false
}

// SetCloudAccountId gets a reference to the given string and assigns it to the CloudAccountId field.
func (o *SubscriptionCommonFields) SetCloudAccountId(v string) {
	o.CloudAccountId = &v
}

// GetCloudProviderId returns the CloudProviderId field value if set, zero value otherwise.
func (o *SubscriptionCommonFields) GetCloudProviderId() string {
	if o == nil || o.CloudProviderId == nil {
		var ret string
		return ret
	}
	return *o.CloudProviderId
}

// GetCloudProviderIdOk returns a tuple with the CloudProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCommonFields) GetCloudProviderIdOk() (*string, bool) {
	if o == nil || o.CloudProviderId == nil {
		return nil, false
	}
	return o.CloudProviderId, true
}

// HasCloudProviderId returns a boolean if a field has been set.
func (o *SubscriptionCommonFields) HasCloudProviderId() bool {
	if o != nil && o.CloudProviderId != nil {
		return true
	}

	return false
}

// SetCloudProviderId gets a reference to the given string and assigns it to the CloudProviderId field.
func (o *SubscriptionCommonFields) SetCloudProviderId(v string) {
	o.CloudProviderId = &v
}

// GetClusterBillingModel returns the ClusterBillingModel field value if set, zero value otherwise.
func (o *SubscriptionCommonFields) GetClusterBillingModel() string {
	if o == nil || o.ClusterBillingModel == nil {
		var ret string
		return ret
	}
	return *o.ClusterBillingModel
}

// GetClusterBillingModelOk returns a tuple with the ClusterBillingModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCommonFields) GetClusterBillingModelOk() (*string, bool) {
	if o == nil || o.ClusterBillingModel == nil {
		return nil, false
	}
	return o.ClusterBillingModel, true
}

// HasClusterBillingModel returns a boolean if a field has been set.
func (o *SubscriptionCommonFields) HasClusterBillingModel() bool {
	if o != nil && o.ClusterBillingModel != nil {
		return true
	}

	return false
}

// SetClusterBillingModel gets a reference to the given string and assigns it to the ClusterBillingModel field.
func (o *SubscriptionCommonFields) SetClusterBillingModel(v string) {
	o.ClusterBillingModel = &v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *SubscriptionCommonFields) GetClusterId() string {
	if o == nil || o.ClusterId == nil {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCommonFields) GetClusterIdOk() (*string, bool) {
	if o == nil || o.ClusterId == nil {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *SubscriptionCommonFields) HasClusterId() bool {
	if o != nil && o.ClusterId != nil {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *SubscriptionCommonFields) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetConsoleUrl returns the ConsoleUrl field value if set, zero value otherwise.
func (o *SubscriptionCommonFields) GetConsoleUrl() string {
	if o == nil || o.ConsoleUrl == nil {
		var ret string
		return ret
	}
	return *o.ConsoleUrl
}

// GetConsoleUrlOk returns a tuple with the ConsoleUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCommonFields) GetConsoleUrlOk() (*string, bool) {
	if o == nil || o.ConsoleUrl == nil {
		return nil, false
	}
	return o.ConsoleUrl, true
}

// HasConsoleUrl returns a boolean if a field has been set.
func (o *SubscriptionCommonFields) HasConsoleUrl() bool {
	if o != nil && o.ConsoleUrl != nil {
		return true
	}

	return false
}

// SetConsoleUrl gets a reference to the given string and assigns it to the ConsoleUrl field.
func (o *SubscriptionCommonFields) SetConsoleUrl(v string) {
	o.ConsoleUrl = &v
}

// GetConsumerUuid returns the ConsumerUuid field value if set, zero value otherwise.
func (o *SubscriptionCommonFields) GetConsumerUuid() string {
	if o == nil || o.ConsumerUuid == nil {
		var ret string
		return ret
	}
	return *o.ConsumerUuid
}

// GetConsumerUuidOk returns a tuple with the ConsumerUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCommonFields) GetConsumerUuidOk() (*string, bool) {
	if o == nil || o.ConsumerUuid == nil {
		return nil, false
	}
	return o.ConsumerUuid, true
}

// HasConsumerUuid returns a boolean if a field has been set.
func (o *SubscriptionCommonFields) HasConsumerUuid() bool {
	if o != nil && o.ConsumerUuid != nil {
		return true
	}

	return false
}

// SetConsumerUuid gets a reference to the given string and assigns it to the ConsumerUuid field.
func (o *SubscriptionCommonFields) SetConsumerUuid(v string) {
	o.ConsumerUuid = &v
}

// GetCpuTotal returns the CpuTotal field value if set, zero value otherwise.
func (o *SubscriptionCommonFields) GetCpuTotal() int32 {
	if o == nil || o.CpuTotal == nil {
		var ret int32
		return ret
	}
	return *o.CpuTotal
}

// GetCpuTotalOk returns a tuple with the CpuTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCommonFields) GetCpuTotalOk() (*int32, bool) {
	if o == nil || o.CpuTotal == nil {
		return nil, false
	}
	return o.CpuTotal, true
}

// HasCpuTotal returns a boolean if a field has been set.
func (o *SubscriptionCommonFields) HasCpuTotal() bool {
	if o != nil && o.CpuTotal != nil {
		return true
	}

	return false
}

// SetCpuTotal gets a reference to the given int32 and assigns it to the CpuTotal field.
func (o *SubscriptionCommonFields) SetCpuTotal(v int32) {
	o.CpuTotal = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *SubscriptionCommonFields) GetCreatorId() string {
	if o == nil || o.CreatorId == nil {
		var ret string
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCommonFields) GetCreatorIdOk() (*string, bool) {
	if o == nil || o.CreatorId == nil {
		return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *SubscriptionCommonFields) HasCreatorId() bool {
	if o != nil && o.CreatorId != nil {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given string and assigns it to the CreatorId field.
func (o *SubscriptionCommonFields) SetCreatorId(v string) {
	o.CreatorId = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *SubscriptionCommonFields) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCommonFields) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *SubscriptionCommonFields) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *SubscriptionCommonFields) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetExternalClusterId returns the ExternalClusterId field value if set, zero value otherwise.
func (o *SubscriptionCommonFields) GetExternalClusterId() string {
	if o == nil || o.ExternalClusterId == nil {
		var ret string
		return ret
	}
	return *o.ExternalClusterId
}

// GetExternalClusterIdOk returns a tuple with the ExternalClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCommonFields) GetExternalClusterIdOk() (*string, bool) {
	if o == nil || o.ExternalClusterId == nil {
		return nil, false
	}
	return o.ExternalClusterId, true
}

// HasExternalClusterId returns a boolean if a field has been set.
func (o *SubscriptionCommonFields) HasExternalClusterId() bool {
	if o != nil && o.ExternalClusterId != nil {
		return true
	}

	return false
}

// SetExternalClusterId gets a reference to the given string and assigns it to the ExternalClusterId field.
func (o *SubscriptionCommonFields) SetExternalClusterId(v string) {
	o.ExternalClusterId = &v
}

// GetLastReconcileDate returns the LastReconcileDate field value if set, zero value otherwise.
func (o *SubscriptionCommonFields) GetLastReconcileDate() time.Time {
	if o == nil || o.LastReconcileDate == nil {
		var ret time.Time
		return ret
	}
	return *o.LastReconcileDate
}

// GetLastReconcileDateOk returns a tuple with the LastReconcileDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCommonFields) GetLastReconcileDateOk() (*time.Time, bool) {
	if o == nil || o.LastReconcileDate == nil {
		return nil, false
	}
	return o.LastReconcileDate, true
}

// HasLastReconcileDate returns a boolean if a field has been set.
func (o *SubscriptionCommonFields) HasLastReconcileDate() bool {
	if o != nil && o.LastReconcileDate != nil {
		return true
	}

	return false
}

// SetLastReconcileDate gets a reference to the given time.Time and assigns it to the LastReconcileDate field.
func (o *SubscriptionCommonFields) SetLastReconcileDate(v time.Time) {
	o.LastReconcileDate = &v
}

// GetLastReleasedAt returns the LastReleasedAt field value if set, zero value otherwise.
func (o *SubscriptionCommonFields) GetLastReleasedAt() time.Time {
	if o == nil || o.LastReleasedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastReleasedAt
}

// GetLastReleasedAtOk returns a tuple with the LastReleasedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCommonFields) GetLastReleasedAtOk() (*time.Time, bool) {
	if o == nil || o.LastReleasedAt == nil {
		return nil, false
	}
	return o.LastReleasedAt, true
}

// HasLastReleasedAt returns a boolean if a field has been set.
func (o *SubscriptionCommonFields) HasLastReleasedAt() bool {
	if o != nil && o.LastReleasedAt != nil {
		return true
	}

	return false
}

// SetLastReleasedAt gets a reference to the given time.Time and assigns it to the LastReleasedAt field.
func (o *SubscriptionCommonFields) SetLastReleasedAt(v time.Time) {
	o.LastReleasedAt = &v
}

// GetLastTelemetryDate returns the LastTelemetryDate field value if set, zero value otherwise.
func (o *SubscriptionCommonFields) GetLastTelemetryDate() time.Time {
	if o == nil || o.LastTelemetryDate == nil {
		var ret time.Time
		return ret
	}
	return *o.LastTelemetryDate
}

// GetLastTelemetryDateOk returns a tuple with the LastTelemetryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCommonFields) GetLastTelemetryDateOk() (*time.Time, bool) {
	if o == nil || o.LastTelemetryDate == nil {
		return nil, false
	}
	return o.LastTelemetryDate, true
}

// HasLastTelemetryDate returns a boolean if a field has been set.
func (o *SubscriptionCommonFields) HasLastTelemetryDate() bool {
	if o != nil && o.LastTelemetryDate != nil {
		return true
	}

	return false
}

// SetLastTelemetryDate gets a reference to the given time.Time and assigns it to the LastTelemetryDate field.
func (o *SubscriptionCommonFields) SetLastTelemetryDate(v time.Time) {
	o.LastTelemetryDate = &v
}

// GetManaged returns the Managed field value
func (o *SubscriptionCommonFields) GetManaged() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Managed
}

// GetManagedOk returns a tuple with the Managed field value
// and a boolean to check if the value has been set.
func (o *SubscriptionCommonFields) GetManagedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Managed, true
}

// SetManaged sets field value
func (o *SubscriptionCommonFields) SetManaged(v bool) {
	o.Managed = v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *SubscriptionCommonFields) GetOrganizationId() string {
	if o == nil || o.OrganizationId == nil {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCommonFields) GetOrganizationIdOk() (*string, bool) {
	if o == nil || o.OrganizationId == nil {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *SubscriptionCommonFields) HasOrganizationId() bool {
	if o != nil && o.OrganizationId != nil {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *SubscriptionCommonFields) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *SubscriptionCommonFields) GetPlanId() string {
	if o == nil || o.PlanId == nil {
		var ret string
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCommonFields) GetPlanIdOk() (*string, bool) {
	if o == nil || o.PlanId == nil {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *SubscriptionCommonFields) HasPlanId() bool {
	if o != nil && o.PlanId != nil {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *SubscriptionCommonFields) SetPlanId(v string) {
	o.PlanId = &v
}

// GetProductBundle returns the ProductBundle field value if set, zero value otherwise.
func (o *SubscriptionCommonFields) GetProductBundle() string {
	if o == nil || o.ProductBundle == nil {
		var ret string
		return ret
	}
	return *o.ProductBundle
}

// GetProductBundleOk returns a tuple with the ProductBundle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCommonFields) GetProductBundleOk() (*string, bool) {
	if o == nil || o.ProductBundle == nil {
		return nil, false
	}
	return o.ProductBundle, true
}

// HasProductBundle returns a boolean if a field has been set.
func (o *SubscriptionCommonFields) HasProductBundle() bool {
	if o != nil && o.ProductBundle != nil {
		return true
	}

	return false
}

// SetProductBundle gets a reference to the given string and assigns it to the ProductBundle field.
func (o *SubscriptionCommonFields) SetProductBundle(v string) {
	o.ProductBundle = &v
}

// GetProvenance returns the Provenance field value if set, zero value otherwise.
func (o *SubscriptionCommonFields) GetProvenance() string {
	if o == nil || o.Provenance == nil {
		var ret string
		return ret
	}
	return *o.Provenance
}

// GetProvenanceOk returns a tuple with the Provenance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCommonFields) GetProvenanceOk() (*string, bool) {
	if o == nil || o.Provenance == nil {
		return nil, false
	}
	return o.Provenance, true
}

// HasProvenance returns a boolean if a field has been set.
func (o *SubscriptionCommonFields) HasProvenance() bool {
	if o != nil && o.Provenance != nil {
		return true
	}

	return false
}

// SetProvenance gets a reference to the given string and assigns it to the Provenance field.
func (o *SubscriptionCommonFields) SetProvenance(v string) {
	o.Provenance = &v
}

// GetRegionId returns the RegionId field value if set, zero value otherwise.
func (o *SubscriptionCommonFields) GetRegionId() string {
	if o == nil || o.RegionId == nil {
		var ret string
		return ret
	}
	return *o.RegionId
}

// GetRegionIdOk returns a tuple with the RegionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCommonFields) GetRegionIdOk() (*string, bool) {
	if o == nil || o.RegionId == nil {
		return nil, false
	}
	return o.RegionId, true
}

// HasRegionId returns a boolean if a field has been set.
func (o *SubscriptionCommonFields) HasRegionId() bool {
	if o != nil && o.RegionId != nil {
		return true
	}

	return false
}

// SetRegionId gets a reference to the given string and assigns it to the RegionId field.
func (o *SubscriptionCommonFields) SetRegionId(v string) {
	o.RegionId = &v
}

// GetReleased returns the Released field value if set, zero value otherwise.
func (o *SubscriptionCommonFields) GetReleased() bool {
	if o == nil || o.Released == nil {
		var ret bool
		return ret
	}
	return *o.Released
}

// GetReleasedOk returns a tuple with the Released field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCommonFields) GetReleasedOk() (*bool, bool) {
	if o == nil || o.Released == nil {
		return nil, false
	}
	return o.Released, true
}

// HasReleased returns a boolean if a field has been set.
func (o *SubscriptionCommonFields) HasReleased() bool {
	if o != nil && o.Released != nil {
		return true
	}

	return false
}

// SetReleased gets a reference to the given bool and assigns it to the Released field.
func (o *SubscriptionCommonFields) SetReleased(v bool) {
	o.Released = &v
}

// GetServiceLevel returns the ServiceLevel field value if set, zero value otherwise.
func (o *SubscriptionCommonFields) GetServiceLevel() string {
	if o == nil || o.ServiceLevel == nil {
		var ret string
		return ret
	}
	return *o.ServiceLevel
}

// GetServiceLevelOk returns a tuple with the ServiceLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCommonFields) GetServiceLevelOk() (*string, bool) {
	if o == nil || o.ServiceLevel == nil {
		return nil, false
	}
	return o.ServiceLevel, true
}

// HasServiceLevel returns a boolean if a field has been set.
func (o *SubscriptionCommonFields) HasServiceLevel() bool {
	if o != nil && o.ServiceLevel != nil {
		return true
	}

	return false
}

// SetServiceLevel gets a reference to the given string and assigns it to the ServiceLevel field.
func (o *SubscriptionCommonFields) SetServiceLevel(v string) {
	o.ServiceLevel = &v
}

// GetSocketTotal returns the SocketTotal field value if set, zero value otherwise.
func (o *SubscriptionCommonFields) GetSocketTotal() int32 {
	if o == nil || o.SocketTotal == nil {
		var ret int32
		return ret
	}
	return *o.SocketTotal
}

// GetSocketTotalOk returns a tuple with the SocketTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCommonFields) GetSocketTotalOk() (*int32, bool) {
	if o == nil || o.SocketTotal == nil {
		return nil, false
	}
	return o.SocketTotal, true
}

// HasSocketTotal returns a boolean if a field has been set.
func (o *SubscriptionCommonFields) HasSocketTotal() bool {
	if o != nil && o.SocketTotal != nil {
		return true
	}

	return false
}

// SetSocketTotal gets a reference to the given int32 and assigns it to the SocketTotal field.
func (o *SubscriptionCommonFields) SetSocketTotal(v int32) {
	o.SocketTotal = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SubscriptionCommonFields) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCommonFields) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SubscriptionCommonFields) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SubscriptionCommonFields) SetStatus(v string) {
	o.Status = &v
}

// GetSupportLevel returns the SupportLevel field value if set, zero value otherwise.
func (o *SubscriptionCommonFields) GetSupportLevel() string {
	if o == nil || o.SupportLevel == nil {
		var ret string
		return ret
	}
	return *o.SupportLevel
}

// GetSupportLevelOk returns a tuple with the SupportLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCommonFields) GetSupportLevelOk() (*string, bool) {
	if o == nil || o.SupportLevel == nil {
		return nil, false
	}
	return o.SupportLevel, true
}

// HasSupportLevel returns a boolean if a field has been set.
func (o *SubscriptionCommonFields) HasSupportLevel() bool {
	if o != nil && o.SupportLevel != nil {
		return true
	}

	return false
}

// SetSupportLevel gets a reference to the given string and assigns it to the SupportLevel field.
func (o *SubscriptionCommonFields) SetSupportLevel(v string) {
	o.SupportLevel = &v
}

// GetSystemUnits returns the SystemUnits field value if set, zero value otherwise.
func (o *SubscriptionCommonFields) GetSystemUnits() string {
	if o == nil || o.SystemUnits == nil {
		var ret string
		return ret
	}
	return *o.SystemUnits
}

// GetSystemUnitsOk returns a tuple with the SystemUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCommonFields) GetSystemUnitsOk() (*string, bool) {
	if o == nil || o.SystemUnits == nil {
		return nil, false
	}
	return o.SystemUnits, true
}

// HasSystemUnits returns a boolean if a field has been set.
func (o *SubscriptionCommonFields) HasSystemUnits() bool {
	if o != nil && o.SystemUnits != nil {
		return true
	}

	return false
}

// SetSystemUnits gets a reference to the given string and assigns it to the SystemUnits field.
func (o *SubscriptionCommonFields) SetSystemUnits(v string) {
	o.SystemUnits = &v
}

// GetTrialEndDate returns the TrialEndDate field value if set, zero value otherwise.
func (o *SubscriptionCommonFields) GetTrialEndDate() time.Time {
	if o == nil || o.TrialEndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.TrialEndDate
}

// GetTrialEndDateOk returns a tuple with the TrialEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCommonFields) GetTrialEndDateOk() (*time.Time, bool) {
	if o == nil || o.TrialEndDate == nil {
		return nil, false
	}
	return o.TrialEndDate, true
}

// HasTrialEndDate returns a boolean if a field has been set.
func (o *SubscriptionCommonFields) HasTrialEndDate() bool {
	if o != nil && o.TrialEndDate != nil {
		return true
	}

	return false
}

// SetTrialEndDate gets a reference to the given time.Time and assigns it to the TrialEndDate field.
func (o *SubscriptionCommonFields) SetTrialEndDate(v time.Time) {
	o.TrialEndDate = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *SubscriptionCommonFields) GetUsage() string {
	if o == nil || o.Usage == nil {
		var ret string
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCommonFields) GetUsageOk() (*string, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *SubscriptionCommonFields) HasUsage() bool {
	if o != nil && o.Usage != nil {
		return true
	}

	return false
}

// SetUsage gets a reference to the given string and assigns it to the Usage field.
func (o *SubscriptionCommonFields) SetUsage(v string) {
	o.Usage = &v
}

func (o SubscriptionCommonFields) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.BillingExpirationDate != nil {
		toSerialize["billing_expiration_date"] = o.BillingExpirationDate
	}
	if o.BillingMarketplaceAccount != nil {
		toSerialize["billing_marketplace_account"] = o.BillingMarketplaceAccount
	}
	if o.CloudAccountId != nil {
		toSerialize["cloud_account_id"] = o.CloudAccountId
	}
	if o.CloudProviderId != nil {
		toSerialize["cloud_provider_id"] = o.CloudProviderId
	}
	if o.ClusterBillingModel != nil {
		toSerialize["cluster_billing_model"] = o.ClusterBillingModel
	}
	if o.ClusterId != nil {
		toSerialize["cluster_id"] = o.ClusterId
	}
	if o.ConsoleUrl != nil {
		toSerialize["console_url"] = o.ConsoleUrl
	}
	if o.ConsumerUuid != nil {
		toSerialize["consumer_uuid"] = o.ConsumerUuid
	}
	if o.CpuTotal != nil {
		toSerialize["cpu_total"] = o.CpuTotal
	}
	if o.CreatorId != nil {
		toSerialize["creator_id"] = o.CreatorId
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.ExternalClusterId != nil {
		toSerialize["external_cluster_id"] = o.ExternalClusterId
	}
	if o.LastReconcileDate != nil {
		toSerialize["last_reconcile_date"] = o.LastReconcileDate
	}
	if o.LastReleasedAt != nil {
		toSerialize["last_released_at"] = o.LastReleasedAt
	}
	if o.LastTelemetryDate != nil {
		toSerialize["last_telemetry_date"] = o.LastTelemetryDate
	}
	if true {
		toSerialize["managed"] = o.Managed
	}
	if o.OrganizationId != nil {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if o.PlanId != nil {
		toSerialize["plan_id"] = o.PlanId
	}
	if o.ProductBundle != nil {
		toSerialize["product_bundle"] = o.ProductBundle
	}
	if o.Provenance != nil {
		toSerialize["provenance"] = o.Provenance
	}
	if o.RegionId != nil {
		toSerialize["region_id"] = o.RegionId
	}
	if o.Released != nil {
		toSerialize["released"] = o.Released
	}
	if o.ServiceLevel != nil {
		toSerialize["service_level"] = o.ServiceLevel
	}
	if o.SocketTotal != nil {
		toSerialize["socket_total"] = o.SocketTotal
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.SupportLevel != nil {
		toSerialize["support_level"] = o.SupportLevel
	}
	if o.SystemUnits != nil {
		toSerialize["system_units"] = o.SystemUnits
	}
	if o.TrialEndDate != nil {
		toSerialize["trial_end_date"] = o.TrialEndDate
	}
	if o.Usage != nil {
		toSerialize["usage"] = o.Usage
	}
	return json.Marshal(toSerialize)
}

type NullableSubscriptionCommonFields struct {
	value *SubscriptionCommonFields
	isSet bool
}

func (v NullableSubscriptionCommonFields) Get() *SubscriptionCommonFields {
	return v.value
}

func (v *NullableSubscriptionCommonFields) Set(val *SubscriptionCommonFields) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionCommonFields) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionCommonFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionCommonFields(val *SubscriptionCommonFields) *NullableSubscriptionCommonFields {
	return &NullableSubscriptionCommonFields{value: val, isSet: true}
}

func (v NullableSubscriptionCommonFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionCommonFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

