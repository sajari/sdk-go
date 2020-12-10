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
	"fmt"
)

// SchemaFieldType Type represents the underlying data type of the field.   - TYPE_UNSPECIFIED: Type not specified.  - STRING: String values.  - INTEGER: Integer values (64-bit).  - FLOAT: Floating point values (32-bit).  - DOUBLE: Double floating point values (64-bit).  - BOOLEAN: Boolean values.  - TIMESTAMP: Timestamp values.
type SchemaFieldType string

// List of SchemaFieldType
const (
	SCHEMAFIELDTYPE_TYPE_UNSPECIFIED SchemaFieldType = "TYPE_UNSPECIFIED"
	SCHEMAFIELDTYPE_STRING           SchemaFieldType = "STRING"
	SCHEMAFIELDTYPE_INTEGER          SchemaFieldType = "INTEGER"
	SCHEMAFIELDTYPE_FLOAT            SchemaFieldType = "FLOAT"
	SCHEMAFIELDTYPE_DOUBLE           SchemaFieldType = "DOUBLE"
	SCHEMAFIELDTYPE_BOOLEAN          SchemaFieldType = "BOOLEAN"
	SCHEMAFIELDTYPE_TIMESTAMP        SchemaFieldType = "TIMESTAMP"
)

func (v *SchemaFieldType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SchemaFieldType(value)
	for _, existing := range []SchemaFieldType{"TYPE_UNSPECIFIED", "STRING", "INTEGER", "FLOAT", "DOUBLE", "BOOLEAN", "TIMESTAMP"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SchemaFieldType", value)
}

// Ptr returns reference to SchemaFieldType value
func (v SchemaFieldType) Ptr() *SchemaFieldType {
	return &v
}

type NullableSchemaFieldType struct {
	value *SchemaFieldType
	isSet bool
}

func (v NullableSchemaFieldType) Get() *SchemaFieldType {
	return v.value
}

func (v *NullableSchemaFieldType) Set(val *SchemaFieldType) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaFieldType) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaFieldType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaFieldType(val *SchemaFieldType) *NullableSchemaFieldType {
	return &NullableSchemaFieldType{value: val, isSet: true}
}

func (v NullableSchemaFieldType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaFieldType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
