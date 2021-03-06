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

// Point struct for Point
type Point struct {
	Time time.Time `json:"time"`
	Value float32 `json:"value"`
}

// NewPoint instantiates a new Point object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoint(time time.Time, value float32, ) *Point {
	this := Point{}
	this.Time = time
	this.Value = value
	return &this
}

// NewPointWithDefaults instantiates a new Point object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPointWithDefaults() *Point {
	this := Point{}
	return &this
}

// GetTime returns the Time field value
func (o *Point) GetTime() time.Time {
	if o == nil  {
		var ret time.Time
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *Point) GetTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *Point) SetTime(v time.Time) {
	o.Time = v
}

// GetValue returns the Value field value
func (o *Point) GetValue() float32 {
	if o == nil  {
		var ret float32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Point) GetValueOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Point) SetValue(v float32) {
	o.Value = v
}

func (o Point) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["time"] = o.Time
	}
	if true {
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


