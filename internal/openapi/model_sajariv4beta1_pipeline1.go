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
	"time"
)

// Sajariv4beta1Pipeline1 struct for Sajariv4beta1Pipeline1
type Sajariv4beta1Pipeline1 struct {
	// Output only. Creation time of the pipeline.
	CreateTime *time.Time          `json:"create_time,omitempty"`
	Type       V4beta1PipelineType `json:"type"`
	// The pipeline's name.  Must start with an alphanumeric character followed by one or more alphanumeric, `_`, `-` or `.` characters. Strictly speaking, it must match the regular expression: `^[a-zA-Z0-9][a-zA-Z0-9_\\-\\.]+$`.
	Name string `json:"name"`
	// The pipeline's version.  Must start with an alphanumeric character followed by one or more alphanumeric, `_`, `-` or `.` characters. Strictly speaking, it must match the regular expression: `^[a-zA-Z0-9][a-zA-Z0-9_\\-\\.]+$`.
	Version string `json:"version"`
	// Description of the pipeline.
	Description *string `json:"description,omitempty"`
	// Pre-steps are run before an indexing operation or query request is sent to the search index.
	PreSteps *[]V4beta1Step1 `json:"pre_steps,omitempty"`
	// Post-steps are run after an indexing operation or query request has been sent to the search index.  For indexing operations, the post-steps only run when creating new records. They do not run when updating records.  For querying, the post-steps have access to the result-set. This makes it possible to act on the results before sending them back to the caller.
	PostSteps *[]V4beta1Step1 `json:"post_steps,omitempty"`
}

// NewSajariv4beta1Pipeline1 instantiates a new Sajariv4beta1Pipeline1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSajariv4beta1Pipeline1(type_ V4beta1PipelineType, name string, version string) *Sajariv4beta1Pipeline1 {
	this := Sajariv4beta1Pipeline1{}
	this.Type = type_
	this.Name = name
	this.Version = version
	return &this
}

// NewSajariv4beta1Pipeline1WithDefaults instantiates a new Sajariv4beta1Pipeline1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSajariv4beta1Pipeline1WithDefaults() *Sajariv4beta1Pipeline1 {
	this := Sajariv4beta1Pipeline1{}
	var type_ V4beta1PipelineType = "TYPE_UNSPECIFIED"
	this.Type = type_
	return &this
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *Sajariv4beta1Pipeline1) GetCreateTime() time.Time {
	if o == nil || o.CreateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sajariv4beta1Pipeline1) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || o.CreateTime == nil {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *Sajariv4beta1Pipeline1) HasCreateTime() bool {
	if o != nil && o.CreateTime != nil {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *Sajariv4beta1Pipeline1) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// GetType returns the Type field value
func (o *Sajariv4beta1Pipeline1) GetType() V4beta1PipelineType {
	if o == nil {
		var ret V4beta1PipelineType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Sajariv4beta1Pipeline1) GetTypeOk() (*V4beta1PipelineType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Sajariv4beta1Pipeline1) SetType(v V4beta1PipelineType) {
	o.Type = v
}

// GetName returns the Name field value
func (o *Sajariv4beta1Pipeline1) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Sajariv4beta1Pipeline1) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Sajariv4beta1Pipeline1) SetName(v string) {
	o.Name = v
}

// GetVersion returns the Version field value
func (o *Sajariv4beta1Pipeline1) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *Sajariv4beta1Pipeline1) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *Sajariv4beta1Pipeline1) SetVersion(v string) {
	o.Version = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Sajariv4beta1Pipeline1) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sajariv4beta1Pipeline1) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Sajariv4beta1Pipeline1) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Sajariv4beta1Pipeline1) SetDescription(v string) {
	o.Description = &v
}

// GetPreSteps returns the PreSteps field value if set, zero value otherwise.
func (o *Sajariv4beta1Pipeline1) GetPreSteps() []V4beta1Step1 {
	if o == nil || o.PreSteps == nil {
		var ret []V4beta1Step1
		return ret
	}
	return *o.PreSteps
}

// GetPreStepsOk returns a tuple with the PreSteps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sajariv4beta1Pipeline1) GetPreStepsOk() (*[]V4beta1Step1, bool) {
	if o == nil || o.PreSteps == nil {
		return nil, false
	}
	return o.PreSteps, true
}

// HasPreSteps returns a boolean if a field has been set.
func (o *Sajariv4beta1Pipeline1) HasPreSteps() bool {
	if o != nil && o.PreSteps != nil {
		return true
	}

	return false
}

// SetPreSteps gets a reference to the given []V4beta1Step1 and assigns it to the PreSteps field.
func (o *Sajariv4beta1Pipeline1) SetPreSteps(v []V4beta1Step1) {
	o.PreSteps = &v
}

// GetPostSteps returns the PostSteps field value if set, zero value otherwise.
func (o *Sajariv4beta1Pipeline1) GetPostSteps() []V4beta1Step1 {
	if o == nil || o.PostSteps == nil {
		var ret []V4beta1Step1
		return ret
	}
	return *o.PostSteps
}

// GetPostStepsOk returns a tuple with the PostSteps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sajariv4beta1Pipeline1) GetPostStepsOk() (*[]V4beta1Step1, bool) {
	if o == nil || o.PostSteps == nil {
		return nil, false
	}
	return o.PostSteps, true
}

// HasPostSteps returns a boolean if a field has been set.
func (o *Sajariv4beta1Pipeline1) HasPostSteps() bool {
	if o != nil && o.PostSteps != nil {
		return true
	}

	return false
}

// SetPostSteps gets a reference to the given []V4beta1Step1 and assigns it to the PostSteps field.
func (o *Sajariv4beta1Pipeline1) SetPostSteps(v []V4beta1Step1) {
	o.PostSteps = &v
}

func (o Sajariv4beta1Pipeline1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreateTime != nil {
		toSerialize["create_time"] = o.CreateTime
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.PreSteps != nil {
		toSerialize["pre_steps"] = o.PreSteps
	}
	if o.PostSteps != nil {
		toSerialize["post_steps"] = o.PostSteps
	}
	return json.Marshal(toSerialize)
}

type NullableSajariv4beta1Pipeline1 struct {
	value *Sajariv4beta1Pipeline1
	isSet bool
}

func (v NullableSajariv4beta1Pipeline1) Get() *Sajariv4beta1Pipeline1 {
	return v.value
}

func (v *NullableSajariv4beta1Pipeline1) Set(val *Sajariv4beta1Pipeline1) {
	v.value = val
	v.isSet = true
}

func (v NullableSajariv4beta1Pipeline1) IsSet() bool {
	return v.isSet
}

func (v *NullableSajariv4beta1Pipeline1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSajariv4beta1Pipeline1(val *Sajariv4beta1Pipeline1) *NullableSajariv4beta1Pipeline1 {
	return &NullableSajariv4beta1Pipeline1{value: val, isSet: true}
}

func (v NullableSajariv4beta1Pipeline1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSajariv4beta1Pipeline1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
