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
	"time"
)

// Record struct for Record
type Record struct {
	// Unique identifier for the object. Not supported for all object kinds.
	Id *string `json:"id,omitempty"`
	Kind *string `json:"kind,omitempty"`
	// Link path to request the object. Not supported for all object kinds.
	Href *string `json:"href,omitempty"`
	// The record's partition within the topic
	Partition *int32 `json:"partition,omitempty"`
	// The record's offset within the topic partition
	Offset *int64 `json:"offset,omitempty"`
	// Timestamp associated with the record. The type is indicated by `timestampType`. When producing a record, this value will be used as the record's `CREATE_TIME`.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// Type of timestamp associated with the record
	TimestampType *string `json:"timestampType,omitempty"`
	// Record headers, key/value pairs
	Headers *map[string]string `json:"headers,omitempty"`
	// Record key
	Key *string `json:"key,omitempty"`
	// Record value
	Value string `json:"value"`
}

// NewRecord instantiates a new Record object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecord(value string) *Record {
	this := Record{}
	this.Value = value
	return &this
}

// NewRecordWithDefaults instantiates a new Record object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordWithDefaults() *Record {
	this := Record{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Record) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Record) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Record) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Record) SetId(v string) {
	o.Id = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *Record) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Record) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *Record) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *Record) SetKind(v string) {
	o.Kind = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *Record) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Record) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *Record) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *Record) SetHref(v string) {
	o.Href = &v
}

// GetPartition returns the Partition field value if set, zero value otherwise.
func (o *Record) GetPartition() int32 {
	if o == nil || o.Partition == nil {
		var ret int32
		return ret
	}
	return *o.Partition
}

// GetPartitionOk returns a tuple with the Partition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Record) GetPartitionOk() (*int32, bool) {
	if o == nil || o.Partition == nil {
		return nil, false
	}
	return o.Partition, true
}

// HasPartition returns a boolean if a field has been set.
func (o *Record) HasPartition() bool {
	if o != nil && o.Partition != nil {
		return true
	}

	return false
}

// SetPartition gets a reference to the given int32 and assigns it to the Partition field.
func (o *Record) SetPartition(v int32) {
	o.Partition = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *Record) GetOffset() int64 {
	if o == nil || o.Offset == nil {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Record) GetOffsetOk() (*int64, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *Record) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *Record) SetOffset(v int64) {
	o.Offset = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *Record) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Record) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *Record) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *Record) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetTimestampType returns the TimestampType field value if set, zero value otherwise.
func (o *Record) GetTimestampType() string {
	if o == nil || o.TimestampType == nil {
		var ret string
		return ret
	}
	return *o.TimestampType
}

// GetTimestampTypeOk returns a tuple with the TimestampType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Record) GetTimestampTypeOk() (*string, bool) {
	if o == nil || o.TimestampType == nil {
		return nil, false
	}
	return o.TimestampType, true
}

// HasTimestampType returns a boolean if a field has been set.
func (o *Record) HasTimestampType() bool {
	if o != nil && o.TimestampType != nil {
		return true
	}

	return false
}

// SetTimestampType gets a reference to the given string and assigns it to the TimestampType field.
func (o *Record) SetTimestampType(v string) {
	o.TimestampType = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *Record) GetHeaders() map[string]string {
	if o == nil || o.Headers == nil {
		var ret map[string]string
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Record) GetHeadersOk() (*map[string]string, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *Record) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]string and assigns it to the Headers field.
func (o *Record) SetHeaders(v map[string]string) {
	o.Headers = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *Record) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Record) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *Record) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *Record) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value
func (o *Record) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Record) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Record) SetValue(v string) {
	o.Value = v
}

func (o Record) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Partition != nil {
		toSerialize["partition"] = o.Partition
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.TimestampType != nil {
		toSerialize["timestampType"] = o.TimestampType
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableRecord struct {
	value *Record
	isSet bool
}

func (v NullableRecord) Get() *Record {
	return v.value
}

func (v *NullableRecord) Set(val *Record) {
	v.value = val
	v.isSet = true
}

func (v NullableRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecord(val *Record) *NullableRecord {
	return &NullableRecord{value: val, isSet: true}
}

func (v NullableRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


