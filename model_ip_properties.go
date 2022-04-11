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

// IpProperties struct for IpProperties
type IpProperties struct {
	Family *string `json:"family,omitempty"`
	Version *float32 `json:"version,omitempty"`
}

// NewIpProperties instantiates a new IpProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpProperties() *IpProperties {
	this := IpProperties{}
	return &this
}

// NewIpPropertiesWithDefaults instantiates a new IpProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpPropertiesWithDefaults() *IpProperties {
	this := IpProperties{}
	return &this
}

// GetFamily returns the Family field value if set, zero value otherwise.
func (o *IpProperties) GetFamily() string {
	if o == nil || o.Family == nil {
		var ret string
		return ret
	}
	return *o.Family
}

// GetFamilyOk returns a tuple with the Family field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpProperties) GetFamilyOk() (*string, bool) {
	if o == nil || o.Family == nil {
		return nil, false
	}
	return o.Family, true
}

// HasFamily returns a boolean if a field has been set.
func (o *IpProperties) HasFamily() bool {
	if o != nil && o.Family != nil {
		return true
	}

	return false
}

// SetFamily gets a reference to the given string and assigns it to the Family field.
func (o *IpProperties) SetFamily(v string) {
	o.Family = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *IpProperties) GetVersion() float32 {
	if o == nil || o.Version == nil {
		var ret float32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpProperties) GetVersionOk() (*float32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *IpProperties) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given float32 and assigns it to the Version field.
func (o *IpProperties) SetVersion(v float32) {
	o.Version = &v
}

func (o IpProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Family != nil {
		toSerialize["family"] = o.Family
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableIpProperties struct {
	value *IpProperties
	isSet bool
}

func (v NullableIpProperties) Get() *IpProperties {
	return v.value
}

func (v *NullableIpProperties) Set(val *IpProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableIpProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableIpProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpProperties(val *IpProperties) *NullableIpProperties {
	return &NullableIpProperties{value: val, isSet: true}
}

func (v NullableIpProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


