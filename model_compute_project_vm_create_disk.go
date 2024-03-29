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

// ComputeProjectVmCreateDisk struct for ComputeProjectVmCreateDisk
type ComputeProjectVmCreateDisk struct {
	Name string `json:"name"`
	Service string `json:"service"`
	Size float32 `json:"size"`
}

// NewComputeProjectVmCreateDisk instantiates a new ComputeProjectVmCreateDisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeProjectVmCreateDisk(name string, service string, size float32) *ComputeProjectVmCreateDisk {
	this := ComputeProjectVmCreateDisk{}
	this.Name = name
	this.Service = service
	this.Size = size
	return &this
}

// NewComputeProjectVmCreateDiskWithDefaults instantiates a new ComputeProjectVmCreateDisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeProjectVmCreateDiskWithDefaults() *ComputeProjectVmCreateDisk {
	this := ComputeProjectVmCreateDisk{}
	return &this
}

// GetName returns the Name field value
func (o *ComputeProjectVmCreateDisk) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ComputeProjectVmCreateDisk) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ComputeProjectVmCreateDisk) SetName(v string) {
	o.Name = v
}

// GetService returns the Service field value
func (o *ComputeProjectVmCreateDisk) GetService() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *ComputeProjectVmCreateDisk) GetServiceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *ComputeProjectVmCreateDisk) SetService(v string) {
	o.Service = v
}

// GetSize returns the Size field value
func (o *ComputeProjectVmCreateDisk) GetSize() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ComputeProjectVmCreateDisk) GetSizeOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ComputeProjectVmCreateDisk) SetSize(v float32) {
	o.Size = v
}

func (o ComputeProjectVmCreateDisk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["service"] = o.Service
	}
	if true {
		toSerialize["size"] = o.Size
	}
	return json.Marshal(toSerialize)
}

type NullableComputeProjectVmCreateDisk struct {
	value *ComputeProjectVmCreateDisk
	isSet bool
}

func (v NullableComputeProjectVmCreateDisk) Get() *ComputeProjectVmCreateDisk {
	return v.value
}

func (v *NullableComputeProjectVmCreateDisk) Set(val *ComputeProjectVmCreateDisk) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeProjectVmCreateDisk) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeProjectVmCreateDisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeProjectVmCreateDisk(val *ComputeProjectVmCreateDisk) *NullableComputeProjectVmCreateDisk {
	return &NullableComputeProjectVmCreateDisk{value: val, isSet: true}
}

func (v NullableComputeProjectVmCreateDisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeProjectVmCreateDisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


