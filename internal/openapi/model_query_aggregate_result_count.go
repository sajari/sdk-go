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

// QueryAggregateResultCount Count contains the counts for the set of values returned.
type QueryAggregateResultCount struct {
	Counts *map[string]int32 `json:"counts,omitempty"`
}

// NewQueryAggregateResultCount instantiates a new QueryAggregateResultCount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryAggregateResultCount() *QueryAggregateResultCount {
	this := QueryAggregateResultCount{}
	return &this
}

// NewQueryAggregateResultCountWithDefaults instantiates a new QueryAggregateResultCount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryAggregateResultCountWithDefaults() *QueryAggregateResultCount {
	this := QueryAggregateResultCount{}
	return &this
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *QueryAggregateResultCount) GetCounts() map[string]int32 {
	if o == nil || o.Counts == nil {
		var ret map[string]int32
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAggregateResultCount) GetCountsOk() (*map[string]int32, bool) {
	if o == nil || o.Counts == nil {
		return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *QueryAggregateResultCount) HasCounts() bool {
	if o != nil && o.Counts != nil {
		return true
	}

	return false
}

// SetCounts gets a reference to the given map[string]int32 and assigns it to the Counts field.
func (o *QueryAggregateResultCount) SetCounts(v map[string]int32) {
	o.Counts = &v
}

func (o QueryAggregateResultCount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Counts != nil {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableQueryAggregateResultCount struct {
	value *QueryAggregateResultCount
	isSet bool
}

func (v NullableQueryAggregateResultCount) Get() *QueryAggregateResultCount {
	return v.value
}

func (v *NullableQueryAggregateResultCount) Set(val *QueryAggregateResultCount) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryAggregateResultCount) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryAggregateResultCount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryAggregateResultCount(val *QueryAggregateResultCount) *NullableQueryAggregateResultCount {
	return &NullableQueryAggregateResultCount{value: val, isSet: true}
}

func (v NullableQueryAggregateResultCount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryAggregateResultCount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
