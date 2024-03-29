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

// NetworkingProjectNetadpUpdate struct for NetworkingProjectNetadpUpdate
type NetworkingProjectNetadpUpdate struct {
	Firewall NullableString `json:"firewall,omitempty"`
}

// NewNetworkingProjectNetadpUpdate instantiates a new NetworkingProjectNetadpUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingProjectNetadpUpdate() *NetworkingProjectNetadpUpdate {
	this := NetworkingProjectNetadpUpdate{}
	return &this
}

// NewNetworkingProjectNetadpUpdateWithDefaults instantiates a new NetworkingProjectNetadpUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingProjectNetadpUpdateWithDefaults() *NetworkingProjectNetadpUpdate {
	this := NetworkingProjectNetadpUpdate{}
	return &this
}

// GetFirewall returns the Firewall field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkingProjectNetadpUpdate) GetFirewall() string {
	if o == nil || o.Firewall.Get() == nil {
		var ret string
		return ret
	}
	return *o.Firewall.Get()
}

// GetFirewallOk returns a tuple with the Firewall field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkingProjectNetadpUpdate) GetFirewallOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Firewall.Get(), o.Firewall.IsSet()
}

// HasFirewall returns a boolean if a field has been set.
func (o *NetworkingProjectNetadpUpdate) HasFirewall() bool {
	if o != nil && o.Firewall.IsSet() {
		return true
	}

	return false
}

// SetFirewall gets a reference to the given NullableString and assigns it to the Firewall field.
func (o *NetworkingProjectNetadpUpdate) SetFirewall(v string) {
	o.Firewall.Set(&v)
}
// SetFirewallNil sets the value for Firewall to be an explicit nil
func (o *NetworkingProjectNetadpUpdate) SetFirewallNil() {
	o.Firewall.Set(nil)
}

// UnsetFirewall ensures that no value is present for Firewall, not even an explicit nil
func (o *NetworkingProjectNetadpUpdate) UnsetFirewall() {
	o.Firewall.Unset()
}

func (o NetworkingProjectNetadpUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Firewall.IsSet() {
		toSerialize["firewall"] = o.Firewall.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkingProjectNetadpUpdate struct {
	value *NetworkingProjectNetadpUpdate
	isSet bool
}

func (v NullableNetworkingProjectNetadpUpdate) Get() *NetworkingProjectNetadpUpdate {
	return v.value
}

func (v *NullableNetworkingProjectNetadpUpdate) Set(val *NetworkingProjectNetadpUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingProjectNetadpUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingProjectNetadpUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingProjectNetadpUpdate(val *NetworkingProjectNetadpUpdate) *NullableNetworkingProjectNetadpUpdate {
	return &NullableNetworkingProjectNetadpUpdate{value: val, isSet: true}
}

func (v NullableNetworkingProjectNetadpUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkingProjectNetadpUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


