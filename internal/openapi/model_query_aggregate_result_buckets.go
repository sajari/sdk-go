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

// QueryAggregateResultBuckets Buckets is a full set of buckets computed in an aggregation.
type QueryAggregateResultBuckets struct {
	Buckets *map[string]BucketsBucket `json:"buckets,omitempty"`
}

// NewQueryAggregateResultBuckets instantiates a new QueryAggregateResultBuckets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryAggregateResultBuckets() *QueryAggregateResultBuckets {
	this := QueryAggregateResultBuckets{}
	return &this
}

// NewQueryAggregateResultBucketsWithDefaults instantiates a new QueryAggregateResultBuckets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryAggregateResultBucketsWithDefaults() *QueryAggregateResultBuckets {
	this := QueryAggregateResultBuckets{}
	return &this
}

// GetBuckets returns the Buckets field value if set, zero value otherwise.
func (o *QueryAggregateResultBuckets) GetBuckets() map[string]BucketsBucket {
	if o == nil || o.Buckets == nil {
		var ret map[string]BucketsBucket
		return ret
	}
	return *o.Buckets
}

// GetBucketsOk returns a tuple with the Buckets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAggregateResultBuckets) GetBucketsOk() (*map[string]BucketsBucket, bool) {
	if o == nil || o.Buckets == nil {
		return nil, false
	}
	return o.Buckets, true
}

// HasBuckets returns a boolean if a field has been set.
func (o *QueryAggregateResultBuckets) HasBuckets() bool {
	if o != nil && o.Buckets != nil {
		return true
	}

	return false
}

// SetBuckets gets a reference to the given map[string]BucketsBucket and assigns it to the Buckets field.
func (o *QueryAggregateResultBuckets) SetBuckets(v map[string]BucketsBucket) {
	o.Buckets = &v
}

func (o QueryAggregateResultBuckets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Buckets != nil {
		toSerialize["buckets"] = o.Buckets
	}
	return json.Marshal(toSerialize)
}

type NullableQueryAggregateResultBuckets struct {
	value *QueryAggregateResultBuckets
	isSet bool
}

func (v NullableQueryAggregateResultBuckets) Get() *QueryAggregateResultBuckets {
	return v.value
}

func (v *NullableQueryAggregateResultBuckets) Set(val *QueryAggregateResultBuckets) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryAggregateResultBuckets) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryAggregateResultBuckets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryAggregateResultBuckets(val *QueryAggregateResultBuckets) *NullableQueryAggregateResultBuckets {
	return &NullableQueryAggregateResultBuckets{value: val, isSet: true}
}

func (v NullableQueryAggregateResultBuckets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryAggregateResultBuckets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}