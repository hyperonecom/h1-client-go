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

// DnsProjectZoneUpdate struct for DnsProjectZoneUpdate
type DnsProjectZoneUpdate struct {
	Name *string `json:"name,omitempty"`
}

// NewDnsProjectZoneUpdate instantiates a new DnsProjectZoneUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsProjectZoneUpdate() *DnsProjectZoneUpdate {
	this := DnsProjectZoneUpdate{}
	return &this
}

// NewDnsProjectZoneUpdateWithDefaults instantiates a new DnsProjectZoneUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsProjectZoneUpdateWithDefaults() *DnsProjectZoneUpdate {
	this := DnsProjectZoneUpdate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DnsProjectZoneUpdate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsProjectZoneUpdate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DnsProjectZoneUpdate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DnsProjectZoneUpdate) SetName(v string) {
	o.Name = &v
}

func (o DnsProjectZoneUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableDnsProjectZoneUpdate struct {
	value *DnsProjectZoneUpdate
	isSet bool
}

func (v NullableDnsProjectZoneUpdate) Get() *DnsProjectZoneUpdate {
	return v.value
}

func (v *NullableDnsProjectZoneUpdate) Set(val *DnsProjectZoneUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsProjectZoneUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsProjectZoneUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsProjectZoneUpdate(val *DnsProjectZoneUpdate) *NullableDnsProjectZoneUpdate {
	return &NullableDnsProjectZoneUpdate{value: val, isSet: true}
}

func (v NullableDnsProjectZoneUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsProjectZoneUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


