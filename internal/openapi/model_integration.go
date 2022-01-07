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
	"time"
)

// Integration Integration contains integration config, e.g. filters and sort options.
type Integration struct {
	// Output only. The ID of the account that owns the collection, e.g. `1618535966441231024`.
	AccountId *string `json:"account_id,omitempty"`
	// Output only. The ID of the collection that owns this integration, e.g. `my-collection`.
	CollectionId *string `json:"collection_id,omitempty"`
	// The integration's config.
	Config map[string]interface{} `json:"config"`
	// Output only. The time the integration was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// The integration's display name.
	DisplayName string `json:"display_name"`
	// Output only. The integration's ID.
	Id *string `json:"id,omitempty"`
	// Output only. The time the integration was last updated.
	UpdateTime *time.Time `json:"update_time,omitempty"`
}

// NewIntegration instantiates a new Integration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegration(config map[string]interface{}, displayName string) *Integration {
	this := Integration{}
	this.Config = config
	this.DisplayName = displayName
	return &this
}

// NewIntegrationWithDefaults instantiates a new Integration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationWithDefaults() *Integration {
	this := Integration{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *Integration) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Integration) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *Integration) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *Integration) SetAccountId(v string) {
	o.AccountId = &v
}

// GetCollectionId returns the CollectionId field value if set, zero value otherwise.
func (o *Integration) GetCollectionId() string {
	if o == nil || o.CollectionId == nil {
		var ret string
		return ret
	}
	return *o.CollectionId
}

// GetCollectionIdOk returns a tuple with the CollectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Integration) GetCollectionIdOk() (*string, bool) {
	if o == nil || o.CollectionId == nil {
		return nil, false
	}
	return o.CollectionId, true
}

// HasCollectionId returns a boolean if a field has been set.
func (o *Integration) HasCollectionId() bool {
	if o != nil && o.CollectionId != nil {
		return true
	}

	return false
}

// SetCollectionId gets a reference to the given string and assigns it to the CollectionId field.
func (o *Integration) SetCollectionId(v string) {
	o.CollectionId = &v
}

// GetConfig returns the Config field value
func (o *Integration) GetConfig() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *Integration) GetConfigOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *Integration) SetConfig(v map[string]interface{}) {
	o.Config = v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *Integration) GetCreateTime() time.Time {
	if o == nil || o.CreateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Integration) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || o.CreateTime == nil {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *Integration) HasCreateTime() bool {
	if o != nil && o.CreateTime != nil {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *Integration) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// GetDisplayName returns the DisplayName field value
func (o *Integration) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *Integration) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *Integration) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Integration) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Integration) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Integration) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Integration) SetId(v string) {
	o.Id = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *Integration) GetUpdateTime() time.Time {
	if o == nil || o.UpdateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Integration) GetUpdateTimeOk() (*time.Time, bool) {
	if o == nil || o.UpdateTime == nil {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *Integration) HasUpdateTime() bool {
	if o != nil && o.UpdateTime != nil {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given time.Time and assigns it to the UpdateTime field.
func (o *Integration) SetUpdateTime(v time.Time) {
	o.UpdateTime = &v
}

func (o Integration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.CollectionId != nil {
		toSerialize["collection_id"] = o.CollectionId
	}
	if true {
		toSerialize["config"] = o.Config
	}
	if o.CreateTime != nil {
		toSerialize["create_time"] = o.CreateTime
	}
	if true {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.UpdateTime != nil {
		toSerialize["update_time"] = o.UpdateTime
	}
	return json.Marshal(toSerialize)
}

type NullableIntegration struct {
	value *Integration
	isSet bool
}

func (v NullableIntegration) Get() *Integration {
	return v.value
}

func (v *NullableIntegration) Set(val *Integration) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegration(val *Integration) *NullableIntegration {
	return &NullableIntegration{value: val, isSet: true}
}

func (v NullableIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}