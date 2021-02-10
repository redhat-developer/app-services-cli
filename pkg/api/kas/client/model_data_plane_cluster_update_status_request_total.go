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

// DataPlaneClusterUpdateStatusRequestTotal struct for DataPlaneClusterUpdateStatusRequestTotal
type DataPlaneClusterUpdateStatusRequestTotal struct {
	IngressEgressThroughputPerSec *string `json:"ingressEgressThroughputPerSec,omitempty"`
	Connections                   *int32  `json:"connections,omitempty"`
	DataRetentionSize             *string `json:"dataRetentionSize,omitempty"`
	Partitions                    *int32  `json:"partitions,omitempty"`
}

// NewDataPlaneClusterUpdateStatusRequestTotal instantiates a new DataPlaneClusterUpdateStatusRequestTotal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataPlaneClusterUpdateStatusRequestTotal() *DataPlaneClusterUpdateStatusRequestTotal {
	this := DataPlaneClusterUpdateStatusRequestTotal{}
	return &this
}

// NewDataPlaneClusterUpdateStatusRequestTotalWithDefaults instantiates a new DataPlaneClusterUpdateStatusRequestTotal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataPlaneClusterUpdateStatusRequestTotalWithDefaults() *DataPlaneClusterUpdateStatusRequestTotal {
	this := DataPlaneClusterUpdateStatusRequestTotal{}
	return &this
}

// GetIngressEgressThroughputPerSec returns the IngressEgressThroughputPerSec field value if set, zero value otherwise.
func (o *DataPlaneClusterUpdateStatusRequestTotal) GetIngressEgressThroughputPerSec() string {
	if o == nil || o.IngressEgressThroughputPerSec == nil {
		var ret string
		return ret
	}
	return *o.IngressEgressThroughputPerSec
}

// GetIngressEgressThroughputPerSecOk returns a tuple with the IngressEgressThroughputPerSec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataPlaneClusterUpdateStatusRequestTotal) GetIngressEgressThroughputPerSecOk() (*string, bool) {
	if o == nil || o.IngressEgressThroughputPerSec == nil {
		return nil, false
	}
	return o.IngressEgressThroughputPerSec, true
}

// HasIngressEgressThroughputPerSec returns a boolean if a field has been set.
func (o *DataPlaneClusterUpdateStatusRequestTotal) HasIngressEgressThroughputPerSec() bool {
	if o != nil && o.IngressEgressThroughputPerSec != nil {
		return true
	}

	return false
}

// SetIngressEgressThroughputPerSec gets a reference to the given string and assigns it to the IngressEgressThroughputPerSec field.
func (o *DataPlaneClusterUpdateStatusRequestTotal) SetIngressEgressThroughputPerSec(v string) {
	o.IngressEgressThroughputPerSec = &v
}

// GetConnections returns the Connections field value if set, zero value otherwise.
func (o *DataPlaneClusterUpdateStatusRequestTotal) GetConnections() int32 {
	if o == nil || o.Connections == nil {
		var ret int32
		return ret
	}
	return *o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataPlaneClusterUpdateStatusRequestTotal) GetConnectionsOk() (*int32, bool) {
	if o == nil || o.Connections == nil {
		return nil, false
	}
	return o.Connections, true
}

// HasConnections returns a boolean if a field has been set.
func (o *DataPlaneClusterUpdateStatusRequestTotal) HasConnections() bool {
	if o != nil && o.Connections != nil {
		return true
	}

	return false
}

// SetConnections gets a reference to the given int32 and assigns it to the Connections field.
func (o *DataPlaneClusterUpdateStatusRequestTotal) SetConnections(v int32) {
	o.Connections = &v
}

// GetDataRetentionSize returns the DataRetentionSize field value if set, zero value otherwise.
func (o *DataPlaneClusterUpdateStatusRequestTotal) GetDataRetentionSize() string {
	if o == nil || o.DataRetentionSize == nil {
		var ret string
		return ret
	}
	return *o.DataRetentionSize
}

// GetDataRetentionSizeOk returns a tuple with the DataRetentionSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataPlaneClusterUpdateStatusRequestTotal) GetDataRetentionSizeOk() (*string, bool) {
	if o == nil || o.DataRetentionSize == nil {
		return nil, false
	}
	return o.DataRetentionSize, true
}

// HasDataRetentionSize returns a boolean if a field has been set.
func (o *DataPlaneClusterUpdateStatusRequestTotal) HasDataRetentionSize() bool {
	if o != nil && o.DataRetentionSize != nil {
		return true
	}

	return false
}

// SetDataRetentionSize gets a reference to the given string and assigns it to the DataRetentionSize field.
func (o *DataPlaneClusterUpdateStatusRequestTotal) SetDataRetentionSize(v string) {
	o.DataRetentionSize = &v
}

// GetPartitions returns the Partitions field value if set, zero value otherwise.
func (o *DataPlaneClusterUpdateStatusRequestTotal) GetPartitions() int32 {
	if o == nil || o.Partitions == nil {
		var ret int32
		return ret
	}
	return *o.Partitions
}

// GetPartitionsOk returns a tuple with the Partitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataPlaneClusterUpdateStatusRequestTotal) GetPartitionsOk() (*int32, bool) {
	if o == nil || o.Partitions == nil {
		return nil, false
	}
	return o.Partitions, true
}

// HasPartitions returns a boolean if a field has been set.
func (o *DataPlaneClusterUpdateStatusRequestTotal) HasPartitions() bool {
	if o != nil && o.Partitions != nil {
		return true
	}

	return false
}

// SetPartitions gets a reference to the given int32 and assigns it to the Partitions field.
func (o *DataPlaneClusterUpdateStatusRequestTotal) SetPartitions(v int32) {
	o.Partitions = &v
}

func (o DataPlaneClusterUpdateStatusRequestTotal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IngressEgressThroughputPerSec != nil {
		toSerialize["ingressEgressThroughputPerSec"] = o.IngressEgressThroughputPerSec
	}
	if o.Connections != nil {
		toSerialize["connections"] = o.Connections
	}
	if o.DataRetentionSize != nil {
		toSerialize["dataRetentionSize"] = o.DataRetentionSize
	}
	if o.Partitions != nil {
		toSerialize["partitions"] = o.Partitions
	}
	return json.Marshal(toSerialize)
}

type NullableDataPlaneClusterUpdateStatusRequestTotal struct {
	value *DataPlaneClusterUpdateStatusRequestTotal
	isSet bool
}

func (v NullableDataPlaneClusterUpdateStatusRequestTotal) Get() *DataPlaneClusterUpdateStatusRequestTotal {
	return v.value
}

func (v *NullableDataPlaneClusterUpdateStatusRequestTotal) Set(val *DataPlaneClusterUpdateStatusRequestTotal) {
	v.value = val
	v.isSet = true
}

func (v NullableDataPlaneClusterUpdateStatusRequestTotal) IsSet() bool {
	return v.isSet
}

func (v *NullableDataPlaneClusterUpdateStatusRequestTotal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataPlaneClusterUpdateStatusRequestTotal(val *DataPlaneClusterUpdateStatusRequestTotal) *NullableDataPlaneClusterUpdateStatusRequestTotal {
	return &NullableDataPlaneClusterUpdateStatusRequestTotal{value: val, isSet: true}
}

func (v NullableDataPlaneClusterUpdateStatusRequestTotal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataPlaneClusterUpdateStatusRequestTotal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
