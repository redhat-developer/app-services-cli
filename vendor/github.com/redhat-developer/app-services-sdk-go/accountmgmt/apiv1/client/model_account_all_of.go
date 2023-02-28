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

// AccountAllOf struct for AccountAllOf
type AccountAllOf struct {
	BanCode *string `json:"ban_code,omitempty"`
	BanDescription *string `json:"ban_description,omitempty"`
	Banned *bool `json:"banned,omitempty"`
	Capabilities *[]Capability `json:"capabilities,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Email *string `json:"email,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	Labels *[]Label `json:"labels,omitempty"`
	LastName *string `json:"last_name,omitempty"`
	Organization *Organization `json:"organization,omitempty"`
	OrganizationId *string `json:"organization_id,omitempty"`
	RhitAccountId *string `json:"rhit_account_id,omitempty"`
	RhitWebUserId *string `json:"rhit_web_user_id,omitempty"`
	ServiceAccount *bool `json:"service_account,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Username string `json:"username"`
}

// NewAccountAllOf instantiates a new AccountAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountAllOf(username string) *AccountAllOf {
	this := AccountAllOf{}
	var banned bool = false
	this.Banned = &banned
	var serviceAccount bool = false
	this.ServiceAccount = &serviceAccount
	this.Username = username
	return &this
}

// NewAccountAllOfWithDefaults instantiates a new AccountAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountAllOfWithDefaults() *AccountAllOf {
	this := AccountAllOf{}
	var banned bool = false
	this.Banned = &banned
	var serviceAccount bool = false
	this.ServiceAccount = &serviceAccount
	return &this
}

// GetBanCode returns the BanCode field value if set, zero value otherwise.
func (o *AccountAllOf) GetBanCode() string {
	if o == nil || o.BanCode == nil {
		var ret string
		return ret
	}
	return *o.BanCode
}

// GetBanCodeOk returns a tuple with the BanCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAllOf) GetBanCodeOk() (*string, bool) {
	if o == nil || o.BanCode == nil {
		return nil, false
	}
	return o.BanCode, true
}

// HasBanCode returns a boolean if a field has been set.
func (o *AccountAllOf) HasBanCode() bool {
	if o != nil && o.BanCode != nil {
		return true
	}

	return false
}

// SetBanCode gets a reference to the given string and assigns it to the BanCode field.
func (o *AccountAllOf) SetBanCode(v string) {
	o.BanCode = &v
}

// GetBanDescription returns the BanDescription field value if set, zero value otherwise.
func (o *AccountAllOf) GetBanDescription() string {
	if o == nil || o.BanDescription == nil {
		var ret string
		return ret
	}
	return *o.BanDescription
}

// GetBanDescriptionOk returns a tuple with the BanDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAllOf) GetBanDescriptionOk() (*string, bool) {
	if o == nil || o.BanDescription == nil {
		return nil, false
	}
	return o.BanDescription, true
}

// HasBanDescription returns a boolean if a field has been set.
func (o *AccountAllOf) HasBanDescription() bool {
	if o != nil && o.BanDescription != nil {
		return true
	}

	return false
}

// SetBanDescription gets a reference to the given string and assigns it to the BanDescription field.
func (o *AccountAllOf) SetBanDescription(v string) {
	o.BanDescription = &v
}

// GetBanned returns the Banned field value if set, zero value otherwise.
func (o *AccountAllOf) GetBanned() bool {
	if o == nil || o.Banned == nil {
		var ret bool
		return ret
	}
	return *o.Banned
}

// GetBannedOk returns a tuple with the Banned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAllOf) GetBannedOk() (*bool, bool) {
	if o == nil || o.Banned == nil {
		return nil, false
	}
	return o.Banned, true
}

// HasBanned returns a boolean if a field has been set.
func (o *AccountAllOf) HasBanned() bool {
	if o != nil && o.Banned != nil {
		return true
	}

	return false
}

// SetBanned gets a reference to the given bool and assigns it to the Banned field.
func (o *AccountAllOf) SetBanned(v bool) {
	o.Banned = &v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *AccountAllOf) GetCapabilities() []Capability {
	if o == nil || o.Capabilities == nil {
		var ret []Capability
		return ret
	}
	return *o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAllOf) GetCapabilitiesOk() (*[]Capability, bool) {
	if o == nil || o.Capabilities == nil {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *AccountAllOf) HasCapabilities() bool {
	if o != nil && o.Capabilities != nil {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given []Capability and assigns it to the Capabilities field.
func (o *AccountAllOf) SetCapabilities(v []Capability) {
	o.Capabilities = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AccountAllOf) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAllOf) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AccountAllOf) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AccountAllOf) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *AccountAllOf) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAllOf) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *AccountAllOf) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *AccountAllOf) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *AccountAllOf) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAllOf) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *AccountAllOf) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *AccountAllOf) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *AccountAllOf) GetLabels() []Label {
	if o == nil || o.Labels == nil {
		var ret []Label
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAllOf) GetLabelsOk() (*[]Label, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *AccountAllOf) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []Label and assigns it to the Labels field.
func (o *AccountAllOf) SetLabels(v []Label) {
	o.Labels = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *AccountAllOf) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAllOf) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *AccountAllOf) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *AccountAllOf) SetLastName(v string) {
	o.LastName = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *AccountAllOf) GetOrganization() Organization {
	if o == nil || o.Organization == nil {
		var ret Organization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAllOf) GetOrganizationOk() (*Organization, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *AccountAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given Organization and assigns it to the Organization field.
func (o *AccountAllOf) SetOrganization(v Organization) {
	o.Organization = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *AccountAllOf) GetOrganizationId() string {
	if o == nil || o.OrganizationId == nil {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAllOf) GetOrganizationIdOk() (*string, bool) {
	if o == nil || o.OrganizationId == nil {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *AccountAllOf) HasOrganizationId() bool {
	if o != nil && o.OrganizationId != nil {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *AccountAllOf) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetRhitAccountId returns the RhitAccountId field value if set, zero value otherwise.
func (o *AccountAllOf) GetRhitAccountId() string {
	if o == nil || o.RhitAccountId == nil {
		var ret string
		return ret
	}
	return *o.RhitAccountId
}

// GetRhitAccountIdOk returns a tuple with the RhitAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAllOf) GetRhitAccountIdOk() (*string, bool) {
	if o == nil || o.RhitAccountId == nil {
		return nil, false
	}
	return o.RhitAccountId, true
}

// HasRhitAccountId returns a boolean if a field has been set.
func (o *AccountAllOf) HasRhitAccountId() bool {
	if o != nil && o.RhitAccountId != nil {
		return true
	}

	return false
}

// SetRhitAccountId gets a reference to the given string and assigns it to the RhitAccountId field.
func (o *AccountAllOf) SetRhitAccountId(v string) {
	o.RhitAccountId = &v
}

// GetRhitWebUserId returns the RhitWebUserId field value if set, zero value otherwise.
func (o *AccountAllOf) GetRhitWebUserId() string {
	if o == nil || o.RhitWebUserId == nil {
		var ret string
		return ret
	}
	return *o.RhitWebUserId
}

// GetRhitWebUserIdOk returns a tuple with the RhitWebUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAllOf) GetRhitWebUserIdOk() (*string, bool) {
	if o == nil || o.RhitWebUserId == nil {
		return nil, false
	}
	return o.RhitWebUserId, true
}

// HasRhitWebUserId returns a boolean if a field has been set.
func (o *AccountAllOf) HasRhitWebUserId() bool {
	if o != nil && o.RhitWebUserId != nil {
		return true
	}

	return false
}

// SetRhitWebUserId gets a reference to the given string and assigns it to the RhitWebUserId field.
func (o *AccountAllOf) SetRhitWebUserId(v string) {
	o.RhitWebUserId = &v
}

// GetServiceAccount returns the ServiceAccount field value if set, zero value otherwise.
func (o *AccountAllOf) GetServiceAccount() bool {
	if o == nil || o.ServiceAccount == nil {
		var ret bool
		return ret
	}
	return *o.ServiceAccount
}

// GetServiceAccountOk returns a tuple with the ServiceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAllOf) GetServiceAccountOk() (*bool, bool) {
	if o == nil || o.ServiceAccount == nil {
		return nil, false
	}
	return o.ServiceAccount, true
}

// HasServiceAccount returns a boolean if a field has been set.
func (o *AccountAllOf) HasServiceAccount() bool {
	if o != nil && o.ServiceAccount != nil {
		return true
	}

	return false
}

// SetServiceAccount gets a reference to the given bool and assigns it to the ServiceAccount field.
func (o *AccountAllOf) SetServiceAccount(v bool) {
	o.ServiceAccount = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AccountAllOf) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAllOf) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AccountAllOf) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *AccountAllOf) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUsername returns the Username field value
func (o *AccountAllOf) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *AccountAllOf) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *AccountAllOf) SetUsername(v string) {
	o.Username = v
}

func (o AccountAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BanCode != nil {
		toSerialize["ban_code"] = o.BanCode
	}
	if o.BanDescription != nil {
		toSerialize["ban_description"] = o.BanDescription
	}
	if o.Banned != nil {
		toSerialize["banned"] = o.Banned
	}
	if o.Capabilities != nil {
		toSerialize["capabilities"] = o.Capabilities
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.FirstName != nil {
		toSerialize["first_name"] = o.FirstName
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.LastName != nil {
		toSerialize["last_name"] = o.LastName
	}
	if o.Organization != nil {
		toSerialize["organization"] = o.Organization
	}
	if o.OrganizationId != nil {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if o.RhitAccountId != nil {
		toSerialize["rhit_account_id"] = o.RhitAccountId
	}
	if o.RhitWebUserId != nil {
		toSerialize["rhit_web_user_id"] = o.RhitWebUserId
	}
	if o.ServiceAccount != nil {
		toSerialize["service_account"] = o.ServiceAccount
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if true {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableAccountAllOf struct {
	value *AccountAllOf
	isSet bool
}

func (v NullableAccountAllOf) Get() *AccountAllOf {
	return v.value
}

func (v *NullableAccountAllOf) Set(val *AccountAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountAllOf(val *AccountAllOf) *NullableAccountAllOf {
	return &NullableAccountAllOf{value: val, isSet: true}
}

func (v NullableAccountAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

