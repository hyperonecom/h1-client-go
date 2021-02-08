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

// NetworkingProjectNetgwUpdate struct for NetworkingProjectNetgwUpdate
type NetworkingProjectNetgwUpdate struct {
	Name *string `json:"name,omitempty"`
}

// NewNetworkingProjectNetgwUpdate instantiates a new NetworkingProjectNetgwUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingProjectNetgwUpdate() *NetworkingProjectNetgwUpdate {
	this := NetworkingProjectNetgwUpdate{}
	return &this
}

// NewNetworkingProjectNetgwUpdateWithDefaults instantiates a new NetworkingProjectNetgwUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingProjectNetgwUpdateWithDefaults() *NetworkingProjectNetgwUpdate {
	this := NetworkingProjectNetgwUpdate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworkingProjectNetgwUpdate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingProjectNetgwUpdate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworkingProjectNetgwUpdate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworkingProjectNetgwUpdate) SetName(v string) {
	o.Name = &v
}

func (o NetworkingProjectNetgwUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkingProjectNetgwUpdate struct {
	value *NetworkingProjectNetgwUpdate
	isSet bool
}

func (v NullableNetworkingProjectNetgwUpdate) Get() *NetworkingProjectNetgwUpdate {
	return v.value
}

func (v *NullableNetworkingProjectNetgwUpdate) Set(val *NetworkingProjectNetgwUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingProjectNetgwUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingProjectNetgwUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingProjectNetgwUpdate(val *NetworkingProjectNetgwUpdate) *NullableNetworkingProjectNetgwUpdate {
	return &NullableNetworkingProjectNetgwUpdate{value: val, isSet: true}
}

func (v NullableNetworkingProjectNetgwUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkingProjectNetgwUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


