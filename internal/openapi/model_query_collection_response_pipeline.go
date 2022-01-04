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

// QueryCollectionResponsePipeline The resolved query pipeline that was used to run the query.
type QueryCollectionResponsePipeline struct {
	// The pipeline's name, e.g. `my-pipeline`.
	Name *string `json:"name,omitempty"`
	// The pipeline's version, e.g. `42`.
	Version *string `json:"version,omitempty"`
}

// NewQueryCollectionResponsePipeline instantiates a new QueryCollectionResponsePipeline object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryCollectionResponsePipeline() *QueryCollectionResponsePipeline {
	this := QueryCollectionResponsePipeline{}
	return &this
}

// NewQueryCollectionResponsePipelineWithDefaults instantiates a new QueryCollectionResponsePipeline object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryCollectionResponsePipelineWithDefaults() *QueryCollectionResponsePipeline {
	this := QueryCollectionResponsePipeline{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *QueryCollectionResponsePipeline) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCollectionResponsePipeline) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *QueryCollectionResponsePipeline) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *QueryCollectionResponsePipeline) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *QueryCollectionResponsePipeline) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCollectionResponsePipeline) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *QueryCollectionResponsePipeline) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *QueryCollectionResponsePipeline) SetVersion(v string) {
	o.Version = &v
}

func (o QueryCollectionResponsePipeline) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableQueryCollectionResponsePipeline struct {
	value *QueryCollectionResponsePipeline
	isSet bool
}

func (v NullableQueryCollectionResponsePipeline) Get() *QueryCollectionResponsePipeline {
	return v.value
}

func (v *NullableQueryCollectionResponsePipeline) Set(val *QueryCollectionResponsePipeline) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryCollectionResponsePipeline) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryCollectionResponsePipeline) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryCollectionResponsePipeline(val *QueryCollectionResponsePipeline) *NullableQueryCollectionResponsePipeline {
	return &NullableQueryCollectionResponsePipeline{value: val, isSet: true}
}

func (v NullableQueryCollectionResponsePipeline) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryCollectionResponsePipeline) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
