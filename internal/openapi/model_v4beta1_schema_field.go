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

// V4beta1SchemaField SchemaField defines the properties of a field in the schema.
type V4beta1SchemaField struct {
	// The name of the field.
	Name string `json:"name"`
	// The description of the field.
	Description *string                `json:"description,omitempty"`
	Type        V4beta1SchemaFieldType `json:"type"`
	Mode        SchemaFieldMode        `json:"mode"`
	// List indicates if the field is a list of values.  A list is also known as an array.  For example, if `type` is string and `list` is `true`, then the field is a list of strings.
	List *bool `json:"list,omitempty"`
	// The required length of the list, if `list` is `true`.  This allows you to enforce that a list contains an exact number of items.  For example, to store a 2x2 vector, you could set `type` to float, `list` to `true` and `list_length` to `4`.
	ListLength *int32 `json:"list_length,omitempty"`
}

// NewV4beta1SchemaField instantiates a new V4beta1SchemaField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV4beta1SchemaField(name string, type_ V4beta1SchemaFieldType, mode SchemaFieldMode) *V4beta1SchemaField {
	this := V4beta1SchemaField{}
	this.Name = name
	this.Type = type_
	this.Mode = mode
	return &this
}

// NewV4beta1SchemaFieldWithDefaults instantiates a new V4beta1SchemaField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV4beta1SchemaFieldWithDefaults() *V4beta1SchemaField {
	this := V4beta1SchemaField{}
	var type_ V4beta1SchemaFieldType = "TYPE_UNSPECIFIED"
	this.Type = type_
	var mode SchemaFieldMode = "MODE_UNSPECIFIED"
	this.Mode = mode
	return &this
}

// GetName returns the Name field value
func (o *V4beta1SchemaField) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *V4beta1SchemaField) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *V4beta1SchemaField) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *V4beta1SchemaField) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V4beta1SchemaField) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *V4beta1SchemaField) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *V4beta1SchemaField) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value
func (o *V4beta1SchemaField) GetType() V4beta1SchemaFieldType {
	if o == nil {
		var ret V4beta1SchemaFieldType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *V4beta1SchemaField) GetTypeOk() (*V4beta1SchemaFieldType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *V4beta1SchemaField) SetType(v V4beta1SchemaFieldType) {
	o.Type = v
}

// GetMode returns the Mode field value
func (o *V4beta1SchemaField) GetMode() SchemaFieldMode {
	if o == nil {
		var ret SchemaFieldMode
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *V4beta1SchemaField) GetModeOk() (*SchemaFieldMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *V4beta1SchemaField) SetMode(v SchemaFieldMode) {
	o.Mode = v
}

// GetList returns the List field value if set, zero value otherwise.
func (o *V4beta1SchemaField) GetList() bool {
	if o == nil || o.List == nil {
		var ret bool
		return ret
	}
	return *o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V4beta1SchemaField) GetListOk() (*bool, bool) {
	if o == nil || o.List == nil {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *V4beta1SchemaField) HasList() bool {
	if o != nil && o.List != nil {
		return true
	}

	return false
}

// SetList gets a reference to the given bool and assigns it to the List field.
func (o *V4beta1SchemaField) SetList(v bool) {
	o.List = &v
}

// GetListLength returns the ListLength field value if set, zero value otherwise.
func (o *V4beta1SchemaField) GetListLength() int32 {
	if o == nil || o.ListLength == nil {
		var ret int32
		return ret
	}
	return *o.ListLength
}

// GetListLengthOk returns a tuple with the ListLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V4beta1SchemaField) GetListLengthOk() (*int32, bool) {
	if o == nil || o.ListLength == nil {
		return nil, false
	}
	return o.ListLength, true
}

// HasListLength returns a boolean if a field has been set.
func (o *V4beta1SchemaField) HasListLength() bool {
	if o != nil && o.ListLength != nil {
		return true
	}

	return false
}

// SetListLength gets a reference to the given int32 and assigns it to the ListLength field.
func (o *V4beta1SchemaField) SetListLength(v int32) {
	o.ListLength = &v
}

func (o V4beta1SchemaField) MarshalJSON() ([]byte, error) {
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
	if o.List != nil {
		toSerialize["list"] = o.List
	}
	if o.ListLength != nil {
		toSerialize["list_length"] = o.ListLength
	}
	return json.Marshal(toSerialize)
}

type NullableV4beta1SchemaField struct {
	value *V4beta1SchemaField
	isSet bool
}

func (v NullableV4beta1SchemaField) Get() *V4beta1SchemaField {
	return v.value
}

func (v *NullableV4beta1SchemaField) Set(val *V4beta1SchemaField) {
	v.value = val
	v.isSet = true
}

func (v NullableV4beta1SchemaField) IsSet() bool {
	return v.isSet
}

func (v *NullableV4beta1SchemaField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV4beta1SchemaField(val *V4beta1SchemaField) *NullableV4beta1SchemaField {
	return &NullableV4beta1SchemaField{value: val, isSet: true}
}

func (v NullableV4beta1SchemaField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV4beta1SchemaField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}