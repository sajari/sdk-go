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

// SetDefaultPipelineRequest struct for SetDefaultPipelineRequest
type SetDefaultPipelineRequest struct {
	// The name of the pipeline to use when not otherwise specified.
	Pipeline string       `json:"pipeline"`
	Type     PipelineType `json:"type"`
}

// NewSetDefaultPipelineRequest instantiates a new SetDefaultPipelineRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetDefaultPipelineRequest(pipeline string, type_ PipelineType) *SetDefaultPipelineRequest {
	this := SetDefaultPipelineRequest{}
	this.Pipeline = pipeline
	this.Type = type_
	return &this
}

// NewSetDefaultPipelineRequestWithDefaults instantiates a new SetDefaultPipelineRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetDefaultPipelineRequestWithDefaults() *SetDefaultPipelineRequest {
	this := SetDefaultPipelineRequest{}
	var type_ PipelineType = "TYPE_UNSPECIFIED"
	this.Type = type_
	return &this
}

// GetPipeline returns the Pipeline field value
func (o *SetDefaultPipelineRequest) GetPipeline() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pipeline
}

// GetPipelineOk returns a tuple with the Pipeline field value
// and a boolean to check if the value has been set.
func (o *SetDefaultPipelineRequest) GetPipelineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pipeline, true
}

// SetPipeline sets field value
func (o *SetDefaultPipelineRequest) SetPipeline(v string) {
	o.Pipeline = v
}

// GetType returns the Type field value
func (o *SetDefaultPipelineRequest) GetType() PipelineType {
	if o == nil {
		var ret PipelineType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SetDefaultPipelineRequest) GetTypeOk() (*PipelineType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SetDefaultPipelineRequest) SetType(v PipelineType) {
	o.Type = v
}

func (o SetDefaultPipelineRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pipeline"] = o.Pipeline
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableSetDefaultPipelineRequest struct {
	value *SetDefaultPipelineRequest
	isSet bool
}

func (v NullableSetDefaultPipelineRequest) Get() *SetDefaultPipelineRequest {
	return v.value
}

func (v *NullableSetDefaultPipelineRequest) Set(val *SetDefaultPipelineRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetDefaultPipelineRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetDefaultPipelineRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetDefaultPipelineRequest(val *SetDefaultPipelineRequest) *NullableSetDefaultPipelineRequest {
	return &NullableSetDefaultPipelineRequest{value: val, isSet: true}
}

func (v NullableSetDefaultPipelineRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetDefaultPipelineRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
