/*
 * Sajari API
 *
 * Sajari is a smart, highly-configurable, real-time search service that enables thousands of businesses worldwide to provide amazing search experiences on their websites, stores, and applications.
 *
 * API version: v4beta1
 * Contact: support@sajari.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// V4beta1BatchPutRecordsResponseKey struct for V4beta1BatchPutRecordsResponseKey
type V4beta1BatchPutRecordsResponseKey struct {
	// Index of the record in `records` that this key corresponds to.
	Index *int32            `json:"index,omitempty"`
	Key   *Sajariv4beta1Key `json:"key,omitempty"`
}

// NewV4beta1BatchPutRecordsResponseKey instantiates a new V4beta1BatchPutRecordsResponseKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV4beta1BatchPutRecordsResponseKey() *V4beta1BatchPutRecordsResponseKey {
	this := V4beta1BatchPutRecordsResponseKey{}
	return &this
}

// NewV4beta1BatchPutRecordsResponseKeyWithDefaults instantiates a new V4beta1BatchPutRecordsResponseKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV4beta1BatchPutRecordsResponseKeyWithDefaults() *V4beta1BatchPutRecordsResponseKey {
	this := V4beta1BatchPutRecordsResponseKey{}
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *V4beta1BatchPutRecordsResponseKey) GetIndex() int32 {
	if o == nil || o.Index == nil {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V4beta1BatchPutRecordsResponseKey) GetIndexOk() (*int32, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *V4beta1BatchPutRecordsResponseKey) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *V4beta1BatchPutRecordsResponseKey) SetIndex(v int32) {
	o.Index = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *V4beta1BatchPutRecordsResponseKey) GetKey() Sajariv4beta1Key {
	if o == nil || o.Key == nil {
		var ret Sajariv4beta1Key
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V4beta1BatchPutRecordsResponseKey) GetKeyOk() (*Sajariv4beta1Key, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *V4beta1BatchPutRecordsResponseKey) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given Sajariv4beta1Key and assigns it to the Key field.
func (o *V4beta1BatchPutRecordsResponseKey) SetKey(v Sajariv4beta1Key) {
	o.Key = &v
}

func (o V4beta1BatchPutRecordsResponseKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

type NullableV4beta1BatchPutRecordsResponseKey struct {
	value *V4beta1BatchPutRecordsResponseKey
	isSet bool
}

func (v NullableV4beta1BatchPutRecordsResponseKey) Get() *V4beta1BatchPutRecordsResponseKey {
	return v.value
}

func (v *NullableV4beta1BatchPutRecordsResponseKey) Set(val *V4beta1BatchPutRecordsResponseKey) {
	v.value = val
	v.isSet = true
}

func (v NullableV4beta1BatchPutRecordsResponseKey) IsSet() bool {
	return v.isSet
}

func (v *NullableV4beta1BatchPutRecordsResponseKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV4beta1BatchPutRecordsResponseKey(val *V4beta1BatchPutRecordsResponseKey) *NullableV4beta1BatchPutRecordsResponseKey {
	return &NullableV4beta1BatchPutRecordsResponseKey{value: val, isSet: true}
}

func (v NullableV4beta1BatchPutRecordsResponseKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV4beta1BatchPutRecordsResponseKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}