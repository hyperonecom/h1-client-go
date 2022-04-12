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

// MetricProperties struct for MetricProperties
type MetricProperties struct {
	Dimension []MetricPropertiesDimension `json:"dimension,omitempty"`
	Unit string `json:"unit"`
}

// NewMetricProperties instantiates a new MetricProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricProperties(unit string) *MetricProperties {
	this := MetricProperties{}
	this.Unit = unit
	return &this
}

// NewMetricPropertiesWithDefaults instantiates a new MetricProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricPropertiesWithDefaults() *MetricProperties {
	this := MetricProperties{}
	return &this
}

// GetDimension returns the Dimension field value if set, zero value otherwise.
func (o *MetricProperties) GetDimension() []MetricPropertiesDimension {
	if o == nil || o.Dimension == nil {
		var ret []MetricPropertiesDimension
		return ret
	}
	return o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricProperties) GetDimensionOk() ([]MetricPropertiesDimension, bool) {
	if o == nil || o.Dimension == nil {
		return nil, false
	}
	return o.Dimension, true
}

// HasDimension returns a boolean if a field has been set.
func (o *MetricProperties) HasDimension() bool {
	if o != nil && o.Dimension != nil {
		return true
	}

	return false
}

// SetDimension gets a reference to the given []MetricPropertiesDimension and assigns it to the Dimension field.
func (o *MetricProperties) SetDimension(v []MetricPropertiesDimension) {
	o.Dimension = v
}

// GetUnit returns the Unit field value
func (o *MetricProperties) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *MetricProperties) GetUnitOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *MetricProperties) SetUnit(v string) {
	o.Unit = v
}

func (o MetricProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Dimension != nil {
		toSerialize["dimension"] = o.Dimension
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableMetricProperties struct {
	value *MetricProperties
	isSet bool
}

func (v NullableMetricProperties) Get() *MetricProperties {
	return v.value
}

func (v *NullableMetricProperties) Set(val *MetricProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricProperties(val *MetricProperties) *NullableMetricProperties {
	return &NullableMetricProperties{value: val, isSet: true}
}

func (v NullableMetricProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

