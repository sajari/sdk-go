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

// Enginev2Value Value is the representation for record field values.
type Enginev2Value struct {
	Null     *bool          `json:"null,omitempty"`
	Single   *string        `json:"single,omitempty"`
	Repeated *ValueRepeated `json:"repeated,omitempty"`
}

// NewEnginev2Value instantiates a new Enginev2Value object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnginev2Value() *Enginev2Value {
	this := Enginev2Value{}
	return &this
}

// NewEnginev2ValueWithDefaults instantiates a new Enginev2Value object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnginev2ValueWithDefaults() *Enginev2Value {
	this := Enginev2Value{}
	return &this
}

// GetNull returns the Null field value if set, zero value otherwise.
func (o *Enginev2Value) GetNull() bool {
	if o == nil || o.Null == nil {
		var ret bool
		return ret
	}
	return *o.Null
}

// GetNullOk returns a tuple with the Null field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Enginev2Value) GetNullOk() (*bool, bool) {
	if o == nil || o.Null == nil {
		return nil, false
	}
	return o.Null, true
}

// HasNull returns a boolean if a field has been set.
func (o *Enginev2Value) HasNull() bool {
	if o != nil && o.Null != nil {
		return true
	}

	return false
}

// SetNull gets a reference to the given bool and assigns it to the Null field.
func (o *Enginev2Value) SetNull(v bool) {
	o.Null = &v
}

// GetSingle returns the Single field value if set, zero value otherwise.
func (o *Enginev2Value) GetSingle() string {
	if o == nil || o.Single == nil {
		var ret string
		return ret
	}
	return *o.Single
}

// GetSingleOk returns a tuple with the Single field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Enginev2Value) GetSingleOk() (*string, bool) {
	if o == nil || o.Single == nil {
		return nil, false
	}
	return o.Single, true
}

// HasSingle returns a boolean if a field has been set.
func (o *Enginev2Value) HasSingle() bool {
	if o != nil && o.Single != nil {
		return true
	}

	return false
}

// SetSingle gets a reference to the given string and assigns it to the Single field.
func (o *Enginev2Value) SetSingle(v string) {
	o.Single = &v
}

// GetRepeated returns the Repeated field value if set, zero value otherwise.
func (o *Enginev2Value) GetRepeated() ValueRepeated {
	if o == nil || o.Repeated == nil {
		var ret ValueRepeated
		return ret
	}
	return *o.Repeated
}

// GetRepeatedOk returns a tuple with the Repeated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Enginev2Value) GetRepeatedOk() (*ValueRepeated, bool) {
	if o == nil || o.Repeated == nil {
		return nil, false
	}
	return o.Repeated, true
}

// HasRepeated returns a boolean if a field has been set.
func (o *Enginev2Value) HasRepeated() bool {
	if o != nil && o.Repeated != nil {
		return true
	}

	return false
}

// SetRepeated gets a reference to the given ValueRepeated and assigns it to the Repeated field.
func (o *Enginev2Value) SetRepeated(v ValueRepeated) {
	o.Repeated = &v
}

func (o Enginev2Value) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Null != nil {
		toSerialize["null"] = o.Null
	}
	if o.Single != nil {
		toSerialize["single"] = o.Single
	}
	if o.Repeated != nil {
		toSerialize["repeated"] = o.Repeated
	}
	return json.Marshal(toSerialize)
}

type NullableEnginev2Value struct {
	value *Enginev2Value
	isSet bool
}

func (v NullableEnginev2Value) Get() *Enginev2Value {
	return v.value
}

func (v *NullableEnginev2Value) Set(val *Enginev2Value) {
	v.value = val
	v.isSet = true
}

func (v NullableEnginev2Value) IsSet() bool {
	return v.isSet
}

func (v *NullableEnginev2Value) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnginev2Value(val *Enginev2Value) *NullableEnginev2Value {
	return &NullableEnginev2Value{value: val, isSet: true}
}

func (v NullableEnginev2Value) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnginev2Value) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
