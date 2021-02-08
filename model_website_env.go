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
)

// WebsiteEnv struct for WebsiteEnv
type WebsiteEnv struct {
	Id *string `json:"id,omitempty"`
	Name string `json:"name"`
	Value string `json:"value"`
}

// NewWebsiteEnv instantiates a new WebsiteEnv object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebsiteEnv(name string, value string, ) *WebsiteEnv {
	this := WebsiteEnv{}
	this.Name = name
	this.Value = value
	return &this
}

// NewWebsiteEnvWithDefaults instantiates a new WebsiteEnv object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebsiteEnvWithDefaults() *WebsiteEnv {
	this := WebsiteEnv{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WebsiteEnv) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteEnv) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WebsiteEnv) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WebsiteEnv) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *WebsiteEnv) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WebsiteEnv) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WebsiteEnv) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *WebsiteEnv) GetValue() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *WebsiteEnv) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *WebsiteEnv) SetValue(v string) {
	o.Value = v
}

func (o WebsiteEnv) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableWebsiteEnv struct {
	value *WebsiteEnv
	isSet bool
}

func (v NullableWebsiteEnv) Get() *WebsiteEnv {
	return v.value
}

func (v *NullableWebsiteEnv) Set(val *WebsiteEnv) {
	v.value = val
	v.isSet = true
}

func (v NullableWebsiteEnv) IsSet() bool {
	return v.isSet
}

func (v *NullableWebsiteEnv) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebsiteEnv(val *WebsiteEnv) *NullableWebsiteEnv {
	return &NullableWebsiteEnv{value: val, isSet: true}
}

func (v NullableWebsiteEnv) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebsiteEnv) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


