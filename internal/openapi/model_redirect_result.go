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

// RedirectResult RedirectResult indicates that a redirect has been triggered for a given query.
type RedirectResult struct {
	// The redirect's ID.
	Id *string `json:"id,omitempty"`
	// The target to redirect the user to.
	Target *string `json:"target,omitempty"`
	// A redirect token.  Call SendEvent with this token to indicate that a redirect has been performed.
	Token *string `json:"token,omitempty"`
}

// NewRedirectResult instantiates a new RedirectResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedirectResult() *RedirectResult {
	this := RedirectResult{}
	return &this
}

// NewRedirectResultWithDefaults instantiates a new RedirectResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedirectResultWithDefaults() *RedirectResult {
	this := RedirectResult{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RedirectResult) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedirectResult) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RedirectResult) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RedirectResult) SetId(v string) {
	o.Id = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *RedirectResult) GetTarget() string {
	if o == nil || o.Target == nil {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedirectResult) GetTargetOk() (*string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *RedirectResult) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *RedirectResult) SetTarget(v string) {
	o.Target = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *RedirectResult) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedirectResult) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *RedirectResult) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *RedirectResult) SetToken(v string) {
	o.Token = &v
}

func (o RedirectResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableRedirectResult struct {
	value *RedirectResult
	isSet bool
}

func (v NullableRedirectResult) Get() *RedirectResult {
	return v.value
}

func (v *NullableRedirectResult) Set(val *RedirectResult) {
	v.value = val
	v.isSet = true
}

func (v NullableRedirectResult) IsSet() bool {
	return v.isSet
}

func (v *NullableRedirectResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedirectResult(val *RedirectResult) *NullableRedirectResult {
	return &NullableRedirectResult{value: val, isSet: true}
}

func (v NullableRedirectResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedirectResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
