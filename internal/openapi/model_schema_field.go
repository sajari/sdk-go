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

// SchemaField SchemaField defines the properties of a field in the schema.
type SchemaField struct {
	// The name of the field.
	Name string `json:"name"`
	// The description of the field.
	Description *string         `json:"description,omitempty"`
	Type        SchemaFieldType `json:"type"`
	Mode        SchemaFieldMode `json:"mode"`
	// Array indicates if the field is an array of values.  For example, if `type` is string and `array` is `true`, then the field is an array of strings.
	Array *bool `json:"array,omitempty"`
	// The required length of the array, if `array` is `true`.  This allows you to enforce that an array contains an exact number of items.  For example, to store a 2x2 vector, you could set `type` to float, `array` to `true` and `array_length` to `4`.
	ArrayLength *int32 `json:"array_length,omitempty"`
}

// NewSchemaField instantiates a new SchemaField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaField(name string, type_ SchemaFieldType, mode SchemaFieldMode) *SchemaField {
	this := SchemaField{}
	this.Name = name
	this.Type = type_
	this.Mode = mode
	return &this
}

// NewSchemaFieldWithDefaults instantiates a new SchemaField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaFieldWithDefaults() *SchemaField {
	this := SchemaField{}
	var type_ SchemaFieldType = "TYPE_UNSPECIFIED"
	this.Type = type_
	var mode SchemaFieldMode = "MODE_UNSPECIFIED"
	this.Mode = mode
	return &this
}

// GetName returns the Name field value
func (o *SchemaField) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SchemaField) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SchemaField) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SchemaField) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaField) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SchemaField) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SchemaField) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value
func (o *SchemaField) GetType() SchemaFieldType {
	if o == nil {
		var ret SchemaFieldType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SchemaField) GetTypeOk() (*SchemaFieldType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SchemaField) SetType(v SchemaFieldType) {
	o.Type = v
}

// GetMode returns the Mode field value
func (o *SchemaField) GetMode() SchemaFieldMode {
	if o == nil {
		var ret SchemaFieldMode
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *SchemaField) GetModeOk() (*SchemaFieldMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *SchemaField) SetMode(v SchemaFieldMode) {
	o.Mode = v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *SchemaField) GetArray() bool {
	if o == nil || o.Array == nil {
		var ret bool
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaField) GetArrayOk() (*bool, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *SchemaField) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given bool and assigns it to the Array field.
func (o *SchemaField) SetArray(v bool) {
	o.Array = &v
}

// GetArrayLength returns the ArrayLength field value if set, zero value otherwise.
func (o *SchemaField) GetArrayLength() int32 {
	if o == nil || o.ArrayLength == nil {
		var ret int32
		return ret
	}
	return *o.ArrayLength
}

// GetArrayLengthOk returns a tuple with the ArrayLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaField) GetArrayLengthOk() (*int32, bool) {
	if o == nil || o.ArrayLength == nil {
		return nil, false
	}
	return o.ArrayLength, true
}

// HasArrayLength returns a boolean if a field has been set.
func (o *SchemaField) HasArrayLength() bool {
	if o != nil && o.ArrayLength != nil {
		return true
	}

	return false
}

// SetArrayLength gets a reference to the given int32 and assigns it to the ArrayLength field.
func (o *SchemaField) SetArrayLength(v int32) {
	o.ArrayLength = &v
}

func (o SchemaField) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["mode"] = o.Mode
	}
	if o.Array != nil {
		toSerialize["array"] = o.Array
	}
	if o.ArrayLength != nil {
		toSerialize["array_length"] = o.ArrayLength
	}
	return json.Marshal(toSerialize)
}

type NullableSchemaField struct {
	value *SchemaField
	isSet bool
}

func (v NullableSchemaField) Get() *SchemaField {
	return v.value
}

func (v *NullableSchemaField) Set(val *SchemaField) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaField) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaField(val *SchemaField) *NullableSchemaField {
	return &NullableSchemaField{value: val, isSet: true}
}

func (v NullableSchemaField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
