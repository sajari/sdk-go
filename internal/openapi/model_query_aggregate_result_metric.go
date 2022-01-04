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

// QueryAggregateResultMetric Metric represents the metric type requested, represented by an Enumeration Type.
type QueryAggregateResultMetric struct {
	Value *float64 `json:"value,omitempty"`
}

// NewQueryAggregateResultMetric instantiates a new QueryAggregateResultMetric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryAggregateResultMetric() *QueryAggregateResultMetric {
	this := QueryAggregateResultMetric{}
	return &this
}

// NewQueryAggregateResultMetricWithDefaults instantiates a new QueryAggregateResultMetric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryAggregateResultMetricWithDefaults() *QueryAggregateResultMetric {
	this := QueryAggregateResultMetric{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *QueryAggregateResultMetric) GetValue() float64 {
	if o == nil || o.Value == nil {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAggregateResultMetric) GetValueOk() (*float64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *QueryAggregateResultMetric) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *QueryAggregateResultMetric) SetValue(v float64) {
	o.Value = &v
}

func (o QueryAggregateResultMetric) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableQueryAggregateResultMetric struct {
	value *QueryAggregateResultMetric
	isSet bool
}

func (v NullableQueryAggregateResultMetric) Get() *QueryAggregateResultMetric {
	return v.value
}

func (v *NullableQueryAggregateResultMetric) Set(val *QueryAggregateResultMetric) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryAggregateResultMetric) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryAggregateResultMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryAggregateResultMetric(val *QueryAggregateResultMetric) *NullableQueryAggregateResultMetric {
	return &NullableQueryAggregateResultMetric{value: val, isSet: true}
}

func (v NullableQueryAggregateResultMetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryAggregateResultMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
