/*
 * Sajari API
 *
 * Sajari is a smart, highly-configurable, real-time search service that enables thousands of businesses worldwide to provide amazing search experiences on their websites, stores, and applications.
 *
 * API version: v4
 * Contact: support@sajari.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// RecordKey struct for RecordKey
type RecordKey struct {
	// A field in your record that uniquely identifies it, e.g. `id`.
	Field string `json:"field"`
	// The value of `field` in your record, e.g. `b217a995-597c-410f-bef2-60e9f8c0aadd`.
	Value string `json:"value"`
}

// NewRecordKey instantiates a new RecordKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecordKey(field string, value string) *RecordKey {
	this := RecordKey{}
	this.Field = field
	this.Value = value
	return &this
}

// NewRecordKeyWithDefaults instantiates a new RecordKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordKeyWithDefaults() *RecordKey {
	this := RecordKey{}
	return &this
}

// GetField returns the Field field value
func (o *RecordKey) GetField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *RecordKey) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *RecordKey) SetField(v string) {
	o.Field = v
}

// GetValue returns the Value field value
func (o *RecordKey) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *RecordKey) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *RecordKey) SetValue(v string) {
	o.Value = v
}

func (o RecordKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["field"] = o.Field
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableRecordKey struct {
	value *RecordKey
	isSet bool
}

func (v NullableRecordKey) Get() *RecordKey {
	return v.value
}

func (v *NullableRecordKey) Set(val *RecordKey) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordKey) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordKey(val *RecordKey) *NullableRecordKey {
	return &NullableRecordKey{value: val, isSet: true}
}

func (v NullableRecordKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}