/*
 * Kafka Instance API
 *
 * API for interacting with Kafka Instance. Includes Produce, Consume and Admin APIs
 *
 * API version: 0.12.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkainstanceclient

import (
	"encoding/json"
)

// ErrorList struct for ErrorList
type ErrorList struct {
	Kind *string `json:"kind,omitempty"`
	Items []Error `json:"items"`
	// Total number of errors returned in this request
	Total int32 `json:"total"`
	// Number of entries per page (returned for fetch requests)
	Size *int32 `json:"size,omitempty"`
	// Current page number (returned for fetch requests)
	Page *int32 `json:"page,omitempty"`
}

// NewErrorList instantiates a new ErrorList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorList(items []Error, total int32) *ErrorList {
	this := ErrorList{}
	this.Items = items
	this.Total = total
	return &this
}

// NewErrorListWithDefaults instantiates a new ErrorList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorListWithDefaults() *ErrorList {
	this := ErrorList{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ErrorList) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorList) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ErrorList) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ErrorList) SetKind(v string) {
	o.Kind = &v
}

// GetItems returns the Items field value
func (o *ErrorList) GetItems() []Error {
	if o == nil {
		var ret []Error
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ErrorList) GetItemsOk() (*[]Error, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *ErrorList) SetItems(v []Error) {
	o.Items = v
}

// GetTotal returns the Total field value
func (o *ErrorList) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ErrorList) GetTotalOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *ErrorList) SetTotal(v int32) {
	o.Total = v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ErrorList) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorList) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ErrorList) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *ErrorList) SetSize(v int32) {
	o.Size = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *ErrorList) GetPage() int32 {
	if o == nil || o.Page == nil {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorList) GetPageOk() (*int32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *ErrorList) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *ErrorList) SetPage(v int32) {
	o.Page = &v
}

func (o ErrorList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["items"] = o.Items
	}
	if true {
		toSerialize["total"] = o.Total
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	return json.Marshal(toSerialize)
}

type NullableErrorList struct {
	value *ErrorList
	isSet bool
}

func (v NullableErrorList) Get() *ErrorList {
	return v.value
}

func (v *NullableErrorList) Set(val *ErrorList) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorList) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorList(val *ErrorList) *NullableErrorList {
	return &NullableErrorList{value: val, isSet: true}
}

func (v NullableErrorList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


