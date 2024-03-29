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

// NetworkingProjectIpAssociate struct for NetworkingProjectIpAssociate
type NetworkingProjectIpAssociate struct {
	Ip string `json:"ip"`
}

// NewNetworkingProjectIpAssociate instantiates a new NetworkingProjectIpAssociate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingProjectIpAssociate(ip string) *NetworkingProjectIpAssociate {
	this := NetworkingProjectIpAssociate{}
	this.Ip = ip
	return &this
}

// NewNetworkingProjectIpAssociateWithDefaults instantiates a new NetworkingProjectIpAssociate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingProjectIpAssociateWithDefaults() *NetworkingProjectIpAssociate {
	this := NetworkingProjectIpAssociate{}
	return &this
}

// GetIp returns the Ip field value
func (o *NetworkingProjectIpAssociate) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *NetworkingProjectIpAssociate) GetIpOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *NetworkingProjectIpAssociate) SetIp(v string) {
	o.Ip = v
}

func (o NetworkingProjectIpAssociate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ip"] = o.Ip
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkingProjectIpAssociate struct {
	value *NetworkingProjectIpAssociate
	isSet bool
}

func (v NullableNetworkingProjectIpAssociate) Get() *NetworkingProjectIpAssociate {
	return v.value
}

func (v *NullableNetworkingProjectIpAssociate) Set(val *NetworkingProjectIpAssociate) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingProjectIpAssociate) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingProjectIpAssociate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingProjectIpAssociate(val *NetworkingProjectIpAssociate) *NullableNetworkingProjectIpAssociate {
	return &NullableNetworkingProjectIpAssociate{value: val, isSet: true}
}

func (v NullableNetworkingProjectIpAssociate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkingProjectIpAssociate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


