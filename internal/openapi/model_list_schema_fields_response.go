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

// ListSchemaFieldsResponse struct for ListSchemaFieldsResponse
type ListSchemaFieldsResponse struct {
	// A token, which can be sent as `page_token` to retrieve the next page.  If this field is omitted, there are no subsequent pages.
	NextPageToken *string `json:"next_page_token,omitempty"`
	// The schema fields.
	SchemaFields *[]SchemaField `json:"schema_fields,omitempty"`
	// Maximum number of fields to return.
	TotalSize *int32 `json:"total_size,omitempty"`
}

// NewListSchemaFieldsResponse instantiates a new ListSchemaFieldsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSchemaFieldsResponse() *ListSchemaFieldsResponse {
	this := ListSchemaFieldsResponse{}
	return &this
}

// NewListSchemaFieldsResponseWithDefaults instantiates a new ListSchemaFieldsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSchemaFieldsResponseWithDefaults() *ListSchemaFieldsResponse {
	this := ListSchemaFieldsResponse{}
	return &this
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *ListSchemaFieldsResponse) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSchemaFieldsResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *ListSchemaFieldsResponse) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *ListSchemaFieldsResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetSchemaFields returns the SchemaFields field value if set, zero value otherwise.
func (o *ListSchemaFieldsResponse) GetSchemaFields() []SchemaField {
	if o == nil || o.SchemaFields == nil {
		var ret []SchemaField
		return ret
	}
	return *o.SchemaFields
}

// GetSchemaFieldsOk returns a tuple with the SchemaFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSchemaFieldsResponse) GetSchemaFieldsOk() (*[]SchemaField, bool) {
	if o == nil || o.SchemaFields == nil {
		return nil, false
	}
	return o.SchemaFields, true
}

// HasSchemaFields returns a boolean if a field has been set.
func (o *ListSchemaFieldsResponse) HasSchemaFields() bool {
	if o != nil && o.SchemaFields != nil {
		return true
	}

	return false
}

// SetSchemaFields gets a reference to the given []SchemaField and assigns it to the SchemaFields field.
func (o *ListSchemaFieldsResponse) SetSchemaFields(v []SchemaField) {
	o.SchemaFields = &v
}

// GetTotalSize returns the TotalSize field value if set, zero value otherwise.
func (o *ListSchemaFieldsResponse) GetTotalSize() int32 {
	if o == nil || o.TotalSize == nil {
		var ret int32
		return ret
	}
	return *o.TotalSize
}

// GetTotalSizeOk returns a tuple with the TotalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSchemaFieldsResponse) GetTotalSizeOk() (*int32, bool) {
	if o == nil || o.TotalSize == nil {
		return nil, false
	}
	return o.TotalSize, true
}

// HasTotalSize returns a boolean if a field has been set.
func (o *ListSchemaFieldsResponse) HasTotalSize() bool {
	if o != nil && o.TotalSize != nil {
		return true
	}

	return false
}

// SetTotalSize gets a reference to the given int32 and assigns it to the TotalSize field.
func (o *ListSchemaFieldsResponse) SetTotalSize(v int32) {
	o.TotalSize = &v
}

func (o ListSchemaFieldsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	if o.SchemaFields != nil {
		toSerialize["schema_fields"] = o.SchemaFields
	}
	if o.TotalSize != nil {
		toSerialize["total_size"] = o.TotalSize
	}
	return json.Marshal(toSerialize)
}

type NullableListSchemaFieldsResponse struct {
	value *ListSchemaFieldsResponse
	isSet bool
}

func (v NullableListSchemaFieldsResponse) Get() *ListSchemaFieldsResponse {
	return v.value
}

func (v *NullableListSchemaFieldsResponse) Set(val *ListSchemaFieldsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListSchemaFieldsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListSchemaFieldsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSchemaFieldsResponse(val *ListSchemaFieldsResponse) *NullableListSchemaFieldsResponse {
	return &NullableListSchemaFieldsResponse{value: val, isSet: true}
}

func (v NullableListSchemaFieldsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSchemaFieldsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
