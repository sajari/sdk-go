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

// PromotionPin struct for PromotionPin
type PromotionPin struct {
	Key *RecordKey `json:"key,omitempty"`
	// Position the record should occupy in search results. The top position is position 1.  Doesn't need to be contiguous with other pins, i.e. there can be gaps in the pinned set that are filled with organic results.  In the case where there are insufficient search results pinned items are collapsed.
	Position *int32 `json:"position,omitempty"`
}

// NewPromotionPin instantiates a new PromotionPin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotionPin() *PromotionPin {
	this := PromotionPin{}
	return &this
}

// NewPromotionPinWithDefaults instantiates a new PromotionPin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotionPinWithDefaults() *PromotionPin {
	this := PromotionPin{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *PromotionPin) GetKey() RecordKey {
	if o == nil || o.Key == nil {
		var ret RecordKey
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionPin) GetKeyOk() (*RecordKey, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *PromotionPin) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given RecordKey and assigns it to the Key field.
func (o *PromotionPin) SetKey(v RecordKey) {
	o.Key = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *PromotionPin) GetPosition() int32 {
	if o == nil || o.Position == nil {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionPin) GetPositionOk() (*int32, bool) {
	if o == nil || o.Position == nil {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *PromotionPin) HasPosition() bool {
	if o != nil && o.Position != nil {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
func (o *PromotionPin) SetPosition(v int32) {
	o.Position = &v
}

func (o PromotionPin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Position != nil {
		toSerialize["position"] = o.Position
	}
	return json.Marshal(toSerialize)
}

type NullablePromotionPin struct {
	value *PromotionPin
	isSet bool
}

func (v NullablePromotionPin) Get() *PromotionPin {
	return v.value
}

func (v *NullablePromotionPin) Set(val *PromotionPin) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotionPin) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotionPin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotionPin(val *PromotionPin) *NullablePromotionPin {
	return &NullablePromotionPin{value: val, isSet: true}
}

func (v NullablePromotionPin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotionPin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
