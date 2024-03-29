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

// NetworkingProjectIpCreate struct for NetworkingProjectIpCreate
type NetworkingProjectIpCreate struct {
	Address *string `json:"address,omitempty"`
	Network *string `json:"network,omitempty"`
	PtrRecord *string `json:"ptrRecord,omitempty"`
	Tag []Tag `json:"tag,omitempty"`
}

// NewNetworkingProjectIpCreate instantiates a new NetworkingProjectIpCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingProjectIpCreate() *NetworkingProjectIpCreate {
	this := NetworkingProjectIpCreate{}
	return &this
}

// NewNetworkingProjectIpCreateWithDefaults instantiates a new NetworkingProjectIpCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingProjectIpCreateWithDefaults() *NetworkingProjectIpCreate {
	this := NetworkingProjectIpCreate{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *NetworkingProjectIpCreate) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingProjectIpCreate) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *NetworkingProjectIpCreate) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *NetworkingProjectIpCreate) SetAddress(v string) {
	o.Address = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *NetworkingProjectIpCreate) GetNetwork() string {
	if o == nil || o.Network == nil {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingProjectIpCreate) GetNetworkOk() (*string, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *NetworkingProjectIpCreate) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *NetworkingProjectIpCreate) SetNetwork(v string) {
	o.Network = &v
}

// GetPtrRecord returns the PtrRecord field value if set, zero value otherwise.
func (o *NetworkingProjectIpCreate) GetPtrRecord() string {
	if o == nil || o.PtrRecord == nil {
		var ret string
		return ret
	}
	return *o.PtrRecord
}

// GetPtrRecordOk returns a tuple with the PtrRecord field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingProjectIpCreate) GetPtrRecordOk() (*string, bool) {
	if o == nil || o.PtrRecord == nil {
		return nil, false
	}
	return o.PtrRecord, true
}

// HasPtrRecord returns a boolean if a field has been set.
func (o *NetworkingProjectIpCreate) HasPtrRecord() bool {
	if o != nil && o.PtrRecord != nil {
		return true
	}

	return false
}

// SetPtrRecord gets a reference to the given string and assigns it to the PtrRecord field.
func (o *NetworkingProjectIpCreate) SetPtrRecord(v string) {
	o.PtrRecord = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *NetworkingProjectIpCreate) GetTag() []Tag {
	if o == nil || o.Tag == nil {
		var ret []Tag
		return ret
	}
	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingProjectIpCreate) GetTagOk() ([]Tag, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *NetworkingProjectIpCreate) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given []Tag and assigns it to the Tag field.
func (o *NetworkingProjectIpCreate) SetTag(v []Tag) {
	o.Tag = v
}

func (o NetworkingProjectIpCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	if o.PtrRecord != nil {
		toSerialize["ptrRecord"] = o.PtrRecord
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkingProjectIpCreate struct {
	value *NetworkingProjectIpCreate
	isSet bool
}

func (v NullableNetworkingProjectIpCreate) Get() *NetworkingProjectIpCreate {
	return v.value
}

func (v *NullableNetworkingProjectIpCreate) Set(val *NetworkingProjectIpCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingProjectIpCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingProjectIpCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingProjectIpCreate(val *NetworkingProjectIpCreate) *NullableNetworkingProjectIpCreate {
	return &NullableNetworkingProjectIpCreate{value: val, isSet: true}
}

func (v NullableNetworkingProjectIpCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkingProjectIpCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


