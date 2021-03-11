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
	"fmt"
)

// GetPipelineRequestView  - VIEW_UNSPECIFIED: The default / unset value. The API defaults to the `BASIC` view.  - BASIC: Include basic information including type, name, version and description but not the full step configuration. This is the default value (for both [ListPipelines](/api#operation/ListPipelines) and [GetPipeline](/api#operation/GetPipeline)).  - FULL: Include the information from `BASIC`, plus full step configuration.
type GetPipelineRequestView string

// List of GetPipelineRequestView
const (
	GETPIPELINEREQUESTVIEW_VIEW_UNSPECIFIED GetPipelineRequestView = "VIEW_UNSPECIFIED"
	GETPIPELINEREQUESTVIEW_BASIC            GetPipelineRequestView = "BASIC"
	GETPIPELINEREQUESTVIEW_FULL             GetPipelineRequestView = "FULL"
)

func (v *GetPipelineRequestView) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GetPipelineRequestView(value)
	for _, existing := range []GetPipelineRequestView{"VIEW_UNSPECIFIED", "BASIC", "FULL"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GetPipelineRequestView", value)
}

// Ptr returns reference to GetPipelineRequestView value
func (v GetPipelineRequestView) Ptr() *GetPipelineRequestView {
	return &v
}

type NullableGetPipelineRequestView struct {
	value *GetPipelineRequestView
	isSet bool
}

func (v NullableGetPipelineRequestView) Get() *GetPipelineRequestView {
	return v.value
}

func (v *NullableGetPipelineRequestView) Set(val *GetPipelineRequestView) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPipelineRequestView) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPipelineRequestView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPipelineRequestView(val *GetPipelineRequestView) *NullableGetPipelineRequestView {
	return &NullableGetPipelineRequestView{value: val, isSet: true}
}

func (v NullableGetPipelineRequestView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPipelineRequestView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}