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

// NetworkingProjectFirewallTransfer struct for NetworkingProjectFirewallTransfer
type NetworkingProjectFirewallTransfer struct {
	Project string `json:"project"`
}

// NewNetworkingProjectFirewallTransfer instantiates a new NetworkingProjectFirewallTransfer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingProjectFirewallTransfer(project string, ) *NetworkingProjectFirewallTransfer {
	this := NetworkingProjectFirewallTransfer{}
	this.Project = project
	return &this
}

// NewNetworkingProjectFirewallTransferWithDefaults instantiates a new NetworkingProjectFirewallTransfer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingProjectFirewallTransferWithDefaults() *NetworkingProjectFirewallTransfer {
	this := NetworkingProjectFirewallTransfer{}
	return &this
}

// GetProject returns the Project field value
func (o *NetworkingProjectFirewallTransfer) GetProject() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *NetworkingProjectFirewallTransfer) GetProjectOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *NetworkingProjectFirewallTransfer) SetProject(v string) {
	o.Project = v
}

func (o NetworkingProjectFirewallTransfer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["project"] = o.Project
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkingProjectFirewallTransfer struct {
	value *NetworkingProjectFirewallTransfer
	isSet bool
}

func (v NullableNetworkingProjectFirewallTransfer) Get() *NetworkingProjectFirewallTransfer {
	return v.value
}

func (v *NullableNetworkingProjectFirewallTransfer) Set(val *NetworkingProjectFirewallTransfer) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingProjectFirewallTransfer) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingProjectFirewallTransfer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingProjectFirewallTransfer(val *NetworkingProjectFirewallTransfer) *NullableNetworkingProjectFirewallTransfer {
	return &NullableNetworkingProjectFirewallTransfer{value: val, isSet: true}
}

func (v NullableNetworkingProjectFirewallTransfer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkingProjectFirewallTransfer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


