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

// ContainerProjectRegistryCreate struct for ContainerProjectRegistryCreate
type ContainerProjectRegistryCreate struct {
	Name string `json:"name"`
	Service string `json:"service"`
	Tag []Tag `json:"tag,omitempty"`
}

// NewContainerProjectRegistryCreate instantiates a new ContainerProjectRegistryCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerProjectRegistryCreate(name string, service string) *ContainerProjectRegistryCreate {
	this := ContainerProjectRegistryCreate{}
	this.Name = name
	this.Service = service
	return &this
}

// NewContainerProjectRegistryCreateWithDefaults instantiates a new ContainerProjectRegistryCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerProjectRegistryCreateWithDefaults() *ContainerProjectRegistryCreate {
	this := ContainerProjectRegistryCreate{}
	return &this
}

// GetName returns the Name field value
func (o *ContainerProjectRegistryCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContainerProjectRegistryCreate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContainerProjectRegistryCreate) SetName(v string) {
	o.Name = v
}

// GetService returns the Service field value
func (o *ContainerProjectRegistryCreate) GetService() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *ContainerProjectRegistryCreate) GetServiceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *ContainerProjectRegistryCreate) SetService(v string) {
	o.Service = v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *ContainerProjectRegistryCreate) GetTag() []Tag {
	if o == nil || o.Tag == nil {
		var ret []Tag
		return ret
	}
	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerProjectRegistryCreate) GetTagOk() ([]Tag, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *ContainerProjectRegistryCreate) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given []Tag and assigns it to the Tag field.
func (o *ContainerProjectRegistryCreate) SetTag(v []Tag) {
	o.Tag = v
}

func (o ContainerProjectRegistryCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["service"] = o.Service
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	return json.Marshal(toSerialize)
}

type NullableContainerProjectRegistryCreate struct {
	value *ContainerProjectRegistryCreate
	isSet bool
}

func (v NullableContainerProjectRegistryCreate) Get() *ContainerProjectRegistryCreate {
	return v.value
}

func (v *NullableContainerProjectRegistryCreate) Set(val *ContainerProjectRegistryCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerProjectRegistryCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerProjectRegistryCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerProjectRegistryCreate(val *ContainerProjectRegistryCreate) *NullableContainerProjectRegistryCreate {
	return &NullableContainerProjectRegistryCreate{value: val, isSet: true}
}

func (v NullableContainerProjectRegistryCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerProjectRegistryCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


