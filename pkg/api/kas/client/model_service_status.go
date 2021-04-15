/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage kafka instances and connectors.
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kasclient

import (
	"encoding/json"
)

// ServiceStatus Schema for the service status response body
type ServiceStatus struct {
	Kafkas *ServiceStatusKafkas `json:"kafkas,omitempty"`
}

// NewServiceStatus instantiates a new ServiceStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceStatus() *ServiceStatus {
	this := ServiceStatus{}
	return &this
}

// NewServiceStatusWithDefaults instantiates a new ServiceStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceStatusWithDefaults() *ServiceStatus {
	this := ServiceStatus{}
	return &this
}

// GetKafkas returns the Kafkas field value if set, zero value otherwise.
func (o *ServiceStatus) GetKafkas() ServiceStatusKafkas {
	if o == nil || o.Kafkas == nil {
		var ret ServiceStatusKafkas
		return ret
	}
	return *o.Kafkas
}

// GetKafkasOk returns a tuple with the Kafkas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceStatus) GetKafkasOk() (*ServiceStatusKafkas, bool) {
	if o == nil || o.Kafkas == nil {
		return nil, false
	}
	return o.Kafkas, true
}

// HasKafkas returns a boolean if a field has been set.
func (o *ServiceStatus) HasKafkas() bool {
	if o != nil && o.Kafkas != nil {
		return true
	}

	return false
}

// SetKafkas gets a reference to the given ServiceStatusKafkas and assigns it to the Kafkas field.
func (o *ServiceStatus) SetKafkas(v ServiceStatusKafkas) {
	o.Kafkas = &v
}

func (o ServiceStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Kafkas != nil {
		toSerialize["kafkas"] = o.Kafkas
	}
	return json.Marshal(toSerialize)
}

type NullableServiceStatus struct {
	value *ServiceStatus
	isSet bool
}

func (v NullableServiceStatus) Get() *ServiceStatus {
	return v.value
}

func (v *NullableServiceStatus) Set(val *ServiceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceStatus(val *ServiceStatus) *NullableServiceStatus {
	return &NullableServiceStatus{value: val, isSet: true}
}

func (v NullableServiceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
