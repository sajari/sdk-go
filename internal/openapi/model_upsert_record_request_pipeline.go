/*
 * Search.io API
 *
 * Search.io is a smart, highly-configurable, real-time search service that enables thousands of businesses worldwide to provide amazing search experiences on their websites, stores, and applications.
 *
 * API version: v4
 * Contact: support@search.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// UpsertRecordRequestPipeline The pipeline to use when upserting the record.  If not provided the default record pipeline is used.
type UpsertRecordRequestPipeline struct {
	// The record pipeline's name, e.g. `my-pipeline`.
	Name string `json:"name"`
	// The record pipeline's version, e.g. `42`.  If not provided the default version is used.
	Version *string `json:"version,omitempty"`
}

// NewUpsertRecordRequestPipeline instantiates a new UpsertRecordRequestPipeline object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpsertRecordRequestPipeline(name string) *UpsertRecordRequestPipeline {
	this := UpsertRecordRequestPipeline{}
	this.Name = name
	return &this
}

// NewUpsertRecordRequestPipelineWithDefaults instantiates a new UpsertRecordRequestPipeline object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpsertRecordRequestPipelineWithDefaults() *UpsertRecordRequestPipeline {
	this := UpsertRecordRequestPipeline{}
	return &this
}

// GetName returns the Name field value
func (o *UpsertRecordRequestPipeline) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpsertRecordRequestPipeline) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpsertRecordRequestPipeline) SetName(v string) {
	o.Name = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *UpsertRecordRequestPipeline) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpsertRecordRequestPipeline) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *UpsertRecordRequestPipeline) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *UpsertRecordRequestPipeline) SetVersion(v string) {
	o.Version = &v
}

func (o UpsertRecordRequestPipeline) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableUpsertRecordRequestPipeline struct {
	value *UpsertRecordRequestPipeline
	isSet bool
}

func (v NullableUpsertRecordRequestPipeline) Get() *UpsertRecordRequestPipeline {
	return v.value
}

func (v *NullableUpsertRecordRequestPipeline) Set(val *UpsertRecordRequestPipeline) {
	v.value = val
	v.isSet = true
}

func (v NullableUpsertRecordRequestPipeline) IsSet() bool {
	return v.isSet
}

func (v *NullableUpsertRecordRequestPipeline) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpsertRecordRequestPipeline(val *UpsertRecordRequestPipeline) *NullableUpsertRecordRequestPipeline {
	return &NullableUpsertRecordRequestPipeline{value: val, isSet: true}
}

func (v NullableUpsertRecordRequestPipeline) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpsertRecordRequestPipeline) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
