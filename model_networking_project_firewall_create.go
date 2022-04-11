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

// NetworkingProjectFirewallCreate struct for NetworkingProjectFirewallCreate
type NetworkingProjectFirewallCreate struct {
	Name string `json:"name"`
	Service *string `json:"service,omitempty"`
	Ingress []NetworkingRule `json:"ingress,omitempty"`
	Egress []NetworkingRule `json:"egress,omitempty"`
	Tag []Tag `json:"tag,omitempty"`
}

// NewNetworkingProjectFirewallCreate instantiates a new NetworkingProjectFirewallCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingProjectFirewallCreate(name string) *NetworkingProjectFirewallCreate {
	this := NetworkingProjectFirewallCreate{}
	this.Name = name
	var service string = "5bacaf7202deee0c100eda3b"
	this.Service = &service
	return &this
}

// NewNetworkingProjectFirewallCreateWithDefaults instantiates a new NetworkingProjectFirewallCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingProjectFirewallCreateWithDefaults() *NetworkingProjectFirewallCreate {
	this := NetworkingProjectFirewallCreate{}
	var service string = "5bacaf7202deee0c100eda3b"
	this.Service = &service
	return &this
}

// GetName returns the Name field value
func (o *NetworkingProjectFirewallCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NetworkingProjectFirewallCreate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NetworkingProjectFirewallCreate) SetName(v string) {
	o.Name = v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *NetworkingProjectFirewallCreate) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingProjectFirewallCreate) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *NetworkingProjectFirewallCreate) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *NetworkingProjectFirewallCreate) SetService(v string) {
	o.Service = &v
}

// GetIngress returns the Ingress field value if set, zero value otherwise.
func (o *NetworkingProjectFirewallCreate) GetIngress() []NetworkingRule {
	if o == nil || o.Ingress == nil {
		var ret []NetworkingRule
		return ret
	}
	return o.Ingress
}

// GetIngressOk returns a tuple with the Ingress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingProjectFirewallCreate) GetIngressOk() ([]NetworkingRule, bool) {
	if o == nil || o.Ingress == nil {
		return nil, false
	}
	return o.Ingress, true
}

// HasIngress returns a boolean if a field has been set.
func (o *NetworkingProjectFirewallCreate) HasIngress() bool {
	if o != nil && o.Ingress != nil {
		return true
	}

	return false
}

// SetIngress gets a reference to the given []NetworkingRule and assigns it to the Ingress field.
func (o *NetworkingProjectFirewallCreate) SetIngress(v []NetworkingRule) {
	o.Ingress = v
}

// GetEgress returns the Egress field value if set, zero value otherwise.
func (o *NetworkingProjectFirewallCreate) GetEgress() []NetworkingRule {
	if o == nil || o.Egress == nil {
		var ret []NetworkingRule
		return ret
	}
	return o.Egress
}

// GetEgressOk returns a tuple with the Egress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingProjectFirewallCreate) GetEgressOk() ([]NetworkingRule, bool) {
	if o == nil || o.Egress == nil {
		return nil, false
	}
	return o.Egress, true
}

// HasEgress returns a boolean if a field has been set.
func (o *NetworkingProjectFirewallCreate) HasEgress() bool {
	if o != nil && o.Egress != nil {
		return true
	}

	return false
}

// SetEgress gets a reference to the given []NetworkingRule and assigns it to the Egress field.
func (o *NetworkingProjectFirewallCreate) SetEgress(v []NetworkingRule) {
	o.Egress = v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *NetworkingProjectFirewallCreate) GetTag() []Tag {
	if o == nil || o.Tag == nil {
		var ret []Tag
		return ret
	}
	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingProjectFirewallCreate) GetTagOk() ([]Tag, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *NetworkingProjectFirewallCreate) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given []Tag and assigns it to the Tag field.
func (o *NetworkingProjectFirewallCreate) SetTag(v []Tag) {
	o.Tag = v
}

func (o NetworkingProjectFirewallCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.Ingress != nil {
		toSerialize["ingress"] = o.Ingress
	}
	if o.Egress != nil {
		toSerialize["egress"] = o.Egress
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkingProjectFirewallCreate struct {
	value *NetworkingProjectFirewallCreate
	isSet bool
}

func (v NullableNetworkingProjectFirewallCreate) Get() *NetworkingProjectFirewallCreate {
	return v.value
}

func (v *NullableNetworkingProjectFirewallCreate) Set(val *NetworkingProjectFirewallCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingProjectFirewallCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingProjectFirewallCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingProjectFirewallCreate(val *NetworkingProjectFirewallCreate) *NullableNetworkingProjectFirewallCreate {
	return &NullableNetworkingProjectFirewallCreate{value: val, isSet: true}
}

func (v NullableNetworkingProjectFirewallCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkingProjectFirewallCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


