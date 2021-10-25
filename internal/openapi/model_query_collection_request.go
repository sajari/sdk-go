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

// QueryCollectionRequest A request to perform a search using a pipeline.
type QueryCollectionRequest struct {
	Pipeline *QueryCollectionRequestPipeline `json:"pipeline,omitempty"`
	// The initial values for the variables the pipeline operates on and transforms throughout its steps.  The most important variable is `q` which is the query the user entered, for example:  ```json { \"q\": \"search terms\" } ```  To paginate through results, set the variables `page` and `resultsPerPage`, for example:  ```json { \"q\": \"search terms\", \"page\": 5, \"resultsPerPage\": 20 } ```  To sort results, set the variable `sort` to the name of one of your collection's schema fields, for example:  ```json { \"q\": \"search terms\", \"sort\": \"name\" } ```  To sort in reverse, prefix the schema field with a minus sign `-`, for example:  ```json { \"q\": \"search terms\", \"sort\": \"-name\" } ```
	Variables map[string]map[string]interface{} `json:"variables"`
	Tracking  *QueryCollectionRequestTracking   `json:"tracking,omitempty"`
}

// NewQueryCollectionRequest instantiates a new QueryCollectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryCollectionRequest(variables map[string]map[string]interface{}) *QueryCollectionRequest {
	this := QueryCollectionRequest{}
	this.Variables = variables
	return &this
}

// NewQueryCollectionRequestWithDefaults instantiates a new QueryCollectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryCollectionRequestWithDefaults() *QueryCollectionRequest {
	this := QueryCollectionRequest{}
	return &this
}

// GetPipeline returns the Pipeline field value if set, zero value otherwise.
func (o *QueryCollectionRequest) GetPipeline() QueryCollectionRequestPipeline {
	if o == nil || o.Pipeline == nil {
		var ret QueryCollectionRequestPipeline
		return ret
	}
	return *o.Pipeline
}

// GetPipelineOk returns a tuple with the Pipeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCollectionRequest) GetPipelineOk() (*QueryCollectionRequestPipeline, bool) {
	if o == nil || o.Pipeline == nil {
		return nil, false
	}
	return o.Pipeline, true
}

// HasPipeline returns a boolean if a field has been set.
func (o *QueryCollectionRequest) HasPipeline() bool {
	if o != nil && o.Pipeline != nil {
		return true
	}

	return false
}

// SetPipeline gets a reference to the given QueryCollectionRequestPipeline and assigns it to the Pipeline field.
func (o *QueryCollectionRequest) SetPipeline(v QueryCollectionRequestPipeline) {
	o.Pipeline = &v
}

// GetVariables returns the Variables field value
func (o *QueryCollectionRequest) GetVariables() map[string]map[string]interface{} {
	if o == nil {
		var ret map[string]map[string]interface{}
		return ret
	}

	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value
// and a boolean to check if the value has been set.
func (o *QueryCollectionRequest) GetVariablesOk() (*map[string]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variables, true
}

// SetVariables sets field value
func (o *QueryCollectionRequest) SetVariables(v map[string]map[string]interface{}) {
	o.Variables = v
}

// GetTracking returns the Tracking field value if set, zero value otherwise.
func (o *QueryCollectionRequest) GetTracking() QueryCollectionRequestTracking {
	if o == nil || o.Tracking == nil {
		var ret QueryCollectionRequestTracking
		return ret
	}
	return *o.Tracking
}

// GetTrackingOk returns a tuple with the Tracking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCollectionRequest) GetTrackingOk() (*QueryCollectionRequestTracking, bool) {
	if o == nil || o.Tracking == nil {
		return nil, false
	}
	return o.Tracking, true
}

// HasTracking returns a boolean if a field has been set.
func (o *QueryCollectionRequest) HasTracking() bool {
	if o != nil && o.Tracking != nil {
		return true
	}

	return false
}

// SetTracking gets a reference to the given QueryCollectionRequestTracking and assigns it to the Tracking field.
func (o *QueryCollectionRequest) SetTracking(v QueryCollectionRequestTracking) {
	o.Tracking = &v
}

func (o QueryCollectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Pipeline != nil {
		toSerialize["pipeline"] = o.Pipeline
	}
	if true {
		toSerialize["variables"] = o.Variables
	}
	if o.Tracking != nil {
		toSerialize["tracking"] = o.Tracking
	}
	return json.Marshal(toSerialize)
}

type NullableQueryCollectionRequest struct {
	value *QueryCollectionRequest
	isSet bool
}

func (v NullableQueryCollectionRequest) Get() *QueryCollectionRequest {
	return v.value
}

func (v *NullableQueryCollectionRequest) Set(val *QueryCollectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryCollectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryCollectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryCollectionRequest(val *QueryCollectionRequest) *NullableQueryCollectionRequest {
	return &NullableQueryCollectionRequest{value: val, isSet: true}
}

func (v NullableQueryCollectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryCollectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
