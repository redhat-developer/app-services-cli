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
	"fmt"
)

// RecordIncludedProperty the model 'RecordIncludedProperty'
type RecordIncludedProperty string

// List of RecordIncludedProperty
const (
	RECORDINCLUDEDPROPERTY_PARTITION RecordIncludedProperty = "partition"
	RECORDINCLUDEDPROPERTY_OFFSET RecordIncludedProperty = "offset"
	RECORDINCLUDEDPROPERTY_TIMESTAMP RecordIncludedProperty = "timestamp"
	RECORDINCLUDEDPROPERTY_TIMESTAMP_TYPE RecordIncludedProperty = "timestampType"
	RECORDINCLUDEDPROPERTY_HEADERS RecordIncludedProperty = "headers"
	RECORDINCLUDEDPROPERTY_KEY RecordIncludedProperty = "key"
	RECORDINCLUDEDPROPERTY_VALUE RecordIncludedProperty = "value"
)

var allowedRecordIncludedPropertyEnumValues = []RecordIncludedProperty{
	"partition",
	"offset",
	"timestamp",
	"timestampType",
	"headers",
	"key",
	"value",
}

func (v *RecordIncludedProperty) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RecordIncludedProperty(value)
	for _, existing := range allowedRecordIncludedPropertyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RecordIncludedProperty", value)
}

// NewRecordIncludedPropertyFromValue returns a pointer to a valid RecordIncludedProperty
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRecordIncludedPropertyFromValue(v string) (*RecordIncludedProperty, error) {
	ev := RecordIncludedProperty(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RecordIncludedProperty: valid values are %v", v, allowedRecordIncludedPropertyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RecordIncludedProperty) IsValid() bool {
	for _, existing := range allowedRecordIncludedPropertyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RecordIncludedProperty value
func (v RecordIncludedProperty) Ptr() *RecordIncludedProperty {
	return &v
}

type NullableRecordIncludedProperty struct {
	value *RecordIncludedProperty
	isSet bool
}

func (v NullableRecordIncludedProperty) Get() *RecordIncludedProperty {
	return v.value
}

func (v *NullableRecordIncludedProperty) Set(val *RecordIncludedProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordIncludedProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordIncludedProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordIncludedProperty(val *RecordIncludedProperty) *NullableRecordIncludedProperty {
	return &NullableRecordIncludedProperty{value: val, isSet: true}
}

func (v NullableRecordIncludedProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordIncludedProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

