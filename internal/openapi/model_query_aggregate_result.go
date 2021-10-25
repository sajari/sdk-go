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

// QueryAggregateResult A query aggregate result contains results of aggregations.
type QueryAggregateResult struct {
	Metric     *QueryAggregateResultMetric     `json:"metric,omitempty"`
	Count      *QueryAggregateResultCount      `json:"count,omitempty"`
	Buckets    *QueryAggregateResultBuckets    `json:"buckets,omitempty"`
	Date       *QueryAggregateResultDate       `json:"date,omitempty"`
	Analysis   *QueryAggregateResultAnalysis   `json:"analysis,omitempty"`
	Percentile *QueryAggregateResultPercentile `json:"percentile,omitempty"`
}

// NewQueryAggregateResult instantiates a new QueryAggregateResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryAggregateResult() *QueryAggregateResult {
	this := QueryAggregateResult{}
	return &this
}

// NewQueryAggregateResultWithDefaults instantiates a new QueryAggregateResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryAggregateResultWithDefaults() *QueryAggregateResult {
	this := QueryAggregateResult{}
	return &this
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *QueryAggregateResult) GetMetric() QueryAggregateResultMetric {
	if o == nil || o.Metric == nil {
		var ret QueryAggregateResultMetric
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAggregateResult) GetMetricOk() (*QueryAggregateResultMetric, bool) {
	if o == nil || o.Metric == nil {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *QueryAggregateResult) HasMetric() bool {
	if o != nil && o.Metric != nil {
		return true
	}

	return false
}

// SetMetric gets a reference to the given QueryAggregateResultMetric and assigns it to the Metric field.
func (o *QueryAggregateResult) SetMetric(v QueryAggregateResultMetric) {
	o.Metric = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *QueryAggregateResult) GetCount() QueryAggregateResultCount {
	if o == nil || o.Count == nil {
		var ret QueryAggregateResultCount
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAggregateResult) GetCountOk() (*QueryAggregateResultCount, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *QueryAggregateResult) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given QueryAggregateResultCount and assigns it to the Count field.
func (o *QueryAggregateResult) SetCount(v QueryAggregateResultCount) {
	o.Count = &v
}

// GetBuckets returns the Buckets field value if set, zero value otherwise.
func (o *QueryAggregateResult) GetBuckets() QueryAggregateResultBuckets {
	if o == nil || o.Buckets == nil {
		var ret QueryAggregateResultBuckets
		return ret
	}
	return *o.Buckets
}

// GetBucketsOk returns a tuple with the Buckets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAggregateResult) GetBucketsOk() (*QueryAggregateResultBuckets, bool) {
	if o == nil || o.Buckets == nil {
		return nil, false
	}
	return o.Buckets, true
}

// HasBuckets returns a boolean if a field has been set.
func (o *QueryAggregateResult) HasBuckets() bool {
	if o != nil && o.Buckets != nil {
		return true
	}

	return false
}

// SetBuckets gets a reference to the given QueryAggregateResultBuckets and assigns it to the Buckets field.
func (o *QueryAggregateResult) SetBuckets(v QueryAggregateResultBuckets) {
	o.Buckets = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *QueryAggregateResult) GetDate() QueryAggregateResultDate {
	if o == nil || o.Date == nil {
		var ret QueryAggregateResultDate
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAggregateResult) GetDateOk() (*QueryAggregateResultDate, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *QueryAggregateResult) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given QueryAggregateResultDate and assigns it to the Date field.
func (o *QueryAggregateResult) SetDate(v QueryAggregateResultDate) {
	o.Date = &v
}

// GetAnalysis returns the Analysis field value if set, zero value otherwise.
func (o *QueryAggregateResult) GetAnalysis() QueryAggregateResultAnalysis {
	if o == nil || o.Analysis == nil {
		var ret QueryAggregateResultAnalysis
		return ret
	}
	return *o.Analysis
}

// GetAnalysisOk returns a tuple with the Analysis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAggregateResult) GetAnalysisOk() (*QueryAggregateResultAnalysis, bool) {
	if o == nil || o.Analysis == nil {
		return nil, false
	}
	return o.Analysis, true
}

// HasAnalysis returns a boolean if a field has been set.
func (o *QueryAggregateResult) HasAnalysis() bool {
	if o != nil && o.Analysis != nil {
		return true
	}

	return false
}

// SetAnalysis gets a reference to the given QueryAggregateResultAnalysis and assigns it to the Analysis field.
func (o *QueryAggregateResult) SetAnalysis(v QueryAggregateResultAnalysis) {
	o.Analysis = &v
}

// GetPercentile returns the Percentile field value if set, zero value otherwise.
func (o *QueryAggregateResult) GetPercentile() QueryAggregateResultPercentile {
	if o == nil || o.Percentile == nil {
		var ret QueryAggregateResultPercentile
		return ret
	}
	return *o.Percentile
}

// GetPercentileOk returns a tuple with the Percentile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAggregateResult) GetPercentileOk() (*QueryAggregateResultPercentile, bool) {
	if o == nil || o.Percentile == nil {
		return nil, false
	}
	return o.Percentile, true
}

// HasPercentile returns a boolean if a field has been set.
func (o *QueryAggregateResult) HasPercentile() bool {
	if o != nil && o.Percentile != nil {
		return true
	}

	return false
}

// SetPercentile gets a reference to the given QueryAggregateResultPercentile and assigns it to the Percentile field.
func (o *QueryAggregateResult) SetPercentile(v QueryAggregateResultPercentile) {
	o.Percentile = &v
}

func (o QueryAggregateResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Metric != nil {
		toSerialize["metric"] = o.Metric
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Buckets != nil {
		toSerialize["buckets"] = o.Buckets
	}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	if o.Analysis != nil {
		toSerialize["analysis"] = o.Analysis
	}
	if o.Percentile != nil {
		toSerialize["percentile"] = o.Percentile
	}
	return json.Marshal(toSerialize)
}

type NullableQueryAggregateResult struct {
	value *QueryAggregateResult
	isSet bool
}

func (v NullableQueryAggregateResult) Get() *QueryAggregateResult {
	return v.value
}

func (v *NullableQueryAggregateResult) Set(val *QueryAggregateResult) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryAggregateResult) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryAggregateResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryAggregateResult(val *QueryAggregateResult) *NullableQueryAggregateResult {
	return &NullableQueryAggregateResult{value: val, isSet: true}
}

func (v NullableQueryAggregateResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryAggregateResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
