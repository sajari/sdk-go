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

// GetDefaultPipelineResponse struct for GetDefaultPipelineResponse
type GetDefaultPipelineResponse struct {
	// The name of the pipeline to use when not otherwise specified.
	Pipeline *string `json:"pipeline,omitempty"`
}

// NewGetDefaultPipelineResponse instantiates a new GetDefaultPipelineResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDefaultPipelineResponse() *GetDefaultPipelineResponse {
	this := GetDefaultPipelineResponse{}
	return &this
}

// NewGetDefaultPipelineResponseWithDefaults instantiates a new GetDefaultPipelineResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDefaultPipelineResponseWithDefaults() *GetDefaultPipelineResponse {
	this := GetDefaultPipelineResponse{}
	return &this
}

// GetPipeline returns the Pipeline field value if set, zero value otherwise.
func (o *GetDefaultPipelineResponse) GetPipeline() string {
	if o == nil || o.Pipeline == nil {
		var ret string
		return ret
	}
	return *o.Pipeline
}

// GetPipelineOk returns a tuple with the Pipeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDefaultPipelineResponse) GetPipelineOk() (*string, bool) {
	if o == nil || o.Pipeline == nil {
		return nil, false
	}
	return o.Pipeline, true
}

// HasPipeline returns a boolean if a field has been set.
func (o *GetDefaultPipelineResponse) HasPipeline() bool {
	if o != nil && o.Pipeline != nil {
		return true
	}

	return false
}

// SetPipeline gets a reference to the given string and assigns it to the Pipeline field.
func (o *GetDefaultPipelineResponse) SetPipeline(v string) {
	o.Pipeline = &v
}

func (o GetDefaultPipelineResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Pipeline != nil {
		toSerialize["pipeline"] = o.Pipeline
	}
	return json.Marshal(toSerialize)
}

type NullableGetDefaultPipelineResponse struct {
	value *GetDefaultPipelineResponse
	isSet bool
}

func (v NullableGetDefaultPipelineResponse) Get() *GetDefaultPipelineResponse {
	return v.value
}

func (v *NullableGetDefaultPipelineResponse) Set(val *GetDefaultPipelineResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDefaultPipelineResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDefaultPipelineResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDefaultPipelineResponse(val *GetDefaultPipelineResponse) *NullableGetDefaultPipelineResponse {
	return &NullableGetDefaultPipelineResponse{value: val, isSet: true}
}

func (v NullableGetDefaultPipelineResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDefaultPipelineResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}