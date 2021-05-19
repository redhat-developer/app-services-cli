/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage kafka instances and connectors.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkamgmtv1

import (
	"encoding/json"
)

// InstantQuery struct for InstantQuery
type InstantQuery struct {
	Metric    *map[string]string `json:"metric,omitempty"`
	Timestamp *int64             `json:"Timestamp,omitempty"`
	Value     float64            `json:"Value"`
}

// NewInstantQuery instantiates a new InstantQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstantQuery(value float64) *InstantQuery {
	this := InstantQuery{}
	this.Value = value
	return &this
}

// NewInstantQueryWithDefaults instantiates a new InstantQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstantQueryWithDefaults() *InstantQuery {
	this := InstantQuery{}
	return &this
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *InstantQuery) GetMetric() map[string]string {
	if o == nil || o.Metric == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstantQuery) GetMetricOk() (*map[string]string, bool) {
	if o == nil || o.Metric == nil {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *InstantQuery) HasMetric() bool {
	if o != nil && o.Metric != nil {
		return true
	}

	return false
}

// SetMetric gets a reference to the given map[string]string and assigns it to the Metric field.
func (o *InstantQuery) SetMetric(v map[string]string) {
	o.Metric = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *InstantQuery) GetTimestamp() int64 {
	if o == nil || o.Timestamp == nil {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstantQuery) GetTimestampOk() (*int64, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *InstantQuery) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *InstantQuery) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetValue returns the Value field value
func (o *InstantQuery) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *InstantQuery) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *InstantQuery) SetValue(v float64) {
	o.Value = v
}

func (o InstantQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Metric != nil {
		toSerialize["metric"] = o.Metric
	}
	if o.Timestamp != nil {
		toSerialize["Timestamp"] = o.Timestamp
	}
	if true {
		toSerialize["Value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableInstantQuery struct {
	value *InstantQuery
	isSet bool
}

func (v NullableInstantQuery) Get() *InstantQuery {
	return v.value
}

func (v *NullableInstantQuery) Set(val *InstantQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableInstantQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableInstantQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstantQuery(val *InstantQuery) *NullableInstantQuery {
	return &NullableInstantQuery{value: val, isSet: true}
}

func (v NullableInstantQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstantQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
