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

// V4beta1UpsertRecordRequestPipeline The pipeline to use when upserting the record.  If not provided the default record pipeline is used.
type V4beta1UpsertRecordRequestPipeline struct {
	// The record pipeline's name, e.g. `my-pipeline`.
	Name string `json:"name"`
	// The record pipeline's version, e.g. `42`.  If not provided the default version is used.
	Version *string `json:"version,omitempty"`
}

// NewV4beta1UpsertRecordRequestPipeline instantiates a new V4beta1UpsertRecordRequestPipeline object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV4beta1UpsertRecordRequestPipeline(name string) *V4beta1UpsertRecordRequestPipeline {
	this := V4beta1UpsertRecordRequestPipeline{}
	this.Name = name
	return &this
}

// NewV4beta1UpsertRecordRequestPipelineWithDefaults instantiates a new V4beta1UpsertRecordRequestPipeline object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV4beta1UpsertRecordRequestPipelineWithDefaults() *V4beta1UpsertRecordRequestPipeline {
	this := V4beta1UpsertRecordRequestPipeline{}
	return &this
}

// GetName returns the Name field value
func (o *V4beta1UpsertRecordRequestPipeline) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *V4beta1UpsertRecordRequestPipeline) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *V4beta1UpsertRecordRequestPipeline) SetName(v string) {
	o.Name = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *V4beta1UpsertRecordRequestPipeline) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V4beta1UpsertRecordRequestPipeline) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *V4beta1UpsertRecordRequestPipeline) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *V4beta1UpsertRecordRequestPipeline) SetVersion(v string) {
	o.Version = &v
}

func (o V4beta1UpsertRecordRequestPipeline) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableV4beta1UpsertRecordRequestPipeline struct {
	value *V4beta1UpsertRecordRequestPipeline
	isSet bool
}

func (v NullableV4beta1UpsertRecordRequestPipeline) Get() *V4beta1UpsertRecordRequestPipeline {
	return v.value
}

func (v *NullableV4beta1UpsertRecordRequestPipeline) Set(val *V4beta1UpsertRecordRequestPipeline) {
	v.value = val
	v.isSet = true
}

func (v NullableV4beta1UpsertRecordRequestPipeline) IsSet() bool {
	return v.isSet
}

func (v *NullableV4beta1UpsertRecordRequestPipeline) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV4beta1UpsertRecordRequestPipeline(val *V4beta1UpsertRecordRequestPipeline) *NullableV4beta1UpsertRecordRequestPipeline {
	return &NullableV4beta1UpsertRecordRequestPipeline{value: val, isSet: true}
}

func (v NullableV4beta1UpsertRecordRequestPipeline) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV4beta1UpsertRecordRequestPipeline) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}