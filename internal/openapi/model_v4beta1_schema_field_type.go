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
	"fmt"
)

// V4beta1SchemaFieldType Type represents the underlying data type of the field.   - TYPE_UNSPECIFIED: Type not specified.  - STRING: String values.  - INTEGER: Integer values (64-bit).  - FLOAT: Floating point values (32-bit).  - DOUBLE: Double floating point values (64-bit).  - BOOLEAN: Boolean values.  - TIMESTAMP: Timestamp values.
type V4beta1SchemaFieldType string

// List of v4beta1SchemaFieldType
const (
	V4BETA1SCHEMAFIELDTYPE_TYPE_UNSPECIFIED V4beta1SchemaFieldType = "TYPE_UNSPECIFIED"
	V4BETA1SCHEMAFIELDTYPE_STRING           V4beta1SchemaFieldType = "STRING"
	V4BETA1SCHEMAFIELDTYPE_INTEGER          V4beta1SchemaFieldType = "INTEGER"
	V4BETA1SCHEMAFIELDTYPE_FLOAT            V4beta1SchemaFieldType = "FLOAT"
	V4BETA1SCHEMAFIELDTYPE_DOUBLE           V4beta1SchemaFieldType = "DOUBLE"
	V4BETA1SCHEMAFIELDTYPE_BOOLEAN          V4beta1SchemaFieldType = "BOOLEAN"
	V4BETA1SCHEMAFIELDTYPE_TIMESTAMP        V4beta1SchemaFieldType = "TIMESTAMP"
)

func (v *V4beta1SchemaFieldType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V4beta1SchemaFieldType(value)
	for _, existing := range []V4beta1SchemaFieldType{"TYPE_UNSPECIFIED", "STRING", "INTEGER", "FLOAT", "DOUBLE", "BOOLEAN", "TIMESTAMP"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V4beta1SchemaFieldType", value)
}

// Ptr returns reference to v4beta1SchemaFieldType value
func (v V4beta1SchemaFieldType) Ptr() *V4beta1SchemaFieldType {
	return &v
}

type NullableV4beta1SchemaFieldType struct {
	value *V4beta1SchemaFieldType
	isSet bool
}

func (v NullableV4beta1SchemaFieldType) Get() *V4beta1SchemaFieldType {
	return v.value
}

func (v *NullableV4beta1SchemaFieldType) Set(val *V4beta1SchemaFieldType) {
	v.value = val
	v.isSet = true
}

func (v NullableV4beta1SchemaFieldType) IsSet() bool {
	return v.isSet
}

func (v *NullableV4beta1SchemaFieldType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV4beta1SchemaFieldType(val *V4beta1SchemaFieldType) *NullableV4beta1SchemaFieldType {
	return &NullableV4beta1SchemaFieldType{value: val, isSet: true}
}

func (v NullableV4beta1SchemaFieldType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV4beta1SchemaFieldType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
