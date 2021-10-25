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

// QueryCollectionResponse struct for QueryCollectionResponse
type QueryCollectionResponse struct {
	Pipeline *QueryCollectionResponsePipeline `json:"pipeline,omitempty"`
	// The modified variables returned by the pipeline after it has finished processing.
	Variables *map[string]map[string]interface{} `json:"variables,omitempty"`
	// The results returned by the query.
	Results *[]QueryResult `json:"results,omitempty"`
	// The total number of results that match the query.
	TotalSize *string `json:"total_size,omitempty"`
	// The total time taken to perform the query.
	ProcessingDuration *string `json:"processing_duration,omitempty"`
	// The aggregates returned by the query.
	Aggregates *map[string]QueryAggregateResult `json:"aggregates,omitempty"`
	// The aggregates run with filters.
	AggregateFilters *map[string]QueryAggregateResult `json:"aggregate_filters,omitempty"`
	// A mapping of redirects triggered for all possible variations of the query.
	Redirects *map[string]RedirectResult `json:"redirects,omitempty"`
	// A list of the promotions activated when running the query.
	ActivePromotions *[]ActivePromotion `json:"active_promotions,omitempty"`
}

// NewQueryCollectionResponse instantiates a new QueryCollectionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryCollectionResponse() *QueryCollectionResponse {
	this := QueryCollectionResponse{}
	return &this
}

// NewQueryCollectionResponseWithDefaults instantiates a new QueryCollectionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryCollectionResponseWithDefaults() *QueryCollectionResponse {
	this := QueryCollectionResponse{}
	return &this
}

// GetPipeline returns the Pipeline field value if set, zero value otherwise.
func (o *QueryCollectionResponse) GetPipeline() QueryCollectionResponsePipeline {
	if o == nil || o.Pipeline == nil {
		var ret QueryCollectionResponsePipeline
		return ret
	}
	return *o.Pipeline
}

// GetPipelineOk returns a tuple with the Pipeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCollectionResponse) GetPipelineOk() (*QueryCollectionResponsePipeline, bool) {
	if o == nil || o.Pipeline == nil {
		return nil, false
	}
	return o.Pipeline, true
}

// HasPipeline returns a boolean if a field has been set.
func (o *QueryCollectionResponse) HasPipeline() bool {
	if o != nil && o.Pipeline != nil {
		return true
	}

	return false
}

// SetPipeline gets a reference to the given QueryCollectionResponsePipeline and assigns it to the Pipeline field.
func (o *QueryCollectionResponse) SetPipeline(v QueryCollectionResponsePipeline) {
	o.Pipeline = &v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *QueryCollectionResponse) GetVariables() map[string]map[string]interface{} {
	if o == nil || o.Variables == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCollectionResponse) GetVariablesOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.Variables == nil {
		return nil, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *QueryCollectionResponse) HasVariables() bool {
	if o != nil && o.Variables != nil {
		return true
	}

	return false
}

// SetVariables gets a reference to the given map[string]map[string]interface{} and assigns it to the Variables field.
func (o *QueryCollectionResponse) SetVariables(v map[string]map[string]interface{}) {
	o.Variables = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *QueryCollectionResponse) GetResults() []QueryResult {
	if o == nil || o.Results == nil {
		var ret []QueryResult
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCollectionResponse) GetResultsOk() (*[]QueryResult, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *QueryCollectionResponse) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []QueryResult and assigns it to the Results field.
func (o *QueryCollectionResponse) SetResults(v []QueryResult) {
	o.Results = &v
}

// GetTotalSize returns the TotalSize field value if set, zero value otherwise.
func (o *QueryCollectionResponse) GetTotalSize() string {
	if o == nil || o.TotalSize == nil {
		var ret string
		return ret
	}
	return *o.TotalSize
}

// GetTotalSizeOk returns a tuple with the TotalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCollectionResponse) GetTotalSizeOk() (*string, bool) {
	if o == nil || o.TotalSize == nil {
		return nil, false
	}
	return o.TotalSize, true
}

// HasTotalSize returns a boolean if a field has been set.
func (o *QueryCollectionResponse) HasTotalSize() bool {
	if o != nil && o.TotalSize != nil {
		return true
	}

	return false
}

// SetTotalSize gets a reference to the given string and assigns it to the TotalSize field.
func (o *QueryCollectionResponse) SetTotalSize(v string) {
	o.TotalSize = &v
}

// GetProcessingDuration returns the ProcessingDuration field value if set, zero value otherwise.
func (o *QueryCollectionResponse) GetProcessingDuration() string {
	if o == nil || o.ProcessingDuration == nil {
		var ret string
		return ret
	}
	return *o.ProcessingDuration
}

// GetProcessingDurationOk returns a tuple with the ProcessingDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCollectionResponse) GetProcessingDurationOk() (*string, bool) {
	if o == nil || o.ProcessingDuration == nil {
		return nil, false
	}
	return o.ProcessingDuration, true
}

// HasProcessingDuration returns a boolean if a field has been set.
func (o *QueryCollectionResponse) HasProcessingDuration() bool {
	if o != nil && o.ProcessingDuration != nil {
		return true
	}

	return false
}

// SetProcessingDuration gets a reference to the given string and assigns it to the ProcessingDuration field.
func (o *QueryCollectionResponse) SetProcessingDuration(v string) {
	o.ProcessingDuration = &v
}

// GetAggregates returns the Aggregates field value if set, zero value otherwise.
func (o *QueryCollectionResponse) GetAggregates() map[string]QueryAggregateResult {
	if o == nil || o.Aggregates == nil {
		var ret map[string]QueryAggregateResult
		return ret
	}
	return *o.Aggregates
}

// GetAggregatesOk returns a tuple with the Aggregates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCollectionResponse) GetAggregatesOk() (*map[string]QueryAggregateResult, bool) {
	if o == nil || o.Aggregates == nil {
		return nil, false
	}
	return o.Aggregates, true
}

// HasAggregates returns a boolean if a field has been set.
func (o *QueryCollectionResponse) HasAggregates() bool {
	if o != nil && o.Aggregates != nil {
		return true
	}

	return false
}

// SetAggregates gets a reference to the given map[string]QueryAggregateResult and assigns it to the Aggregates field.
func (o *QueryCollectionResponse) SetAggregates(v map[string]QueryAggregateResult) {
	o.Aggregates = &v
}

// GetAggregateFilters returns the AggregateFilters field value if set, zero value otherwise.
func (o *QueryCollectionResponse) GetAggregateFilters() map[string]QueryAggregateResult {
	if o == nil || o.AggregateFilters == nil {
		var ret map[string]QueryAggregateResult
		return ret
	}
	return *o.AggregateFilters
}

// GetAggregateFiltersOk returns a tuple with the AggregateFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCollectionResponse) GetAggregateFiltersOk() (*map[string]QueryAggregateResult, bool) {
	if o == nil || o.AggregateFilters == nil {
		return nil, false
	}
	return o.AggregateFilters, true
}

// HasAggregateFilters returns a boolean if a field has been set.
func (o *QueryCollectionResponse) HasAggregateFilters() bool {
	if o != nil && o.AggregateFilters != nil {
		return true
	}

	return false
}

// SetAggregateFilters gets a reference to the given map[string]QueryAggregateResult and assigns it to the AggregateFilters field.
func (o *QueryCollectionResponse) SetAggregateFilters(v map[string]QueryAggregateResult) {
	o.AggregateFilters = &v
}

// GetRedirects returns the Redirects field value if set, zero value otherwise.
func (o *QueryCollectionResponse) GetRedirects() map[string]RedirectResult {
	if o == nil || o.Redirects == nil {
		var ret map[string]RedirectResult
		return ret
	}
	return *o.Redirects
}

// GetRedirectsOk returns a tuple with the Redirects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCollectionResponse) GetRedirectsOk() (*map[string]RedirectResult, bool) {
	if o == nil || o.Redirects == nil {
		return nil, false
	}
	return o.Redirects, true
}

// HasRedirects returns a boolean if a field has been set.
func (o *QueryCollectionResponse) HasRedirects() bool {
	if o != nil && o.Redirects != nil {
		return true
	}

	return false
}

// SetRedirects gets a reference to the given map[string]RedirectResult and assigns it to the Redirects field.
func (o *QueryCollectionResponse) SetRedirects(v map[string]RedirectResult) {
	o.Redirects = &v
}

// GetActivePromotions returns the ActivePromotions field value if set, zero value otherwise.
func (o *QueryCollectionResponse) GetActivePromotions() []ActivePromotion {
	if o == nil || o.ActivePromotions == nil {
		var ret []ActivePromotion
		return ret
	}
	return *o.ActivePromotions
}

// GetActivePromotionsOk returns a tuple with the ActivePromotions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCollectionResponse) GetActivePromotionsOk() (*[]ActivePromotion, bool) {
	if o == nil || o.ActivePromotions == nil {
		return nil, false
	}
	return o.ActivePromotions, true
}

// HasActivePromotions returns a boolean if a field has been set.
func (o *QueryCollectionResponse) HasActivePromotions() bool {
	if o != nil && o.ActivePromotions != nil {
		return true
	}

	return false
}

// SetActivePromotions gets a reference to the given []ActivePromotion and assigns it to the ActivePromotions field.
func (o *QueryCollectionResponse) SetActivePromotions(v []ActivePromotion) {
	o.ActivePromotions = &v
}

func (o QueryCollectionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Pipeline != nil {
		toSerialize["pipeline"] = o.Pipeline
	}
	if o.Variables != nil {
		toSerialize["variables"] = o.Variables
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	if o.TotalSize != nil {
		toSerialize["total_size"] = o.TotalSize
	}
	if o.ProcessingDuration != nil {
		toSerialize["processing_duration"] = o.ProcessingDuration
	}
	if o.Aggregates != nil {
		toSerialize["aggregates"] = o.Aggregates
	}
	if o.AggregateFilters != nil {
		toSerialize["aggregate_filters"] = o.AggregateFilters
	}
	if o.Redirects != nil {
		toSerialize["redirects"] = o.Redirects
	}
	if o.ActivePromotions != nil {
		toSerialize["active_promotions"] = o.ActivePromotions
	}
	return json.Marshal(toSerialize)
}

type NullableQueryCollectionResponse struct {
	value *QueryCollectionResponse
	isSet bool
}

func (v NullableQueryCollectionResponse) Get() *QueryCollectionResponse {
	return v.value
}

func (v *NullableQueryCollectionResponse) Set(val *QueryCollectionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryCollectionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryCollectionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryCollectionResponse(val *QueryCollectionResponse) *NullableQueryCollectionResponse {
	return &NullableQueryCollectionResponse{value: val, isSet: true}
}

func (v NullableQueryCollectionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryCollectionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
