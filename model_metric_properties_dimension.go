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

// MetricPropertiesDimension struct for MetricPropertiesDimension
type MetricPropertiesDimension struct {
	Description string `json:"description"`
	Name string `json:"name"`
}

// NewMetricPropertiesDimension instantiates a new MetricPropertiesDimension object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricPropertiesDimension(description string, name string) *MetricPropertiesDimension {
	this := MetricPropertiesDimension{}
	this.Description = description
	this.Name = name
	return &this
}

// NewMetricPropertiesDimensionWithDefaults instantiates a new MetricPropertiesDimension object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricPropertiesDimensionWithDefaults() *MetricPropertiesDimension {
	this := MetricPropertiesDimension{}
	return &this
}

// GetDescription returns the Description field value
func (o *MetricPropertiesDimension) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *MetricPropertiesDimension) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *MetricPropertiesDimension) SetDescription(v string) {
	o.Description = v
}

// GetName returns the Name field value
func (o *MetricPropertiesDimension) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MetricPropertiesDimension) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MetricPropertiesDimension) SetName(v string) {
	o.Name = v
}

func (o MetricPropertiesDimension) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableMetricPropertiesDimension struct {
	value *MetricPropertiesDimension
	isSet bool
}

func (v NullableMetricPropertiesDimension) Get() *MetricPropertiesDimension {
	return v.value
}

func (v *NullableMetricPropertiesDimension) Set(val *MetricPropertiesDimension) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricPropertiesDimension) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricPropertiesDimension) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricPropertiesDimension(val *MetricPropertiesDimension) *NullableMetricPropertiesDimension {
	return &NullableMetricPropertiesDimension{value: val, isSet: true}
}

func (v NullableMetricPropertiesDimension) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricPropertiesDimension) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


