/*
 * Kafka Management API
 *
 * Kafka Management API is a REST API to manage Kafka instances
 *
 * API version: 1.15.0
 * Contact: rhosak-support@redhat.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkamgmtclient

import (
	"encoding/json"
)

// EnterpriseClusterListItem struct for EnterpriseClusterListItem
type EnterpriseClusterListItem struct {
	Id string `json:"id"`
	Kind string `json:"kind"`
	Href string `json:"href"`
	// Indicates whether Kafkas created on this data plane cluster have to be accessed via private network
	AccessKafkasViaPrivateNetwork bool `json:"access_kafkas_via_private_network"`
	// The OCM's cluster id of the registered Enterprise cluster.
	ClusterId *string `json:"cluster_id,omitempty"`
	// The status of Enterprise cluster registration
	Status *string `json:"status,omitempty"`
	// The cloud provider for this cluster. This valus will be used as the Kafka's cloud provider value when a Kafka is created on this cluster
	CloudProvider *string `json:"cloud_provider,omitempty"`
	// The region of this cluster. This valus will be used as the Kafka's region value when a Kafka is created on this cluster
	Region *string `json:"region,omitempty"`
	// A flag indicating whether this cluster is available on multiple availability zones or not
	MultiAz bool `json:"multi_az"`
}

// NewEnterpriseClusterListItem instantiates a new EnterpriseClusterListItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnterpriseClusterListItem(id string, kind string, href string, accessKafkasViaPrivateNetwork bool, multiAz bool) *EnterpriseClusterListItem {
	this := EnterpriseClusterListItem{}
	this.Id = id
	this.Kind = kind
	this.Href = href
	this.AccessKafkasViaPrivateNetwork = accessKafkasViaPrivateNetwork
	this.MultiAz = multiAz
	return &this
}

// NewEnterpriseClusterListItemWithDefaults instantiates a new EnterpriseClusterListItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnterpriseClusterListItemWithDefaults() *EnterpriseClusterListItem {
	this := EnterpriseClusterListItem{}
	return &this
}

// GetId returns the Id field value
func (o *EnterpriseClusterListItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EnterpriseClusterListItem) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EnterpriseClusterListItem) SetId(v string) {
	o.Id = v
}

// GetKind returns the Kind field value
func (o *EnterpriseClusterListItem) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *EnterpriseClusterListItem) GetKindOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *EnterpriseClusterListItem) SetKind(v string) {
	o.Kind = v
}

// GetHref returns the Href field value
func (o *EnterpriseClusterListItem) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *EnterpriseClusterListItem) GetHrefOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *EnterpriseClusterListItem) SetHref(v string) {
	o.Href = v
}

// GetAccessKafkasViaPrivateNetwork returns the AccessKafkasViaPrivateNetwork field value
func (o *EnterpriseClusterListItem) GetAccessKafkasViaPrivateNetwork() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AccessKafkasViaPrivateNetwork
}

// GetAccessKafkasViaPrivateNetworkOk returns a tuple with the AccessKafkasViaPrivateNetwork field value
// and a boolean to check if the value has been set.
func (o *EnterpriseClusterListItem) GetAccessKafkasViaPrivateNetworkOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessKafkasViaPrivateNetwork, true
}

// SetAccessKafkasViaPrivateNetwork sets field value
func (o *EnterpriseClusterListItem) SetAccessKafkasViaPrivateNetwork(v bool) {
	o.AccessKafkasViaPrivateNetwork = v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *EnterpriseClusterListItem) GetClusterId() string {
	if o == nil || o.ClusterId == nil {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseClusterListItem) GetClusterIdOk() (*string, bool) {
	if o == nil || o.ClusterId == nil {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *EnterpriseClusterListItem) HasClusterId() bool {
	if o != nil && o.ClusterId != nil {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *EnterpriseClusterListItem) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EnterpriseClusterListItem) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseClusterListItem) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EnterpriseClusterListItem) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *EnterpriseClusterListItem) SetStatus(v string) {
	o.Status = &v
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise.
func (o *EnterpriseClusterListItem) GetCloudProvider() string {
	if o == nil || o.CloudProvider == nil {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseClusterListItem) GetCloudProviderOk() (*string, bool) {
	if o == nil || o.CloudProvider == nil {
		return nil, false
	}
	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *EnterpriseClusterListItem) HasCloudProvider() bool {
	if o != nil && o.CloudProvider != nil {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *EnterpriseClusterListItem) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *EnterpriseClusterListItem) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseClusterListItem) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *EnterpriseClusterListItem) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *EnterpriseClusterListItem) SetRegion(v string) {
	o.Region = &v
}

// GetMultiAz returns the MultiAz field value
func (o *EnterpriseClusterListItem) GetMultiAz() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MultiAz
}

// GetMultiAzOk returns a tuple with the MultiAz field value
// and a boolean to check if the value has been set.
func (o *EnterpriseClusterListItem) GetMultiAzOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MultiAz, true
}

// SetMultiAz sets field value
func (o *EnterpriseClusterListItem) SetMultiAz(v bool) {
	o.MultiAz = v
}

func (o EnterpriseClusterListItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["href"] = o.Href
	}
	if true {
		toSerialize["access_kafkas_via_private_network"] = o.AccessKafkasViaPrivateNetwork
	}
	if o.ClusterId != nil {
		toSerialize["cluster_id"] = o.ClusterId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.CloudProvider != nil {
		toSerialize["cloud_provider"] = o.CloudProvider
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if true {
		toSerialize["multi_az"] = o.MultiAz
	}
	return json.Marshal(toSerialize)
}

type NullableEnterpriseClusterListItem struct {
	value *EnterpriseClusterListItem
	isSet bool
}

func (v NullableEnterpriseClusterListItem) Get() *EnterpriseClusterListItem {
	return v.value
}

func (v *NullableEnterpriseClusterListItem) Set(val *EnterpriseClusterListItem) {
	v.value = val
	v.isSet = true
}

func (v NullableEnterpriseClusterListItem) IsSet() bool {
	return v.isSet
}

func (v *NullableEnterpriseClusterListItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnterpriseClusterListItem(val *EnterpriseClusterListItem) *NullableEnterpriseClusterListItem {
	return &NullableEnterpriseClusterListItem{value: val, isSet: true}
}

func (v NullableEnterpriseClusterListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnterpriseClusterListItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

