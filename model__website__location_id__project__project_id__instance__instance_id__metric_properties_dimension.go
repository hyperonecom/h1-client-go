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

// WebsiteLocationIdProjectProjectIdInstanceInstanceIdMetricPropertiesDimension struct for WebsiteLocationIdProjectProjectIdInstanceInstanceIdMetricPropertiesDimension
type WebsiteLocationIdProjectProjectIdInstanceInstanceIdMetricPropertiesDimension struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewWebsiteLocationIdProjectProjectIdInstanceInstanceIdMetricPropertiesDimension instantiates a new WebsiteLocationIdProjectProjectIdInstanceInstanceIdMetricPropertiesDimension object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebsiteLocationIdProjectProjectIdInstanceInstanceIdMetricPropertiesDimension() *WebsiteLocationIdProjectProjectIdInstanceInstanceIdMetricPropertiesDimension {
	this := WebsiteLocationIdProjectProjectIdInstanceInstanceIdMetricPropertiesDimension{}
	return &this
}

// NewWebsiteLocationIdProjectProjectIdInstanceInstanceIdMetricPropertiesDimensionWithDefaults instantiates a new WebsiteLocationIdProjectProjectIdInstanceInstanceIdMetricPropertiesDimension object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebsiteLocationIdProjectProjectIdInstanceInstanceIdMetricPropertiesDimensionWithDefaults() *WebsiteLocationIdProjectProjectIdInstanceInstanceIdMetricPropertiesDimension {
	this := WebsiteLocationIdProjectProjectIdInstanceInstanceIdMetricPropertiesDimension{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WebsiteLocationIdProjectProjectIdInstanceInstanceIdMetricPropertiesDimension) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteLocationIdProjectProjectIdInstanceInstanceIdMetricPropertiesDimension) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WebsiteLocationIdProjectProjectIdInstanceInstanceIdMetricPropertiesDimension) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WebsiteLocationIdProjectProjectIdInstanceInstanceIdMetricPropertiesDimension) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WebsiteLocationIdProjectProjectIdInstanceInstanceIdMetricPropertiesDimension) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteLocationIdProjectProjectIdInstanceInstanceIdMetricPropertiesDimension) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WebsiteLocationIdProjectProjectIdInstanceInstanceIdMetricPropertiesDimension) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WebsiteLocationIdProjectProjectIdInstanceInstanceIdMetricPropertiesDimension) SetDescription(v string) {
	o.Description = &v
}

func (o WebsiteLocationIdProjectProjectIdInstanceInstanceIdMetricPropertiesDimension) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableWebsiteLocationIdProjectProjectIdInstanceInstanceIdMetricPropertiesDimension struct {
	value *WebsiteLocationIdProjectProjectIdInstanceInstanceIdMetricPropertiesDimension
	isSet bool
}

func (v NullableWebsiteLocationIdProjectProjectIdInstanceInstanceIdMetricPropertiesDimension) Get() *WebsiteLocationIdProjectProjectIdInstanceInstanceIdMetricPropertiesDimension {
	return v.value
}

func (v *NullableWebsiteLocationIdProjectProjectIdInstanceInstanceIdMetricPropertiesDimension) Set(val *WebsiteLocationIdProjectProjectIdInstanceInstanceIdMetricPropertiesDimension) {
	v.value = val
	v.isSet = true
}

func (v NullableWebsiteLocationIdProjectProjectIdInstanceInstanceIdMetricPropertiesDimension) IsSet() bool {
	return v.isSet
}

func (v *NullableWebsiteLocationIdProjectProjectIdInstanceInstanceIdMetricPropertiesDimension) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebsiteLocationIdProjectProjectIdInstanceInstanceIdMetricPropertiesDimension(val *WebsiteLocationIdProjectProjectIdInstanceInstanceIdMetricPropertiesDimension) *NullableWebsiteLocationIdProjectProjectIdInstanceInstanceIdMetricPropertiesDimension {
	return &NullableWebsiteLocationIdProjectProjectIdInstanceInstanceIdMetricPropertiesDimension{value: val, isSet: true}
}

func (v NullableWebsiteLocationIdProjectProjectIdInstanceInstanceIdMetricPropertiesDimension) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebsiteLocationIdProjectProjectIdInstanceInstanceIdMetricPropertiesDimension) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


