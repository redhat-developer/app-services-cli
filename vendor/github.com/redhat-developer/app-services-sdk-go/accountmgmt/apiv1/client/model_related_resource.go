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
)

// RelatedResource struct for RelatedResource
type RelatedResource struct {
	Href *string `json:"href,omitempty"`
	Id *string `json:"id,omitempty"`
	Kind *string `json:"kind,omitempty"`
	AvailabilityZoneType string `json:"availability_zone_type"`
	BillingModel string `json:"billing_model"`
	Byoc string `json:"byoc"`
	CloudProvider string `json:"cloud_provider"`
	Cost int32 `json:"cost"`
	Product string `json:"product"`
	ProductId *string `json:"product_id,omitempty"`
	ResourceName *string `json:"resource_name,omitempty"`
	ResourceType string `json:"resource_type"`
}

// NewRelatedResource instantiates a new RelatedResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelatedResource(availabilityZoneType string, billingModel string, byoc string, cloudProvider string, cost int32, product string, resourceType string) *RelatedResource {
	this := RelatedResource{}
	this.AvailabilityZoneType = availabilityZoneType
	this.BillingModel = billingModel
	this.Byoc = byoc
	this.CloudProvider = cloudProvider
	this.Cost = cost
	this.Product = product
	this.ResourceType = resourceType
	return &this
}

// NewRelatedResourceWithDefaults instantiates a new RelatedResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelatedResourceWithDefaults() *RelatedResource {
	this := RelatedResource{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *RelatedResource) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedResource) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *RelatedResource) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *RelatedResource) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RelatedResource) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedResource) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RelatedResource) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RelatedResource) SetId(v string) {
	o.Id = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *RelatedResource) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedResource) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *RelatedResource) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *RelatedResource) SetKind(v string) {
	o.Kind = &v
}

// GetAvailabilityZoneType returns the AvailabilityZoneType field value
func (o *RelatedResource) GetAvailabilityZoneType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AvailabilityZoneType
}

// GetAvailabilityZoneTypeOk returns a tuple with the AvailabilityZoneType field value
// and a boolean to check if the value has been set.
func (o *RelatedResource) GetAvailabilityZoneTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AvailabilityZoneType, true
}

// SetAvailabilityZoneType sets field value
func (o *RelatedResource) SetAvailabilityZoneType(v string) {
	o.AvailabilityZoneType = v
}

// GetBillingModel returns the BillingModel field value
func (o *RelatedResource) GetBillingModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BillingModel
}

// GetBillingModelOk returns a tuple with the BillingModel field value
// and a boolean to check if the value has been set.
func (o *RelatedResource) GetBillingModelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BillingModel, true
}

// SetBillingModel sets field value
func (o *RelatedResource) SetBillingModel(v string) {
	o.BillingModel = v
}

// GetByoc returns the Byoc field value
func (o *RelatedResource) GetByoc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Byoc
}

// GetByocOk returns a tuple with the Byoc field value
// and a boolean to check if the value has been set.
func (o *RelatedResource) GetByocOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Byoc, true
}

// SetByoc sets field value
func (o *RelatedResource) SetByoc(v string) {
	o.Byoc = v
}

// GetCloudProvider returns the CloudProvider field value
func (o *RelatedResource) GetCloudProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *RelatedResource) GetCloudProviderOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *RelatedResource) SetCloudProvider(v string) {
	o.CloudProvider = v
}

// GetCost returns the Cost field value
func (o *RelatedResource) GetCost() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Cost
}

// GetCostOk returns a tuple with the Cost field value
// and a boolean to check if the value has been set.
func (o *RelatedResource) GetCostOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cost, true
}

// SetCost sets field value
func (o *RelatedResource) SetCost(v int32) {
	o.Cost = v
}

// GetProduct returns the Product field value
func (o *RelatedResource) GetProduct() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *RelatedResource) GetProductOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value
func (o *RelatedResource) SetProduct(v string) {
	o.Product = v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *RelatedResource) GetProductId() string {
	if o == nil || o.ProductId == nil {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedResource) GetProductIdOk() (*string, bool) {
	if o == nil || o.ProductId == nil {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *RelatedResource) HasProductId() bool {
	if o != nil && o.ProductId != nil {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *RelatedResource) SetProductId(v string) {
	o.ProductId = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *RelatedResource) GetResourceName() string {
	if o == nil || o.ResourceName == nil {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedResource) GetResourceNameOk() (*string, bool) {
	if o == nil || o.ResourceName == nil {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *RelatedResource) HasResourceName() bool {
	if o != nil && o.ResourceName != nil {
		return true
	}

	return false
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *RelatedResource) SetResourceName(v string) {
	o.ResourceName = &v
}

// GetResourceType returns the ResourceType field value
func (o *RelatedResource) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *RelatedResource) GetResourceTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *RelatedResource) SetResourceType(v string) {
	o.ResourceType = v
}

func (o RelatedResource) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["availability_zone_type"] = o.AvailabilityZoneType
	}
	if true {
		toSerialize["billing_model"] = o.BillingModel
	}
	if true {
		toSerialize["byoc"] = o.Byoc
	}
	if true {
		toSerialize["cloud_provider"] = o.CloudProvider
	}
	if true {
		toSerialize["cost"] = o.Cost
	}
	if true {
		toSerialize["product"] = o.Product
	}
	if o.ProductId != nil {
		toSerialize["product_id"] = o.ProductId
	}
	if o.ResourceName != nil {
		toSerialize["resource_name"] = o.ResourceName
	}
	if true {
		toSerialize["resource_type"] = o.ResourceType
	}
	return json.Marshal(toSerialize)
}

type NullableRelatedResource struct {
	value *RelatedResource
	isSet bool
}

func (v NullableRelatedResource) Get() *RelatedResource {
	return v.value
}

func (v *NullableRelatedResource) Set(val *RelatedResource) {
	v.value = val
	v.isSet = true
}

func (v NullableRelatedResource) IsSet() bool {
	return v.isSet
}

func (v *NullableRelatedResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelatedResource(val *RelatedResource) *NullableRelatedResource {
	return &NullableRelatedResource{value: val, isSet: true}
}

func (v NullableRelatedResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelatedResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

