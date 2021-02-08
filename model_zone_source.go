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

// ZoneSource struct for ZoneSource
type ZoneSource struct {
	DnsProbing *bool `json:"dnsProbing,omitempty"`
}

// NewZoneSource instantiates a new ZoneSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneSource() *ZoneSource {
	this := ZoneSource{}
	var dnsProbing bool = false
	this.DnsProbing = &dnsProbing
	return &this
}

// NewZoneSourceWithDefaults instantiates a new ZoneSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneSourceWithDefaults() *ZoneSource {
	this := ZoneSource{}
	var dnsProbing bool = false
	this.DnsProbing = &dnsProbing
	return &this
}

// GetDnsProbing returns the DnsProbing field value if set, zero value otherwise.
func (o *ZoneSource) GetDnsProbing() bool {
	if o == nil || o.DnsProbing == nil {
		var ret bool
		return ret
	}
	return *o.DnsProbing
}

// GetDnsProbingOk returns a tuple with the DnsProbing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneSource) GetDnsProbingOk() (*bool, bool) {
	if o == nil || o.DnsProbing == nil {
		return nil, false
	}
	return o.DnsProbing, true
}

// HasDnsProbing returns a boolean if a field has been set.
func (o *ZoneSource) HasDnsProbing() bool {
	if o != nil && o.DnsProbing != nil {
		return true
	}

	return false
}

// SetDnsProbing gets a reference to the given bool and assigns it to the DnsProbing field.
func (o *ZoneSource) SetDnsProbing(v bool) {
	o.DnsProbing = &v
}

func (o ZoneSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DnsProbing != nil {
		toSerialize["dnsProbing"] = o.DnsProbing
	}
	return json.Marshal(toSerialize)
}

type NullableZoneSource struct {
	value *ZoneSource
	isSet bool
}

func (v NullableZoneSource) Get() *ZoneSource {
	return v.value
}

func (v *NullableZoneSource) Set(val *ZoneSource) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneSource) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneSource(val *ZoneSource) *NullableZoneSource {
	return &NullableZoneSource{value: val, isSet: true}
}

func (v NullableZoneSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


