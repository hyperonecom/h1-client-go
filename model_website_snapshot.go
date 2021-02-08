/*
 * HyperOne
 *
 * HyperOne API
 *
 * API version: 0.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package h1

import (
	"encoding/json"
	"time"
)

// WebsiteSnapshot struct for WebsiteSnapshot
type WebsiteSnapshot struct {
	Id *string `json:"id,omitempty"`
	Name string `json:"name"`
	Creation *time.Time `json:"creation,omitempty"`
	Used *float32 `json:"used,omitempty"`
}

// NewWebsiteSnapshot instantiates a new WebsiteSnapshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebsiteSnapshot(name string, ) *WebsiteSnapshot {
	this := WebsiteSnapshot{}
	this.Name = name
	return &this
}

// NewWebsiteSnapshotWithDefaults instantiates a new WebsiteSnapshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebsiteSnapshotWithDefaults() *WebsiteSnapshot {
	this := WebsiteSnapshot{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WebsiteSnapshot) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteSnapshot) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WebsiteSnapshot) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WebsiteSnapshot) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *WebsiteSnapshot) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WebsiteSnapshot) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WebsiteSnapshot) SetName(v string) {
	o.Name = v
}

// GetCreation returns the Creation field value if set, zero value otherwise.
func (o *WebsiteSnapshot) GetCreation() time.Time {
	if o == nil || o.Creation == nil {
		var ret time.Time
		return ret
	}
	return *o.Creation
}

// GetCreationOk returns a tuple with the Creation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteSnapshot) GetCreationOk() (*time.Time, bool) {
	if o == nil || o.Creation == nil {
		return nil, false
	}
	return o.Creation, true
}

// HasCreation returns a boolean if a field has been set.
func (o *WebsiteSnapshot) HasCreation() bool {
	if o != nil && o.Creation != nil {
		return true
	}

	return false
}

// SetCreation gets a reference to the given time.Time and assigns it to the Creation field.
func (o *WebsiteSnapshot) SetCreation(v time.Time) {
	o.Creation = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *WebsiteSnapshot) GetUsed() float32 {
	if o == nil || o.Used == nil {
		var ret float32
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteSnapshot) GetUsedOk() (*float32, bool) {
	if o == nil || o.Used == nil {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *WebsiteSnapshot) HasUsed() bool {
	if o != nil && o.Used != nil {
		return true
	}

	return false
}

// SetUsed gets a reference to the given float32 and assigns it to the Used field.
func (o *WebsiteSnapshot) SetUsed(v float32) {
	o.Used = &v
}

func (o WebsiteSnapshot) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Creation != nil {
		toSerialize["creation"] = o.Creation
	}
	if o.Used != nil {
		toSerialize["used"] = o.Used
	}
	return json.Marshal(toSerialize)
}

type NullableWebsiteSnapshot struct {
	value *WebsiteSnapshot
	isSet bool
}

func (v NullableWebsiteSnapshot) Get() *WebsiteSnapshot {
	return v.value
}

func (v *NullableWebsiteSnapshot) Set(val *WebsiteSnapshot) {
	v.value = val
	v.isSet = true
}

func (v NullableWebsiteSnapshot) IsSet() bool {
	return v.isSet
}

func (v *NullableWebsiteSnapshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebsiteSnapshot(val *WebsiteSnapshot) *NullableWebsiteSnapshot {
	return &NullableWebsiteSnapshot{value: val, isSet: true}
}

func (v NullableWebsiteSnapshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebsiteSnapshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


