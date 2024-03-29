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

// ComputeProjectVmConnectOpen struct for ComputeProjectVmConnectOpen
type ComputeProjectVmConnectOpen struct {
	Protocol *string `json:"protocol,omitempty"`
}

// NewComputeProjectVmConnectOpen instantiates a new ComputeProjectVmConnectOpen object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeProjectVmConnectOpen() *ComputeProjectVmConnectOpen {
	this := ComputeProjectVmConnectOpen{}
	var protocol string = "http"
	this.Protocol = &protocol
	return &this
}

// NewComputeProjectVmConnectOpenWithDefaults instantiates a new ComputeProjectVmConnectOpen object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeProjectVmConnectOpenWithDefaults() *ComputeProjectVmConnectOpen {
	this := ComputeProjectVmConnectOpen{}
	var protocol string = "http"
	this.Protocol = &protocol
	return &this
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *ComputeProjectVmConnectOpen) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeProjectVmConnectOpen) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *ComputeProjectVmConnectOpen) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *ComputeProjectVmConnectOpen) SetProtocol(v string) {
	o.Protocol = &v
}

func (o ComputeProjectVmConnectOpen) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Protocol != nil {
		toSerialize["protocol"] = o.Protocol
	}
	return json.Marshal(toSerialize)
}

type NullableComputeProjectVmConnectOpen struct {
	value *ComputeProjectVmConnectOpen
	isSet bool
}

func (v NullableComputeProjectVmConnectOpen) Get() *ComputeProjectVmConnectOpen {
	return v.value
}

func (v *NullableComputeProjectVmConnectOpen) Set(val *ComputeProjectVmConnectOpen) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeProjectVmConnectOpen) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeProjectVmConnectOpen) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeProjectVmConnectOpen(val *ComputeProjectVmConnectOpen) *NullableComputeProjectVmConnectOpen {
	return &NullableComputeProjectVmConnectOpen{value: val, isSet: true}
}

func (v NullableComputeProjectVmConnectOpen) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeProjectVmConnectOpen) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


