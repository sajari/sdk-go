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
	"time"
)

// Redirect Redirect contains a target that you can redirect users to if their search query matches a certain condition.
type Redirect struct {
	// Output only. The redirect's ID.
	Id *string `json:"id,omitempty"`
	// Output only. The ID of the collection that owns this redirect.
	CollectionId *string `json:"collection_id,omitempty"`
	// Output only. Time the redirect was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Output only. Time the redirect was last updated.
	UpdateTime *time.Time `json:"update_time,omitempty"`
	// Condition expression applied to a search request that determines whether a search is redirected.  For example, to redirect if the user's query is `apples`, set condition to `q = 'apples'`.
	Condition string `json:"condition"`
	// Target to redirect the user to if their query matches `condition`.  For searches performed in a browser, target is usually a URL but it can be any value that your integration can interpret as a redirect.  For example, for URLs that you need to resolve at runtime, target might be a URL template string. For apps, target might be a unique identifier used to send the user to the correct view.
	Target string `json:"target"`
	// If disabled, the redirect will never be triggered.
	Disabled *bool `json:"disabled,omitempty"`
}

// NewRedirect instantiates a new Redirect object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedirect(condition string, target string) *Redirect {
	this := Redirect{}
	this.Condition = condition
	this.Target = target
	return &this
}

// NewRedirectWithDefaults instantiates a new Redirect object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedirectWithDefaults() *Redirect {
	this := Redirect{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Redirect) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Redirect) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Redirect) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Redirect) SetId(v string) {
	o.Id = &v
}

// GetCollectionId returns the CollectionId field value if set, zero value otherwise.
func (o *Redirect) GetCollectionId() string {
	if o == nil || o.CollectionId == nil {
		var ret string
		return ret
	}
	return *o.CollectionId
}

// GetCollectionIdOk returns a tuple with the CollectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Redirect) GetCollectionIdOk() (*string, bool) {
	if o == nil || o.CollectionId == nil {
		return nil, false
	}
	return o.CollectionId, true
}

// HasCollectionId returns a boolean if a field has been set.
func (o *Redirect) HasCollectionId() bool {
	if o != nil && o.CollectionId != nil {
		return true
	}

	return false
}

// SetCollectionId gets a reference to the given string and assigns it to the CollectionId field.
func (o *Redirect) SetCollectionId(v string) {
	o.CollectionId = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *Redirect) GetCreateTime() time.Time {
	if o == nil || o.CreateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Redirect) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || o.CreateTime == nil {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *Redirect) HasCreateTime() bool {
	if o != nil && o.CreateTime != nil {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *Redirect) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *Redirect) GetUpdateTime() time.Time {
	if o == nil || o.UpdateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Redirect) GetUpdateTimeOk() (*time.Time, bool) {
	if o == nil || o.UpdateTime == nil {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *Redirect) HasUpdateTime() bool {
	if o != nil && o.UpdateTime != nil {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given time.Time and assigns it to the UpdateTime field.
func (o *Redirect) SetUpdateTime(v time.Time) {
	o.UpdateTime = &v
}

// GetCondition returns the Condition field value
func (o *Redirect) GetCondition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Condition
}

// GetConditionOk returns a tuple with the Condition field value
// and a boolean to check if the value has been set.
func (o *Redirect) GetConditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Condition, true
}

// SetCondition sets field value
func (o *Redirect) SetCondition(v string) {
	o.Condition = v
}

// GetTarget returns the Target field value
func (o *Redirect) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *Redirect) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *Redirect) SetTarget(v string) {
	o.Target = v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *Redirect) GetDisabled() bool {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Redirect) GetDisabledOk() (*bool, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *Redirect) HasDisabled() bool {
	if o != nil && o.Disabled != nil {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *Redirect) SetDisabled(v bool) {
	o.Disabled = &v
}

func (o Redirect) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CollectionId != nil {
		toSerialize["collection_id"] = o.CollectionId
	}
	if o.CreateTime != nil {
		toSerialize["create_time"] = o.CreateTime
	}
	if o.UpdateTime != nil {
		toSerialize["update_time"] = o.UpdateTime
	}
	if true {
		toSerialize["condition"] = o.Condition
	}
	if true {
		toSerialize["target"] = o.Target
	}
	if o.Disabled != nil {
		toSerialize["disabled"] = o.Disabled
	}
	return json.Marshal(toSerialize)
}

type NullableRedirect struct {
	value *Redirect
	isSet bool
}

func (v NullableRedirect) Get() *Redirect {
	return v.value
}

func (v *NullableRedirect) Set(val *Redirect) {
	v.value = val
	v.isSet = true
}

func (v NullableRedirect) IsSet() bool {
	return v.isSet
}

func (v *NullableRedirect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedirect(val *Redirect) *NullableRedirect {
	return &NullableRedirect{value: val, isSet: true}
}

func (v NullableRedirect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedirect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
