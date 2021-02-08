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

// IpAssociated struct for IpAssociated
type IpAssociated struct {
	Ip *string `json:"ip,omitempty"`
	Fip *string `json:"fip,omitempty"`
	Netadp *string `json:"netadp,omitempty"`
}

// NewIpAssociated instantiates a new IpAssociated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpAssociated() *IpAssociated {
	this := IpAssociated{}
	return &this
}

// NewIpAssociatedWithDefaults instantiates a new IpAssociated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpAssociatedWithDefaults() *IpAssociated {
	this := IpAssociated{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *IpAssociated) GetIp() string {
	if o == nil || o.Ip == nil {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpAssociated) GetIpOk() (*string, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *IpAssociated) HasIp() bool {
	if o != nil && o.Ip != nil {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *IpAssociated) SetIp(v string) {
	o.Ip = &v
}

// GetFip returns the Fip field value if set, zero value otherwise.
func (o *IpAssociated) GetFip() string {
	if o == nil || o.Fip == nil {
		var ret string
		return ret
	}
	return *o.Fip
}

// GetFipOk returns a tuple with the Fip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpAssociated) GetFipOk() (*string, bool) {
	if o == nil || o.Fip == nil {
		return nil, false
	}
	return o.Fip, true
}

// HasFip returns a boolean if a field has been set.
func (o *IpAssociated) HasFip() bool {
	if o != nil && o.Fip != nil {
		return true
	}

	return false
}

// SetFip gets a reference to the given string and assigns it to the Fip field.
func (o *IpAssociated) SetFip(v string) {
	o.Fip = &v
}

// GetNetadp returns the Netadp field value if set, zero value otherwise.
func (o *IpAssociated) GetNetadp() string {
	if o == nil || o.Netadp == nil {
		var ret string
		return ret
	}
	return *o.Netadp
}

// GetNetadpOk returns a tuple with the Netadp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpAssociated) GetNetadpOk() (*string, bool) {
	if o == nil || o.Netadp == nil {
		return nil, false
	}
	return o.Netadp, true
}

// HasNetadp returns a boolean if a field has been set.
func (o *IpAssociated) HasNetadp() bool {
	if o != nil && o.Netadp != nil {
		return true
	}

	return false
}

// SetNetadp gets a reference to the given string and assigns it to the Netadp field.
func (o *IpAssociated) SetNetadp(v string) {
	o.Netadp = &v
}

func (o IpAssociated) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ip != nil {
		toSerialize["ip"] = o.Ip
	}
	if o.Fip != nil {
		toSerialize["fip"] = o.Fip
	}
	if o.Netadp != nil {
		toSerialize["netadp"] = o.Netadp
	}
	return json.Marshal(toSerialize)
}

type NullableIpAssociated struct {
	value *IpAssociated
	isSet bool
}

func (v NullableIpAssociated) Get() *IpAssociated {
	return v.value
}

func (v *NullableIpAssociated) Set(val *IpAssociated) {
	v.value = val
	v.isSet = true
}

func (v NullableIpAssociated) IsSet() bool {
	return v.isSet
}

func (v *NullableIpAssociated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpAssociated(val *IpAssociated) *NullableIpAssociated {
	return &NullableIpAssociated{value: val, isSet: true}
}

func (v NullableIpAssociated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpAssociated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


