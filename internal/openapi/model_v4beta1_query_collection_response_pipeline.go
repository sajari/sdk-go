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

// V4beta1QueryCollectionResponsePipeline The resolved query pipeline that was used to run the query.
type V4beta1QueryCollectionResponsePipeline struct {
	// The pipeline's name, e.g. `my-pipeline`.
	Name *string `json:"name,omitempty"`
	// The pipeline's version, e.g. `42`.
	Version *string `json:"version,omitempty"`
}

// NewV4beta1QueryCollectionResponsePipeline instantiates a new V4beta1QueryCollectionResponsePipeline object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV4beta1QueryCollectionResponsePipeline() *V4beta1QueryCollectionResponsePipeline {
	this := V4beta1QueryCollectionResponsePipeline{}
	return &this
}

// NewV4beta1QueryCollectionResponsePipelineWithDefaults instantiates a new V4beta1QueryCollectionResponsePipeline object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV4beta1QueryCollectionResponsePipelineWithDefaults() *V4beta1QueryCollectionResponsePipeline {
	this := V4beta1QueryCollectionResponsePipeline{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V4beta1QueryCollectionResponsePipeline) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V4beta1QueryCollectionResponsePipeline) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V4beta1QueryCollectionResponsePipeline) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V4beta1QueryCollectionResponsePipeline) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *V4beta1QueryCollectionResponsePipeline) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V4beta1QueryCollectionResponsePipeline) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *V4beta1QueryCollectionResponsePipeline) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *V4beta1QueryCollectionResponsePipeline) SetVersion(v string) {
	o.Version = &v
}

func (o V4beta1QueryCollectionResponsePipeline) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableV4beta1QueryCollectionResponsePipeline struct {
	value *V4beta1QueryCollectionResponsePipeline
	isSet bool
}

func (v NullableV4beta1QueryCollectionResponsePipeline) Get() *V4beta1QueryCollectionResponsePipeline {
	return v.value
}

func (v *NullableV4beta1QueryCollectionResponsePipeline) Set(val *V4beta1QueryCollectionResponsePipeline) {
	v.value = val
	v.isSet = true
}

func (v NullableV4beta1QueryCollectionResponsePipeline) IsSet() bool {
	return v.isSet
}

func (v *NullableV4beta1QueryCollectionResponsePipeline) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV4beta1QueryCollectionResponsePipeline(val *V4beta1QueryCollectionResponsePipeline) *NullableV4beta1QueryCollectionResponsePipeline {
	return &NullableV4beta1QueryCollectionResponsePipeline{value: val, isSet: true}
}

func (v NullableV4beta1QueryCollectionResponsePipeline) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV4beta1QueryCollectionResponsePipeline) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
