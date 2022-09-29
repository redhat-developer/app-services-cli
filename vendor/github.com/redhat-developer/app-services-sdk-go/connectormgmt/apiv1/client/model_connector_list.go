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

// ConnectorList struct for ConnectorList
type ConnectorList struct {
	Kind string `json:"kind"`
	Page int32 `json:"page"`
	Size int32 `json:"size"`
	Total int32 `json:"total"`
	Items []Connector `json:"items"`
}

// NewConnectorList instantiates a new ConnectorList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorList(kind string, page int32, size int32, total int32, items []Connector) *ConnectorList {
	this := ConnectorList{}
	this.Kind = kind
	this.Page = page
	this.Size = size
	this.Total = total
	this.Items = items
	return &this
}

// NewConnectorListWithDefaults instantiates a new ConnectorList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorListWithDefaults() *ConnectorList {
	this := ConnectorList{}
	return &this
}

// GetKind returns the Kind field value
func (o *ConnectorList) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ConnectorList) GetKindOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *ConnectorList) SetKind(v string) {
	o.Kind = v
}

// GetPage returns the Page field value
func (o *ConnectorList) GetPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *ConnectorList) GetPageOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *ConnectorList) SetPage(v int32) {
	o.Page = v
}

// GetSize returns the Size field value
func (o *ConnectorList) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ConnectorList) GetSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ConnectorList) SetSize(v int32) {
	o.Size = v
}

// GetTotal returns the Total field value
func (o *ConnectorList) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ConnectorList) GetTotalOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *ConnectorList) SetTotal(v int32) {
	o.Total = v
}

// GetItems returns the Items field value
func (o *ConnectorList) GetItems() []Connector {
	if o == nil {
		var ret []Connector
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ConnectorList) GetItemsOk() (*[]Connector, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *ConnectorList) SetItems(v []Connector) {
	o.Items = v
}

func (o ConnectorList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["page"] = o.Page
	}
	if true {
		toSerialize["size"] = o.Size
	}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableConnectorList struct {
	value *ConnectorList
	isSet bool
}

func (v NullableConnectorList) Get() *ConnectorList {
	return v.value
}

func (v *NullableConnectorList) Set(val *ConnectorList) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorList) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorList(val *ConnectorList) *NullableConnectorList {
	return &NullableConnectorList{value: val, isSet: true}
}

func (v NullableConnectorList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


