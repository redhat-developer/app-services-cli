/*
 * Strimzi Kubernetes REST API
 *
 * An API to provide k8s REST endpoints for query
 *
 * API version: 0.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package strimziadminclient

import (
	"encoding/json"
)

// Partition Kafka topic partition
type Partition struct {
	// Uniquie id for the partition
	Id int32 `json:"id"`
	// List of replicas for the partition
	Replicas *[]map[string]interface{} `json:"replicas,omitempty"`
	// List isync-replicas for this partition.
	Isr *[]map[string]interface{} `json:"isr,omitempty"`
	// Kafka server / broker.
	Leader *map[string]interface{} `json:"leader,omitempty"`
}

// NewPartition instantiates a new Partition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartition(id int32) *Partition {
	this := Partition{}
	this.Id = id
	return &this
}

// NewPartitionWithDefaults instantiates a new Partition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartitionWithDefaults() *Partition {
	this := Partition{}
	return &this
}

// GetId returns the Id field value
func (o *Partition) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Partition) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Partition) SetId(v int32) {
	o.Id = v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *Partition) GetReplicas() []map[string]interface{} {
	if o == nil || o.Replicas == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Partition) GetReplicasOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Replicas == nil {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *Partition) HasReplicas() bool {
	if o != nil && o.Replicas != nil {
		return true
	}

	return false
}

// SetReplicas gets a reference to the given []map[string]interface{} and assigns it to the Replicas field.
func (o *Partition) SetReplicas(v []map[string]interface{}) {
	o.Replicas = &v
}

// GetIsr returns the Isr field value if set, zero value otherwise.
func (o *Partition) GetIsr() []map[string]interface{} {
	if o == nil || o.Isr == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Isr
}

// GetIsrOk returns a tuple with the Isr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Partition) GetIsrOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Isr == nil {
		return nil, false
	}
	return o.Isr, true
}

// HasIsr returns a boolean if a field has been set.
func (o *Partition) HasIsr() bool {
	if o != nil && o.Isr != nil {
		return true
	}

	return false
}

// SetIsr gets a reference to the given []map[string]interface{} and assigns it to the Isr field.
func (o *Partition) SetIsr(v []map[string]interface{}) {
	o.Isr = &v
}

// GetLeader returns the Leader field value if set, zero value otherwise.
func (o *Partition) GetLeader() map[string]interface{} {
	if o == nil || o.Leader == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Leader
}

// GetLeaderOk returns a tuple with the Leader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Partition) GetLeaderOk() (*map[string]interface{}, bool) {
	if o == nil || o.Leader == nil {
		return nil, false
	}
	return o.Leader, true
}

// HasLeader returns a boolean if a field has been set.
func (o *Partition) HasLeader() bool {
	if o != nil && o.Leader != nil {
		return true
	}

	return false
}

// SetLeader gets a reference to the given map[string]interface{} and assigns it to the Leader field.
func (o *Partition) SetLeader(v map[string]interface{}) {
	o.Leader = &v
}

func (o Partition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Replicas != nil {
		toSerialize["replicas"] = o.Replicas
	}
	if o.Isr != nil {
		toSerialize["isr"] = o.Isr
	}
	if o.Leader != nil {
		toSerialize["leader"] = o.Leader
	}
	return json.Marshal(toSerialize)
}

type NullablePartition struct {
	value *Partition
	isSet bool
}

func (v NullablePartition) Get() *Partition {
	return v.value
}

func (v *NullablePartition) Set(val *Partition) {
	v.value = val
	v.isSet = true
}

func (v NullablePartition) IsSet() bool {
	return v.isSet
}

func (v *NullablePartition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartition(val *Partition) *NullablePartition {
	return &NullablePartition{value: val, isSet: true}
}

func (v NullablePartition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
