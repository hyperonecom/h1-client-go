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

// Point struct for Point
type Point struct {
	Dimension *map[string]string `json:"dimension,omitempty"`
	Value []PointValue `json:"value,omitempty"`
}

// NewPoint instantiates a new Point object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoint() *Point {
	this := Point{}
	return &this
}

// NewPointWithDefaults instantiates a new Point object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPointWithDefaults() *Point {
	this := Point{}
	return &this
}

// GetDimension returns the Dimension field value if set, zero value otherwise.
func (o *Point) GetDimension() map[string]string {
	if o == nil || o.Dimension == nil {
		var ret map[string]string
		return ret
	}
	return *o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Point) GetDimensionOk() (*map[string]string, bool) {
	if o == nil || o.Dimension == nil {
		return nil, false
	}
	return o.Dimension, true
}

// HasDimension returns a boolean if a field has been set.
func (o *Point) HasDimension() bool {
	if o != nil && o.Dimension != nil {
		return true
	}

	return false
}

// SetDimension gets a reference to the given map[string]string and assigns it to the Dimension field.
func (o *Point) SetDimension(v map[string]string) {
	o.Dimension = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Point) GetValue() []PointValue {
	if o == nil || o.Value == nil {
		var ret []PointValue
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Point) GetValueOk() ([]PointValue, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Point) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []PointValue and assigns it to the Value field.
func (o *Point) SetValue(v []PointValue) {
	o.Value = v
}

func (o Point) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Dimension != nil {
		toSerialize["dimension"] = o.Dimension
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullablePoint struct {
	value *Point
	isSet bool
}

func (v NullablePoint) Get() *Point {
	return v.value
}

func (v *NullablePoint) Set(val *Point) {
	v.value = val
	v.isSet = true
}

func (v NullablePoint) IsSet() bool {
	return v.isSet
}

func (v *NullablePoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoint(val *Point) *NullablePoint {
	return &NullablePoint{value: val, isSet: true}
}

func (v NullablePoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


