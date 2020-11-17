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

// BatchUpsertRecordsResponseVariables struct for BatchUpsertRecordsResponseVariables
type BatchUpsertRecordsResponseVariables struct {
	// Index of the record in `records` that these variables correspond to.
	Index *int32 `json:"index,omitempty"`
	// The variables.
	Variables *map[string]interface{} `json:"variables,omitempty"`
}

// NewBatchUpsertRecordsResponseVariables instantiates a new BatchUpsertRecordsResponseVariables object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchUpsertRecordsResponseVariables() *BatchUpsertRecordsResponseVariables {
	this := BatchUpsertRecordsResponseVariables{}
	return &this
}

// NewBatchUpsertRecordsResponseVariablesWithDefaults instantiates a new BatchUpsertRecordsResponseVariables object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchUpsertRecordsResponseVariablesWithDefaults() *BatchUpsertRecordsResponseVariables {
	this := BatchUpsertRecordsResponseVariables{}
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *BatchUpsertRecordsResponseVariables) GetIndex() int32 {
	if o == nil || o.Index == nil {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchUpsertRecordsResponseVariables) GetIndexOk() (*int32, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *BatchUpsertRecordsResponseVariables) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *BatchUpsertRecordsResponseVariables) SetIndex(v int32) {
	o.Index = &v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *BatchUpsertRecordsResponseVariables) GetVariables() map[string]interface{} {
	if o == nil || o.Variables == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchUpsertRecordsResponseVariables) GetVariablesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Variables == nil {
		return nil, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *BatchUpsertRecordsResponseVariables) HasVariables() bool {
	if o != nil && o.Variables != nil {
		return true
	}

	return false
}

// SetVariables gets a reference to the given map[string]interface{} and assigns it to the Variables field.
func (o *BatchUpsertRecordsResponseVariables) SetVariables(v map[string]interface{}) {
	o.Variables = &v
}

func (o BatchUpsertRecordsResponseVariables) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	if o.Variables != nil {
		toSerialize["variables"] = o.Variables
	}
	return json.Marshal(toSerialize)
}

type NullableBatchUpsertRecordsResponseVariables struct {
	value *BatchUpsertRecordsResponseVariables
	isSet bool
}

func (v NullableBatchUpsertRecordsResponseVariables) Get() *BatchUpsertRecordsResponseVariables {
	return v.value
}

func (v *NullableBatchUpsertRecordsResponseVariables) Set(val *BatchUpsertRecordsResponseVariables) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchUpsertRecordsResponseVariables) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchUpsertRecordsResponseVariables) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchUpsertRecordsResponseVariables(val *BatchUpsertRecordsResponseVariables) *NullableBatchUpsertRecordsResponseVariables {
	return &NullableBatchUpsertRecordsResponseVariables{value: val, isSet: true}
}

func (v NullableBatchUpsertRecordsResponseVariables) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchUpsertRecordsResponseVariables) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}