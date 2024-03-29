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

// ComputeProjectVmCreateNetadp struct for ComputeProjectVmCreateNetadp
type ComputeProjectVmCreateNetadp struct {
	Firewall *string `json:"firewall,omitempty"`
	Ip []string `json:"ip,omitempty"`
	Network string `json:"network"`
}

// NewComputeProjectVmCreateNetadp instantiates a new ComputeProjectVmCreateNetadp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeProjectVmCreateNetadp(network string) *ComputeProjectVmCreateNetadp {
	this := ComputeProjectVmCreateNetadp{}
	this.Network = network
	return &this
}

// NewComputeProjectVmCreateNetadpWithDefaults instantiates a new ComputeProjectVmCreateNetadp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeProjectVmCreateNetadpWithDefaults() *ComputeProjectVmCreateNetadp {
	this := ComputeProjectVmCreateNetadp{}
	return &this
}

// GetFirewall returns the Firewall field value if set, zero value otherwise.
func (o *ComputeProjectVmCreateNetadp) GetFirewall() string {
	if o == nil || o.Firewall == nil {
		var ret string
		return ret
	}
	return *o.Firewall
}

// GetFirewallOk returns a tuple with the Firewall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeProjectVmCreateNetadp) GetFirewallOk() (*string, bool) {
	if o == nil || o.Firewall == nil {
		return nil, false
	}
	return o.Firewall, true
}

// HasFirewall returns a boolean if a field has been set.
func (o *ComputeProjectVmCreateNetadp) HasFirewall() bool {
	if o != nil && o.Firewall != nil {
		return true
	}

	return false
}

// SetFirewall gets a reference to the given string and assigns it to the Firewall field.
func (o *ComputeProjectVmCreateNetadp) SetFirewall(v string) {
	o.Firewall = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *ComputeProjectVmCreateNetadp) GetIp() []string {
	if o == nil || o.Ip == nil {
		var ret []string
		return ret
	}
	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeProjectVmCreateNetadp) GetIpOk() ([]string, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *ComputeProjectVmCreateNetadp) HasIp() bool {
	if o != nil && o.Ip != nil {
		return true
	}

	return false
}

// SetIp gets a reference to the given []string and assigns it to the Ip field.
func (o *ComputeProjectVmCreateNetadp) SetIp(v []string) {
	o.Ip = v
}

// GetNetwork returns the Network field value
func (o *ComputeProjectVmCreateNetadp) GetNetwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *ComputeProjectVmCreateNetadp) GetNetworkOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *ComputeProjectVmCreateNetadp) SetNetwork(v string) {
	o.Network = v
}

func (o ComputeProjectVmCreateNetadp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Firewall != nil {
		toSerialize["firewall"] = o.Firewall
	}
	if o.Ip != nil {
		toSerialize["ip"] = o.Ip
	}
	if true {
		toSerialize["network"] = o.Network
	}
	return json.Marshal(toSerialize)
}

type NullableComputeProjectVmCreateNetadp struct {
	value *ComputeProjectVmCreateNetadp
	isSet bool
}

func (v NullableComputeProjectVmCreateNetadp) Get() *ComputeProjectVmCreateNetadp {
	return v.value
}

func (v *NullableComputeProjectVmCreateNetadp) Set(val *ComputeProjectVmCreateNetadp) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeProjectVmCreateNetadp) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeProjectVmCreateNetadp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeProjectVmCreateNetadp(val *ComputeProjectVmCreateNetadp) *NullableComputeProjectVmCreateNetadp {
	return &NullableComputeProjectVmCreateNetadp{value: val, isSet: true}
}

func (v NullableComputeProjectVmCreateNetadp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeProjectVmCreateNetadp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


