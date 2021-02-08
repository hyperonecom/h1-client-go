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

// ComputeProjectVmIsoCreate struct for ComputeProjectVmIsoCreate
type ComputeProjectVmIsoCreate struct {
	Iso *string `json:"iso,omitempty"`
}

// NewComputeProjectVmIsoCreate instantiates a new ComputeProjectVmIsoCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeProjectVmIsoCreate() *ComputeProjectVmIsoCreate {
	this := ComputeProjectVmIsoCreate{}
	return &this
}

// NewComputeProjectVmIsoCreateWithDefaults instantiates a new ComputeProjectVmIsoCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeProjectVmIsoCreateWithDefaults() *ComputeProjectVmIsoCreate {
	this := ComputeProjectVmIsoCreate{}
	return &this
}

// GetIso returns the Iso field value if set, zero value otherwise.
func (o *ComputeProjectVmIsoCreate) GetIso() string {
	if o == nil || o.Iso == nil {
		var ret string
		return ret
	}
	return *o.Iso
}

// GetIsoOk returns a tuple with the Iso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeProjectVmIsoCreate) GetIsoOk() (*string, bool) {
	if o == nil || o.Iso == nil {
		return nil, false
	}
	return o.Iso, true
}

// HasIso returns a boolean if a field has been set.
func (o *ComputeProjectVmIsoCreate) HasIso() bool {
	if o != nil && o.Iso != nil {
		return true
	}

	return false
}

// SetIso gets a reference to the given string and assigns it to the Iso field.
func (o *ComputeProjectVmIsoCreate) SetIso(v string) {
	o.Iso = &v
}

func (o ComputeProjectVmIsoCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Iso != nil {
		toSerialize["iso"] = o.Iso
	}
	return json.Marshal(toSerialize)
}

type NullableComputeProjectVmIsoCreate struct {
	value *ComputeProjectVmIsoCreate
	isSet bool
}

func (v NullableComputeProjectVmIsoCreate) Get() *ComputeProjectVmIsoCreate {
	return v.value
}

func (v *NullableComputeProjectVmIsoCreate) Set(val *ComputeProjectVmIsoCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeProjectVmIsoCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeProjectVmIsoCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeProjectVmIsoCreate(val *ComputeProjectVmIsoCreate) *NullableComputeProjectVmIsoCreate {
	return &NullableComputeProjectVmIsoCreate{value: val, isSet: true}
}

func (v NullableComputeProjectVmIsoCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeProjectVmIsoCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


