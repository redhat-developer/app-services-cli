/*
 * Managed Service API
 *
 * Managed Service API
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kasclient

import (
	"encoding/json"
)

// DataPlaneClusterUpdateStatusRequestConditions struct for DataPlaneClusterUpdateStatusRequestConditions
type DataPlaneClusterUpdateStatusRequestConditions struct {
	Type               *string `json:"type,omitempty"`
	Reason             *string `json:"reason,omitempty"`
	Message            *string `json:"message,omitempty"`
	Status             *string `json:"status,omitempty"`
	LastTransitionTime *string `json:"lastTransitionTime,omitempty"`
}

// NewDataPlaneClusterUpdateStatusRequestConditions instantiates a new DataPlaneClusterUpdateStatusRequestConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataPlaneClusterUpdateStatusRequestConditions() *DataPlaneClusterUpdateStatusRequestConditions {
	this := DataPlaneClusterUpdateStatusRequestConditions{}
	return &this
}

// NewDataPlaneClusterUpdateStatusRequestConditionsWithDefaults instantiates a new DataPlaneClusterUpdateStatusRequestConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataPlaneClusterUpdateStatusRequestConditionsWithDefaults() *DataPlaneClusterUpdateStatusRequestConditions {
	this := DataPlaneClusterUpdateStatusRequestConditions{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DataPlaneClusterUpdateStatusRequestConditions) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataPlaneClusterUpdateStatusRequestConditions) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DataPlaneClusterUpdateStatusRequestConditions) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DataPlaneClusterUpdateStatusRequestConditions) SetType(v string) {
	o.Type = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *DataPlaneClusterUpdateStatusRequestConditions) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataPlaneClusterUpdateStatusRequestConditions) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *DataPlaneClusterUpdateStatusRequestConditions) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *DataPlaneClusterUpdateStatusRequestConditions) SetReason(v string) {
	o.Reason = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DataPlaneClusterUpdateStatusRequestConditions) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataPlaneClusterUpdateStatusRequestConditions) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DataPlaneClusterUpdateStatusRequestConditions) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *DataPlaneClusterUpdateStatusRequestConditions) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DataPlaneClusterUpdateStatusRequestConditions) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataPlaneClusterUpdateStatusRequestConditions) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DataPlaneClusterUpdateStatusRequestConditions) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DataPlaneClusterUpdateStatusRequestConditions) SetStatus(v string) {
	o.Status = &v
}

// GetLastTransitionTime returns the LastTransitionTime field value if set, zero value otherwise.
func (o *DataPlaneClusterUpdateStatusRequestConditions) GetLastTransitionTime() string {
	if o == nil || o.LastTransitionTime == nil {
		var ret string
		return ret
	}
	return *o.LastTransitionTime
}

// GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataPlaneClusterUpdateStatusRequestConditions) GetLastTransitionTimeOk() (*string, bool) {
	if o == nil || o.LastTransitionTime == nil {
		return nil, false
	}
	return o.LastTransitionTime, true
}

// HasLastTransitionTime returns a boolean if a field has been set.
func (o *DataPlaneClusterUpdateStatusRequestConditions) HasLastTransitionTime() bool {
	if o != nil && o.LastTransitionTime != nil {
		return true
	}

	return false
}

// SetLastTransitionTime gets a reference to the given string and assigns it to the LastTransitionTime field.
func (o *DataPlaneClusterUpdateStatusRequestConditions) SetLastTransitionTime(v string) {
	o.LastTransitionTime = &v
}

func (o DataPlaneClusterUpdateStatusRequestConditions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.LastTransitionTime != nil {
		toSerialize["lastTransitionTime"] = o.LastTransitionTime
	}
	return json.Marshal(toSerialize)
}

type NullableDataPlaneClusterUpdateStatusRequestConditions struct {
	value *DataPlaneClusterUpdateStatusRequestConditions
	isSet bool
}

func (v NullableDataPlaneClusterUpdateStatusRequestConditions) Get() *DataPlaneClusterUpdateStatusRequestConditions {
	return v.value
}

func (v *NullableDataPlaneClusterUpdateStatusRequestConditions) Set(val *DataPlaneClusterUpdateStatusRequestConditions) {
	v.value = val
	v.isSet = true
}

func (v NullableDataPlaneClusterUpdateStatusRequestConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableDataPlaneClusterUpdateStatusRequestConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataPlaneClusterUpdateStatusRequestConditions(val *DataPlaneClusterUpdateStatusRequestConditions) *NullableDataPlaneClusterUpdateStatusRequestConditions {
	return &NullableDataPlaneClusterUpdateStatusRequestConditions{value: val, isSet: true}
}

func (v NullableDataPlaneClusterUpdateStatusRequestConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataPlaneClusterUpdateStatusRequestConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
