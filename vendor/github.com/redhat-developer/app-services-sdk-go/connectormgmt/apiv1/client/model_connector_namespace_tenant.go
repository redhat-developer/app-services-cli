/*
 * Connector Management API
 *
 * Connector Management API is a REST API to manage connectors.
 *
 * API version: 0.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package connectormgmtclient

import (
	"encoding/json"
)

// ConnectorNamespaceTenant struct for ConnectorNamespaceTenant
type ConnectorNamespaceTenant struct {
	Kind ConnectorNamespaceTenantKind `json:"kind"`
	// Either user or organisation id depending on the value of kind
	Id string `json:"id"`
}

// NewConnectorNamespaceTenant instantiates a new ConnectorNamespaceTenant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorNamespaceTenant(kind ConnectorNamespaceTenantKind, id string) *ConnectorNamespaceTenant {
	this := ConnectorNamespaceTenant{}
	this.Kind = kind
	this.Id = id
	return &this
}

// NewConnectorNamespaceTenantWithDefaults instantiates a new ConnectorNamespaceTenant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorNamespaceTenantWithDefaults() *ConnectorNamespaceTenant {
	this := ConnectorNamespaceTenant{}
	return &this
}

// GetKind returns the Kind field value
func (o *ConnectorNamespaceTenant) GetKind() ConnectorNamespaceTenantKind {
	if o == nil {
		var ret ConnectorNamespaceTenantKind
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ConnectorNamespaceTenant) GetKindOk() (*ConnectorNamespaceTenantKind, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *ConnectorNamespaceTenant) SetKind(v ConnectorNamespaceTenantKind) {
	o.Kind = v
}

// GetId returns the Id field value
func (o *ConnectorNamespaceTenant) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConnectorNamespaceTenant) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConnectorNamespaceTenant) SetId(v string) {
	o.Id = v
}

func (o ConnectorNamespaceTenant) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableConnectorNamespaceTenant struct {
	value *ConnectorNamespaceTenant
	isSet bool
}

func (v NullableConnectorNamespaceTenant) Get() *ConnectorNamespaceTenant {
	return v.value
}

func (v *NullableConnectorNamespaceTenant) Set(val *ConnectorNamespaceTenant) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorNamespaceTenant) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorNamespaceTenant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorNamespaceTenant(val *ConnectorNamespaceTenant) *NullableConnectorNamespaceTenant {
	return &NullableConnectorNamespaceTenant{value: val, isSet: true}
}

func (v NullableConnectorNamespaceTenant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorNamespaceTenant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


