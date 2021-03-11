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

// QueryCollectionRequestTracking struct for QueryCollectionRequestTracking
type QueryCollectionRequestTracking struct {
	Type *QueryCollectionRequestTrackingType `json:"type,omitempty"`
	// Query ID of the query. If this is empty, then one is generated.
	QueryId *string `json:"query_id,omitempty"`
	// Sequence number of query.
	Sequence *int32 `json:"sequence,omitempty"`
	// Tracking field used to identify records in the collection.  Must be unique schema field.
	Field *string `json:"field,omitempty"`
	// Custom values to be included in tracking data.
	Data *map[string]string `json:"data,omitempty"`
}

// NewQueryCollectionRequestTracking instantiates a new QueryCollectionRequestTracking object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryCollectionRequestTracking() *QueryCollectionRequestTracking {
	this := QueryCollectionRequestTracking{}
	var type_ QueryCollectionRequestTrackingType = "TYPE_UNSPECIFIED"
	this.Type = &type_
	return &this
}

// NewQueryCollectionRequestTrackingWithDefaults instantiates a new QueryCollectionRequestTracking object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryCollectionRequestTrackingWithDefaults() *QueryCollectionRequestTracking {
	this := QueryCollectionRequestTracking{}
	var type_ QueryCollectionRequestTrackingType = "TYPE_UNSPECIFIED"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *QueryCollectionRequestTracking) GetType() QueryCollectionRequestTrackingType {
	if o == nil || o.Type == nil {
		var ret QueryCollectionRequestTrackingType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCollectionRequestTracking) GetTypeOk() (*QueryCollectionRequestTrackingType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *QueryCollectionRequestTracking) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given QueryCollectionRequestTrackingType and assigns it to the Type field.
func (o *QueryCollectionRequestTracking) SetType(v QueryCollectionRequestTrackingType) {
	o.Type = &v
}

// GetQueryId returns the QueryId field value if set, zero value otherwise.
func (o *QueryCollectionRequestTracking) GetQueryId() string {
	if o == nil || o.QueryId == nil {
		var ret string
		return ret
	}
	return *o.QueryId
}

// GetQueryIdOk returns a tuple with the QueryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCollectionRequestTracking) GetQueryIdOk() (*string, bool) {
	if o == nil || o.QueryId == nil {
		return nil, false
	}
	return o.QueryId, true
}

// HasQueryId returns a boolean if a field has been set.
func (o *QueryCollectionRequestTracking) HasQueryId() bool {
	if o != nil && o.QueryId != nil {
		return true
	}

	return false
}

// SetQueryId gets a reference to the given string and assigns it to the QueryId field.
func (o *QueryCollectionRequestTracking) SetQueryId(v string) {
	o.QueryId = &v
}

// GetSequence returns the Sequence field value if set, zero value otherwise.
func (o *QueryCollectionRequestTracking) GetSequence() int32 {
	if o == nil || o.Sequence == nil {
		var ret int32
		return ret
	}
	return *o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCollectionRequestTracking) GetSequenceOk() (*int32, bool) {
	if o == nil || o.Sequence == nil {
		return nil, false
	}
	return o.Sequence, true
}

// HasSequence returns a boolean if a field has been set.
func (o *QueryCollectionRequestTracking) HasSequence() bool {
	if o != nil && o.Sequence != nil {
		return true
	}

	return false
}

// SetSequence gets a reference to the given int32 and assigns it to the Sequence field.
func (o *QueryCollectionRequestTracking) SetSequence(v int32) {
	o.Sequence = &v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *QueryCollectionRequestTracking) GetField() string {
	if o == nil || o.Field == nil {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCollectionRequestTracking) GetFieldOk() (*string, bool) {
	if o == nil || o.Field == nil {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *QueryCollectionRequestTracking) HasField() bool {
	if o != nil && o.Field != nil {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *QueryCollectionRequestTracking) SetField(v string) {
	o.Field = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *QueryCollectionRequestTracking) GetData() map[string]string {
	if o == nil || o.Data == nil {
		var ret map[string]string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCollectionRequestTracking) GetDataOk() (*map[string]string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *QueryCollectionRequestTracking) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]string and assigns it to the Data field.
func (o *QueryCollectionRequestTracking) SetData(v map[string]string) {
	o.Data = &v
}

func (o QueryCollectionRequestTracking) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.QueryId != nil {
		toSerialize["query_id"] = o.QueryId
	}
	if o.Sequence != nil {
		toSerialize["sequence"] = o.Sequence
	}
	if o.Field != nil {
		toSerialize["field"] = o.Field
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableQueryCollectionRequestTracking struct {
	value *QueryCollectionRequestTracking
	isSet bool
}

func (v NullableQueryCollectionRequestTracking) Get() *QueryCollectionRequestTracking {
	return v.value
}

func (v *NullableQueryCollectionRequestTracking) Set(val *QueryCollectionRequestTracking) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryCollectionRequestTracking) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryCollectionRequestTracking) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryCollectionRequestTracking(val *QueryCollectionRequestTracking) *NullableQueryCollectionRequestTracking {
	return &NullableQueryCollectionRequestTracking{value: val, isSet: true}
}

func (v NullableQueryCollectionRequestTracking) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryCollectionRequestTracking) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}