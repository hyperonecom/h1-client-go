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

// ComputeProjectVmSerialport struct for ComputeProjectVmSerialport
type ComputeProjectVmSerialport struct {
	Number *string `json:"number,omitempty"`
}

// NewComputeProjectVmSerialport instantiates a new ComputeProjectVmSerialport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeProjectVmSerialport() *ComputeProjectVmSerialport {
	this := ComputeProjectVmSerialport{}
	var number string = "1"
	this.Number = &number
	return &this
}

// NewComputeProjectVmSerialportWithDefaults instantiates a new ComputeProjectVmSerialport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeProjectVmSerialportWithDefaults() *ComputeProjectVmSerialport {
	this := ComputeProjectVmSerialport{}
	var number string = "1"
	this.Number = &number
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *ComputeProjectVmSerialport) GetNumber() string {
	if o == nil || o.Number == nil {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeProjectVmSerialport) GetNumberOk() (*string, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *ComputeProjectVmSerialport) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *ComputeProjectVmSerialport) SetNumber(v string) {
	o.Number = &v
}

func (o ComputeProjectVmSerialport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	return json.Marshal(toSerialize)
}

type NullableComputeProjectVmSerialport struct {
	value *ComputeProjectVmSerialport
	isSet bool
}

func (v NullableComputeProjectVmSerialport) Get() *ComputeProjectVmSerialport {
	return v.value
}

func (v *NullableComputeProjectVmSerialport) Set(val *ComputeProjectVmSerialport) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeProjectVmSerialport) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeProjectVmSerialport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeProjectVmSerialport(val *ComputeProjectVmSerialport) *NullableComputeProjectVmSerialport {
	return &NullableComputeProjectVmSerialport{value: val, isSet: true}
}

func (v NullableComputeProjectVmSerialport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeProjectVmSerialport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


