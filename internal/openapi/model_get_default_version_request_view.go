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

// GetDefaultVersionRequestView  - VIEW_UNSPECIFIED: The default / unset value. The API defaults to the `BASIC` view.  - BASIC: Include basic information including type, name, version and description but not the full step configuration. This is the default value (for both [ListPipelines](/api#operation/ListPipelines) and [GetPipeline](/api#operation/GetPipeline)).  - FULL: Include the information from `BASIC`, plus full step configuration.
type GetDefaultVersionRequestView string

// List of GetDefaultVersionRequestView
const (
	GETDEFAULTVERSIONREQUESTVIEW_VIEW_UNSPECIFIED GetDefaultVersionRequestView = "VIEW_UNSPECIFIED"
	GETDEFAULTVERSIONREQUESTVIEW_BASIC            GetDefaultVersionRequestView = "BASIC"
	GETDEFAULTVERSIONREQUESTVIEW_FULL             GetDefaultVersionRequestView = "FULL"
)

func (v *GetDefaultVersionRequestView) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GetDefaultVersionRequestView(value)
	for _, existing := range []GetDefaultVersionRequestView{"VIEW_UNSPECIFIED", "BASIC", "FULL"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GetDefaultVersionRequestView", value)
}

// Ptr returns reference to GetDefaultVersionRequestView value
func (v GetDefaultVersionRequestView) Ptr() *GetDefaultVersionRequestView {
	return &v
}

type NullableGetDefaultVersionRequestView struct {
	value *GetDefaultVersionRequestView
	isSet bool
}

func (v NullableGetDefaultVersionRequestView) Get() *GetDefaultVersionRequestView {
	return v.value
}

func (v *NullableGetDefaultVersionRequestView) Set(val *GetDefaultVersionRequestView) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDefaultVersionRequestView) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDefaultVersionRequestView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDefaultVersionRequestView(val *GetDefaultVersionRequestView) *NullableGetDefaultVersionRequestView {
	return &NullableGetDefaultVersionRequestView{value: val, isSet: true}
}

func (v NullableGetDefaultVersionRequestView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDefaultVersionRequestView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
