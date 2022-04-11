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

// ContainerProjectRegistryUpdate struct for ContainerProjectRegistryUpdate
type ContainerProjectRegistryUpdate struct {
	Name *string `json:"name,omitempty"`
}

// NewContainerProjectRegistryUpdate instantiates a new ContainerProjectRegistryUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerProjectRegistryUpdate() *ContainerProjectRegistryUpdate {
	this := ContainerProjectRegistryUpdate{}
	return &this
}

// NewContainerProjectRegistryUpdateWithDefaults instantiates a new ContainerProjectRegistryUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerProjectRegistryUpdateWithDefaults() *ContainerProjectRegistryUpdate {
	this := ContainerProjectRegistryUpdate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ContainerProjectRegistryUpdate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerProjectRegistryUpdate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ContainerProjectRegistryUpdate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ContainerProjectRegistryUpdate) SetName(v string) {
	o.Name = &v
}

func (o ContainerProjectRegistryUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableContainerProjectRegistryUpdate struct {
	value *ContainerProjectRegistryUpdate
	isSet bool
}

func (v NullableContainerProjectRegistryUpdate) Get() *ContainerProjectRegistryUpdate {
	return v.value
}

func (v *NullableContainerProjectRegistryUpdate) Set(val *ContainerProjectRegistryUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerProjectRegistryUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerProjectRegistryUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerProjectRegistryUpdate(val *ContainerProjectRegistryUpdate) *NullableContainerProjectRegistryUpdate {
	return &NullableContainerProjectRegistryUpdate{value: val, isSet: true}
}

func (v NullableContainerProjectRegistryUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerProjectRegistryUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


