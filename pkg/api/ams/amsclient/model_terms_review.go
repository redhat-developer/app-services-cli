/*
 * Account Management Service API
 *
 * Manage user subscriptions and clusters
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package amsclient

import (
	"encoding/json"
)

// TermsReview struct for TermsReview
type TermsReview struct {
	AccountUsername string `json:"account_username"`
}

// NewTermsReview instantiates a new TermsReview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTermsReview(accountUsername string) *TermsReview {
	this := TermsReview{}
	this.AccountUsername = accountUsername
	return &this
}

// NewTermsReviewWithDefaults instantiates a new TermsReview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTermsReviewWithDefaults() *TermsReview {
	this := TermsReview{}
	return &this
}

// GetAccountUsername returns the AccountUsername field value
func (o *TermsReview) GetAccountUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountUsername
}

// GetAccountUsernameOk returns a tuple with the AccountUsername field value
// and a boolean to check if the value has been set.
func (o *TermsReview) GetAccountUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountUsername, true
}

// SetAccountUsername sets field value
func (o *TermsReview) SetAccountUsername(v string) {
	o.AccountUsername = v
}

func (o TermsReview) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_username"] = o.AccountUsername
	}
	return json.Marshal(toSerialize)
}

type NullableTermsReview struct {
	value *TermsReview
	isSet bool
}

func (v NullableTermsReview) Get() *TermsReview {
	return v.value
}

func (v *NullableTermsReview) Set(val *TermsReview) {
	v.value = val
	v.isSet = true
}

func (v NullableTermsReview) IsSet() bool {
	return v.isSet
}

func (v *NullableTermsReview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTermsReview(val *TermsReview) *NullableTermsReview {
	return &NullableTermsReview{value: val, isSet: true}
}

func (v NullableTermsReview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTermsReview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
