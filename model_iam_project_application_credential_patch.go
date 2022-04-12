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

// IamProjectApplicationCredentialPatch struct for IamProjectApplicationCredentialPatch
type IamProjectApplicationCredentialPatch struct {
	Name string `json:"name"`
}

// NewIamProjectApplicationCredentialPatch instantiates a new IamProjectApplicationCredentialPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamProjectApplicationCredentialPatch(name string) *IamProjectApplicationCredentialPatch {
	this := IamProjectApplicationCredentialPatch{}
	this.Name = name
	return &this
}

// NewIamProjectApplicationCredentialPatchWithDefaults instantiates a new IamProjectApplicationCredentialPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamProjectApplicationCredentialPatchWithDefaults() *IamProjectApplicationCredentialPatch {
	this := IamProjectApplicationCredentialPatch{}
	return &this
}

// GetName returns the Name field value
func (o *IamProjectApplicationCredentialPatch) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IamProjectApplicationCredentialPatch) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IamProjectApplicationCredentialPatch) SetName(v string) {
	o.Name = v
}

func (o IamProjectApplicationCredentialPatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableIamProjectApplicationCredentialPatch struct {
	value *IamProjectApplicationCredentialPatch
	isSet bool
}

func (v NullableIamProjectApplicationCredentialPatch) Get() *IamProjectApplicationCredentialPatch {
	return v.value
}

func (v *NullableIamProjectApplicationCredentialPatch) Set(val *IamProjectApplicationCredentialPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableIamProjectApplicationCredentialPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableIamProjectApplicationCredentialPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamProjectApplicationCredentialPatch(val *IamProjectApplicationCredentialPatch) *NullableIamProjectApplicationCredentialPatch {
	return &NullableIamProjectApplicationCredentialPatch{value: val, isSet: true}
}

func (v NullableIamProjectApplicationCredentialPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamProjectApplicationCredentialPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

