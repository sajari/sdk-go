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

// PromotionExclusion struct for PromotionExclusion
type PromotionExclusion struct {
	Key *RecordKey `json:"key,omitempty"`
}

// NewPromotionExclusion instantiates a new PromotionExclusion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotionExclusion() *PromotionExclusion {
	this := PromotionExclusion{}
	return &this
}

// NewPromotionExclusionWithDefaults instantiates a new PromotionExclusion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotionExclusionWithDefaults() *PromotionExclusion {
	this := PromotionExclusion{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *PromotionExclusion) GetKey() RecordKey {
	if o == nil || o.Key == nil {
		var ret RecordKey
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionExclusion) GetKeyOk() (*RecordKey, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *PromotionExclusion) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given RecordKey and assigns it to the Key field.
func (o *PromotionExclusion) SetKey(v RecordKey) {
	o.Key = &v
}

func (o PromotionExclusion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

type NullablePromotionExclusion struct {
	value *PromotionExclusion
	isSet bool
}

func (v NullablePromotionExclusion) Get() *PromotionExclusion {
	return v.value
}

func (v *NullablePromotionExclusion) Set(val *PromotionExclusion) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotionExclusion) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotionExclusion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotionExclusion(val *PromotionExclusion) *NullablePromotionExclusion {
	return &NullablePromotionExclusion{value: val, isSet: true}
}

func (v NullablePromotionExclusion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotionExclusion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
