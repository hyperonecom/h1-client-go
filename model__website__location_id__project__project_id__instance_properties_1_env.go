/*
HyperOne

HyperOne API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// WebsiteLocationIdProjectProjectIdInstanceProperties1Env struct for WebsiteLocationIdProjectProjectIdInstanceProperties1Env
type WebsiteLocationIdProjectProjectIdInstanceProperties1Env struct {
	Value string `json:"value"`
}

// NewWebsiteLocationIdProjectProjectIdInstanceProperties1Env instantiates a new WebsiteLocationIdProjectProjectIdInstanceProperties1Env object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebsiteLocationIdProjectProjectIdInstanceProperties1Env(value string) *WebsiteLocationIdProjectProjectIdInstanceProperties1Env {
	this := WebsiteLocationIdProjectProjectIdInstanceProperties1Env{}
	this.Value = value
	return &this
}

// NewWebsiteLocationIdProjectProjectIdInstanceProperties1EnvWithDefaults instantiates a new WebsiteLocationIdProjectProjectIdInstanceProperties1Env object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebsiteLocationIdProjectProjectIdInstanceProperties1EnvWithDefaults() *WebsiteLocationIdProjectProjectIdInstanceProperties1Env {
	this := WebsiteLocationIdProjectProjectIdInstanceProperties1Env{}
	return &this
}

// GetValue returns the Value field value
func (o *WebsiteLocationIdProjectProjectIdInstanceProperties1Env) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *WebsiteLocationIdProjectProjectIdInstanceProperties1Env) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *WebsiteLocationIdProjectProjectIdInstanceProperties1Env) SetValue(v string) {
	o.Value = v
}

func (o WebsiteLocationIdProjectProjectIdInstanceProperties1Env) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableWebsiteLocationIdProjectProjectIdInstanceProperties1Env struct {
	value *WebsiteLocationIdProjectProjectIdInstanceProperties1Env
	isSet bool
}

func (v NullableWebsiteLocationIdProjectProjectIdInstanceProperties1Env) Get() *WebsiteLocationIdProjectProjectIdInstanceProperties1Env {
	return v.value
}

func (v *NullableWebsiteLocationIdProjectProjectIdInstanceProperties1Env) Set(val *WebsiteLocationIdProjectProjectIdInstanceProperties1Env) {
	v.value = val
	v.isSet = true
}

func (v NullableWebsiteLocationIdProjectProjectIdInstanceProperties1Env) IsSet() bool {
	return v.isSet
}

func (v *NullableWebsiteLocationIdProjectProjectIdInstanceProperties1Env) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebsiteLocationIdProjectProjectIdInstanceProperties1Env(val *WebsiteLocationIdProjectProjectIdInstanceProperties1Env) *NullableWebsiteLocationIdProjectProjectIdInstanceProperties1Env {
	return &NullableWebsiteLocationIdProjectProjectIdInstanceProperties1Env{value: val, isSet: true}
}

func (v NullableWebsiteLocationIdProjectProjectIdInstanceProperties1Env) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebsiteLocationIdProjectProjectIdInstanceProperties1Env) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


